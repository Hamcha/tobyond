package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)
const index = `<!doctypehtml><html lang=en><meta charset=UTF-8><title>Byond redirect service</title><style>body,html{height:100%;overflow:hidden}body{background:linear-gradient(to top,#012,#456) #123;text-align:center;font-family:sans-serif;display:flex;flex-direction:column;justify-content:center;color:#eee}h1,p{line-height:1.5rem;padding:.6rem 1rem;margin:0}b{background-color:#345;padding:3px;color:#fff}</style><h1>BYOND redirector service</h1><p>Having issues linking to BYOND servers on Telegram / etc?<p>Replace <b>byond://<span style=color:#ff0>your.server.url:port</span></b> with <b>https://byond.ovo.ovh/<span style=color:#ff0>your.server.url:port </span></b>and you're set!`
const template = `<!doctypehtml><html lang=en><meta charset=UTF-8><title>Byond redirect</title><meta content="1;url=byond://%s"http-equiv=refresh><style>body,html{height:100%%;overflow:hidden}body{background:linear-gradient(to top,#012,#456) #123;text-align:center;font-family:sans-serif;display:flex;flex-direction:column;justify-content:center;color:#eee}h1,p{line-height:1.5rem;padding:.6rem 1rem;margin:0}</style><h1>You are being redirected to BYOND!</h1><p>If a popup doesn't come up or the client doesn't open in the next few seconds something has probably gone wrong!`

func main() {
	bind := os.Getenv("BIND")
	if bind == "" {
		bind = ":8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Path) < 2 {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprint(w, index)
			return
		}
		// Make sure URL is well-formed
		if strings.IndexRune(r.URL.Path, ':') <= 0 {
			http.Error(w, "URL is malformed, should be server.address:port", 400)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, template, r.URL.Path[1:])
	})
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		panic(err)
	}
}