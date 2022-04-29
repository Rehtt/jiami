/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/4/29 14:46
 */

package aisle

import (
	"jiami/aisle/util"
	"jiami/aisle/util/ecc"
)

type Server struct {
	Universal
}

func (s *Server) Send(data []byte) ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.index++

	return nil, nil
}

func (s *Server) TakeOver(data []byte) ([]byte, bool, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	// 第一次连接
	data, err := s.First(data)
	if err != nil {
		return nil, false, err
	}
	if data != nil {
		return data, true, nil
	}

	return nil, false, nil
}

func (s *Server) First(data []byte) (out []byte, err error) {
	if s.index != 0 {
		return nil, nil
	}
	s.index++
	s.key = util.GenerateRandomBytes(5)
	pub, err := ecc.ConvertToPubKey(data)
	if err != nil {
		return
	}
	eData, err := pub.Encrypt(s.key)
	return eData, err
}
