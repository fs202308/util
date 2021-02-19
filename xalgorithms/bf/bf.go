package bf

import (
	"fmt"
	"sync"

	"golang.org/x/crypto/blowfish"
)

var wg sync.WaitGroup

var chanBuffer chan []byte = make(chan []byte, 5000)

var cryptKey = []byte{0x00, 0x00, 0x00, 0x00}

func makeBuffer(size int) []byte {
	return make([]byte, size)
}

func copy_safe(dst []byte, src []byte) {
	srcLen := len(src)
	dstLen := len(dst)
	minLen := srcLen

	if srcLen > dstLen {
		minLen = dstLen
	}

	copy(dst, src[:minLen])
}

func WithKey(key []byte) {
	cryptKey = key
}

func EncryptData(data []byte) []byte {
	c, err := blowfish.NewCipher(cryptKey)
	if err != nil {
		fmt.Printf("NewCipher(%d bytes) = %s\n", len(cryptKey), err)
		return nil
	}

	encryptLen := len(data)
	encryptTimes := encryptLen / 8

	tmpEnBuffer := makeBuffer(encryptLen)
	for i := 0; i < encryptTimes; i += 1 {
		ct := make([]byte, 8)
		c.Encrypt(ct, data[(i*8):])
		copy_safe(tmpEnBuffer[(i*8):], ct)
	}
	copy_safe(tmpEnBuffer[encryptTimes*8:], data[encryptTimes*8:])

	return tmpEnBuffer
}

func DecryptData(data []byte) []byte {
	c, err := blowfish.NewCipher(cryptKey)
	if err != nil {
		fmt.Printf("NewCipher(%d bytes) = %s\n", len(cryptKey), err)
		return nil
	}

	readytoDecrypt := data
	decryptLen := len(readytoDecrypt)
	encryptTimes := decryptLen / 8

	tmpDebuffer := makeBuffer(decryptLen)
	for i := 0; i < encryptTimes; i += 1 {
		pt := make([]byte, 8)
		c.Decrypt(pt, data[(i*8):])
		copy_safe(tmpDebuffer[(i*8):], pt)
	}
	copy_safe(tmpDebuffer[encryptTimes*8:], data[encryptTimes*8:])

	return tmpDebuffer
}
