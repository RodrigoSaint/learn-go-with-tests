package area

import "testing"

func compareArea(t *testing.T, shape Shape, expected float32) {
	got := shape.Area()

	if got != expected {
		t.Errorf("%#v area Expected %f got %f", shape, expected, got)
	}
}

func TestAreaReactangle(t *testing.T) {
	rec := Rectangle{Width: 10, Height: 10}
	var expected float32 = 100.0
	compareArea(t, rec, expected)
}

func TestCircleArea(t *testing.T) {
	circ := Circle{Radius: 2}
	var expected float32 = 12.566371
	compareArea(t, circ, expected)
}
