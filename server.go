/**
* Author: CZ cz.theng@gmail.com
 */

package pbrpc

import (
	"errors"
	"io"
	"net/rpc"
	"sync"
)

var errMissingParams = errors.New("pbrpc: request body missing params")

type serverCodec struct {
	c io.Closer

	mutex sync.Mutex // protects seq, pending
}

// NewServerCodec returns a new rpc.ServerCodec using JSON-RPC on conn.
func NewServerCodec(conn io.ReadWriteCloser) rpc.ServerCodec {
	return &serverCodec{
		c: conn,
	}
}

type serverRequest struct {
}

func (r *serverRequest) reset() {
}

type serverResponse struct {
}

func (c *serverCodec) ReadRequestHeader(r *rpc.Request) error {
	c.mutex.Lock()
	c.mutex.Unlock()

	return nil
}

func (c *serverCodec) ReadRequestBody(x interface{}) error {
	if x == nil {
		return nil
	}
	return nil
}

func (c *serverCodec) WriteResponse(r *rpc.Response, x interface{}) error {
	c.mutex.Lock()
	c.mutex.Unlock()

	return nil
}

func (c *serverCodec) Close() error {
	return c.c.Close()
}

// ServeConn runs the PB-RPC server on a single connection.
// ServeConn blocks, serving the connection until the client hangs up.
// The caller typically invokes ServeConn in a go statement.
func ServeConn(conn io.ReadWriteCloser) {
	rpc.ServeCodec(NewServerCodec(conn))
}
