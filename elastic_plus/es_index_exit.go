package elastic_plus

import (
	"context"
	"github.com/elastic/go-elasticsearch/v6/esapi"
	"io"
	"log"
)

// IndexExit 判断索引是否存在
func (c *EsClient) IndexExit(indices ...string) (exitIdex bool) {
	
	req := esapi.IndicesExistsRequest{
		Index: indices,
	}
	
	res, _ := req.Do(context.Background(), c.Transport)
	
	// 关闭链接
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(res.Body)
	
	// 判断是否存在索引
	if res.IsError() {
		exitIdex = false
		return
	}
	
	exitIdex = true
	
	return
}
