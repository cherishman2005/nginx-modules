#include <iostream>
#include <list>
#include <deque>

int main() {
    // std::list 示例
    std::list<int> myList = {1, 2, 3};
    myList.push_front(0);
    myList.push_back(4);
    for (auto it = myList.begin(); it != myList.end(); ++it) {
        std::cout << *it << " ";
    }
    std::cout << std::endl;

    // std::deque 示例
    std::deque<int> myDeque = {1, 2, 3};
    myDeque.push_front(0);
    myDeque.push_back(4);
    for (size_t i = 0; i < myDeque.size(); ++i) {
        std::cout << myDeque[i] << " ";
    }
    std::cout << std::endl;

    return 0;
}
