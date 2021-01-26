package client

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hewenyu/toolspackage/logger"
)

var tag string

const HAND_SHAKE_MSG = "我是打洞消息"

/*
HeatCheck 心跳检测
*/
func health_check(conn *net.UDPConn) {
	var (
		data []byte
		err  error
	)
	data = make([]byte, 1024)
	for {
		// 发送 health 检查
		if _, err = conn.Write([]byte("health")); err != nil {
			logger.Panic(err)
		}
		n, remoteAddr, err := conn.ReadFromUDP(data)

		if err != nil {
			fmt.Printf("error during read: %s", err)
		}

		logger.Info(string(data[:n]), remoteAddr.String())

		time.Sleep(time.Second * 5)
	}
}

func Client() {
	if len(os.Args) < 2 {
		fmt.Println("请输入一个客户端标志")
		os.Exit(0)
	}

	// 当前进程标记字符串,便于显示
	tag = os.Args[1]

	logger.Info(tag)
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 9901} // 注意端口必须固定

	dstAddr := &net.UDPAddr{IP: net.ParseIP("124.71.182.117"), Port: 9527}

	logger.Info("本地 " + srcAddr.String() + "链接 远程 " + dstAddr.String())

	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	defer conn.Close()

	if err != nil {
		fmt.Println(err)
	}
	if _, err = conn.Write([]byte("hello, I'm new peer:" + tag)); err != nil {
		logger.Panic(err)
	}
	data := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Printf("error during read: %s", err)
	}

	// 获取UDP 的信息
	logger.Info(string(data[:n]))

	logger.Info(n, remoteAddr.String())

	go health_check(conn)

	// anotherPeer := parseAddr(string(data[:n]))
	// fmt.Printf("local:%s server:%s another:%s\n", srcAddr, remoteAddr, anotherPeer.String())
	// // 开始打洞
	// bidirectionHole(srcAddr, &anotherPeer)
}
func parseAddr(addr string) net.UDPAddr {
	t := strings.Split(addr, ":")
	port, _ := strconv.Atoi(t[1])
	return net.UDPAddr{
		IP:   net.ParseIP(t[0]),
		Port: port,
	}
}
func bidirectionHole(srcAddr *net.UDPAddr, anotherAddr *net.UDPAddr) {
	conn, err := net.DialUDP("udp", srcAddr, anotherAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	// 向另一个peer发送一条udp消息(对方peer的nat设备会丢弃该消息,非法来源),用意是在自身的nat设备打开一条可进入的通道,这样对方peer就可以发过来udp消息
	if _, err = conn.Write([]byte(HAND_SHAKE_MSG)); err != nil {
		logger.Info("send handshake:", err)
	}
	go func() {
		for {
			time.Sleep(10 * time.Second)
			if _, err = conn.Write([]byte("from [" + tag + "]")); err != nil {
				logger.Info("send msg fail", err)
			}
		}
	}()
	for {
		data := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(data)
		if err != nil {
			logger.Info("error during read: %s\n", err)
		} else {
			logger.Info("收到数据:%s\n", data[:n])
		}
	}
}
