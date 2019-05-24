package main

import "fmt"

func findSubstring(s string, words []string) []int {
	if s == "" || len(words) == 0 {
		return nil
	}
	var k int
	var status =true
	var result = []int{}
	wordsLen := len(words)
	wordLen := len(words[0])
	maps := initMap( words)
	mapLen := len(maps)
	for i:=0; i <= len(s)-wordLen*wordsLen;i++{
		if _,ok := maps[s[i:i+wordLen]]; !ok{
			continue
		}
		start := i
		for status  {
			if k == mapLen {
				break
			}
			if _, ok := maps[s[start:start+wordLen]];!ok{
				break
			}
			maps[s[start:start+wordLen]] -=1
			start = start+wordLen
			status, k = completed(maps)
		}
		if k == mapLen {
			result = append(result, i)
		}
		maps = initMap(words)
		k = 0
		status = true
	}
	return result
}

func completed(m map[string]int) (bool,int) {
	i := 0
	for _, v := range m{
		if v < 0{
			return false, i
		}
		if v == 0 {
			i++
		}
	}
	return true, i
}

func initMap(words []string) map[string]int{
	var maps = make(map[string]int)
	for _,v := range words {
		maps[v]+=1
	}
	return maps
}

func main() {
	s := "a"
	words := []string{"a"}
	fmt.Println(findSubstring(s, words))
}

/*
"a"
["a"]

"barfoofoobarthefoobarman"
["bar","foo","the"]

"wordgoodgoodgoodbestword"
["word","good","best","good"]
[8]
*/