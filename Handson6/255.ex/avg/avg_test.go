package avg

import (
	"fmt"
	"testing"
)

func TestCenteredavg(t *testing.T) {
	xi := []int{1002, 2002, 3002}
	var expected float64 = 2002
	result := Centeredavg(xi)
	if expected != result {
		t.Error("expected value :", expected, "Output :", result)
	}
}

func BenchmarkCenteredavg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Centeredavg([]int{5002, 2002, 3002, 1002})
	}
}

func ExampleCenteredavg() {
	fmt.Println(Centeredavg([]int{1002, 2002, 3002, 4002}))
	//Output
	//2502
}
