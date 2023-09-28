package encoding_test

import (
	"testing"
	"uuid-converter/encoding"
)

func TestDecode(t *testing.T) {
	input := "0523baa5e0c54e9aa4d93306cf8325be"
	expectedOutput := "0523baa5-e0c5-4e9a-a4d93-306cf8325be"

	got, err := encoding.Decode(input)
	if err != nil {
		t.Fatal("err should be nil")
	}

	if got != expectedOutput {
		t.Fatalf("unexpected output %q, expected: %q", got, expectedOutput)
	}
}
