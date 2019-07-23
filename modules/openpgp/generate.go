package openpgp

import (
	"bytes"
	"github.com/keybase/go-crypto/openpgp"
	"github.com/keybase/go-crypto/openpgp/armor"
)

// KeyPair ...
type KeyPair struct {
	PublicKey  string
	PrivateKey string
}

// Options ...
type Options struct {
	KeyOptions *KeyOptions
	Name       string
	Comment    string
	Email      string
	Passphrase string
}

// Generate ...
func (o *OpenPGP) Generate(options *Options) (*KeyPair, error) {

	var keyPair *KeyPair
	config := generatePacketConfig(options.KeyOptions)
	entity, err := openpgp.NewEntity(options.Name, options.Comment, options.Email, config)
	if err != nil {
		return keyPair, err
	}

	for _, id := range entity.Identities {
		err := id.SelfSignature.SignUserId(id.UserId.Id, entity.PrimaryKey, entity.PrivateKey, nil)
		if err != nil {
			return keyPair, err
		}
	}

	if options.Passphrase != "" {
		err = entity.PrivateKey.Encrypt([]byte(options.Passphrase), nil)
		if err != nil {
			return keyPair, err
		}
	}

	keyPair = &KeyPair{}
	privateKeyBuf := bytes.NewBuffer(nil)
	writerPrivate, err := armor.Encode(privateKeyBuf, openpgp.PrivateKeyType, headers)
	if err != nil {
		return keyPair, err
	}
	defer writerPrivate.Close()

	err = entity.SerializePrivate(writerPrivate, nil)
	if err != nil {
		return keyPair, err
	}
	writerPrivate.Close()
	keyPair.PrivateKey = privateKeyBuf.String()

	publicKeyBuf := bytes.NewBuffer(nil)
	writerPublic, err := armor.Encode(publicKeyBuf, openpgp.PublicKeyType, headers)
	if err != nil {
		return keyPair, err
	}
	defer writerPublic.Close()

	err = entity.Serialize(writerPublic)
	if err != nil {
		return keyPair, err
	}
	writerPublic.Close()
	keyPair.PublicKey = publicKeyBuf.String()

	return keyPair, nil
}
