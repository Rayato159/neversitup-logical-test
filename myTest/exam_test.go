package myTest

import (
	"fmt"
	"testing"

	"github.com/Rayato159/neversitup-logical-test/exam"
	"github.com/stretchr/testify/assert"
)

type testValidatePincode struct {
	input  string
	expect bool
}

func TestValidatePincode(t *testing.T) {
	tests := []testValidatePincode{
		{
			input:  "nervergonnagiveyouup",
			expect: false,
		},
		{
			input:  "-17283",
			expect: false,
		},
		{
			input:  "172839",
			expect: true,
		},
		{
			input:  "111822",
			expect: false,
		},
		{
			input:  "112762",
			expect: true,
		},
		{
			input:  "123743",
			expect: false,
		},
		{
			input:  "321895",
			expect: false,
		},
		{
			input:  "124578",
			expect: true,
		},
		{
			input:  "112233",
			expect: false,
		},
		{
			input:  "882211",
			expect: false,
		},
		{
			input:  "887712",
			expect: true,
		},
	}

	for _, test := range tests {
		fmt.Println("test ->", test.input)
		assert.Equal(t, test.expect, exam.ValidatePincode(test.input))
	}
}
