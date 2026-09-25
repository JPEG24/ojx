package new

import (
	"fmt"
	"net/http"
	"strings"

	"ojx/bin/util"

	"github.com/PuerkitoBio/goquery"
)

func NewCodeforces(contestID string) (*util.Contest, error) {
	url := fmt.Sprintf("https://codeforces.com/contest/%s", contestID)

	resp, err := NewRequest(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	contest := &util.Contest{
		Platform:  "codeforces",
		ContestID: contestID,
	}

	doc.Find("table.problems tr").Each(func(_ int, tr *goquery.Selection) {
		cols := tr.Find("td")

		if cols.Length() < 2 {
			return
		}

		taskID := strings.ToLower(strings.TrimSpace(cols.Eq(0).Text()))
		taskURL, exists := cols.Eq(0).Find("a").Attr("href")
		if !exists {
			return
		}

		notice := strings.TrimSpace(cols.Eq(1).Find(".notice").Text())
		lines := strings.Split(notice, "\n")
		if len(lines) == 0 {
			return
		}

		timeLimit, memoryLimit := ParseCodeforcesLimits(lines[len(lines)-1])

		contest.Tasks = append(contest.Tasks, util.Task{
			TaskID:      taskID,
			URL:         fmt.Sprintf("https://codeforces.com%s", taskURL),
			TimeLimit:   timeLimit,
			MemoryLimit: memoryLimit,
		})
	})
	return contest, nil
}

func ParseCodeforcesLimits(limitsText string) (float64, int) {
	var timeLimit float64
	var memoryLimit int

	fmt.Sscanf(strings.TrimSpace(limitsText), "%f s, %d MB", &timeLimit, &memoryLimit)
	return timeLimit, memoryLimit
}
