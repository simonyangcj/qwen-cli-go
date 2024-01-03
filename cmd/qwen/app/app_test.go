package app

import (
	"path"
	"qwen-cli/cmd/qwen/app/option"
	api "qwen-cli/pkg/api"
	conversationLib "qwen-cli/pkg/conversation"
	"qwen-cli/pkg/util"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("qwen-cli-go test", func() {
	conversationDir := "/tmp/conversation"
	BeforeEach(func() {
		err := util.CreateDirIfNotExists(conversationDir)
		Expect(err).To(BeNil())
	})

	Describe("create,delete,list,show conversation", func() {
		It("create conversation no error", func() {
			err := createConversationRun(conversationDir)(nil, []string{"c1"})
			Expect(err).To(BeNil())
			exists := util.PathExist(path.Join(conversationDir, "c1"))
			Expect(exists).To(Equal(true))
			exists = util.PathExist(path.Join(conversationDir, "c1", "conversation.json"))
			Expect(exists).To(Equal(true))
		})
		It("delete conversation no error", func() {
			err := createConversationRun(conversationDir)(nil, []string{"c1"})
			Expect(err).To(BeNil())
			exists := util.PathExist(path.Join(conversationDir, "c1"))
			Expect(exists).To(Equal(true))
			err = deleteConversationRun(conversationDir)(nil, []string{"c1"})
			Expect(err).To(BeNil())
			exists = util.PathExist(path.Join(conversationDir, "c1"))
			Expect(exists).To(Equal(false))
		})
		It("show conversation no error", func() {
			err := createConversationRun(conversationDir)(nil, []string{"c1"})
			Expect(err).To(BeNil())
			exists := util.PathExist(path.Join(conversationDir, "c1"))
			Expect(exists).To(Equal(true))
			conversationLib.WriteToConversationFile(conversationDir, "c1", []api.Message{{Content: "hello", Role: "user"}})
			Expect(err).To(BeNil())
			err = showConversationRun(conversationDir)(nil, []string{"c1"})
			Expect(err).To(BeNil())
		})
		AfterEach(func() {
			err := util.DeleteDirs(conversationDir)
			Expect(err).To(BeNil())
		})
	})

	Describe("prompt", func() {
		It("preCheck test should be no error", func() {
			err := preCheckArgs(&option.PromptOptions{Model: "qwen-turbo"})
			Expect(err).To(BeNil())
		})
		It("preCheck test should be error", func() {
			err := preCheckArgs(&option.PromptOptions{Model: "qwen-turbo1"})
			Expect(err).ToNot(BeNil())
		})
	})
})
