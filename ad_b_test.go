package ad

// go test -bench=".*"

import (
	"encoding/base64"
	"testing"
)

func BenchmarkSIDFromString(b *testing.B) {
	sidString := "S-1-5-21-1117333035-483950394-1849977318-285965"
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SIDFromString(sidString)
	}
}

func BenchmarkSIDToString(b *testing.B) {
	objectSid := "AQUAAAAAAAUVAAAAKyaZQjp/2BzmaURuDV0EAA=="
	decodedObjectSid, _ := base64.StdEncoding.DecodeString(objectSid)
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SIDToString(decodedObjectSid)
	}
}
