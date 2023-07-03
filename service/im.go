package service

import (
	"encoding/json"
	"fmt"
	"ginchat/models"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer func(ws *websocket.Conn) {
	// 	err := ws.Close()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }(ws)

	iuserId, _ := strconv.Atoi(c.Query("userId"))
	userId := uint(iuserId)
	// targetId, _ := strconv.Atoi(c.Query("targetId"))

	node := Node{
		Conn:      ws,
		DataQueue: make(chan []byte),
		GroupSet:  set.New(set.ThreadSafe),
	}
	rwLocker.Lock()
	clientMap[userId] = &node
	rwLocker.Unlock()

	go sendProc(&node)
	go recvProc(&node)

	sendP2PMsg(userId, userId, []byte("hello"))

	// for {
	// 	utils.Publish(c, utils.PublishKey, "hello")
	// 	msg, err := utils.Subscribe(c, utils.PublishKey)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	err = ws.WriteMessage(websocket.TextMessage, []byte(msg))
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// }
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSet  set.Interface
}

var clientMap map[uint]*Node = make(map[uint]*Node)

var rwLocker sync.RWMutex

func sendProc(node *Node) {
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(node.Conn)

	for {
		msg := <-node.DataQueue
		fmt.Println("sendProc <-", msg)
		err := node.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func recvProc(node *Node) {
	for {
		_, p, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		broadcast(p)
	}
}

var udpsendChan = make(chan []byte)

func broadcast(p []byte) {
	udpsendChan <- p
}

func init() {
	go udpSendProc()
	go udpRecvProc()
}

func udpSendProc() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8081,
	})
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		select {
		case data := <-udpsendChan:
			conn.Write(data)
		}
	}
}

func udpRecvProc() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 8081,
	})
	if err != nil {
		panic("udp listen err: " + err.Error())
	}
	defer conn.Close()
	for {
		data := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err)
			break
		}
		dispatch(data[:n])
	}
}

func dispatch(p []byte) {
	msg := models.Message{}
	err := json.Unmarshal(p, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("dispatch", msg)
	switch msg.Type {
	case models.P2P:
		sendP2PMsg(msg.FromId, msg.TargetId, []byte(msg.Content))
		// case models.GROUP: sendGroupMsg(msg)
		// case models.BROADCAST: sendBroadCastMsg(msg)
	}
}

func sendP2PMsg(fromId uint, targetId uint, data []byte) {
	rwLocker.RLock()
	node, ok := clientMap[targetId]
	rwLocker.RUnlock()
	if !ok {
		return
	}
	node.DataQueue <- data
}
