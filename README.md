# goldmark-mermaid

[![Sync to Gitee](https://github.com/OhYee/goldmark-mermaid/workflows/Sync%20to%20Gitee/badge.svg)](https://gitee.com/OhYee/goldmark-mermaid) [![w
orkflow state](https://github.com/OhYee/goldmark-mermaid/workflows/test/badge.svg)](https://github.com/OhYee/goldmark-mermaid/actions) [![codecov](https://codecov.io/gh/OhYee/goldmark-mermaid/branch/master/graph/badge.svg)](https://codecov.io/gh/OhYee/goldmark-mermaid) [![version](https://img.shields.io/github/v/tag/OhYee/goldmark-mermaid)](https://github.com/OhYee/goldmark-mermaid/tags)

goldmark-mermaid is an extension for [goldmark](https://github.com/yuin/goldmark).  

You can use [mermaid](https://github.com/mermaid-js/mermaid) to build svg image in your markdown like [mume](https://github.com/shd101wyy/mume)

## screenshot

There are two demo(using `'` instead of &#8242; in the code block)

1. default config

[Demo1](demo/demo1/main.go)
[Output1](demo/demo1/output.html)

```markdown
'''go
package main

import ()

func main(){}
'''

'''mermaid
pie
    title Key elements in Product X
    "Calcium" : 42.96
    "Potassium" : 50.05
    "Magnesium" : 10.01
    "Iron" :  5
'''
```

![](img/default.png)

2. using `mermaid-svg` and [goldmark-highlighting extension](https://github.com/yuin/goldmark-highlighting)

[Demo2](demo/demo1/main.go)
[Output2](demo/demo1/output.html)

```markdown
'''go
package main

import ()

func main(){}
'''

'''mermaid-svg
pie
    title Key elements in Product X
    "Calcium" : 42.96
    "Potassium" : 50.05
    "Magnesium" : 10.01
    "Iron" :  5
'''
```

![](img/highlighting.png)

## Installation

```bash
go get -u github.com/OhYee/goldmark-mermaid
```

## License

[MIT](LICENSE)
