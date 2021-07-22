package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	// string type
	user := User{
		"mohammad",
		21,
	}
	j, _ := json.Marshal(user)

	result, _ := rdb.Set(ctx, "person", j, 0).Result()
	result, _ = rdb.Set(ctx, "name", "mohammad", 0).Result()
	_ = result
	result, _ = rdb.Get(ctx, "person").Result()
	result, _ = rdb.Get(ctx, "name").Result()

	result, _ = rdb.MSet(ctx, "user1", "ali", "user2", "mohammad").Result()
	users, _ := rdb.MGet(ctx, "user1", "user2").Result()
	_ = users

	// list type
	rdb.RPush(ctx, "users", "mohammad", "ali")
	rdb.LPush(ctx, "users", "erfan")
	userList, _ := rdb.LRange(ctx, "users", 0, -1).Result()
	_ = userList
	result, _ = rdb.LPop(ctx, "users").Result()

	// hashes
	rdb.HSet(ctx, "user:1", "username", "mohammad")
	rdb.HMSet(ctx, "user:1", "email", "m.dehghanpour10@gmail.com", "age", "21")
	rdb.HIncrBy(ctx, "user:1", "age", 10)
	result, _ = rdb.HGet(ctx, "user:1", "username").Result()
	userTable, _ := rdb.HGetAll(ctx, "user:1").Result()
	_ = userTable

	//sets
	rdb.SAdd(ctx, "numberSet", 1, 2, 3, 4)
	setMember, _ := rdb.SMembers(ctx, "numberSet").Result()
	isMebmer, _ := rdb.SIsMember(ctx, "numberSet", 5).Result()
	_, _ = setMember, isMebmer

	//sorted set

	rdb.ZAdd(ctx, "sorted", &redis.Z{Score: 5, Member: "mohammad"}, &redis.Z{Score: 1, Member: "reza"}, &redis.Z{Score: 10, Member: "ali"})

	zs, _ := rdb.ZRangeWithScores(ctx, "sorted", 0, -1).Result()

	fmt.Println(zs)
	// general
	rdb.Expire(ctx, "name", 10*time.Second)
	exists, _ := rdb.Exists(ctx, "user3").Result()
	_ = exists
	keys, _ := rdb.Keys(ctx, "*").Result() //get all key
	_ = keys
	del, _ := rdb.Del(ctx, "name").Result() //delete key
	_ = del
	rdb.FlushAll(ctx)                //clear redis
	rdb.Rename(ctx, "name", "title") //rename key
}
