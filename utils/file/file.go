package file

import (
	"fmt"
	g "oj/app/global"
	"oj/utils/uuid"
	"os"
)

// CheckNotExist 检查指定文件是否存在
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// MkDir 建立文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// IsNotExistMkDir 检查文件夹是否存在, 不存在则创建
func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// CodeSave 保存代码
func CodeSave(code []byte) (string, error) {
	dirName := "code/" + uuid.GetUUid()
	path := dirName + "/main.go"
	err := os.Mkdir(dirName, 0777)
	if err != nil {
		g.Logger.Errorf("mkdir err:%v", err)
		return "", fmt.Errorf("internal err")
	}
	f, err := os.Create(path)
	if err != nil {
		g.Logger.Errorf("create err:%v", err)
		return "", fmt.Errorf("internal err")
	}
	f.Write(code)
	defer f.Close()
	return path, nil
}
