package data

type Color struct {
	Red   int
	Green int
	Blue  int
}

type HEX struct {
	Value string
}

type Cycle struct {
	PinRcycle int
	PinGcycle int
	PinBcycle int
}

type RGBLinker struct {
	RGB          Color
	Freq         int
	PinRGBlayout PinRGBlayout
	Cycle        Cycle
}

type PinRGBlayout struct {
	PinR int
	PinG int
	PinB int
}
