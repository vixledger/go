package xdr_test

import (
	. "github.com/vixledger/go/xdr"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("xdr.AccountId#Address()", func() {
	It("returns an empty string when account id is nil", func() {
		addy := (*AccountId)(nil).Address()
		Expect(addy).To(Equal(""))
	})

	It("returns a strkey string when account id is valid", func() {
		var aid AccountId
		aid.SetAddress("GBMZ7CITNRMRQ3ZFRJP6TIFGVKKE4MH7KCCVRCISHEPSCHFOF4JA767I")
		addy := aid.Address()
		Expect(addy).To(Equal("GBMZ7CITNRMRQ3ZFRJP6TIFGVKKE4MH7KCCVRCISHEPSCHFOF4JA767I"))
	})
})

var _ = Describe("xdr.AccountId#Equals()", func() {
	It("returns true when the account ids have equivalent values", func() {
		var l, r AccountId
		l.SetAddress("GBMZ7CITNRMRQ3ZFRJP6TIFGVKKE4MH7KCCVRCISHEPSCHFOF4JA767I")
		r.SetAddress("GBMZ7CITNRMRQ3ZFRJP6TIFGVKKE4MH7KCCVRCISHEPSCHFOF4JA767I")
		Expect(l.Equals(r)).To(BeTrue())
	})

	It("returns false when the account ids have different values", func() {
		var l, r AccountId
		l.SetAddress("GBMZ7CITNRMRQ3ZFRJP6TIFGVKKE4MH7KCCVRCISHEPSCHFOF4JA767I")
		r.SetAddress("GBTBXQEVDNVUEESCTPUT3CHJDVNG44EMPMBELH5F7H3YPHXPZXOTEWB4")
		Expect(l.Equals(r)).To(BeFalse())
	})
})

var _ = Describe("xdr.AccountId#LedgerKey()", func() {
	It("works", func() {
		var aid AccountId
		aid.SetAddress("GBMZ7CITNRMRQ3ZFRJP6TIFGVKKE4MH7KCCVRCISHEPSCHFOF4JA767I")

		key := aid.LedgerKey()
		packed := key.MustAccount().AccountId
		Expect(packed.Equals(aid)).To(BeTrue())
	})
})
