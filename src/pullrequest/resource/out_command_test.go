package resource_test

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"pullrequest/resource"
	"pullrequest/resource/fake"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OutCommand", func() {
	var fakeGithub resource.Github
	var tempPath string
	var outCommand *resource.OutCommand

	BeforeEach(func() {
		fakeGithub = &fake.FGithub{}
		outCommand = resource.NewOutCommand(fakeGithub)
		tempPath = os.TempDir()
		os.Mkdir(path.Join(tempPath, "path"), os.ModePerm)

		fakeGithub.(*fake.FGithub).GetPRResult = &resource.Pull{
			Ref: "ref",
		}
		err := ioutil.WriteFile(path.Join(tempPath, "path", "pr_id"), []byte("ref"), 0644)
		Expect(err).ToNot(HaveOccurred())

		err = ioutil.WriteFile(path.Join(tempPath, "path", "pr_number"), []byte("12"), 0644)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		os.Remove(tempPath)
	})

	Context("when the resource is valid", func() {
		It("should return no error", func() {
			// Preparation
			outRequest := resource.OutRequest{
				OutParams: resource.OutParams{
					Status: "passing",
					Path:   "path",
				},
			}

			fakeRef := "fake-ref"
			fakeGithub.(*fake.FGithub).UpdatePRError = nil
			fakeGithub.(*fake.FGithub).UpdatePRResult = fakeRef

			// Execution
			outResponse, err := outCommand.Run(tempPath, outRequest)

			// Verification
			Expect(err).ToNot(HaveOccurred())
			Expect(outResponse.Version.Ref).To(Equal(fakeRef))
		})
	})

	Context("when the resource is invalid", func() {
		It("Should return error", func() {
			// Preparation
			outCommand := resource.NewOutCommand(fakeGithub)
			outRequest := resource.OutRequest{
				OutParams: resource.OutParams{
					Path: "path",
				},
			}

			fakeGithub.(*fake.FGithub).UpdatePRError = errors.New("An error occurred")

			// Execution
			_, err := outCommand.Run(tempPath, outRequest)

			// Verification
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("updating pr: An error occurred"))
		})
	})

	Context("when the folder is invalid", func() {
		It("Should return error", func() {
			// Preparation
			outCommand := resource.NewOutCommand(fakeGithub)
			outRequest := resource.OutRequest{
				OutParams: resource.OutParams{
					Status: "passing",
					Path:   "path",
				}}
			fakePath := "/fakepath"

			fakeGithub.(*fake.FGithub).UpdatePRError = errors.New("stat /fakepath/path: no such file or directory")

			// Execution
			_, err := outCommand.Run(fakePath, outRequest)

			// Verification
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("stat /fakepath/path: no such file or directory"))
		})
	})

	Context("When the pull request is outdated", func() {
		It("should return no error", func() {
			// Preparation
			outRequest := resource.OutRequest{
				OutParams: resource.OutParams{
					Status: "passing",
					Path:   "path",
				},
			}

			fakeGithub.(*fake.FGithub).GetPRResult = &resource.Pull{
				Ref: "latest-fake-ref",
			}

			fakeRef := "fake-ref"
			fakeGithub.(*fake.FGithub).UpdatePRError = nil
			fakeGithub.(*fake.FGithub).UpdatePRResult = fakeRef

			// Execution
			_, err := outCommand.Run(tempPath, outRequest)

			// Verification
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("PR is out of date"))
		})
	})
})
