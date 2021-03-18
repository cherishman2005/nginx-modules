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
    BusinessAuthRouteDao(): i_(0) {}
    BusinessAuthRouteDao(const BusinessAuthRouteDao&);
    int i_;
};

int main() {
    //BusinessAuthRouteDao obj1 = BusinessAuthRouteDao::Instance();
    //BusinessAuthRouteDao obj2 = BusinessAuthRouteDao::Instance();

    BusinessAuthRouteDao::Instance().setNum(1111);
    BusinessAuthRouteDao::Instance().showNum();

    return 0;
}