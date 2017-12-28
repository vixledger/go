package xdr_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/vixledger/go/xdr"
)

var _ = Describe("xdr.AccountEntry#SignerSummary()", func() {
	const address = "GBMZ7CITNRMRQ3ZFRJP6TIFGVKKE4MH7KCCVRCISHEPSCHFOF4JA767I"
	var account AccountEntry

	BeforeEach(func() {
		account.AccountId.SetAddress(address)
	})

	It("adds the master signer when non-zero", func() {
		account.Thresholds[0] = 1
		summary := account.SignerSummary()
		Expect(summary).To(HaveKey(address))
		Expect(summary[address]).To(Equal(int32(1)))
	})

	It("doesn't have the master signer when zero", func() {
		account.Thresholds[0] = 0
		summary := account.SignerSummary()
		Expect(summary).ToNot(HaveKey(address))
	})

	It("includes every secondary signer", func() {
		account.Signers = []Signer{
			signer("GDMGDPXGPZ7UMKF6PPQULMENZQMMMUZITSNRNTVCMDLSM4FCHWJ5SIJY", 2),
			signer("GCHN24SEPBQGC3JJIB2YVTT2OVPO7FHOOH5ZMOIBKS4DANTR33U7S25Q", 4),
		}
		summary := account.SignerSummary()
		for _, signer := range account.Signers {
			addy := signer.Key.Address()
			Expect(summary).To(HaveKey(addy))
			Expect(summary[addy]).To(Equal(int32(signer.Weight)))
		}
	})
})

func signer(address string, weight int) (ret Signer) {

	ret.Key.SetAddress(address)
	ret.Weight = Uint32(weight)
	return
}
