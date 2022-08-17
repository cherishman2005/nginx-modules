# g++4.6 make_pair error

**问题描述**
```
error: no matching function for call to ‘make_pair(int&, int&)’
note: candidate is:
/usr/include/c++/4.6/bits/stl_pair.h:262:5: note: template<class _T1, class _T2> std::pair<typename std::__decay_and_strip<_T1>::__type, typename std::__decay_and_strip<_T2>::__type> std::make_pair(_T1&&, _T2&&)
UhsProxyStat.cpp:46:70: error: no matching function for call to ‘make_pair(std::string&, server::uhs_proxy::UhsStatOp2Counter&)’
```

```
#include <atomic>
#include <map>
using namespace std;

int main(void)
{
  std::atomic<bool> data_ready(false);
  std::map<int, int> m;
  m.insert(std::make_pair<int, int>(1, 1));
  
  return 0;
}
```

## 解决方法

```
  m.insert(std::make_pair<int, int>(1, 1));
```
改为
```
  m.insert(std::pair<int, int>(1, 1));
```

![image](https://user-images.githubusercontent.com/17688273/185116320-49d0d30e-3939-48bf-ab39-92591fa91dac.png)

```
template <class T1, class T2>
pair<V1, V2> make_pair(T1&& x, T2&& y) noexcept;
```

# 参考链接

- [https://stackoverflow.com/questions/3559344/error-no-matching-function-for-call-to-make-pairint-quest](https://stackoverflow.com/questions/3559344/error-no-matching-function-for-call-to-make-pairint-quest)
