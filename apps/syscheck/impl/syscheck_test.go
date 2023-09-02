package impl_test

import "testing"

func TestStopFirewall(t *testing.T) {
	err := svc.StopFirewall(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStopSelinux(t *testing.T) {
	err := svc.StopSelinux(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
