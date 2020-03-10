package krpc

import "net"

type RpcContext struct {
	index int
	chain []RpcHandlerFunc
	addr  *net.UDPAddr
	req   Request
	resp  Response
}

func NewReqContext(chain []RpcHandlerFunc, req Request, resp Response, addr *net.UDPAddr) *RpcContext {
	return &RpcContext{
		chain: chain,
		index: -1,
		req:   req,
		resp:  resp,
		addr:  addr,
	}
}

func (c *RpcContext) Next() {
	c.index++
	for ; c.index < len(c.chain); c.index++ {
		c.chain[c.index](c)
	}
}

func (c *RpcContext) Request() Request {
	return c.req
}

func (c *RpcContext) Response() Response {
	return c.resp
}

func (c *RpcContext) WriteAs(resp Response) {
	c.resp = resp
}

func (c *RpcContext) RemoteAddr() *net.UDPAddr {
	return c.addr
}
