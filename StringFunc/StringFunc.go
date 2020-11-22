package main

import (
	"fmt"
	s "strings"
)

// 이..이런 방법이 있었군!!! 왜 생각못했을까..
var fp = fmt.Println

func main() {
	fp(`===== STRING FUNC =====`)

	fp("Contains:  ", s.Contains("test", "es"))        // true
	fp("Count:     ", s.Count("test", "t"))            // 2
	fp("HasPrefix: ", s.HasPrefix("test", "te"))       // true
	fp("HasSuffix: ", s.HasSuffix("test", "st"))       // true
	fp("Index:     ", s.Index("test", "e"))            // 1
	fp("Join:      ", s.Join([]string{"a", "b"}, "-")) // a-b
	fp("Repeat:    ", s.Repeat("a", 5))                // aaaaa
	fp("Replace:   ", s.Replace("foo", "o", "0", -1))  // f00
	fp("Replace:   ", s.Replace("foo", "o", "0", 1))   // f0o
	fp("Split:     ", s.Split("a-b-c-d-e", "-"))       // [a,b,c,d,e]
	fp("ToLower:   ", s.ToLower("TEST"))               // test
	fp("ToUpper:   ", s.ToUpper("test"))               // TEST
	fp()

	fp("Len: ", len("hello")) // 5
	fp("Char:", "hello"[1])   // e
}
