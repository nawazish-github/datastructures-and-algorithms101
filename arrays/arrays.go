package arrays

import (
	"errors"
)

const initCap = 10

//DynamicArrays represents an elastic array
type DynamicArrays struct {
	Size     int           //Size is the current number of elements in the DynamicArray
	Capacity int           //Capacity total number of elements which can be accomodated into the DynamicArray
	Data     []interface{} //Data is the underlying storage for the DynamicArray
}

//NewArray returns a new DynamicArray initialized to its default values
func NewArray() *DynamicArrays {
	return &DynamicArrays{
		Size:     0,
		Capacity: initCap,
		Data:     make([]interface{}, initCap),
	}
}

//Add adds the given element to the end of the DynamicArray
//If there is not enough space to add, it grows itself.
func (d *DynamicArrays) Add(elem interface{}) {
	if d.Size == d.Capacity {
		enlargeArray(d)
		d.Data[d.Size] = elem
		d.Size++
		return
	}

	d.Data[d.Size] = elem
	d.Size++
}

func enlargeArray(d *DynamicArrays) {
	d.Size = 0
	newCap := d.Capacity * 2
	tempData := make([]interface{}, newCap)
	for idx, val := range d.Data {
		tempData[idx] = val
		d.Size++
	}
	d.Capacity = newCap
	d.Data = tempData
}

//Insert inserts given element at the given index
func (d *DynamicArrays) Insert(elem interface{}, idx int) {
	switch {
	//check if index is empty or occupied
	case idx < d.Size && d.Size < d.Capacity:
		shiftAndPutElement(d, idx, elem)
	//check whether there is any space for insertion
	case idx < d.Size && d.Size == d.Capacity:
		enlargeArray(d)
		//make space by shifting
		shiftAndPutElement(d, idx, elem)
	//check whether space is available and index is not occupied
	case idx >= d.Size && d.Size < d.Capacity && idx < d.Capacity:
		d.Data[idx] = elem
		d.Size++
	default:
		panic(errors.New("Cannot insert beyond the capacity of array"))
	}

}

func shiftAndPutElement(d *DynamicArrays, idx int, elem interface{}) {
	for i := d.Size - 1; i >= idx; i-- {
		d.Data[i+1] = d.Data[i]
	}
	d.Data[idx] = elem
	d.Size++
}

//RotateClockwise rotates the internal array clockwise n times
func (d *DynamicArrays) RotateClockwise(n int) {
	currIdx := len(d.Data) - 1
	for i := 1; i <= n; i++ {
		val := d.Data[currIdx].(int)
		enqueue(val)
		d.Data[currIdx] = d.Data[0]
		for currIdx > 0 {

		}
	}
}

var temp = make([]int, 2)

func enqueue(num int) {
	temp = append(temp, num)
}

func dequeue() int {
	t := temp[0]
	temp[0] = temp[1]
	return t
}
