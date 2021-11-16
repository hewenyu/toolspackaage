package map_shard_example

import "github.com/hewenyu/toolspackage/systemctl_plus/hash"

// GetShard 获取分片
func (q *CountShard) GetShard(key string) *CountMap {
	return q.data[uint(hash.Fnv32(key))%uint(q.shardTotal)]
}
