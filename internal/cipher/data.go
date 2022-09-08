package cipher

import (
	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/keyset"
	"github.com/google/tink/go/tink"
	"main/internal/cipher/master_key"
	"os"
)

var aeadCipher tink.AEAD

func createOrRead(path string) *keyset.Handle {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return create(path)
	}
	return read(path)
}

func create(path string) *keyset.Handle {
	kh, err := keyset.NewHandle(aead.AES256GCMKeyTemplate())
	if err != nil {
		panic(err)
	}
	file, _ := os.Create(path)
	w := keyset.NewBinaryWriter(file)
	if err := kh.Write(w, master_key.New()); err != nil {
		panic(err)
	}
	return kh
}

func read(path string) *keyset.Handle {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	bin := keyset.NewBinaryReader(file)
	kh, err := keyset.Read(bin, master_key.New())
	if err != nil {
		panic(err)
	}
	return kh
}
