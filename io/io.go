package main

import "os"

// os.Stdin：标准输入的文件实例，类型为*File
// os.Stdout：标准输出的文件实例，类型为*File
// os.Stderr：标准错误输出的文件实例，类型为*File
func main() {

	var buf [16]byte
	os.Stdin.Read(buf[:])
	os.Stdin.WriteString(string(buf[:]))
}
