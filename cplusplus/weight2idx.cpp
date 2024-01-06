#include <iostream>
#include <stdlib.h>
#include <string.h>
#include <vector>
#include <set>
#include <map>
#include <unordered_map>
#include <algorithm>
#include <cmath>
#include <iterator>

using namespace std;

int main()
{

    map<int, int> idx2weight = {{0,5}, {1, 3}, {2,5}, {3, 4}};
    multimap<int, int, greater<int> > weight2idx;
    //multimap<int, int> weight2idx;
    
    for (auto & e : idx2weight) {
        int weight = e.second;
        weight2idx.insert(make_pair(weight,e.first));
    }

    for (const auto & e : weight2idx) {
        cout << e.first << " -> " << e.second << endl;
    }

    return 0;
}
