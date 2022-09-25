package controllers

import (
	"encoding/json"
	"fmt"
	"mustafa_m/models"
	"mustafa_m/views"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var wsChan = make(chan WsPayload)

var clients = make(map[*websocket.Conn]string)

type MessageBoard struct {
	MessageBoard   *views.View
	MessageService *models.MessageService
}

// Defines the response that we send to the frontend
type WsResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

type WsPayload struct {
	Action   string          `json:"action"`
	Message  string          `json:"message"`
	UserName string          `json:"username"`
	Conn     *websocket.Conn `json:"-"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func NewMessageController(ms *models.MessageService) *MessageBoard {
	return &MessageBoard{
		MessageBoard:   views.NewView("bootstrap", "messageBoard/messageBoard.gohtml"),
		MessageService: ms,
	}
}

func (mb *MessageBoard) GetMessageBoard(w http.ResponseWriter, r *http.Request) {
	mb.MessageBoard.Render(w, nil)
}

func (mb *MessageBoard) GetRecentMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := mb.MessageService.GetRecentMessages()
	if err != nil {
		//
	}
	jsonEncoding, err := json.Marshal(messages)
	fmt.Fprintln(w, string(jsonEncoding))
}

func (mb *MessageBoard) Websocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		panic(err)
	}

	resp := WsResponse{
		Message: "This is the response from the server",
	}

	err = ws.WriteJSON(resp)

	if err != nil {
		panic(err)
	}

	clients[ws] = ""

	// creating a function that runs behind the scenes and holds the websocket

	go keepWebsocketConnection(ws)
}

func keepWebsocketConnection(ws *websocket.Conn) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	var payload WsPayload
	for {
		err := ws.ReadJSON(&payload)
		if err != nil {
		} else {
			payload.Conn = ws
			wsChan <- payload
		}
	}
}

func ListenToChannel(ms *models.MessageService) {
	response := WsResponse{}

	for {
		e := <-wsChan
		switch e.Action {
		case "username":
			clients[e.Conn] = e.UserName
			response.Action = "list_users"
			response.Message = strconv.Itoa(len(clients))
			broadCastToAll(response)

		case "broadcast":
			response.Action = "list_message"
			response.Message = fmt.Sprintf("<strong>%s</strong>: %s", e.UserName, e.Message)
			broadCastToAll(response)
			go ms.WriteMessage(&models.Message{Username: e.UserName, Message: e.Message})
		}

	}
}

func broadCastToAll(response WsResponse) {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			fmt.Println("This is where the error is")
			fmt.Println(err)
			fmt.Println("This is where the error is x2")
			client.Close()
			// delete from the map since their connection is no longer valid
			delete(clients, client)
		}
	}
}

func AddMessageBoardRoutes(r *mux.Router, mbC *MessageBoard) {
	r.HandleFunc("/msg", mbC.GetMessageBoard)
	r.HandleFunc("/api/v1/msg/ws", mbC.Websocket)
	r.HandleFunc("/api/v1/msg/recent", mbC.GetRecentMessages)
}
