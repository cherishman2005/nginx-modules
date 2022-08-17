#include <iostream>
#include <atomic>
#include <map>
using namespace std;

int main(void)
{
  std::atomic<bool> data_ready(false);
  std::map<int, int> m;
  m.insert(std::make_pair<int, int>(1, 2));
  
  for (std::map<int, int>::const_iterator it = m.begin(); it != m.end(); it++) {
      cout << it->first << " : " << it->second << endl;
  }
  
  return 0;
}

// g++ atomic_test.cpp -o atomic_test -std=c++0x
