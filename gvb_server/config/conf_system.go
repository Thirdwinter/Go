package config

import (
	"strconv"

	log "github.com/ThirdWinter/Go/mylog"
)

type System struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Env         string `yaml:"env"`
	JwtKey      string `yaml:"jwtkey"`
	Accesskey   string `yaml:"accesskey"`
	SecretKey   string `yaml:"secretkey"`
	Bucket      string `yaml:"bucket"`
	Qiniuserver string `yaml:"qiniuserver"`
}

func (s System) Addr() string {
	return s.Host + ":" + strconv.Itoa(s.Port)
}

func (s System) Jk() string {
	log.Error("jwtkey:%s", s.JwtKey)
	return s.JwtKey
}
