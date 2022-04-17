package ctrl

import (
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"log"
	"net/http"
	"strconv"
	"sync"
)

//本核心在于形成userid和Node的映射关系
type Node struct {
	Conn *websocket.Conn
	//并行转串行,
	DataQueue chan []byte
	GroupSets set.Interface
}

//映射关系表
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

//读写锁
var rwlocker sync.RWMutex

//
// ws://127.0.0.1/chat?id=1&token=xxxx
func Chat(writer http.ResponseWriter,
	request *http.Request) {

	//todo 获取用户id和token
	query := request.URL.Query()
	id := query.Get("id")
	token := query.Get("token")

	//todo 检验接入是否合法
	userId, _ := strconv.ParseInt(id, 10, 64)
	isvalida := checkToken(userId, token)

	//todo 获得conn
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return isvalida
	}}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	//todo userid和node形成绑定关系
	rwlocker.Lock()
	clientMap[userId] = node
	rwlocker.Unlock()
	//todo 完成发送逻辑,con
	sendproc(node)
	//todo 完成接收逻辑
	recvproc(node)
	//发送测试内容
	sendMsg(7, []byte("hello world!"))
}

//发送协程
func sendproc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}

}

//接收协程
func recvproc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}
		//todo 对data进一步处理
		fmt.Printf("recv<=%s", data)
	}
}

//todo 发送消息
func sendMsg(userId int64, msg []byte) {
	rwlocker.RLock()
	node, ok := clientMap[userId]
	rwlocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}

//检测是否有效
func checkToken(userId int64, token string) bool {
	user := userService.Find(userId)
	return user.Token == token
}
