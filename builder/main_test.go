package builder

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBuild(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Package: github.com/vixledger/go/build")
}

// ExampleTransactionBuilder creates and signs a simple transaction, and then
// encodes it into a base64 string capable of being submitted to vixal-core.
//
// It uses the transaction builder system
func ExampleTransactionBuilder() {
	seed := "SDL6ZQL52IEX34E574CGWKYX5MVPILN7F44QESY7VLWLKUSXZNBVERDY"
	tx := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		Payment(
			Destination{"GCGZMI5I2SEV2DEF3477UI3PAC2UJUZECK36KK4XFY6C2YPTSCNKRDA6"},
			NativeAmount{"50"},
		),
	)

	txe := tx.Sign(seed)
	txeB64, _ := txe.Base64()

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAALy+aX1ueJ5Q7EcqD543OWZszFYHG2zJdImOryX8vwQsAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAAjZYjqNSJXQyF3z/6I28AtUTTJBK35SuXLjwtYfOQmqgAAAAAAAAAAB3NZQAAAAAAAAAAAfy/BCwAAABAynhW6w86xIFEqFT1h6zzBYb86K9VG2shmIjewByGbenV1hHR6d0rA2XxtwzDCTAkvDQ4e9jqIQd+NLvU1jLVCA==
}

// ExamplePathPayment creates and signs a simple transaction with PathPayment operation, and then
// encodes it into a base64 string capable of being submitted to vixal-core.
func ExamplePathPayment() {
	seed := "SDL6ZQL52IEX34E574CGWKYX5MVPILN7F44QESY7VLWLKUSXZNBVERDY"
	tx := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		Payment(
			Destination{"GAI5G7FANU5M3SPUB2RQBOKCACITUOHOIGAKLG6F6MJPIUB3N2ICGSGB"},
			CreditAmount{"USD", "GCGZMI5I2SEV2DEF3477UI3PAC2UJUZECK36KK4XFY6C2YPTSCNKRDA6", "50"},
			PayWith(CreditAsset("EUR", "GBGRQD3EWQZBDBEWFRP5IABAD2G2BYKCROXP2GDWCOCTTHC5DJCKXRDQ"), "100").
				Through(Asset{Native: true}).
				Through(CreditAsset("BTC", "GCWZWELRUSWEWAW6THI46HVZQITUMLVIE67PO6MYMW6AHWUF6ZM3SERL")),
		),
	)

	txe := tx.Sign(seed)
	txeB64, _ := txe.Base64()

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAALy+aX1ueJ5Q7EcqD543OWZszFYHG2zJdImOryX8vwQsAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAIAAAABRVVSAAAAAABNGA9ktDIRhJYsX9QAIB6NoOFCi679GHYThTmcXRpEqwAAAAA7msoAAAAAABHTfKBtOs3J9A6jALlCAJE6OO5BgKWbxfMS9FA7bpAjAAAAAVVTRAAAAAAAjZYjqNSJXQyF3z/6I28AtUTTJBK35SuXLjwtYfOQmqgAAAAAHc1lAAAAAAIAAAAAAAAAAUJUQwAAAAAArZsRcaSsSwLemdHPHrmCJ0YuqCe+93mYZbwD2oX2WbkAAAAAAAAAAfy/BCwAAABANz7oKnHaQo0S6czk0XDkWw/DfCcp+38Pen4hfbFG19j8SbLOuTc+jQiE1IwlhY6mH9thHluR63F9SQIGV0jSAg==
}

// ExampleSetOptions creates and signs a simple transaction with SetOptions operation, and then
// encodes it into a base64 string capable of being submitted to vixal-core.
func ExampleSetOptions() {
	seed := "SDL6ZQL52IEX34E574CGWKYX5MVPILN7F44QESY7VLWLKUSXZNBVERDY"
	tx := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		SetOptions(
			InflationDest("GCRGX2ZL6U7DE6IVKPK3FO7O5DILNMLVLRGGO4J5OCMN2SZ6TAJD7F7L"),
			SetAuthRequired(),
			SetAuthRevocable(),
			SetAuthImmutable(),
			ClearAuthRequired(),
			ClearAuthRevocable(),
			ClearAuthImmutable(),
			MasterWeight(1),
			SetThresholds(2, 3, 4),
			HomeDomain("ngboss.com"),
			AddSigner("GDDIOEMAXD5SWVVUZCZFBSYXD3ITL2ROQ3GGI3WMRDD6G5VTC4XQ7DTT", 5),
		),
	)

	txe := tx.Sign(seed)
	txeB64, _ := txe.Base64()

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAALy+aX1ueJ5Q7EcqD543OWZszFYHG2zJdImOryX8vwQsAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAUAAAABAAAAAKJr6yv1PjJ5FVPVsrvu6NC2sXVcTGdxPXCY3Us+mBI/AAAAAQAAAAcAAAABAAAABwAAAAEAAAABAAAAAQAAAAIAAAABAAAAAwAAAAEAAAAEAAAAAQAAAApuZ2Jvc3MuY29tAAAAAAABAAAAAMaHEYC4+ytWtMiyUMsXHtE16i6GzGRuzIjH43azFy8PAAAABQAAAAAAAAAB/L8ELAAAAEBRAWIkipgpakN0OZkynQN2MVmz59mV3ZTQQxQdyP9siSNtMPeBe5FyaDXKV7Pam48tPl2FDm68HI82LFeuy+oO
}

// ExampleSetOptions_manyOperations creates and signs a simple transaction with many SetOptions operations, and then
// encodes it into a base64 string capable of being submitted to vixal-core.
func ExampleSetOptions_manyOperations() {
	seed := "SDL6ZQL52IEX34E574CGWKYX5MVPILN7F44QESY7VLWLKUSXZNBVERDY"
	tx := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		InflationDest("GCRGX2ZL6U7DE6IVKPK3FO7O5DILNMLVLRGGO4J5OCMN2SZ6TAJD7F7L"),
		SetAuthRequired(),
		SetAuthRevocable(),
		SetAuthImmutable(),
		ClearAuthRequired(),
		ClearAuthRevocable(),
		ClearAuthImmutable(),
		MasterWeight(1),
		SetThresholds(2, 3, 4),
		HomeDomain("ngboss.com"),
		RemoveSigner("GDDIOEMAXD5SWVVUZCZFBSYXD3ITL2ROQ3GGI3WMRDD6G5VTC4XQ7DTT"),
	)

	txe := tx.Sign(seed)
	txeB64, _ := txe.Base64()

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAALy+aX1ueJ5Q7EcqD543OWZszFYHG2zJdImOryX8vwQsAAAETAAAAAAAAAABAAAAAAAAAAAAAAALAAAAAAAAAAUAAAABAAAAAKJr6yv1PjJ5FVPVsrvu6NC2sXVcTGdxPXCY3Us+mBI/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAQAAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAQAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAQAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAABAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAABAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAABAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAEAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAABAAAAAgAAAAEAAAADAAAAAQAAAAQAAAAAAAAAAAAAAAAAAAAFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAAKbmdib3NzLmNvbQAAAAAAAAAAAAAAAAAFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAAAAMaHEYC4+ytWtMiyUMsXHtE16i6GzGRuzIjH43azFy8PAAAAAAAAAAAAAAAB/L8ELAAAAECMuRKUPbiK545kwCbf0pEzTbZ4S/1zbMBrzxqN6+oVxM1C+y9eioRquMZLwCzEEKw1bdk+xF/VcP6OYto1iG8I
}

// ExampleChangeTrust creates and signs a simple transaction with ChangeTrust operation, and then
// encodes it into a base64 string capable of being submitted to vixal-core.
func ExampleChangeTrust() {
	seed := "SDL6ZQL52IEX34E574CGWKYX5MVPILN7F44QESY7VLWLKUSXZNBVERDY"
	tx := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		Trust("USD", "GCGZMI5I2SEV2DEF3477UI3PAC2UJUZECK36KK4XFY6C2YPTSCNKRDA6", Limit("100.25")),
	)

	txe := tx.Sign(seed)
	txeB64, _ := txe.Base64()

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAALy+aX1ueJ5Q7EcqD543OWZszFYHG2zJdImOryX8vwQsAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAYAAAABVVNEAAAAAACNliOo1IldDIXfP/ojbwC1RNMkErflK5cuPC1h85CaqAAAAAA7wO+gAAAAAAAAAAH8vwQsAAAAQBrY3JSXAKcAnxUDAnX4E1V7EEH1eDIlgaHbz3giD9sBZeou40zNmI1kGzTM/h3KSW4snAFLwDHNGSN0I53VTQU=
}

// ExampleChangeTrust_maxLimit creates and signs a simple transaction with ChangeTrust operation (maximum limit), and then
// encodes it into a base64 string capable of being submitted to vixal-core.
func ExampleChangeTrust_maxLimit() {
	seed := "SDL6ZQL52IEX34E574CGWKYX5MVPILN7F44QESY7VLWLKUSXZNBVERDY"
	tx := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		Trust("USD", "GCGZMI5I2SEV2DEF3477UI3PAC2UJUZECK36KK4XFY6C2YPTSCNKRDA6"),
	)

	txe := tx.Sign(seed)
	txeB64, _ := txe.Base64()

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAALy+aX1ueJ5Q7EcqD543OWZszFYHG2zJdImOryX8vwQsAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAYAAAABVVNEAAAAAACNliOo1IldDIXfP/ojbwC1RNMkErflK5cuPC1h85CaqH//////////AAAAAAAAAAH8vwQsAAAAQOQeMy29nTZc4n5PoE/OkcVNHvOqNx6T+qqoe833uB8uwLCZu3h63jCPZGeIeAwNYiImwTjaxRcUD/jPWLjxpA0=
}

// ExampleRemoveTrust creates and signs a simple transaction with ChangeTrust operation (remove trust), and then
// encodes it into a base64 string capable of being submitted to vixal-core.
func ExampleRemoveTrust() {
	seed := "SDL6ZQL52IEX34E574CGWKYX5MVPILN7F44QESY7VLWLKUSXZNBVERDY"
	operationSource := "GDEGETQRNUMSRKOKSIED2PSYZL7M5KFSH65KUYB4Y72RR36NVUMYXFU2"
	tx := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		RemoveTrust(
			"USD",
			"GCGZMI5I2SEV2DEF3477UI3PAC2UJUZECK36KK4XFY6C2YPTSCNKRDA6",
			SourceAccount{operationSource},
		),
	)

	txe := tx.Sign(seed)
	txeB64, _ := txe.Base64()

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAALy+aX1ueJ5Q7EcqD543OWZszFYHG2zJdImOryX8vwQsAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAQAAAADIYk4RbRkoqcqSCD0+WMr+zqiyP7qqYDzH9Rjvza0ZiwAAAAYAAAABVVNEAAAAAACNliOo1IldDIXfP/ojbwC1RNMkErflK5cuPC1h85CaqAAAAAAAAAAAAAAAAAAAAAH8vwQsAAAAQM3qBd6NaOyfsa8J+ry0yGMWdIuhNNP9Ol1gfjr6Ripw923akorY4gTlnC223w/2yeVhmCKxIE/tH97hGKgMfgs=
}

// ExampleManageOffer creates and signs a simple transaction with ManageOffer operations, and then
// encodes it into a base64 string capable of being submitted to vixal-core.
func ExampleManageOffer() {
	rate := Rate{
		Selling: NativeAsset(),
		Buying:  CreditAsset("USD", "GCGZMI5I2SEV2DEF3477UI3PAC2UJUZECK36KK4XFY6C2YPTSCNKRDA6"),
		Price:   Price("125.12"),
	}

	seed := "SDL6ZQL52IEX34E574CGWKYX5MVPILN7F44QESY7VLWLKUSXZNBVERDY"
	tx := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		CreateOffer(rate, "20"),
		UpdateOffer(rate, "40", OfferID(2)),
		DeleteOffer(rate, OfferID(1)),
	)

	txe := tx.Sign(seed)
	txeB64, _ := txe.Base64()

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAALy+aX1ueJ5Q7EcqD543OWZszFYHG2zJdImOryX8vwQsAAABLAAAAAAAAAABAAAAAAAAAAAAAAADAAAAAAAAAAMAAAAAAAAAAVVTRAAAAAAAjZYjqNSJXQyF3z/6I28AtUTTJBK35SuXLjwtYfOQmqgAAAAAC+vCAAAADDgAAAAZAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAABVVNEAAAAAACNliOo1IldDIXfP/ojbwC1RNMkErflK5cuPC1h85CaqAAAAAAX14QAAAAMOAAAABkAAAAAAAAAAgAAAAAAAAADAAAAAAAAAAFVU0QAAAAAAI2WI6jUiV0Mhd8/+iNvALVE0yQSt+Urly48LWHzkJqoAAAAAAAAAAAAAAw4AAAAGQAAAAAAAAABAAAAAAAAAAH8vwQsAAAAQOEzXtyiLFp4C7Qr75rz6BjTcpkKUsctC18UVn71SGD6bwYfHF4bV3C38eWuBk9/LbMzJvKnC1ELr8aFb7zn8Aw=
}

// ExampleCreatePassiveOffer creates and signs a simple transaction with CreatePassiveOffer operation, and then
// encodes it into a base64 string capable of being submitted to vixal-core.
func ExampleCreatePassiveOffer() {
	rate := Rate{
		Selling: NativeAsset(),
		Buying:  CreditAsset("USD", "GCGZMI5I2SEV2DEF3477UI3PAC2UJUZECK36KK4XFY6C2YPTSCNKRDA6"),
		Price:   Price("125.12"),
	}

	seed := "SDL6ZQL52IEX34E574CGWKYX5MVPILN7F44QESY7VLWLKUSXZNBVERDY"
	tx := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		CreatePassiveOffer(rate, "20"),
	)

	txe := tx.Sign(seed)
	txeB64, _ := txe.Base64()

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAALy+aX1ueJ5Q7EcqD543OWZszFYHG2zJdImOryX8vwQsAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAQAAAAAAAAAAVVTRAAAAAAAjZYjqNSJXQyF3z/6I28AtUTTJBK35SuXLjwtYfOQmqgAAAAAC+vCAAAADDgAAAAZAAAAAAAAAAH8vwQsAAAAQCxn3ppN3kUQf7tmG/SC8wepgFu8MLfscndzW2/xnKlIRQ7I5i2vqKPCzVnPvXbfUQciDBtRWZix5yZv/rV7KQg=
}

// ExampleAccountMerge creates and signs a simple transaction with AccountMerge operation, and then
// encodes it into a base64 string capable of being submitted to vixal-core.
func ExampleAccountMerge() {
	seed := "SDL6ZQL52IEX34E574CGWKYX5MVPILN7F44QESY7VLWLKUSXZNBVERDY"
	tx := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		AccountMerge(
			Destination{"GAI5G7FANU5M3SPUB2RQBOKCACITUOHOIGAKLG6F6MJPIUB3N2ICGSGB"},
		),
	)

	txe := tx.Sign(seed)
	txeB64, _ := txe.Base64()

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAALy+aX1ueJ5Q7EcqD543OWZszFYHG2zJdImOryX8vwQsAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAgAAAAAEdN8oG06zcn0DqMAuUIAkTo47kGApZvF8xL0UDtukCMAAAAAAAAAAfy/BCwAAABAZrTLRkiX7GYwzyKl2WMXz9huT/r8uv0JWWUMj5VMLR1ZjRKCIMT7qFMCgcjfdd6eX8otOiQpgOy137vafGxFCg==
}

// ExampleInflation creates and signs a simple transaction with Inflation operation, and then
// encodes it into a base64 string capable of being submitted to vixal-core.
func ExampleInflation() {
	seed := "SDL6ZQL52IEX34E574CGWKYX5MVPILN7F44QESY7VLWLKUSXZNBVERDY"
	tx := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		Inflation(),
	)

	txe := tx.Sign(seed)
	txeB64, _ := txe.Base64()

	fmt.Printf("tx base64: %s", txeB64)

	// Output: tx base64: AAAAALy+aX1ueJ5Q7EcqD543OWZszFYHG2zJdImOryX8vwQsAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAkAAAAAAAAAAfy/BCwAAABAXq/Tx5BUD5ji4060MN5xhXe4BNMdz5FMl/+Sjl9/a9npXZUh695GKfX/ZFJhM4BF2fuWBlGF5zeicMo8ezOgCw==

}
