package main

import (
	"fmt"
	"strings"
)

func findDivisor(n int) {
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
func ExtractDigit (n int) {
	for n > 0 {
		digit := n % 10
		fmt.Print(digit, " ")
		n = n / 10
	}
	fmt.Println()
}
func TriangleRight(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= i {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
func TriangleLeft(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+j >= n-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func NumberPyramid(n int) {
	for i := 0; i < n; i++ {
		for j := n - i; j >= 1; j-- {
			fmt.Print(j, " ")
		}
		for j := 2; j <= n-i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

func Pattern5a(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j%2 == 1 {
				fmt.Printf("%d ", i)
			} else {
				fmt.Printf("%d ", n-i+1)
			}
		}
		fmt.Println()
	}
}

func pattern6b(n int) {
	for i := 1; i <= n; i++ {
		even := 2
		odd := 1
		for j := 1; j <= n; j++ {
			if i%2 == 1 { // baris ganjil
				if j%2 == 0 {
					fmt.Printf("%d ", even)
					even += 2
				} else {
					fmt.Print("- ")
				}
			} else { // baris genap
				if j%2 == 1 {
					fmt.Printf("%d ", odd)
					odd += 2
				} else {
					fmt.Print("- ")
				}
			}
		}
		fmt.Println()
	}
}

func IsPalindrome(s string) bool {
	clean := ""
	for _, ch := range s {
		if ch != ' ' {
			clean += strings.ToLower(string(ch))
		}
	}

	i, j := 0, len(clean)-1
	for i < j {
		if clean[i] != clean[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func CheckBraces(s string) bool {
	count := 0
	for _, ch := range s {
		if ch == '(' {
			count++
		} else if ch == ')' {
			count--
			if count < 0 {
				return false
			}
		}
	}
	return count == 0
}
func IsNumberPalindrome(n int) bool {
	if n < 0 {
		return false
	}

	original := n
	reversed := 0

	for n > 0 {
		reversed = reversed*10 + (n % 10)
		n /= 10
	}

	return original == reversed
}


func main() {
	//No. 1 INTEGER = findDivisor
	fmt.Print("findDivisor(6)  -> ")
	findDivisor(6)
	fmt.Print("findDivisor(24) -> ")
	findDivisor(24)
	fmt.Print("findDivisor(7)  -> ")
	findDivisor(7)
	//No. 2 ExtractDigit
	fmt.Print ("ExtractDigit(12234) -> ")
	ExtractDigit(12234)
	fmt.Print ("ExtractDigit(5432) -> ")
	ExtractDigit(5432)
	fmt.Print ("ExtractDigit(1278)) -> ")
	ExtractDigit(1278)
	//No. 3 Triangle
	fmt.Print ("TriangleRight 1\n")
	TriangleRight (5)
	fmt.Print ("TriangleLeft 2\n")
	//No. 4 NumberPyramid
	TriangleLeft (5)
	fmt.Print ("\nNumberPyramid\n")
	NumberPyramid (8)
	//No. 5 Pattern
	fmt.Print ("Pattern\n")
	Pattern5a (9)
	fmt.Print ("Pattern\n")
	Pattern5a (5)
	// No. 6 Pattern
	fmt.Print ("Pattern\n")
	pattern6b (9)
	fmt.Print ("Pattern\n")
	pattern6b (5)
	//No. 8 STRING = isPalindrome
	fmt.Print ("isPalindrome(“Kasur ini rusak”)")
	fmt.Print ("sPalindrome(“tamaT”)")
	fmt.Print ("isPalindrome(“Aku Usa”)")
	//No. 9 - reverse
	fmt.Println(`reverse("ABCD")   =`, reverse("ABCD"))
	fmt.Println(`reverse("tamaT")  =`, reverse("tamaT"))
	fmt.Println(`reverse("XYnb")   =`, reverse("XYnb"))
	// No. 10 — CheckBraces
	fmt.Println(`CheckBraces("(())")       =`, CheckBraces("(())"))
	fmt.Println(`CheckBraces("()()")       =`, CheckBraces("()()"))
	fmt.Println(`CheckBraces("(()")        =`, CheckBraces("(()"))
	fmt.Println(`CheckBraces("(()))(()")   =`, CheckBraces("(()))(()"))
	// No. 11 — IsNumberPalindrome
	fmt.Println("IsNumberPalindrome(121)", IsNumberPalindrome(121))
	fmt.Println("IsNumberPalindrome(2147447412)", IsNumberPalindrome(2147447412))
	fmt.Println("IsNumberPalindrome(333)", IsNumberPalindrome(333))
	fmt.Println("IsNumberPalindrome(110)", IsNumberPalindrome(110))
	fmt.Println("IsNumberPalindrome(11)", IsNumberPalindrome(11))
}
