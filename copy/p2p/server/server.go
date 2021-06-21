package server

// type ServerD struct {
// 	Conn *net.UDPConn
// }

// /*
// ListenUDPAndServe 监听端口并打开服务
// */
// func (this ServerD) ListenUDPAndServe(serverNet, laddr string) error {
// 	c, err := net.ListenPacket(serverNet, laddr)
// 	if err != nil {
// 		return err
// 	}
// 	s := &Server{
// 		log: defaultLogger,
// 	}
// 	return s.Serve(c)
// }

// /*
// RunServer 启动服务,目前只支持udp协议
// */
// func RunServer() {
// 	flag.Parse()
// 	if *profile {
// 		go func() {
// 			log.Println(http.ListenAndServe("localhost:9527", nil))
// 		}()
// 	}
// 	switch *network {
// 	case "udp":
// 		normalized := normalize(*address)
// 		logger.Infof("Server listening on", normalized, "via", *network)
// 		logger.Fatal(ListenUDPAndServe(*network, normalized))
// 	default:
// 		logger.Fatal("unsupported network:", *network)
// 	}
// }

// // func normalize(address string) string {
// // 	if len(address) == 0 {
// // 		address = "0.0.0.0"
// // 	}
// // 	if !strings.Contains(address, ":") {
// // 		address = fmt.Sprintf("%s:%d", address, NewPort)
// // 	}
// // 	return address
// // }

// // func RunServer() {
// // 	flag.Parse()
// // 	if *profile {
// // 		go func() {
// // 			log.Println(http.ListenAndServe("localhost:6060", nil))
// // 		}()
// // 	}
// // 	switch *network {
// // 	case "udp":
// // 		normalized := normalize(*address)
// // 		fmt.Println("gortc/stund listening on", normalized, "via", *network)
// // 		log.Fatal(ListenUDPAndServe(*network, normalized))
// // 	default:
// // 		log.Fatalln("unsupported network:", *network)
// // 	}
// // }

// func ServerDemo() {

// 	// var peers []net.UDPAddr
// 	// 初始化
// 	// peers = make([]net.UDPAddr, 0)

// 	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 9527})
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	logger.Info("本地地址: <" + listener.LocalAddr().String() + ">")

// 	data := make([]byte, 1024)
// 	for {
// 		n, remoteAddr, err := listener.ReadFromUDP(data)
// 		if err != nil {
// 			fmt.Printf("error during read: %s", err)
// 		}
// 		// 传递消息
// 		message := string(data[:n])

// 		switch message {
// 		case "health":
// 			logger.Infof("来自 <%s> 的心跳检查", remoteAddr.String())
// 			listener.WriteToUDP([]byte("心跳检查OK"), remoteAddr)
// 		default:
// 			logger.Infof("<%s>", message, remoteAddr.String())

// 			listener.WriteToUDP([]byte("已经收到消息:"+remoteAddr.String()), remoteAddr)
// 			// 将打洞加入server
// 			// peers = append(peers, *remoteAddr)
// 			// if len(peers) == 2 {
// 			// 	logger.Infof("进行UDP打洞,建立 %s <--> %s 的连接\n", peers[0].String(), peers[1].String())
// 			// 	listener.WriteToUDP([]byte(peers[1].String()), &peers[0])
// 			// 	listener.WriteToUDP([]byte(peers[0].String()), &peers[1])
// 			// }
// 		}

// 	}
// }
