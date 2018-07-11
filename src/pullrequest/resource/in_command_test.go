package resource_test

import (
	"errors"
	"os"
	"path"
	"pullrequest/resource"
	"pullrequest/resource/fake"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("InCommand", func() {
	var fakeGithub *fake.FGithub
	var fakeDestDir string

	BeforeEach(func() {
		fakeGithub = &fake.FGithub{}
		fakeDestDir = path.Join(os.TempDir(), "fake-dir")
	})

	AfterEach(func() {
		os.Remove(fakeDestDir)
	})

	Context("when version is valid", func() {
		It("should return no error", func() {
			// Preparation
			fakeRef := "fake-ref"

			fakeGithub.DownloadPRResult = &resource.Pull{
				Ref: fakeRef,
			}
			inCommand := resource.NewInCommand(fakeGithub)

			inRequest := resource.InRequest{
				Version: resource.Version{
					Ref: fakeRef,
					PR:  "89",
				},
			}

			// Execution
			inResponse, err := inCommand.Run(fakeDestDir, inRequest)

			// Verification
			Expect(err).ToNot(HaveOccurred())
			Expect(inResponse.Version.Ref).To(Equal(fakeRef))
			Expect(inResponse.Version.PR).To(Equal("89"))
		})
	})

	Context("when destDir is not valid", func() {
		It("should return error", func() {
			fakeGithub := &fake.FGithub{}
			inCommand := resource.NewInCommand(fakeGithub)

			inRequest := resource.InRequest{}

			_, err := inCommand.Run("/fake/not/exist/dir", inRequest)
			Expect(err).To(MatchError("making dest dir: mkdir /fake/not/exist/dir: no such file or directory"))
		})
	})

	Context("when download failed", func() {
		It("should return error", func() {
			fakeGithub.DownloadPRError = errors.New("fake-download-pr-error")
			inCommand := resource.NewInCommand(fakeGithub)

			inRequest := resource.InRequest{
				Version: resource.Version{
					Ref: "fake-ref",
					PR:  "100",
				},
			}

			_, err := inCommand.Run(fakeDestDir, inRequest)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError("downloading pr: fake-download-pr-error"))
		})
	})

	Context("when pr ref is not valid", func() {
		It("should return error", func() {
			fakeGithub.DownloadPRResult = &resource.Pull{
				Ref: "fake-ref-001",
			}
			inCommand := resource.NewInCommand(fakeGithub)
			inRequest := resource.InRequest{
				Version: resource.Version{
					Ref: "fake-ref",
					PR:  "100",
				},
			}

			_, err := inCommand.Run(fakeDestDir, inRequest)
			Expect(err).To(MatchError("pr ref is not valid"))
		})
	})
})
