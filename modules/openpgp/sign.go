package openpgp

import (
	"bytes"
	"github.com/keybase/go-crypto/openpgp"
)

// Sign ...
func (o *OpenPGP) Sign(message, publicKey, privateKey, passphrase string) (string, error) {

	entity, err := o.ReadSignKey(publicKey, privateKey, passphrase)
	if err != nil {
		return "", err
	}

	writer := new(bytes.Buffer)
	reader := bytes.NewReader([]byte(message))
	err = openpgp.ArmoredDetachSign(writer, entity, reader, nil)
	if err != nil {
		return "", err
	}

	return writer.String(), nil
}
