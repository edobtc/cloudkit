package rpc

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

type Credentials struct {
	Username     string
	Password     string
	Salt         string
	HMACPassword string
	RPCAuth      string
}

func (c *Credentials) GenerateSalt(size int) (string, error) {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (c *Credentials) GeneratePassword() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

func (c *Credentials) PasswordToHmac(salt, password string) string {
	mac := hmac.New(sha256.New, []byte(salt))
	mac.Write([]byte(password))
	return hex.EncodeToString(mac.Sum(nil))
}

func (c *Credentials) Generate() (string, error) {
	if c.Password == "" {
		password, err := c.GeneratePassword()
		if err != nil {
			return "", err
		}
		c.Password = password
	}

	salt, err := c.GenerateSalt(16)
	if err != nil {
		return "", err
	}
	c.Salt = salt
	c.HMACPassword = c.PasswordToHmac(salt, c.Password)
	details := fmt.Sprintf("rpcauth=%s:%s$%s", c.Username, c.Salt, c.HMACPassword)
	c.RPCAuth = details
	return c.RPCAuth, nil
}
