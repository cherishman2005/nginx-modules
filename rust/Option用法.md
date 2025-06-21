# Option用法

在 Rust 中，Option<T>是一个极为重要的枚举类型，用于表示可能存在或不存在的值。它通过类型系统强制处理空值情况，避免了其他语言中常见的空指针异常。以下是其核心用法和实践指南：

## 一、Option<T>的基本定义
```rust
enum Option<T> {
    Some(T),  // 包含一个有效值T
    None,     // 表示值不存在
}
```

* Option<T>是泛型类型，可包裹任意类型T。
* Some(T)表示存在值，None表示值缺失。

## 二、创建 Option 实例

1. 直接构造
```rust
let some_num = Some(5);          // Option<i32>
let no_num: Option<i32> = None;  // 显式指定类型
```

2. 函数返回值
```rust
fn divide(a: f64, b: f64) -> Option<f64> {
    if b == 0.0 {
        None
    } else {
        Some(a / b)
    }
}


fn main() {
	let result = divide(10.0, 2.0);  // Some(5.0)
	if let Some(value) = result {
		println!("Some result={}", value);
	}
	println!("result={:?}", result);
}
```

从其他类型转换
```rust
let maybe_num = "42".parse::<i32>().ok();  // Some(42)
let empty_str = "".parse::<i32>().ok();     // None
```

### 三、解包 Option 值
1. 模式匹配（最安全的方式）
```rust
let opt = Some("hello");
match opt {
    Some(val) => println!("值是: {}", val),
    None => println!("没有值"),
}
```

2. if let 简化匹配
```rust
if let Some(val) = opt {
    println!("简化匹配: {}", val);
}
```

3. 组合子方法（链式操作）

map(f)：转换Some(T)中的值，None保持不变

```rust
let opt = Some(5);
let doubled = opt.map(|x| x * 2);  // Some(10)

let none_opt: Option<i32> = None;
let still_none = none_opt.map(|x| x * 2);  // None
```

and_then(f)：链式调用返回Option的函数
```rust
fn square(x: i32) -> Option<i32> { Some(x * x) }
fn double(x: i32) -> Option<i32> { Some(x * 2) }

let result = Some(3)
    .and_then(square)  // Some(9)
    .and_then(double); // Some(18)
```

filter(predicate)：过滤不符合条件的值
```rust
let opt = Some(5);
let even = opt.filter(|x| x % 2 == 0);  // None

let opt = Some(4);
let even = opt.filter(|x| x % 2 == 0);  // Some(4)
```

4. 解包工具方法
unwrap()：有值时返回值，无值时 panic（仅用于调试或确定有值时）
```rust
let val = Some(5).unwrap();  // 5
// let val = None.unwrap();  // 运行时panic: called `Option::unwrap()` on a `None`

unwrap_or(default)：无值时返回默认值
rust
let val = Some(5).unwrap_or(0);  // 5
let default = None.unwrap_or(10); // 10
```

unwrap_or_else(f)：无值时执行函数生成默认值
```rust
let val = None.unwrap_or_else(|| {
    println!("生成默认值");
    100
});  // 100
```

## 四、Option 与错误处理

通过?操作符传播None，替代异常：

```rust
fn read_file(path: &str) -> Option<String> {
    let file = std::fs::File::open(path).ok()?;
    let mut content = String::new();
    file.read_to_string(&mut content).ok()?;
    Some(content)
}
```

## 五、常见场景与最佳实践
1. 集合操作
```rust
let numbers = vec![1, 2, 3];
let first = numbers.first();       // Option<&i32>
let third = numbers.get(2);        // Option<&i32>
```

2. 函数返回值（可能失败的操作）
```rust
fn find_user(id: u32) -> Option<User> {
    // 从数据库查询用户，未找到返回None
}
```

3. 可选参数
```rust
fn greet(name: Option<&str>) {
    let greeting = name.unwrap_or("陌生人");
    println!("你好，{}！", greeting);
}
```

4. 避免嵌套 Option
```rust
// 不好的写法：嵌套Option
fn bad_example() -> Option<Option<i32>> {
    Some(Some(5))
}

// 好的写法：扁平化Option
fn good_example() -> Option<i32> {
    Some(5)
}
```

## 六、与 Result<T, E> 的区别

* Option<T>：表示值存在或不存在（适用于 “有无” 场景）。
* Result<T, E>：表示操作成功（Ok(T)）或失败（Err(E)），需指定失败原因。

```rust
// 示例：文件读取
fn read_file(path: &str) -> Result<String, std::io::Error> {
    let mut file = std::fs::File::open(path)?;
    let mut content = String::new();
    file.read_to_string(&mut content)?;
    Ok(content)
}
```

## 七、进阶技巧

批量处理 Option

```rust
let opt1 = Some(1);
let opt2 = Some(2);
let opt3 = None;

// 所有Option都为Some时，合并为一个元组
let result = opt1.zip(opt2).zip(opt3);  // None
let result = opt1.zip(opt2);            // Some((1, 2))
```

遍历 Option 中的值
```rust
let opt = Some(10);
opt.iter().for_each(|x| println!("值: {}", x));  // 打印10

let none_opt: Option<i32> = None;
none_opt.iter().for_each(|x| println!("值: {}", x));  // 无输出
```


# 总结

Option<T>是 Rust 避免空指针异常的核心机制，通过类型系统强制开发者处理值缺失的情况。合理使用模式匹配、组合子方法和?操作符，能写出更安全、简洁的代码。记住：永远不要在生产环境中盲目使用unwrap()，除非你 100% 确定值存在。
