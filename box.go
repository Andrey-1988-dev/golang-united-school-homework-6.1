package golang_united_school_homework

import (
	"errors"
	"reflect"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return errors.New("out of the shapesCapacity range")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	var shape Shape
	if len(b.shapes) <= i {
		return shape, errors.New("doesn't exist")
	}
	shape = b.shapes[i]
	return shape, nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	var shape Shape
	if len(b.shapes) <= i {
		return shape, errors.New("doesn't exist")
	} else if i > b.shapesCapacity {
		return shape, errors.New("index went out of the range")
	}
	shape = b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	var oldShape Shape
	if len(b.shapes) <= i {
		return oldShape, errors.New("doesn't exist")
	} else if i > b.shapesCapacity {
		return oldShape, errors.New("index went out of the range")
	}
	oldShape = b.shapes[i]
	b.shapes[i] = shape
	return oldShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for _, shape := range b.shapes {
		sum += shape.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for _, shape := range b.shapes {
		sum += shape.CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var newShaps []Shape
	for _, shape := range b.shapes {
		if reflect.TypeOf(shape).String() != "*golang_united_school_homework.Circle" {
			newShaps = append(newShaps, shape)
		}
	}
	if len(newShaps) == len(b.shapes) {
		return errors.New("circles are not exist in the list")
	}
	b.shapes = newShaps
	return nil
}
