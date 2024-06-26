package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{12, 6}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("Got '%.2f' want '%.2f'.", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("Got '%g' want '%g'.", got, tt.want)
		}
	}
}
