package crypto

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"strconv"
)

const (
	saltSize = 32
)

// HashString returns the string sha1 on the hashed
// input string which we can use for labels and tags
// for resources included in an experiment
func HashString(input string) string {
	// Compute 20- byte sha1
	var x = sha1.Sum([]byte(input))

	// Get the first 15 characters of the hexdigest.
	var y = fmt.Sprintf("%x", x[0:8])
	return y[0 : len(y)-1]
}

// Hash returns a uint64 from sha1 hash of any input string
// which is the basic for taking experiment parameter inpt
// and choosing a variant based on index
func Hash(input string) uint64 {
	y := HashString(input)
	var z uint64
	z, _ = strconv.ParseUint(y, 16, 64)

	return z
}

// GenerateSalt will generate a salt for use when one is not provided
func GenerateSalt() string {
	saltLetters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	salt := make([]rune, saltSize)
	size := len(saltLetters)

	for i := range salt {
		salt[i] = saltLetters[rand.Intn(size)]
	}
	return string(salt)
}
