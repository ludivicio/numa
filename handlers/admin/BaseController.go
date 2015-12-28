package admin

import (
	"fmt"
	"strings"
	"io"
	"crypto/md5"
    "crypto/rand"
    "encoding/base64"
    "encoding/hex"
	"github.com/astaxie/beego"
)

const (
	BIG_IMAGE_PATH   = "./static/upload/bimage/"
	SMALL_IMAGE_PATH = "./static/upload/simage/"
	FILE_PATH        = "./static/upload/attachment/"
)

var pathArr []string = []string{"", BIG_IMAGE_PATH, SMALL_IMAGE_PATH, FILE_PATH}

type BaseController struct {
	beego.Controller
	userID         int32
	userName       string
	moduleName     string
	controllerName string
	actionName     string
}

// Prepare function
func (this *BaseController) Prepare() {

	controllerName, actionName := this.GetControllerAndAction()
	this.moduleName = "admin"
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)

}

func (this *BaseController) loginVerify() {
	if this.controllerName == "account" && (this.actionName == "login" || this.actionName == "logout") {
		fmt.Println("login or logout...")
	} else {
		arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
		if len(arr) == 2 {

		}
	}
}


// 随机生成UUID
func (this *BaseController) GenUid() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	b64 := base64.URLEncoding.EncodeToString(b)
	h := md5.New()
	h.Write([]byte(b64))
	return hex.EncodeToString(h.Sum(nil))
} 
