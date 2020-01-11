package main

import (
	"strings"

	"github.com/axaxs/wordtrie"
)

func buildChildWords(t *wordtrie.Trie) []string {
	curTrie := t
	words := []string{}
	if curTrie.IsWord {
		words = append(words, curTrie.BuildWord())
	}

	for _, trie := range curTrie.Children {
		words = append(words, buildChildWords(trie)...)
	}

	return words
}

func buildChildWordsContaining(t *wordtrie.Trie, curs string) []string {
	words := []string{}
	curs = strings.Replace(curs, string(t.Chr), "", 1)
	if t.IsWord {
		words = append(words, t.BuildWord())
	}

	for _, trie := range t.Children {
		if strings.Contains(curs, string(trie.Chr)) {
			words = append(words, buildChildWordsContaining(trie, curs)...)
		}
	}

	return words
}

func canMake(t *wordtrie.Trie, s string) []string {
	res := make([]string, 0, 10)
	tested := make(map[rune]bool)
	for _, v := range s {
		if _, ok := tested[v]; ok {
			continue
		}
		tested[v] = true
		parent, _ := t.TrieAt(string(v))
		res = append(res, buildChildWordsContaining(parent, s)...)
	}

	return res
}

func anagrams(t *wordtrie.Trie, s string) []string {
	ana := canMake(t, s)
	res := make([]string, 0, len(ana)/2)
	for _, v := range ana {
		if len(v) == len(s) {
			res = append(res, v)
		}
	}

	return res
}

func startsWith(t *wordtrie.Trie, in string) []string {
	curTrie := t
	for in != "" {
		ok := false
		for _, node := range curTrie.Children {
			if node.Chr == rune(in[0]) {
				ok = true
				curTrie = node
				if len(in) == 1 {
					in = ""
				} else {
					in = in[1:]
				}
				break
			}
		}
		if !ok {
			return []string{}
		}
	}

	return buildChildWords(curTrie)
}
