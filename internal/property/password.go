package property

import (
	"database/sql/driver"
	"encoding/base64"
	"main/internal/cipher"
)

type Password string

// Value store value to database
func (p Password) Value() (driver.Value, error) {
	bytes, err := cipher.Encrypt(string(p))
	return base64.StdEncoding.EncodeToString(bytes), err
}

// Scan get value from database
func (p *Password) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	var ciphertext string
	switch src.(type) {
	case string:
		ciphertext = src.(string)
	case []byte:
		ciphertext = string(src.([]byte))
	}
	val, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return err
	}
	decrypted, err := cipher.Decrypt(string(val))
	if err != nil {
		*p = Password(ciphertext)
		return nil
	}
	*p = Password(decrypted)
	return nil
}
