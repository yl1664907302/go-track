package elastics

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/olivere/elastic/v7"
	"go-track/pojo"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestA(t *testing.T) {
	url := "http://192.168.163.66:9200"
	_, err := elastic.NewClient(
		//elastics 服务地址
		elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetBasicAuth("elastic", "123456"),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	log.Println(err)
}

func TestSelectNewDocByindex(t *testing.T) {
	type args struct {
		index string
		key   string
		any   interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			args: args{index: "l1_t",
				key: "maketime.keyword",
				any: pojo.Markdown{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectNewDocByindex(tt.args.index, tt.args.key, tt.args.any)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectNewDocByindex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectNewDocByindex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexESForAlert(t *testing.T) {
	var robot pojo.Robot
	byindex, err := SelectNewDocByindex("l1_r", "robot_id.keyword", &pojo.Robot{})
	if err != nil {
		log.Println(err)
	}
	err = sonic.Unmarshal(byindex, &robot)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(robot)
}
