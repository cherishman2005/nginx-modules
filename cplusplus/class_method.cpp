#include<iostream>
using namespace std;
class Parent
{
public:
    int mi;
    void add(int a)
    {
        mi += a;
    }
    void add(int a, int b)
    {
        mi += (a + b);
    }
};
class Child : public Parent
{
public:
    int mi;
    void add(int a, int b, int c)
    {
        mi += (a + b + c);
    }
};
int main()
{
    Child c;
    c.mi = 10;
    c.Parent::mi = 100;
    c.Parent::add(1);
    c.Parent::add(2, 3);
    c.add(4, 5, 6);
    cout << "c.mi = " << c.mi << endl;
    cout << "c.Parent::mi = " << c.Parent::mi << endl;
    return 0;
}