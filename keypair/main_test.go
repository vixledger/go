package keypair

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/types"
)

func TestBuild(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Package: github.com/vixledger/go/keypair")
}

var (
	address   = "GCNGTQWZVUMP6Z3GUQWDNZ55FBILPPIMRY5COVGKNXJNYNGVM7JQIEV5"
	seed      = "SBXNYNCP3FKBVR4GTZG3K7WOGG7EGH3A6WLL5CLZMGNGEPFXKKLTB3ZL"
	hint      = [4]byte{0xd5, 0x67, 0xd3, 0x4}
	message   = []byte("hello")
	signature = []byte{
		0x94, 0xbc, 0x89, 0x44, 0xf6, 0xd3, 0xff, 0x7f, 0xfa, 0xa0, 0xc1, 0xf4,
		0x81, 0x82, 0x62, 0x4f, 0x0e, 0x05, 0xc2, 0x04, 0xc4, 0x01, 0xec, 0xb4,
		0x1c, 0x93, 0xfe, 0xd1, 0x05, 0x98, 0xb5, 0x76, 0xba, 0xf3, 0xcb, 0xfe,
		0x98, 0xa5, 0xde, 0xac, 0x29, 0x51, 0xc2, 0x28, 0xaa, 0xe6, 0xcf, 0x86,
		0x6d, 0x9f, 0x65, 0xe3, 0xc5, 0xf4, 0x38, 0xee, 0x5a, 0x60, 0x7b, 0xe6,
		0x3f, 0xc9, 0xa8, 0x07,
	}
)

func ItBehavesLikeAKP(subject *KP) {

	// NOTE: subject will only be valid to dereference when inside am "It"
	// example.

	Describe("Address()", func() {
		It("returns the correct address", func() {
			Expect((*subject).Address()).To(Equal(address))
		})
	})

	Describe("Hint()", func() {
		It("returns the correct hint", func() {
			Expect((*subject).Hint()).To(Equal(hint))
		})
	})

	type VerifyCase struct {
		Message   []byte
		Signature []byte
		Case      types.GomegaMatcher
	}

	DescribeTable("Verify()",
		func(vc VerifyCase) {
			Expect((*subject).Verify(vc.Message, vc.Signature)).To(vc.Case)
		},
		Entry("correct", VerifyCase{message, signature, BeNil()}),
		Entry("empty signature", VerifyCase{message, []byte{}, Equal(ErrInvalidSignature)}),
		Entry("empty message", VerifyCase{[]byte{}, signature, Equal(ErrInvalidSignature)}),
		Entry("different message", VerifyCase{[]byte("diff"), signature, Equal(ErrInvalidSignature)}),
		Entry("malformed signature", VerifyCase{message, signature[0:10], Equal(ErrInvalidSignature)}),
	)
}

type ParseCase struct {
	Input    string
	TypeCase types.GomegaMatcher
	ErrCase  types.GomegaMatcher
}

var _ = DescribeTable("keypair.Parse()",
	func(c ParseCase) {
		kp, err := Parse(c.Input)

		Expect(kp).To(c.TypeCase)
		Expect(err).To(c.ErrCase)
	},

	Entry("a valid address", ParseCase{
		Input:    "GCNGTQWZVUMP6Z3GUQWDNZ55FBILPPIMRY5COVGKNXJNYNGVM7JQIEV5",
		TypeCase: BeAssignableToTypeOf(&FromAddress{}),
		ErrCase:  BeNil(),
	}),
	Entry("a corrupted address", ParseCase{
		Input:    "GCNGTQWZVUMP6Z3GUQWDNZ56FBILPPIMRY5COVGKNXJNYNGVM7JQIEV5",
		TypeCase: BeNil(),
		ErrCase:  HaveOccurred(),
	}),
	Entry("a valid seed", ParseCase{
		Input:    "SBXNYNCP3FKBVR4GTZG3K7WOGG7EGH3A6WLL5CLZMGNGEPFXKKLTB3ZL",
		TypeCase: BeAssignableToTypeOf(&Full{}),
		ErrCase:  BeNil(),
	}),
	Entry("a corrupted seed", ParseCase{
		Input:    "SBXNYNCP3FKBVR4GTZG3K7WOGG1EGH3A6WLL5CLZMGNGEPFXKKLTB3ZL",
		TypeCase: BeNil(),
		ErrCase:  HaveOccurred(),
	}),
	Entry("a blank string", ParseCase{
		Input:    "",
		TypeCase: BeNil(),
		ErrCase:  HaveOccurred(),
	}),
)
