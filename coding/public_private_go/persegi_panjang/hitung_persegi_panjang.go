package persegipanjang

type Hitung interface {
	Luas() float64
	Keliling() float64
}

type hitung struct {
	lebar   float64
	panjang float64
}

func NewHitung(parameterLebar float64, parameterPanjang float64) *hitung {
	return &hitung{lebar: parameterLebar, panjang: parameterPanjang}
}

func (h *hitung) Luas() float64 {
	return h.panjang * h.lebar
}

func (h *hitung) Keliling() float64 {
	return (2 * h.panjang) + (2 * h.lebar)
}
