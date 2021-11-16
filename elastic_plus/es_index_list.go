package elastic_plus

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v6/esapi"
	"io"
	"log"
)

/*
SearchIndex 获取所有的索引
*/
func (c *EsClient) SearchIndex() (hitting ESIndexList, err error) {
	
	req := esapi.CatIndicesRequest{
		Pretty: true,
		Format: "json",
	}
	
	res, err := req.Do(context.Background(), c.Transport)
	
	if err != nil {
		return
	}
	
	// 关闭链接
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(res.Body)
	
	if res.IsError() {
		var e map[string]interface{}
		if errs := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, errs
		} else {
			errs = errors.New(fmt.Sprintf("[%s] %s: %s", res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"]),
			)
			return nil, errs
		}
		
	}
	
	// 返回结果json处理
	if err = json.NewDecoder(res.Body).Decode(&hitting); err != nil {
		
		return nil, err
	}
	return
}
