package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abcabcbb"
	//str1 := "bbbbb"
	//str2 := "pwwkew"
	lengthOfLongestSubstring(str)

}

func lengthOfLongestSubstring(s string) int {

	longest := ""
	for i := range s {
		if strings.Index(longest, string(s[i])) != -1 {
			//如果位置小于1，将重复位置之前的所有字符全部剔除掉，
			//如果位置大于1，不提出，保持不变,作为备选项
			fmt.Println(fmt.Errorf("%c exists in %s", s[i], longest))
			continue
		}
		longest = fmt.Sprintf("%s%c", longest, s[i])
		fmt.Printf("longest=%s\n", longest)
	}
	return len(longest)

}
