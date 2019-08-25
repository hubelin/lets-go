package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("test rectange", func(t *testing.T) {
		r := Rectangle{10.0, 10.0}
		got := r.Area()
		want := 100.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("test circle", func(t *testing.T) {
		c := Circle{10.0}
		got := c.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
