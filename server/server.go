package main

import (
	"net"
	"fmt"
	"github.com/baiyunpeng/chatRoom/const"
	"github.com/baiyunpeng/chatRoom/common"
	"github.com/baiyunpeng/chatRoom/modes"
	"encoding/json"
)

//储存在线用户的map
var onlieclient = make(map[string]modes.Client)
//管道对象
var channelChatChannel = make(chan modes.Chat)

//遍历管道消息
func handelData() {
	for {
		chat := <-channelChatChannel
		//连接
		if chat.CallType == constant.CALL_TYPE_COMMOND {
			handelCommonData(chat)
		} else if chat.CallType == constant.CALL_TYPE_GROUP {
			fmt.Println("CALL_TYPE_GROUP", chat)
		} else if chat.CallType == constant.CALL_TYPE_BROADCAST {
			fmt.Println("CALL_TYPE_BROADCAST", chat)
		} else if chat.CallType == constant.CALL_TYPE_P2P {
			fmt.Println("CALL_TYPE_P2P", chat)
		}
	}
}

func handelCommonData(chat modes.Chat){
	if chat.Message == constant.CALL_COMMON_CONNECTION {
		registerClietn(chat);
	}else if chat.Message == constant.CALL_COMMON_CLOSE {
		fmt.Println("CALL_COMMON_CLOSE",chat)
	}
}
func registerClietn(chat modes.Chat) {
	client := modes.Client{chat.Name, chat.Name, make(chan string)}
	onlieclient[chat.Name] = client;
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
func monitorConn(chat modes.Chat) {
	channelChatChannel <- chat;
}

/**
广播消息
 */
func broadcast(message string) {
	chat := modes.Chat{"", "", constant.CALL_TYPE_BROADCAST, message, ""}
	messageByte,err:=json.Marshal(chat);
	if common.CheckError(err,"数据转换失败"){
		for _, client := range onlieclient {
			client.Channel
		}
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
	go handelData();
	for {
		listenerConn(listen_socket);
	}
}

func lisenerClose(socket net.Listener) {
	socket.Close();
	fmt.Println("socket连接关闭")
}
