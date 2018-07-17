package fake

import (
	"pullrequest/resource"
)

// FGithub is
type FGithub struct {
	ListPRResult []*resource.Pull
	ListPRError  error

	DownloadPRError error

	UpdatePRResult string
	UpdatePRError  error

	GetPRError  error
	GetPRResult *resource.Pull
}

// Get a single PR
func (fg *FGithub) GetPR(prNumber int) (*resource.Pull, error) {
	return fg.GetPRResult, fg.GetPRError
}

// ListPRs is
func (fg *FGithub) ListPRs() ([]*resource.Pull, error) {
	return fg.ListPRResult, fg.ListPRError
}

// DownloadPR is
func (fg *FGithub) DownloadPR(destDir string, prNumber int) error {
	return fg.DownloadPRError
}

// UpdatePR is
func (fg *FGithub) UpdatePR(sourceDir, status, path string) (string, error) {
	return fg.UpdatePRResult, fg.UpdatePRError
}
