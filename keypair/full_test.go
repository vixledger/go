package keypair

import (
	"encoding/hex"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("keypair.Full", func() {
	var subject KP

	JustBeforeEach(func() {
		subject = &Full{seed}
	})

	ItBehavesLikeAKP(&subject)

	type SignCase struct {
		Message   string
		Signature string
	}

	DescribeTable("Sign()",
		func(c SignCase) {
			sig, err := subject.Sign([]byte(c.Message))
			actual := hex.EncodeToString(sig)

			Expect(actual).To(Equal(c.Signature))
			Expect(err).To(BeNil())
		},

		Entry("hello", SignCase{
			"hello",
			"94bc8944f6d3ff7ffaa0c1f48182624f0e05c204c401ecb41c93fed10598b576baf3cbfe98a5deac2951c228aae6cf866d9f65e3c5f438ee5a607be63fc9a807",
		}),
		Entry("this is a message", SignCase{
			"this is a message",
			"c8f7b7926faf57afb99185c72f9e133262726fd0f74055d58ecaa7ce7860a7e0cd58bba7abd39290948be939b968e43b0f453a9360ea7089e0748598d521190a",
		}),
	)

	Describe("SignDecorated()", func() {
		It("returns the correct xdr struct", func() {
			sig, err := subject.SignDecorated(message)
			Expect(err).To(BeNil())
			Expect(sig.Hint).To(BeEquivalentTo(hint))
			Expect(sig.Signature).To(BeEquivalentTo(signature))
		})
	})

})
