package tonogo

import "testing"

//TestHitungKeliling
func TestCalc(t *testing.T) {
	t.Logf("Keliling : %d", Calc(4, 2))

	if Calc(4, 2) != 6 {
		t.Errorf("SALAH! harusnya %d", 11)
	}
}

func BenchmarkCalc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calc(4, 2)
	}
}
