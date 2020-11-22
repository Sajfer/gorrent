package bitfield_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/sajfer/gorrent/bitfield"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	rc := m.Run()

	// rc 0 means we've passed,
	// and CoverMode will be non empty if run with -cover
	if rc == 0 && testing.CoverMode() != "" {
		c := testing.Coverage()
		if c < 0.8 {
			fmt.Println("Tests passed but coverage failed at", c)
			rc = -1
		}
	}
	os.Exit(rc)
}

func TestHasPiece(t *testing.T) {
	bf := bitfield.Bitfield{0b01010100, 0b01010100}
	outputs := []bool{false, true, false, true, false, true, false, false, false, true, false, true, false, true, false, false, false, false, false, false}
	for i := 0; i < len(outputs); i++ {
		assert.Equal(t, outputs[i], bf.HasPiece(i))
	}
}

func TestSetPiece(t *testing.T) {
	tests := []struct {
		input  bitfield.Bitfield
		index  int
		output bitfield.Bitfield
	}{
		{
			input:  bitfield.Bitfield{0b01010100, 0b01010100},
			index:  4,
			output: bitfield.Bitfield{0b01011100, 0b01010100},
		},
		{
			input:  bitfield.Bitfield{0b01010100, 0b01010100},
			index:  9,
			output: bitfield.Bitfield{0b01010100, 0b01010100},
		},
		{
			input:  bitfield.Bitfield{0b01010100, 0b01010100},
			index:  15,
			output: bitfield.Bitfield{0b01010100, 0b01010101},
		},
		{
			input:  bitfield.Bitfield{0b01010100, 0b01010100},
			index:  19, //                            v (noop)
			output: bitfield.Bitfield{0b01010100, 0b01010100},
		},
	}
	for _, test := range tests {
		bf := test.input
		bf.SetPiece(test.index)
		assert.Equal(t, test.output, bf)
	}
}
