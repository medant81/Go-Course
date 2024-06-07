/*
Измените код примера 4 так, чтобы лог попадал не в stdout, а в файл log.
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	fileLog, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer fileLog.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	l := log.New(fileLog, "", log.LstdFlags)
	logHandler := logMiddleware(l)

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: authHandler(logHandler(mux), l),
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalln(fmt.Errorf("Не удалось запустить сервер: %w", err))
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	msg := "Hello, Go!"
	log.Println("resp:", msg)
	fmt.Fprint(res, msg)
}

func authHandler(h http.Handler, l *log.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			xId := r.Header.Get("x-my-app-id")
			if xId != "my_secret" {
				http.Error(w, "пользователь не авторизован", http.StatusUnauthorized)
				l.Println("Unauthorized url: ", r.URL, "x-my-app-id: ", xId)
				return
			}
			h.ServeHTTP(w, r)
		})
}

func logMiddleware(l *log.Logger) func(h http.Handler) http.Handler {

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				l.Println("url:", r.URL)
				h.ServeHTTP(w, r)
			})
	}

}
