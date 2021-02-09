package ui

import (
	"net/http"
	"path"

	"github.com/ddvk/rmfakecloud/internal/common"
	"github.com/ddvk/rmfakecloud/internal/config"
	"github.com/ddvk/rmfakecloud/internal/db"
	"github.com/ddvk/rmfakecloud/internal/webassets"
	"github.com/gin-gonic/gin"
)

type CodeGenerator interface {
	NewCode(string) (string, error)
}

/// ReactAppWrapper wrap some stuff
type ReactAppWrapper struct {
	fs            http.FileSystem
	prefix        string
	cfg           *config.Config
	userStorer    db.UserStorer
	codeConnector CodeGenerator
}

const indexReplacement = "/default"

func (w ReactAppWrapper) Open(filepath string) (http.File, error) {
	fullpath := filepath
	//index.html hack
	if filepath != indexReplacement {
		fullpath = path.Join(w.prefix, filepath)
	} else {
		fullpath = "/index.html"
	}
	f, err := w.fs.Open(fullpath)
	return f, err
}
func badReq(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": message})
}

// RegisterUI add the react ui
func New(cfg *config.Config, userStorer db.UserStorer, codeConnector common.CodeConnector) *ReactAppWrapper {
	staticWrapper := ReactAppWrapper{
		fs:            webassets.Assets,
		prefix:        "/static",
		cfg:           cfg,
		userStorer:    userStorer,
		codeConnector: codeConnector,
	}
	return &staticWrapper
}
