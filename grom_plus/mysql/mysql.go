package mysql

// MYSQL 默认的配置文件
type MYSQL struct {
	Name        string `json:"name"`         // 数据库名称
	Password    string `json:"password"`     // 数据库密码
	User        string `json:"user"`         // 数据库用户
	Host        string `json:"host"`         // host
	Port        string `json:"port"`         // 端口
	TablePrefix string `json:"table_prefix"` // 表前缀
	ParseTime   string `json:"parseTime"`    // 时间修改
	Loc         string `json:"loc"`          // default Local
}
