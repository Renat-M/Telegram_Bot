package telegram

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func GetAnekdot() string {
	resp, err := http.Get("https://www.anekdot.ru/random/anekdot/")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	re := regexp.MustCompile(`<div class="text">(.+?)</div>`)
	match := re.FindStringSubmatch(string(body))
	if len(match) > 1 {
		anek := match[1]
		anek = strings.ReplaceAll(anek, "<br>", "\n")
		return anek
	} else {
		return "Анекдот не найден"
	}
}
