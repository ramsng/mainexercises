package main

import (
	"fmt"
	"testing"
)

type input struct {
	test   []int64
	testop int
}

func TestAverage(t *testing.T) {

	test01 := input{
		test:   []int64{40, 50, 60},
		testop: 55,
	}

	test02 := input{
		test:   []int64{2, 15, 20, 3},
		testop: 10,
	}

	test03 := input{
		test:   []int64{400, 500, 300, 900},
		testop: 525,
	}

	if a := Average(test01.test); test01.testop == int(a) {
		fmt.Println("success")
	} else {
		t.Error("input values:", test01.test, "expected output:", test01.testop, "actual output :", int(a))
	}

	if a := Average(test02.test); test02.testop == int(a) {
		fmt.Println("success")
	} else {
		t.Error("input values:", test02.test, "expected output:", test02.testop, "actual output :", int(a))
	}
	if a := Average(test03.test); test03.testop == int(a) {
		fmt.Println("input values:", test03.test, "expected output:", test03.testop, "actual output :", int(a))
	} else {
		t.Error("input values:", test03.test, "expected output:", test03.testop, "actual output :", int(a))
	}
}
