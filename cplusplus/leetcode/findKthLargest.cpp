#include <iostream>
#include <string>
#include <vector>
#include <memory>
#include <iterator>
#include <queue>

using namespace std;

class Solution {
public:
    int findKthLargest(vector<int>& nums, int k) {
        // desc 最小堆
        priority_queue<int, vector<int>, greater<int>> pq;
        for (const auto & e : nums) {
            if (pq.size() < k) {
                pq.push(e);
                continue;
            }
            if (pq.top() < e) {
                pq.pop();
                pq.push(e);
            }
        }

        return pq.top();
    }
};

int main() {
    vector<int> nums = {3,2,1,5,6,4};
    int k = 2;
    shared_ptr<Solution> solution = make_shared<Solution>();

    auto result = solution->findKthLargest(nums, k);
    cout << result << endl;
    //copy(result.begin(), result.end(), ostream_iterator<int>(cout, ","));
    cout << endl;
    
    return 0;
}
