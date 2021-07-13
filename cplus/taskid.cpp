#include <iostream>
#include <string>
#include <set>
#include <map>
#include <unordered_map>

using namespace std;

struct TaskID
{
    uint32_t taskType;
    std::string taskId;

    TaskID()
        : taskType(0)
    {
    }

    TaskID(uint32_t type, std::string id)
    {
        taskType = type;
        taskId = id;
    }
    
    bool operator<(const TaskID& right) const
    {
        return (taskType < right.taskType) ||
            ((taskType == right.taskType) && (taskId < right.taskId));
    }

    bool operator==(const TaskID& right) const
    {
        return ((taskType == right.taskType) && (taskId == right.taskId));
    }
};


int main() {
    std::map<TaskID, string> task;


    task.insert({{1, "123"}, "555"});
    
    cout << "size=" << task.size() << endl;
    
    for (const auto & e : task) {
        cout << e.first.taskId << ":" << e.second << endl;
    }
    
    TaskID id(1, "123");
    if (task.find(id) == task.end()) {
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
