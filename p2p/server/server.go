package server

import (
	"fmt"
	"net"

	"github.com/hewenyu/toolspackage/logger"
)

func ServerDemo() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 9527})
	if err != nil {
		fmt.Println(err)
		return
	}
	logger.Info("本地地址: <" + listener.LocalAddr().String() + ">")

	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		// 传递消息
		message := string(data[:n])

		switch message {
		case "health":
			logger.Infof("来自 <%s> 的心跳检查", remoteAddr.String())
			listener.WriteToUDP([]byte("心跳检查OK"), remoteAddr)
		default:
			logger.Infof("<%s>", message, remoteAddr.String())
			listener.WriteToUDP([]byte("已经收到消息:"+remoteAddr.String()), remoteAddr)
		}

		// peers = append(peers, *remoteAddr)
		// if len(peers) == 2 {
		// 	logger.Infof("进行UDP打洞,建立 %s <--> %s 的连接\n", peers[0].String(), peers[1].String())
		// 	listener.WriteToUDP([]byte(peers[1].String()), &peers[0])
		// 	listener.WriteToUDP([]byte(peers[0].String()), &peers[1])
		// 	time.Sleep(time.Second * 8)
		// 	logger.Info("中转服务器退出,仍不影响peers间通信")
		// 	return
		// }
	}
}
