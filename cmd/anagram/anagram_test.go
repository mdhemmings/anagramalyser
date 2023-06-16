package anagram

import (
	"fmt"
	"testing"
)

func TestCheckAnagrams(t *testing.T) {
	got := CheckAnagrams("abc", []string{"cba", "def"})
	want := []string{"cba"}
	if !arrayChecker(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCheckIsAnagram(t *testing.T) {
	got := isAnagram("cba", "abc")
	if got != true {
		t.Errorf("Got false when this should have been true")
	}
}

func arrayChecker(array1, array2 []string) bool {
	fmt.Println(len(array1))
	fmt.Println(len(array2))
	if len(array1) != len(array2) {
		return false
	}
	for i, array := range array1 {
		if array != array2[i] {
			return false
		}
	}
	return true
}
