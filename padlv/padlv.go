package padlv

import (
        "net/http"
        "regexp"
)

func init() {
        http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
        path := r.URL.Path
        if path[0] != '/' {
                http.NotFound(w, r)
        }
        path = path[1:]

        if matched, _ := regexp.MatchString("^[0-9]+$", path); matched {
                http.Redirect(w, r,
                        "http://pad.lv/"+path,
                        http.StatusTemporaryRedirect)
        } else {
                http.NotFound(w, r)
        }
}
