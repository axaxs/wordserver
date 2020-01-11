package main

import (
	//	"fmt"

	"net/http"
	"sort"
	"strings"

	"github.com/axaxs/wordtrie"
)

var (
	prefixTree *wordtrie.Trie
	suffixTree *wordtrie.Trie
)

func reverseString(in string) string {
	rs := []rune(in)
	n := len(rs)
	for i := 0; i < len(in)/2; i++ {
		rs[i], rs[n-1-i] = rs[n-1-i], rs[i]
	}
	return string(rs)
}

func contains(in string) []string {
	res := make([]string, 0, 20)
	for _, v := range allWords {
		if strings.Contains(v, in) {
			res = append(res, v)
		}
	}
	return res
}

func handler(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	broken := strings.Split(req.URL.Path, "/")
	if len(broken) != 3 {
		w.Write([]byte("Invalid URI\n"))
		return
	}
	match := strings.ToLower(broken[2])
	if !onlyLowerChars(match) {
		w.Write([]byte("Only enter ascii characters\n"))
		return
	}
	var res []string
	switch broken[1] {
	case "anagrams":
		res = anagrams(prefixTree, match)
	case "canmake":
		res = canMake(prefixTree, match)
	case "startswith":
		res = startsWith(prefixTree, match)
	case "endswith":
		res = startsWith(suffixTree, reverseString(match))
		newRes := make([]string, 0, len(res))
		for _, v := range res {
			newRes = append(newRes, reverseString(v))
		}
		res = newRes
	case "contains":
		res = contains(match)
	default:
		res = []string{"Unknown path"}
	}
	if len(res) == 0 {
		w.Write([]byte("No results found\n"))
	}
	sort.Strings(res)
	w.Write([]byte(strings.Join(res, "\n")))
}

func main() {
	err := wordsFromFile("aswordlist.txt")
	if err != nil {
		panic(err)
	}

	prefixTree = wordtrie.NewTrie()
	suffixTree = wordtrie.NewTrie()
	for _, v := range allWords {
		prefixTree.Insert(v)
		suffixTree.Insert(reverseString(v))
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
