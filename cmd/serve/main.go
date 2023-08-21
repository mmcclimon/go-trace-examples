package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
)

const (
	startPort = 9011
	indexPort = 8999
)

type trace struct {
	Name string
	Port int
}

var tpl = template.Must(template.New("html").Parse(`
<html>
<head>
	<title>megatrace</title>
<body>
	<ul>
	{{ range .}}
		<li><a href="http://localhost:{{.Port}}">{{.Name}}</a></li>
	{{ end }}
	</ul>
</html>
`))

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	traces := startTraceServers(ctx, &wg)

	addr := fmt.Sprintf("localhost:%d", indexPort)
	log.Printf("will listen on http://%s", addr)
	go listen(addr, traces)

	// wait for cancellation
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println("cleaning up")
	cancel()
	wg.Wait()

	log.Println("so long!")
}

func ordie(err error) {
	if err == nil {
		return
	}

	fmt.Println(err)
	os.Exit(1)
}

func startTraceServers(ctx context.Context, wg *sync.WaitGroup) []trace {
	files, err := os.ReadDir("traces")
	ordie(err)

	traces := make([]trace, len(files))

	// prevent the stupid thing from popping open browser windows
	os.Setenv("BROWSER", "/usr/bin/true")

	for i, ent := range files {
		t := trace{ent.Name(), startPort + i}
		traces[i] = t

		addr := fmt.Sprintf("-http=localhost:%d", t.Port)
		path := filepath.Join("traces", t.Name)
		args := []string{"tool", "trace", addr, path}

		goCmd := "go"

		if t.Name == "superstress-gc.trace" {
			if _, err := exec.LookPath("gotip"); err == nil {
				goCmd = "gotip"
			} else {
				log.Println("cannot find gotip in path; superstress-gc trace might fail")
			}
		}

		cmd := exec.CommandContext(ctx, goCmd, args...)
		cmd.Cancel = func() error {
			log.Println("killing process for " + t.Name)
			wg.Done()
			return cmd.Process.Kill()
		}

		log.Println("run:", goCmd, strings.Join(args, " "))

		err = cmd.Start()
		ordie(err)

		wg.Add(1)
	}

	return traces
}

func listen(addr string, traces []trace) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("handling request")
		tpl.Execute(w, traces)
	})

	err := http.ListenAndServe(addr, nil)
	if err != http.ErrServerClosed {
		ordie(err)
	}
}
