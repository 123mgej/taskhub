package requestid

import (
	"crypto/rand"
	"encoding/hex"
)

func New() string {
	b:=make([]byte,16) // 16 bytes
	_,err :=rand.Read(b)
	if err !=nil {
		return "requestid-error"
	}
	return hex.EncodeToString(b)
}

