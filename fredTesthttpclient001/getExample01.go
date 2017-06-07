package main

import (
	"fmt"
	"github.com/lixiangzhong/httpclient"
	"time"
)

func main() {
	c := httpclient.New()
	c.Get("www.google.com")//if url Scheme=="" default:http://www.google.com
	c.SetTimeout(3 * time.Second)
	c.Query.Add("search", "news")
	res, err := c.Do()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.StatusCode)
}
