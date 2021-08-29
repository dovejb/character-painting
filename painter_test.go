package painting

import (
	"testing"
)

func TestString(t *testing.T) {
	s := "hello world"
	ss := String(s)
	for i := range ss {
		t.Log(ss[i])
	}
}

func TestRainbow(t *testing.T) {
	s := "hello world"
	ss := Rainbow(s)
	for i := range ss {
		t.Log(ss[i])
	}
}

func TestRainbowStagger(t *testing.T) {
	s := "hello world"
	ss := RainbowStagger(s)
	for i := range ss {
		t.Log(ss[i])
	}
}

func TestColor(t *testing.T) {
	s := "hello world"
	ss := Color(s, COLOR_GREEN)
	for i := range ss {
		t.Log(ss[i])
	}
}

func TestColorLoop(t *testing.T) {
	s := "hello world"
	ss := ColorLoop(s, []string{COLOR_GREEN, COLOR_RED, COLOR_YELLOW, COLOR_BLUE, COLOR_CYAN})
	for i := range ss {
		t.Log(ss[i])
	}
}
