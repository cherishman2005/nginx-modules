#include <iostream>
#include <sstream>
#include <string>
#include <set>
#include <map>
#include <unordered_map>
//#include <tr1/unordered_map>

using namespace std;

namespace Stream {
struct TaskID
{
    uint32_t type;
    std::string id;

    TaskID()
        : type(0)
    {
    }

    TaskID(uint32_t t, std::string i)
    {
        type = t;
        id = i;
    }
    
    bool operator < (const TaskID& right) const
    {
        return (type < right.type) ||
            ((type == right.type) && (id < right.id));
    }

    bool operator == (const TaskID& right) const
    {
        return ((type == right.type) && (id == right.id));
    }
    
    bool operator != (const TaskID& right) const
    {
        return ((type != right.type) || (id != right.id));
    }
    
    string dump() const
    {
        ostringstream os;
        os << type << ":" << id;
        return os.str();
    }
};
}


namespace std {
template<>
class hash<Stream::TaskID> {
public:
    size_t operator()(const Stream::TaskID &t) const {
        return std::hash<int>()(t.type) ^ std::hash<string>()(t.id);
    }
};
}


int main() {
    unordered_map<Stream::TaskID, string> task;

    task.insert({{1, "123"}, "555"});
    
    cout << "size=" << task.size() << endl;
    
    for (const auto & e : task) {
        cout << e.first.id << ":" << e.second << endl;
    }
    
    Stream::TaskID taskId(1, "123");
    //cout << "taskId=" << taskId << endl;
    if (task.find(taskId) == task.end()) {
        cout << "not found" << endl;
    } else {
        cout << "found" << endl;
    }
    
    if (task.find({1, "123"}) == task.end()) {
        cout << "not found" << endl;
    } else {
        cout << "found" << endl;
    }
    
    if (task.find({1, "1234"}) == task.end()) {
        cout << "not found" << endl;
    } else {
        cout << "found" << endl;
    }
    
    return 0;
}

//g++ taskid.cpp -o taskid -std=c++11
