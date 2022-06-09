/*
 * File: \src\ch8\du1\main.go                                                  *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/19 , 19:19:40                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/19 , 19:29:52                              *
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
)

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

	//输出结果
	var nfiles, nbytes int64

	for size := range fileSizes {
		nfiles++
		nbytes += size
	}

	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
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
