package main

import (
	"fmt"
	"os"
	"os/user"
	"syscall"
)

/*
Write a program wordfreq to report the frequency of each word in an input text file.
Call input.Split(Buffio.ScanWords) before the first call to Scan to break the input
into words instead of lines.
*/

func main() {
	//words := make(map[string]int)
	//os.ReadDir()
	j, err := os.ReadDir("/etc")
	if err != nil {
		fmt.Printf("problem houston\n")
	}
	for _, v := range j {
		f, e := v.Info()
		if e != nil {
			fmt.Println(e)
		}
		stat, err := f.Sys().(*syscall.Stat_t)
		if !err {
			fmt.Println("syscall.Stat_t")
		}
		userinfo, _ := user.LookupId(fmt.Sprint(stat.Uid))
		groupinfo, _ := user.LookupGroupId(fmt.Sprint(stat.Gid))
		fmt.Printf("%s %d %s %-s %*d %s %s\n", f.Mode(), stat.Nlink, userinfo.Name, groupinfo.Name, 10, stat.Size, fmt.Sprint(f.ModTime()), f.Name())
	}
}
