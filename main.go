/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/4/29 14:58
 */

package main

import (
	"fmt"
	"jiami/aisle"
	"log"
)

func main() {
	c := aisle.NewClient()
	s := aisle.NewServer()
	data, err := c.First()
	if err != nil {
		log.Fatalln(err)
	}

	sData, _, _ := s.TakeOver(data)

	fmt.Println(c.FirstTakeOver(sData))

}
