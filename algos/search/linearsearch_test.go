package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearSearch(t *testing.T) {
	array := []int{101, 102, 103}

	index := linearsearch(array, 101)
	assert.Equal(t, 0, index)

	index = linearsearch(array, 103)
	assert.Equal(t, 2, index)

	index = linearsearch(array, 999)
	assert.Equal(t, -1, index)
}
