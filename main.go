package main

import (
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gomodule/redigo/redis"
)

var (
	REDIS_HOST string
	REDIS_PORT string
	REDIS_PASS string
	RedisUrl   string
	REDIS_DB   string
	// err        error
)

/*
redisGet 从redis中获取数据
*/
func redisGet(keys string) (result_str string, err error) {

	conn, err := redis.DialURL(RedisUrl)

	if err != nil {
		return
	}
	defer conn.Close()

	result, err := conn.Do("GET", keys)

	if err != nil {
		return "", err
	}
	if result == nil {
		err = errors.New("redis GET 的数据为空")
		return "", err
	}
	result_str = string(result.([]byte))
	return result_str, nil
}

/*
FormatRedisUrl 格式化 redis的链接
redis_pwd redis 链接密码
redis_host redis 链接地址
redis_port redis 链接端口
redis_db redis 使用的数据库
*/
func formatRedisUrl(redis_pwd, redis_host, redis_port, redis_db string) (format_url string) {

	if redis_pwd == "" || redis_pwd == "nil" {
		format_url = "redis://" + redis_host + ":" + redis_port + "/" + redis_db
	} else {
		format_url = "redis://:" + redis_pwd + "@" + redis_host + ":" + redis_port + "/" + redis_db
	}
	return

}

/*
redisGet 往redis写入数据获取数据
*/
func RedisSet(keys string, values string) (result bool, err error) {

	var Expire = 60
	conn, err := redis.DialURL(RedisUrl)

	if err != nil {
		return
	}
	defer conn.Close()

	_, err = conn.Do("SET", keys, values)

	if err != nil {
		return false, err
	}
	_, err = conn.Do("Expire", keys, Expire)

	if err != nil {
		return false, err
	}

	return true, nil

}

func main() {
	// export REDIS_HOST="124.71.182.117" REDIS_PORT="6379"  REDIS_PASS="WGRnZnVYZW9sRW81VEwwTgo=" REDIS_DB="0"

	var keys = flag.String("k", "", "redis获取的key")

	flag.Parse()
	REDIS_HOST = os.Getenv("REDIS_HOST")
	REDIS_PORT = os.Getenv("REDIS_PORT")
	REDIS_PASS = os.Getenv("REDIS_PASS")
	REDIS_DB = os.Getenv("REDIS_DB")

	sDec, _ := base64.StdEncoding.DecodeString(REDIS_PASS)

	RedisUrl = formatRedisUrl(strings.Replace(string(sDec), "\n", "", -1), REDIS_HOST, REDIS_PORT, REDIS_DB)

	result, err := redisGet(*keys)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

}
