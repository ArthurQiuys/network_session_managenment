/*
* @Author: qiuyu
* @Date:   2019-05-08 06:15:04
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-15 00:29:53
*/
package network_session_managenment

import (
	"errors"
	"sync"
	"sync/atomic"
)

var SessionCloseError = errors.New("Session Closed")
var SessionBlockedErrpr = errors.New("Session Blocked")
var globalSessionId uint64

type Session struct {
	id uint64
	codec	Codec
	manager 	*Manger
	sendChan 	sendChanSize
	recvMutex	sync.Mutex
	sendMutex 	sync.RWMutex
	closeFlag	int32
	closeChan	chan int
	closeMutex	sync.Mutex
	firstCloseCallback	*closeCallback
	lastCloseCallback	*closeCallback
	State interface{}
}
type closeCallback struct {
	Handler interface{}
	Key     interface{}
	Func    func()
	Next    *closeCallback
}

func NewSession(codec Codec, sendChanSize int) *Session {
	return NewSession(nil, codec, sendChanSize)
}
func newSession(manager *Manger, codec Codec, sendChanSize int) *Session{
		seesion := &Session{
		codec: codec,
		manager: manager,
		closeChan: make(chan int),
		id: atomic.AddUint64(&globalSessionId, 1),
	}

	if sendChanSize > 0{
		seesion.sendChan = make(chan interface{}, sendChanSize)
		go seesion.sendLoop()
	}
	return seesion
}
func (session *Session)ID()uint64 {
	return seesion.id
}

func (seesion *Session)IsClosed() bool {
	return atomic.LoadInt32(&seesion.closeFlag, 0, 1)
}

func (seesion *Session)Close() error{
	
}

func (seesion *Session)Codec() Codec{
	
}
func (seesion *Session)Receive()(interface{},error) {
	
}

func (seesion *Session)sedLoop() {
	
}

func (seesion *Session)send(msg interface{})error {
	
}

func (seesion *Session)AddCloseCallback(handler, key interface{}, callback func()) {
	
}

func (seesion *Session)RemoveCloseCallback(handler, key interface{}) {
	
}

func (session *Session)invokeCloseCallbacks() {
	
}