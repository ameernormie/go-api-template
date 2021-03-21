package encryptionservice

import (
	"testing"

	"github.com/ameernormie/go-api-template/internal/testutil"
)

func TestEncryption(t *testing.T) {
	testutil.InitDb()

	plainText := "text to be encrypted"

	encrypted, err := Encrypt(plainText)
	if err != nil {
		t.Errorf("Error while encrypting: %v", err)
		return
	}

	decrypted, err := Decrypt(encrypted)
	if err != nil {
		t.Errorf("Error while decrypting: %v", err)
		return
	}

	if decrypted != plainText {
		t.Errorf("Decrypted text is not equal to original text")
		return
	}
}
