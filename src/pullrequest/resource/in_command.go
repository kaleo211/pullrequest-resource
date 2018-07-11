package resource

import (
	"errors"
	"fmt"
	"strconv"
)

// InCommand is
type InCommand struct {
	github Github
}

// NewInCommand is
func NewInCommand(g Github) *InCommand {
	return &InCommand{g}
}

// Run is
func (ic *InCommand) Run(destDir string, req InRequest) (InResponse, error) {

	pr, err := strconv.Atoi(req.Version.PR)
	if err != nil {
		return InResponse{}, fmt.Errorf("converting pr number: %+v", err)
	}

	pull, err := ic.github.DownloadPR(destDir, pr)
	if err != nil {
		return InResponse{}, fmt.Errorf("downloading pr: %+v", err)
	}

	if pull.Ref != req.Version.Ref {
		return InResponse{}, errors.New("pr ref is not valid")
	}

	return InResponse{
		Version: req.Version,
	}, nil
}
