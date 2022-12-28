package ssh

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionsDefaults(t *testing.T) {
	o := Options{}
	o.SetDefaults()

	if o.BitSize != DefaultBitSize {
		t.Errorf("expected bit size to be %d, got %d", DefaultBitSize, o.BitSize)
	}

	if o.PriKeyPath != DefaultPrivateKeyPath {
		t.Errorf("expected private key path to be %s, got %s", DefaultPrivateKeyPath, o.PriKeyPath)
	}

	if o.PubKeyPath != DefaultPublicKeyPath {
		t.Errorf("expected public key path to be %s, got %s", DefaultPublicKeyPath, o.PubKeyPath)
	}
}

func TestGenerate(t *testing.T) {
	s := NewSSHKey(DefaultOptions())
	err := s.Generate()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, s.Priv, "should have a private key")
	assert.NotNil(t, s.Pub, "should have a public key")
}

func TestWriteFiles(t *testing.T) {
	o := DefaultOptions()
	s := NewSSHKey(o)
	err := s.Generate()
	if err != nil {
		t.Error(err)
	}

	err = s.SaveToFile()
	if err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(o.PriKeyPath); err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(o.PubKeyPath); err != nil {
		t.Error(err)
	}

	os.Remove(o.PriKeyPath)
	os.Remove(o.PubKeyPath)
}

func TestValues(t *testing.T) {
	o := DefaultOptions()
	s := NewSSHKey(o)
	err := s.Generate()
	if err != nil {
		t.Error(err)
	}

	err = s.SaveToFile()
	if err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(o.PriKeyPath); err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(o.PubKeyPath); err != nil {
		t.Error(err)
	}

	fPriv, err := os.ReadFile(o.PriKeyPath)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, s.Priv, fPriv, "should write correct data to file")

	fPub, err := os.ReadFile(o.PubKeyPath)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, s.Pub, fPub, "should write correct data to file")
	os.Remove(o.PriKeyPath)
	os.Remove(o.PubKeyPath)
}
