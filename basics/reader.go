package basics

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	res, _ := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	defer res.Body.Close()
	if bs, err := ioutil.ReadAll(res.Body); err == nil {
		fmt.Print(string(bs))
	}
}