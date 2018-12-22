package common

import (
	"net"
	"fmt"
	"github.com/baiyunpeng/chatRoom/const"
	"sync"
)

var readLock sync.Mutex
var writeLock sync.Mutex
/**
接收数据
 */
func Receive(conn net.Conn) string {
	data := make([]byte, 255)
	index, err := conn.Read(data);
	if (err != nil) {
		fmt.Println("数据读取失败")
	}
	var msg = string(data[0:index])
	fmt.Printf("接收到数据%v\n", msg)
	return msg;
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
