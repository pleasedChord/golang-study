package main

import (
	"fmt"
	"strings"
)

func find1(ss []string) {
	strMinLen := len(ss[0])

	for _, cli := range ss {
		if len(cli) < strMinLen {
			strMinLen = len(cli)
		}
	}

	minStr := ""
	for i := 0; i < strMinLen; i++ {
		m := make(map[string]string)

		for _, cli := range ss {
			m[cli[:i+1]] = cli
		}

		if len(m) == 1 {
			for k, _ := range m {
				minStr = k
			}
		} else if len(minStr) == 0 {
			fmt.Println("没有最长公共前缀")
			break
		} else {
			fmt.Println("最长公共前缀是:", minStr)
			break
		}
	}
}

func find2(ss []string) string {
	if len(ss) == 0 {
		fmt.Println("非法输入")
		return ""
	}

	prefix := ss[0]
	for _, cli := range ss[1:] {
		for len(prefix) > 0 {
			if strings.HasPrefix(cli, prefix) {
				return prefix
			}

			prefix = prefix[:len(prefix)-1]
		}

		if len(prefix) == 0 {
			return ""
		}
	}

	return prefix
}

/*
最长公共前缀
查找字符串数组中的最长公共前缀
*/
func main() {
	ss := []string{"abcg", "abcde", "abcdf"}
	// find1(ss)
	prefix := find2(ss)
	fmt.Println("最长公共前缀是:", prefix)
}
