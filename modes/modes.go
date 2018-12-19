package modes

/**
客户端结构体
 */
type Client struct {
	//地址
	addr string
	//姓名
	name string
	//通讯管道
	channel chan string
}
