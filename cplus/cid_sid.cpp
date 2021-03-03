#include <iostream>
#include <set>

using namespace std;

struct CidSid
{
    uint32_t cid;
    uint32_t sid;

    CidSid()
        : cid(0)
        , sid(0)
    {
    }

    CidSid(uint32_t c, uint32_t s)
    {
        cid = c;
        sid = s;
    }
    
    bool operator<(const CidSid& right) const
    {
        return (cid < right.cid) ||
            ((cid == right.cid) && (sid < right.sid));
    }

    bool operator==(const CidSid& right) const
    {
        return ((cid == right.cid) && (sid == right.sid));
    }
};


int main() {
    std::set<CidSid> cidSidSet;
    
    cidSidSet.insert({123, 456});
    cidSidSet.insert({123, 456});
    cidSidSet.insert({123, 555});
    
    cidSidSet.insert({111, 555});
    
    
    cout << "size=" << cidSidSet.size() << endl;
    
    for (const auto & e : cidSidSet) {
        cout << e.cid << ":" << e.sid << endl;
    }
    
    CidSid cidSid(123, 456);
    if (cidSidSet.find(cidSid) == cidSidSet.end()) {
        cout << "not found" << endl;
    } else {
        cout << "found" << endl;
    }
    
    return 0;
}

//g++ cid_sid.cpp -o cid_sid -std=c++11