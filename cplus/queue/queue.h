#pragma once

#include <limits.h>
#include <stdint.h>
#include <inttypes.h>
#include <assert.h>

#include <chrono>
#include <condition_variable>
#include <mutex>
#include <map>
#include <utility>
#include <string>
#include <sys/time.h>
#include <time.h>
#include <iostream>

using namespace std;

namespace Stream
{
    static inline uint64_t getNowMs()
    {
        timeval tv;
        gettimeofday(&tv, nullptr);

        return tv.tv_sec * 1000 + tv.tv_usec / 1000;
    }

    constexpr size_t kDefaultMaxSize = 600 * 6;
    constexpr size_t kMaxRollbackCount = 60 * 1;
    constexpr int DEFAULT_QUEUE_TIMEOUT_MS = 5;

    template <typename VAL>
    class Queue
    {
    public:
        Queue()
            : m_name("")
            , m_prePushKey(INT64_MAX)
            , m_prePopTimeMs(0)
            , m_preTimeRef(0)
            , m_maxSize(kDefaultMaxSize)
            , m_pushRollBackCount(0)
        {
        }

        Queue(size_t max)
            : m_name("")
            , m_prePushKey(INT64_MAX)
            , m_prePopTimeMs(0)
            , m_preTimeRef(0)
            , m_maxSize(max)
            , m_pushRollBackCount(0)
        {
        }

        ~Queue()
        {
        }


        void setName(const std::string &name)
        {
            m_name = name;
        }

        std::string getName() const
        {
            return m_name;
        }

        void setMaxSize(size_t max)
        {
            m_maxSize = max;
        }

        size_t size()
        {
            std::unique_lock<std::mutex> lockGuard(m_mutex);
            return m_queue.size();
        }

        bool empty()
        {
            std::unique_lock<std::mutex> lockGuard(m_mutex);
            return m_queue.empty();
        }

        bool push(int64_t key, const VAL &val)
        {
            std::unique_lock<std::mutex> lockGuard(m_mutex);
            m_maxKey = m_maxKey > key ? m_maxKey : key;

            bool forceClean = false;

            if (m_prePushKey != INT64_MAX)
            {
                // 针对audio/video场景：
                // 1. key值小于先前值，push失败；
                // 2. key比先前值一直小，次数达到kMaxRollbackCount则重置
                if (key < m_prePushKey)
                {
                    cout << "key rollback: " << m_prePushKey << " -> " << key << endl;

                    ++m_pushRollBackCount;

                    if (m_pushRollBackCount == kMaxRollbackCount)
                    {
                        m_pushRollBackCount = 0;
                        forceClean = true;

                        cout << "rollback count -> " << kMaxRollbackCount << ", force clean" << endl;
                    }
                    else
                    {
                        return false;
                    }
                }
            }

            if (forceClean)
            {
                // 次数达到kMaxRollbackCount则认为翻转、重置
                m_prePushKey = INT64_MAX;
                m_prePopTimeMs = 0;
                m_preTimeRef = 0;

                cout << "rollback" << endl;
            }

            m_prePushKey = key;

            m_queue.insert({key, val});

            while (m_queue.size() >= m_maxSize)
            {
                cout << "queue size too big: " << m_queue.size() << endl;
                m_queue.erase(m_queue.begin());
            }

            m_cond.notify_one();

            return true;
        }

        bool pop(VAL &val, int timeoutInMs = 0)
        {
            std::unique_lock<std::mutex> lockGuard(m_mutex);

            if (m_queue.empty())
            {
                if (timeoutInMs == 0)
                {
                    return false;
                }
                else if (timeoutInMs == -1)
                {
                    m_cond.wait(lockGuard);
                }
                else
                {
                    if (m_cond.wait_for(lockGuard, std::chrono::microseconds(timeoutInMs))
                         == std::cv_status::timeout)
                    {
                        return false;
                    }
                }
            }

            if (m_queue.empty())
            {
                return false;
            }

            m_prePopTimeMs = getNowMs();
            val = m_queue.begin()->second;

            assert(!m_queue.empty());

            m_queue.erase(m_queue.begin());

            return true;
        }

        bool pop_tail(VAL &val)
        {
            return pop_tail_n(val, 0);
        }

        bool pop_tail_n(VAL &val, int n)
        {
            std::unique_lock<std::mutex> lockGuard(m_mutex);

            if (m_queue.empty())
            {
                return false;
            }

            m_prePopTimeMs = getNowMs();

            int index = n;
            if (index >= m_queue.size())
            {
                index = m_queue.size() - 1;
            }

            if (index < 0)
            {
                index = 0;
            }

            auto iter = m_queue.rbegin();
            for (int i = 0; i != index; ++i)
            {
                ++iter;
            }
            val = iter->second;

            assert(!m_queue.empty());

            m_queue.erase(--iter.base(), m_queue.end());

            return true;
        }

        bool pop_by_given_time_ref(int64_t timeRef, VAL &val)
        {
            int64_t now_ms = getNowMs();

            std::unique_lock<std::mutex> lockGuard(m_mutex);

            if (m_queue.empty())
            {
                return false;
            }

            int64_t fixedTime = timeRef;
            if (fixedTime <= m_preTimeRef)
            {
                if (m_prePopTimeMs == 0)
                {
                    fixedTime = 0;
                }
                else
                {
                    fixedTime = m_preTimeRef + now_ms - m_prePopTimeMs;
                }
            }

            chooseFrame(fixedTime, val);

            m_prePopTimeMs = now_ms;
            m_preTimeRef = timeRef;

            return true;
        }

        void chooseFrame(int64_t fixedTime, VAL &val)
        {
            if (fixedTime == 0)
            {
#if 0
                auto iter = m_queue.begin();
                for (size_t i = 0; i < m_queue.size() / 2; ++i)
                {
                    ++iter;
                }

                val = iter->second;
                m_queue.erase(m_queue.begin(), iter);
#else
                auto iter = m_queue.rbegin();
                val = iter->second;
                m_queue.clear();
#endif
            }
            else
            {
                auto iter = m_queue.upper_bound(fixedTime);

                if (iter == m_queue.end())
                {
                    val = m_queue.rbegin()->second;
                    m_queue.clear();
                }
                else
                {
                    val = iter->second;
                    m_queue.erase(m_queue.begin(), iter);
                }
            }
        }

    private:
        std::string m_name;
        std::map<int64_t, VAL> m_queue;
        std::mutex m_mutex;
        std::condition_variable m_cond;

        int64_t m_prePushKey;
        int64_t m_prePopTimeMs;
        int64_t m_preTimeRef;
        size_t m_maxSize;

        uint32_t m_pushRollBackCount;

        int64_t m_maxKey;
    };


}
