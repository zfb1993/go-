//统计某个文件的重复行数
//打开文件后一行一行扫描
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
	counts := make(map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts,files[1])
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts,arg)
            f.Close()
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

func countLines(f *os.File, counts map[string]int,fileName string) {
    input := bufio.NewScanner(f)
    for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()]  > 1{
			fmt.Println(fileName)
		}
    }
}
//执行 go run .\dup2.go F:\laragon\www\golang\test.txt
//输出
/**
 *3       qqqqqq
 *2       eeeeeeeee
 */