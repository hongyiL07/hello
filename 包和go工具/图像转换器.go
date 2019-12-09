package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"os"
)

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "jpeg: %v\n", err) //输出JPEG图像
		os.Exit(1)
	}
}
func toJPEG(in io.Reader, out io.Writer) error { //读取并返回图像格式
	img, Kind, err := image.Decode(in) //image.Decode函数查阅一个关于支持格式的表格
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format=", Kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
