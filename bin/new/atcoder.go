package new

import (
	"fmt"
	"net/http"
	"strings"

	"ojx/bin/util"

	"github.com/PuerkitoBio/goquery"
)

func NewAtCoder(contestID string) (*util.Contest, error) {
	url := fmt.Sprintf("https://atcoder.jp/contests/%s/tasks", contestID)

	config, err := util.LoadConfig()
	if err != nil {
		return nil, err
	}

	session := config.Cookie.AtCoder

	resp, err := NewAtCoderRequest(url, session)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"unexpected status: %d",
			resp.StatusCode,
		)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	contest := &util.Contest{
		Platform:  "atcoder",
		ContestID: contestID,
	}

	doc.Find("table tbody tr").Each(func(_ int, tr *goquery.Selection) {
		cols := tr.Find("td")

		if cols.Length() < 4 {
			return
		}

		taskID := strings.ToLower(strings.TrimSpace(cols.Eq(0).Text()))

		taskURL, exists := cols.
			Eq(1).
			Find("a").
			Attr("href")

		if !exists {
			return
		}

		contest.Tasks = append(
			contest.Tasks,
			util.Task{
				TaskID: taskID,
				URL: fmt.Sprintf(
					"https://atcoder.jp%s",
					taskURL,
				),
				TimeLimit: parseAtCoderTimeLimit(
					cols.Eq(2).Text(),
				),
				MemoryLimit: parseAtCoderMemoryLimit(
					cols.Eq(3).Text(),
				),
			},
		)
	})

	if len(contest.Tasks) == 0 {
		return nil, fmt.Errorf(
			"no tasks found",
		)
	}

	return contest, nil
}

func parseAtCoderTimeLimit(timeStr string) float64 {
	var timeLimit float64

	fmt.Sscanf(
		strings.TrimSpace(timeStr),
		"%f sec",
		&timeLimit,
	)

	return timeLimit
}

func parseAtCoderMemoryLimit(memoryStr string) int {
	var memoryLimit int

	fmt.Sscanf(
		strings.TrimSpace(memoryStr),
		"%d MiB",
		&memoryLimit,
	)

	return memoryLimit
}
