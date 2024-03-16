package ftpclient

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
