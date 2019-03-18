package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abcabcbb"
	str1 := "bbbbb"
	str2 := "pwwkew"
	lengthOfLongestSubstring(str)
	lengthOfLongestSubstring(str1)
	lengthOfLongestSubstring(str2)

}

func lengthOfLongestSubstring(s string) int {
	max_size := 0

	longest := ""
	for i := range s {
		index := strings.Index(longest, string(s[i]))
		if index != -1 {
			if longest[0] == s[i] {
				//如果首位与当前字符相同，将首位剔除，末尾补充该字符，继续遍历
				fmt.Println(fmt.Errorf("first:%c exists in %s", s[i], longest))
			} else {
				fmt.Println(fmt.Errorf("no first:%c exists in %s", s[i], longest))
				//如果位置大于1，将当前字符串的长度作为备选项，然后剔除重复位置之前的所有字符(包含重复位置)，末尾补充该字符,继续遍历
				if len(longest) > max_size {
					max_size = len(longest)
				}
				//然后剔除重复位置之前的所有字符(包含重复位置)，末尾补充该字符,继续遍历
			}
			longest = longest[index+1:]
		} else {
			fmt.Println(fmt.Errorf("%c no exists in %s", s[i], longest))
		}
		longest = fmt.Sprintf("%s%c", longest, s[i])
		fmt.Printf("longest=%s,max_size:%d\n", longest, max_size)
	}
	if len(longest) > max_size {
		max_size = len(longest)
	}

	return max_size

}
