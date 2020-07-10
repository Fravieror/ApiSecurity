package domain

type Cipher interface {
	Encrypt(string) []byte
	Decrypt(string) []byte
}
