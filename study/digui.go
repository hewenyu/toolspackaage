package study

import "fmt"

func reverseString(s []byte) []byte {

	var (
		fix  int
		back []byte
	)

	fix = len(s)
	back = make([]byte, 0)

	for k := range s {

		am := fix - k - 1
		back = append(back, s[am])

	}
	return back
}

func main() {
	a := []byte("hello")

	fmt.Println(reverseString(a))
}
