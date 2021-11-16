package elastic_plus

import (
	"log"
	"testing"
)

func TestListIndices(t *testing.T) {
	
	cfg := NewEsConfig("http://localhost:9200", "test", "test")
	
	newClient, err := NewClient(cfg)
	if err != nil {
		return
	}
	
	index, err := newClient.SearchIndex()
	if err != nil {
		return
	}
	
	log.Println(index)
	
}
