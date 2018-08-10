package main_test

import (
	"math/rand"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "ginexample"
)

var source = rand.NewSource(time.Now().UnixNano())

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[source.Int63()%int64(len(charset))]
	}
	return string(b)
}

var _ = Describe("ContructWebsealNames", func() {
	var message Message

	BeforeEach(func() {
		message = Message{
			Event: "teardown",
			Profile: &Profile{
				Branch:          RandString(rand.Intn(5)),
				BuildNumber:     string(rand.Intn(100)),
				ApplicationName: RandString(rand.Intn(10)),
				ServiceName:     RandString(rand.Intn(10)),
				Domain:          RandString(rand.Intn(5)) + "." + RandString(rand.Intn(10)),
			},
		}
	})

	Describe("given a Message", func() {
		Context("with random generated Profile", func() {
			It("should return constructed Webseal names", func() {
				Expect(ConstructWebsealNames(".",
					string(message.Profile.Branch), "-",
					string(message.Profile.BuildNumber), ".",
					string(message.Profile.ApplicationName), ".",
					string(message.Profile.ServiceName), ".",
					string(message.Profile.Domain),
				)).To(Equal([2]string{
					"AWS-1-1.webseald." + message.Profile.Branch +
						"-" + message.Profile.BuildNumber + "." + message.Profile.ApplicationName +
						"." + message.Profile.ServiceName + "." + message.Profile.Domain,
					"AWS-2-1.webseald." + message.Profile.Branch +
						"-" + message.Profile.BuildNumber + "." + message.Profile.ApplicationName +
						"." + message.Profile.ServiceName + "." + message.Profile.Domain}))
			})
		})
	})
})
