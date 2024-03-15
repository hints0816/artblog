package model

import (
	"time"
	"fmt"
	"strconv"
	"github.com/patrickmn/go-cache"
)

var goCahce = cache.New(5*time.Minute, 600*time.Second)

func InitCronTime() {
	//过期时间，清理过期的时间
	ticker := time.NewTicker(time.Millisecond * 10000)
    go func() {
        for t := range ticker.C {
			for e := range goCahce.Items() {
				fmt.Println("flushed",t)
				count, _ := goCahce.Get(e)
				id, _ := strconv.Atoi(e)
				count1, _ := count.(int)
				UpdateArtView(id, count1)
				goCahce.Flush()
			}
		}
    }()
}

func SetViewCache(id string) {
	var found bool
	_, found = goCahce.Get(id)
	if found {
		goCahce.IncrementInt(id, 1)
	} else {
		goCahce.Set(id, 1, cache.DefaultExpiration)
	}
}

func GetGoCahce() *cache.Cache  {
	return goCahce
}