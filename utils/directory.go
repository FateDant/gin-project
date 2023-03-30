package utils

import "os"

// PathExists 判断路径是否存在
func PathExists(path string) (bool, error) {
	//Stat获取文件属性
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
