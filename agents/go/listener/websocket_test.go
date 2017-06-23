package listener

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/maxlaverse/reverse-shell/common"
	"github.com/maxlaverse/reverse-shell/message"
)

func init() {
	log.SetFlags(0)
}

func TestStructComparison(t *testing.T) {
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandleFunc("/agent/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := common.Upgrader.Upgrade(w, r, nil)
		_, m, err := conn.ReadMessage()
		if err != nil {
			common.Logger.Debugf("ReadMessage error: %s", err)
		}
		b := message.FromBinary(m)
		switch v := b.(type) {
		case *message.ProcessOutput:
			os.Stdout.Write(v.Data)
		case *message.ProcessCreated:
			log.Printf("New session is named: %s\n", v.Id)
		case *message.ProcessTerminated:
			log.Printf("Session closed: %s\n", v.Id)
		default:
			common.Logger.Debugf("Received an unknown message type: %v", v)
		}

	})
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
