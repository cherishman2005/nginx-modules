
# C++对map遍历删除

## C++11对map遍历删除
```
#include <iostream>
#include <stdlib.h>
#include <string.h>
#include <vector>
#include <set>
#include <map>
#include <unordered_map>
#include <algorithm>
#include <cmath>
#include <iterator>

using namespace std;

int main()
{

    map<int, int> tasks = {{1,1}, {2, 2}, {3,3}, {4, 4}};
    
    // 形如如下逻辑
    for (const auto & e : tasks) {
        // 根据业务场景，满足条件的元素进行删除操作
        if (e.first != 3) continue;
        
        const auto & it = tasks.find(e.first);
        if (it != tasks.end()) {
            tasks.erase(it);
        }
    }
    
    

    for (const auto & e : tasks) {
        cout << e.first << endl;
    }

    return 0;
}
```



