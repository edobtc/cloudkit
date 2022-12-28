package ssh

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	mrand "math/rand"
	"os"

	"golang.org/x/crypto/ssh"
)

const (
	DefaultBitSize        = 4096
	DefaultPrivateKeyPath = "id_rsa"
	DefaultPublicKeyPath  = "id_rsa.pub"
)

// SSHKey is a struct that holds the options when
// creating a new ssh key
type Options struct {
	BitSize    int    `json:"bit_size"`
	PriKeyPath string `json:"pri_key_path"`
	PubKeyPath string `json:"pub_key_path"`
}

// SetDefaults sets the default values for the options
func (o *Options) SetDefaults() {
	if o.BitSize == 0 {
		o.BitSize = DefaultBitSize
	}

	if o.PriKeyPath == "" {
		o.PriKeyPath = DefaultPrivateKeyPath
	}

	if o.PubKeyPath == "" {
		o.PubKeyPath = DefaultPublicKeyPath
	}
}

// DefaultOptions returns the default options
func DefaultOptions() Options {
	return Options{
		BitSize:    DefaultBitSize,
		PriKeyPath: DefaultPrivateKeyPath,
		PubKeyPath: DefaultPublicKeyPath,
	}
}

type SshKey struct {
	options Options
	PrivKey *rsa.PrivateKey
	Priv    []byte
	Pub     []byte
}

func NewSSHKey(opt Options) *SshKey {
	return &SshKey{
		options: opt,
	}
}

func (s *SshKey) Generate() error {
	privateKey, err := s.generatePrivateKey()
	if err != nil {
		return err
	}

	s.PrivKey = privateKey

	s.generatePublicKey()
	if err != nil {
		log.Fatal(err.Error())
	}

	s.Priv = s.encodePrivateKeyToPEM()

	return nil
}

// generatePrivateKey creates a RSA Private Key of specified byte size
func (s *SshKey) generatePrivateKey() (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, s.options.BitSize)
	if err != nil {
		return nil, err
	}

	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// encodePrivateKeyToPEM encodes Private Key from RSA to PEM format
func (s *SshKey) encodePrivateKeyToPEM() []byte {
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PrivateKey(s.PrivKey),
	}

	return pem.EncodeToMemory(&privBlock)
}

// generatePublicKey take a rsa.PublicKey and return
// bytes suitable for writing to .pub file
// returns in the format "ssh-rsa ..."
func (s *SshKey) generatePublicKey() ([]byte, error) {
	pub, err := ssh.NewPublicKey(s.PrivKey.Public())
	if err != nil {
		return []byte{}, err
	}

	s.Pub = ssh.MarshalAuthorizedKey(pub)

	return s.Pub, nil
}

func (s *SshKey) SaveToFile() error {
	err := write(s.Priv, s.options.PriKeyPath)
	if err != nil {
		return err
	}

	err = write(s.Pub, s.options.PubKeyPath)
	if err != nil {
		return err
	}

	return nil
}

func write(keyBytes []byte, saveFileTo string) error {
	err := os.WriteFile(saveFileTo, keyBytes, 0600)
	if err != nil {
		return err
	}

	return nil
}

func GenerateED25519Key() ([]byte, []byte) {
	pubKey, privKey, _ := ed25519.GenerateKey(rand.Reader)
	publicKey, _ := ssh.NewPublicKey(pubKey)

	pemKey := &pem.Block{
		Type:  "OPENSSH PRIVATE KEY",
		Bytes: MarshalED25519PrivateKey(privKey),
	}
	privateKey := pem.EncodeToMemory(pemKey)
	authorizedKey := ssh.MarshalAuthorizedKey(publicKey)

	return privateKey, authorizedKey
}

func MarshalED25519PrivateKey(key ed25519.PrivateKey) []byte {
	// Add our key header (followed by a null byte)
	magic := append([]byte("openssh-key-v1"), 0)

	var w struct {
		CipherName   string
		KdfName      string
		KdfOpts      string
		NumKeys      uint32
		PubKey       []byte
		PrivKeyBlock []byte
	}

	// Fill out the private key fields
	pk1 := struct {
		Check1  uint32
		Check2  uint32
		Keytype string
		Pub     []byte
		Priv    []byte
		Comment string
		Pad     []byte `ssh:"rest"`
	}{}

	// Set our check ints
	ci := mrand.Uint32()
	pk1.Check1 = ci
	pk1.Check2 = ci

	// Set our key type
	pk1.Keytype = ssh.KeyAlgoED25519

	// Add the pubkey to the optionally-encrypted block
	pk, ok := key.Public().(ed25519.PublicKey)
	if !ok {
		//fmt.Fprintln(os.Stderr, "ed25519.PublicKey type assertion failed on an ed25519 public key. This should never ever happen.")
		return nil
	}
	pubKey := []byte(pk)
	pk1.Pub = pubKey

	// Add our private key
	pk1.Priv = []byte(key)

	// Might be useful to put something in here at some point
	pk1.Comment = ""

	// Add some padding to match the encryption block size within PrivKeyBlock (without Pad field)
	// 8 doesn't match the documentation, but that's what ssh-keygen uses for unencrypted keys. *shrug*
	bs := 8
	blockLen := len(ssh.Marshal(pk1))
	padLen := (bs - (blockLen % bs)) % bs
	pk1.Pad = make([]byte, padLen)

	// Padding is a sequence of bytes like: 1, 2, 3...
	for i := 0; i < padLen; i++ {
		pk1.Pad[i] = byte(i + 1)
	}

	// Generate the pubkey prefix "\0\0\0\nssh-ed25519\0\0\0 "
	prefix := []byte{0x0, 0x0, 0x0, 0x0b}
	prefix = append(prefix, []byte(ssh.KeyAlgoED25519)...)
	prefix = append(prefix, []byte{0x0, 0x0, 0x0, 0x20}...)

	// Only going to support unencrypted keys for now
	w.CipherName = "none"
	w.KdfName = "none"
	w.KdfOpts = ""
	w.NumKeys = 1
	w.PubKey = append(prefix, pubKey...)
	w.PrivKeyBlock = ssh.Marshal(pk1)

	magic = append(magic, ssh.Marshal(w)...)

	return magic
}
