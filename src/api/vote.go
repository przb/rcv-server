package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

type SingleRank struct {
	Name string `json:"name"`
	Rank int    `json:"rank"`
}

type voteSubmission struct {
	// doing weird formatting to make the json better
	Rankings []SingleRank `json:"rankings"`
}

type voteOptions struct {
	Names []string `json:"options"`
}

func VoteSubmit(ctx *gin.Context) {
	var in voteSubmission
	ctx.BindJSON(&in)

	// TODO
}

func VoteGetOptions(ctx *gin.Context) {
	var vo voteOptions
	id := ctx.Query("id")
	val, ok := db[id]
	options := val.Options
	// technically this could just return _n, since the option map is mapped by name, but this would make it a little easier to change the key if we so desire
	vo.Names = funk.Map(options, func(_n string, o Option) string { return o.Name }).([]string)
	if ok {
		ctx.IndentedJSON(http.StatusOK, vo)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"details": "Id not found"})
	}
}
