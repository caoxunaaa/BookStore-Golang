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

func UploadFile(r *http.Request, name string) (fileName, storePath string) {
	_, fh, err := r.FormFile(name)
	if err != nil {
		return "", ""
	}
	fileName = fh.Filename
	nameSplit := strings.Split(fileName, ".")
	dir := nameSplit[len(nameSplit)-1]
	storePath = "assets/" + dir + "/" + fileName

	_, err = os.Stat(Services.C.FileStorage.Path + "assets")
	if os.IsNotExist(err) {
		fmt.Println("目录不存在,创建目录")
		err = os.Mkdir(Services.C.FileStorage.Path+"assets", 0777)
		if err != nil {
			return "", ""
		}
	}
	_, err = os.Stat(Services.C.FileStorage.Path + "assets/" + dir)
	if os.IsNotExist(err) {
		fmt.Println("文件不存在,创建目录")
		err = os.Mkdir(Services.C.FileStorage.Path+"assets/"+dir, 0777)
		if err != nil {
			return "", ""
		}
	}

	err = SaveUploadedFile(fh, Services.C.FileStorage.Path+storePath)
	if err != nil {
		return "", ""
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
