package selinux_test

import (
	"testing"

	"kubevirt.io/client-go/testutils"
)

func TestSelinux(t *testing.T) {
	testutils.KubeVirtTestSuiteSetup(t, "Selinux Suite")
}
