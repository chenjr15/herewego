package method

// Int 使用type定义的一个基于int的新类型
type Int int

// 绑定在Int类型上的自增方法
func (i *Int) Increase(num int) *Int {
	*i += Int(num)
	return i
}
