package models

import (
	"fmt"
	"github.com/ildomm/linx_challenge/config"
	"crypto/md5"
)


func (m *TokenInfo) GenerateHash() string {
	h := md5.New()
	h.Write([]byte(m.Token))
	m.Hash = config.App.Runtime.Url + fmt.Sprintf("%x", h.Sum(nil))
	return m.Hash
}
