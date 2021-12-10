package x509key

//go:generate stringer -type=KeyType
type KeyType int

const (
	RSA KeyType = iota + 1
	ECDSA
	Ed25519
)
