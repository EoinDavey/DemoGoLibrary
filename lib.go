package lib

type connection struct{}

type HttpClient struct {
	LastRequest string
	Conn        connection
}

func NewClient(address string) HttpClient {
}

func OldClient(address string) HttpClient {
}

func AmazingClient(address string) HttpClient {
}

func SpeactularClient(address string) HttpClient {
}

func ElderlyClient(address string) HttpClient {
}

func TcpClient(address string) HttpClient {
}
