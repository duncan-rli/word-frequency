package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
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

	// open input file
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file: ", filename)
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	var tree *Node
	// make a read buffer
	r := bufio.NewReader(fi)

	const bufSize = int(1024)
	var partWord = make([]byte, 0)

	// make a buffer to keep chunks that are read
	buf := make([]byte, bufSize)

	// read chunks from file
	for {
		// read a chunk
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// remove non ascii chars and convert to lower case,
		// changes passed back in same buffer
		convertBuffer(&buf, n)

		// split on space and join to any previous partwords
		var words *[][]byte
		partWord, words = splitBuffer(partWord, &buf)

		// add words in slice array to tree
		for _, word := range *words {
			if CheckTreeContainsAndUpdate(tree, word) == false {
				// word not in tree so add it
				var data Data
				data.word = append(data.word, word...)
				data.count = 1
				tree = Add(tree, data)
			}
		}
	}

	// display 20 most common
	freqList := [20]Data{}
	FindTwentyMostCommon(tree, &freqList)

	for _, v := range freqList {
		fmt.Println(v.count, string(v.word))
	}
}

// split buffer into words and incorporate any part word from the previous buffer
func splitBuffer(part []byte, buf *[]byte) ([]byte, *[][]byte) {
	if buf == nil {
		return nil, nil
	}
	l := len(*buf)
	bptr := buf

	wordsl := make([][]byte, 0)

	// check for part word from previous buffer split, join if required
	if len(part) > 0 {
		l = l + len(part)
		// append buffer to the word part from last time
		part = append(part, *buf...)
		bptr = &part
	}

	start := 0
	var remainingPart []byte

	for i, c := range *bptr {

		if c == ' ' {
			if start == i {
				// found a space on its own
				start++
				continue
			}
			// found word break
			sl := (*bptr)[start:i]
			start = i + 1
			// append word slice to slice of words
			wordsl = append(wordsl, sl)
		}
		if start < (l) && i == l-1 {
			remainingPart = make([]byte, i+1-start)
			copy(remainingPart, (*bptr)[start:i+1])
		}
	}
	return remainingPart, &wordsl
}

// convert buffer chars, clear any unused buffer
func convertBuffer(buf *[]byte, bytesRead int) {
	if buf == nil {
		return
	}
	l := len(*buf)
	if l == 0 {
		return
	}

	for i, c := range *buf {
		if i < bytesRead {
			(*buf)[i] = convertChar(c)
		} else {
			(*buf)[i] = 32 // ' '
		}
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
