package m3u8

import "github.com/Nyloner/ntool/logs"

type Options struct {
	Command string `goptions:"--cmd, description='Command to exec', obligatory"`
	Url     string `goptions:"--url, description='Which url need download', obligatory"`
	Tasks   int    `goptions:"--tasks, description='Task size.'"`
}

type M3U8 struct {
	Options Options
}

func (m *M3U8) CmdDownload() error {
	logs.Info("call m3u8 download.[url]=%#v", m.Options.Url)
	return nil
}
