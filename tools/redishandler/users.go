package redishandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

func RandomUser() string {
	rand.Seed(time.Now().UTC().UnixNano())
	uc, _ := getSetting("USER_COUNT")
	n := int64(rand.Intn(uc))
	un, _ := Client.ZRange(Ctx, "username", n, n).Result()
	return un[0]
}

func GenerateFakeUsers(num int) error {
	u := UsersStruct{}
	url := fmt.Sprintf("https://randomuser.me/api/?results=%v&inc=login,name,email,registered", num+1)
	err := gfuFillStruct(url, &u)
	if err != nil {
		log.Print("Error related with random user API. Try again later")
		return err
	}

	log.Print("Generating random users...")
	failed := 0
	for i := 0; i < num; i++ {
		if CheckZSet(u.Results[i].Login.Username, "username") == true {
			failed++
			continue
		}
		insertUser(i, &u)
	}
	log.Printf("%v random users successfully generated", num-failed)
	return nil
}

func gfuFillStruct(url string, Users *UsersStruct) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &Users)
	if err != nil {
		return err
	}
	if len(Users.Results) == 0 {
		err = errors.New("Error with getting users")
	}
	return nil
}

func insertUser(i int, u *UsersStruct) {
	id, _ := getSetting("USER_COUNT")
	Client.Incr(Ctx, "USER_COUNT")
	Client.ZAdd(Ctx, "username", &redis.Z{
		Score:  float64(id),
		Member: u.Results[i].Login.Username,
	})
	Client.RPush(Ctx, "firstname", u.Results[i].Name.First)
	Client.RPush(Ctx, "lastname", u.Results[i].Name.Last)
	Client.ZAdd(Ctx, "email", &redis.Z{
		Score:  float64(id),
		Member: u.Results[i].Email,
	})
	Client.RPush(Ctx, "regdate", u.Results[i].Registered.Date)
}

type UsersStruct struct {
	Results []struct {
		Name struct {
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Email string `json:"email"`
		Login struct {
			Username string `json:"username"`
		} `json:"login"`
		Registered struct {
			Date time.Time `json:"date"`
		} `json:"registered"`
	} `json:"results"`
}
