package server

// func Server() {
// 	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 9527})
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	logger.Infof("本地地址: <%s> \n", listener.LocalAddr().String())
// 	// peers := make([]net.UDPAddr, 0, 2)
// 	data := make([]byte, 1024)
// 	for {
// 		n, remoteAddr, err := listener.ReadFromUDP(data)
// 		if err != nil {
// 			fmt.Printf("error during read: %s", err)
// 		}
// 		logger.Infof("<%s> %s\n", remoteAddr.String(), data[:n])

// 		listener.WriteToUDP([]byte("已经收到消息了"), remoteAddr)

// 		// peers = append(peers, *remoteAddr)
// 		// if len(peers) == 2 {
// 		// 	logger.Infof("进行UDP打洞,建立 %s <--> %s 的连接\n", peers[0].String(), peers[1].String())
// 		// 	listener.WriteToUDP([]byte(peers[1].String()), &peers[0])
// 		// 	listener.WriteToUDP([]byte(peers[0].String()), &peers[1])
// 		// 	time.Sleep(time.Second * 8)
// 		// 	logger.Info("中转服务器退出,仍不影响peers间通信")
// 		// 	return
// 		// }
// 	}
// }
