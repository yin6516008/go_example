package stringutil_test

import (
	"fmt"
	"testing"
)

func TestStringutil(t *testing.T) {
	r := []rune("hello")
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		// 第一轮i=0,j=4
		// 第二轮i=1,j=3
		r[i], r[j] = r[j], r[i]
		fmt.Println(string(r[i]), string(r[j]), i, j)
	}
	fmt.Println(string(r))
}
