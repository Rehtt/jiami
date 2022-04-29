/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/4/29 17:19
 */

package ecc

import (
	"fmt"
	"log"
	"testing"
)

func TestT(t *testing.T) {
	key := GenerateKey()
	pubKey := key.GeneratePublicKey()
	fmt.Println(pubKey.ConvertToBytes())
	fmt.Println(key.ConvertToBytes())

	eData, err := pubKey.Encrypt([]byte("test"))
	if err != nil {
		log.Fatalln(eData)
	}
	fmt.Println("加密：", eData)

	data, err := key.Decrypt(eData)
	if err != nil {
		log.Fatalln(data)
	}
	fmt.Println("解密：", string(data))
}
