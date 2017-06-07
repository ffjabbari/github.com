package main

import (
	"fmt"
	"github.com/lixiangzhong/httpclient"
)

func main() {
	c := httpclient.New()
	c.PostForm("www.google.com/example/api")
	c.Param.Add("key", "value")
	res, err := c.Do()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.String())
}
