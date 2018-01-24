package main

import (
    "log"
    "fmt"
    "github.com/valyala/fasthttp"
)

func main() {
    h := H
    if e := fasthttp.ListenAndServe("0.0.0.0:5000", h); e != nil {
        log.Fatalf("%s", e)
    }
}

func H(ctx *fasthttp.RequestCtx) {
    switch string(ctx.Path()) {
    case "/helloworld":
        ctx.SetContentType("text/json; charset=utf-8")
        fmt.Fprintf(ctx, `{"message":"Hello World!"}`)
    //default: commented out to improve the performance.
    //  ctx.Error("Unsupported path", fasthttp.StatusNotFound)
    }
}
