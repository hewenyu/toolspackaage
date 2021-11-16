package httpc

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

/*
WebGet 获取信息
*/
type WebGet struct {
	Host    string            `json:"host"`
	Url     string            `json:"url"`
	UrlData map[string]string `json:"urldata"`
}

/*
ResPonse 返回信息
*/
type ResPonse struct {
	Status string `json:"status"`
	Body   []byte `json:"body"`
}

/*
FormatUrl 输出完整的url
官方推荐 build 然每次分配的内存次数有点多，但是每次分配的内存大小比buffer要少。
*/
func (v *WebGet) FormatUrl() (b strings.Builder) {

	values := url.Values{}

	for ks, vs := range v.UrlData {
		values.Set(ks, vs)

	}

	formatUrl := values.Encode()

	b.WriteString(v.Host)
	b.WriteString(v.Url)
	b.WriteString("?")
	b.WriteString(formatUrl)

	return

}

func (w *WebGet) InvokeGet() (res ResPonse, err error) {

	url := w.FormatUrl()

	log.Println("请求接口", url.String())
	// 参数格式化
	req, err := http.NewRequest("GET", url.String(), nil)

	if err != nil {
		return
	}
	// 发起请求
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Trident/7.0; rv:11.0) like Gecko")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// 输出

	res.Body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	res.Status = resp.Status

	return
}

func StructToMaps(b interface{}) (res map[string]string) {
	val := reflect.ValueOf(b)

	res = make(map[string]string)

	for i := 0; i < val.Type().NumField(); i++ {
		// 获取名称
		key := val.Type().Field(i).Name
		// 获取数据
		value := reflect.Indirect(val).FieldByName(key)
		// 获取json的key
		jsonKey := val.Type().Field(i).Tag.Get("json")
		// 设置map 强制转换成string 格式
		strValue := fmt.Sprintf("%v", value)
		// 空的数据就不需要
		if strValue != "" {
			res[jsonKey] = strValue
		}

	}

	return
}
