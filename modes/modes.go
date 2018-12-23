package modes

/**
客户端结构体
 */
type Client struct {
	//地址
	Addr string
	//姓名
	Name string
	//通讯内容
	Channel chan Chat;
}

type Chat struct {
	Sender   string
	Receiver string
	CallType string
	Message  string
	Group    string
}
