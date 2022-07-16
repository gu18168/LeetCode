package progress

import (
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5/plumbing/filemode"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Solution struct {
	ProblemId    string
	SolutionId   string
	SolutionRoot string
	SolutionFile string
}

func List(tree *object.Tree, f func(solution Solution)) {
	for _, entry := range tree.Entries {
		if entry.Mode == filemode.Dir && entry.Name == "src" {
			sourceTree, _ := tree.Tree(entry.Name)

			// problem_{problem_id}
			for _, source := range sourceTree.Entries {
				if source.Mode != filemode.Dir {
					continue
				}

				problemId := strings.TrimPrefix(source.Name, "problem_")
				solutionTree, _ := sourceTree.Tree(source.Name)

				// solution.go file
				for _, solution := range solutionTree.Entries {
					solutionId := strings.TrimSuffix(solution.Name, ".go")
					f(Solution{
						ProblemId:    strings.ReplaceAll(problemId, "_", "-"),
						SolutionId:   solutionId,
						SolutionRoot: filepath.Join(entry.Name, source.Name),
						SolutionFile: solution.Name,
					})
				}
			}
		}
	}
}
