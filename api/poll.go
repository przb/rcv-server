package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Used Internally
var db = make(map[string]Poll) // map id to poll

type Poll struct {
	Options map[string]Option `json:"options"` // map name to option. this implies names must be unique
	Id      string            `json:"id"`
}

type Option struct {
	Rank     int    `json:"rank"`
	Name     string `json:"name"`
	NumVotes int    `json:"num_votes"`
}

// Used when making requests
type PollInput struct {
	Options []string `json:"options"`
}

type PollOutput struct {
	Options []Option `json:"options"`
	Id      string   `json:"id"`
}

func pollInToPoll(i PollInput) (p Poll) {
	p.Id = uuid.NewString()
  p.Options = make(map[string]Option)
	for i, on := range i.Options {
		o := Option{
			Rank:     i,
			Name:     on,
			NumVotes: 0,
		}
    fmt.Println("appended option")
		p.Options[on] = o
	}
	return
}

func pollToOutPoll(p Poll) (o PollOutput) {
	o.Id = p.Id
	for _, option := range p.Options {
		o.Options = append(o.Options, option)
	}
	return
}

func PollCreateHandler(ctx *gin.Context) {
	var in PollInput

	ctx.BindJSON(&in)
	p := pollInToPoll(in)
  db[p.Id] = p

	out := pollToOutPoll(p)
	ctx.IndentedJSON(http.StatusOK, out)
	return
}
