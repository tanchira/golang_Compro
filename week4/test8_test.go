package test1

import "testing"

func TestSum(t *testing.T) {
	t.Run("Hello World", func(t *tesing.T) {
		want := true
		got := Sum("Hello World")
		if want != got {
			t.Error("want true but got false", want, got)
		}

	})
	t.Run("Happy birthday", func(t *testing.T) {
		want := false
		got := Sum("Happy birthday")
		if want != got {
			t.Error("want false but got true", want, got)
		}

	})

}