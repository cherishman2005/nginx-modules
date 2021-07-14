#include <unordered_map>
#include <utility>
#include <cstdint>
#include <iostream>
#include <functional>
 
namespace test{
  enum COLOR{ WHITE, BLAC };
}
 
#if 0  // 如果没有这里，G++4.8.4和G++5.4.0会报错
namespace std {
template<>
struct hash<test::COLOR> {
   typedef test::COLOR argument_type;
   typedef size_t result_type;
 
   result_type operator () (const argument_type& x) const {
      using type = typename std::underlying_type<argument_type>::type;
      return std::hash<type>()(static_cast<type>(x));
   }
};
}
#endif
 
namespace test{
class mytest{
 public:
  std::unordered_map<COLOR, int> id_map_;
};
}
 
int main(){
    test::mytest t;
    return 0;
}