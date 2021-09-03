package architecture

import (
	"runtime"
	"testing"
)

func TestDependent(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Not work in arch amd64")
	}

	t.Fail()
}
