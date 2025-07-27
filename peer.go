package main

import "net"

type Peer struct {
	conn  net.Conn
	msgCh chan Message
	delCh chan *Peer
}

func (p *Peer) Send(msg []byte) (int,error){
	return p.conn.Write(msg)
}

func NewPeer(conn net.Conn, msgCh chan Message, delCh chan *Peer) *Peer {
	return &Peer{
		conn:  conn,
		msgCh: msgCh,
		delCh: delCh,
	}
}
