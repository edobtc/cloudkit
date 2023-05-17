package rpc

import "testing"

func TestGenerateSalt(t *testing.T) {
	c := Credentials{}
	salt, err := c.GenerateSalt(16)
	if err != nil {
		t.Error(err)
	}
	if len(salt) != 32 {
		t.Errorf("Expected salt to be 32 bytes long, got %d", len(salt))
	}
}

func TestGeneratePassword(t *testing.T) {
	c := Credentials{}
	password, err := c.GeneratePassword()
	if err != nil {
		t.Error(err)
	}
	if len(password) != 44 {
		t.Errorf("Expected password to be 44 bytes long, got %d", len(password))
	}
}

func TestPasswordToHmac(t *testing.T) {
	c := Credentials{}
	salt, err := c.GenerateSalt(16)
	if err != nil {
		t.Error(err)
	}
	password, err := c.GeneratePassword()
	if err != nil {
		t.Error(err)
	}
	hmac := c.PasswordToHmac(salt, password)
	if len(hmac) != 64 {
		t.Errorf("Expected hmac to be 64 bytes long, got %d", len(hmac))
	}
}
