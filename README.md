# Markdown Previewer

This tool is a simple markdown file previewer inspired by hexo, and practice on web programming.

![Image](https://github.com/tz70s/markdown-preview/blob/terminate/shots/print.png)

### Dependencies
I use the markdown transformer via [Blackfriday](https://github.com/russross/blackfriday) and [Bluemonday](https://github.com/microcosm-cc/bluemonday)

### Usage
```Bash
git clone https://github.com/tz70s/markdown-preview.git

# check dependencies
go get -u "github.com/russross/blackfriday"
go get -u "github.com/microcosm-cc/bluemonday"

# build
go build
```

```Bash
# example
# ./markdown-preview [inputfile] [port(default = 9090)]
./markdown-preview README.md
# check out usage
./markdown-preview

```

### To Do
Add github api into command line interface
