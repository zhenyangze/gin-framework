// Package providers provides ...
package providers

import (
	"gitee.com/zhenyangze/gin-framework/configs"
	"gopkg.in/gomail.v2"
)

var (
//Mail *gomail.Message
)

func init() {
	//Mail = gomail.NewMessage()
}

/*ret, err := providers.SendMail("zhenyangze@163.com", []string{
	"zhenyangze@163.com",
}, []string{"2597762726@qq.com", "阳泽"}, "测试", "<p>测试</p>", "png")

if err != nil {
	fmt.Println(err.Error())
}*/

/**
* SendMail
*
* @param string
* @param []string
* @param []string
* @param string
* @param string
* @param
* @param error
*
* @return
 */
func SendMail(from string, to []string, cc []string, subject string, body string, attach string) (bool, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)

	m.SetAddressHeader("Cc", cc[0], cc[1])
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	m.Attach(attach)

	mailConfig := configs.GetMailConfig()
	d := gomail.NewDialer(mailConfig["host"].(string), mailConfig["port"].(int), mailConfig["username"].(string), mailConfig["password"].(string))

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return false, err
	}

	return true, nil
}
