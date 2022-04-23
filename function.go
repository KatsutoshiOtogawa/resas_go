// Package p contains an HTTP Cloud Function.
package p

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HelloWorld", HelloWorld)
}

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func HelloWorld(w http.ResponseWriter, r *http.Request) {

	api_url := "https://opendata.resas-portal.go.jp/api/v1/prefectures"
	req, err := http.NewRequest("GET", api_url, nil)

	if err != nil {

		// bad requestの原因を書く。
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	req.Header.Set("X-API-KEY", os.Getenv("PREFECUTRE_API_KEY"))
	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// reponse
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// jsonをクライアントに返す。
	fmt.Fprint(w, string(byteArray))
}
