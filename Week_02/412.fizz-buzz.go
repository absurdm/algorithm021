/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 *
 * https://leetcode-cn.com/problems/fizz-buzz/description/
 *
 * algorithms
 * Easy (65.21%)
 * Likes:    77
 * Dislikes: 0
 * Total Accepted:    51.3K
 * Total Submissions: 78.6K
 * Testcase Example:  '1'
 *
 * 写一个程序，输出从 1 到 n 数字的字符串表示。
 * 
 * 1. 如果 n 是3的倍数，输出“Fizz”；
 * 
 * 2. 如果 n 是5的倍数，输出“Buzz”；
 * 
 * 3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。
 * 
 * 示例：
 * 
 * n = 15,
 * 
 * 返回:
 * [
 * ⁠   "1",
 * ⁠   "2",
 * ⁠   "Fizz",
 * ⁠   "4",
 * ⁠   "Buzz",
 * ⁠   "Fizz",
 * ⁠   "7",
 * ⁠   "8",
 * ⁠   "Fizz",
 * ⁠   "Buzz",
 * ⁠   "11",
 * ⁠   "Fizz",
 * ⁠   "13",
 * ⁠   "14",
 * ⁠   "FizzBuzz"
 * ]
 * 
 * 
 */

// @lc code=start
func fizzBuzz(n int) []string {
	// 暴力解法
	// golang map 无序
	res := []string{}
	for i := 1; i <= n; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			res = append(res, "FizzBuzz")
		} else if i % 3 == 0 {
			res = append(res, "Fizz")
		} else if i % 5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}

// var wg sync.WaitGroup
// func fizzBuzz(n int) []string {
// 	ret := make([]string, n)
// 	for i := 0; ; {
// 		size := 1000
// 		if n-i < 1000 {
// 			size = n - i
// 		}
// 		if size <= 0 {
// 			break
// 		}

// 		wg.Add(1)
// 		go quickFizzBuzz(i, i+size, ret)

// 		i += size
// 	}

// 	wg.Wait()

// 	return ret
// }

// func quickFizzBuzz(start, end int, ret []string) {
// 	for i := start; i < end; i++ {
// 		fizz := (i+1)%3 == 0
// 		buzz := (i+1)%5 == 0
// 		if fizz {
// 			ret[i] = "Fizz"
// 		}
// 		if buzz {
// 			ret[i] += "Buzz"
// 		}
// 		if !fizz && !buzz {
// 			ret[i] = strconv.Itoa(i + 1)
// 		}
// 	}
// 	wg.Done()
// }

// @lc code=end

