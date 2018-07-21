package arrays

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
