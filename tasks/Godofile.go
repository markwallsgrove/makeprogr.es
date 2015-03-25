package main

import (
	"fmt"
	. "gopkg.in/godo.v1"
)

func tasks(p *Project) {
	Env = `GOPATH=.vendor::$GOPATH`

	p.Task("default", D{"dev"})

	p.Task("init?", func() {
		fmt.Println(" -> creating directories")
		Bash("mkdir pkg bin")
		fmt.Println(" -> installing libs")
		Bash("go get github.com/nitrous-io/goop")
		fmt.Println(" -> running goop")
		Bash("bin/goop install")
		fmt.Println(" -> running bower")
		Bash("bower install")
	})

	p.Task("generate", func() {
		Bash("bin/esc -o src/github.com/talis/makeprogr.es/main/static.go -pkg main -prefix src/github.com/talis/makeprogr.es src/github.com/talis/makeprogr.es/static")
		// TODO: minify css
		// TODO: minify html
		// TODO: minify js
	})

	p.Task("validate", func() {
		// TODO: lint js
		// TODO: lint css
	})

	p.Task("build", D{"init", "generate", "validate"}, func() {
		Run("GOOS=linux GOARCH=amd64 go install", In{"src/github.com/talis/makeprogr.es/main"})
	})

	p.Task("dev", D{"build"}, func() {
		Start("main.go", M{"$in": "src/github.com/talis/makeprogr.es/main"})
	}).Watch("src/github.com/talis/makeprogr.es/**/*.go").
		Debounce(3000)

}

func main() {
	Godo(tasks)
}
