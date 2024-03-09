package service

import (
	"context"
	"mime/multipart"

	"github.com/ThirdWinter/Go/gvb_server/global"
	"github.com/ThirdWinter/Go/gvb_server/utils/errmsg"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)


func UpLoadFile(file multipart.File,filesize int64) (string,int){
	var (
		// Accesskey = global.Config.System.Accesskey
		// Secretkey = global.Config.System.SecretKey
		Bucket    = global.Config.System.Bucket
		ImgUrl    = global.Config.System.Qiniuserver
		Mac = auth.New(global.Config.System.Accesskey, global.Config.System.SecretKey)
	)
	
	PutPolicy:=storage.PutPolicy{
		Scope: Bucket,
	}
	uptoken:=PutPolicy.UploadToken(Mac)
	cfg:=storage.Config{
		Zone: &storage.ZoneHuabei,
		UseCdnDomains: false,
		UseHTTPS: false,
	}

	PutExtra:=storage.PutExtra{}
	formUploader:=storage.NewFormUploader(&cfg)
	ret:=storage.PutRet{}
	//formUploader.Put(context.Background(), &ret, uptoken, flie,filesize,&PutExtra)
	err:=formUploader.PutWithoutKey(context.Background(), &ret, uptoken, file, filesize, &PutExtra)
	if err != nil {
		return "",errmsg.ERROR
	}
	url:=ImgUrl+ret.Key
	return url, errmsg.SUCCESS
}