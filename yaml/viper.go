package yaml

import (
	"log"
	
	"github.com/spf13/viper"
)

var (
	ConfigName string
	SetFile    *viper.Viper
)

// SetConfigName 读取配置文件
func SetConfigName(in string) { SetFile.SetConfigName(in) }

func init() {
	SetFile.SetConfigName(ConfigName)
	// viper.SetConfigName("config")     // 读取json配置文件
	//viper.AddConfigPath("/etc/appname/")   //设置配置文件的搜索目录
	//viper.AddConfigPath("$HOME/.appname")  // 设置配置文件的搜索目录
	SetFile.AddConfigPath(".") // 设置配置文件和可执行二进制文件在用一个目录
	if err := SetFile.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
		log.Fatal(err) // 读取配置文件失败致命错误
	}
}

// func main() {
// 	fmt.Println("获取配置文件的port", viper.GetInt("port"))
// 	fmt.Println("获取配置文件的mysql.url", viper.GetString(`mysql.url`))
// 	fmt.Println("获取配置文件的mysql.username", viper.GetString(`mysql.username`))
// 	fmt.Println("获取配置文件的mysql.password", viper.GetString(`mysql.password`))
// 	fmt.Println("获取配置文件的redis", viper.GetStringSlice("redis"))
// 	fmt.Println("获取配置文件的smtp", viper.GetStringMap("smtp"))
// }
