package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var projectName string

func init(){
	flag.StringVar(&projectName,"project_name","sunshine","")
}

func main(){
	flag.Parse()//解析参数
	fileName := "./project_template"
	exist, err := PathExists(fileName)
	if err != nil {
		fmt.Println(err)
	}
	if !exist {
		DownloadTemplate()
	}
	// 修改文件名
	err = os.Rename(fileName, projectName)
	if err != nil {
		fmt.Println(err)
	}
}

// 下载文件
func DownloadTemplate() {
	url := "https://github.com/cocofree1/project_template.git"
	cmd := exec.Command("git", "clone", url)
	err := cmd.Run()
	if err != nil {
		fmt.Println("请检查网络")
	}
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
