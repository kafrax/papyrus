package payrus

import "math"

const (
	abortIndex uint8 = math.MaxUint8 / 2
)

type HandlerFunc func(context *Context)
type HandlerList []HandlerFunc

// Endpoint to add middleware what you want to do when pay channel action will doing before.
// like secret,validator,insert order to db etc.
type Endpoint func(ctx Context, requestBody interface{}) (ret interface{}, err Error)
type Middleware func(endpoint Endpoint)Endpoint

// Context is all of the pay channel to use for execute operation like pay or transfer etc.
type Context struct {
	handlers HandlerList
	errors   Error
	abortIndex    uint8 //now handlerFunc's index.
	params   map[string]string
}

func (c *Context) Reset() {

}

func (c *Context) Clone() *Context {
	var cp = *c
	cp.handlers=nil
	cp.errors=Error{}
	cp.abortIndex=-1
	return &cp
}

func (c *Context) Next() {

}

func (c *Context) Abort() {
	c.abortIndex = abortIndex
}

func (c *Context) AbortWithErr(msg string, t ErrorType, err error) Error {
	c.Abort()
	return ErrorsNew(msg, t, err)
}

func (c *Context) IsAbort() bool {
	return c.abortIndex >= abortIndex
}

func (c *Context) LenHandler() int {
	return len(c.handlers)
}

//set other param for action
func (c *Context) SetParam(key, value string) {

}

//get other param from pay channel
func (c *Context) GetParam(key string) string {
	return ""
}
