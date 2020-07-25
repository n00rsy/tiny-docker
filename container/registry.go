package container

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	url := "https://raw.githubusercontent.com/n00rsy/ubernetes/master/ExampleImage.txt"
	out, err := os.Create("ExampleImage.txt")
	must(err)

	defer out.Close()

	resp, err := http.Get(url)
	must(err)
	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)
	must(err)
	fmt.Println(n)
}

