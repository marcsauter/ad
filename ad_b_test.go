package ad

// go test -bench=".*"

import (
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
