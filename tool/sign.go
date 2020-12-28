package tool

import (
	"fmt"
	"sort"
	"strings"
	"github.com/davecgh/go-spew/spew"
)

type SignParams struct {
	m    map[string]interface{} // 参数
	apiKey string                 // apiKey

}

func (p *SignParams) InitSign(mData map[string]interface{}, apiKey, sign string) (bool) {
	p.m = mData
	p.apiKey = apiKey


	var keys []string
	for k, _ := range p.m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	builder := strings.Builder{}
	for _, v := range keys {
		builder.WriteString(v)
		builder.WriteString("=")
		builder.WriteString(fmt.Sprint(p.m[v]))
		builder.WriteString("&")
	}
	builder.WriteString("apikey=" + p.apiKey)
	signData := strings.ToUpper(Md5(builder.String()))
	spew.Dump(signData)
	if signData == sign{
		return true
	}
	return false
}

func (p *SignParams) InitSignVosSign(mData map[string]interface{}, apiKey, sign string) (bool) {
	p.m = mData
	p.apiKey = apiKey


	var keys []string
	for k, _ := range p.m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	builder := strings.Builder{}
	builder.WriteString( p.apiKey)
	for _, v := range keys {
		builder.WriteString(v)
		builder.WriteString(fmt.Sprint(p.m[v]))
	}
	spew.Dump(builder.String())
	signData := Md5(builder.String())
	spew.Dump(mData,"InitSignVosSign",signData)
	if signData == sign{
		return true
	}
	return false
}
