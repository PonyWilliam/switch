// This file is auto-generated, don't edit it. Thanks.
package message

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func SendMsg(phone string, msg string) (_err error) {
	client, _err := CreateClient(tea.String("LTAI5tE3tZqxQbjj4BhsvfDV"), tea.String("nworMiBq9GZdZvKj0LZJ29GwOlQnkm"))
	if _err != nil {
		return _err
	}
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("小伟王"),
		TemplateCode:  tea.String("SMS_171859237"),
		PhoneNumbers:  tea.String(phone),
		TemplateParam: tea.String("{\"code\":" + msg + "}"),
	}
	// 复制代码运行请自行打印 API 的返回值
	_, _err = client.SendSms(sendSmsRequest)
	if _err != nil {
		return _err
	}
	return _err
}
