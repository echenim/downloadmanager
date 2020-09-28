package enginee

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func (d Downloader) getHTTPRequest(method string) (*http.Request, error) {
	r, er := http.NewRequest(method, d.From, nil)
	check(er, "error occured while retriving the file:")
	r.Header.Set("User-Agent", "download manager")

	return r, nil
}

func (d Downloader) act() (string, error) {
	fmt.Println("\n making connection to ......")
	r, er := d.getHTTPRequest("HEAD")
	check(er, "fail to open connection : ")

	resp, er := http.DefaultClient.Do(r)
	check(er, "fail to get header : ")
	log.Printf("\n Response : %v\n ", resp.StatusCode)

	if resp.StatusCode > 299 {
		log.Printf("Can not process, response is %v\n", resp.StatusCode)
	}
	size, er := strconv.Atoi(resp.Header.Get("Content-Length"))
	check(er, "fail to get content-length")
	log.Printf("Size is %v bytes\n", size)
	var sections = make([][2]int, d.Section)
	var eachSize = size / d.Section
	sections = d.FormSections(sections, eachSize)

	log.Println(sections)
	var wg sync.WaitGroup
	// download each section concurrently
	for i, s := range sections {
		wg.Add(1)
		go func(i int, s [2]int) {
			defer wg.Done()
			_,err = d.downloadSection(i, s)
			if err != nil {
				panic(err)
			}
		}(i, s)
	}
	wg.Wait()


	return "", nil
}
