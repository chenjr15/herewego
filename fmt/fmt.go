package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stack := []string{}
	var tabCnt, maxLen int
	var line, name string
	var err error
	r := bufio.NewReader(os.Stdin)
	for {

		// 按行读取
		line, err = r.ReadString('\n')
		// fmt.Println("Got:", line, err)

		if err != nil {
			break
		}
		// 去除换行
		line = line[:len(line)-1]

		// 读取\t数量
		for tabCnt = 0; tabCnt < len(line); {
			if line[tabCnt] == '\t' {
				tabCnt++
			} else {
				break
			}
		}
		name = line[tabCnt:]
		if isFile(name) {
			// 计算长度
			size := len(name) + len(stack)
			for _, n := range stack {
				size += len(n)
			}
			if size > maxLen {
				maxLen = size

			}
			fmt.Println(tabCnt, stack, name, size)
			continue

		}
		// 处理文件栈

		if tabCnt >= len(stack) {
			// 下一层
			stack = append(stack, name)
		} else {
			stack = stack[0 : tabCnt+1]
			stack[tabCnt] = name

		}

	}

	fmt.Println(maxLen)

}
func isFile(name string) bool {
	for _, c := range name {
		if c == '.' {
			return true
		}

	}
	return false
}
