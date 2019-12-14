package ocr

import "github.com/Nyloner/ntool/logs"

type Options struct {
	Command string `goptions:"--cmd, description='Command to exec', obligatory"`
	Url     string `goptions:"--url, description='Which url need download', obligatory"`
}

type OCR struct {
	Options Options
}

func init() {
	InitBaiDuOCR()
}

func (o *OCR) CmdOcr() error {
	text, err := OCRImage(o.Options.Url)
	if err != nil {
		return nil
	}
	logs.Info("OCR Result.[text]=%#v", text)
	return nil
}
