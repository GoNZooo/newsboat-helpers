package header

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

type Header struct {
	Feed       string
	Title      string
	Author     string
	Date       time.Time
	Link       string
	PodcastUrl string
}

func (h Header) AsFilepath() string {
	sanitizer := strings.NewReplacer(
		" ", "-",
		"/", "_",
		"\\", "_",
		":", "_",
		"*", "_",
		"?", "_",
		"\"", "_",
		"<", "_",
		">", "_",
		"|", "_",
	)

	feed := sanitizer.Replace(h.Feed)
	title := sanitizer.Replace(h.Title)
	author := sanitizer.Replace(h.Author)
	date := h.Date.Format("2006-01-02")
	filename := fmt.Sprintf("%s/%s/%s_%s.mp3", author, feed, date, title)

	return strings.ToLower(filename)
}

func Parse(r io.Reader) (Header, error) {
	header := Header{}
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if header.Feed == "" && strings.HasPrefix(line, "Feed: ") {
			header.Feed = strings.TrimPrefix(line, "Feed: ")
			continue
		}

		if header.Title == "" && strings.HasPrefix(line, "Title: ") {
			header.Title = strings.TrimPrefix(line, "Title: ")
			continue
		}

		if header.Author == "" && strings.HasPrefix(line, "Author: ") {
			header.Author = strings.TrimPrefix(line, "Author: ")
			continue
		}

		if header.Date.IsZero() && strings.HasPrefix(line, "Date: ") {
			date, err := time.Parse(
				"Mon, 2 Jan 2006 15:04:05 -0700",
				strings.TrimPrefix(line, "Date: "),
			)
			if err != nil {
				return header, err
			}
			header.Date = date
			continue
		}

		if header.Link == "" && strings.HasPrefix(line, "Link: ") {
			header.Link = strings.TrimPrefix(line, "Link: ")
			continue
		}

		if header.PodcastUrl == "" && strings.HasPrefix(line, "Podcast Download URL: ") {
			split := strings.TrimPrefix(line, "Podcast Download URL: ")
			header.PodcastUrl = strings.Split(split, " ")[0]
			continue
		}
	}

	return header, nil
}
