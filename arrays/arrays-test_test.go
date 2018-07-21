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

	fillUpDynamicArray(da)

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

func TestInsertMethodToInsertElementWhenSpaceIsAvailable(t *testing.T) {
	da := NewArray()
	da.Add(1)
	da.Add(2)
	da.Add(4)

	da.Insert(3, 2)

	if da.Data[2] != 3 {
		t.Errorf("Insertion failed")
	}
}

func TestInsertWhenSpaceIsNotAvailable(t *testing.T) {
	da := NewArray()
	fillUpDynamicArray(da)
	da.Insert(2.5, 2)

	if da.Data[2] != 2.5 {
		t.Errorf("Could not insert 2.5 between 2 and 3")
	}
}

func fillUpDynamicArray(da *DynamicArrays) {
	for i := 1; i <= 10; i++ {
		da.Add(i)
	}
}

func TestInsertWhenElementIsToBeInsertedAtTheEndWithSpaceAvailable(t *testing.T) {
	da := NewArray()
	da.Add(1)
	da.Add(2)

	da.Insert(3, 2)

	if da.Data[2] != 3 {
		t.Errorf("Could not insert 3 at the end")
	}
}

func TestInsertToPanicWhenAttemptToInsertBeyondCapacityOfArray(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("Insertion beyond the capacity did not fail graciously")
		}
	}()

	da := NewArray()
	da.Insert(14, 43)
}
