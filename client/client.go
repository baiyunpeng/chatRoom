package main

import (
	"net"
	"fmt"
	"github.com/baiyunpeng/chatRoom/const"
	"github.com/baiyunpeng/chatRoom/common"
)

func main() {
	serverAddr, _ := getInitData();
	fmt.Printf("正在连接服务器，服务器地址：%v\n", serverAddr)
	conn, err := net.Dial(constant.SERVER_PROTOCOL, "10.168.62.33:8088")
	if (err != nil) {
		fmt.Println("客户端连接服务器失败")
		return;
	}
	defer connectionClose(conn)
/*	var clientConnStatus = connServer(conn, serverAddr, nick);
	fmt.Println("clientConnStatus:" + clientConnStatus)
	if (constant.CLIENT_CONN_SUCCESS == clientConnStatus) {
		fmt.Println("连接服务器成功")
	} else {
		fmt.Println("连接服务器失败")
	}*/
	go func() {
		for {
			message := common.Receive(conn);
			fmt.Printf("接收到服务器消息：%v\n", message)
		}
	}()
	for {
		var message = "";
		fmt.Println("请输入消息内容：")
		fmt.Scan(&message)
		fmt.Printf("发送消息内容：%v\n", message)
		common.SendMessage(conn, message)

	}
	fmt.Println("服务器连接断开")
	conn.Close()
}

func connServer(conn net.Conn, serverAddr, nick string) string {
	common.SendMessage(conn, constant.CLIENT_LOGIN_PROTOCOL+nick)
	return common.Receive(conn);
}

func connectionClose(conn net.Conn) {
	conn.Close();
	fmt.Println("conn连接关闭")
}

/**
初始化用户信息
 */
func getInitData() (string, string) {
	var addr, nick string;
	fmt.Println("欢迎使用聊天室...")
	fmt.Println("------------------------------")
	/*fmt.Println("请输入连接服务器地址：")
	//服务器地址
	fmt.Scan(&addr)
	fmt.Printf("您输入的服务器地址是：%v\n", addr)*/
	fmt.Println("请输入您的昵称：")
	fmt.Scan(&nick)
	return addr, nick;
}
