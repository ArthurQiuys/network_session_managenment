/*
* @Author: qiuyu
* @Date:   2019-05-08 06:03:23
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-08 06:12:43
*/
package network_session_managenment

import (
	"sync"
)

type KEY interface{}

type Channel struct {
	mutex sync.RWMutex
	sessions map[KEY]*Session
	State interface{}
}

func NewChannel()*Channel {
	return &Channel{
		sessions: make(map[KEY]Session),
	}
}

func (channel *Channel)Len() int{
	channel.mutex.RLock()
	defer channel.mutex.RUnlock()
	return len(channel.sessions)
}

func (channel *Channel)Fetch(callback func(*Session)) {
	channel.mutex.RLock()
	defer channel.mutex.RUnlock()
	for _ , seesion := range channel.sessions{
		callback(seesion)
	}
}

func (channel *Channel)Get( key KEY)* Session {
	
}

func (channel *Channel)Put(key KEY, seesion *Session) {

}

func (channel *Channel)remove(key KEY, seesion *Session) {
	
}

func (channel *Channel)Remove(key KEY) bool {
	
}

func (channel *Channel)FetchAndRemove(callback func(*Session)) {
	
}

func (channel *Channel)Close() {
	
}