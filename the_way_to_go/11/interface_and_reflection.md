## 接口与反射

## 接口 Interface

### Interface '不需要'  implements 关键字

也就是不用显式声明实现, 就像是Python 的duck type, 只要能够满足interface 描述的方法即可


有相应接口需要的方法就算是实现了， 不需要特地的去实现, 先有方法实现之后再写一个Interface 也是可以的


### 任何类型都实现了空接口 `interface {}`

那这个就可以传任何类型了, 类似于java 中的`Object`类

