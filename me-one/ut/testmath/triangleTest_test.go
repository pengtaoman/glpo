package testmath

import (
	"testing"
)

func sTriangle(t *testing.T) {
	test := []struct{ a, b, c int }{
		{3, 4, 5},
		{6, 8, 1},
		{10, 24, 26},
		{12, 35, 37},
	}

	for _, tt := range test {
		actual := CalcTriangle(tt.a, tt.b)
		println("!!!!!!!!!!!!!!!!!!!!", actual)
		if actual != tt.c {
			t.Errorf("计算三角形弦错误(%d, %d);"+"got %d; expect %d",
				tt.a, tt.b, actual, tt.c)
		}
	}
}
