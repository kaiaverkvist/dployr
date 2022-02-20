package utils_test

import (
	"github.com/kaiaverkvist/dployr/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testSliceWithDuplicates = []string {
	"a",
	"b",
	"c",
	"c",
	"d",
}

var expectedSlice = []string {
	"a",
	"b",
	"c",
	"d",
}

func TestUniqueNonEmptyElementsOf(t *testing.T) {
	m := testSliceWithDuplicates
	result := utils.UniqueNonEmptyElementsOf(m)

	assert.Equal(t, expectedSlice, result)
}