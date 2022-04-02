// 非单例模式（non-singleton-instance）
#include <iostream>
using namespace std;

class BusinessAuthRouteDao {
 public:

  ~BusinessAuthRouteDao() {}

  static BusinessAuthRouteDao& Instance() {
    //static BusinessAuthRouteDao _inst;
    return _inst;
  }
  
  int i_;
  void setNum(int i) {
    i_ = i;
  }
  
  void showNum() {
    cout << i_ << endl;
  }
private:
    //私有化构造函数
    BusinessAuthRouteDao(): i_(0) {}
    BusinessAuthRouteDao(const BusinessAuthRouteDao&);
    //使用一个静态成员维护单例
    static BusinessAuthRouteDao _inst;
};

BusinessAuthRouteDao BusinessAuthRouteDao::_inst;

/*
int main() {
    BusinessAuthRouteDao obj1 = BusinessAuthRouteDao::Instance();
    BusinessAuthRouteDao obj2 = BusinessAuthRouteDao::Instance();
    
    obj1.setNum(1111);
    obj1.showNum();
    obj2.showNum();
    
    obj2.setNum(200);
    obj1.showNum();
    obj2.showNum();
    
    return 0;
}
*/


int main() {
    //BusinessAuthRouteDao obj1 = BusinessAuthRouteDao::Instance();

    BusinessAuthRouteDao::Instance().setNum(1111);
    BusinessAuthRouteDao::Instance().showNum();
    
    
    return 0;
}