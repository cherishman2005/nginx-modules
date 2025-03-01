#include <iostream>
#include <memory>

using namespace std;

class MyClass {
public:
    MyClass() { std::cout << "MyClass constructor" << std::endl; }
    ~MyClass() { std::cout << "MyClass destructor" << std::endl; }
    void doSomething() { std::cout << "Doing something..." << std::endl; }
};

int main() {
    std::shared_ptr<MyClass> ptr1 = std::make_shared<MyClass>();
    std::shared_ptr<MyClass> ptr2 = ptr1;
    std::cout << "Use count: " << ptr1.use_count() << std::endl;
    ptr1->doSomething();
    ptr2->doSomething();
    return 0;
}
