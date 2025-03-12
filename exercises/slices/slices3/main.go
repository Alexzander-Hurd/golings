// slices3
// Make me compile!

package main

import "fmt"

func main() {
	names := []string{"John", "Maria", "Carl", "Peter"}
	names = append(names[:len(names)], "Jimmy")
	fmt.Println(names)
}
