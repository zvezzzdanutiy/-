package AnekdotProviders

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (d *Domain) GetAnekdot(ctx context.Context, category string) (string, error) {
	response, err := d.client.Get(d.GenerateURL())
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("wrong response got, code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), err
}
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
func (d *Domain) GenerateURL() string {
	const apiKey = "7c72cdbe4f22a675723cfccb9669cbe31d7c1aed2abf8ba146f7fa24426b920a"
	x := fmt.Sprintf("%d", time.Now().Unix())
	str := "pid=s157zmtnncbm4m15b40k&method=getRandItem&uts=" + x // Установите значение num равным желаемому количеству анекдотов
	hash := GetMD5Hash(str + apiKey)
	return "http://anecdotica.ru/api?" + str + "&hash=" + hash
}
