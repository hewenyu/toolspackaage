package client

import (
	"fmt"
	"time"

	"gortc.io/stun"
)

/*
HeatCheck 心跳检测
*/
func HeatCheck(c *stun.Client) {
	var xorAddr stun.XORMappedAddress
	for {
		message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)

		if err := c.Do(message, func(res stun.Event) {
			if res.Error != nil {
				panic(res.Error)
			}
			// Decoding XOR-MAPPED-ADDRESS attribute from message.

			if err := xorAddr.GetFrom(res.Message); err != nil {
				panic(err)
			}
			fmt.Println("your IP is", xorAddr.IP, "PORT", xorAddr.Port, "心跳存活")
		}); err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 5)
	}
}

func StunClient() {

	// 创建一个stun链接
	c, err := stun.Dial("udp", "124.71.182.117:9527")
	if err != nil {
		panic(err)
	}

	go HeatCheck(c)
	// message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)

	// c.Start(message, func(res stun.Event) {
	// 	if res.Error != nil {
	// 		panic(res.Error)
	// 	}

	// 	fmt.Println(res.Message)
	// })

	for {
		time.Sleep(time.Second * 10)
	}
}
