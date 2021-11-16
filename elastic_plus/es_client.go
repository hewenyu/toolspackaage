package elastic_plus

import (
	"github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/esapi"
	"github.com/elastic/go-elasticsearch/v6/estransport"
)

// EsClient 客户端struct
type EsClient struct {
	*esapi.API // 启用es api 模式
	Transport  estransport.Interface
}

// NewClient 构造客户端
func NewClient(cfg *EsConfig) (*EsClient, error) {
	
	var tmpcfg elasticsearch.Config
	tmpcfg.Addresses = cfg.Addresses
	tmpcfg.Username = cfg.Username
	tmpcfg.Password = cfg.Password
	
	client, err := elasticsearch.NewClient(tmpcfg)
	
	esClient := &EsClient{Transport: client.Transport, API: client.API}
	
	return esClient, err
}
