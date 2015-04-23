package vector_test

import (
	"testing"

	. "github.com/outofjungle/gotest/vector"
)

var val = 0.70710678118655

func TestQ1Q1(t *testing.T) {
	p1 := Point{X: val, Y: val}
	p2 := Point{X: val, Y: val}

	expected := 1.000000000000007
	if scalar := p1.Dot(p2); scalar != expected {
		t.Errorf("%v • %v returned %f, expected %f", p1, p2, scalar, expected)
	}
}

func TestQ1Q2(t *testing.T) {
	p1 := Point{X: val, Y: val}
	p2 := Point{X: -val, Y: val}

	expected := 0.0
	if scalar := p1.Dot(p2); scalar != expected {
		t.Errorf("%v • %v returned %f, expected %f", p1, p2, scalar, expected)
	}
}

func TestQ1Q3(t *testing.T) {
	p1 := Point{X: val, Y: val}
	p2 := Point{X: -val, Y: -val}

	expected := -1.000000000000007
	if scalar := p1.Dot(p2); scalar != expected {
		t.Errorf("%v • %v returned %f, expected %f", p1, p2, scalar, expected)
	}
}

func TestQ1Q4(t *testing.T) {
	p1 := Point{X: val, Y: val}
	p2 := Point{X: val, Y: -val}

	expected := 0.0
	if scalar := p1.Dot(p2); scalar != expected {
		t.Errorf("%v • %v returned %f, expected %f", p1, p2, scalar, expected)
	}
}
