package cnc

import (
	"encoding/json"
	"fmt"
)

func jsonDecoding(msg BotVerify, buf []byte, len int) string {
	err := json.Unmarshal(buf[:len], &msg)
	if err != nil {
		fmt.Println("[-]json解码失败", err)
		return ""
	}
	return msg.Verify
}
