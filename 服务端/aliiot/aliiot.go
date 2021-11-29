// This file is auto-generated, don't edit it. Thanks.
package aliiot

import (
	"encoding/json"
	"fmt"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	iot20180120 "github.com/alibabacloud-go/iot-20180120/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

var client *iot20180120.Client
var _err error
var AllDevices []*iot20180120.QueryDeviceResponseBodyDataDeviceInfo

type data struct {
	Data int `json:"data"`
}
type Indentifier struct {
	Indentifier string `json:"Indentifier"`
}

type Value struct {
	Value data `json:"Value"`
}
type Times struct {
	Time int64 `json:"Time"`
}
type info struct {
	Indentifier string `json:"Indentifier"`
	Value       data   `json:"Value"`
	Time        int64  `json:"Time"`
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *iot20180120.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("iot.cn-shanghai.aliyuncs.com")
	_result = &iot20180120.Client{}
	_result, _err = iot20180120.NewClient(config)
	return _result, _err
}
func Ali_init() {
	//首先获取产品类目下所有设备,在暂时只负责一个产品
	client, _err = CreateClient(tea.String("LTAI5tE3tZqxQbjj4BhsvfDV"), tea.String("nworMiBq9GZdZvKj0LZJ29GwOlQnkm"))
	_getAllDevices("a1J4LfeKtKw")
}
func Smart_setDeviceProperty(name string, key string, k string, v string) error {
	fmt.Println("{\"" + k + "\":" + v + "}")
	setDevicePropertyRequest := &iot20180120.SetDevicePropertyRequest{
		DeviceName: tea.String(name),
		ProductKey: tea.String(key),
		Items:      tea.String("{\"" + k + "\":" + v + "}"),
	}
	_, _err = client.SetDeviceProperty(setDevicePropertyRequest)
	if _err != nil {
		return _err
	}
	return _err
}
func Getproperty(name string, key string) (results []info) {
	queryDeviceOriginalPropertyStatusRequest := &iot20180120.QueryDeviceOriginalPropertyStatusRequest{
		Asc:        tea.Int32(1),
		PageSize:   tea.Int32(10),
		ProductKey: tea.String(key),
		DeviceName: tea.String(name)}
	res, _ := client.QueryDeviceOriginalPropertyStatus(queryDeviceOriginalPropertyStatusRequest)
	for _, v := range res.Body.Data.List.PropertyStatusDataInfo {
		var temp info
		temp.Time = *v.Time
		temp.Indentifier = *v.Identifier
		json.Unmarshal([]byte(*v.Value), &temp.Value)
		results = append(results, temp)
	}
	return results
}
func Reflash() {
	_getAllDevices("a1J4LfeKtKw")
}
func _getAllDevices(key string) {
	_, res := getAllDevices(key)
	AllDevices = res.Body.Data.DeviceInfo
}
func getAllDevices(key string) (error, *iot20180120.QueryDeviceResponse) {
	queryDeviceRequest := &iot20180120.QueryDeviceRequest{
		ProductKey: tea.String(key),
	}
	res, _err := client.QueryDevice(queryDeviceRequest)
	if _err != nil {
		return _err, nil
	}
	return _err, res
}
