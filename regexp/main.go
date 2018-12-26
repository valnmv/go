package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	needle := "(?i)chocolate"
	haystack := "Chocolate is my favorite"
	matched, err := regexp.MatchString(needle, haystack)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(matched)

	re := regexp.MustCompile("^[a-zA-Z0-9]{5,12}")
	fmt.Println(re.MatchString("slimshady99"))
	fmt.Println(re.MatchString("!asdf£33£3"))
	fmt.Println(re.MatchString("roger"))
	fmt.Println(re.MatchString("iamthebestuseofthisappevaaaar"))

	re = regexp.MustCompile("^[\u0000-\u00FF\u0400-\u04FF]{5,12}")

	words := "Düsseldorf, Köln, Москва,!@#$"
	words2 := "北京市, إسرائيل"
	fmt.Println(words, re.MatchString(words))
	fmt.Println(words2, re.MatchString(words2))
}
