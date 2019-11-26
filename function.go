package vader

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

//GoogleCloudFunctionHTTP because we wrote it in go already
func GoogleCloudFunctionHTTP(w http.ResponseWriter, r *http.Request) {
	text := ""
	reqPath, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		http.Error(w, "Please specify a sentence as the body or path of your http POST request "+err.Error(), http.StatusBadRequest)
		return
	}

	//find this dang sentence somewhere, anywhere!
	switch {
	case len(reqPath.Path) > 10:
		text = reqPath.Path
	case len(reqPath.RawQuery) > 10:
		text = reqPath.RawQuery
	case len(reqPath.Fragment) > 10:
		text = reqPath.Fragment
	}

	//check the body
	if text == "" {
		if r.Body == nil {
			http.Error(w, "Please specify a sentence as the body or path of your http POST request", http.StatusBadRequest)
			return
		}
		bodyReader, err := r.GetBody()
		if err != nil {
			http.Error(w, "Please specify a sentence as the body or path of your http POST request "+err.Error(), http.StatusBadRequest)
			return
		}
		body, err := ioutil.ReadAll(bodyReader)
		if err != nil {
			http.Error(w, "Please specify a sentence as the body or path of your http POST request "+err.Error(), http.StatusBadRequest)
			return
		}
		text = string(body)
	}
	out := json.NewEncoder(w)
	err = out.Encode(GetSentiment(text))
	if err != nil {
		http.Error(w, "Please specify a sentence as the body or path of your http POST request "+err.Error(), http.StatusBadRequest)
	}
}
