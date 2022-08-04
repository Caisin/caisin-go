package rsakey

import "testing"

func TestGenKey(t *testing.T) {
	GenerateRSAKey("pub.pem", "pri.pem", 2048)
}
