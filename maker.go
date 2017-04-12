package xid

import (
	"encoding/binary"
	"sync/atomic"
	"time"
)

// Make generates an ID from given time & counter
func Make(t time.Time, counter ...uint32) ID {
	var id ID
	// Timestamp, 4 bytes, big endian
	binary.BigEndian.PutUint32(id[:], uint32(t.Unix()))
	// Machine, first 3 bytes of md5(hostname)
	id[4] = machineID[0]
	id[5] = machineID[1]
	id[6] = machineID[2]
	// Pid, 2 bytes, specs don't specify endianness, but we use big endian.
	id[7] = byte(pid >> 8)
	id[8] = byte(pid)
	// 3 bytes, big endian
	var i uint32
	if len(counter) > 0 {
		i = counter[0]
	} else {
		i = atomic.AddUint32(&objectIDCounter, 1)
	}
	id[9] = byte(i >> 16)
	id[10] = byte(i >> 8)
	id[11] = byte(i)
	return id
}
