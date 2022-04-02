#include <iostream>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <vector>
#include <set>
#include <map>
#include <unordered_map>
#include <algorithm>
#include <cmath>
#include <sstream>
#include <iterator>
#include <math.h>
#include <string>

using namespace std;

// alloc-dealloc-mismatch

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */

struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution {
public:
    ListNode* removeElements(ListNode* head, int val) {
        vector<int> vec;
        for (ListNode *p = head; p; p = p->next) {
            ListNode *q = p;
            if (p->val != val) {
                vec.push_back(p->val);
            }
            //free(q);
        }
        copy(vec.begin(), vec.end(), ostream_iterator<int>(cout, ","));
        cout << endl;

        ListNode *h = nullptr;
        ListNode *p = nullptr;
        for (auto e : vec) {
            ListNode *q = new ListNode(e);
            //q->val = e;
            //q->next = nullptr;
            if (!p) {
                p = q;
                h = p;
            } else {
                p->next = q;
                p = q;
            }
        }

        return h;
    }

//private:
    ListNode* createList(const vector<int> &vec) {
        ListNode *h = nullptr;
        ListNode *p = nullptr;
        for (auto e : vec) {
            ListNode *q = new ListNode(e);
            //q->val = e;
            //q->next = nullptr;
            if (!p) {
                p = q;
                h = p;
            } else {
                p->next = q;
                p = q;
            }
        }
        return h;
    }

};

int main()
{
    int s = 7;
    vector<int> nums = {2,3,1,2,4,3};
    
    Solution solution;
    
    ListNode* head = solution.createList(nums);
        
    auto o = solution.removeElements(head, 1);
    
    cout << o << endl;

    //std::copy(o.begin(), o.end(), ostream_iterator<int>(std::cout, ","));

    return 0;
}
