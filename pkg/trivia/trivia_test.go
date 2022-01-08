package trivia

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPI_GetQuestions(t *testing.T) {
	// Start a local HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		rw.Write([]byte(sample_response))
	}))
	// Close the server when test finishes
	defer ts.Close()

	// Use Client & URL from our local test server
	api := API{ts.Client(), DefaultOption, ts.URL}
	got, err := api.GetQuestions(context.Background())

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(test_1_expected, got)
}

func TestPrettyPrint(t *testing.T) {
	got := PrettyPrint(sample_question)
	assert.Equal(t, test_2_expected, got)
}

var sample_question = []*Question{
	{
		Category:      "Entertainment: Music",
		Type:          "multiple",
		Difficulty:    "medium",
		Question:      "Which country is singer Kyary Pamyu Pamyu from?",
		CorrectAnswer: "Japan",
		IncorrectAnswers: []string{
			"South Korea",
			"China",
			"Vietnam",
		},
	},
}

var sample_response = `
{
	"response_code": 0,
	"results": [
		{
			"category": "Entertainment: Music",
        	"type": "multiple",
        	"difficulty": "medium",
        	"question": "Which country is singer Kyary Pamyu Pamyu from?",
        	"correct_answer": "Japan",
        	"incorrect_answers": [
            	"South Korea",
            	"China",
            	"Vietnam"
        	]
    	}
	]
}`

var test_1_expected = sample_question

var test_2_expected = `[
    {
        "category": "Entertainment: Music",
        "type": "multiple",
        "difficulty": "medium",
        "question": "Which country is singer Kyary Pamyu Pamyu from?",
        "correct_answer": "Japan",
        "incorrect_answers": [
            "South Korea",
            "China",
            "Vietnam"
        ]
    }
]`
