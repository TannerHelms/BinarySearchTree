package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func getFile(file string) string{
	letter, _ := ioutil.ReadFile(file)

	return string(letter)

}

func main() {
	start := time.Now()
	binarySearchTree := &Node{}

	getDictionary := getFile("words.txt")

	getLetter := getFile("letter.txt")

	var builder strings.Builder
	for _, byte := range getLetter {
		if byte == 10 {
			builder.WriteRune(rune(32))
		}  else if unicode.IsLetter(byte) || unicode.Is(unicode.Zs, byte) {
			builder.WriteRune(byte)
		}

	}

	dictionary := strings.Split(getDictionary, "\n")
	letter := strings.Split(builder.String() , " ")


	rand.Shuffle(len(dictionary), func(i, j int) {
		dictionary[i], dictionary[j] = dictionary[j], dictionary[i]
	})

	for i := 0; i < len(dictionary); i++ {
		binarySearchTree.insert(strings.ToLower(dictionary[i]))
	}
	for i := 0; i < len(letter); i++ {
		if !binarySearchTree.search(strings.ToLower(letter[i])) {
			fmt.Println(letter[i])
		}
	}

	//inOrderTraversal(binarySearchTree)

	fmt.Println(binarySearchTree.numberNodes())
	log.Printf("Spell Checker took: %s", time.Since(start))

}


