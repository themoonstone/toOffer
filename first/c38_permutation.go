package first

import "strings"

/*
	输入一个字符串，打印出该字符串中字符的所有排列。

	你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

	示例:

	输入：s = "abc"
	输出：["abc","acb","bac","bca","cab","cba"]
	 

	限制：

	1 <= s 的长度 <= 8
*/
/*
	1. 拆解s为单个字母的列表slice
	2. 遍历slice,取出其中的每一个字符作为排列的第一个字母
	3. 将剩余列表传入下层继续操作
	4. 结束条件: 生成的字符串长度与s相等
	5. 判断:给定map进行判断
*/
var flagMap = make(map[string]bool)
func Permutation(s string) []string {
	return permutation(s)
}

func permutation(s string) []string {
	result := make([]string, 0)
	curPermutation:=make([]string,0)
	leave := strings.Split(s,"")
	_permutation(s,leave,curPermutation, &result)

	return result
}

func _permutation(s string, leave []string, cur []string, res *[]string) {
	// return
	if len(leave) == 0 {
		curString := strings.Join(cur, "")
		if !flagMap[curString] {
			*res = append(*res, curString)
			flagMap[curString] = true
		}
		return
	}
	for i := 0; i < len(leave); i++ {
		cur = append(cur, leave[i])
		newLeave := append([]string{},leave[:i]...)
		newLeave = append(newLeave, leave[i+1:]...)
		_permutation(s, newLeave, cur, res)
		cur = cur[:len(cur) - 1] // 回溯
	}
}
