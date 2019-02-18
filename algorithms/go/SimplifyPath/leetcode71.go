package SimplifyPath

import "strings"

func simplifyPath(path string) string {
	length := len(path)
	stack := make([]string, 0, length >> 1)
	pathComponent := make([]byte, 0, length)

	for i:=0; i < length; i++ {
		for i < length && path[i] != '/' {
			pathComponent = append(pathComponent, path[i])
		}
		s := string(pathComponent)
		switch s {
		case ".", "":
		case "..":
			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, s)
		}
	}
	return "/" + strings.Join(stack, "/")
}