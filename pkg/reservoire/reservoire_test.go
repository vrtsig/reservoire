package reservoire_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vrtsig/reservoire/pkg/reservoire"
)

func TestStringReservoire(t *testing.T) {
	testLines := make([]string, 1000)
	for i := 0; i < len(testLines); i++ {
		testLines[i] = "line " + strconv.Itoa(i+1)
	}

	seed := int64(0)
	r1, _ := reservoire.NewStringReservoire(1, seed)
	r10, _ := reservoire.NewStringReservoire(10, seed)
	r100, _ := reservoire.NewStringReservoire(100, seed)
	r100b, _ := reservoire.NewStringReservoire(100, seed)
	r100c, _ := reservoire.NewStringReservoire(100, seed+1)
	r1000, _ := reservoire.NewStringReservoire(1000, seed)
	r2000, _ := reservoire.NewStringReservoire(2000, seed)

	for _, s := range testLines {
		r1.Add(s)
		r10.Add(s)
		r100.Add(s)
		r100b.Add(s)
		r100c.Add(s)
		r1000.Add(s)
		r2000.Add(s)
	}

	// assert the correct length
	assert.Equal(t, r1.Len(), 1)
	assert.Equal(t, r100.Len(), 100)
	assert.Equal(t, r1000.Len(), 1000)
	assert.Equal(t, r2000.Len(), 1000)

	// big enough reservoires should contain the whole input set
	assert.Equal(t, r1000.GetAll(), testLines)
	assert.Equal(t, r2000.GetAll(), testLines)

	// seeds should work as intended
	assert.Equal(t, r100.GetAll(), r100b.GetAll())
	assert.NotEqual(t, r100.GetAll(), r100c.GetAll())
}
