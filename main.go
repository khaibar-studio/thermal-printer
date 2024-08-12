package main

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

import (
	"net"

	"github.com/hennedo/escpos"
)

func main() {
	socket, err := net.Dial("tcp", "192.168.8.40:9100")
	if err != nil {
		println(err.Error())
	}
	defer socket.Close()

	p := escpos.New(socket)

	p.Bold(true).Size(2, 2).Write("Hello World")
	p.LineFeed()
	p.Bold(false).Underline(2).Justify(escpos.JustifyCenter).Write("this is underlined")
	p.LineFeed()
	p.QRCode("https://github.com/hennedo/escpos", true, 10, escpos.QRCodeErrorCorrectionLevelH)

	// You need to use either p.Print() or p.PrintAndCut() at the end to send the data to the printer.
	p.PrintAndCut()
}
