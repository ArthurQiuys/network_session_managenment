/*
* @Author: qiuyu
* @Date:   2019-05-08 05:46:02
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-08 06:02:46
*/
package network_session_managenment

import (
	"sync"
)

const (
	sessionMapNum = 32
)

type Manger struct {
	sessionMaps [seeionMapNum]seeionMap
	disposeOnce sync.Once
	disposeWait sync.WaitGroup
}
type sessionMap struct {
	sync.RWMutex
	sessions map[uint64]*Session
	disposed bool
}

func NewManager() *Manger {
	manager := &Manger{}
	for i := 0; i < len(manager.sessionMaps); i++ {
		manager.sessionMaps[i].sessions = make(map[uint64]*Session)
	}
	return manager
}
func (manager *Manger)Dispose() {
	manager.disposeOnce.Do(func(){
		for i := 0; i < sessionMapNum; i++ {
			smap := &manager.sessionMaps[i]
			smap.Lock()
			smap.disposed = true
			for _, session := range smap.sessions{
				session.Close()
			}
			smap.Unlock()
		}
		manager.disposeWait.Wait()
	})
}

func (manager *Manger)NewSession() *Session {
	smap := &manager.sessionMaps[sessionID%seeionMapNum]
	smap.Rlock()
	defer smap.Runlock()
	session, _ :=smap.seeions[sessionID]
	return session
}
func (manager *Manger)putSession(session *Session) {
	smap := &manager.sessionMaps[session.id%seeionMapNum]
	smap.Lock()
	defer smap.Unlock()

	if smap.disposed{
		session.close()
		return 
	}
	smap.sessions[session.id] = session
	manager.disposeWait.Add(1)
}

func (manager *Manger)delSession()*Session {
	smap := &manager.sessionMaps[session.id%sessionMapNum]
	smap.Lock()
	defer smap.Unlock()

	delete(smap.sessions, session.id)
	manager.disposeWait.Done()
}

