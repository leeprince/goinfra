package ftpclient

import (
	"github.com/leeprince/goinfra/utils/dumputil"
	"net/url"
	"strings"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/16 12:32
 * @Desc:
 */

func TestName(t *testing.T) {
	imageURL := "https://dd.com/ddd/2024/dd.png"
	urlParse, _ := url.Parse(imageURL)
	dumputil.P(urlParse)
	
	urlPathList := strings.Split(urlParse.Path, "/")
	fileName := urlPathList[len(urlPathList)-1]
	dumputil.P(fileName)
}
