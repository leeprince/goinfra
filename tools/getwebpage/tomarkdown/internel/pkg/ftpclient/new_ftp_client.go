package ftpclient

import "errors"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/16 03:20
 * @Desc:
 */

type FtpClient struct {
	Conf       Conf
	AccessHost string
}

type Conf struct {
	Host     string
	Port     string
	Username string
	Password string
}

func NewFtpClient(conf Conf, accessHost string) *FtpClient {
	return &FtpClient{
		Conf:       conf,
		AccessHost: accessHost,
	}
}

func (r *FtpClient) checkInit() (err error) {
	if r.Conf.Host == "" {
		err = errors.New("ftpConf.Host must config")
		return
	}
	if r.Conf.Port == "" {
		err = errors.New("ftpConf.Port must config")
		return
	}
	if r.Conf.Username == "" {
		err = errors.New("ftpConf.Username must config")
		return
	}
	if r.Conf.Password == "" {
		err = errors.New("ftpConf.Password must config")
		return
	}
	return
}
