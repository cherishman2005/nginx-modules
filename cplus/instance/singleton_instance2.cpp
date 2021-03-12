#include <iostream>
using namespace std;

#include <iostream>
using namespace std;
//饿汉式
class Singleton{
public:
    //提供访问单例对象的结果
    static Singleton &Instance(void){
        return s_instance;
    }
    void print(void)const{
        cout << m_data << endl;
    }
    
    void setNum(int i) {
        i_ = i;
    }
      
    void showNum() {
        cout << i_ << endl;
    }
    
private:
    //私有化构造函数
    Singleton(int data):m_data(data){}
    Singleton(const Singleton&);
    //使用一个静态成员维护单例
    static Singleton s_instance;
    int m_data;
    int i_;
};
Singleton Singleton::s_instance(12345);
//静态成员在类的外部的全局区初始化，是语法规定。并不是在真正意义上在类的外部访问了私有成员变量。

/*
int main() {
    Singleton obj1 = Singleton::Instance();
    Singleton obj2 = Singleton::Instance();
    
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
    Singleton::Instance();
    
    Singleton::Instance().setNum(1111);
    Singleton::Instance().showNum();
    
    Singleton::Instance().setNum(200);
    Singleton::Instance().showNum();
    
    return 0;
}
