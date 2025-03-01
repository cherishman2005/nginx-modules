#include <iostream>

using namespace std;

#include <iostream>
#include <memory>

class MyClass {
public:
    MyClass() { std::cout << "MyClass constructor" << std::endl; }
    ~MyClass() { std::cout << "MyClass destructor" << std::endl; }
    void doSomething() { std::cout << "Doing something..." << std::endl; }
};

int main() {
    std::unique_ptr<MyClass> ptr = std::make_unique<MyClass>();
    ptr->doSomething();

    // 移动所有权
    std::unique_ptr<MyClass> ptr2 = std::move(ptr);
    if (ptr == nullptr) {
        std::cout << "ptr is null after move" << std::endl;
    }
    ptr2->doSomething();
    return 0;
}

//  g++ .\unique_ptr.cpp -o .\unique_ptr -std=c++14