package main

import (
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"github.com/hewenyu/toolspackage/global_variable"
	"os"
	"strings"
	
	"github.com/gomodule/redigo/redis"
)

var (
	RedisHost string
	RedisPort string
	RedisPass string
	RedisUrl  string
	RedisDb   string
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
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			global_variable.CLOG.Error(err.Error())
		}
	}(conn)
	
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
func formatRedisUrl(redisPwd, redisHost, redisPort, redisDb string) (formatUrl string) {
	
	if redisPwd == "" || redisPwd == "nil" {
		formatUrl = "redis://" + redisHost + ":" + redisPort + "/" + redisDb
	} else {
		formatUrl = "redis://:" + redisPwd + "@" + redisHost + ":" + redisPort + "/" + redisDb
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
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			global_variable.CLOG.Error(err.Error())
		}
	}(conn)
	
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
	RedisHost = os.Getenv("REDIS_HOST")
	RedisPort = os.Getenv("REDIS_PORT")
	RedisPass = os.Getenv("REDIS_PASS")
	RedisDb = os.Getenv("REDIS_DB")
	
	sDec, _ := base64.StdEncoding.DecodeString(RedisPass)
	
	RedisUrl = formatRedisUrl(strings.Replace(string(sDec), "\n", "", -1), RedisHost, RedisPort, RedisDb)
	
	result, err := redisGet(*keys)
	
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
	
}
