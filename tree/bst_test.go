package tree

import "testing"

func TestInsertMethod(t *testing.T) {
	var bst Bst
	bst.Insert(3, bst.Root)
	bst.Insert(1, bst.Root)
	bst.Insert(2, bst.Root)

	if bst.Size != 3 {
		t.Errorf("Insert failed; in correct insertion count: %v", bst.Size)
		return
	}

	if bst.Root.value != 3 {
		t.Errorf("Incorrect insertion of %v", bst.Root.value)
		return
	}

	if bst.Root.leftChild.value != 1 {
		t.Errorf("Incorrect insertion of %v", bst.Root.value)
		return
	}
	if bst.Root.leftChild.rightChild.value != 2 {
		t.Errorf("Incorrect insertion of %v", bst.Root.value)
		return
	}
}
