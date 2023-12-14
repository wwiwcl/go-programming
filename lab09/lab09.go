package main

import (
    "flag"
    "fmt"
    "github.com/gocolly/colly"
	"os"
)

func main() {
    max := flag.Int("max", 10, "Max number of comments to show")
    flag.Parse()

	if len(flag.Args()) > 0 {
        fmt.Fprintf(os.Stderr, "flag provided but not defined: -%s\n", flag.Args()[0])
        os.Exit(2)
    }

    c := colly.NewCollector()

    num := 0

    c.OnHTML(".push", func(e *colly.HTMLElement) {
        if num >= *max {
            return
        }
        name := e.ChildText(".push-userid")
        comment := e.ChildText(".push-content")
        time := e.ChildText(".push-ipdatetime")
        fmt.Printf("%d. 名字：%s，留言%s，時間： %s\n", num+1, name, comment, time)
        num++
    })

    c.Visit("https://www.ptt.cc/bbs/joke/M.1481217639.A.4DF.html")
}