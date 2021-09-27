package test

import "testing"

func TestGcDetail001(t *testing.T) {
	// GOMAXPROCS=8 GODEBUG=schedtrace=500,scheddetail=1 godoc -http=:6060
	// GOMAXPROCS=8 GODEBUG=schedtrace=1000 godoc -http=:6060
	//  wrk -t20 -c10000 -d10s http://localhost:6060
}
