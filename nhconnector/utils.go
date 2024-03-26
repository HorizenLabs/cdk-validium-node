package nhconnector

import "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"

// Bytes800 represents an 800 byte array
type Bytes800 [800]byte

// NewBytes800 creates a new Bytes800 type
func NewBytes800(b [800]byte) Bytes800 {
	return Bytes800(b)
}

// Convert a 800 bytes hex string into Bytes800
func Bytes800FromHex(str string) (Bytes800, error) {
	b, err := codec.HexDecodeString(str)
	if err != nil {
		return Bytes800{}, err
	}
	return Bytes800(b), err
}
