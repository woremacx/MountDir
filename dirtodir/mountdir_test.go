package dirtodir

import (
	"os"
	"path/filepath"
	"syscall"
	"testing"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

var cwd, _ = os.Getwd()
var a lbFS

func TestGetAttr(t *testing.T) {
	attr, err := a.GetAttr("t/f1.txt", nil)
	st := syscall.Stat_t{}
	_ = syscall.Stat(cwd+"/t/f1.txt", &st)

	at := &fuse.Attr{}
	at.FromStat(&st)
	if *at != *attr || err != fuse.OK {
		t.Error("Wrong")
	}
}

func TestOpenDir(t *testing.T) {
	stream, err := a.OpenDir("t", nil)
	if stream[0].Name != "f1.txt" || err != fuse.OK {
		t.Error("Wrong")
	}
}

func TestOpen(t *testing.T) {
	path, _ := filepath.Abs("t/f1.txt")
	nf, err := a.Open(path, 0, nil)
	file, er2 := os.Open(path)
	result := nodefs.NewLoopbackFile(file).String()
	expected := nf.String()
	if result != expected {
		t.Logf("result   [%s]", result)
		t.Logf("expected [%s]", expected)
		t.Errorf("result not match expected %s, result %s", expected, result)
	}
	if er2 != nil {
		t.Error("er2 is not nil")
	}
	if err != fuse.OK {
		t.Error("err is not fuse.OK")
	}
}

func TestGetPath(t *testing.T) {
	a.Root = cwd
	path := a.GetPath("t")
	if path != cwd+"/t" {
		t.Error("Wrong")
	}
}
