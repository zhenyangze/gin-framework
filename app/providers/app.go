package providers

var (
	appPath    string
	configPath string
)

func init() {
}

func SetAppPath(path string) {
	appPath = path
}

func GetAppPath() string {
	return appPath
}

func SetConfigPath(path string) {
	configPath = path
}

func GetConfigPath() string {
	return configPath
}
