package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func main() {
	fmt.Println("调用redis")
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if (err != nil) {
		panic(err)
	}

	// 设置字符串
	val, _ := client.Set("name", "chukui", 10000*time.Millisecond).Result()
	fmt.Println("设置字符串返回：", val)

	// 获取字符串
	val, _ = client.Get("name").Result()
	fmt.Println("获取字符串返回", val)

	// 整数字符串
	age := 18
	client.Set("age", age, 0)
	ret, _ := client.Incr("age").Result()
	fmt.Println("年龄增长的一岁", ret)

	// list
	listKey := "name:chukui:list:todo"
	client.RPush(listKey, "play basketball", "swimming")
	list, _ := client.LRange(listKey, 0, -1).Result()
	fmt.Println("todo list:", list)
}
