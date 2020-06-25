package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
	"strconv"
)

func main() {

	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{"clustercfg.awrxipp5sxfb6qy.gguve8.usw2.cache.amazonaws.com:6379"},
		Password: "ipmvwwb1kapi^r^z#tMja^w#ue^kf!ox",
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
        j := 100
	for i := 0; i < j; i++ {
		key := "nikkey"+strconv.Itoa(i)
	err = client.Set(ctx, key, strconv.Itoa(i), time.Minute).Err()
        if err != nil {
	      panic(err)
	}

//	res, _ := client.ClusterInfo(ctx).Result()
//	fmt.Println(res)
	hashSlot, _ := client.ClusterKeySlot(ctx,key).Result()
	fmt.Printf("Key %v is in slot %v \n" ,key, hashSlot)

     //  keys, err := client.ClusterGetKeysInSlot(hashtag.Slot("key"+strconv.Itoa(i)), 1).Result()
      // fmt.Println(keys)

}


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
