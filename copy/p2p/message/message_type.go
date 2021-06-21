package message

/*
Method is uint16 representation of 12-bit STUN method.
*/
type Method uint16

/*
MessageClass是STUN消息类的2位类的8位表示形式。
*/
type MessageClass byte

/*
MessageType 是 STUN 消息中的请求区域
*/
type MessageType struct {
	Method Method       // e.g. binding
	Class  MessageClass // e.g. request
}
