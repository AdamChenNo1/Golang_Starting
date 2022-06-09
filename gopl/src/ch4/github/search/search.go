/*
 * File: \src\ch4\github\search\search.go                                      *
 * Project: gopl                                                               *
 * Created At: Tuesday, 2021/08/10 , 18:30:11                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/06/2 , 20:38:29                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package search

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
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    //Markdown 格式
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues 函数查询 Github 的 issue 跟踪接口
func SearchIssues(terms map[string][]string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms["q"], " "))
	sort := url.QueryEscape(strings.Join(terms["sort"], " "))
	res, err := http.Get(IssuesURL + "?q=" + q + "&sort=" + sort)
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
