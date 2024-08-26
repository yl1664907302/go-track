package elastics

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"go-track/global"
	"log"
	"os"
	"testing"
)

func TestA(t *testing.T) {
	url := "http://192.168.163.66:9200"
	_, err := elastic.NewClient(
		//elastics 服务地址
		elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(global.CONF.System.Elasticsearch.Username, global.CONF.System.Elasticsearch.Password),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	log.Println(err)
}

func TestB(t *testing.T) {
	some, err := SelectDocidBySome("r1_1", "robot_id", "1")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(some)
}
