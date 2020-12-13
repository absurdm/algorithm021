/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 *
 * https://leetcode-cn.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (63.22%)
 * Likes:    315
 * Dislikes: 0
 * Total Accepted:    186.7K
 * Total Submissions: 294.9K
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
 * 
 * 示例 1:
 * 
 * 输入: s = "anagram", t = "nagaram"
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: s = "rat", t = "car"
 * 输出: false
 * 
 * 说明:
 * 你可以假设字符串只包含小写字母。
 * 
 * 进阶:
 * 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
 * 
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash := [26]int{}
	for _, char := range s {
		// hash 函数:  映射每个小写字母到 hash 这个数组
		// 字符串只有小写字母, 所以可以直接用 字母的 ascii 码减去a 映射到数组
		// ascii 码 97～122 号为26个小写英文字母
		hash[char-'a']++
	}

	// 遍历第二个字符, 计算重复
	// 进行计算减一
	for _, char := range t {
		hash[char-'a']--
	}

	// hash 表所有想等于 0 是异位
	for _, num := range hash {
		if num != 0 {
			return false
		}
	}
	return true
}
// @lc code=end

