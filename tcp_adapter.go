package go2p

import (
	"fmt"
	"net"
)

type adapterTCP struct {
	conn net.Conn
}

func NewAdapter(conn net.Conn) Adapter {
	a := new(adapterTCP)
	a.conn = conn
	return a
}

func (a *adapterTCP) ReadMessage() (*Message, error) {
	m := NewMessage()
	err := m.ReadFromConn(a.conn)
	return m, err
}

func (a *adapterTCP) WriteMessage(m *Message) error {
	err := m.WriteIntoConn(a.conn)
	return err
}

func (a *adapterTCP) Close() {
	a.conn.Close()
}

func (a *adapterTCP) Address() string {
	addr := a.conn.RemoteAddr()
	res := fmt.Sprintf("%s:%s", addr.Network(), addr.String())
	return res
}
