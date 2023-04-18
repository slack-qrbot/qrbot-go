package methods

import "testing"

func Test_Payload_String(t *testing.T) {
	p := Tag{
		Id:    "00",
		Value: "01",
	}

	want := "000201"

	if p.Encode() != want {
		t.Errorf("got %q, wanted %q", p.String(), want)
	}
}

func TestConsumerAccount_Encode(t *testing.T) {
	c := ConsumerAccount{
		guid: string(Napas),
		BankAccount: BankAccount{
			BankId:        "970468",
			AccountNumber: "0011009950446",
		},
		ServiceCode: TransferToAccount,
	}

	want := "0010A00000072701270006970468011300110099504460208QRIBFTTA"

	if c.Encode() != want {
		t.Errorf("got %q, wanted %q", c.Encode(), want)
	}
}

func TestBankTransfer_Encode(t *testing.T) {
	b := BankTransfer{
		version:     V1,
		Method:      Static,
		Currency:    VND,
		Amount:      180000,
		Purpose:     "thanh toan don hang",
		CountryCode: Vietnam,
		ConsumerAccount: ConsumerAccount{
			guid: "A000000727",
			BankAccount: BankAccount{
				BankId:        "970403",
				AccountNumber: "0011012345678",
			},
			ServiceCode: TransferToAccount,
		},
	}

	want := "00020101021238570010A00000072701270006970403011300110123456780208QRIBFTTA530370454061800005802VN62230819thanh toan don hang63045FAB"

	if b.Encode() != want {
		t.Errorf("got %q, wanted %q", b.Encode(), want)
	}

}
