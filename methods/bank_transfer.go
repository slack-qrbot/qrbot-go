package methods

import (
	"fmt"

	"github.com/sigurn/crc16"
)

type TlvEncoding interface {
	Encode() string
}

type Tag struct {
	Id, Value string
}

type Method string

const (
	Dynamic Method = "11"
	Static  Method = "12"
)

type guid string

const Napas guid = "A000000727"

type version uint8

const V1 version = 1

type BankTransfer struct {
	version         version
	Method          Method
	ConsumerAccount ConsumerAccount
	Currency        Currency
	Amount          uint32
	CountryCode     CountryCode
	Purpose         string
}

func NewBankTransfer(
	serviceCode ServiceCode,
	bankId string,
	accountNumber string,
	amount uint32,
	purpose string,
) BankTransfer {
	return BankTransfer{
		version: V1,
		Method:  Dynamic,
		ConsumerAccount: ConsumerAccount{
			guid: string(Napas),
			BankAccount: BankAccount{
				BankId:        bankId,
				AccountNumber: accountNumber,
			},
			ServiceCode: serviceCode,
		},
		Currency:    VND,
		CountryCode: Vietnam,
		Amount:      amount,
		Purpose:     purpose,
	}
}

func (t BankTransfer) Encode() string {

	purpose := ""
	if t.Purpose != "" {
		purpose = Tag{Id: "62", Value: Tag{Id: "08", Value: t.Purpose}.Encode()}.Encode()
	}

	amount := ""
	if t.Amount > 0 {
		amount = Tag{Id: "54", Value: fmt.Sprintf("%d", t.Amount)}.Encode()
	}

	payload := fmt.Sprintf("%s%s%s%s%s%s%s6304",
		Tag{Id: "00", Value: fmt.Sprintf("%02d", t.version)},
		Tag{Id: "01", Value: string(t.Method)}.Encode(),
		Tag{Id: "38", Value: t.ConsumerAccount.Encode()},
		Tag{Id: "53", Value: string(t.Currency)},
		amount,
		Tag{Id: "58", Value: string(t.CountryCode)},
		purpose,
	)

	table := crc16.MakeTable(crc16.CRC16_CCITT_FALSE)
	h := crc16.New(table)
	h.Write([]byte(payload))

	return fmt.Sprintf("%s%X", payload, h.Sum16())
}

type ServiceCode string

const (
	TransferToCard    ServiceCode = "QRIBFTTC"
	TransferToAccount ServiceCode = "QRIBFTTA"
)

type Currency string

const (
	JPY Currency = "392"
	KRW Currency = "410"
	MYR Currency = "458"
	CNY Currency = "156"
	IDR Currency = "360"
	PHP Currency = "608"
	SGD Currency = "702"
	THB Currency = "764"
	VND Currency = "704"
)

type CountryCode string

const (
	Japan       CountryCode = "JP"
	Korea       CountryCode = "KR"
	Malaysia    CountryCode = "MY"
	China       CountryCode = "CN"
	Indonesia   CountryCode = "ID"
	Philippines CountryCode = "PH"
	Singapore   CountryCode = "SG"
	Thailand    CountryCode = "TH"
	Vietnam     CountryCode = "VN"
)

type ConsumerAccount struct {
	guid        string
	BankAccount BankAccount
	ServiceCode ServiceCode
}

type BankAccount struct {
	BankId, AccountNumber string
}

func (b BankAccount) Encode() string {
	return fmt.Sprintf("%s%s", Tag{Id: "00", Value: b.BankId}.Encode(), Tag{Id: "01", Value: b.AccountNumber}.Encode())
}

func (p Tag) Encode() string {
	return fmt.Sprintf(p.Id + fmt.Sprintf("%02d", len(p.Value)) + p.Value)
}

func (t Tag) String() string {
	return t.Encode()
}

func (a ConsumerAccount) Encode() string {
	return fmt.Sprintf("%s%s%s",
		Tag{Id: "00", Value: a.guid},
		Tag{Id: "01", Value: a.BankAccount.Encode()},
		Tag{Id: "02", Value: string(a.ServiceCode)},
	)
}

func GetBankIdFromCode(bankCode string) string {
	// create a map of bank codes and bank ids
	bankCodes := map[string]string{
		"ABB": "970425",
		"ACB": "970416",
		"Agribank": "970405",
		"BAB": "970409",
		"BaoVietBank": "970438",
		"BIDV": "970418",
		"BVB": "970454",
		"CB": "970444",
		"CTG": "970415",
		"Dong A Bank": "970406",
		"EIB": "970431",
		"GPBank": "970408",
		"HDB": "970437",
		"Hong Leong Bank": "970442",
		"IVB": "970434",
		"KLB": "970452",
		"LPB": "970449",
		"MBB": "970422",
		"MSB": "970426",
		"NAB": "970428",
		"NCB": "970419",
		"OCB": "970448",
		"OceanBank": "970414",
		"PGB": "970430",
		"Public Bank": "970439",
		"PVcombank": "970412",
		"SCB": "970429",
		"SGB": "970400",
		"SHB": "970443",
		"Shinhan Bank": "970424",
		"SSB": "970440",
		"STB": "970403",
		"TCB": "970407",
		"TPB": "970423",
		"VAB": "970427",
		"VBB": "970433",
		"VCB": "970436",
		"VIB": "970441",
		"VPB": "970432",
		"VRB": "970421",
	}

	return bankCodes[bankCode]
}

