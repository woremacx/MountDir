package dirtodir

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
)

type lbFS struct {
	pathfs.FileSystem
	Root string
}

func Linkfs(r string) *lbFS {
	root, err := filepath.Abs(r)
	if err != nil {
		fmt.Println("ERROR LINKING")
		os.Exit(1)
	}
	return &lbFS{
		FileSystem: pathfs.NewDefaultFileSystem(),
		Root:       root,
	}
}

func (fs *lbFS) GetAttr(name string, context *fuse.Context) (a *fuse.Attr, code fuse.Status) {
	fullpath := fs.GetPath(name)
	var err error = nil
	st := syscall.Stat_t{}
	err = syscall.Stat(fullpath, &st)
	if err != nil {
		return nil, fuse.ToStatus(err)
	}
	a = &fuse.Attr{}
	a.FromStat(&st)
	return a, fuse.OK
}

func (fs *lbFS) OpenDir(name string, context *fuse.Context) (stream []fuse.DirEntry, status fuse.Status) {
	f, err := os.Open(fs.GetPath(name))
	if err != nil {
		return nil, fuse.ToStatus(err)
	}
	want := 50
	output := make([]fuse.DirEntry, 0, want)
	infos, err := f.Readdir(want)
	for i := range infos {
		n := infos[i].Name()
		d := fuse.DirEntry{
			Name: n}
		s := fuse.ToStatT(infos[i])
		d.Mode = uint32(s.Mode)
		output = append(output, d)
	}
	f.Close()
	return output, fuse.OK
}

func (fs *lbFS) Open(name string, flags uint32, context *fuse.Context) (fuseFile nodefs.File, status fuse.Status) {
	f, err := os.OpenFile(fs.GetPath(name), int(flags), 0)
	if err != nil {
		return nil, fuse.ToStatus(err)
	}
	nf := nodefs.NewLoopbackFile(f)

	return nf, fuse.OK
}

func (fs *lbFS) GetPath(relPath string) string {
	a := filepath.Join(fs.Root, relPath)
	return a
}

func M2(mnt string, fs *lbFS) *fuse.Server {
	nfs := pathfs.NewPathNodeFs(fs, nil)
	server, _, err := nodefs.MountRoot(mnt, nfs.Root(), nil)
	if err != nil {
		log.Fatalf("ERROR MOUNTING %s", err)
	}
	fmt.Println("MOUNTED")
	return server
}
