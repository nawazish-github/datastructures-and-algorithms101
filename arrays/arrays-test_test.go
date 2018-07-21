package arrays

import "testing"

func TestAddMethodToAddElementWhenSpaceIsAvailable(t *testing.T) {

	da := NewArray()
	da.Add(1)

	if da.Size != 1 {
		t.Errorf("Size of Dynamic Array is not 1")
	}
}

func TestAddMethodToAddElementWhenThereIsNotEnoughSpaceAvailable(t *testing.T) {

	da := NewArray()

	for i := 1; i <= 10; i++ {
		da.Add(i)
	}

	da.Add(11)
	da.Add(12)
	if da.Size != 12 {
		t.Errorf("Size of Dynamic Array is not 1")
	}
	if da.Data[10] != 11 {
		t.Errorf("11 was not found at index 10")
	}

	if da.Data[11] != 12 {
		t.Errorf("11 was not found at index 10")
	}

}
