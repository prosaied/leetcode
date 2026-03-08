package main

import (
	"fmt"
)

//	func maxFreqSum(s string) int {
//		vowels := make(map[string]int, 0)
//		consonants := make(map[string]int, 0)
//		var maxVowel, maxConsonant int
//		for _, c := range s {
//			if string(c) == "a" || string(c) == "e" || string(c) == "i" || string(c) == "o" || string(c) == "u" {
//				vowels[string(c)]++
//				if vowels[string(c)] > maxVowel {
//					maxVowel = vowels[string(c)]
//				}
//			} else {
//				consonants[string(c)]++
//				if consonants[string(c)] > maxConsonant {
//					maxConsonant = consonants[string(c)]
//				}
//			}
//		}
//		fmt.Println(vowels, consonants)
//		fmt.Println(maxVowel, maxConsonant)
//		return maxVowel + maxConsonant
//	}
func maxFreqSum(s string) int {
	alphabet := make(map[rune]int)
	var maxVol, maxCon int
	for _, r := range s {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			alphabet[r]++
			if alphabet[r] > maxVol {
				maxVol = alphabet[r]
			}
		default:
			alphabet[r]++
			if alphabet[r] > maxCon {
				maxCon = alphabet[r]
			}
		}
	}
	return maxCon + maxVol
}

func main() {
	fmt.Println(maxFreqSum("aeiaeia"))
}
