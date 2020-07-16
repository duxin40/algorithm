package tools

import (
	"archive/zip"
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// 打包成zip文件
func Zip(srcDir string, zipFileName string) error {

	// 预防：旧文件无法覆盖
	// os.RemoveAll(zip_file_name)

	// 创建：zip文件
	zipFile, err := os.Create(zipFileName)
	defer zipFile.Close()
	if err != nil {
		return errors.Wrapf(err, "创建目标压缩文件失败")
	} else {

	}

	// 打开：zip文件
	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	// 遍历路径信息
	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, _ error) error {

		// 如果是源路径，提前进行下一个遍历
		if path == srcDir {
			return nil
		}

		// 获取：文件头信息
		header, _ := zip.FileInfoHeader(info)
		header.Name = strings.TrimPrefix(path, srcDir+`/`)

		// 判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += `/`
		} else {
			// 设置：zip的文件压缩算法
			header.Method = zip.Deflate
		}

		// 创建：压缩包头部信息
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})
	if err != nil {
		return errors.Wrapf(err, "遍历路径信息失败")
	}
	return nil
}
