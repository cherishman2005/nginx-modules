// https://www.educative.io/edpresso/what-is-the-stdbadalloc-exception-in-cpp

#include <iostream>
#include <new>
#include <vector>
#include <set>
using namespace std;

#define SIZE 53771823
//#define SIZE 1000000000000

vector<int>   a(SIZE, 0);

 // Driver code 
int main () {
  try
  {
    a[100] = 1111;
    //int * myarray = new int[1000000000000]; 
    set<int> s;
    s.insert(a.begin(),a.end());

    set<int> tmp, tmp1;
    tmp = s;
    tmp1 = s;
  }
  catch (std::bad_alloc & exception)
  {
     std::cerr << "bad_alloc detected: " << exception.what() << std::endl;
  }
  return 0;
}