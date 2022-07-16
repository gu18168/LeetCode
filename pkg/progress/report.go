package progress

import (
	"html/template"
	"io"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func makeSolution(tree *object.Tree) map[string][]Solution {
	result := make(map[string][]Solution)
	List(tree, func(solution Solution) {
		haveSolutions := result[solution.ProblemId]
		result[solution.ProblemId] = append(haveSolutions, solution)
	})

	return result
}

type Elements struct {
	Chart     string
	Solutions map[string][]Solution
	Problems  []Problem
}

func writeHTML(w io.Writer, problems *Problems, tree *object.Tree, chart string) error {
	elements := Elements{
		Chart:     chart,
		Solutions: makeSolution(tree),
		Problems:  problems.Problems,
	}

	tmpl_path, _ := filepath.Abs("pkg/progress/template.html")
	tmpl := template.Must(template.ParseFiles(tmpl_path))
	tmpl.Execute(w, elements)

	return nil
}

func Report(repository *git.Repository, problems *Problems, chart, output string) error {
	f, err := os.Create(output)
	if err != nil {
		return err
	}

	ref, err := repository.Head()
	if err != nil {
		return err
	}

	lastCommit, err := repository.CommitObject(ref.Hash())
	if err != nil {
		return err
	}

	tree, err := lastCommit.Tree()
	if err != nil {
		return err
	}

	writeHTML(f, problems, tree, chart)
	return nil
}
