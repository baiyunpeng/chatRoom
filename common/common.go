package common

import (
	"net"
	"fmt"
	"../const"
	"../modes"
	"encoding/json"
)

type connListener func(modes.Chat, net.Conn)

/**
接收数据
 */
func Receive(conn net.Conn) ([]byte, error) {
	var message [] byte = nil;
	receiveData := make([]byte, 255)
	index, err := conn.Read(receiveData);
	message = receiveData[0:index];
	return message, err;
}

/**
发送消息
 */
func SendMessage(conn net.Conn, message string) {
	conn.Write([]byte(message))
}

/**
hu获取服务器地址
 */
func ServerAddr() string {
	return constant.SERVER_ADDR + ":" + constant.SERVER_PORT;
}

func MonitorConn(conn net.Conn, listener connListener) {
	go func() {
		chat := &modes.Chat{};
		var sender string = "";
		for {
			message, err := Receive(conn);
			if CheckError(err, "读取数据失败") {
				err := json.Unmarshal(message, chat);
				sender = chat.Sender;
				if CheckError(err, "转换消息失败") {
					listener(*chat, conn);
				}
			} else {
				fmt.Println("数据连接断开")
				chat.Sender = sender
				chat.CallType = constant.CALL_TYPE_COMMOND
				chat.Message = constant.CALL_COMMON_CLOSE
				listener(*chat, conn)
				break;
			}

		}
	}()
}

func MonitorChat(conn net.Conn, client modes.Client) {
	go func() {
		for {
			chat := <-client.Channel
			messageByte, err := json.Marshal(chat)
			if CheckError(err, "转换消息失败") {
				if (chat.CallType == constant.CALL_TYPE_COMMOND && chat.Message == constant.CALL_COMMON_CLOSE) {
					fmt.Println("关闭监听")
					break;
				}
				conn.Write(messageByte)
			}
		}
	}()
}

/**
错误检查
 */
func CheckError(err error, message string) bool {
	if (nil != err) {
		fmt.Println(message + err.Error())
		return false;
	}
	return true;
}
