package persegi

import "math"

type Hitung interface {
	Luas() float64
	Keliling() float64
}

type hitung struct {
	sisi float64
}

func NewHitung(nilai float64) *hitung {
	return &hitung{sisi: nilai}
}

func (p *hitung) Luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p *hitung) Keliling() float64 {
	return p.sisi * 4
}
