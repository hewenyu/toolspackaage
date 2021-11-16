package elastic_plus

import (
	"context"
	"encoding/json"
	"github.com/hewenyu/toolspackage/elastic_plus/bluk"
	"github.com/hewenyu/toolspackage/global_variable"
	"io"
	"strings"
)

// BulkNewJson ES 创建数据 批量
func (c *EsClient) BulkNewJson(insertData bluk.BulkDataList) bool {
	
	var ctx = context.Background()
	
	res, err := c.Bulk(
		strings.NewReader(insertData.ToString()),
		c.Bulk.WithRefresh("true"),
		c.Bulk.WithPretty(),
		c.Bulk.WithTimeout(100),
		c.Bulk.WithContext(ctx),
	)
	
	if err != nil {
		global_variable.CLOG.Sugar().Errorf("Unexpected error when getting a response: %s", err)
	}
	
	// req := esapi.BulkRequest{
	// 	Body:    strings.NewReader(insert_data.ToString()),
	// 	Refresh: "true",
	// 	Pretty:  true,
	// 	Timeout: 100,
	// }
	
	// res, _ := req.Do(ctx, c.Transport)
	
	// 关闭链接
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			global_variable.CLOG.Error(err.Error())
		}
	}(res.Body)
	
	if res.IsError() {
		var e map[string]interface{}
		
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			global_variable.CLOG.Sugar().Infof("Error parsing the response body: %s \n", err)
		} else {
			// 日志答应详细信息
			global_variable.CLOG.Sugar().Infof("[%s] %s: %s \n", res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
		return false
	}
	
	global_variable.CLOG.Info("日志归档完成")
	
	return true
	
}
