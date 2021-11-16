package elastic_plus

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"github.com/hewenyu/toolspackage/elastic_plus/bluk"
	"github.com/hewenyu/toolspackage/systemctl_plus/time_plus"
	"testing"
	"time"
)

type ForTestModels struct {
	ResultStatus string    `json:"resultStatus"`
	Spanid       string    `json:"spanId"`
	Elapsed      int       `json:"elapsed"`
	TraceId      string    `json:"traceId"`
	Hostname     string    `json:"hostname"`
	Appname      string    `json:"appname"`
	RequestId    string    `json:"requestId"`
	ErrorCode    string    `json:"errorCode"`
	MethodName   string    `json:"methodName"`
	Classname    string    `json:"className"`
	Userid       string    `json:"userId"`
	LogTime      time.Time `json:"logTime"`
	Mydate       time.Time `json:"mydate"`
}

func (a ForTestModels) GetJson() ([]byte, error) {
	
	return json.Marshal(a)
}

func TestBulk(t *testing.T) {
	
	timeLocal, err := time_plus.StrToTimeLocal("2021/11/01 16:17:11.122")
	if err != nil {
		return
	}
	
	news1 := ForTestModels{
		ResultStatus: "Y",
		Spanid:       "1",
		Elapsed:      4,
		TraceId:      "ac11000116357546292727666d00db",
		Hostname:     "localhost",
		Appname:      "test",
		RequestId:    "46B7390F78E647E79A587234AB79C67G",
		ErrorCode:    "",
		MethodName:   "aaaa",
		Classname:    "bbbbb",
		Userid:       "123456",
		LogTime:      timeLocal,
		Mydate:       timeLocal,
	}
	
	newRes := make(bluk.BulkDataList, 0)
	
	newRes = append(newRes, bluk.BulkData{
		//
		//Index: "account_audit_2021-11-02",
		Index: "aaa_provider_2021-11-01",
		Data:  bluk.InsertBody{ID: uuid.NewV4().String(), Body: news1},
	})
	
	cfg := NewEsConfig("http://lcoalhost:9200", "aaa", "aaa")
	
	cli, _ := NewClient(cfg)
	
	cli.BulkNewJson(newRes)
	
}
