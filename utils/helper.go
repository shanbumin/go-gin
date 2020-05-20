package utils

import (
	"path/filepath"
	"runtime"
)

/*获取当前文件执行的路径*/
//func GetCurPath() string {
//	file, _ := exec.LookPath(os.Args[0])
//	//得到全路径，
//	path, _ := filepath.Abs(file)
//	rst := filepath.Dir(path)
//	return rst
//}





func GetCurrentPath(skip int) string {
	//获取当前函数Caller reports，取得当前调用对应的文件
	_, f, _, _ := runtime.Caller(skip)
	//解析出所在目录
	dir := filepath.Dir(f)
	return dir
}

