/**
* Author: CZ cz.theng@gmail.com
 */

package pbrpc

import (
	"io"
	"net"
	"net/rpc"
	"sync"
)

type clientCodec struct {
	c io.Closer

	// temporary work space
	req clientRequest
	rsp clientResponse

	mutex sync.Mutex // protects pending
}

// NewClientCodec returns a new rpc.ClientCodec using JSON-RPC on conn.
func NewClientCodec(conn io.ReadWriteCloser) rpc.ClientCodec {
	return &clientCodec{
		c: conn,
	}
}

type clientRequest struct {
}

func (c *clientCodec) WriteRequest(r *rpc.Request, param interface{}) error {
	c.mutex.Lock()
	c.mutex.Unlock()
	return nil
}

type clientResponse struct {
}

func (r *clientResponse) reset() {
}

func (c *clientCodec) ReadResponseHeader(r *rpc.Response) error {
	c.mutex.Lock()
	c.mutex.Unlock()

	return nil
}

func (c *clientCodec) ReadResponseBody(x interface{}) error {
	if x == nil {
		return nil
	}
	return nil
}

func (c *clientCodec) Close() error {
	return c.c.Close()
}

// NewClient returns a new rpc.Client to handle requests to the
// set of services at the other end of the connection.
func NewClient(conn io.ReadWriteCloser) *rpc.Client {
	return rpc.NewClientWithCodec(NewClientCodec(conn))
}

// Dial connects to a PB-RPC server at the specified network address.
func Dial(network, address string) (*rpc.Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return NewClient(conn), err
}
