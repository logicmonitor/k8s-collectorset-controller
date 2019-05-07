package err

import (
	"testing"
)

func TestRecoverError(t *testing.T) {
	run(t)
	t.Logf("Test success")

}

func run(t *testing.T) {
	defer RecoverError("Test msg")
	t.Logf("Test run start...")
	panic("Test panic")
	t.Logf("Test run end...")
}
