package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
)

var c string = `
反射是一门非常有用的技术
Go自然也提供了反射在这里介绍一下
学习Go的反射类型之前首先必须要知道 Type 和 Value
比如一个interface{}  obj使用反射 那么可以这样
 t := reflect.TypeOf(obj)  这个t就是Type类型
v := reflect.ValueOf(obj) 这个v就是Value类型
以上例为例
Type是一个接口  主要用于存储obj 中的变量 方法等得类型信息 但无法通过Type接口去使用函数调用赋值等操作
Value是一个结构体 主要用于存储obj中的变量值和方法等信息 通过Value获取的方法和变量是可以进行调用和赋值的
你会发现通过v也是就Value 许多方法的返回值都是Value 比如v.Elem() v.Field() v.Method()等等等等 而我们就可以通过
这些返回的Value对当前的这个字段进行操作 这些Value每个都对应了一个对应的字段
并且你无法从Value中拿到这些字段或方法的信息而是只能够读取其的值或修改之和调用函数
而你想要获取的所有反射信息必须要通过Type来获取  而同样的既然
每一个字段或方法都对应一个Value 那么 也就同时对应了一个Type
下面介绍几个反射的常用操作 读取全部的属性值
代码如下
最好在赋值之前先通过Value CanSet判断是否可以进行赋值
另外对于反射修改值时需要注意一点就是  若是想要队只进行修改那么 需要经过以下步骤 比如操作对象为o
1 v := reflect.ValueOf(&o)你没看错 反射一个指针的Value出来
2 v = v.Elem()取出指针指向的o对应的Value
3 最好在通过v.CanSet进行判断是否可以修改值
4 然后就可以进一下一些列的设置操作了
终于知道是什么原因要让通过反射修改值时
必须先传入对象指针在通过Elem的方式取出对应的Value 才能进行修改值得操作
因为若是修改值
直接传值本身就是值拷贝  所以反射回来的Value也就不是与传进去的值相对应的那个Value了 也就是根本无法寻址 所以必须通过传入地址 的方式保证可寻址
然后再通过Elem将对应到实际对象的Value取出来 进行修改值的操作
`;

func main() {

	db, err := sql.Open("mysql", "root:wjb@/BrainWu_Blog?charset=utf8")
	defer db.Close()
	checkErr(err)

	stmt, err := db.Prepare("insert article set author=?,date=?,title=?,content=?")
	checkErr(err)

	res, err := stmt.Exec("BrainWu", "2014-04-30", "Go-Happy", c)
	//checkErr(err)
	fmt.Println(res)
	fmt.Println(err)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//create table article (
//uid int(10) not null auto_increment,
//author varchar(64) default null,
//date varchar(64) default null,
//title varchar(64) default null,
//content text,
//primary key (uid)
//) CHARSET=utf8;
