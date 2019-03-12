// Part of the program to find the top most common words in a text file
package main

import (
	"sort"
)

type Data struct {
	word  []byte
	count int
}

type Node struct {
	value  Data
	pLeft  *Node
	pRight *Node
}

func Add(tree *Node, val Data) *Node {

	tempNodePtr := new(Node)
	tempNodePtr.value = val
	tempNodePtr.pLeft = nil
	tempNodePtr.pRight = nil
	return addNode(tree, tempNodePtr)
}

func addNode(tree *Node, toAdd *Node) *Node {
	if tree == nil {
		return toAdd
	} else {
		comp := ByteCompare(toAdd.value.word, tree.value.word)
		if comp == 1 {
			tree.pLeft = addNode(tree.pLeft, toAdd)
		} else if comp == 0 {
			tree.value.count++
		} else {
			tree.pRight = addNode(tree.pRight, toAdd)
		}
		return tree
	}
}

func TreeContains(tree *Node, word []byte) bool {
	// if parameter word is in the binary tree
	// then true is returned.
	// else false is returned

	if tree == nil {
		// Tree is empty
		return false
	} else if ByteCompare(word, tree.value.word) == 0 {
		//The word matches to one in the root node.
		tree.value.count++
		return true
	} else if ByteCompare(word, tree.value.word) == 1 {
		// The word is less than the one in the root node
		// and must be sent to the left subtree.
		return TreeContains(tree.pLeft, word)
	} else {
		// The word is more than the one in the root node
		// and must be sent to the right subtree.
		return TreeContains(tree.pRight, word)
	}
}

func FindTwentyMostCommon(tree *Node, freqList *[20]Data) {
	if tree != nil {
		FindTwentyMostCommon(tree.pLeft, freqList)
		FindTwentyMostCommon(tree.pRight, freqList)

		addWordToList(freqList, tree.value)
	}
}

func addWordToList(freqList *[20]Data, val Data) {
	// as list is sorted, the bottom of the list (index19) should be the current min
	if (*freqList)[19].count < val.count {
		(*freqList)[19] = val
		sort.Slice(freqList[:], func(i, j int) bool {
			return freqList[i].count > freqList[j].count
		})
	}
}

// compare bytes arrays, return 1 if lhs is greater, -1 if rhs is greater
// return 0 if equal
func ByteCompare(lhs []byte, rhs []byte) int {
	lenlhs := len(lhs)
	lenrhs := len(rhs)
	if lenlhs == 0 && lenrhs == 0 {
		return 0
	}
	if lenlhs == 0 {
		return -1
	}
	if lenrhs == 0 {
		return 1
	}

	maxlen := lenlhs
	if lenlhs < lenrhs {
		maxlen = lenrhs
	}

	for i := 0; i < maxlen; i++ {
		//	for i, val := range lhs {
		if i < lenlhs && i < lenrhs {
			if lhs[i] > rhs[i] {
				return 1
			} else if lhs[i] < rhs[i] {
				return -1
			}
		} else {
			if i >= lenlhs {
				return -1
			} else {
				return 1
			}
		}
	}
	return 0
}
