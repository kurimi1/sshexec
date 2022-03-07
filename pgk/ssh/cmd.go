package ssh

import (
	"github.com/zeromicro/go-zero/core/logx"
)

//Cmd is in host exec cmd
func (ss *SSH) Cmd(host string, cmd string) string {
	session, err := ss.NewSession(host)
	if err != nil {
		logx.Errorf("[ssh][%s]Error create ssh session failed,%s", host, err)
	}
	defer session.Close()
	b, err := session.CombinedOutput(cmd)
	if err != nil {
		logx.Errorf("[ssh][%s]Error exec command failed: %s", host, err)
	}

	if b != nil {
		str := string(b)
		// str = strings.ReplaceAll(str, "\r\n", spilt)
		return str
	}

	return ""
}
