package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"

	"github.com/bradhe/stopwatch"
)

const (
	url      = "http://localhost:3000/hashpls"
	keyLen   = 32
	password = "hunter2"
	passLen  = 7
)

func main() {
	repeat := int(math.Ceil(float64(passLen) / float64(len(password))))
	payload := strings.NewReader(fmt.Sprintf("password=%s&keyLen=%d", strings.Repeat(password, repeat), keyLen))
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	sw := stopwatch.Start()
	res, _ := http.DefaultClient.Do(req)
	sw.Stop()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Printf("Request done in %dms. Result: %s\n", sw.Milliseconds(), body)
}
