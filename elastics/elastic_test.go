package elastics

import (
	"fmt"
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

func TestB(t *testing.T) {
	some, err := SelectDocidBySome("l1_r", "robot_id", "1")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(some)
}

func TestSelectDocidBySome(t *testing.T) {
	type args struct {
		index string
		key   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			args: args{
				index: "l1_t",
				key:   "maketime",
				value: "2024-07-17 03:42:26",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectDocidBySome(tt.args.index, tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectDocidBySome() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SelectDocidBySome() got = %v, want %v", got, tt.want)
			}
		})
	}
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
	type args struct {
		message []pojo.Alerts
		index   string
	}
	tests := []struct {
		name  string
		args  args
		want  error
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CreateIndexESForAlert(tt.args.message, tt.args.index)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexESForAlert() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CreateIndexESForAlert() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
