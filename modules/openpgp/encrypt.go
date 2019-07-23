package openpgp

import (
	"bytes"
	"github.com/keybase/go-crypto/openpgp"
	"github.com/keybase/go-crypto/openpgp/armor"
	"io/ioutil"
)

// Encrypt ...
func (o *OpenPGP) Encrypt(message, publicKey string) (string, error) {

	entityList, err := o.ReadPublicKey(publicKey)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	w, err := openpgp.Encrypt(buf, entityList, nil, nil, nil)
	if err != nil {
		return "", err
	}
	_, err = w.Write([]byte(message))
	if err != nil {
		return "", err
	}
	err = w.Close()
	if err != nil {
		return "", err
	}

	output, err := ioutil.ReadAll(buf)
	if err != nil {
		return "", err
	}

	pubKeyBuf := bytes.NewBuffer(nil)
	pubKeyWriter, err := armor.Encode(pubKeyBuf, messageHeader, headers)
	if err != nil {
		return "", err
	}
	_, err = pubKeyWriter.Write(output)
	if err != nil {
		return "", err
	}
	err = pubKeyWriter.Close()
	if err != nil {
		return "", err
	}

	outputString := pubKeyBuf.String()

	return outputString, nil
}
