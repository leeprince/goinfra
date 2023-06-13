# OCR
---



# 添加git仓库子模块
## 添加 gosseracttest 下 gosseractocrserver为子模块
```
# 进入要克隆另Git仓库的目录。
$ cd goinfra/ocr/gosseracttest

# 执行命令。<URL of the Git repository>是要克隆的Git仓库的URL，<name of the subdirectory>是要将Git仓库克隆到的子目录的
$ git submodule add https://github.com/otiai10/ocrserver ocrserver

# 初始化子模块
$ git submodule init

# 更新子模块
$ git submodule update
```