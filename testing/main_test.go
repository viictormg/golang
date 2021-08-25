package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(2, 2)
	if total != 4 {
		t.Errorf("Sum was incorrect, got  %d expected %d", total, 4)
	}

	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{5, 5, 10},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrect, got  %d expected %d", total, item.n)

		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
		{5, 2, 5},
		{7, 2, 7},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.n {
			t.Errorf("Get max was incorrect, got %d expected %d", max, item.n)
		}
	}
}

func TestFib(t *testing.T) {
	tables := []struct {
		n int
		r int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}
	for _, item := range tables {
		fib := Fibonacci(item.n)
		if fib != item.r {
			t.Errorf("Get max was incorrect, got %d expected %d", fib, item.n)
		}
	}
}
