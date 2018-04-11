package main

import (
	"sync"

	"github.com/gorilla/websocket"
)

type AgentTable struct {
	sync.Mutex
	agents map[string]*websocket.Conn
}

func NewAgentTable() AgentTable {
	return AgentTable{
		agents: make(map[string]*websocket.Conn),
	}
}

func (s *AgentTable) AddAgent(conn *websocket.Conn) {
	s.Lock()
	defer s.Unlock()
	s.agents[conn.RemoteAddr().String()] = conn

}

func (s *AgentTable) RemoveAgent(conn *websocket.Conn) {
	s.Lock()
	defer s.Unlock()
	delete(s.agents, conn.RemoteAddr().String())
}

func (b *AgentTable) FindAgent(address string) *websocket.Conn {
	b.Lock()
	defer b.Unlock()
	return b.agents[address]
}
