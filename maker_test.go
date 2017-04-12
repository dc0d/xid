package xid

import (
	"testing"
	"time"
)

func TestMake(t *testing.T) {
	tm := time.Now()
	var c uint32 = 110
	if Make(tm, c) != Make(tm, c) {
		t.Fail()
	}
}
