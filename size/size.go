package size

type unit int

const (
	Cell unit = iota
	Percent
	Fraction
)

var Auto = Size{Unit: Fraction, Value: 1}

type Size struct {
	Value int
	Unit  unit
}

func WithFraction(value int) Size {
	return Size{Value: value, Unit: Fraction}
}

func WithPercent(value int) Size {
	return Size{Value: value, Unit: Percent}
}

func WithCells(value int) Size {
	return Size{Value: value, Unit: Cell}
}
