package area

import "testing"

func TestAreaReactangle(t *testing.T) {
	rec := Rectangle{Width: 10, Height: 10}
	got := rec.Area()
	var expected float32 = 100.0

	if got != expected {
		t.Errorf("Expected %f got %f", expected, got)
	}
}

func TestCircleArea(t *testing.T) {
	circ := Circle{Radius: 2}
	got := circ.Area()
	var expected float32 = 12.566371
	if got != expected {
		t.Errorf("Expected %f got %f", expected, got)
	}
}
