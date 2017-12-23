package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"ninex"
	"path"

	qrcode "github.com/skip2/go-qrcode"
)

func serve(ctx context.Context, listen string, pingCh chan<- struct{}) {
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    listen,
		Handler: mux,
	}

	mux.HandleFunc("/", logHandler(serveHome))
	mux.HandleFunc("/info", logHandler(serveInfo))
	mux.HandleFunc("/ui.js", logHandler(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Join(*contentDir, "ui.js"))
	}))
	mux.HandleFunc("/qr", logHandler(serveQR))
	if setCookiePath != "" {
		mux.HandleFunc(setCookiePath, logHandler(setCookie))
		mux.HandleFunc("/admin/shutdown", logHandler(adminHandler(shutdown(ctx, srv))))
		mux.HandleFunc("/admin/ping", logHandler(adminHandler(ping(pingCh))))
	}

	err := srv.ListenAndServe()
	if err == http.ErrServerClosed {
		log.Print("orderly shutdown")
		ctxCancel()
		return
	}
	log.Fatal(err)
}

func logHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("serving %s", r.URL)
		next(w, r)
	}
}

func adminHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("admin")
		if err != nil {
			httpErr(w, "missing admin cookie", http.StatusForbidden)
			return
		}
		if c.Value != hex.EncodeToString(adminSecret[:]) {
			httpErr(w, "invalid admin cookie", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}

func shutdown(ctx context.Context, srv *http.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := srv.Shutdown(ctx)
		if err != nil {
			httpErr(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func ping(pingCh chan<- struct{}) http.HandlerFunc {
	return func(http.ResponseWriter, *http.Request) {
		pingCh <- struct{}{}
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	tmplName := path.Join(*contentDir, "home.html.tmpl")
	tmpl, err := template.ParseFiles(tmplName)
	if err != nil {
		httpErr(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmplData := map[string]interface{}{
		"contractAddr": hex.EncodeToString(contractAddr[:]),
		"abi":          ninex.NinexABI,
	}
	err = tmpl.Execute(w, tmplData)
	if err != nil {
		httpErr(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func serveInfo(w http.ResponseWriter, r *http.Request) {
	info, err := getInfo()
	if err != nil {
		httpErr(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(info)
	if err != nil {
		httpErr(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func serveQR(w http.ResponseWriter, r *http.Request) {
	q, err := url.QueryUnescape(r.URL.RawQuery)
	if err != nil {
		httpErr(w, err.Error(), http.StatusInternalServerError)
		return
	}
	png, err := qrcode.Encode(q, qrcode.Medium, 256)
	if err != nil {
		httpErr(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	_, err = w.Write(png)
	if err != nil {
		httpErr(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

var setCookieEnabled = true

func setCookie(w http.ResponseWriter, r *http.Request) {
	if !setCookieEnabled {
		httpErr(w, "setcookie endpoint disabled", http.StatusForbidden)
		return
	}
	cookie := &http.Cookie{
		Name:  "admin",
		Value: hex.EncodeToString(adminSecret[:]),
	}
	http.SetCookie(w, cookie)
	setCookieEnabled = false
}

func httpErr(w http.ResponseWriter, msg string, code int) {
	log.Print(msg)
	http.Error(w, msg, code)
}
