package main

import (
	"encoding/binary"
	"flag"
	"net"
	"os"

	"leafserver/src/server/conf"
	"leafserver/src/server/game"
	"leafserver/src/server/gate"
	"leafserver/src/server/login"

	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	//"github.com/gorilla/websocket"
)

func clientwsrun() (ret int, err error) {
  //FIXNE NOT YET IMPEMENT
  ret = 0
	err = nil
	return
}

func clientrun() (ret int, err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}

	// Hello message (JSON-encoded)
	// The structure of the message
	data := []byte(`{
        "Hello": {
            "Name": "leaf"
        }
    }`)

	// len + data
	m := make([]byte, 2+len(data))

	// BigEndian encoded
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)

	// Send message
	conn.Write(m)

	ret = 0
	err = nil
	return
}

func main() {

	clientFlag := flag.Bool("client", false, "launch client")
	flag.Parse()
	if *clientFlag {
		ret, err := clientrun()
		if err != nil {
			panic("leafclient: exit:" + string(ret) + " ,error: " + err.Error())
		}
		os.Exit(ret)
	}

	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
