package persistence

import (
	"time"
)

func AddCommandLastUsed(userId string, command string) error {
	res := redisClient.HSet(ctx, "commandLastUsed:"+userId, command, time.Now().Unix())
	return res.Err()
}

func GetCommandLastUsed(userId string, command string) (time int64, err error) {
	if redisClient.Exists(ctx, "commandLastUsed:"+userId).Val() == 0 {
		return 0, nil
	}
	res := redisClient.HGet(ctx, "commandLastUsed:"+userId, command)
	return res.Int64()
}
