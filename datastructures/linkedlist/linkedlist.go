package linkedlist

import "fmt"

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func NewLinkedListEntry(value int, previous *LinkedList) *LinkedList {
	entry := LinkedList{
		Value: value,
		Next:  nil,
	}

	if previous != nil {
		previous.Next = &entry
	}

	return &entry
}

func NewLinkedListEntryFromArray(values []int) *LinkedList {
	var start, cur *LinkedList

	for i, value := range values {
		next := &LinkedList{
			Value: value,
			Next:  nil,
		}

		if i == 0 {
			start = next
		} else {
			cur.Next = next
		}

		cur = next
	}

	return start
}

func (l *LinkedList) Get(index int) (int, error) {
	if index == 0 {
		return l.Value, nil
	}

	cur := l.Next

	for i := 1; i <= index; i++ {
		if i == index {
			return cur.Value, nil
		}

		cur = cur.Next
		if cur == nil {
			break
		}
	}

	return -1, fmt.Errorf("index of %v is out of bounds", index)
}

func (l *LinkedList) Append(value int) {
	cur := l

	for {
		if cur.Next == nil {
			cur.Next = &LinkedList{
				Value: value,
				Next:  nil,
			}
			break
		}

		cur = cur.Next
	}
}

func (l *LinkedList) Insert(index int, value int) error {
	if index == 0 {
		l.Next = &LinkedList{
			Value: l.Value,
			Next:  l.Next,
		}

		l.Value = value
		return nil
	}

	cur := l.Next

	for i := 1; i <= index; i++ {
		if i == index {
			next := &LinkedList{
				Value: value,
				Next:  cur.Next,
			}
			cur.Next = next
			break
		}

		cur = cur.Next
		if cur == nil {
			return fmt.Errorf("index of %v is out of bounds", index)
		}
	}

	return nil
}

func (l *LinkedList) Length() int {
	cur := l
	length := 0

	if cur == nil {
		return length
	}

	for {
		cur = cur.Next
		length++

		if cur == nil {
			break
		}
	}

	return length
}

func (l *LinkedList) Remove(index int) error {
	if index == 0 {
		if l.Next == nil {
			return fmt.Errorf("cannot delete only entry")
		}

		l.Value = l.Next.Value
		l.Next = l.Next.Next
		return nil
	}

	cur := l

	for i := 1; i <= index; i++ {
		if cur.Next == nil {
			return fmt.Errorf("index of %v is out of bounds", index)
		}

		if i == index {
			cur.Next = cur.Next.Next
			break
		}

		cur = cur.Next
	}

	return nil
}
