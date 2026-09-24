package new

import (
	"fmt"
	"ojx/bin/util"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func NewYukicoder(contestID string) (*util.Contest, error) {
	url := fmt.Sprintf("https://yukicoder.me/contests/%s", contestID)

	resp, err := NewRequest("GET", url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	contest := &util.Contest{
		Platform:  "yukicoder",
		ContestID: contestID,
	}

	doc.Find("table tbody tr").Each(func(_ int, tr *goquery.Selection) {
		cols := tr.Find("td")

		if cols.Length() < 3 {
			return
		}

		taskID := strings.ToLower(strings.TrimSpace(cols.Eq(0).Text()))
		taskURL, exists := cols.Eq(2).Find("a").Attr("href")
		if !exists {
			return
		}

		task, err := LoadYukicoderTask(taskID, taskURL)
		if err != nil {
			return
		}

		contest.Tasks = append(contest.Tasks, task)
	})

	if len(contest.Tasks) == 0 {
		return nil, fmt.Errorf(
			"no tasks found",
		)
	}

	return contest, nil
}

func LoadYukicoderTask(taskID, taskURL string) (util.Task, error) {
	resp, err := NewRequest("GET", fmt.Sprintf("https://yukicoder.me%s", taskURL))
	if err != nil {
		return util.Task{}, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return util.Task{}, err
	}

	header := strings.TrimSpace(doc.Find(".problem-header-main").First().Text())

	return util.Task{
		TaskID:      taskID,
		URL:         fmt.Sprintf("https://yukicoder.me%s", taskURL),
		TimeLimit:   parseYukicoderTimeLimit(header),
		MemoryLimit: parseYukicoderMemoryLimit(header),
	}, nil
}

func parseYukicoderTimeLimit(header string) float64 {
	re := regexp.MustCompile(`実行時間制限\s*:\s*1ケース\s*([0-9.]+)秒`)

	match := re.FindStringSubmatch(header)
	if len(match) < 2 {
		return 0
	}

	f, err := strconv.ParseFloat(match[1], 64)

	if err != nil {
		return 0
	}
	return f
}

func parseYukicoderMemoryLimit(header string) int {
	re := regexp.MustCompile(`メモリ制限\s*:\s*([0-9]+)\s*MB`)

	match := re.FindStringSubmatch(header)
	if len(match) < 2 {
		return 0
	}

	f, err := strconv.Atoi(match[1])

	if err != nil {
		return 0
	}
	return f
}
