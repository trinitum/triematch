package triematch

import (
	"testing"
)

func TestAddStrings(t *testing.T) {
	trie := NewTrie()

	err := trie.AddString("foo", 1)
	if err != nil {
		t.Errorf("Couldn't add foo to the empty trie")
	}

	err = trie.AddString("foo", 1)
	if err == nil {
		t.Errorf("Added foo second time without getting an error")
	}

	err = trie.AddString("fooBar", 2)
	if err != nil {
		t.Errorf("Couldn't add fooBar to the trie")
	}

	err = trie.AddByteString([]byte{0, 1, 2, 3}, 2)
	if err != nil {
		t.Errorf("Couldn't add bytes to the trie")
	}
}

func cmpMatch(a interface{}, b interface{}) bool {
	return a == b
}

func TestLongestMatch(t *testing.T) {
	trie := NewTrie()

	trie.AddString("foo", 1)
	trie.AddString("fooBar", 2)
	trie.AddString("fooBarBaz", 3)

	trie.AddString("bar", 4)

	trie.AddByteString([]byte{1, 2, 3, 0, 1}, 5)
	trie.AddByteString([]byte{1, 2, 3, 0, 1, 2, 3}, 6)

	res := trie.LongestMatch("fooBar")
	if ! cmpMatch(res, 2) {
		t.Errorf("match for fooBar is %v", res)
	}

	res = trie.LongestMatch("fooBarBa")
	if ! cmpMatch(res, 2) {
		t.Errorf("match for fooBarBa is %v", res)
	}

	res = trie.LongestMatch("fooBarBazuka")
	if ! cmpMatch(res, 3) {
		t.Errorf("match for fooBarBazuka is %v", res)
	}

	res = trie.LongestMatch("barber")
	if ! cmpMatch(res, 4) {
		t.Errorf("match for barber is %v", res)
	}

	res = trie.LongestMatch("AfooBarBa")
	if ! cmpMatch(res, nil) {
		t.Errorf("Unexpected match for AfooBarBa: %v", res)
	}

	res = trie.LongestByteMatch([]byte{1, 2, 3, 0, 1, 2, 4, 5, 6})
	if ! cmpMatch(res, 5) {
		t.Errorf("Longest byte match for 1,2,3,0,1,2,4,5,6 is %v", res)
	}
}

func TestManyPatterns(t *testing.T) {
	trie := NewTrie()

	trie.AddString("aoo", 1)
	trie.AddString("boo", 2)
	trie.AddString("coo", 3)
	trie.AddString("doo", 4)
	trie.AddString("eoo", 5)
	trie.AddString("foo", 6)
	trie.AddString("goo", 7)
	trie.AddString("hoo", 8)
	trie.AddString("ioo", 9)
	trie.AddString("joo", 10)
	trie.AddString("koo", 11)

	res := trie.LongestMatch("koo")
	if ! cmpMatch(res, 11) {
		t.Errorf("match for koo is %v", res)
	}
}
