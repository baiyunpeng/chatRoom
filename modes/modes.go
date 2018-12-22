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
	Channel chan string;
}

type Chat struct {
	Name     string
	TellName string
	CallType string
	Message  string
	Group    string
}
