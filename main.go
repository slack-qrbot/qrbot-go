package main

import (
	"image/png"
	"net/http"
	"qrbot/methods"
	"strconv"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {

	http.HandleFunc("/bank", func(w http.ResponseWriter, r *http.Request) {

		amount, _ := strconv.ParseUint(r.FormValue("amount"), 10, 32)

		bank := methods.NewBankTransfer(
			methods.TransferToAccount,
			methods.GetBankIdFromCode(r.FormValue("bank_code")),
			r.FormValue("bank_account"),
			uint32(amount),
			r.FormValue("purpose"),
		)

		qrCode, _ := qr.Encode(bank.Encode(), qr.M, qr.Auto)

		qrCode, _ = barcode.Scale(qrCode, 300, 300)

		png.Encode(w, qrCode)
	})

	http.ListenAndServe(":8080", nil)
}
