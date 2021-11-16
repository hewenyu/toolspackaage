package elastic_plus

// EsConfig ES 默认配置
type EsConfig struct {
	Addresses []string `json:"addresses"` // ES的API节点
	Username  string   `json:"username"`  // 账号
	Password  string   `json:"password"`  // 密码
}

// NewEsConfig 初始化
func NewEsConfig(address string, username string, password string) *EsConfig {
	return &EsConfig{
		Addresses: []string{address},
		Username:  username,
		Password:  password,
	}
}
