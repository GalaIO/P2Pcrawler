package krpc

type RpcContext struct {
	index int
	chain []RpcHandlerFunc
	req   Request
	resp  Response
}

func NewReqContext(chain []RpcHandlerFunc, req Request, resp Response) *RpcContext {
	return &RpcContext{chain: chain, index: -1, req: req, resp: resp}
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
