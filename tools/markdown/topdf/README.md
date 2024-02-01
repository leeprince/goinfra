# macOS 将 markdown 转为 pdf

```shell
brew install pandoc
```

```
# 未包含中文时，正常转换
pandoc 1.md -o 1.pdf
```

## pdflatex not found. Please select a different --pdf-engine or install pdflatex
这个错误信息表明，在尝试编译LaTeX文档生成PDF文件时，系统没有找到`pdflatex`编译器。`pdflatex`是LaTeX工作流程中用来直接生成PDF格式文档的工具，它是TeX发行版（如TeX Live或MiKTeX）的一部分。

要解决这个问题，请按照以下步骤操作：

1. **安装TeX发行版**：
- 对于Windows用户，可以安装MiKTeX（https://miktex.org/download）。
- 对于Linux用户，可以通过包管理器安装TeX Live（在Ubuntu/Debian上使用`sudo apt-get install texlive-full`命令；在Fedora/CentOS上使用`sudo dnf install texlive-scheme-full`命令）。
- 对于macOS用户，推荐安装MacTeX（https://tug.org/mactex/）。

2. **验证安装**：
安装完成后，确保`pdflatex`已经被添加到系统的PATH环境变量中，你可以打开终端（命令行）并输入`pdflatex --version`来检查是否已经成功安装和可用。

3. **重新运行Pandoc**：
一旦`pdflatex`安装完成并且能在命令行下正常运行，再次尝试用Pandoc将Markdown转换为PDF，并指定PDF引擎为pdflatex，例如：

```sh
pandoc mydocument.md -o output.pdf --pdf-engine=pdflatex
```

如果之前已经在使用Pandoc并且仍然遇到问题，确认一下你正在使用的Pandoc版本与当前系统上的`pdflatex`兼容，并且路径设置正确。


## macOS 的操作系统，我不想安装 mactex，因为它很大。我有个想法 typora 既然可以正常将 markdown 转为 pdf 是否有其他更小的版本可以使用
是的，如果你不想安装完整的MacTeX发行版（因为它确实比较大），你可以选择安装一个较小的TeX环境来满足基本的PDF编译需求。对于只涉及Markdown到PDF转换且不需要处理复杂的LaTeX文档功能的情况，可以考虑以下替代方案：

BasicTeX： BasicTeX是MacTeX的一个轻量级版本，包含最少的基本包，并允许按需安装额外的宏包。可以在MacTeX官网找到BasicTeX的下载链接。

TinyTeX： TinyTeX是由Yihui Xie开发的一个专门为R和Markdown用户设计的小型TeX发行版。它会根据需要自动下载必要的宏包，从而保持小巧。TinyTeX同样适用于macOS系统，你可以在TinyTeX项目页面上找到安装指南。

依赖于Typora内置引擎： Typora本身在渲染Markdown到PDF时并不一定需要你在本地安装pdflatex等工具。Typora内部集成了必要的库来处理转换，因此理论上即使不单独安装TeX环境也能正常工作。但请注意，Typora可能没有提供像完整LaTeX发行版那样的全部排版定制能力。

然而，请注意，这些小型发行版虽然初始安装小，但如果在转换过程中遇到需要用到非默认提供的LaTeX宏包或样式文件时，它们将会自动下载所需的组件以完成编译，这可能会消耗网络流量并最终占用一定的磁盘空间。

最后添加环境变量：
```
# MacTex 的基础版本 BasicTeX：用于将 markdown 转为 pdf
export PATH=${PATH}:/usr/local/texlive/2023basic/bin/universal-darwin
```

# 需要支持中文
包含中文时，直接使用 pandoc 命令会报错：
```
➜  sudo pandoc 1.md -o 1.pdf                          

Error producing PDF.
! LaTeX Error: Unicode character 我 (U+6211)
               not set up for use with LaTeX.

See the LaTeX manual or LaTeX Companion for explanation.
Type  H <return>  for immediate help.
 ...                                              
                                                  
l.58 hello word 我
```

```
➜  sudo pandoc 1.md -o 1.pdf --pdf-engine=xelatex     

[WARNING] Missing character: There is no 我 (U+6211) (U+6211) in font [lmroman10-regular]:mapping=t
[WARNING] Missing character: There is no 爱 (U+7231) (U+7231) in font [lmroman10-regular]:mapping=t
[WARNING] Missing character: There is no 中 (U+4E2D) (U+4E2D) in font [lmroman10-regular]:mapping=t
[WARNING] Missing character: There is no 国 (U+56FD) (U+56FD) in font [lmroman10-regular]:mapping=t
```


## 解决步骤
按如下步骤解决后，得到完整的指令为：
```
sudo pandoc 1.md -o 1.pdf --pdf-engine=xelatex --template=pandoc-template.tex
```

## tlmgr （TeX Live Manager）
```

sudo tlmgr install sourcesanspro
```

## 使用 `pandoc-template.tex` 模版
> 备份为 `pandoc-template.bak.tex`

修改 `pandoc-template.bak.tex` 中的 `\setCJKmainfont`， 并直接使用 MacOS 中的字体

修改前：
```
\setCJKmainfont{宋体}[BoldFont=SimHei, ItalicFont=KaiTi] %配置中文字体
```

修改后：
```
\setCJKmainfont[
    Path = /Library/Fonts/,
    BoldFont = Arial Unicode,
    ItalicFont = Arial Unicode
]{Arial Unicode}
```

## 通过 `tlmgr （TeX Live Manager）` 安装依赖 
```
sudo tlmgr update --self --all
sudo tlmgr install xecjk
sudo tlmgr install footnotebackref
sudo tlmgr install csquotes
sudo tlmgr install mdframed
sudo tlmgr install zref
sudo tlmgr install needspace
sudo tlmgr install sourcecodepro
sudo tlmgr install titling
```







