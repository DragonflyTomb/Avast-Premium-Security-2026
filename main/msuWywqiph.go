package main
import (
	"fmt"
	"strings"
)
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
func toUpper(s string) string {
	return strings.ToUpper(s)
}
func toLower(s string) string {
	return strings.ToLower(s)
}
func isPalindrome(s string) bool {
	cleaned := strings.ToLower(strings.Join(strings.Fields(s), ""))
	return cleaned == reverse(cleaned)
}
func main() {
	input := "A man a plan a canal Panama"
	fmt.Println("Original:", input)
	fmt.Println("Reversed:", reverse(input))
	fmt.Println("Uppercase:", toUpper(input))
	fmt.Println("Lowercase:", toLower(input))
	fmt.Println("Is Palindrome:", isPalindrome(input))
	input2 := "Hello"
	fmt.Println("Original:", input2)
	fmt.Println("Reversed:", reverse(input2))
	fmt.Println("Uppercase:", toUpper(input2))
	fmt.Println("Lowercase:", toLower(input2))
	fmt.Println("Is Palindrome:", isPalindrome(input2))
}
