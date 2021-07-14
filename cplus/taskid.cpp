#include <iostream>
#include <string>
#include <set>
#include <map>
#include <unordered_map>

using namespace std;

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
        type = type;
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
};


int main() {
    std::map<TaskID, string> task;


    task.insert({{1, "123"}, "555"});
    
    cout << "size=" << task.size() << endl;
    
    for (const auto & e : task) {
        cout << e.first.id << ":" << e.second << endl;
    }
    
    TaskID taskId(1, "123");
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
