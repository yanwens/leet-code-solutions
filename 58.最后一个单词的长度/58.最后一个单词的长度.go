/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 *
 * https://leetcode-cn.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (28.61%)
 * Total Accepted:    16.7K
 * Total Submissions: 58.5K
 * Testcase Example:  '"Hello World"'
 *
 * 给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
 *
 * 如果不存在最后一个单词，请返回 0 。
 *
 * 说明：一个单词是指由字母组成，但不包含任何空格的字符串。
 *
 * 示例:
 *
 * 输入: "Hello World"
 * 输出: 5
 *
 *
 */
func lengthOfLastWord(s string) int {
	n := len(s) - 1
	flag := true
	res := 0
	for n >= 0 {
		if string(s[n]) == " " {
			if flag {
				s = s[0:n]
			} else {
				return res
			}
		} else {
			flag = false
			res++
		}
		n--
	}
	return res
}
