#!/bin/bash


# ----------- 兼容macOS/Linux的`sed -i` 参数 -------
# Linux: 不创建备份的-i 参数的使用`-i`
sedi=(-i)
case "$(uname)" in
  # macOS: 不创建备份的-i 参数的使用`-i ""`
  Darwin*) sedi=(-i "")
esac

docPath=proto/$1/$2/*.json
#echo $docPath

# Swagger Json 的字段简介需要的是description
sed "${sedi[@]}" s/title/description/ $docPath
# ----------- 兼容macOS/Linux的`sed -i` 参数-end -------
