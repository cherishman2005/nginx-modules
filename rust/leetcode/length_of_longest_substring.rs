use std::collections::HashMap;

struct Solution; // 定义 Solution 结构体（Rust 中 trait/impl 的常见用法）

impl Solution {
    pub fn length_of_longest_substring(s: String) -> i32 {
        let mut mp: HashMap<char, usize> = HashMap::new();
        let mut start = 0;
        let mut max_len = 0;
        
        for (i, c) in s.chars().enumerate() {
            if let Some(&prev_idx) = mp.get(&c) {
                start = start.max(prev_idx + 1); // 更新左边界
            }
            mp.insert(c, i); // 记录字符最新位置
            max_len = max_len.max(i - start + 1); // 计算当前窗口长度
        }
        
        max_len as i32 // 转换为 i32 类型返回
    }
}

// 辅助函数：获取两个 usize 的最大值（也可直接使用标准库的 std::cmp::max）
fn max(a: usize, b: usize) -> usize {
    if a > b { a } else { b }
}

fn main() {
    // 测试用例 1: 无重复字符的最长子串为 "abc"，长度 3
    let s1 = "abcabcbb".to_string();
    let result1 = Solution::length_of_longest_substring(s1);
    println!("测试用例 1 结果: {}", result1); // 输出: 3

    // 测试用例 2: 全重复字符，最长子串长度为 1
    let s2 = "bbbbb".to_string();
    let result2 = Solution::length_of_longest_substring(s2);
    println!("测试用例 2 结果: {}", result2); // 输出: 1

    // 测试用例 3: 最长子串为 "wke"，长度 3
    let s3 = "pwwkew".to_string();
    let result3 = Solution::length_of_longest_substring(s3);
    println!("测试用例 3 结果: {}", result3); // 输出: 3

    // 空字符串测试
    let s4 = "".to_string();
    let result4 = Solution::length_of_longest_substring(s4);
    println!("空字符串测试结果: {}", result4); // 输出: 0
}
