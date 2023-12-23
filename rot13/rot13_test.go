package kata

import "testing"

func TestRot13(t *testing.T) {
	t.Run("short example", func(t *testing.T) {
		assertExpectedResult(t, Rot13("EBG13 rknzcyr."), "ROT13 example.")
	})
	t.Run("long example with line breaks", func(t *testing.T) {
		assertExpectedResult(t, Rot13("How can you tell an extrovert from an\nintrovert at NSA? Va gur ryringbef,\ngur rkgebireg ybbxf ng gur BGURE thl'f fubrf."), "Ubj pna lbh gryy na rkgebireg sebz na\nvagebireg ng AFN? In the elevators,\nthe extrovert looks at the OTHER guy's shoes.")
	})
	t.Run("number example", func(t *testing.T) {
		assertExpectedResult(t, Rot13("123"), "123")
	})
	t.Run("special symbol example", func(t *testing.T) {
		assertExpectedResult(t, Rot13("@[`{"), "@[`{")
	})
}

func assertExpectedResult(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected '%s', but got '%s'", want, got)
	}
}
