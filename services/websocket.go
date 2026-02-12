package services

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type ServerService struct {
	server     *http.Server
	stopServer chan bool
	upgrader   websocket.Upgrader
}

// NewServerService creates a new ServerService
func NewServerService() *ServerService {
	return &ServerService{
		stopServer: make(chan bool),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

// StartServer starts the server
func (s *ServerService) StartServer(port int) {

	s.stopServer = make(chan bool)

	http.HandleFunc("/", s.handleConnections)
	s.server = &http.Server{Addr: fmt.Sprintf(":%d", port)}

	go func() {
		fmt.Printf("Server started on port %d\n", port)
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("ListenAndServe: %v\n", err)
		}
	}()

	<-s.stopServer
	if err := s.server.Close(); err != nil {
		fmt.Printf("Server Close: %v\n", err)
	}
	fmt.Println("Server stopped")
}

// StopServer stops the server
func (s *ServerService) StopServer() {
	s.stopServer <- true
}

// handleConnections handles incoming websocket connections
func (s *ServerService) handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	for {
		// Read message from browser
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}

		msgString := string(msg)
		WailsApp.Event.EmitEvent(&application.CustomEvent{
			Name: "socket:message",
			Data: map[string]interface{}{
				"data": msgString,
			},
		})

	}
}
