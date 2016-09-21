// Package triematch provides a simple functionality to check if string matches
// one of the patterns.
package triematch

import (
	"errors"
	"fmt"
)

type Trie struct {
	match    interface{}
	outgoing []*Trie
	outCount uint16
	alpha    uint8
}

// NewTrie creates a new trie structure and returns it
func NewTrie() (n *Trie) {
	return &Trie{
		outgoing: make([]*Trie, 0, 8),
	}
}

// AddByteString method adds a new pattern to the trie. match is an arbitrary
// value returned if the string matches this pattern
func (n *Trie) AddByteString(str []byte, match interface{}) error {
	nn := n
	for _, b := range str {
		nn = nn.getChildNode(b)
	}
	if nn.match != nil {
		return errors.New("The pattern already exists!")
	}
	nn.match = match

	return nil
}

// AddString method adds a new pattern to the trie. Same as AddByteString, but
// the first argument is string and not []byte
func (n *Trie) AddString(str string, match interface{}) error {
	return n.AddByteString([]byte(str), match)
}

// LongestByteMatch method finds the longest pattern that the string matches
// and returns assosiated value. Note, that the pattern should start at the
// beginning of the string.
func (n *Trie) LongestByteMatch(str []byte) (res interface{}) {
	nn := n
	res = n.match
	for _, b := range str {
		nn = nn.findChildNode(b)
		if nn == nil {
			return res
		}
		if nn.match != nil {
			res = nn.match
		}
	}
	return res
}

// LongestMatch is the same as LongestByteMatch, but accepts string instead of
// []byte
func (n *Trie) LongestMatch(str string) interface{} {
	return n.LongestByteMatch([]byte(str))
}

func (n *Trie) addChildNode(nn *Trie) {
	n.outgoing = append(n.outgoing, nn)
	i := n.outCount
	n.outCount++
	for i > 0 && n.outgoing[i].alpha < n.outgoing[i-1].alpha {
		n.outgoing[i], n.outgoing[i-1] = n.outgoing[i-1], n.outgoing[i]
		i--
	}
}

func (n *Trie) dump(indent int) {
	for _, nn := range n.outgoing {
		fmt.Printf("%*s%d (%c)\n", indent, "", nn.alpha, nn.alpha)
		nn.dump(indent+1)
	}
}

func (n *Trie) findChildNode(b byte) *Trie {
	if n.outCount == 0 {
		return nil
	}
	min := int32(0)
	max := int32(n.outCount - 1)

	for min <= max {
		mid := (min + max) / 2
		if n.outgoing[mid].alpha > b {
			max = mid - 1
		} else if n.outgoing[mid].alpha < b {
			min = mid + 1
		} else {
			return n.outgoing[mid]
		}
	}

	return nil
}

func (n *Trie) getChildNode(b byte) *Trie {
	nn := n.findChildNode(b)
	if nn == nil {
		nn = &Trie{
			outgoing: make([]*Trie, 0, 8),
			alpha:    b,
		}
		n.addChildNode(nn)
	}

	return nn
}
