/*
* @Author: qiuyu
* @Date:   2019-05-08 05:24:33
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-08 05:45:38
*/
package network_session_managenment

import (
	"net"
)

type Server struct {
	manager	*Manager
	listerner	net.Listener
	protocol	Protocol
	handler	    Handler
	sendChanSize	int
}

type Handler interface{
	HandleSession(*Session)
}

var _ Handler = HandlerFunc(nil)

type HandlerFunc func(*Session)

func (f HandlerFunc)HandleSession(seesion *Session) {
	f(seesion)
}

func NewServer(listerner net.Listener, protocol Protocol,sendChanSize int, handler Handler)*Server {
	return &Server{
		manager:   NewManager(),
		listerner: listerner,
		protocol: protocol,
		handler: handler,
		sendChanSize: sendChanSize,
	}
}

func (server *Server)Listener()net.Listener {
	return server.listerner
}

func (server *Server)Serve() error{
	for{
		connter, err := Accept(server.listerner)
        if err != nil {
        	return err
        }

        go func() {
        	codec, err := server.protocol.NewCodec(connter)
        	if err != nil {
        		connter.Close()
        	    return 
        	}
        	seesion := server.manager.NewSession(codec, server.sendChanSize)
        	server.handler.HandleSession(seesion)
        }()
	}
}

func (server *Server) GetSession(sessionID uint64)*Session{
	return server.manager.GetSession(sessionID)
}

func (server *Server) Stop() {
	server.listerner.Close()
	server.manager.Dispose()
}