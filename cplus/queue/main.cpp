#include "queue.h"
#include <thread>
#include <chrono>
#include <iostream>

using namespace std;
using namespace Stream;

typedef Queue<int, 36> TestQueue;

static int test() {
    TestQueue q;
    int val = 0;
    bool b = false;
    
    // front: no value
    b = q.pop(val);
    if (b) {
        cout << "pop: val=" << val << ", size=" << q.size() << endl;
    } else {
        cout << "no value" << endl;
    }
    
    for (int i = 1; i <= 100; i++) {
        q.push(i, {i});
    }
    
    // front
    b = q.pop(val);
    if (b) {
        cout << "pop: val=" << val << ", size=" << q.size() << endl;
    } else {
        cout << "no value" << endl;
    }
    
    // tail 
    b = q.pop_tail(val);
    if (b) {
        cout << "pop_tail: val=" << val << ", size=" << q.size() << endl;
    } else {
        cout << "no value" << endl;
    }
    
    return 0;
}

static int test1() {
    TestQueue q;
    int val = 0;
    bool b = false;
    
    for (int i = 1; i <= 100; i++) {
        q.push(i, {i});
    }
    
    // tail_n
    b = q.pop_tail_n(val, 5);
    if (b) {
        cout << "pop_tail_n: val=" << val << ", size=" << q.size() << endl;
    } else {
        cout << "no value" << endl;
    }
    
    return 0;
}


static int test2() {
    TestQueue q;
    int val = 0;
    bool b = false;
    
    for (int i = 1; i <= 100; i++) {
        q.push(i, {i});
    }
    
    // pop_by_given_time_ref
    b = q.pop_by_given_time_ref(10, val);
    if (b) {
        cout << "pop_by_given_time_ref: val=" << val << ", size=" << q.size() << endl;
    } else {
        cout << "no value" << endl;
    }

    // pop_by_given_time_ref
    b = q.pop_by_given_time_ref(10, val);
    if (b) {
        cout << "pop_by_given_time_ref: val=" << val << ", size=" << q.size() << endl;
    } else {
        cout << "no value" << endl;
    }

    return 0;
}


static void update_data(uint32_t timeout, TestQueue *q)
{
    if (nullptr == q) {
        return;
    }
    std::this_thread::sleep_for(std::chrono::milliseconds(timeout));
    q->push(9, {9});
    q->push(99, {99});
    
    std::this_thread::sleep_for(std::chrono::milliseconds(timeout));
    cout << "update queue size=" << q->size() << endl;
}

static void delete_data(uint32_t timeout, TestQueue *q)
{
    if (nullptr == q) {
        return;
    }
    
    int val;
    bool b = q->pop(val, timeout);
    if (b) {
        cout << "pop: val=" << val << ", size=" << q->size() << endl;
    } else {
        cout << "no value" << endl;
    }
}

static int test3() {
    TestQueue q;
    int val = 0;
    
    std::thread t1(delete_data, 30*1000, &q);
    std::thread t(update_data, 10, &q);
    
    // tail_n
//    bool b = q.pop(val, 30);
//    if (b) {
//        cout << "pop: val=" << val << ", size=" << q.size() << endl;
//    } else {
//        cout << "no value" << endl;
//    }
    
    cout << "queue size=" << q.size() << endl;
    
    t.join();
    t1.join();
    
//    while (1) {
//        std::this_thread::sleep_for(std::chrono::milliseconds(3));
//    }
    return 0;
}


int main() {
    test();
    
    test1();
    
    test2();
    
    test3();
    
    return 0;
}
