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

## c++0x和c++11

上一个版本的C++国际标准是2003年发布的，所以叫C++ 03。

然后C++国际标准委员会在研究C++ 03的下一个版本的时候，一开始计划是07年发布，所以最初这个标准叫C++ 07。

但是到06年的时候，官方觉得07年肯定完不成C++ 07，而且官方觉得08年可能也完不成。最后干脆叫C++ 0x。x的意思是不知道到底能在07还是08还是09年完成。

结果2010年的时候也没完成，最后在2011年终于完成了C++标准。所以最终定名为C++11。

The old -std=c++0x is only needed for older compiler versions that did not support -std=c++11 and they chose that name to express the preliminary and unstable nature of features (and the ABI) of the then upcoming C++11 (and when it was still unclear whether that would eventually become C++10 or C++12). They changes some of the details adapting to the changing working drafts of the standard at the time before the C++11 standard was officially released.If your compiler supports -std=c++11, there is no reason to use -std=c++0x. Concerning compatibility: There might even be differences and incompatibilities, but these are not just bound to the use of -std=c++0x, but to specific versions of the compiler. When the compiler supports both, they should be identical.

# 参考链接

- [https://stackoverflow.com/questions/3559344/error-no-matching-function-for-call-to-make-pairint-quest](https://stackoverflow.com/questions/3559344/error-no-matching-function-for-call-to-make-pairint-quest)

- [https://www.zhihu.com/question/20141092](https://www.zhihu.com/question/20141092)

