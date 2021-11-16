package client

// func XClient() {
// 	var serverAddr = flag.String("s", "stun.sipgate.net:3478", "STUN server address")
// 	var v = flag.Bool("v", false, "verbose mode")
// 	var vv = flag.Bool("vv", false, "double verbose mode (includes -v)")
// 	var vvv = flag.Bool("vvv", false, "triple verbose mode (includes -v and -vv)")
// 	flag.Parse()

// 	// Creates a STUN client. NewClientWithConnection can also be used if you want to handle the UDP listener by yourself.
// 	client := stun.NewClient()
// 	// The default addr (stun.DefaultServerAddr) will be used unless we call SetServerAddr.
// 	client.SetServerAddr(*serverAddr)
// 	// Non verbose mode will be used by default unless we call SetVerbose(true) or SetVVerbose(true).
// 	client.SetVerbose(*v || *vv || *vvv)
// 	client.SetVVerbose(*vv || *vvv)
// 	// Discover the NAT and return the result.
// 	nat, host, err := client.Discover()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println("NAT Type:", nat)
// 	if host != nil {
// 		fmt.Println("External IP Family:", host.Family())
// 		fmt.Println("External IP:", host.IP())
// 		fmt.Println("External Port:", host.Port())
// 	}

// }
