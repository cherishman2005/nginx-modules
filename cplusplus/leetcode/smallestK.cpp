#include <iostream>
#include <string>
#include <vector>
#include <memory>
#include <iterator>
#include <map>
#include <set>
#include <algorithm>

using namespace std;

class Solution {
public:
    vector<int> smallestK(vector<int>& arr, int k) {
        // 采用multimap模拟大顶堆
        multiset<int, greater<int>> mp;
        for (const auto & e : arr) {
            if (mp.size() < k) {
                mp.insert(e);
                continue;
            }

            auto itBegin = mp.begin();
            if (*itBegin > e) {
                mp.erase(itBegin);  // 挖掉最大值
                mp.insert(e);
            }
        }

        vector<int> res;
        for (const auto & e : mp) {
            res.push_back(e);
        }

        std::reverse(res.begin(), res.end());
        return res;
    }
};

int main() {
    vector<int> arr = {1,3,5,7,2,4,6,8};
    int k = 4;
    shared_ptr<Solution> solution = make_shared<Solution>();

    auto result = solution->smallestK(arr, k);
    //cout << result << endl;
    copy(result.begin(), result.end(), ostream_iterator<int>(cout, ","));
    cout << endl;
    
    return 0;
}
