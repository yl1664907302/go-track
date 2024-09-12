package utils

import (
	"github.com/bytedance/sonic"
	"go-track/pojo"
	"io"
	"log"
	"net/http"
)

// 创建http请求
func SelectAlertsByKey(receiver string, key string, value string) (int, error) {
	var num int
	var alerts []pojo.Alert2
	req, err := http.Get("http://zxknsn.natappfree.cc" + ":" + "80" + "/api/v2/alerts")
	if err != nil {
		return 0, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(req.Body)
	all, err := io.ReadAll(req.Body)
	if err != nil {
		return 0, err
	}

	err = sonic.Unmarshal(all, &alerts)
	if err != nil {
		return 0, err
	}

	for _, a := range alerts {
		key2 := false
		for _, r := range a.Receivers {
			if r.Name == receiver {
				key2 = true
			}
		}
		if key2 && key == "status" {
			if a.Status.State == value {
				num++
			}
		}
	}
	log.Println(num)
	return num, err
}
