package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	appName := "anagramalyser"
	fmt.Println("-> Building ...")
	fmt.Printf("OS is %v\n", runtime.GOOS)
	if runtime.GOOS == "windows" {
		appName += ".exe"
	}
	_, err := exec.Command("go", "build", "-o", "../"+appName).Output()
	if err != nil {
		t.Errorf("Failed to build - %v", err)
	}
	fmt.Println("-> Running test ...")
	result, err := exec.Command("../"+appName, "-file", "../testfile.txt", "-showempty").Output()
	if err != nil {
		t.Errorf("Failed to execute, exit code was - %v", err)
	}
	got := string(result[:])
	if !strings.Contains(got, "Found anagrams for abc are [cba]") {
		t.Errorf("Test failed to find all anagrams for abc word")
	}
	if !strings.Contains(got, "Found anagrams for cba are [abc]") {
		t.Errorf("Test failed to find all anagams for cba word")
	}
	if !strings.Contains(got, "No anagrams found for def") {
		t.Errorf("Test failed to find empty anagrams for def word")
	}
	if !strings.Contains(got, "No anagrams found for asd") {
		t.Errorf("Test failed to find empty anagrams for asd word")
	}
}
