package leetcode

import (
	"strings"
)

// 赎金信
func CanConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}
	dic := make(map[byte]int)
	for _, value := range magazine {
		dic[byte(value)] += 1
	}
	for _, value := range ransomNote {
		if num, ok := dic[byte(value)]; ok {
			if num == 0 {
				return false
			}
			dic[byte(value)] -= 1
		} else {
			return false
		}
	}
	return true
}

// 同构字符串
func IsIsomorphic(s string, t string) bool {
	dic1 := make(map[byte]byte)
	dic2 := make(map[byte]byte)
	for i, value := range s {
		if word, ok := dic1[byte(value)]; ok {
			if word != t[i] {
				return false
			}
		} else if word, ok := dic2[t[i]]; ok {
			if word != byte(value) {
				return false
			}
		} else {
			dic1[byte(value)] = t[i]
			dic2[t[i]] = byte(value)
		}
	}
	return true
}

// 单词规律
func WordPattern(pattern string, s string) bool {
	words := strings.Split(pattern, " ")
	length := len(words)
	if len(s) != length {
		return false
	}
	dic1, dic2 := make(map[byte]string), make(map[string]byte)
	for index, value := range s {
		if dic1[byte(value)] != "" && dic1[byte(value)] != words[index] {
			return false
		} else if dic2[words[index]] > 0 && dic2[words[index]] != byte(value) {
			return false
		} else {
			dic1[byte(value)] = words[index]
			dic2[words[index]] = byte(value)
		}
	}
	return true
}

// 有效的字母异构词
func IsAnagram(s string, t string) bool {
	hashmap := make(map[byte]int)
	if len(s) != len(t) {
		return false
	}
	for _, value := range s {
		hashmap[byte(value)] += 1
	}
	for _, value := range t {
		hashmap[byte(value)]--
		if hashmap[byte(value)] < 0 {
			return false
		}
	}
	return true
}

// 字母异构词分组
func GroupAnagrams(strs []string) [][]string {
	dic := make(map[[26]int][]string)
	for _, str := range strs {
		temp := [26]int{}
		for _, v := range str {
			index := v - 'a'
			temp[index] += 1
		}
		dic[temp] = append(dic[temp], str)
	}
	result := make([][]string, 0)
	for _, v := range dic {
		result = append(result, v)
	}
	return result
}
