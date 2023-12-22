package consts

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/12/21 17:00
 * @Desc:
 */

/*
基于数据 URI 方案的图片编码格式。它将图片文件的内容转换成了 Base64 编码的字符串，并将其嵌入到一个数据 URI 中。
解析这段代码：
- data: 数据 URI 方案的标识符。
- image/png: 表示嵌入的数据是 PNG 图片格式。在这里的 png 可能是图片的实际格式，也可能是其他格式，比如 jpeg 或 gif。
- base64: 表示编码方式采用了 Base64。

> 读取图片为字节后，直接转为 base64 编码，即可得到`base64_encoded_data`

这段数据 URI 可以直接用于在 HTML 或 CSS 中嵌入图片，以减少网页加载时对服务器的请求次数。它可以作为图片的源（src 属性）或背景图片等
例如：
```
<img src="data:image/png;base64,base64_encoded_data" alt="Base64 Image">
```
或者在 CSS 中：
```
div {
    background-image: url('data:image/png;base64,base64_encoded_data');
}
```
这样的数据 URI 方式可以减少页面加载时对服务器的请求次数，但有时会增加 HTML 或 CSS 文件的大小。这种方法适合小型的图片或者需要动态生成图片的场景。
*/
const (
	DATA_IMAGE_URI_BASE64_PNG  = "data:image/png;base64,"
	DATA_IMAGE_URI_BASE64_JPEG = "data:image/jpeg;base64,"
	DATA_IMAGE_URI_BASE64_JPG  = "data:image/jpg;base64,"
	DATA_IMAGE_URI_BASE64_GIF  = "data:image/gif;base64,"
)
