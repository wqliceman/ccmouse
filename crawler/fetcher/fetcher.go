package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

func Fetch(url string) ([]byte, error)  {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("抓取出错了, 返回码[%d]", resp.StatusCode)
	}
	bufBody := bufio.NewReader(resp.Body)
	utf8Reader := transform.NewReader(bufBody, determineEncoding(bufBody).NewDecoder())
	body, err := ioutil.ReadAll(utf8Reader)
	return body, err
}


func determineEncoding(r *bufio.Reader) encoding.Encoding {
	data, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(data, "")
	return e
}