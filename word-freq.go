// Program to find the top most common words in a text file

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)



func usageText() {
	fmt.Println("Count word-frequency program")
	fmt.Println("   word-freq filename")
}

func main() {
	flag.Usage = usageText
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Missing filename")
		usageText()
		return
	}

	filename := args[0]

	// read input file
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error readinging file: ", filename)
		panic(err)
	}

	var tree * Node

	convertBuffer(&buf)

	// split on space and join to any previous partwords
	var words *[][]byte
	words = splitBuffer(&buf)

	// add words in slice array to tree, increment count if already exists
	for _, word := range *words {
		if TreeContains(tree, word) == false {
			var data Data
			data.word = append(data.word, word...)
			data.count = 1
			tree = Add(tree, data)
		}
	}

	// display 20 most common
	freqList := [20]Data{}
	FindTwentyMostCommon(tree, &freqList)

	for _, v := range freqList {
		fmt.Println(v.count, string(v.word))
	}
}


// split buffer into words
func splitBuffer(buf *[]byte) *[][]byte {
	wordsl := make([][]byte, 0)
	if buf == nil {
		return &wordsl
	}
	bptr := buf

	start := 0
	for i, c := range *bptr {

		if c == ' ' {
			if start == i {
				// found a space on its own
				start++
				continue
			}
			// found word break
			sl := (*bptr)[start: i]
			start = i+1
			// append word slice to slice of words
			wordsl = append(wordsl, sl)
		}
	}
	return &wordsl
}

// convert buffer chars, clear any unused buffer
func convertBuffer(buf *[]byte) {
	if buf == nil {
		return
	}
	l := len(*buf)
	if l == 0 {
		return
	}

	for i, c := range *buf {
		(*buf)[i] = convertChar(c)
	}

}

// remove non ascii chars and punctuation, convert upper case to lower
func convertChar(c byte) byte {
	switch {
	case 'A' <= c && c <= 'Z':
		// convert to lower case
		return c + ' '
	case 'a' <= c && c <= 'z':
		// pass lower case thru
		return c
	default:
		// all other chars to space
		return ' '
	}
	return 0
}
