package unit_test

import (
	"testing"

	"github.com/miyabiii1210/ulala/go/library/encrypt"
)

const (
	DEFAULT_API_SECRET string = "ABCDEFG123456789"
)

var (
	plaintext     string = "gedgdre3r51sada45bhsacb21xasd45vgghdhf"
	encryptedText string = "AAAAAAAAAAAAAAAAAAAAAJd03iDllrlVskP5+Gy81fAu40pGkBf8hmLbkB5PPROh/nAS3RJl"
)

func TestEncrypt(t *testing.T) {
	encryptedText, err := encrypt.Encrypt(plaintext, DEFAULT_API_SECRET)
	if err != nil {
		t.Errorf("Encrypt failed: %v", err)
		return
	}

	t.Logf("Encrypt text: %v", encryptedText)
}

func TestDecrypt(t *testing.T) {
	decryptedText, err := encrypt.Decrypt(encryptedText, DEFAULT_API_SECRET)
	if err != nil {
		t.Errorf("decrypt error: %v", err)
		return
	}

	t.Logf("Decrypted text: %v", decryptedText)
}
