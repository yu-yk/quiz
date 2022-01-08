package trivia

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

type Response struct {
	ResponseCode int         `json:"response_code,omitempty"`
	Results      []*Question `json:"results,omitempty"`
}

type Question struct {
	Category         string   `json:"category,omitempty"`
	Type             string   `json:"type,omitempty"`
	Difficulty       string   `json:"difficulty,omitempty"`
	Question         string   `json:"question,omitempty"`
	CorrectAnswer    string   `json:"correct_answer,omitempty"`
	IncorrectAnswers []string `json:"incorrect_answers,omitempty"`
}

type Options struct {
	Amount int
	Type   string
	// Category   string
	// Difficulty string
	// Encoding   string
}

var DefaultOption = Options{
	Amount: 1,
	Type:   "",
}

type API struct {
	Client  *http.Client
	Opt     Options
	baseURL string
}

func NewAPI() *API {
	return &API{
		Client:  http.DefaultClient,
		Opt:     DefaultOption,
		baseURL: "https://opentdb.com/api.php",
	}
}

func (api *API) GetQuestions(ctx context.Context) ([]*Question, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, api.baseURL, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	q := req.URL.Query()
	q.Add("amount", strconv.Itoa(api.Opt.Amount))
	q.Add("type", api.Opt.Type)

	req.URL.RawQuery = q.Encode()

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()

	var r Response
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, errors.WithStack(err)
	}

	return r.Results, nil
}

func PrettyPrint(qq []*Question) string {
	s, _ := json.MarshalIndent(qq, "", "    ")
	return string(s)
}
