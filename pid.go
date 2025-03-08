package running

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type Pid struct {
	log         *log.Logger
	pidFilePath string
	iPID        int
}

func NewPid(pid int, filePath string) *Pid {
	szPath := []byte(filePath)
	if szPath[len(szPath)-1] != '/' {
		filePath = filePath + "/"
	}
	_, file := filepath.Split(os.Args[0])
	filePath = filePath + file + ".pid"
	return &Pid{
		iPID:        pid,
		pidFilePath: filePath,
	}
}

func (p *Pid) SetErrorLog(err *log.Logger) {
	p.log = err
}

func (p *Pid) logf(format string, args ...interface{}) {
	if p.log == nil {
		return
	}
	p.log.Printf(format, args...)
}

// Init create pid file, set working dir, setpid.
func (p *Pid) Create() error {
	strPID := strconv.Itoa(p.iPID)
	if err := ioutil.WriteFile(p.pidFilePath, []byte(strPID), 0644); err != nil {
		p.logf("ioutil.WriteFile(%v,%v,0644) error(%v)", p.pidFilePath, strPID, err)
		return err
	}
	return nil
}

func (p *Pid) Close() error {
	buf, err := ioutil.ReadFile(p.pidFilePath)
	if err != nil {
		p.logf("ioutil.ReadFile(%v) error(%v)", p.pidFilePath, err)
		return err
	}
	strPID := string(buf)
	pid, err := strconv.Atoi(strPID)
	if err != nil {
		p.logf("strconv.Atoi(%v) error(%v)", strPID, err)
		return err
	}
	if pid != p.iPID {
		return nil
	}
	if err := os.Remove(p.pidFilePath); err != nil {
		p.logf("os.Remove(%v) error(%v)", p.pidFilePath, err)
		return err
	}
	return nil
}
