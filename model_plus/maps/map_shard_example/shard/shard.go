package shard

import "github.com/hewenyu/toolspackage/model_plus/maps/map_shard_example/base_shard"

type CountShard struct {
	data       []*base_shard.CountMap
	shardTotal int
}

// NewCountShard
// 使用分片机制,减少锁的MAP
func NewCountShard(shardTotal int) *CountShard {
	
	_shard := CountShard{
		data:       make([]*base_shard.CountMap, 0),
		shardTotal: shardTotal,
	}
	
	for i := 0; i < shardTotal; i++ {
		_shard.data = append(_shard.data, base_shard.NewCountMap())
	}
	
	return &_shard
	
}
