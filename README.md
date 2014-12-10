# ad

## Description
package `ad` implements helper function for the Microsoft Active Directory.

SID = Microsoft Security Identifier

## Usage
### String to SID
    sidString := "S-1-5-21-1117333035-483950394-1849977318-285965"
    sid, _ := ad.SIDFromString(sidString)
    encodedObjectSid := base64.StdEncoding.EncodeToString(sid)

### SID to String
    objectSid := "AQUAAAAAAAUVAAAAKyaZQjp/2BzmaURuDV0EAA=="
    decodedObjectSid, _ := base64.StdEncoding.DecodeString(objectSid)
    sid, _ := ad.SIDToString(decodedObjectSid)

## Author
* [Marc Sauter](mailto:marc.sauter@bluewin.ch)

