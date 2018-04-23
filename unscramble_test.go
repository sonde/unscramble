package main

import (
	"testing"
)

func Test_unscramble(t *testing.T) {
	tables := []struct {
		scrambled string
		dict      []wordByLen
		words     string
		ok        bool
	}{
		{
			"elhloothtedrowl",
			[]wordByLen{
				{0, map[string]string{}},
				{1, map[string]string{}},
				{2, map[string]string{"ot": "to"}},
				{3, map[string]string{"eht": "the"}},
				{4, map[string]string{}},
				{5, map[string]string{"ehllo": "hello", "dlorw": "world"}},
			},
			"hello to the world",
			true,
		},
		{
			"iamacatnotadog",
			[]wordByLen{
				{0, map[string]string{}},
				{1, map[string]string{"i": "i", "a": "a"}},
				{2, map[string]string{"am": "am"}},
				{3, map[string]string{"not": "not", "dgo": "dog", "act": "cat"}},
			},
			"i am a cat not a dog",
			true,
		},
		{
			"elhloothtedrowl",
			[]wordByLen{},
			"othello worthed",
			false,
		},
		{
			"csurlmnnbagiteennesc",
			[]wordByLen{},
			"unscrambling sentence",
			false,
		},
	}
	for _, table := range tables {

		if len(table.dict) == 0 { // Read the dict if dict is uninitialized
			table.dict = *readDict("dict.large")
		}

		words, ok := unscramble(table.scrambled, &table.dict)
		if table.words != words {
			t.Errorf("Wanted: %s, %t; Got %s, %t\n", table.words, table.ok, words, ok)
		}
	}
}
