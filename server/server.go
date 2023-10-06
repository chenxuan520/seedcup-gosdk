package server

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"

	"github.com/chenxuan520/seedcup-gosdk/config"
	"github.com/chenxuan520/seedcup-gosdk/elements"
)

type Conn struct {
	conn net.Conn
}

func CreateConn(config *config.Config) (conn *Conn, err error) {
	conn = &Conn{}
	conn.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", config.Ip, config.Port))
	return conn, err
}

func (conn *Conn) RecvPacket() (*elements.RespPacket, error) {
	packet := &elements.RespPacket{}
	dataBuf, err := conn.recvPacket()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dataBuf, packet)
	return packet, err
}

func (conn *Conn) recvPacket() ([]byte, error) {
	lenBuf := make([]byte, 8)
	_, err := conn.conn.Read(lenBuf)
	if err != nil {
		return nil, err
	}
	length := binary.LittleEndian.Uint64(lenBuf)

	reader := bufio.NewReader(conn.conn)
	readerBuffer, err := ioutil.ReadAll(io.LimitReader(reader, int64(length)))
	if err != nil {
		return nil, err
	}
	return readerBuffer, nil
}

func (conn *Conn) sendPacket(buf []byte) (err error) {
	lenBuf := make([]byte, 8)
	binary.LittleEndian.PutUint64(lenBuf, uint64(len(buf)))
	lenBuf = append(lenBuf, buf...)
	_, err = conn.conn.Write(lenBuf)
	return
}

func (conn *Conn) UpstreamAction(playerID int32, action elements.ActionType) (err error) {
	packet := elements.ReqPacket{
		Type: elements.ActionReq,
		Data: elements.ReqAction{
			PlayerID:   playerID,
			ActionType: action,
		},
	}
	buf, err := json.Marshal(packet)
	if err != nil {
		return nil
	}
	return conn.sendPacket(buf)
}

func (conn *Conn) UpstreamInit() (err error) {
	reqInit := elements.ReqPacket{
		Type: elements.InitReq, Data: elements.ReqInit{},
	}
	buf, err := json.Marshal(reqInit)
	if err != nil {
		return
	}
	err = conn.sendPacket(buf)
	if err != nil {
		return
	}
	return
}

func (conn *Conn) Close() error {
	return conn.conn.Close()
}
