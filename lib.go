package lib

type connection struct{}

type HttpClient struct {
	LastRequest string
	Conn        connection
}

func NewClient(address string) HttpClient {
}
