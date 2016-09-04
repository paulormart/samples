package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/context"
	"time"
	"log"
)

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/earth", handler2)
	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hellow World")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hellow Earth")
}

// Handler using context package
// https://blog.golang.org/context
// e.g.
// www.aukbit.com/search?q=golang&timeout=1s
//
func handlerAukbit(w http.ResponseWriter, req *http.Request) {
	// ctx is the Context for this handler. Calling cancel closes the
    // ctx.Done channel, which is the cancellation signal for requests
    // started by this handler.
    var (
        ctx    context.Context
        cancel context.CancelFunc
    )
    timeout, err := time.ParseDuration(req.FormValue("timeout"))
    if err == nil {
        // The request has a timeout, so create a context that is
        // canceled automatically when the timeout expires.
        ctx, cancel = context.WithTimeout(context.Background(), timeout)
    } else {
        ctx, cancel = context.WithCancel(context.Background())
    }
    defer cancel() // Cancel ctx as soon as handlerAukbit returns.
	// Check the search query.
    query := req.FormValue("q")
    if query == "" {
        http.Error(w, "no query", http.StatusBadRequest)
        return
    }

    // Store the user IP in ctx for use by code in other packages.
    userIP, err := userip.FromRequest(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    ctx = userip.NewContext(ctx, userIP)

	//// Run the Google search and print the results.
    //start := time.Now()
    //results, err := google.Search(ctx, query)
    //elapsed := time.Since(start)
	//if err := resultsTemplate.Execute(w, struct {
    //    Results          google.Results
    //    Timeout, Elapsed time.Duration
    //}{
    //    Results: results,
    //    Timeout: timeout,
    //    Elapsed: elapsed,
    //}); err != nil {
    //    log.Print(err)
    //    return
    //}
	return
}
