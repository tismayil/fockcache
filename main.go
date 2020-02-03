package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	"github.com/christophwitzko/go-curl"
)

var InjecTedHeaderSite = "cachefock.dotphp.net"
var userAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0_1 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A470a Safari/9537.53"

func randomZen() string {
	var RandomNums = rand.Intn(9999999)
	return strconv.Itoa(RandomNums)
}

func sendReq(hostname string, num int) string {

	urlez := hostname + "/" + randomZen() + "=" + randomZen()
	if num == 1 {

		header := http.Header{
			"X-Forwarded-Host": {InjecTedHeaderSite},
			"Accept":           {"*/*; q=0.01"},
			"Accept-Language":  {"en-US,en;q=0.5"},
			"Accept-Encoding":  {"gzip, deflate"},
			"User-Agent":       {userAgent},
		}
		err, str, _ := curl.String(urlez, "file", "header=", header)
		if err != nil {
			return ""
		}

		R := checkResponseFock(str)

		if R == true {
			return urlez
		} else {
			return "false"
		}

	} else {

		header := http.Header{
			"X-Forwarded-Host":   {InjecTedHeaderSite},
			"X-Forwarded-Scheme": {"nothttps"},
			"Accept":             {"*/*; q=0.01"},
			"Accept-Language":    {"en-US,en;q=0.5"},
			"Accept-Encoding":    {"gzip, deflate"},
			"User-Agent":         {userAgent},
		}
		err, str, _ := curl.String(urlez, "file", "header=", header)
		if err != nil {
			return ""
		}

		R := checkResponseFock(str)

		if R == true {
			return urlez
		} else {
			return "false"
		}
	}

}

func checkResponseFock(response string) bool {
	cacheFocked := regexp.MustCompile(InjecTedHeaderSite)
	if len(cacheFocked.FindStringSubmatch(response)) > 0 {
		return true
	}
	return false
}

func main() {

	hostname := flag.String("hostname", "", "Please input hostname")
	flag.Parse()

	fmt.Println(`
	______         _     _____           _          
	|  ____|       | |   / ____|         | |         
	| |__ ___   ___| | _| |     __ _  ___| |__   ___ 
	|  __/ _ \ / __| |/ / |    / _' |/ __| '_ \ / _ \
	| | | (_) | (__|   <| |___| (_| | (__| | | |  __/
	|_|  \___/ \___|_|\_\\_____\__,_|\___|_| |_|\___|
													 
	Cache Poisoning Tester
	Github : github.com/tismayil
	Host : ` + *hostname + `
	`)

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	s.Restart()

	fmt.Printf("\033[1;34m%s\033[0m", "1 - Test Check Without Any Extra Param : \n")
	for i := 1; i <= 10; i++ {
		req := sendReq(*hostname, 1)

		if req != "false" {
			s.Restart()
			fmt.Printf("\033[1;36m%s\033[0m", "Testing Injecting /X-Forwarded-Host/ : Focked : "+req+"\n")

		} else {
			s.Restart()
			fmt.Printf("\033[1;31m%s\033[0m", "Testing Injecting /X-Forwarded-Host/ : Req #"+strconv.Itoa(i)+" False\n")
		}
	}

	s.Restart()
	fmt.Printf("\033[1;34m%s\033[0m", "\n2 - Test Add New Header /X-Forwarded-Scheme/ Start Test : \n")
	for i := 1; i <= 10; i++ {
		req := sendReq(*hostname, 2)

		if req != "false" {
			s.Restart()
			fmt.Printf("\033[1;36m%s\033[0m", "Testing Injecting /X-Forwarded-Scheme/ : Focked : "+req+"\n")
		} else {
			s.Restart()
			fmt.Printf("\033[1;31m%s\033[0m", "Testing Injecting /X-Forwarded-Scheme/ : Req #"+strconv.Itoa(i)+" False\n")
		}
	}

	fmt.Printf("\033[1;34m%s\033[0m", "\n Scan Completed. \n")
	s.Restart()
	s.Stop()

}
