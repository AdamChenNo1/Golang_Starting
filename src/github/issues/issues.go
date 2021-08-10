/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-07-15 17:48:55
 * @LastEditors: Elon C
 * @LastEditTime: 2021-07-15 21:33:03
 * @FilePath: \Golang_Starting\github\issues\issues.go
 */
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:total_count`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:html_url`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:created_at`
	Body      string    //Markdown 格式
}

type User struct {
	Login   string
	HTMLURL string `json:html_url`
}

// SearchIssues 函数查询 Github 的 issue 跟踪接口
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	res, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", res.Status)
	}

	var result IssuesSearchResult

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		res.Body.Close()
		return nil, err
	}

	res.Body.Close()
	return &result, nil
}
