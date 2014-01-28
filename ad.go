// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ad provides Microsoft AD specific stuff
package ad

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

// SIDFromString creates a SID (Microsoft Security Identifier)
// from its string representation.
// It returns the SID as []byte
func SIDFromString(sid string) ([]byte, error) {
	// TODO: check for valid sid ... what is valid?
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
	return buf.Bytes(), nil
}

// SIDToString creates the string
// representation of a SID (Microsoft Security Identifier)
// It returns the SID as string
func SIDToString(sid []byte) (string, error) {
	// TODO: check for valid sid ... what is valid?
	buf := new(bytes.Buffer)
	// sid[0] - 1 byte, revision must be 1
	buf.WriteString("S-1")
	// sid[1] - 1 byte, number of sub authorities
	// not used in string representation
	num := uint8(sid[1])
	// sid[2:7] - 6 byte (big endian) IdentifierAuthority
	ia := sid[2 : 2+6]
	buf.WriteString(fmt.Sprintf("-%d", binary.BigEndian.Uint64(append([]byte{0, 0}, ia...))))
	// sid[8:] - 4 byte (little endian) each SubAuthority
	var i uint8
	for i = 0; i < num; i++ {
		buf.WriteString(fmt.Sprintf("-%d", binary.LittleEndian.Uint32(sid[(i*4)+8:8+(i*4)+4])))
	}
	return buf.String(), nil
}
