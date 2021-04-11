package message

/*
Message 消息属性抽象

所有的STUN消息都包含20个字节的消息头：16位的消息类型(MessageType)、16位的消息长度(MessageLength)、128位的事务ID(Transaction ID)。

消息类型许可值：0x0001(捆绑请求)、0x0101(捆绑响应)、0x0111(捆绑错误响应)、0x0002(共享私密请求)、0x0102(共享私密响应)、0x0112(共享私密错误响应)。

消息长度：消息大小的字节数，不包括20字节的头部。

事务ID：128位的标识符，用于随机请求和响应，请求与其相应的所有响应具有相同的标识符。
*/
type Message struct {
	Type   MessageType
	Length uint32
	Raw    []byte
}

/*
Reset 将消息设置成空的
*/
func (m *Message) Reset() {
	m.Raw = m.Raw[:0]
	m.Length = 0
}
