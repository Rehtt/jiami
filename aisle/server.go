/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/4/29 14:46
 */

package model

type Server struct {
	Universal
}

func (s *Server) Send(str string) {
	s.index++
}

func (s *Server) TakeOver(data []byte) []byte {

}
