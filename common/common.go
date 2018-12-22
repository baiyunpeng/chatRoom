package common

import (
	"net"
	"fmt"
	"github.com/baiyunpeng/chatRoom/const"
	"github.com/baiyunpeng/chatRoom/modes"
	"github.com/baiyunpeng/chatRoom/common"
	"encoding/json"
)

type ConnListener func(modes.Chat)

/**
接收数据
 */
func Receive(conn net.Conn) []byte {
	var message [] byte = nil;
	receiveData := make([]byte, 255)
	index, err := conn.Read(receiveData);
	if (CheckError(err, "数据读取失败")) {
		message = receiveData[0:index];
	}
	return message;
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

func MonitorConn(conn net.Conn, listener ConnListener) {
	go func() {
		chat := &modes.Chat{};
		for {
			message := common.Receive(conn);
			err := json.Unmarshal(message, chat);
			if common.CheckError(err, "转换消息失败") {
				listener(*chat);
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
