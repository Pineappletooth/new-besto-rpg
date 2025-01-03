package persistence

import (
	"strconv"
	"time"
)

func AddCommandLastUsed(userId uint32, command string)error{
	res := redisClient.HSet(ctx, "commandLastUsed:"+strconv.FormatUint(uint64(userId),10), command, time.Now().Unix())
	return res.Err()
}

func GetCommandLastUsed(userId uint32, command string) (time int64, err error){
	if(redisClient.Exists(ctx, "commandLastUsed:"+strconv.FormatUint(uint64(userId),10)).Val() == 0){
		return 0, nil
	}
	res := redisClient.HGet(ctx, "commandLastUsed:"+strconv.FormatUint(uint64(userId),10), command)
	return res.Int64()
}