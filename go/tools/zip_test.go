package tools

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestZip(t *testing.T) {
	// 源档案（准备压缩的文件或目录）
	var src = "/xxx/source"
	// 目标文件，压缩后的文件
	var dst = "/xxx/dst.zip"

	// 压缩单独文件夹或单独文件
	if err := Zip(src, dst); err != nil {
		fmt.Println(errors.Wrapf(err, "压缩文件失败").Error())
	}
	fmt.Println("压缩成功")
}
