package internal

import "bytes"

// CompareString 比较字符串，来保证排序
func CompareString(a, b string) int {
	return bytes.Compare([]byte(a), []byte(b))
}
