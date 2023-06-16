package anagram

func isAnagram(s string, t string) bool {
	lenS := len(s)
	lenT := len(t)
	if s == t {
		return false
	}
	if lenS != lenT {
		return false
	}
	anagramMap := make(map[string]int)
	for i := 0; i < lenS; i++ {
		anagramMap[string(s[i])]++
	}
	for i := 0; i < lenT; i++ {
		anagramMap[string(t[i])]--
	}
	for i := 0; i < lenS; i++ {
		if anagramMap[string(s[i])] != 0 {
			return false
		}
	}
	return true
}

func CheckAnagrams(checkWord string, words []string) (anagrams []string) {
	for _, word := range words {
		if len(word) != len(checkWord) {
			// if the length of the two words doesnt match then we immediately discard this iteration of the loop and break
			break
		}
		// pass over to 'isAnagram' function to do some math related stuff that I don't 100% understand
		if isAnagram(checkWord, word) {
			// if this check comes back as true then we'll add it to our array that we'll return
			anagrams = append(anagrams, word)
		}
	}
	return
}
