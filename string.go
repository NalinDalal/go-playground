package main

import (
	"fmt"
	"unicode/utf8"
)

// examineRune demonstrates comparing rune values
func examineRune(r rune) {
	if r == 'ส' {
		fmt.Println("found so sua")
	} else if r == 't' {
		fmt.Println("found tee")
	} else {
		fmt.Println("rune:", string(r))
	}
}

func main() {
	// ----------------------
	// Strings: UTF-8 literals
	// ----------------------
	const s = "สวัสดี" // "hello" in Thai
	fmt.Println("String:", s)

	// ----------------------
	// Length in bytes
	// ----------------------
	fmt.Println("Len (bytes):", len(s)) // 18

	// ----------------------
	// Indexing gives bytes
	// ----------------------
	fmt.Println("Bytes (hex):")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// ----------------------
	// Count runes
	// ----------------------
	fmt.Println("Rune count:", utf8.RuneCountInString(s)) // 6

	// ----------------------
	// Iterate runes with range
	// ----------------------
	fmt.Println("\nRange iteration:")
	for idx, r := range s {
		fmt.Printf("%#U starts at byte %d\n", r, idx)
	}

	// ----------------------
	// Iterate runes with utf8.DecodeRuneInString
	// ----------------------
	fmt.Println("\nDecodeRuneInString iteration:")
	for i, w := 0, 0; i < len(s); i += w {
		r, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at byte %d\n", r, i)
		w = width

		// Examine rune
		examineRune(r)
	}

	// ----------------------
	// Rune literals
	// ----------------------
	var r1 rune = 'ส'
	var r2 rune = 't'
	fmt.Printf("\nRune literals: %c (%#U), %c (%#U)\n", r1, r1, r2, r2)

	// ----------------------
	// Example with non-ASCII symbol
	// ----------------------
	const cmd = "⌘"
	fmt.Println("\nNon-ASCII string:", cmd)
	fmt.Printf("Quoted: %+q\n", cmd)
	fmt.Printf("Hex bytes:")
	for i := 0; i < len(cmd); i++ {
		fmt.Printf(" %x", cmd[i])
	}
	fmt.Println()
}

