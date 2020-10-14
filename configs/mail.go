// Package configs provides ...
package configs

func GetMailConfig() map[string]interface{} {
	mailConfig := make(map[string]interface{})
	//host string, port int, username, password string
	mailConfig["host"] = "smtp.163.com"
	mailConfig["port"] = "25"
	mailConfig["username"] = ""
	mailConfig["password"] = ""

	return mailConfig
}
