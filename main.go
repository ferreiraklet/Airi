package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

func init() {
	flag.Usage = func() {
		help := []string{
			"Airi Hidden Params Finder",
			"",
			"Usage:",
			"+=======================================================+",
			"       -o,     --output     Save output to a file",
			"       -h                   Show This Help Message",
			"",
			"+=======================================================+",
			"",
		}

		fmt.Fprintf(os.Stderr, strings.Join(help, "\n"))
	}

}

func main() {

	var output string
	flag.StringVar(&output, "o", "", "")
	flag.StringVar(&output, "output", "", "")
	// var target string
	// flag.StringVar(&target, "u", "","")
	// flag.StringVar(&target, "url", "", "")
	flag.Parse()

	var urls []string
	std := bufio.NewScanner(os.Stdin)
	for std.Scan() {
		var line string = std.Text()
		hline := strings.Replace(line, "%2F", "/", -1)
		line = hline
		// fmt.Println(line)

		urls = append(urls, line)

	}
	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go func(url string) {

			defer wg.Done()

			x := getParams(url, output)
			if x != "ERROR" {
				fmt.Println(x)
			}
			//fmt.Println(pms, url)

		}(u)
	}

	wg.Wait()

}

func getParams(url string, out string) string {

	var trans = &http.Transport{
		MaxIdleConns:      30,
		IdleConnTimeout:   time.Second,
		DisableKeepAlives: true,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout:   3 * time.Second,
			KeepAlive: time.Second,
		}).DialContext,
	}

	client := &http.Client{
		Transport: trans,
		Timeout:   3 * time.Second,
	}

	res, err := http.NewRequest("GET", url, nil)
	res.Header.Set("Connection", "close")
	resp, err := client.Do(res)
	// res.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36")

	if err != nil {
		return "ERROR"
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "ERROR"
	}

	page := string(body)
	r, _ := regexp.Compile(`[/^name=$/]{5}.*\s?\s?\s?\s?\s?\s?value=""`)
	match := r.FindAllString(page, -1)

	var params []string
	
	for _, param := range match {
		var cont int
		if strings.Contains(url,"?"){
			cont = 2 
		}else{
			cont = 1
		}
		s, _ := regexp.Compile(`name=".*?"`)
		match2 := s.FindString(param)
		name := strings.Split(match2, `"`)
		//matchSplitted := strings.Split(param, " ")
		//name := strings.Split(matchSplitted[2], `"`)
		if !strings.Contains(name[1], "__") {
			if cont > 1{
				fullparam := "&" + name[1] + "=airi"
				params = append(params, fullparam)
		}else{
			fullparam := "?" + name[1] + "=airi"
			params = append(params, fullparam)

		}}

	}
	urlFinal := url + strings.Join(params, "")
	if urlFinal != url {
		//fmt.Println(urlFinal)
		if out != "" {

			/*path, err := os.Getwd()
			  if err != nil {
			  fmt.Println(err)
			          } */
			data := urlFinal + "\n"
			file, err := os.Create(out)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()

			_, errg := file.WriteString(data)
			if errg != nil {
				fmt.Println(errg)
			}

		}
		return urlFinal
	}
	return "ERROR"
	// if out != false{
	// save file
	//}

	//matchSplitted := strings.Split(match, " ")
	//name := strings.Split(matchSplitted[2], `"`)

	//value[1]
	// return "NOT VULNERABLE!"

}

