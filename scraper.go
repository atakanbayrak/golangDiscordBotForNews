package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

type FeedItem struct {
	Title string
	URL   string
}

func readFile(fname string) string {
	databyte, err := ioutil.ReadFile(fname)

	if err != nil {
		panic(err)
	}
	return string(databyte)
}
func ParseRSS() {
	blogList := [1]string{"https://pwnlab.me/feed/"}
	fp := gofeed.NewParser()
	fp.Client = &http.Client{Timeout: time.Second * 5}

	feed_items := make([]FeedItem, 1)

	for true {
		for k := 0; k < len(blogList); k++ {
			feed, err := fp.ParseURL(blogList[k])

			if err == nil {
				// No error
				l.Printf("[INFO] RSS Parser %s icin basladi", blogList[k])
				items := feed.Items
				for i := 0; i < len(items); i++ {
					if strings.Contains(readFile("feed_item.list"), items[i].Link) {
						l.Printf("[WARN] FeedItem zaten olusturuldu, Link %s", items[i].Link)
					} else {
						// FeedItem olusturulur, kanala gönderilir ve feed_item.list dosyasına link yazılır
						feedItem := FeedItem{Title: items[i].Title, URL: items[i].Link}
						feed_items = append(feed_items, feedItem)
						l.Printf("[INFO] FeedItem olusturuldu. Title %s, URL %s", feedItem.Title, feedItem.URL)
						// feed Item objesinin URL ini feed_item.list dosyasına yazalım (Tekrar göndermemek için)
						file, err := os.OpenFile("feed_item.list", os.O_APPEND|os.O_WRONLY, 0644)
						if err != nil {
							panic(err)
						}
						defer file.Close()
						if _, err := file.WriteString(items[i].Link + "\n"); err != nil {
							l.Fatal(err)
						}

						//DC Kanalına mesaj gönderilir
						msg := "Yeni bir gönderi paylasildi: **" + items[i].Title + "**\n" + items[i].Link
						Dg.ChannelMessageSend(botChID, msg)
					}
				}
			} else {
				l.Printf("[ERR] FeedItems olusturulamadı URL: %s", blogList[k])
			}
			time.Sleep(time.Hour * 8)
		}

	}
}
