package main

import (
	"net"
	"fmt"
	"github.com/baiyunpeng/chatRoom/const"
	"github.com/baiyunpeng/chatRoom/common"
	"strings"
)

/**
客户端结构体
 */
type Client struct {
	//地址
	addr string
	//姓名
	name string
	//通讯管道
	channel chan string
	//连接信息
	conn net.Conn
}

//储存在线用户的map
var onlieclient = make(map[string]Client)
//管道间的消息
var channelMessage = make(chan string)

//遍历管道消息
func pushonline() {
	for {
		pushmsg := <-channelMessage
		for _, client := range onlieclient {
			client.channel <- pushmsg
		}
	}
}

func listenerConn(socket net.Listener) {
	conn, err := socket.Accept();
	if (err != nil) {
		fmt.Println("服务器监听失败")
		return
	}
	fmt.Println("连接服务器成功")
	initClient(conn);
}

func initClient(conn net.Conn) {
	defer conn.Close()
	//获取客户端地址
	address := conn.RemoteAddr().String();
	//打印连接成功
	fmt.Printf("[%s]:%s\n", address, "连接成功")
	clientLoginRotocol := common.Receive(conn);
	//有登录协议
	if (strings.HasPrefix(clientLoginRotocol, constant.CLIENT_LOGIN_PROTOCOL)) {
		common.SendMessage(conn,constant.CLIENT_CONN_SUCCESS)
		clientName := clientLoginRotocol[len(constant.CLIENT_LOGIN_PROTOCOL):]
		//创建Client对象
		client := Client{address, clientName, make(chan string), conn}
		//将client放入map集合
		onlieclient[clientName+":"+address] = client
		fmt.Println("更新后:", onlieclient)
		message := fmt.Sprintf("用户%v登录了，地址:%v", clientName, address)
		//广播上线信息
		broadcast(message)
	}
}

func handelData() {
	for _, client := range onlieclient {
		fmt.Println("开始遍历客户端")
		message := common.Receive(client.conn);
		broadcast(message)
	}
}

/**
广播消息
 */
func broadcast(message string) {
	for _, client := range onlieclient {
		common.SendMessage(client.conn, message);
	}
}

func main() {
	fmt.Printf("服务端启动，地址：%v,端口号%v\n", constant.SERVER_ADDR, constant.SERVER_PORT)
	listen_socket, err := net.Listen(constant.SERVER_PROTOCOL, common.ServerAddr())
	if (err != nil) {
		fmt.Println("服务器启动监听失败")
	}
	//延时关闭监听
	defer listen_socket.Close();
	fmt.Println("服务监听启动成功，等待消息")
	go handelData()

	for {
		listenerConn(listen_socket);
	}

}
