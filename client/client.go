package main

import (
	"net"
	"fmt"
	"../const"
	"../common"
	"../modes"
	"encoding/json"
	"strings"
)

//管道对象
var channelChatChannel = make(chan modes.Chat)

func main() {
	serverAddr, nick := getInitData();
	fmt.Printf("正在连接服务器，服务器地址：%v\n", serverAddr)
	conn, err := net.Dial(constant.SERVER_PROTOCOL, serverAddr)
	if !common.CheckError(err, "客户端连接服务器失败") {
		return
	}

	defer connectionClose(conn)
	connServer(conn, serverAddr, nick);
	common.MonitorConn(conn, monitorConn)

	go handelData();
	for {
		var message = "";
		fmt.Println("请输入消息内容：")
		fmt.Scan(&message)
		fmt.Printf("发送消息内容：%v\n", message)
		sendMessage(nick, message, conn)

	}
	fmt.Println("服务器连接断开")
	conn.Close()
}

func connServer(conn net.Conn, serverAddr, nick string) {
	chat := modes.Chat{nick, "", constant.CALL_TYPE_COMMOND, constant.CALL_COMMON_CONNECTION, ""}
	message, err := json.Marshal(chat);
	if common.CheckError(err, "转换对象失败") {
		conn.Write(message);
	}
}

func connectionClose(conn net.Conn) {
	conn.Close();
	fmt.Println("conn连接关闭")
}

//监听连接
func monitorConn(chat modes.Chat, conn net.Conn) {
	channelChatChannel <- chat;
}

func handelData() {
	for {
		chat := <-channelChatChannel
		fmt.Println(chat.Sender, ":", chat.Message)
	}
}

func sendMessage(nick string, message string, conn net.Conn) {
	var receiver string = "";
	var callType string = constant.CALL_TYPE_BROADCAST;
	if strings.HasPrefix(message, "@") {
		index := strings.Index(message, ":")
		receiver = message[1:index]
		message = message[index+1:]
		callType = constant.CALL_TYPE_P2P
	}
	chat := modes.Chat{nick, receiver, callType, message, ""}
	messageByte, err := json.Marshal(chat);
	if common.CheckError(err, "数据转换失败") {
		conn.Write(messageByte);
	}
}

/**
初始化用户信息
 */
func getInitData() (string, string) {
	var addr, nick string;
	fmt.Println("欢迎使用聊天室...")
	fmt.Println("------------------------------")
	fmt.Println("请输入连接服务器地址：")
	//服务器地址
	fmt.Scan(&addr)
	fmt.Printf("您输入的服务器地址是：%v\n", addr)
	fmt.Println("请输入您的昵称：")
	fmt.Scan(&nick)
	return addr, nick;
}
