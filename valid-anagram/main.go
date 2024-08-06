package main

import "fmt"

func main() {
	res := isAnagram("anagram", "nagaram")
	fmt.Println(res)
}

func isAnagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}

	letter1 := make(map[rune]int)
	letter2 := make(map[rune]int)
	for _, k := range t {
		if _, ok := letter2[k]; ok {
			letter2[k]++
			continue
		}
		letter2[k]++
	}

	for _, v := range s {
		if _, ok := letter1[v]; ok {
			letter1[v]++
			continue
		}
		letter1[v]++
	}

	for _, v := range s {
		if _, ok := letter2[v]; ok {
			if letter2[v] != letter1[v] {
				return false
			}
		} else if _, ok := letter2[v]; !ok {
			return false
		}

	}

	return true
}
