package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()

	feed, err := fp.ParseURL("https://owllife.kennesaw.edu/events.rss")
	if err != nil {
		fmt.Println("Error parsing RSS feed:", err)
		return
	}

	if len(feed.Items) == 0 {
		fmt.Println("No items found in the RSS feed.")
		return
	}

	//fmt.Println(feed.String())
	var mar, ken, other, online int = 0, 0, 0, 0

	for _, item := range feed.Items {
		fmt.Println("-------------------------")
		fmt.Println("\nTitle: ", item.Title)

		fmt.Println("\nLink: ", item.Link)
		/*	fmt.Println("\nAuthor ", item.Author)
			fmt.Println("\nImage ", item.Image)*/
		for index, category := range item.Categories {
			fmt.Printf("Categ. %d: %s\n", index+1, category)
		}
		//fmt.Println("\nCategories", item.Categories)
		fmt.Println("\nLocation: ", item.Custom["location"])

		//fmt.Println("\nDesc. ", item.Description)
		fmt.Println("\nDesc: ", parseDesc(item.Description))
		fmt.Println("\nDateTime: ", parseDate(item.Description))
		//fmt.Println("\nLocation. ", parseStrings(item.Custom, `location:`, `start`))

		//fmt.Println("Desc:", item.Description)
		campus := ("\nModality: " + modality(item.Categories, item.Custom["location"]))
		if modality(item.Categories, item.Custom["location"]) == `Kennesaw` {
			ken++
		}
		if modality(item.Categories, item.Custom["location"]) == `Online` {
			online++
		}
		if modality(item.Categories, item.Custom["location"]) == `Marietta` {
			mar++
		}
		if modality(item.Categories, item.Custom["location"]) == `Other` {
			other++
		}

		fmt.Println(campus)

	}

	fmt.Printf("Mar. %d, Ken. %d, Onl. %d, Other %d", mar, ken, online, other)

}

func modality(categories []string, location string) (out string) {
	if strings.Contains(location, `.com`) {
		return `Online`
	}
	if strings.Contains(location, `Marietta`) {
		return `Marietta`
	}
	if strings.Contains(location, `Kennesaw`) {
		return `Kennesaw`
	}

	for _, category := range categories {
		if strings.Contains(category, `Kennesaw`) {
			return `Kennesaw`
		}
		if strings.Contains(category, `Marietta`) {
			return `Marietta`
		}
	}
	return `Other`
}

func parseDesc(inString string) (out string) {
	startIndex := strings.Index(inString, `<div class="p-description description"><p>`)
	if startIndex == -1 {
		return "N/A"
	}

	endIndex := startIndex + strings.Index(inString[startIndex:], "</div>")
	if endIndex == -1 {
		return "N/A"
	}

	strippedString := inString[startIndex:endIndex]
	re := regexp.MustCompile(`<.*?>`)
	return re.ReplaceAllString(strippedString, "")
}

func parseStrings(inString, before, after string) (out string) {
	startIndex := strings.Index(inString, before)
	if startIndex == -1 {
		return
	}

	endIndex := startIndex + strings.Index(inString[startIndex:], after)
	if endIndex == -1 {
		return
	}

	return inString[(startIndex + len(before)):endIndex]
}

func parseDate(inString string) (out string) {
	startIndex := strings.Index(inString, `<time class="dt-start dtstart"`)
	if startIndex == -1 {
		return "N/A"
	}

	endIndex := startIndex + strings.Index(inString[startIndex:], `</time>`)
	if endIndex == -1 {
		return "N/A"
	}

	strippedString := inString[startIndex:endIndex]
	re := regexp.MustCompile(`<.*?>`)
	return re.ReplaceAllString(strippedString, "")
}
