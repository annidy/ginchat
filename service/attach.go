package service

import (
	"fmt"
	"ginchat/utils"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		utils.RespFail(ctx.Writer, err.Error())
		return
	}
	suffix := "png"
	fileName := header.Filename
	list := strings.Split(fileName, ".")
	if len(list) > 1 {
		suffix = list[len(list)-1]
	}
	fileName = fmt.Sprintf("%d%04d.%s", time.Now().Unix(), rand.Int31(), suffix)
	url := "./asset/upload/" + fileName
	dstFile, err := os.Create(url)
	if err != nil {
		utils.RespFail(ctx.Writer, err.Error())
		return
	}
	_, err = io.Copy(dstFile, file)
	if err != nil {
		utils.RespFail(ctx.Writer, err.Error())
		return
	}
	utils.RespOk(ctx.Writer, url, "success")
}
