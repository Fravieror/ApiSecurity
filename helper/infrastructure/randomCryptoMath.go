package infrastructure

import (
	crand "crypto/rand"
	rand "math/rand"

	"encoding/binary"
	"log"
)

// CryptoSource struct.
type CryptoSource struct{}

// GenRandom obtiene random number combinando librerias math and crypto.
func (s *CryptoSource) GenRandom() int {
	rnd := rand.New(s)
	return rnd.Intn(1000) // a truly random number 0 to 999
}

// Seed no se que hace.
func (s *CryptoSource) Seed(seed int64) {}

// Int63 tampoco se.
func (s *CryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

// Uint64 menos se.
func (s *CryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
