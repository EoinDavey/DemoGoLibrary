package lib

type connection struct{}

type HttpClient struct {
	LastRequest string
	Conn        connection
}

func NewClient(address string) HttpClient {
    return HttpClient{}
}

func (h *HttpClient) RunDownload() {
}

func (h *HttpClient) ChangeAddress(newAddr, oldAddr string) bool {
    return true
}

type UploadClient struct {
    Conn connection
    Details struct{
        Target string
    }
}

func unexported() error {
    return nil
}

func Upload(u *UploadClient) error {
    return nil
}
