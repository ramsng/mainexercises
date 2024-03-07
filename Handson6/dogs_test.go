package dogs

import (
	"fmt"
	"testing"
)

func ExampleYears1(t *testing.T) {
	fmt.Println(Years1(10))
	//Output:
	//100
}

func BenchmarkYears1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years1(10)
	}
}
