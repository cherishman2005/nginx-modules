# setprecision-format

## setprecision格式化
```
#include <iostream>
#include <iomanip>
using namespace std;

void main()
{
    double pi = 3.14159265;
    cout << pi << endl;                     // 默认以6精度，所以输出为 3.14159
    cout << setprecision(4) << pi << endl;  // 改成4精度，所以输出为3.142
    cout << setprecision(8) << pi << endl;  // 改成8精度，所以输出为3.1415927
    cout << fixed << setprecision(4) << pi << endl; // 加了fixed意味着是固定点方式显示，所以这里的精度指的是小数位，输出为3.1416
    cout << pi << endl;                     // fixed和setprecision的作用还在，依然显示3.1416
    cout.unsetf(ios::fixed);                // 去掉了fixed，所以精度恢复成整个数值的有效位数，显示为3.142
    cout << pi << endl;
    cout.precision(6);                      // 恢复成原来的样子，输出为3.14159
    cout << pi << endl;
}
```

## 取消c++所设置的cout中setprecision输出的格式

C++11以前是这样写的：
```
std::cout.unsetf(std::ios_base::floatfield);
 ```
 
C++11：
```
std::cout << std::defaultfloat;
```


