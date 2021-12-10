package x509key

//go:generate stringer -type=KeyType
type KeyType int

const (
	Uknown KeyType = iota
	RSA
	ECDSA
	Ed25519
)
