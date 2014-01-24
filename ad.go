// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ad provides Microsoft AD specific stuff
package ad

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"strings"
)

// SIDFromString creates a SID (Microsoft Security Identifier)
// from its string representation.
// It returns the SID as []byte
func SIDFromString(sid string) []byte {
	buf := new(bytes.Buffer)
	parts := strings.Split(sid, "-")
	// 1 byte, revision must be 1
	buf.WriteByte(byte(1))
	// 1 byte, number of sub authorities
	buf.WriteByte(byte(len(parts[3:])))
	// 6 byte (big endian) IdentifierAuthority
	b := [8]byte{}
	x, _ := strconv.ParseUint(parts[2], 10, 48)
	binary.BigEndian.PutUint64(b[:], x)
	buf.Write(b[2:]) //
	// 4 byte (little endian) each SubAuthority
	for _, s := range parts[3:] {
		x, _ := strconv.ParseUint(s, 10, 32)
		binary.Write(buf, binary.LittleEndian, uint32(x))
	}
	return buf.Bytes()
}
