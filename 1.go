package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {

	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{"awr1819fghzlemyw-0001-001.awr1819fghzlemyw.coupte.usw2.cache.amazonaws.com:6379"},
		Password: "Tpbys6pg!wwsrrfpviypakgvgcnjbbe#",
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	})

	ctx := context.TODO()
	pong, err := client.Ping(ctx).Result()

	if err != nil {
		fmt.Println("Redis ping failed", err)
	}
	fmt.Println(pong)
	res, _ := client.ClusterInfo(ctx).Result()
	fmt.Println(res)
	//hashSlot, _ := client.ClusterKeySlot(ctx,key).Result()
	//fmt.Printf("Key %v is in slot %v" ,key, hashSlot)

     //  keys, err := client.ClusterGetKeysInSlot(hashtag.Slot("key"+strconv.Itoa(i)), 1).Result()
      // fmt.Println(keys)




subPattern := "__key*__:*"
sub := client.PSubscribe(ctx,subPattern)
fmt.Println("watching %v", subPattern)

ctx1 := context.TODO()
for {
  msg, err := sub.Receive(ctx1)
  if err != nil {
    fmt.Println(err)
    continue
  }

  fmt.Println(msg)
}
}
