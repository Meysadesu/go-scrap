package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Meysadesu/go-scrap/config/database"
	"github.com/gocolly/colly/v2"
)

type Chapter struct {
	ID         int
	ID_Novels  int
	Header     string
	Ctx        string
	Created_at time.Time
}

func main() {
	c := colly.NewCollector()

	// @ Visited to link in href
	// @ Done
	var chapter string
	c.OnHTML("div.entry-content", func(h *colly.HTMLElement) {
		h.ForEach("p", func(i int, h *colly.HTMLElement) {
			text := h.ChildAttr("a", "href")
			if text == "" {
				return
			}

			chap := h.ChildText("a")
			if strings.Contains(chap, "Chapter") {
				removeChapter := strings.ReplaceAll(chap, "Chapter ", "")
				chapter = removeChapter
			}

			if strings.Contains(text, "holypantsu.wordpress.com") {
				h.Request.Visit(text)
				fmt.Println("visited to -> : ", text)
			}
		})
	})

	// @ scrapping content website novels
	c.OnHTML("div.entry-content", func(h *colly.HTMLElement) {

		// @ ambil semua novelnya
		// @ convert array to string
		text := h.ChildTexts("p strong")
		convArrtoString := strings.Join(text, "<enter>")

		// struct
		novels := Chapter{
			ID_Novels:  1,
			Header:     chapter,
			Ctx:        convArrtoString,
			Created_at: time.Now(),
		}

		db, _ := database.DBConnect()
		db.Create(&novels)
	})

	c.Visit("https://holypantsu.wordpress.com/kono-subarashii-sekai-ni-shukufuku-wo/")
}
