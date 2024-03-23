package main

import (
	"bytes"
	"fmt"
)

func ensureNullTerminatedKey(key []byte) []byte {
	index := bytes.Index(key, []byte{0})
	if index < 0 {
		key = append(key, byte(0))

	}
	return key
}

func main() {
	n1 := []byte("test")
	n2 := []byte("tests")

	// for i := 0; i < len(n1); i++ {
	// 	for j := 0; j < len(n1); i++ {
	// 	}
	// }

	fmt.Printf("%08b\n", []byte(n1))
	fmt.Printf("%08b\n", []byte(n2))

	temp := ensureNullTerminatedKey(n2)
	fmt.Printf("%d\n", temp)
}
