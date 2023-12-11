# go-bool

# 一、这是什么？

更节省内容的布尔值：

- 数组（已实现）
- 矩阵（待实现）

# 二、安装

```bash
go get -u github.com/compression-algorithm-research-lab/go-bool
```

# 三、API示例

```go
package main

import (
	"fmt"
	"github.com/compression-algorithm-research-lab/go-bool/bool_array"
)

func main() {

	array := bool_array.New(10)
	array.Set(1, true)
	fmt.Println(array.Get(1))

}
```





