package handler

import (
	"sync"
)

type ProcessTable struct {
	sync.Mutex
	processes map[string]*Process
}

func newProcessTable() *ProcessTable {
	processes := make(map[string]*Process)

	return &ProcessTable{
		processes,
	}
}

func (s *ProcessTable) New(id string, command string) *Process {
	s.Lock()
	defer s.Unlock()
	newProcess := NewProcess(id, command)
	s.processes[newProcess.Id] = newProcess
	return newProcess
}

func (s *ProcessTable) List() []string {
	s.Lock()
	defer processes.Unlock()
	keys := make([]string, 0, len(s.processes))
	for k := range s.processes {
		keys = append(keys, k)
	}
	return keys
}

func (s *ProcessTable) Get(id string) *Process {
	s.Lock()
	defer s.Unlock()
	return s.processes[id]
}

func (s *ProcessTable) Remove(process *Process) {
	s.Lock()
	defer s.Unlock()
	delete(s.processes, process.Id)
}
