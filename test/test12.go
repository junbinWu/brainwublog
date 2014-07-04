package main

import (
	"os"
	"bufio"
	"fmt"
	"time"
)

func main() {
	bufT()
//		iT()
	//	bufStr()
}

func bufStr() {
	if f, err := os.Open("/Users/Berkeley/Desktop/osArgs"); err == nil {
		t1 := time.Now()
		defer f.Close()
		r := bufio.NewReader(f)
		defer func() {

			fmt.Println(time.Now().Sub(t1))
		}()
		for {
			n, e := r.ReadString(byte('\n'))
			fmt.Print(n)
			if e != nil {
				break
			}
		}

	} else {
		fmt.Println(err)
	}
}

func iT() {
	if f, err := os.Open("/Users/Berkeley/Desktop/bytes.txt"); err == nil {
		t1 := time.Now()
		defer f.Close()
		var buf []byte = make([]byte, 4080)
		for {
			n, _ := f.Read(buf)
			if (n == 0) {
				break
			}
			os.Stdout.Write(buf[:n])
		}
		fmt.Println(time.Now().Sub(t1))
	}
}

func bufT() {
	if f, err := os.Open("/Users/Berkeley/Desktop/bytes.txt"); err == nil {
		t1 := time.Now()
		defer f.Close()
		r := bufio.NewReader(f)
		w := bufio.NewWriter(os.Stdout)
		var buf []byte = make([]byte, 300)
		defer func() {
			w.Flush()
			fmt.Println(time.Now().Sub(t1))
		}()
		for {
			n, _ := r.Read(buf)
			if n == 0 {
				break
			}
			w.Write(buf[:n])
		}

	} else {
		fmt.Println(err)
	}
}
