package main

import (
	"os"
	"time"

	"github.com/Nyloner/ntool/cmd/cli"
	"github.com/Nyloner/ntool/cmd/m3u8"
	"github.com/Nyloner/ntool/cmd/ocr"
	"github.com/Nyloner/ntool/logs"
	"github.com/voxelbrain/goptions"
)

type flagOptions struct {
	Help    goptions.Help `goptions:"-h, --help"`
	Version bool          `goptions:"-v, --version"`
	goptions.Verbs
	M3u8 m3u8.Options `goptions:"m3u8"`
	OCR  ocr.Options  `goptions:"ocr"`
}

func main() {
	options := flagOptions{}
	goptions.ParseAndFail(&options)
	s := time.Now().Unix()
	logs.Info("[START]Start flash.")
	defer func() {
		e := time.Now().Unix()
		logs.Info("[End]Finish Run CMD.[cost]=%#v s", (e - s))
	}()
	if err := cli.ExecCommand(&m3u8.M3U8{Options: options.M3u8}, options.M3u8.Command); err != nil {
		logs.Error("Fail to exec m3u8 cmd.[options]=%#v", options.M3u8)
		os.Exit(-1)
	}
	if err := cli.ExecCommand(&ocr.OCR{Options: options.OCR}, options.OCR.Command); err != nil {
		logs.Error("Fail to exec ocr cmd.[options]=%#v", options.OCR)
		os.Exit(-1)
	}
}
