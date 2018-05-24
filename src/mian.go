package main

import (
	 "fmt"
		"github.com/go-redis/redis"
	"strings"
	"strconv"
)
var client *redis.Client

func main()  {
	client = redis.NewClient(&redis.Options{
		Addr: "192.168.0.147:6379",
		Password: "",
		DB:	0,
	})
	pong , err := client.Ping().Result()
	fmt.Println(pong, err)
	ExampleClient()
}
func ExampleClient() {
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	val2, err := client.Get("key2").Result()
	if err == redis.Nil{
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("Key2", val2)
	}
	val3, err := client.Info("Persistence").Result()
	rows := strings.Split(val3, "\n")
	rows = append(rows[:0], rows[1:17]... )

	fmt.Println(val3)

	m := make(map[string]int)
	for i := 0; i < len(rows); i++ {
		col := strings.Split(rows[i], ":")
		//fmt.Println(col[0])
		//fmt.Println(col[1])
		b, _ := strconv.Atoi(col[1])
		m[col[0]] = b
		fmt.Println(m[col[0]])
	}
	if m["aof_enabled"] == 0 {


	}

}
