package fake

import "pullrequest/resource"

// FGithub is
type FGithub struct {
	ListPRResult []*resource.Pull
	ListPRError  error

	DownloadPRError  error
	DownloadPRResult *resource.Pull

	UpdatePRResult string
	UpdatePRError  error
}

// ListPRs is
func (fg *FGithub) ListPRs() ([]*resource.Pull, error) {
	return fg.ListPRResult, fg.ListPRError
}

// DownloadPR is
func (fg *FGithub) DownloadPR(destDir string, prNumber int) (*resource.Pull, error) {
	return fg.DownloadPRResult, fg.DownloadPRError
}

// UpdatePR is
func (fg *FGithub) UpdatePR(sourceDir, status, path string) (string, error) {
	return fg.UpdatePRResult, fg.UpdatePRError
}
