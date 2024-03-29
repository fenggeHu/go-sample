# 内建函数 new 和 make区别
- new。这个用来分配内存的内建函数，但与其它语言中的同名函数不同，它不会初始化内存，只会将内存置零。
也就是说，new(T) 会为类型为 T 的新项分配已置零的内存空间，并返回它的地址，也就是一个类型为 *T 的值。
用 Go 的术语来说，它返回一个指针，该指针指向新分配的，类型为 T 的零值。
- make。它只用于创建切片（slices）、映射（maps）和信道（channels），并返回类型为 T（而非 *T）的一个已初始化 （而非置零）的值。
出现这种用差异的原因在于，这三种类型本质上为引用数据类型，它们在使用前必须初始化。
- 总之，make 只适用于切片（slices）、映射（maps）和信道（channels）且不返回指针。若要获得明确的指针，请使用 new 分配内存。

# 数组和切片
- 数组是值，将一个数组赋予另一个数组会复制其所有元素。 
  若将某个数组传入某个函数，它将接收到该数组的一份副本而非指针。
  数组类型 [10]int 和 [20]int 是不同的。
- 切片通过对数组进行封装，为数据序列提供了更通用、强大而方便的接口。
  除了矩阵变换这类需要明确维度的情况外，Go中的大部分数组编程都是通过切片来完成的。
- 切片保存了对底层数组的引用，若你将某个切片赋予另一个切片，它们会引用同一个数组。
  若某个函数将一个切片作为参数传入，则它对该切片元素的修改对调用者而言同样可见，这可以理解为传递了底层数组的指针。
