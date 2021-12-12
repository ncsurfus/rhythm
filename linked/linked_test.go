package linked

import "testing"

func TestInsertAfter(t *testing.T) {
	values := []int{0, 1, 2, 3, 4, 6}
	root := FromSlice(values)
	root.Next.Next.InsertAfter(7)
	expectedValues := []int{0, 1, 2, 7, 3, 4, 6}
	checkValue := root
	for _, value := range expectedValues {
		if value != checkValue.Value {
			t.Fatal("Wrong LinkedElement value!")
		}
		checkValue = checkValue.Next
	}
}

func TestLast(t *testing.T) {
	values := []int{0, 1, 2, 3, 4, 6}
	root := FromSlice(values)
	if root.Last().Value != 6 {
		t.Fatal("Wrong last value!")
	}
}

func TestFindLinkedListElementByValue(t *testing.T) {
	values := []int{0, 1, 2, 3, 4, 6}
	root := FromSlice(values)
	result, err := FindElementByValue(root, 3)
	if err != nil {
		t.Fatal("LinkedList didn't have nil error!")
	}
	if result.Value != 3 {
		t.Fatal("Wrong value!")
	}
	if result.Next.Value != 4 {
		t.Fatal("Next value was not 4!")
	}
}
