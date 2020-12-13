/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 *
 * https://leetcode-cn.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (64.06%)
 * Likes:    539
 * Dislikes: 0
 * Total Accepted:    125.3K
 * Total Submissions: 195.4K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
 * 
 * 示例:
 * 
 * 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
 * 输出:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 * 
 * 说明：
 * 
 * 
 * 所有输入均为小写字母。
 * 不考虑答案输出的顺序。
 * 
 * 
 */

// @lc code=start
type bytes []byte

// 类型排序接口实现, Len(), Less(), Swap() 接口需要函数
func (b bytes) Len() int { return len(b) }

func (b bytes) Less(i, j int) bool { return b[i] < b[j] }

func (b bytes) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	m := map[string]int{}
	for _, s := range strs {
		kByte := bytes(s)
		// 排序
		sort.Sort(kByte)
		// 排序后的结果字符串做 map k
		k := string(kByte)
		if idx, ok := m[k]; ok {
			// map 中已经出现过得元素, 根据 idx 放入结果字符串
			res[idx] = append(res[idx], s)
		} else {
			// 新元素放入到 map 中去, 并记录在结果中的下标为 len(res)
			m[k] = len(res)
			res = append(res, []string{s})
		}
	}
	return res
}
// @lc code=end

