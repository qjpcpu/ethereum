package etherscan

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/http/cookiejar"
	"time"
)

var httpClient *http.Client

func init() {
	jar, _ := cookiejar.New(nil)
	httpClient = &http.Client{
		Timeout: 10 * time.Second,
		Jar:     jar,
	}
}

type Env string

const (
	Online  Env = `https://etherscan.io`
	Ropsten     = `https://ropsten.etherscan.io`
)

func PendingTxs(env Env, owner string) ([]string, error) {
	var uri string
	if owner != "" {
		uri = fmt.Sprintf("%s/txsPending?a=%s", env, owner)
	} else {
		uri = fmt.Sprintf("%s/txsPending", env)
	}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.79 Safari/537.36`)
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	var hashes []string
	doc.Find(`tbody .address-tag a[href^="/tx"]`).Each(func(i int, s *goquery.Selection) {
		hashes = append(hashes, s.Text())
	})
	return hashes, nil
}
