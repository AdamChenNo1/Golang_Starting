/*
 * File: \src\ch8\du3\main.go                                                  *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/19 , 19:57:50                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/19 , 22:28:55                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type dirinfo struct {
	path   string
	nfile  int64
	nbytes int64
}

var verbose = flag.Bool("v", false, "show verbose progress messages")
var sema = make(chan struct{}, 20)

func main() {
	//确定起始目录
	flag.Parse()
	roots := flag.Args()
	l := len(roots)
	if l == 0 {
		roots = []string{"."} //默认为当前目录
		l = 1
	}

	//遍历文件树
	var n sync.WaitGroup
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	dirs := map[string]*dirinfo{}
	fileSizes := make(chan dirinfo, l*10)
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	for _, root := range roots {
		dirs[root] = &dirinfo{path: root}
		n.Add(1)
		go walkDir(root, root, &n, fileSizes)
	}

	//输出结果
loop:
	for {
		select {
		case info, ok := <-fileSizes:
			if !ok {
				break loop
			}
			dirs[info.path].nfile += info.nfile
			dirs[info.path].nbytes += info.nbytes
		case <-tick:
			printDiskUsage(dirs)
		}
	}

	printDiskUsage(dirs)

}

func printDiskUsage(dirs map[string]*dirinfo) {
	for _, info := range dirs {
		fmt.Printf("%s contains %d files  %.6f GB\n", info.path, info.nfile, float64(info.nbytes)/1e9)
	}
}

// walkDir递归地遍历以dir为根目录的整个文件树
func walkDir(ancester, dir string, n *sync.WaitGroup, fileSizes chan<- dirinfo) {
	defer n.Done()
	// sema2 <- struct{}{}
	// defer func() { <-sema2 }()

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(ancester, subdir, n, fileSizes)
		} else {
			fileSizes <- dirinfo{ancester, 1, entry.Size()}
		}
	}
}

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}

	return entries
}
