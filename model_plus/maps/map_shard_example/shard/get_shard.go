package shard

import (
	"github.com/hewenyu/toolspackage/model_plus/maps/map_shard_example/base_shard"
	"github.com/hewenyu/toolspackage/systemctl_plus/hash"
)

// GetShard 获取分片
func (q *CountShard) GetShard(key string) *base_shard.CountMap {
	return q.data[uint(hash.Fnv32(key))%uint(q.shardTotal)]
}
