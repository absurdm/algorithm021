/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 *
 * https://leetcode-cn.com/problems/ugly-number-ii/description/
 *
 * algorithms
 * Medium (54.58%)
 * Likes:    440
 * Dislikes: 0
 * Total Accepted:    40.4K
 * Total Submissions: 73.8K
 * Testcase Example:  '10'
 *
 * 编写一个程序，找出第 n 个丑数。
 * 
 * 丑数就是质因数只包含 2, 3, 5 的正整数。
 * 
 * 示例:
 * 
 * 输入: n = 10
 * 输出: 12
 * 解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
 * 
 * 说明:  
 * 
 * 
 * 1 是丑数。
 * n 不超过1690。
 * 
 * 
 */

// @lc code=start
func nthUglyNumber(n int) int {
    a,b,c:= 0,0,0
    dp:= make([]int,n,n)
    for i:=0;i<n;i++{
        dp[i]=1
    }
    min := func(a,b int)int{
        if a>b{
            return b
        }
        return a
    }
    for i:=1;i<n;i++{
        t1,t2,t3:=dp[a]*2,dp[b]*3,dp[c]*5
        dp[i]=min(min(t1,t2),t3)
        if dp[i] ==t1 {
            a++
        }
        if dp[i] ==t2 {
            b++
        }
        if dp[i] ==t3 {
            c++
        }
    }
    return dp[n-1]
}
// @lc code=end

