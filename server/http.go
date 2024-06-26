package server

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/devopsext/trends-back/trends"
)

func Serve() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	http.HandleFunc("/summary", func(w http.ResponseWriter, r *http.Request) {
		trendsSummary, _ := trends.GetTrendsSummary()
		fmt.Fprintf(w, "%v", string(trendsSummary))
	})
	http.HandleFunc("/trends", func(w http.ResponseWriter, r *http.Request) {
		trends, _ := trends.GetTrends()
		fmt.Fprintf(w, "%v", string(trends))
	})

	http.ListenAndServe(":8080", nil)
}
