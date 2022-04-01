// Package configs provides ...
package configs

import "gitee.com/zhenyangze/gin-framework/helpers"

func GetMailConfig() map[string]interface{} {
	mailConfig := make(map[string]interface{})

	var config *helpers.Config
	//host string, port int, username, password string
	mailConfig["host"] = config.GetString("mail.host")
	mailConfig["port"] = config.GetString("mail.port")
	mailConfig["username"] = config.GetString("mail.username")
	mailConfig["password"] = config.GetString("mail.password")

	return mailConfig
}
