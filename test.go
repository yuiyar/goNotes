package main

import "fmt"

func main() {
	s := "abcaav"
	t := "aaaabcv"
	fmt.Println(string(findTheDifference(s, t)))
}

func findTheDifference(s, t string) (diff byte) {

	for i := range s {
		diff ^= s[i] ^ t[i]
	}

	return diff ^ t[len(t)-1]

}

/*func findTheDifference(s string, t string) byte {
	var sum int

	for i := 0; i < len(s); i++ {
		sum += int(t[i]) - int(s[i])
	}

	sum += int(t[len(s)])
	return byte(sum)

}*/

/*func findTheDifference(s string, t string) byte {
	maps := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		maps[s[i]]++
	}
	for i := 0; ; i++ {
		maps[t[i]]--
		if maps[t[i]] < 0 {
			return t[i]
		}
	}
}*/
