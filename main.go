package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/url"
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
			"[buffers] | airi",
			"+=======================================================+",
			"",
			" -h                   Show This Help Message",
			"",
			"+=======================================================+",
			"",
		}

		fmt.Fprintf(os.Stderr, strings.Join(help, "\n"))
	}

}

func main() {

	// var target string
	// flag.StringVar(&target, "u", "","")
	// flag.StringVar(&target, "url", "", "")
	flag.Parse()

	
	targets := make(chan string)

	std := bufio.NewScanner(os.Stdin)
	
	var wg sync.WaitGroup
	for i:=0;i<50;i++ {
		
			wg.Add(1)
			go func() {

				defer wg.Done()
				for v := range targets{

					_, err := url.Parse(v)
					if err != nil{
						continue
					}

					x := getParams(v)
					if x != "ERROR" {
						fmt.Println(x)
					}



				}

				
				//fmt.Println(pms, url)

			}()
		
		}
	for std.Scan() {
        var line string = std.Text()
        targets <- line

    	}
    close(targets)
    wg.Wait()

}

func getParams(url string) string {

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
	r, err := regexp.Compile(`\<input.*\s?\s?\s?\s?\s?\s?\s?\s?\s?(name=.*\s?\s?\s?\s?\s?value=.*)`)
	//[/^name=$/]{5}.*\s?\s?\s?\s?\s?\s?value=""
	//\s?\s?\s?\s?\s?\s?(name=".*value="")
	//\s?\s?\s?\s?\s?\s?(name=".*value="")
	//<input\s?\s?\s?type=.*(\s?\s?\s?\s?\s?\s?(name=".*)value="")
	if err != nil{
		return "ERROR"
	}
	match := r.FindAllString(page, -1)

	var params []string
	var cont2 int
	for _, param := range match {
		var cont int
		cont2 += 1
		if strings.Contains(url,"?"){
			cont = 2 
		}else{
			cont = 1
		}

		pattern := regexp.MustCompile(`\s+`)
		newstr := pattern.ReplaceAllString(param, " ")
		param_regex := regexp.MustCompile(`name="?'?[^\"\']+`)
		match2 := param_regex.FindString(newstr)
		paran := match2[6:]

		//fmt.Println(paran)

		if !strings.Contains(paran, "__") {
			if cont > 1{
				fullparam := "&" + paran + "=airi"
				params = append(params, fullparam)
			}else{
				if cont2 >= 2{
                        fullparam := "&" + paran + "=airi"
                        params = append(params, fullparam)
                }else{
                        fullparam := "?" + paran + "=airi"
                        //fmt.Println(cont2)
                        params = append(params, fullparam)
                                }

		}}

	}
	urlFinal := url + strings.Join(params, "")
	if urlFinal != url {
		//fmt.Println(urlFinal)
		return urlFinal
	}
	return "ERROR"
	
}
