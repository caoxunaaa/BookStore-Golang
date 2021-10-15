package Utils

import (
	"WebApi/Services"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func UploadFile(r *http.Request, name string) (fileName, storePath string, err error) {
	_, fh, err := r.FormFile(name)
	if err != nil {
		return "", "", err
	}
	fileName = fh.Filename
	nameSplit := strings.Split(fileName, ".")
	dir := nameSplit[len(nameSplit)-1]
	storePath = "Assets/" + dir + "/" + fileName
	fmt.Println(dir)
	fmt.Println(Services.C.FileStorage.Path + "Assets")
	_, err = os.Stat(Services.C.FileStorage.Path + "Assets")
	if os.IsNotExist(err) {
		fmt.Println("目录不存在,创建目录")
		err = os.Mkdir(Services.C.FileStorage.Path+"Assets", 0777)
		if err != nil {
			fmt.Println("目录不存在,创建目录")
			return "", "", err
		}
	}
	_, err = os.Stat(Services.C.FileStorage.Path + "Assets/" + dir)
	if os.IsNotExist(err) {
		fmt.Println("文件不存在,创建目录")
		err = os.Mkdir(Services.C.FileStorage.Path+"Assets/"+dir, 0777)
		if err != nil {
			return "", "", err
		}
	}

	err = SaveUploadedFile(fh, Services.C.FileStorage.Path+storePath)
	if err != nil {
		return "", "", err
	}
	return
}

func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
