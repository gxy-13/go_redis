package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化链接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password
		DB:       0,  // default db
		PoolSize: 100,
	})

	_, err = rdb.Ping().Result()
	return err
}

func main() {
	err := initClient()
	if err != nil {
		fmt.Printf("init redis client failed,err:%v\n", err)
		return
	}
	fmt.Println("connect redis success....")
	defer rdb.Close()

}
