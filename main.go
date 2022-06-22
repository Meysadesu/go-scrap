package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	// @ Visited to link in href
	// @ Done
	c.OnHTML("div.entry-content", func(h *colly.HTMLElement) {
		h.ForEach("p", func(i int, h *colly.HTMLElement) {
			text := h.ChildAttr("a", "href")
			if text == "" {
				return
			}

			if strings.Contains(text, "holypantsu.wordpress.com") {
				h.Request.Visit(text)
				fmt.Println("visited to -> : ", text, " : -> scrapping success")
			}
		})
	})

	// @ scrapping content website novels
	c.OnHTML("div.entry-content", func(h *colly.HTMLElement) {
		h.ForEach("p", func(i int, h *colly.HTMLElement) {
			text := h.ChildText("strong")
			fmt.Println(text)
		})
	})

	c.Visit("https://holypantsu.wordpress.com/kono-subarashii-sekai-ni-shukufuku-wo/")
}
