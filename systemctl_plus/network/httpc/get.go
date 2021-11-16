package httpc

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Get() {

	params := url.Values{}

	Url, err := url.Parse("http://baidu.com?fd=fdsf")
	if err != nil {
		panic(err.Error())

	}
	params.Set("a", "fdfds")
	params.Set("id", string("1"))
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)

	if err != nil {
		panic(err.Error())

	}

	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())

	}
	fmt.Println(string(s))
}
