package redishandler

import (
	"fmt"
	"log"
	"math/rand"
	"path"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var ServerPath string = "niqurl-server:8081"

func PrintShortURL(url string) error {
	x, _ := Client.ZScan(Ctx, "longurl", 0, url, 0).Val()
	i, err := strconv.Atoi(x[1])
	if err != nil {
		return err
	}
	shorturl := Client.ZRange(Ctx, "shorturl", int64(i), int64(i)).Val()
	fmt.Println("URL [" + url + "] shortened before to: " + path.Join(ServerPath, shorturl[0]))
	return nil
}

func InsertURL(url, shorturl, user string) {
	fmt.Println("Creating short URL for [" + url + "]: " + path.Join(ServerPath, shorturl))
	wt, _ := getSetting("USER_WAIT_TIME")
	uc, _ := getSetting("URL_COUNT")
	Client.Incr(Ctx, "URL_COUNT")
	Client.ZAdd(Ctx, "longurl", &redis.Z{Score: float64(uc), Member: url})
	Client.ZAdd(Ctx, "shorturl", &redis.Z{Score: float64(uc), Member: shorturl})
	Client.RPush(Ctx, "createdby", user)
	Client.Set(Ctx, user, true, time.Duration(wt*1000000000))
}

func ShortURL(url string) string {

	var shrt string
	n, _ := getSetting("SHORT_URL_LEN")

	for i := 0; true; i++ {
		if i == 100 {
			log.Printf("Can't find available url with lenght of : %v\n", n)
			log.Println("Increase url length by setlen command")
			break
		}
		shrt = shortURLGenerate(n)
		if CheckZSet(shrt, "shorturl") == true {
			continue
		}
		break
	}
	return shrt
}

// returns random
func shortURLGenerate(n int) string {
	chr := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	u := make([]byte, n)
	for i := range u {
		rand.Seed(time.Now().UTC().UnixNano())
		u[i] = chr[rand.Intn(len(chr))]
	}

	return string(u)
}
