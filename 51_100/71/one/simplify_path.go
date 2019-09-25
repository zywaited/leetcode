package one

import "strings"

func SimplifyPath(path string) string {
	if len(path) == 0 {
		return "/"
	}
	dirNames := strings.Split(path, "/")
	stack := make([]string, 0, len(dirNames))
	for _, dirName := range dirNames {
		if dirName == ".." {
			// 返回上一层
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		if dirName != "" && dirName != "." {
			stack = append(stack, dirName)
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	pathName := &strings.Builder{}
	for _, dirName := range stack {
		pathName.WriteString("/")
		pathName.WriteString(dirName)
	}
	return pathName.String()
}
