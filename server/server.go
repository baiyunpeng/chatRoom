package main

import (
	"net"
	"fmt"
	"github.com/baiyunpeng/chatRoom/const"
	"github.com/baiyunpeng/chatRoom/common"
)

func main() {
	listen_socket, err := net.Listen(constant.SERVER_PROTOCOL, common.ServerAddr())
	if (err != nil) {
		fmt.Println("启动监听错误")
	}
	//延时关闭监听
	defer listen_socket.Close();
	fmt.Println("服务监听启动中...")
	for {
		conn, err := listen_socket.Accept();
		if (err != nil) {
			fmt.Println("启动监听失败")
		}
		fmt.Println("连接服务器成功")
		for {
			message := common.Receive(conn)
			if (message == constant.SERVER_CLOSE) {
				common.SendMessage(conn, message)
				break;
			}
			fmt.Printf("接收到客户端数据：%v\n", message)
			fmt.Printf("请输入回复客户端的消息：")
			fmt.Scan(&message)
			common.SendMessage(conn, message)
		}
		conn.Close()
	}

}
