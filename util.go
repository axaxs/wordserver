package main

import (
	"io/ioutil"
	"strings"
)

var allWords = []string{}

func wordsFromFile(filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	for _, v := range strings.Split(string(b), "\n") {
		if v == "" {
			continue
		}
		if !onlyLowerChars(v) {
			continue
		}
		allWords = append(allWords, v)
	}

	return nil
}

func onlyLowerChars(s string) bool {
	for _, c := range s {
		if c < 97 || c > 122 {
			return false
		}
	}
	return true
}

func containsSlice(sl []string, s string) bool {
	for _, v := range sl {
		if v == s {
			return true
		}
	}
	return false
}

func containsString(in string) []string {
	res := make([]string, 0, 50)
	for _, v := range allWords {
		if strings.Contains(v, in) {
			res = append(res, v)
		}
	}
	return res
}
