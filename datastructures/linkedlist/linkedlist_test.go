package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLinkedListEntry(t *testing.T) {
	list := NewLinkedListEntry(0, nil)
	next := NewLinkedListEntry(1, list)
	NewLinkedListEntry(2, next)

	assert.Equal(t, 0, list.Value)
	assert.Equal(t, 1, list.Next.Value)
	assert.Equal(t, 2, list.Next.Next.Value)
	assert.Nil(t, list.Next.Next.Next)
}

func TestNewLinkedListEntryFromArray(t *testing.T) {
	array := []int{0, 1, 2}
	list := NewLinkedListEntryFromArray(array)

	assert.Equal(t, 0, list.Value)
	assert.Equal(t, 1, list.Next.Value)
	assert.Equal(t, 2, list.Next.Next.Value)
	assert.Nil(t, list.Next.Next.Next)
}

func TestGetFromLinkedList(t *testing.T) {
	array := []int{101, 102, 103}
	list := NewLinkedListEntryFromArray(array)

	for i, v := range array {
		got, err := list.Get(i)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, v, got)
	}
}

func TestGetOverflowFromLinkedList(t *testing.T) {
	array := []int{101, 102, 103}
	list := NewLinkedListEntryFromArray(array)

	_, err := list.Get(100)
	if err == nil {
		t.Fatal("expected out of bounds error")
	}
}

func TestLengthOfLinkedList(t *testing.T) {
	array := []int{}
	list := NewLinkedListEntryFromArray(array)
	assert.Equal(t, 0, list.Length())

	array = []int{101}
	list = NewLinkedListEntryFromArray(array)
	assert.Equal(t, 1, list.Length())

	array = []int{101, 102, 103}
	list = NewLinkedListEntryFromArray(array)
	assert.Equal(t, 3, list.Length())
}

func TestAppendLinkedList(t *testing.T) {
	array := []int{101}
	list := NewLinkedListEntryFromArray(array)

	list.Append(102)
	assert.Equal(t, 2, list.Length())

	list.Append(103)
	assert.Equal(t, 3, list.Length())
}

func TestInsertLinkedList(t *testing.T) {
	array := []int{101, 103}
	list := NewLinkedListEntryFromArray(array)

	err := list.Insert(1, 102)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, list.Length())

	err = list.Insert(0, 100)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 4, list.Length())

	err = list.Insert(100, 102)
	if err == nil {
		t.Fatal("expected out of bounds error")
	}
}

func TestRemoveFromLinkedList(t *testing.T) {
	array := []int{101, 102}
	list := NewLinkedListEntryFromArray(array)

	err := list.Remove(0)
	if err != nil {
		t.Fatal(err)
	}

	err = list.Remove(0)
	if err == nil {
		t.Fatal("expected cannot delete only entry error")
	}

	array = []int{101, 102, 103}
	list = NewLinkedListEntryFromArray(array)

	err = list.Remove(2)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, list.Length())

	err = list.Remove(0)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, list.Length())
}
