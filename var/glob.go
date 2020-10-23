package glob

import (
	"errors"
	"strings"
)

//addr hash
var AddrPrefix string = "gech"

//url namespace
var Symbol string = AddrPrefix

//sys hash
var Hash string = AddrPrefix

func IsAddr(bytes []byte) bool {
	return len(bytes) == (40 + len(AddrPrefix))
}

func IsHash(byte2 []byte) bool {
	return strings.HasPrefix(strings.ToUpper(string(byte2)), strings.ToUpper(Hash))
}

func GetHashStr(str string) string {
	if strings.HasPrefix(strings.ToUpper(str), strings.ToUpper(AddrPrefix)) {
		return str[len(AddrPrefix):]
	}
	if strings.HasPrefix(strings.ToUpper(str), strings.ToUpper(Hash)) {
		return str[len(Hash):]
	}
	return str
}

func GetHash(byte2 []byte, wantPrefix bool) ([]byte, error) {
	if IsAddr(byte2) {
		return byte2[len(AddrPrefix):], nil
	}
	if IsHash(byte2) {
		return byte2[len(Hash):], nil
	}
	if wantPrefix {
		return nil, errors.New("not hash")
	} else {
		return byte2, nil
	}
}
