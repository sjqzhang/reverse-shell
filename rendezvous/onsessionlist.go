package main

import (
	"encoding/json"
	"net/http"

	"github.com/maxlaverse/reverse-shell/rendezvous/api"
)

type onSessionList struct{}

func (h onSessionList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := []api.SessionListResponseAgent{}
	for k, c := range sessionTable.sessionTable {
		m := []string{}
		for _, c1 := range c.masterConn {
			m = append(m, c1.RemoteAddr().String())
		}
		s := api.SessionListResponseAgent{
			Name:    k,
			Agent:   c.agentConn.RemoteAddr().String(),
			Masters: m,
			State:   c.State.String(),
		}
		p = append(p, s)
	}
	json.NewEncoder(w).Encode(p)
}
