package constant

import "testing"

const(
	MONDAY = 1 << iota
	TUESDAY
	WEDNESDAY
)

func TestConstants(t *testing.T) {
	t.Log(MONDAY, TUESDAY, WEDNESDAY)
}