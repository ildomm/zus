package models

import (
	"fmt"
	"golang.org/x/crypto/sha3"
)

func (m *TokenInfo) GenerateHash() string {
	encoder := sha3.New256()
	encoder.Write([]byte(m.Token))
	result := encoder.Sum(nil)

	return fmt.Sprintf("%x", result)
}
