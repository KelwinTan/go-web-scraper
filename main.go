package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	colly "github.com/gocolly/colly/v2"
	debug "github.com/gocolly/colly/v2/debug"
)

func main() {

	// c := colly.NewCollector(
	// 	colly.AllowedDomains("indeed.com"),
	// )

	// url := "https://www.indeed.com/jobs?q=USA+Landscaping&redirected=1"

	fName := "data5.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))
	// c.OnHTML("ul.hover", func(e *colly.HTMLElement) {
	// 	e.ForEach("li.react-job-listing", func(_ int, el *colly.HTMLElement) {
	// 		writer.Write([]string{
	// 			el.ChildText("a.jobLink:nth-child(1)"),
	// 			el.ChildText("a.jobLink:nth-child(2)"),
	// 			// el.ChildText("td:nth-child(2)"),
	// 			// el.ChildText("td:nth-child(3)"),
	// 		})
	// 	})
	// 	fmt.Println("Scrapping Complete")
	// })
	// c.Visit("https://www.glassdoor.com/Job/us-landscaping-jobs-SRCH_IL.0,2_IN1_KO3,14.htm")

	c.OnHTML("ul#job-list.jobs", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			writer.Write([]string{
				el.ChildText("h3."),
				el.ChildText("a.jobLink:nth-child(2)"),
				// el.ChildText("td:nth-child(2)"),
				// el.ChildText("td:nth-child(3)"),
			})
		})
		fmt.Println("Scrapping Complete")
	})
	c.Visit("https://www.simplyhired.com/search?q=landscaping&l=united+states&job=0S2qm_8MTqqdt6lHKf-70GLK93WeBquSW8ea8Enm_muHbXz00prB5A")

}
