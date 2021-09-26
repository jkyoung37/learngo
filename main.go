package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	link     string
	title    string
	location string
	summary  string
}

var baseURL string = "https://jp.indeed.com/求人?q=python&limit=50"

func main() {
	var jobs []extractedJob
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...)
	}
	fmt.Println(jobs)
	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)
	utf8bom := []byte{0xEF, 0xBB, 0xBF}
	file.Write(utf8bom)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"title", "location", "summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{job.title, job.location, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}
func getPage(page int) []extractedJob {
	var jobs []extractedJob
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)

	checkErr(err)
	searchCards := doc.Find(".tapItem")

	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})
	return jobs
}

func extractJob(card *goquery.Selection) extractedJob {
	link, _ := card.Attr("href")
	title := cleanString(card.Find(".jobTitle>span").Text())
	location := cleanString(card.Find(".companyLocation").Text())
	summary := cleanString(card.Find(".job-snippet").Text())
	return extractedJob{
		link:     link,
		title:    title,
		location: location,
		summary:  summary,
	}
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)

	checkErr(err)
	doc.Find(".pagination-list").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatal("Request failed with Status", res.StatusCode)
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
