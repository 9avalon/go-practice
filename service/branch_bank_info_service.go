package service

import (
	"bytes"
	"context"
	"encoding/json"
	"go-practice/resource"
	"go-practice/utils"
	"log"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esutil"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// BranchBankInfo ...
type BranchBankInfo struct {
	ID             int64  `json:"id" db:"id"`
	BranchBankName string `json:"branch_bank_name" db:"branch_bank_name"`
	BankName       string `json:"bank_name"        db:"bank_name"`
	BankId         int64  `json:"bank_id"          db:"bank_id"`
	ProvinceName   string `json:"province_name"    db:"province_name"`
	ProvinceCode   string `json:"province_code"    db:"province_code"`
	CityName       string `json:"city_name"        db:"city_name"`
	CityCode       string `json:"city_code"        db:"city_code"`
}

// mock支行数据
func mockBranchBankTemplateData() []*BranchBankInfo {
	var list []*BranchBankInfo

	b1 := &BranchBankInfo{
		ID:             429064,
		BranchBankName: "南洋商业银行（中国）有限公司北京五路居支行",
		BankName:       "南洋商业银行（中国）",
		BankId:         503,
		ProvinceName:   "北京市",
		ProvinceCode:   "10001",
		CityName:       "北京市",
		CityCode:       "10001",
	}

	b2 := &BranchBankInfo{
		ID:             429110,
		BranchBankName: "恒生银行(中国)有限公司上海分行",
		BankName:       "恒生银行（中国）",
		BankId:         504,
		ProvinceName:   "上海市",
		ProvinceCode:   "10010",
		CityName:       "上海市",
		CityCode:       "10082",
	}

	list = append(list, b1)
	list = append(list, b2)
	return list
}

func EsInfo() {
	res, err := resource.GetEsClient().Client.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)
}

func InsertEsData() {
	data := mockBranchBankTemplateData()

	// 单条单条插入
	for _, d := range data {
		request := esapi.IndexRequest{
			Index:      "idx_branch_bank_info",
			DocumentID: strconv.Itoa(int(d.ID)),
			Body:       bytes.NewReader([]byte(utils.ToJSON(d))),
			Refresh:    "true",
		}

		// 请求
		res, err := request.Do(context.Background(), resource.GetEsClient().Client)
		if err != nil {
			log.Fatalf("index request fail, err=%+v", err)
			return
		}

		if res.IsError() {
			log.Printf("[%s] Error indexing document ID=%d", res.Status(), d.ID)
		} else {
			// Deserialize the response into a map.
			var r map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
				log.Printf("Error parsing the response body: %s", err)
			} else {
				// Print the response status and indexed document version.
				log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
			}
		}

		err = res.Body.Close()
		if err != nil {
			return
		}
	}
}

func InsertBulkData() {
	data := mockBranchBankTemplateData()

	indexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:      "idx_branch_bank_info",
		NumWorkers: 4,
		Client:     resource.GetEsClient().Client,
	})
	if err != nil {
		return
	}

	for _, v := range data {
		err := indexer.Add(context.Background(), esutil.BulkIndexerItem{
			Action:     "index",
			DocumentID: strconv.Itoa(int(v.ID)),
			Body:       strings.NewReader(utils.ToJSON(v)),
		})
		if err != nil {
			return
		}
	}

	err = indexer.Close(context.Background())
	if err != nil {
		return
	}

	biStats := indexer.Stats()
	// Report the results: number of indexed docs, number of errors, duration, indexing rate
	//
	log.Println(biStats.NumAdded)

}

//
//func SearchEsData() {
//	// Build the request body.
//	var buf bytes.Buffer
//	query := map[string]interface{}{
//		"query": map[string]interface{}{
//			"match": map[string]interface{}{
//				"title": "test",
//			},
//		},
//	}
//
//	if err := json.NewEncoder(&buf).Encode(query); err != nil {
//		log.Fatalf("Error encoding query: %s", err)
//	}
//
//	// Perform the search request.
//	res, err = es.Search(
//		es.Search.WithContext(context.Background()),
//		es.Search.WithIndex("test"),
//		es.Search.WithBody(&buf),
//		es.Search.WithTrackTotalHits(true),
//		es.Search.WithPretty(),
//	)
//	if err != nil {
//		log.Fatalf("Error getting response: %s", err)
//	}
//	defer res.Body.Close()
//
//	if res.IsError() {
//		var e map[string]interface{}
//		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
//			log.Fatalf("Error parsing the response body: %s", err)
//		} else {
//			// Print the response status and error information.
//			log.Fatalf("[%s] %s: %s",
//				res.Status(),
//				e["error"].(map[string]interface{})["type"],
//				e["error"].(map[string]interface{})["reason"],
//			)
//		}
//	}
//
//	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
//		log.Fatalf("Error parsing the response body: %s", err)
//	}
//	// Print the response status, number of results, and request duration.
//	log.Printf(
//		"[%s] %d hits; took: %dms",
//		res.Status(),
//		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
//		int(r["took"].(float64)),
//	)
//	// Print the ID and document source for each hit.
//	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
//		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
//	}
//
//	log.Println(strings.Repeat("=", 37))
//}
