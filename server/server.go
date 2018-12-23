package main

import (
	"net"
	"fmt"
	"../const"
	"../common"
	"../modes"
)

//储存在线用户的map
var onlieclient = make(map[string]modes.Client)

//遍历管道消息
func handelData(chat modes.Chat, conn net.Conn) {
	//连接
	if chat.CallType == constant.CALL_TYPE_COMMOND {
		handelCommonData(chat, conn)
	} else if chat.CallType == constant.CALL_TYPE_GROUP {
		fmt.Println("CALL_TYPE_GROUP", chat)
	} else if chat.CallType == constant.CALL_TYPE_BROADCAST {
		broadcast(chat);
	} else if chat.CallType == constant.CALL_TYPE_P2P {
		p2p(chat)
	}
}

func handelCommonData(chat modes.Chat, conn net.Conn) {
	if chat.Message == constant.CALL_COMMON_CONNECTION {
		registerClietn(chat, conn);
	} else if chat.Message == constant.CALL_COMMON_CLOSE {
		closeClietn(chat)
	}
}
func registerClietn(chat modes.Chat, conn net.Conn) {
	client := modes.Client{chat.Sender, chat.Sender, make(chan modes.Chat)}
	onlieclient[chat.Sender] = client
	common.MonitorChat(conn, client)
	onlineChat := modes.Chat{chat.Sender, "", constant.CALL_TYPE_BROADCAST, chat.Sender + "上线了", ""}
	broadcast(onlineChat);
}

func closeClietn(chat modes.Chat) {
	client := onlieclient[chat.Sender]
	client.Channel <- chat;
	delete(onlieclient, chat.Sender)
}

func listenerConn(socket net.Listener) {
	fmt.Println("服务监听启动成功，等待消息...")
	conn, err := socket.Accept();
	if !common.CheckError(err, "服务器监听失败") {
		return
	}
	fmt.Println("连接服务器成功")
	common.MonitorConn(conn, monitorConn)
}

//监听连接
func monitorConn(chat modes.Chat, conn net.Conn) {
	handelData(chat, conn)
}

/**
广播消息
 */
func broadcast(chat modes.Chat) {
	broadcastChat := modes.Chat{chat.Sender, chat.Receiver, constant.CALL_TYPE_BROADCAST, chat.Message, ""}
	fmt.Println(chat.Sender, "对", "所有人说：", chat.Message)
	for _, client := range onlieclient {
		if client.Name != chat.Sender {
			client.Channel <- broadcastChat
		}
	}
}

func p2p(chat modes.Chat) {
	p2pChat := modes.Chat{chat.Sender, chat.Receiver, constant.CALL_TYPE_P2P, chat.Message, ""}
	if client, ok := onlieclient[chat.Receiver]; ok {
		client.Channel <- p2pChat
		fmt.Println(chat.Sender, "对", chat.Receiver, "说：", chat.Message)
	} else {
		fmt.Println(chat.Receiver, "不存在")
	}

}

func main() {
	fmt.Printf("服务端启动，地址：%v,端口号%v\n", constant.SERVER_ADDR, constant.SERVER_PORT)
	listen_socket, err := net.Listen(constant.SERVER_PROTOCOL, common.ServerAddr())
	if !common.CheckError(err, "服务器启动监听失败") {
		return;
	}
	//延时关闭监听
	defer lisenerClose(listen_socket)
	for {
		listenerConn(listen_socket);
	}
}

func lisenerClose(socket net.Listener) {
	socket.Close();
	fmt.Println("socket连接关闭")
}
