/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/4/29 14:54
 */

package ecc

import (
	"errors"
	ecies "github.com/ecies/go/v2"
	"log"
)

type Key[T PublicKey | PrivateKey] struct {
	key T
}
type PublicKey struct {
	*ecies.PublicKey
}
type PrivateKey struct {
	*ecies.PrivateKey
}

func GenerateKey() Key[PrivateKey] {
	k, err := ecies.GenerateKey()
	if err != nil {
		log.Fatalln(err)
	}
	return Key[PrivateKey]{PrivateKey{k}}
}

func ConvertToPubKey(pub []byte) (Key[PublicKey], error) {
	p, err := ecies.NewPublicKeyFromBytes(pub)
	return Key[PublicKey]{PublicKey{p}}, err
}
func (key Key[T]) ConvertToBytes() []byte {
	switch key := any(key.key).(type) {
	case PrivateKey:
		return key.Bytes()
	case PublicKey:
		return key.Bytes(true)
	}
	return nil
}

func (key Key[T]) GeneratePublicKey() Key[PublicKey] {
	switch key := any(key.key).(type) {
	case PrivateKey:
		return Key[PublicKey]{PublicKey{key.PublicKey}}
	case PublicKey:
		return Key[PublicKey]{key}
	}
	return Key[PublicKey]{}
}

func (key Key[T]) Encrypt(data []byte) (eData []byte, err error) {
	switch key := any(key.key).(type) {
	case PrivateKey:
		eData, err = ecies.Encrypt(key.PublicKey, data)
	case PublicKey:
		eData, err = ecies.Encrypt(key.PublicKey, data)
	}
	return
}

func (key Key[T]) Decrypt(data []byte) (dData []byte, err error) {
	err = errors.New("没有私钥")

	switch key := any(key.key).(type) {
	case PrivateKey:
		dData, err = ecies.Decrypt(key.PrivateKey, data)
	}
	return
}
