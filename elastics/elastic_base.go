package elastics

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

func DelIndex(index string) error {
	ESclient, err := GetEsClient()
	ctx := context.Background()
	if err != nil {
		return err
	}
	_, err = ESclient.DeleteIndex(index).Do(ctx)
	if err != nil {
		return err
	} else {
		log.Printf("已删除索引: %s", index)
	}
	return err
}

func DelDocByKey(index string, key string, value string) error {
	//查询是否存在该文档
	doc_id, err := SelectDocidBySome(index, key, value)
	if doc_id == "" {
		log.Printf("索引%s中不存在%s值为%s的doc", index, key, value)
		return fmt.Errorf("索引%s中不存在%s值为%s的doc", index, key, value)
	} else {
		err := DelDoc(index, doc_id)
		if err != nil {
			return err
		}
	}
	return err
}

func DelDoc(index string, doc_id string) error {
	ESclient, err := GetEsClient()
	if err != nil {
		log.Println(err)
	}
	do, err := ESclient.Delete().Index(index).Id(doc_id).Do(context.Background())
	if do != nil {
		log.Printf("doc%s删除成功", do.Id)
	}
	return err
}

func UpdateDoc(index string, doc_id string, any interface{}) error {
	ESclient, err := GetEsClient()
	if err != nil {
		log.Println(err)
	}
	_, err = ESclient.Update().Index(index).Id(doc_id).Doc(any).Do(context.Background())
	return err
}

func SelectNewDocidByindex(index string, key string) (string, error) {
	var doc_id string
	ESclient, err := GetEsClient()
	if err != nil {
		return "", err
	}
	searchResult, err := ESclient.Search().
		Index(index).            // 设置索引名称
		Sort(key, false).        // 根据时间戳字段排序，false表示降序
		Size(1).                 // 只获取一条记录
		Do(context.Background()) // 执行查询
	if err != nil {
		return "", err
	}

	fmt.Println(searchResult.Hits.TotalHits)
	for _, hit := range searchResult.Hits.Hits {
		if hit == nil {
			return "", err
		}
		doc_id = hit.Id
	}
	return doc_id, err
}

func SelectNewDocByindex(index string, key string, any interface{}) (json.RawMessage, error) {
	ESclient, err := GetEsClient()
	if err != nil {
		return nil, err
	}
	searchResult, err := ESclient.Search().
		Index(index).            // 设置索引名称
		Sort(key, false).        // 根据时间戳字段排序，false表示降序
		Size(1).                 // 只获取一条记录
		Do(context.Background()) // 执行查询
	if err != nil {
		return nil, err
	}

	for _, hit := range searchResult.Hits.Hits {
		if hit == nil {
			return nil, err
		}
		any = hit.Source
	}
	return any.(json.RawMessage), err
}
func SelectDocidBySome(index string, key string, value string) (string, error) {
	var doc_id string
	ESclient, err := GetEsClient()
	if err != nil {
		log.Println(err)
	}
	//创建复合查询
	boolquery := elastic.NewBoolQuery()
	boolquery.Must(elastic.NewTermsQuery(key, value))

	do, err := ESclient.Search(index).Query(boolquery).Do(context.Background())
	for _, hit := range do.Hits.Hits {
		if hit == nil {
			return "", err
		}
		doc_id = hit.Id
	}
	return doc_id, err
}
