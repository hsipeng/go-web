package services

import (
	"context"
	"encoding/json"
	"fmt"

	database "lirawx.cn/go-web/db"
	"lirawx.cn/go-web/models"
)

// GetATodoService 获取todoList，设置缓存
func GetATodoService(id string) (todo *models.Todo, err error) {
	ctx := context.Background()

	val, err := database.Rdb.Get(ctx, "id_"+id).Result()
	json.Unmarshal([]byte(val), &todo)
	if err != nil {
		fmt.Printf("init redis failed, err:%v\n, read from db ----- ", err)
		todo, err = models.GetATodo(id)
		if err != nil {
			fmt.Printf("db find todo failed, err:%v\n,", err)
			return nil, err
		}
		// 存入redis
		// 序列化
		bytes, err := json.Marshal(todo)
		err = database.Rdb.Set(ctx, "id_"+id, string(bytes), 0).Err()
		if err != nil {
			fmt.Printf("save todo in redis failed, err:%v\n,", err)
			return nil, err
		}
		return todo, nil
	}
	fmt.Println("read from redis", todo)
	return
}
