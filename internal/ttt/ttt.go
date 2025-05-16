package ttt

type Test struct {
	windowWidth  int32
	windowHeight int32
}

func NewTest(windowWidth, windowHeight int32) *Test {
	return &Test{
		windowWidth:  windowWidth,
		windowHeight: windowHeight,
	}
}
