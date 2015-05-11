package base
import "net/http"

// http 请求
type Request struct {
	httpReq *http.Request
	depth uint32
}

func NewRequest(httpReq *http.Request,depth uint32) *Request {
	return &Request{httpReq,depth}
}

func (req *Request) HttpReq() *http.Request {
	return req.httpReq
}

func (req *Request) Depth() uint32 {
	return req.depth
}