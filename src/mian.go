package main

import (
	 "fmt"
		"github.com/go-redis/redis"
	"strings"
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
	val3 := client.Info("Persistence")

	fmt.Println(val3)

	val3 := map[string]string{}

}
