


## Content-Disposition
`Content-Disposition` HTTP 响应头可以有以下几种形式：

1. **附件形式（attachment）**：
   ```http
   Content-Disposition: attachment; filename="example.pdf"
   ```
   当 `disposition-type` 设置为 `attachment` 时，指示浏览器或其他用户代理应该以下载的形式处理响应体内容。这意味着浏览器不会直接在页面中显示内容，而是提示用户保存文件到本地磁盘，并且可以指定一个默认的文件名，如上述示例中的 "example.pdf"。

2. **内联形式（inline）**：
   ```http
   Content-Disposition: inline; filename="image.jpg"
   ```
   当设置为 `inline` 或者省略 `disposition-type` 时，通常表示内容应当在浏览器内部显示或处理，例如图像或HTML文档。尽管如此，浏览器可能会根据MIME类型和自身能力来决定是否打开一个新窗口、插入图片还是下载文件。同样也可以提供一个建议的文件名用于保存或引用该内容。

3. **额外参数**：
    - `filename` 参数是常见的，用于指定默认的文件名。
    - `filename*` 参数用来支持包含非ASCII字符的文件名，格式通常是 `filename*=UTF-8''encoded-filename`，其中 `encoded-filename` 是编码后的文件名。
    - 其他可选参数包括但不限于 `creation-date`, `modification-date`, 和 `size` 等，但这些并不常用。

例如：

```http
Content-Disposition: attachment; filename="report.docx"; filename*=UTF-8''%E6%8A%A5%E5%91%8A.docx
```
在这个例子中，服务器指定了一个包含中文字符的文件名，并同时提供了两种编码方式供不同的客户端解析。