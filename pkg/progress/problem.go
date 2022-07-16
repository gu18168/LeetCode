package progress

import "fmt"

type Stat struct {
	Title      string `json:"question__title"`
	TitleSlug  string `json:"question__title_slug"`
	QuestionId int    `json:"frontend_question_id"`
}

type Difficulty struct {
	Level int `json:"level"`
}

type Problem struct {
	Stat       Stat       `json:"stat"`
	Difficulty Difficulty `json:"difficulty"`
	PaidOnly   bool       `json:"paid_only"`
}

func (p *Problem) GetId() string {
	return fmt.Sprintf("%04d-%s", p.Stat.QuestionId, p.Stat.TitleSlug)
}

type Problems struct {
	Problems []Problem `json:"stat_status_pairs"`
}

func (ps *Problems) RetrainFree() {
	free := make([]Problem, 0)
	for _, p := range ps.Problems {
		if !p.PaidOnly {
			free = append(free, p)
		}
	}
	ps.Problems = free
}
