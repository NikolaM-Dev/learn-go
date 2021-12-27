package main

import (
	"fmt"
	"unicode/utf8"
)

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printCharsOld(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func charsAndBytePosition(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func compareStrings(str1, str2 string) {
	if str1 == str2 {
		fmt.Printf("%s and %s are equal\n", str1, str2)
		return
	}
	fmt.Printf("%s and %s are not equal\n", str1, str2)
}

func mutate(s []rune) string {
	s[0] = 'a' // any valid unicode character within single quote is a rune
	return string(s)
}

func main() {
	// Accessing individual bytes of a string
	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printBytes(name)

	// Accessing individual characters of a string
	fmt.Printf("\n")
	printCharsOld(name)
	fmt.Printf("\n")
	printBytes(name)
	fmt.Printf("\n\n")

	// Although the above program looks like a legitimate way to access
	// the individual characters of a string, this has a serious bug.
	// Let's find out what that bug is.
	name = "Señor"
	fmt.Printf("String: %s\n", name)
	printCharsOld(name)
	fmt.Printf("\n")
	printBytes(name)

	// Rune
	// A rune is a builtin type in Go and it's the alias of int32.
	// Rune represents a Unicode code point in Go. It doesn't matter
	// how many bytes the code point occupies, it can be represented
	// by a rune. Let's modify the above program to print characters
	// using a rune.
	fmt.Printf("\n\n")
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)

	// Accessing individual runes using for range loop
	fmt.Printf("\n\n")
	charsAndBytePosition(name)

	// Creating a string from a slice of bytes
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9} // hexadecimales
	byteSliceDecimal := []byte{67, 97, 102, 195, 169} // decimal equivalent
	str := string(byteSlice)
	str2 := string(byteSliceDecimal)
	fmt.Println(str)
	fmt.Println(str2)

	// String length
	word1 := "Señor"
	fmt.Printf("String: %s\n", word1)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word1))
	fmt.Printf("Number of bytes: %d\n", len(word1))

	fmt.Printf("\n")
	word2 := "Pets"
	fmt.Printf("String: %s\n", word2)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word2))
	fmt.Printf("Number of bytes: %d\n", len(word2))

	// String comparison
	string1 := "Go"
	string2 := "Go"
	compareStrings(string1, string2)

	string3 := "hello"
	string4 := "world"
	compareStrings(string3, string4)

	// String concatenation
	str3 := "Go"
	str4 := "is awesome"
	result := str3 + " " + str4
	fmt.Println(result)

	result2 := fmt.Sprintf("%s %s", str3, str4)
	fmt.Println(result2)

	// Strings are immutable
	h := "hello"
	fmt.Println(mutate([]rune(h)))
}
