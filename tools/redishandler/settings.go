package redishandler

import (
	"errors"
	"log"
)

func getSetting(setting string) (int, error) {

	s, err := Client.Get(Ctx, setting).Result()
	if err != nil {
		return 0, err
	}
	if len(s) == 0 {
		return 0, errors.New("Value " + setting + " not found")
	}
	v, _ := Client.Get(Ctx, setting).Int()
	return v, nil
}

func CheckSettings() {
	checkSetting("SHORT_URL_LEN", 4)
	checkSetting("USER_WAIT_TIME", 30)
	checkSetting("URL_COUNT", 0)
	checkSetting("USER_COUNT", 0)
}

func checkSetting(setting string, def int) {
	_, err := getSetting(setting)
	if err != nil {
		Client.Set(Ctx, setting, def, 0)
		log.Printf("init var %q not found", setting)
		log.Printf("initializing default value %v", def)

		if setting == "USER_COUNT" {
			log.Println("USER_COUNT = 0, " +
				"initializing database key and " +
				"generating 5 random users" +
				"(see -generate-fake-users flag)")
			GenerateFakeUsers(5)
		}
	}
}
