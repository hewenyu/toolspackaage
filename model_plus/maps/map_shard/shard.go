package map_shard

type CountShard struct {
	data       []*CountMessage
	shardTotal int
}

// NewCountShard
// 使用分片机制,减少锁的MAP
func NewCountShard(shardTotal int) *CountShard {
	
	_shard := CountShard{
		data:       make([]*CountMap, 0),
		shardTotal: shardTotal,
	}
	
	for i := 0; i < shardTotal; i++ {
		_shard.data = append(_shard.data, NewCountMap())
	}
	
	return &_shard
	
}
