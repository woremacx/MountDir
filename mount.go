package main

import (
	"flag"
	"fmt"

	"github.com/woremacx/MountDir/dirtodir"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 2 {
		fmt.Println("ERROR")
	}
	mntd := flag.Arg(1)
	mnt := flag.Arg(0)
	LkFS := dirtodir.Linkfs(mntd)
	dirtodir.M2(mnt, LkFS).Serve()
}
