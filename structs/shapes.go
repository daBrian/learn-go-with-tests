package structs

type Rectangle struct {
	width  float64
	height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

func Area(rectangle Rectangle) interface{} {
	return rectangle.width * rectangle.height
}
