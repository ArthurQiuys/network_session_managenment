/*
* @Author: qiuyu
* @Date:   2019-05-08 06:15:29
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-08 06:37:37
*/
package network_session_managenment

import (
	"io"
	"net"
	"strings"
	"time"
)

type Protocol interface{
	NewCodec(rw io.ReadWriter)(Codec ,error)
}
type ProtocolFunc func(rw io.ReadWriter)(Codec,error)

func (pf ProtocolFunc)NewCodec(re io.ReadWriter)(Codec,error) {
	return pf(rw)
}

type Codec interface {
	Receive() (interface{}, error)
	Send(interface{}) error
	Close() error
}
type ClearSendChan interface{
		ClearSendChan(<-chan interface{})
}
func Listen(network, address string, protocol Protocol,sendChanSize int,handler Handler)(*Server ,error) {
	
}

func Dial(network, address string,protocol Protocol,sendChanSize int)(*Server ,error) {
	
}
func DialTimeout(network, address string, timeout time.Durtion,protocol Protocol,sendChanSize int)(*Server ,error) {
	
}
func Accept(listerner net.Listener)(net.Conn, error) {
	
}