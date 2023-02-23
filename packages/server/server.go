package server

import (
    "embed"
    "net/http"
    "D3VL/WebWrap/packages/logging"
)

func Start(files embed.FS, port string) {
    log.Debug("Starting server...")

    // serve all files in the static folder at the root of the server
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = "/static" + r.URL.Path
        log.Debug("Serving file: " + r.URL.Path)
        http.FileServer(http.FS(files)).ServeHTTP(w, r)
    })

    log.Debug("Server started.")

    log.Debug("Listening on port " + port + "...")

    err := http.ListenAndServe("127.0.0.1:" + port, nil)
    if err != nil {
        log.Error("Error starting server: " + err.Error())
    }
}
