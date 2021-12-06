package aocinput

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetInputFromDay(day int, session_cookie string) (string, error) {
	const base_url = "https://adventofcode.com/2021/day/"
	day_url := base_url + strconv.Itoa(day) + "/input"

	reader, err := downloadFile(day_url, session_cookie)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[X] ERROR: GetInputFromDay failed to download file.")
		return "", err
	}

	rc := *reader
	contents, err := ioutil.ReadAll(rc)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[X] ERROR: GetInputFromDay failed to read file: ", err)
		return "", err
	}

	return string(contents), nil
}

func downloadFile(url string, session_cookie string) (*io.ReadCloser, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Cookie", "session="+session_cookie)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Fprintln(os.Stderr, "[X] ERROR: DownloadFile failed when getting http response: \n", err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		fmt.Println("[!] WARNING: DownloadFile response code: ", strconv.Itoa(resp.StatusCode))
		fmt.Println("Attempting to continue...")
	}
	if resp.Body != nil {
		return &resp.Body, nil
	}
	return nil, errors.New("DownloadFile failed due to no response body from url.")
}

func main() {
	cookie := os.Args[1]
	input, err := GetInputFromDay(2, cookie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)
}
