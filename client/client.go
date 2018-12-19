package main

import (
	"net"
	"fmt"
	"github.com/baiyunpeng/chatRoom/const"
	"github.com/baiyunpeng/chatRoom/common"
)

func main() {
	serverAddr := getServerAddr();
	fmt.Println("正在连接服务器，服务器地址：%v\n", serverAddr)
	conn, err := net.Dial(constant.SERVER_PROTOCOL, serverAddr)
	if (err != nil) {
		fmt.Println("客户端连接服务器失败")
		return;
	}
	defer conn.Close()
	for {
		var message = "";
		fmt.Println("请输入消息内容：")
		fmt.Scan(&message)
		fmt.Printf("发送消息内容：%v\n", message)
		common.SendMessage(conn, message)
		message = common.Receive(conn);
		fmt.Printf("接收到服务器消息：%v\n", message)
		if (message == constant.SERVER_CLOSE) {
			common.SendMessage(conn, constant.SERVER_CLOSE)
			break;
		}
	}
	fmt.Println("服务器连接断开")
	conn.Close()
}

func getServerAddr() string {
	fmt.Println("欢迎使用聊天室...")
	fmt.Println("------------------------------")
	fmt.Println("请输入连接服务器地址：")
	//服务器地址
	var serverAddr = "";
	fmt.Scan(&serverAddr)
	fmt.Printf("您输入的服务器地址是：%v\n", serverAddr)
	return serverAddr;
}
