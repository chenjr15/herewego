# 方法

## 方法是绑定在参数类型上的！

不是绑定在`struct`上的. 任何是`type`定义的类型是可以被绑定方法的.

```go
type TZ int

func (tz *TZ)Print(){
    print(tz)
}
```

所以`type`定义的类型和原来的类型底层是不一样的, 方法不会被`type`带过去.

## 方法绑定只能作用于相同包内

就是说别的包的绑定不了.

## 方法可以通过类型名直接调用

但是要给方法传一个`receiver`

```go
var a TZ
(*TZ).Print(&a)

```

## 方法可以访问未导出的属性

因为导出不到出是相对于包而言的

