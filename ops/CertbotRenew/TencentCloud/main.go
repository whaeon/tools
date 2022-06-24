package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

type DomainInfo struct {
	Request *dnspod.CreateRecordRequest
}

func NewDomainInfo() *DomainInfo {
	return &DomainInfo{
		Request: dnspod.NewCreateRecordRequest(),
	}
}

type ConnectToken struct {
	SecretId  string
	SecretKey string
}

func NewConnectToken(id string, key string) *ConnectToken {
	return &ConnectToken{
		SecretId:  id,
		SecretKey: key,
	}
}

func (c *ConnectToken) Connect(info *DomainInfo) {
	credential := common.NewCredential(c.SecretId, c.SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	client, _ := dnspod.NewClient(credential, "", cpf)

	response, err := client.CreateRecord(info.Request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

func main() {
	conn := NewConnectToken(os.Getenv("SecretID"), os.Getenv("SecretKey"))

	domain := flag.String("D", "", "Domain")
	subDomain := flag.String("d", "", "SubDomain")
	recordType := flag.String("t", "", "RecordType")
	value := flag.String("v", "", "Value")
	recordLine := flag.String("l", "", "RecordLine")
	flag.Parse()

	info := NewDomainInfo()
	info.Request.Domain = domain
	info.Request.SubDomain = subDomain
	info.Request.RecordType = recordType
	info.Request.Value = value
	info.Request.RecordLine = recordLine

	conn.Connect(info)

}
