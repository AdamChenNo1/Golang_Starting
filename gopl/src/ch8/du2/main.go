/*
 * File: \src\ch8\du2\main.go                                                  *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/19 , 19:39:22                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/19 , 19:52:27                              *
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
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	//确定起始目录
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."} //默认为当前目录
	}

	//遍历文件树
	fileSizes := make(chan int64)

	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}

		close(fileSizes)
	}()

	var tick <-chan time.Time

	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	//输出结果
	var nfiles, nbytes int64

loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.6f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir递归地遍历以dir为根目录的整个文件树
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return entries
}
