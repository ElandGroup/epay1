package helper

import (
	"encoding/hex"

	uuid "github.com/satori/go.uuid"
)

func Uuid32() string {
	u1 := uuid.NewV4()
	return hex.EncodeToString(u1.Bytes())
}
