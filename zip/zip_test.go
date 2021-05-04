// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2021/2/22 10:37 上午
package zip_test

import (
	"github.com/comeonjy/util/zip"
	"log"
	"testing"
)

func TestZip(t *testing.T) {
	// 源档案（准备压缩的文件或目录）
	var src = "./exam_1/1/2"
	// 目标文件，压缩后的文件
	var dst = "log.zip"

	if err := zip.Zip(dst, src); err != nil {
		log.Fatalln(err)
	}

}
