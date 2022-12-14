package main

import (
	"bufio"
	"flag"
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Setting up Gin
	r := gin.Default()

	// CORS
	r.Use(cors.Default())

	r.GET("/", ledger)
	r.GET("/hello", hello_img)
	r.GET("/ping", ping)
	r.GET("/shutdown", shutdown)

	return r
}

func main() {
	// parse parameters
	var portParam = flag.Int("p", 8080, "Specify port. The default is 8080.")
	flag.Parse()
	var port = strconv.Itoa(*portParam)

	// define routes
	r := setupRouter()

	// todo: make the port configurable
	//r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	var url = "localhost:" + port
	r.Run(url)
}

// The Go way to split the string at newline characters.
func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
