package etherscan

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"strings"
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

type PendingTx struct {
	Hash     string `json:"hash"`
	From     string `json:"from"`
	To       string `json:"to"`
	Value    string `json:"value"`
	GasLimit uint64 `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Nonce    uint64 `json:"nonce"`
}

func GetBlockedPendingTx(env Env, owner string, pending_nonce uint64) (PendingTx, error) {
	if owner == "" {
		return PendingTx{}, errors.New("no owner")
	}
	hashes, err := PendingTxs(env, owner)
	if err != nil {
		return PendingTx{}, err
	}
	if len(hashes) == 0 {
		return PendingTx{}, errors.New("no pending tx")
	}
	for _, hash := range hashes {
		detail, err := PendingTxDetail(env, hash)
		if err != nil {
			return detail, err
		}
		if pending_nonce == 0 || pending_nonce == detail.Nonce {
			return detail, nil
		}
	}
	return PendingTx{}, errors.New("no pending tx")
}

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

func PendingTxDetail(env Env, txhash string) (PendingTx, error) {
	var detail PendingTx
	uri := fmt.Sprintf("%s/tx/%s", env, txhash)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return detail, err
	}
	req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.79 Safari/537.36`)
	res, err := httpClient.Do(req)
	if err != nil {
		return detail, err
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return detail, err
	}
	detail.Hash = txhash
	for loop := true; loop; loop = false {
		doc.Find(`.container a[href^="/address"]`).Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				detail.From = s.Text()
			case 1:
				detail.To = s.Text()
			}
		})
		detail.Value = trimEmpty(doc.Find(`.container span[title^="The amount of ETH to be transferred to the recipient"]`).First().Text())
		limit := trimEmpty(doc.Find(`.container span[title^="The amount of GAS supplied for this transaction to happen"]`).First().Text())
		detail.GasLimit, _ = strconv.ParseUint(limit, 10, 64)
		detail.GasPrice = trimEmpty(doc.Find(`.container span[title^="The price offered to the miner to purchase this amount of GAS"]`).First().Text())
		doc.Find(`.container div .col-sm-9.cbs`).Each(func(i int, s *goquery.Selection) {
			if strings.Contains(s.Text(), "{Pending}") {
				str := strings.Replace(s.Text(), "{Pending}", "", -1)
				str = strings.Replace(str, "|", "", -1)
				detail.Nonce, _ = strconv.ParseUint(trimEmpty(str), 10, 64)
			}
		})
	}
	if detail.Nonce == 0 {
		return detail, errors.New("can't find tx nonce")
	}
	return detail, nil
}

func trimEmpty(raw string) string {
	raw = strings.Replace(raw, "\n", "", -1)
	raw = strings.Replace(raw, "\r", "", -1)
	return strings.TrimSpace(raw)
}
