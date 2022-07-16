package progress

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeserializeStat(t *testing.T) {
	source := `{
            "question_id": 1,
            "question__article__live": true,
            "question__article__slug": "two-sum",
            "question__title": "Two Sum",
            "question__title_slug": "two-sum",
            "question__hide": false,
            "total_acs": 2713958,
            "total_submitted": 6004485,
            "frontend_question_id": 1,
            "is_new_question": false
        }`
	var stat Stat
	json.Unmarshal([]byte(source), &stat)

	assert.Equal(t, "Two Sum", stat.Title)
	assert.Equal(t, "two-sum", stat.TitleSlug)
	assert.Equal(t, 1, stat.QuestionId)
}

func TestDeserializeProblem(t *testing.T) {
	source := `{
            "stat": {
                "question_id": 1,
                "question__article__live": true,
                "question__article__slug": "two-sum",
                "question__title": "Two Sum",
                "question__title_slug": "two-sum",
                "question__hide": false,
                "total_acs": 2713958,
                "total_submitted": 6004485,
                "frontend_question_id": 1,
                "is_new_question": false
            },
            "status": null,
            "difficulty": {
                "level": 1
            },
            "paid_only": false,
            "is_favor": false,
            "frequency": 0,
            "progress": 0
        }`

	var problem Problem
	json.Unmarshal([]byte(source), &problem)

	assert.Equal(t, "Two Sum", problem.Stat.Title)
	assert.Equal(t, 1, problem.Difficulty.Level)
	assert.False(t, problem.PaidOnly)
}
