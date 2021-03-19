// 单例模式（singleton-instance）
#include <iostream>
using namespace std;

class BusinessAuthRouteDao {
public:
    ~BusinessAuthRouteDao() {}

    static BusinessAuthRouteDao& Instance() {
        static BusinessAuthRouteDao _inst;
        return _inst;
    }

    void setNum(int i) {
        i_ = i;
    }

    void showNum() {
        cout << i_ << endl;
    }

private:
    BusinessAuthRouteDao(): i_(0) {}    // Private so that it can  not be called
    BusinessAuthRouteDao(const BusinessAuthRouteDao&);  // copy constructor is private
    BusinessAuthRouteDao& operator=(BusinessAuthRouteDao const&){}; // assignment operator is private
    int i_;
};

int main() {
    //BusinessAuthRouteDao obj1 = BusinessAuthRouteDao::Instance();
    //BusinessAuthRouteDao obj2 = BusinessAuthRouteDao::Instance();


    //obj1.setNum(200); 
    //obj2.showNum();

    BusinessAuthRouteDao::Instance().setNum(1111);
    BusinessAuthRouteDao::Instance().showNum();

    return 0;
}