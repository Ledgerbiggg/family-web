package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ReadPathAllDir 读取目录下面所有的子目录和文件,并通过回调函数来处理
// path: 目录的路径
// dirFn: 处理子目录的回调函数 dirName: 子目录的名称
// fileFn: 处理文件的回调函数 dirId: 子目录的id fileName: 文件的名称
func ReadPathAllDir(
	path string,
	dirFn func(dirName string) int,
	fileFn func(dirId int, fileName string)) {
	// 读取目录内容
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Failed to read directory: %v\n", err)
		return
	}

	fmt.Println("Directories in", path, ":")
	for _, entry := range entries {
		// 只处理文件夹
		if entry.IsDir() {
			dirId := dirFn(entry.Name())
			readPathAllPic(path+"/"+entry.Name(), dirId, fileFn)
		}
	}
}

// readPathAllPic 读取给定目录下面所有的图片文件，并通过回调函数来处理
// path: 目录的路径
// dirId: 目录id
// fn: 处理图片文件的回调函数 dirId: 目录id fileName: 文件的名称
func readPathAllPic(path string, dirId int, fn func(dirId int, fileName string)) {
	// 定义支持的图片后缀
	imageExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
		".webp": true,
	}

	// 读取当前目录内容
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Failed to read directory: %v\n", err)
		return
	}

	// 遍历当前目录内容
	for _, entry := range entries {
		// 只处理文件
		if !entry.IsDir() {
			// 获取文件扩展名并转换为小写
			ext := strings.ToLower(filepath.Ext(entry.Name()))
			if imageExtensions[ext] {
				// 调用 fn 函数处理图片文件
				fn(dirId, entry.Name())
			}
		}
	}
}
