package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"ncaptcha/question"
	"net/http"
	"net/url"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

type challenge struct {
	answers []int
	expire  time.Time
}

func randomToken() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

func main() {
	publicApi := os.Getenv("NCAPTCHA_API")

	challenges := sync.Map{} // make(map[string]*challenge)
	tokens := sync.Map{}     // make(map[string]time.Time) token and expiry time

	mux := &http.ServeMux{}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, "404 not found")
			return
		}
		io.WriteString(w, "hello ncaptcha")
	})

	mux.HandleFunc("/assets/ncaptcha.js", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("./assets/ncaptcha.js")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		content = bytes.Replace(content, []byte("<PUBLIC_API>"), []byte(publicApi), 1)
		w.Header().Set("Content-Type", "application/javascript")
		w.WriteHeader(http.StatusOK)
		w.Write(content)
	})

	mux.HandleFunc("/assets/ncaptcha.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/ncaptcha.css")
	})

	mux.HandleFunc("/assets/checkmark-circle.svg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/checkmark-circle.svg")
	})
	mux.HandleFunc("/assets/checkmark.svg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/checkmark.svg")
	})

	mux.HandleFunc("/assets/icon.svg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/icon.svg")
	})

	mux.HandleFunc("/demo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			form := url.Values{}
			form.Set("token", r.FormValue("ncaptcha-response"))

			resp, err := http.PostForm("http://127.0.0.1:8080/verify", form)
			if err != nil {
				log.Println("error while verifying token", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println("error while verifying token", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if string(body) != "ok" {
				io.WriteString(w, "your ncaptcha response is invalid")
				return
			}

			w.Header().Set("Content-Type", "text/html")

			name := html.EscapeString(r.FormValue("name"))
			hello := html.EscapeString(r.FormValue("hello"))

			io.WriteString(w, fmt.Sprintf("<h1>Hello, %s</h1><p>hello, %s</p>", name, hello))

		} else {

			http.ServeFile(w, r, "./assets/demo.html")
		}

	})

	// verify a token
	mux.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		token := r.FormValue("token")
		t, ok := tokens.LoadAndDelete(token)
		if !ok {
			io.WriteString(w, "invalid")
			return
		}

		if t.(time.Time).Before(time.Now()) {
			io.WriteString(w, "invalid")
			return
		}

		log.Println("verify", token)

		io.WriteString(w, "ok")
	})

	// generate a question
	mux.HandleFunc("/challenge", func(w http.ResponseWriter, r *http.Request) {
		img, s, ans, err := question.NewQuestion()
		if err != nil {
			log.Println("error while generating challenge:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		challengeId := randomToken()
		challenges.Store(challengeId, &challenge{
			answers: ans,
			expire:  time.Now().Add(time.Minute * 3),
		})

		log.Println("new challenge", challengeId, ans)

		w.Header().Set("Content-Type", "application/json")

		resp := struct {
			Challenge string `json:"challenge"`
			Id        string `json:"id"`
			Select    string `json:"select"`
		}{}

		resp.Challenge = base64.StdEncoding.EncodeToString(img)
		resp.Id = challengeId
		resp.Select = s

		json.NewEncoder(w).Encode(resp)
	})

	// answer the question and obtain a token
	mux.HandleFunc("/answer", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		// retrieve the challenge
		q, ok := challenges.LoadAndDelete(r.FormValue("challenge"))
		if !ok {
			io.WriteString(w, "challenge not found")
			return
		}

		theChallenge := q.(*challenge)

		if theChallenge.expire.Before(time.Now()) {
			io.WriteString(w, "challenge was expired")
			return
		}

		// read answers
		ansString := strings.Split(r.FormValue("ans"), ",")
		ans := make([]int, 0)

		// convert strings to integers
		for _, v := range ansString {
			if v == "" {
				continue
			}
			n, err := strconv.Atoi(v)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			ans = append(ans, n)
		}

		// compare answers
		slices.Sort(ans)
		if slices.Compare(ans, theChallenge.answers) != 0 {
			// wrong answer
			io.WriteString(w, "wrong answer")
			return
		}

		token := randomToken()
		tokens.Store(token, time.Now().Add(time.Minute*5))

		log.Println("new token", token)

		io.WriteString(w, "TOKEN_"+token)
	})

	log.Println("listening at 127.0.0.1:8080")

	err := http.ListenAndServe("127.0.0.1:8080", mux)
	if err != nil {
		log.Println("could not start server: " + err.Error())
	}
}
