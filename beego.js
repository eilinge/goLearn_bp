// GO FILETYPE
/*

package main

var a uint  // 定义a为uint类型

const b string

type c uint  // uint取别名c

type d struct{}

type interface{}

func main() {}
*/

// Go baseType

/*
intN
    N: wei
    byte: N/8
    valueRange: -2^N/2-1 ~ 2^N/2, 0 ~ 2^N-1

-- 布尔型: bool
    长度: 1 byte
    取值: true, false
    注意事项: 不可以用数字代表true, false

-- 整型: int, uint
    根据平台可能是32, 64 bit

-- 8位整型: int8, uint8
    长度: 1字节
    取值范围: -128~127/0~255

-- 字节型: byte(uint8别名)

-- 16位整型: int16/uint16
    长度: 2byte
    取值范围: -32768~32767/0~65535

-- 32位整型: int32( rune ) / uint32
    长度: 8字节
    取值范围: -2^32/ 2~2^32/2-1 / 0~2^32-1

-- 64位整型: int64/uint64
    长度: 8byte
    取值范围: -2^64/2 ~2^64/2-1 / 0~2^64-1

-- 浮点型: float32/float64
    长度: 4/8字节
    小数位: 精确到7/15小数位(2*len-1)

-- 复数: complex64/complex128
    长度: 8/16byte

-- 足够保存指针的32位或64位整数型: uintptr

-- 其他值类型
    array/struct/string

-- 引用类型
    slice/map(hash表)/chan(通道/实现并发)

-- 接口类型: interface
-- 函数类型: func

// 类型别名
type (
    byte int8
    rune int32
    ByteSize int64
)

// 多个变量声明与赋值: var
    全局变量的声明可使用var()方式
    全局变量的声明不可以省略var, 但可以使用并行方式
    所有变量都可以使用类型推断
    局部变量不可以使用var()的方式简写, 只能使用并行方式

// 变量的类型转换
    Go中不存在隐式转换, 所有类型转换必须显式声明
    转换只能发生在2种相互兼容的类型之间
    类型转化的格式:
        <ValueA> [:]= <TypeOfValueA>(<Valueb>)

// 常量的定义: const
    常量的值在编译时就已经确定
    常量的定义格式与变量基本相同
    等号右侧必须是常量或常量表达式
    常量表达式中的函数必须是内置函数(len,cap,unsafe.SizeOf())

// 常量的初始化规则与枚举
    在定义常量组时, 如果不提供初始值, 则表示将使用上行的表达式
    使用相同的表达式不代表具有相同的值
    iota是常量的计数器, 从0开始, 组中每定义1个常量自动递增1
    通过初始化规则与iota可以达到枚举的效果
    每遇到一个const关键字, iota就会重置为0
*/