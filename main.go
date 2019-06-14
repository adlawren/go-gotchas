package main

// Working through: https://yourbasic.org/golang/gotcha-assignment-entry-nil-map/

import (
	"fmt"
	"math"
	"strings"
	"time"
)

//// #2
type Point struct {
	X, Y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

//// #4
// func Foo(a [2]int) {
// 	a[0] = 8
// }
func Foo(a []int) {
	a[0] = 8
}

func main() {
	fmt.Println("#1")
	{
		//// 'panic: assignment to entry in nil map'
		// var m map[string]float64
		// m["pi"] = 3.1416
		//// Solution: create the map with make (or using a literal)
		m := make(map[string]float64)
		m["pi"] = 3.1416
	}

	fmt.Println("#2")
	{
		//// 'panic: runtime error: invalid memory address or nil pointer dereference'
		//var p *Point
		//fmt.Println(p.Abs())
		//// Solution: initialize the pointer
		var p *Point = &Point{} // or: new(Point)
		fmt.Println(p.Abs())
	}

	fmt.Println("#3")
	{
		//// 'multiple-value time.Parse() in single-value context'
		// t := time.Parse(time.RFC3339, "2018-04-06T10:49:05Z")
		// fmt.Println(t)
		//// Solution: assign the second return value (error)
		t, _ := time.Parse(time.RFC3339, "2018-04-06T10:49:05Z")
		fmt.Println(t)
	}

	fmt.Println("#4")
	{
		//// a[0] doesn't change, even though reassigned in Foo
		// a := [2]int{1, 2}
		// Foo(a)         // Try to change a[0].
		// fmt.Println(a) // Output: [1 2]
		//// Solution: arrays are values, slices are not
		a := []int{1, 2}
		Foo(a)         // Try to change a[0].
		fmt.Println(a) // Output: [1 2]
	}

	fmt.Println("#5")
	{
		//// n doesn't change
		// n := 0
		// if true {
		// 	n := 1
		// 	n++
		// }
		// fmt.Println(n) // 0
		//// Solution: use '=' in if, not ':='
		n := 0
		if true {
			n = 1
			n++
		}
		fmt.Println(n) // 0

		//// Note:
		// You can detect shadowed variables by running:
		// go vet -shadow main.go
		//// In Go 1.12:
		// go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow
		// go vet -vettool=$(which shadow)
	}

	fmt.Println("#6")
	{
		//// 'syntax error: unexpected newline, expecting comma or }'
		// fruit := []string{
		// 	"apple",
		// 	"banana",
		// 	"cherry"
		// }
		// fmt.Println(fruit)
		//// Solution: add comma after final array item
		fruit := []string{
			"apple",
			"banana",
			"cherry",
		}
		fmt.Println(fruit)

		//// Note:
		// In multiline slice, array & map literals, every line must end with a comma
	}

	fmt.Println("#7")
	{
		//// 'cannot assign to s[0]'
		// s := "hello"
		// s[0] = 'H'
		// fmt.Println(s)
		//// Solution: use a rune slice
		s := []rune("hello")
		s[0] = 'H'
		fmt.Println(string(s))

		//// Note:
		// Strings are immutable; they're like read-only byte slices
	}

	fmt.Println("#8")
	{
		//// Different results despite equivalent (?) types
		fmt.Println("H" + "i")
		fmt.Println('H' + 'i')
		//// Solution:
		// On statement is adding strings, the other, *characters
		// *runes, actually
		// Turn the runes into strings before concatenation, or use Sprintf
	}

	fmt.Println("#9")
	{
		//// Output: ""
		fmt.Println(strings.TrimRight("ABBA", "BA"))
		//// Solution: ...
		fmt.Println(strings.TrimRight("ABBA", "BA"))
	}
}
