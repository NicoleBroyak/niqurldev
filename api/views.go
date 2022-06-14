package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/nicolebroyak/niqurldev/tools/redishandler"

	"github.com/gin-gonic/gin"
)

func RedirectURL(c *gin.Context) {
	url, err := FindShortURL(c.Param("url"))
	if err != nil {
		notFound(c)
		return
	}
	c.Redirect(http.StatusMovedPermanently, url)
}

func InspectURL(c *gin.Context) {
	x, b := FindShortURLInfo(c.Param("url"))
	if b != nil {
		notFound(c)
		return
	}
	c.HTML(http.StatusOK, "inspectURL.html", x)
}

func NotFound(c *gin.Context) {
	c.HTML(404, "404.html", "")
}

func FindShortURLInfo(url string) (map[string]interface{}, error) {
	z := map[string]interface{}{}

	scanVal, _ := redishandler.Client.ZScan(
		redishandler.Ctx,
		"shorturl",
		0,
		url,
		0,
	).Val()

	// check if value exists
	if len(scanVal) > 0 {
		i, err := strconv.Atoi(scanVal[1])
		if err != nil {
			return z, err
		}
		z["shorturl"] = url

		z["longurl"] = redishandler.Client.ZRange(
			redishandler.Ctx,
			"longurl",
			int64(i),
			int64(i),
		).Val()[0]

		z["user"] = redishandler.Client.ZRange(
			redishandler.Ctx,
			"username",
			int64(i),
			int64(i),
		).Val()[0]

		return z, nil
	}
	return z, errors.New("shorturl not found")
}

func FindShortURL(url string) (string, error) {
	x, _ := redishandler.Client.ZScan(redishandler.Ctx, "shorturl", 0, url, 0).Val()

	// check if value exists
	if len(x) > 0 {
		i, err := strconv.Atoi(x[1])
		if err != nil {
			return "", err
		}
		url = redishandler.Client.ZRange(
			redishandler.Ctx,
			"longurl",
			int64(i),
			int64(i),
		).Val()[0]

		return url, nil
	}
	return "", errors.New("shorturl not found")
}
