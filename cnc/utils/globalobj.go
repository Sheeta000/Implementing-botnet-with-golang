package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type GlobalObj struct {
	ServerAddr string
	ServerPort string

	TelnetName     string
	TelnetPassWord string
	TelnetPort     string
	TelnetAddr     string

	ClientVerify         string
	ServerVerifyCallBack string
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	file, err := ioutil.ReadFile("info.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &GlobalObject)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	obj := GlobalObj{
		ServerPort:           "12312",
		ServerAddr:           "127.0.0.1",
		TelnetAddr:           "127.0.0.1",
		TelnetPort:           "23",
		TelnetName:           "boxboz",
		TelnetPassWord:       "wowxc60",
		ClientVerify:         "EDGVTDYTHWOTCCVCXIFHOSUOTPKAFOMX",
		ServerVerifyCallBack: "XPTMTVKSPKZGOQODIOUJOXXACMOWLNTK",
	}
	obj.Reload()
}
