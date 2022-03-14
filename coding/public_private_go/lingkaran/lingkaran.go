package lingkaran

type Hitung interface {
	JariJari() float64
}

type hitung struct {
	diameter float64
}

func NewHitung(parameterDiameter float64) *hitung {
	return &hitung{diameter: parameterDiameter}
}

func (h *hitung) JariJari() float64 {
	return h.diameter / 2
}
