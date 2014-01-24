package ad

import (
	"encoding/base64"
	"testing"
)

// some real world examples:
// "S-1-5-21-1117333035-483950394-1849977318-285965" -> "AQUAAAAAAAUVAAAAKyaZQjp/2BzmaURuDV0EAA=="
// "S-1-5-21-1117333035-483950394-1849977318-216321" -> "AQUAAAAAAAUVAAAAKyaZQjp/2BzmaURuAU0DAA=="
// "S-1-5-21-1117333035-483950394-1849977318-35238"  -> "AQUAAAAAAAUVAAAAKyaZQjp/2BzmaURupokAAA=="

func TestSIDFromString(t *testing.T) {
	sidString := "S-1-5-21-1117333035-483950394-1849977318-285965"
	objectSid := "AQUAAAAAAAUVAAAAKyaZQjp/2BzmaURuDV0EAA=="
	encodedSid := base64.StdEncoding.EncodeToString(SIDFromString(sidString))
	if encodedSid != objectSid {
		t.Errorf("wrong sid\n")
	}
}
