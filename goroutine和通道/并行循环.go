package main

import "fmt"

//处理文件的顺序没有关系，因为每一个缩放操作和其他的操作独立
//像这样一些完全独立的子问题组成的问题叫做 高度并行
//高度并行的问题是最容易实现并行的，有许多并行机制来实现线性扩展

func main() {
	fmt.Println("并行循环")

}

//没有一个直接的访问等待 goroutine 结束，但是可以修改内层 goroutine，通过一个共享的通道发送事件来向外层 goroutine 报告它的完成
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func() {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}
	//等待 goroutine 完成
	for range filenames {
		<-ch
	}
} //这里作为一个字面量函数的显式参数传递 f ，而不是在 for 循环中声明 f

//下面这个使用一个缓冲通道来返回生成的图像文件的名称以及任何错误消息
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) { //为指定文件并行的生成缩略图
	type item struct { //它以任意顺序返回生成的文件名
		thumbfile string //如果任何一个步骤出错就返回一个错误
		err       error
	}
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
		}(f)
	}
	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}
