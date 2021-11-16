package bluk

import "strings"

// ToString string 格式插入
func (b *BulkData) ToString() string {
	var body strings.Builder
	body.Reset()
	
	body.WriteString(`{"index" : { "_index" : "` + b.Index + `", "_type" : "_doc", "_id" : "` + b.Data.ID + `" }}`)
	body.WriteString("\n")
	
	jsonBody, _ := b.Data.Body.GetJson()
	body.Write(jsonBody)
	
	// 这个需要换行作为终止符
	body.WriteString("\n")
	
	return body.String()
}

// ToString string 格式插入
func (b *BulkDataList) ToString() string {
	var body strings.Builder
	body.Reset()
	
	for _, v := range *b {
		body.WriteString(v.ToString())
	}
	
	return body.String()
}
