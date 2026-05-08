package permissionbolt

import (
	"testing"

	"github.com/xyproto/pinterface/v2"
)

func TestInterface(t *testing.T) {
	perm, err := New()
	if err != nil {
		t.Fatal(err)
	}

	// Check that the value qualifies for the interface
	var _ pinterface.IPermissions = perm
}
