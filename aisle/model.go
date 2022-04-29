/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/4/29 14:46
 */

package aisle

import (
	"jiami/aisle/util/ecc"
	"sync"
)

type Universal struct {
	index  uint64
	key    []byte
	lock   sync.Mutex
	eccKey *ecc.Key[ecc.PrivateKey]
}
