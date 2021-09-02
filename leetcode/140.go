package leetcode

import "strings"

// 单词拆分 II
// 给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。

// 说明：

// 分隔时可以重复使用字典中的单词。
// 你可以假设字典中没有重复的单词。
// 示例 1：

// 输入:
// s = "catsanddog"
// wordDict = ["cat", "cats", "and", "sand", "dog"]
// 输出:
// [
//   "cats and dog",
//   "cat sand dog"
// ]
// 示例 2：

// 输入:
// s = "pineapplepenapple"
// wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
// 输出:
// [
//   "pine apple pen apple",
//   "pineapple pen apple",
//   "pine applepen apple"
// ]
// 解释: 注意你可以重复使用字典中的单词。
// 示例 3：

// 输入:
// s = "catsandog"
// wordDict = ["cats", "dog", "sand", "and", "cat"]
// 输出:
// []
// 通过次数51,664提交次数106,874

func scheduler(s string, base []string, Dict map[string]map[string]int, lines *[]string) {
	if _, ok := Dict[s[0:1]]; ok {
		for word, length := range Dict[s[0:1]] {
			if len(s) >= length && word == s[:length] {
				if s == word {
					*lines = append(*lines, strings.Join(append(base, word), " "))
					continue
				}
				scheduler(s[length:], append(base, word), Dict, lines)
			}
		}
	}
}

func wordBreak(s string, wordDict []string) []string {

	lines := make([]string, 0)
	Dict := make(map[string]map[string]int)
	chars := make(map[string]struct{})
	sub := ""
	for _, word := range wordDict {
		for i := 0; i < len(word); i++ {
			if _, ok := chars[word[i:i+1]]; !ok {
				chars[word[i:i+1]] = struct{}{}
				sub += word[i : i+1]
			}

		}
		if _, ok := Dict[word[0:1]]; !ok {
			Dict[word[0:1]] = make(map[string]int)
		}
		Dict[word[0:1]][word] = len(word)
	}
	for j := 0; j < len(s); j++ {
		if !strings.Contains(sub, s[j:j+1]) {
			return lines
		}
	}
	base := make([]string, 0)
	scheduler(s, base, Dict, &lines)
	return lines
}
