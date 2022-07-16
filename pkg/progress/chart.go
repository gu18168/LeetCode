package progress

import (
	"fmt"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

type Progress struct {
	time time.Time
	hit  uint
}

func getHits(wg *sync.WaitGroup, tree *object.Tree, date time.Time, ch chan<- Progress) {
	defer wg.Done()

	cache := make(map[string]struct{})
	List(tree, func(solution Solution) {
		cache[solution.ProblemId] = struct{}{}
	})

	hits := len(cache)
	ch <- Progress{date.UTC(), uint(hits)}
}

func getPert(hit, total uint) float64 {
	return float64(hit*100) / float64(total)
}

type pertTicks struct{}

func (pertTicks) Ticks(min, max float64) []plot.Tick {
	tks := plot.DefaultTicks{}.Ticks(min, max)
	for i, t := range tks {
		if t.Value == 0 {
			continue
		}

		tks[i].Label = fmt.Sprintf("%d%%", int(t.Value))
	}

	return tks
}

type dateTicks struct{}

func (dateTicks) Ticks(min, max float64) []plot.Tick {
	tks := plot.DefaultTicks{}.Ticks(min, max)
	for i, t := range tks {
		tks[i].Label = time.Unix(int64(t.Value), 0).Format("2006-01-02")
	}

	return tks
}

func drawChart(progress []Progress, total uint, output string) error {
	const zoom = 1
	const width, height = 987, 610
	const titleFont = 16.0
	const tMargin = 12

	length := len(progress)
	latestDate, latestHit := progress[length-1].time, progress[length-1].hit

	p := plot.New()
	p.Title.Text = fmt.Sprintf("Total: %d, Done: %d, Remaining: %d, Progress: %.2f, Updated: %s",
		total, latestHit, total-latestHit, getPert(latestHit, total), latestDate.Format("2006-01-02"))
	p.Title.Padding = tMargin * zoom
	p.Title.TextStyle.Font.Size = titleFont * 1.24 * font.Length(zoom)

	p.Y.Min, p.Y.Max = 0, 100
	p.Y.Tick.Marker = pertTicks{}

	p.X.Min, p.X.Max = float64(progress[0].time.Unix()), float64(latestDate.Unix())
	p.X.Tick.Marker = dateTicks{}

	p.Add(plotter.NewGrid())

	points := make(plotter.XYs, length)
	for i, prog := range progress {
		points[i].X = float64(prog.time.Unix())
		points[i].Y = getPert(prog.hit, total)
	}
	line, err := plotter.NewLine(points)
	if err != nil {
		return err
	}
	line.Color = plotutil.Color(3)

	p.Add(line)
	wc, err := p.WriterTo(width*zoom, height*zoom, "svg")
	if err != nil {
		return err
	}

	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	wc.WriteTo(f)

	return nil
}

func Draw(repository *git.Repository, problems *Problems, output string) error {
	ref, err := repository.Head()
	if err != nil {
		return err
	}

	commitIter, err := repository.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return err
	}

	progress := make([]Progress, 0)
	ch := make(chan Progress, 10)

	// count the date and hit corresponding to the commit
	go func(ch chan Progress) {
		defer close(ch)

		var wg sync.WaitGroup
		commitIter.ForEach(func(commit *object.Commit) error {
			wg.Add(1)

			date := commit.Author.When
			tree, err := commit.Tree()
			if err != nil {
				return nil
			}

			go getHits(&wg, tree, date, ch)

			return nil
		})

		wg.Wait()
	}(ch)
	
	for prog := range ch {
		progress = append(progress, prog)
	}

	// sort daily hits
	sort.Slice(progress, func(i, j int) bool {
		return progress[i].time.Before(progress[j].time)
	})

	drawChart(progress, uint(len(problems.Problems)), output)

	return nil
}
