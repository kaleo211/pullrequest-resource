package resource

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

// OutCommand is
type OutCommand struct {
	github Github
}

// NewOutCommand is
func NewOutCommand(g Github) *OutCommand {
	return &OutCommand{g}
}

// Run is
func (oc *OutCommand) Run(sourceDir string, req OutRequest) (OutResponse, error) {
	if _, err := os.Stat(path.Join(sourceDir, req.Path)); err != nil {
		return OutResponse{}, err
	}

	id, err := ioutil.ReadFile(path.Join(sourceDir, req.Path, "pr_id"))
	if err != nil {
		return OutResponse{}, err
	}
	num, err := ioutil.ReadFile(path.Join(sourceDir, req.Path, "pr_number"))
	if err != nil {
		return OutResponse{}, err
	}

	number, err := strconv.Atoi(string(num))
	if err != nil {
		return OutResponse{}, fmt.Errorf("converting PR number: %+v", err)
	}

	pull, err := oc.github.GetPR(number)
	if err != nil {
		return OutResponse{}, fmt.Errorf("getting PR: %+v", err)
	}

	if string(id) != pull.Ref {
		return OutResponse{}, errors.New("PR is out of date")
	}

	ref, err := oc.github.UpdatePR(sourceDir, req.OutParams.Status, req.OutParams.Path)
	if err != nil {
		return OutResponse{}, fmt.Errorf("updating pr: %+v", err)
	}

	response := OutResponse{
		Version: Version{
			Ref: ref,
		},
	}
	return response, err
}
