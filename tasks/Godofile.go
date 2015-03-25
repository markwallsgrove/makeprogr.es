package main

import (
	"fmt"
	. "gopkg.in/godo.v1"
)

func tasks(p *Project) {
	Env = `GOPATH=$GOPATH/.vendor::$GOPATH`

	p.Task("default", D{"dev"})

	p.Task("init?", func() {
		fmt.Println(" -> creating directories")
		Bash("mkdir pkg bin src/github.com/markwallsgrove/makeprogr.es/static/bower")
		fmt.Println(" -> installing libs")
		Bash("go get github.com/nitrous-io/goop")
		fmt.Println(" -> running goop")
		Bash("bin/goop install")
		fmt.Println(" -> running bower")
		Bash("bower install")
	})

	p.Task("generate", func() {
		Bash(`
			bin/esc -o src/github.com/markwallsgrove/makeprogr.es/main/static.go \
			-pkg main -prefix src/github.com/markwallsgrove/makeprogr.es \
			src/github.com/markwallsgrove/makeprogr.es/static
		`)
		// TODO: minify css
		// TODO: minify html
		// TODO: minify js
	})

	p.Task("validate", func() {
		// TODO: lint js
		// TODO: lint css
	})

	p.Task("build", D{"init", "generate", "validate"}, func() {
		Run("GOOS=linux GOARCH=amd64 go install", In{"src/github.com/markwallsgrove/makeprogr.es/main"})
	})

	p.Task("dev", D{"build"}, func() {
		Start("main.go", M{"$in": "src/github.com/markwallsgrove/makeprogr.es/main"})
	}).Watch("src/github.com/markwallsgrove/makeprogr.es/**/*.go").
		Debounce(3000)

}

func main() {
	Godo(tasks)
}
