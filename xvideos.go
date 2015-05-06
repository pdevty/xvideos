// This package scrapes the HTML of xvideos.com and gives you information you can use in your go programs
//
// Refer to http://github.com/pdevty/xvideos for more infomation
//
//  package main
//
//  import (
// 		"fmt"
// 		"github.com/pdevty/xvideos"
// 		"log"
//  )
//
//  func main() {
//  	// url to be set xvideos list page
// 		xv, err := xvideos.Get("http://jp.xvideos.com/c/asian_woman-32/")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
//  	for _, v := range xv {
// 			fmt.Println(v.Id, v.Url, v.Duration, v.Rating, v.Thumbnail, v.Title, v.Tags)
// 		}
//  }
package xvideos

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"regexp"
	"runtime"
	"strings"
)

// get func return xvideo data struct
type Xvideo struct {
	Id        string
	Title     string
	Thumbnail string
	Url       string
	Duration  string
	Rating    string
	Tags      []string
}

// get xvideos tags
func getTags(url string) (tags []string, err error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile("/tags/[a-zA-Z]+")
	doc.Find("ul#video-tags li a").Each(func(_ int, s *goquery.Selection) {
		tag := s.Text()
		href, _ := s.Attr("href")
		if re.MatchString(href) {
			tags = append(tags, tag)
		}
	})
	return tags, nil
}

// get xvideos information. url to be set xvideos list page
func Get(rawurl string) (xvideos []Xvideo, err error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	base := u.Scheme + "://" + u.Host
	doc, err := goquery.NewDocument(u.String())
	if err != nil {
		return nil, err
	}
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	re := regexp.MustCompile("[0-9]*%")
	ch := make(chan map[int][]string)
	doc.Find("div.thumbBlock").Each(func(i int, s *goquery.Selection) {
		xv := Xvideo{}
		url, _ := s.Find("p a").Attr("href")
		xv.Url = base + url
		go func() {
			tags, _ := getTags(xv.Url)
			ch <- map[int][]string{i: tags}
		}()
		thumbnail, _ := s.Find("img").Attr("src")
		xv.Thumbnail = thumbnail
		xv.Title = s.Find("p a").Text()
		xv.Duration = s.Find("span.duration").Text()
		xv.Rating = re.FindString(s.Find("span.bg").Text())
		id, _ := s.Attr("id")
		xv.Id = strings.TrimLeft(id, "video_")
		xvideos = append(xvideos, xv)
	})
	for i := 0; i < len(xvideos); i++ {
		m := <-ch
		for k := range m {
			xvideos[k].Tags = m[k]
		}
	}
	return xvideos, nil
}
