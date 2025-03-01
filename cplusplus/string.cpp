#include <iostream>
#include <string>

using namespace std;

class Singleton {

};

int main() {
    string str = "hello";
    str[0] = 'z';
    cout << str << endl;
    
    return 0;
}
