/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/4/29 14:46
 */

package aisle

import (
	"jiami/aisle/util/ecc"
)

type Client struct {
	Universal
}

func (c *Client) Send(data []byte) ([]byte, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	return nil, nil
}
func (s *Client) TakeOver(data []byte) ([]byte, bool, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	return nil, false, nil
}

func (s *Client) First() (out []byte, err error) {
	if s.index != 0 {
		s.eccKey = nil
		return nil, nil
	}
	e := ecc.GenerateKey()
	s.eccKey = &e
	return e.GeneratePublicKey().ConvertToBytes(), nil
}

func (s *Client) FirstTakeOver(data []byte) (err error) {
	if s.index != 0 {
		s.eccKey = nil
		return nil
	}
	dData, err := s.eccKey.Decrypt(data)
	if err != nil {
		return err
	}
	s.key = dData
	s.index++
	return nil
}

func NewClient() *Client {
	return &Client{}
}
