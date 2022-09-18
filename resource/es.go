package resource

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

var esClient *EsClient
var esClientOnce sync.Once

type EsClient struct {
	Client *elasticsearch.Client
}

//InitEsClient ...
func InitEsClient() error {
	cli, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"https://localhost:9200"},
		Username:  "elastic",
		Password:  "uR-ky0GwoszdkjSANt+9",
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion:         tls.VersionTLS12,
				InsecureSkipVerify: true, //添加这一行跳过验证
			},
		},
	})
	if err != nil {
		return err
	}

	esClient = &EsClient{Client: cli}
	return nil
}

// GetEsClient 获取es实例
func GetEsClient() *EsClient {
	esClientOnce.Do(func() {
		err := InitEsClient()
		if err != nil {
			log.Fatalf("init es fail, err=%+v", err)
			return
		}
	})

	return esClient
}
