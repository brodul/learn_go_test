package perimeter

import "testing"

func TestPerimiter(t *testing.T) {

	t.Run("perimiter of a rectengle",
		func(t *testing.T) {
			rectangle := Rectangle{10.0, 10.0}
			got := Perimeter(rectangle)
			want := 40.0

			if got != want {
				t.Errorf("got %.2f want %.2f ", got, want)
			}
		})
	// t.Run("perimiter of a circle",
	// 	func(t *testing.T) {
	// 		circle := Circle{10.0}
	// 		got := Perimeter(circle)
	// 		want := 314.1592653589793

	// 		if got != want {
	// 			t.Errorf("got %.2f want %.2f ", got, want)
	// 		}
	// 	})

}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f ", got, want)
		}

	}
	t.Run("area of a rectengle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})
	t.Run("area of a circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})

}
