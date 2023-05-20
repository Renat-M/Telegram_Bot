package telegram

import (
	"github.com/go-rod/rod"
)

func GetMeme() string {

	//making new browser
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	//making new page
	page := browser.MustPage()

	//go to page
	err := page.Navigate("https://admem.ru/rndm")
	if err != nil {
		return "Нет мема :("
	}

	//looking for img
	imgElement := page.MustElement("img")
	imgURL := imgElement.MustProperty("src").String()

	return imgURL
}
