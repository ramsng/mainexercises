package word

import (
	"fmt"
	"testing"

	"github.com/ramsng/mainexercises/Handson6/254.ex/quote"
)

var xs = []string{"Hello,", "Bar", "Nice,", "Sweet", "Candy", "Honey"}
var xi = []int{1, 1, 1, 2, 4, 1}

func ExampleCount() {
	fmt.Println(Count(quote.SunAlso))
	// Output:
	// 10
}
func TestWordcount(t *testing.T) {
	m := make(map[string]int)
	for a, b := range xs {
		m[b] = xi[a]
	}
	n := Wordcount(quote.SunAlso)
	for a, b := range m {
		if b != n[a] {
			t.Error("For ", a, "expected value :", m[a], "received output :", n[a])
		}
	}
}
func TestWrdcount(t *testing.T) {
	m := make(map[string]int)
	xi := []int{1, 1, 1, 2, 4, 1}
	for a, b := range xs {
		m[b] = xi[a]
	}
	n := Wrdcount(quote.SunAlso)
	for a, b := range m {
		if b != n[a] {
			t.Error("For ", a, "expected value :", m[a], "received output :", n[a])
		}
	}
}
func BenchmarkWrdcount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Wrdcount(quote.SunAlso)
	}
}
func BenchmarkWordcount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Wordcount(quote.SunAlso)
	}
}
