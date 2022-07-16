package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"

	"github.com/go-git/go-git/v5"

	"github.com/gu18168/leetcode-status/pkg/progress"
)

func getProblems() (progress.Problems, error) {
	var result progress.Problems

	resp, err := http.Get("https://leetcode.com/api/problems/algorithms/")
	if err != nil {
		fmt.Println(fmt.Errorf("fetch Problems failed: %v", err))
		return result, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("read Problems failed: %v", err))
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(fmt.Errorf("unmarshal Problems failed: %v", err))
		return result, err
	}

	result.RetrainFree()
	return result, nil
}

func generate_report(repository, target string) {
	problems, err := getProblems()
	if err != nil {
		return
	}

	rep, err := git.PlainOpen(repository)
	if err != nil {
		fmt.Println(fmt.Errorf("open repository failed: %v", err))
		return
	}

	// sort problems by id
	sort.Slice(problems.Problems, func(i, j int) bool {
		return problems.Problems[i].Stat.QuestionId < problems.Problems[j].Stat.QuestionId
	})

	// make the target directory
	err = os.MkdirAll(target, os.ModePerm)
	if err != nil {
		fmt.Println(fmt.Errorf("mkdir %s failed: %v", target, err))
		return
	}

	// generate progress chart
	if err := progress.Draw(rep, &problems, filepath.Join(target, "progress.svg")); err != nil {
		fmt.Println(fmt.Errorf("draw progress failed: %v", err))
		return
	}

	// generate report
	if err := progress.Report(rep, &problems, "progress.svg", filepath.Join(target, "index.html")); err != nil {
		fmt.Println(fmt.Errorf("generate report failed: %v", err))
		return
	}
}

func main() {
	args := os.Args
	repository, target := args[1], args[2]

	generate_report(repository, target)
}
