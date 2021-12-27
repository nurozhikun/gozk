package ztest

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	// fmt.Println(s1)
	// s2 := s1[1:3]
	// fmt.Println(s2)
	// s2[1] = 0
	// fmt.Println(s1)
	// fmt.Println(s2)
	// s1[1] = 10
	// fmt.Println(s1)
	// fmt.Println(s2)
	// fmt.Println(s1[5:], len(s1[5:]))
	// s2 = nil
	// fmt.Println(s2)
	//
	s3 := s1[1:2]
	fmt.Println(s3)
	// s1 = s3
	fmt.Println(append(s3, s3...))
	fmt.Println(s1)
	fmt.Println(s3)
	fmt.Println(s1[1:1] == nil)
	s3 = nil
	fmt.Println(s3 == nil)
}
