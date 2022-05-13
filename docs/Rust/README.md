# Rust

可在 https://crates.io/ 查找要用的 crate.

[async-std](https://github.com/async-rs/async-std):  Async version of the Rust standard library. ([crate doc](https://docs.rs/async-std/latest/async_std/))

## Variable

Unused variable raise warning. If that is intentional, prefix the variable name with an underscore `_` .

### let 和 mut

使用 `let` 声明的变量默认不可二次赋值。

```rust
let x = 1;
x = 2;   // compile error
```

编译报错如下：

```rust
error[E0384]: cannot assign twice to immutable variable `x`
 --> src\main.rs:5:5
  |
4 |     let x = 1;
  |         -
  |         |
  |         first assignment to `x`
  |         help: consider making this binding mutable: `mut x`
5 |     x = 2;
  |     ^^^^^ cannot assign twice to immutable variable
```

使用 `mut` 声明可二次赋值的变量

```rust
let mut x = 1;
println!("The value of x is: {}", x);
x = 2;
println!("The value of x is: {}", x);
```

运行结果如下：

```rust
The value of x is: 1
The value of x is: 2
```

### Shadowing

Declare a new variable with the same name as a previous variable. Shadow a variable by using the same variable’s name and repeating the use of the `let` keyword as follows:

```rust
let y = true;
println!("The value of y is: {}", y);
let y = false; // Shadow previous y
println!("The value of y is: {}", y);
```

### Constant

使用 `const` 声明常量，并且必须带有类型标注。

```rust
const STRING: &str = "hello";
println!("The value of the string constant is: {}", STRING);
```

运行结果如下：

```rust
The value of the string constant is: hello
```

## Scalar Data Type

### Integer

Rust defaults to type `i32` .

| Length | Signed | Unsigned |
| :----: | :----: | :----: |
| 8 bit | i8 | u8 |
| 16 bit | i16 | u16 |
| 32 bit | i32 | u32 |
| 64 bit | i64 | u64 |
| 128 bit | i128 | u128 |
| arch | isize | usize |

### Floating-point

`f32` and `f64` .

Default type is `f64` .

### Boolean

Have a value of either `true` or `false` .

Specified using the keyword `bool` .

### Character

- Represent letters.

- Specified using the keyword `char` .

- Use single quotes.

- Four bytes in size.

## Compound Data Type

### Array

- Continuous group of items

- Fixed length

- Length known at compile time

- 元素具有相同类型

```rust
let array = [1, 2, 3];
```

```rust
let array: [u32; 3] = [1, 2, 3];
```

Access items by index

```rust
let first_item = array[0];
```

Index out-of-bounds:

数组越界会在编译时警告，运行时报错。

以下代码会在编译时报错并停止编译：

```rust
let will_warn = array[100];  // warns by compiler
```

以下代码会在编译时成功编译（无警告），而在运行时报错：

```rust
let len = "Some text".len();
[1][len]; // compile success without warning, panic at runtime
```

### Tuple

- Continuous group of items

- Fixed length

- Length known at compile time

- 元素可以具有不同类型

```rust
let tuple = (true, 2, 3);
```

```rust
let tuple: (bool, u16, u8) = (true, 2, 3);
```

Empty tuple called "unit"

Access item:

```rust
let first_item = tuple.0;
```

```rust
let error = tuple.100; // error of out-of-bounds
```

## Function

- Argument types always required.
- Return type required if value returned. Otherwise, return type is unit.

```rust
fn main() {
    let c = last_char(String::from("Hello"));
    println!("{}", c);  // print 'o'
}

fn last_char(string: String) -> char {
    if string.is_empty() {
        return 'a'
    }
    string.chars().next_back().unwrap()
}
```

## Struct

A type that's composed of other types

Can contain different types

Three flavors of struct:

- Classic struct: Most commonly used. Each field has a name and a type.
- Tuple stuct: Similar to classic structs. Their fields have no names.
- Unit struct: Have no fields. Similar to the `()` unit type.

Define a struct: Use the keyword `struct` followed by the name of the struct.

### Classic Struct

```rust
struct Car {
    make: String,
    model: String,
    year: u32,
}
```

Create an instance of the struct by supplying `key:value` pairs.

```rust
let car = Car {
    make: String::from("Ford"),
    model: String::from("Mustang"),
    year: 1967,
};
```

Can get specific values using dot notation.

```rust
println!("The car is make in {}", car.year);
```

### Tuple Struct

```rust
struct Point2D(u32, u32);
```

Create an instance:

```rust
let origin = Point2D(100, 200);
```

Get value as how tuple do:

```rust
println!("Point contains {:?} and {:?}", origin.0, origin.1);
// Point contains 100 and 200
```

Destructuring assignment:

```rust
let Point2D(x, y) = origin;
println!("Point contains {:?} and {:?}", x, y);
// Point contains 100 and 200
```

## Enum

List all variations of some data.

Referred to as algebraic data types.

Define an enum: 

1. Use the keyword `enum` followed by the name. 
2. Then list all variations. 
3. The variants can also specify their type. Each variant can have a different type.

```rust
enum CardinalDirections {
    North,
    South,
    East,
    West,
}
fn main() {
    let north = CardinalDirections::North;
}
```

```rust
enum CardinalDirections {
    North(String),
    South(String,
    East(String),
    West(String),
}
fn main() {
    let west = CardinalDirections::West(String::from("West"));
}
```

The enum is a custom data type that can be used in code.

More example:

```rust
enum WebEvent {
    PageLoad,
    PageUnload,
    KeyPress(char),
    Paste(String),
    Click { x: i64, y: i64 },
}

fn main() {
    let quit = WebEvent::KeyPress('q');
}
```

### The `Option` enum

```rust
enum Option<T> {
    None,
    Some(T),
}
let something = Some(1);
```

```rust
// 在以下代码中
// 隐式地使用了 Option 进行错误处理
// fruits.get() 的定义为
// pub fun get<I>(&self, index: I) -> Option<&I::Output>
fn main() {
    let fruits = vec!["banana", "apple", "coconut"];
    let first = fruits.get(0);
    println!("{:?}", first);

    let third = fruits.get(2);
    println!("{:?}", third);

    let non_existent = fruits.get(99);
    println!("{:?}", non_existent);
}
// Some("banana")
// Some("coconut")
// None
```

### The `Result` enum

Used for input/output operations

- Parsing strings
- File access
- Data validation

```rust
enum Result<T, E> {
    Ok(T),
    Err(E),
}
```

```rust
// 在以下代码中
// 隐式地使用了 Result 进行错误处理
use std::fs::File;
fn main() {
    let f = File::open("hello.txt");
    let f = match f {
        Ok(file) => file,
        Err(error) => panic!("Can't open the file: {:?}", error),
    };
}
// thread 'main' panicked at 'Can't open the file: Os { code: 2, kind: NotFound, message: "No such file or directory" }', src\main.rs:8:23
```

### `unwrap()` 和 `expect()`

`unwrap()` and `expect()` are helper methods of the Result type.

`unwrap()` returns the value inside the `Ok` variant. Returns a `panic!` macro for the `Err` variant.

- If an `Option` type has **`Some`** value or a `Result` type has a **`Ok`** value, **the value inside them** passes to the next step.
- If the `Option` type has **`None`** value or the `Result` type has **`Err`** value, **program panics**; If `Err`, panics with the error message.

```rust
use std::fs::File;
fn main() {
    let f = File::open("hello.txt").unwrap();
}
```

`expect` returns a value or called the `panic!` macro with a detailed error message. Similar to `unwrap()` but can set a custom message for the panics.

```rust
use std::fs::File;
fn main() {
    let f = File::open("hello.txt").expect("Failed to open hello.txt.");
}
```

### The `?` operator

for `Result` type:

1. Unwraps the value if `Ok` variant
2. Returns an error if `Err` variant

for `Option` type:

1.  Returns a value is with the `Some` variant
2. Returns nothing for the `None` variant

## if 和 else if 和 else

`if` 、 `else if` 、 `else` block can return value, acting as expressions.

```rust
let formal = true;
let greeting = if formal {
    String::from("Good evening")
} else {
    String::from("Hey, friend!")
};
println!("{}", greeting);  // Good evening
```

## match

keyword `match`

a `scrutinee` expression is provided to compare to the patterns.

Arms are evaluated and compared with the `scrutinee` expression.

`match` can return any type, acting as expressions.

```rust
let boolean = true;
let binary = match boolean {
    false => 0,
    true => 1,
};
println!("{}", binary);  // 1
```

## loop

Used to execute over a block of code forever. Or until it is stopped, or the program quits.

```rust
loop {
    println!("I loop forever");
}
```

We can use keyword `break` to stop the loop:

```rust
let mut i = 0;
loop {
    println!("I love loop {} !", i);
    i += 1;
    if i > 2 {
        break;
    }
}
// I love loop 0 !
// I love loop 1 !
// I love loop 2 !
```

we can return value when we use `break`

```rust
let mut i = 1;
let something = loop {
    i *= 2;
    if i > 100 {
        break i;
    }
};
assert_eq!(something, 128);
println!("The value of something is: {}", something);
// The value of somethin is: 128
```

## while

```rust
let mut number = 3;
while number != 0 {
    println!("{}", number);
    number -= 1;
}
// 3
// 2
// 1
```

## for

Iterate over elements in a collection.

With each pass of the for loop, values are extracted from an iterator.

```rust
let a = [10, 20, 30, 40, 50];
for element in a.iter() {
    println!("The value is: {}", element);
}
// The value is: 10
// The value is: 20
// The value is: 30
// The value is: 40
// The value is: 50
```

```rust
for item in 0..5 {
    println!("{}", item * 2);
}
// 0
// 2
// 4
// 6
// 8
```

## panic

Simplest way to handle errors.

`division by zero`、`index-out-of-bound raise` raise panic.

```rust
fn main() {
    let v = vec![1, 2, 3];
    println!("{}", v[6]);
}
// 'main' panicked at 'index out of bounds: the len is 3 but the index is 6
```

## Ownership

Each value has a variable that is called its owner. There can only be one owner at a time. When the owner goes out of scope, the value will be dropped.

Definition of variable's scope: Range within a program for which that variable and the value are valid.

ownership prevents memory safety issues:

- Dangling pointer (including Use-After-Free)
- Double-free
- Memory leaks

### Memory Stack

Stores data with a known, fixed size.

```rust
// An i32 has a known, fixed size at compile time.
// Will take up 32 bits of memory.
// Both the variable and the value of
// the variable would be stored on the stack.
let a: i32 = 5;
```

### Memory Heap

```rust
// Vector is mutable
// Vector size can change when the program is running.
// Vector object stored on stack with pointer to heap.
// Value of vector is stored on heap.
let mut vec1 = vec![1, 2, 3];
vec1.push(4);
println!("{:?}", vec1);  // [1, 2, 3, 4]
drop(vec1);  // Both the data on the stack and the heap would no longer be accessible. 
```

### Ownership of String

```rust
// This String is mutable
// A String stores data on both the stack and the heap
// String object stored on stack with pointer to heap
let mut say1 = String::from("Ca");
say.push('t');
println!("{}", say1);  // Cat
let say2 = say1;  // move the ownership to say2
// say1 is no longer available.
```

```rust
fn main() {
    let mut say = String::from("Ca");
    say.push('t');
    print_out(say);  // variable say has been moved.
    // variable say is now invalid.
    println!("{}", say);  // use after moved, raise error
}

fn print_out(to_print: String) {
    println!("{}", to_print);
}
```

### Clone

It would copy the value in the heap. Really deep copy an instance, rather than moving its ownership.

```rust
let say = String::from("Cat");
let say2 = say.clone();
// Both say and say2 are now available.
drop(say);  // Invalidate the variable say
// Now say is no longer available.
println!("{}", say2);  // Cat
```

```rust
fn main() {
    let mut say = String::from("Ca");
    say.push('t');
    print_out(say.clone());  // Cat
    println!("{}", say);  // Cat
}

fn print_out(to_print: String) {
    println!("{}", to_print);
}
```

### Borrowing

Putting an `&` in front of a variable to borrow it.

It does not transfer ownership, nor does it make a new copy of the value on the heap.

#### Rules of borrowing

1.At any given time, you can have either:

- One mutable reference, or
- Any number of immutable references.

2.References must always be valid.

反例：

```rust
let say1 = String::from("Cat");
let say2 = &say1;
println!("{}", say1);
drop(say1);  // move out of `say1` occurs here
println!("{}", say2);  // borrow later used here
```

编译报错如下：

```rust
error[E0505]: cannot move out of `say1` because it is borrowed
```

#### immutable borrowing

```rust
fn main() {
    let mut say = String::from("Ca");
    say.push('t');
    print_out(&say);  // Cat
    println!("{}", say);  // Cat
}

fn print_out(to_print: &String) {
    println!("{}", to_print);
}
```

#### mutable borrowing

```rust
fn main() {
    let mut my_vec=  vec![1, 2, 3];
    println!("{:?}", my_vec);  // [1, 2, 3]
    add_to_vec(&mut my_vec);
    println!("{:?}", my_vec);  // [1, 2, 3]
}

fn add_to_vec(a_vec: &mut Vec<i32>) {
    a_vec.push(4);
}
```

## String

- UTF-8 Encoded: Rust string type is always valid UTF-8.

- Non-Null-Byte Terminated: Rust string is not null byte terminated.

- Not collections of chars: Rust string is not simply collections of character values.

|  String  |  &str  |
| :------: | :----: |
| CString  |  CStr  |
| OsString | &OsStr |

```rust
fn main() {
    let text = "Hello\nworld\n!";
    let upper = text.to_uppercase();
    let stripped = upper.strip_prefix("HELLO\n").unwrap();
    println!("{}", first_line(stripped));
}

pub fn first_line(string: &str) -> &str {
    string.lines().next().unwrap()
}
```

运行结果如下

```rust
WORLD
```

### The `String` type

- An owned string
- Owns string data
- Data freed when dropped or out of scope

`String` memory consists of three parts

- Length
- Capacity
- Data pointer

```rust
fn string_len(s: String) -> usize {
    return s.len();
}
```

```rust
let s = String::from("Hi\nBye");
```

### The `&str` type

- A borrowed string slice
- Does not own string data
- Data not freed when dropped

`&str` memory consists of two parts

- Length
- Data pointer

```rust
fn string_len(s: &str) -> usize {
    return s.len;
}
```

```rust
let s = String::from("Hi\nBye");
let l = s.lines().next().unwrap();
// When we have L,
// We won't be able to change the 
// data S owns.
// L is an immutable view into data S owns.
```

string literal is embedded into the binary, have type `&str`.

备忘：use `to_own()` method to convert `&str` into `String`

```rust
pub fn first_line(string: String) -> String {
    return string.lines().next().unwrap().to_owned();
}
```

## Collection

请查看文档：[std::collections](https://doc.rust-lang.org/stable/std/collections/)

### `Vec<T>`

`Vec<T>` memory consists of three parts

- Length
- Capacity
- Data pointer

```rust
let v = vec![1u8, 2, 3];
```

如果使用 `v[index]` 来访问其中的元素，当 `index` 越界时，程序崩溃。

如果使用 `v.get(index)` 来访问其中的元素，它返回 `Option` type，当 `index` 越界时，返回 `Option::None`

一个只读的例子：

```rust
fn main() {
    let mut students = vec![ Student {
        name:"Ryan".to_string(),
    }];
    students.push(Student {
        name: "Lisa".to_string(),
    });
    assert!(&students[0] == &Student { name: "Ryan".to_string() });
    assert!(students.get(0) == Some(&Student { name: "Ryan".to_string() }));
    assert!(students.get(10000) == None);
    for student in students.iter() {
        println!("Student name: {}", student.name);
    }
}

#[derive(PartialEq, Eq)]
struct Student {
    name: String
}
```

一个更新元素值的例子：

```rust
let mut v = vec![100, 32, 57];
for i in &mut v {
    *i += 50;
}
```

### `HashMap<K, V>`

```rust
fn main() {
    let mut students = vec![ Student {
        name:"Ryan".to_string(),
    }];
    students.push(Student {
        name: "Lisa".to_string(),
    });
    for student in students.iter() {
        println!("Student name: {}", student.name);
    }

    use std::collections::HashMap;
    let mut enrollment = HashMap::new();
    enrollment.insert("biology".to_string(), students);
    let bio_students = enrollment.get("biology");
    let students = enrollment.remove("biology");
}

#[derive(PartialEq, Eq)]
struct Student {
    name: String
}
```

### `HashSet<T>`

### `VecDeque<T>`

### `LinkedList<T>`

## Slice

Written as `&[T]`

`Slice` are references, items live continuously memory, refer to data owned by another value.

```rust
let v = vec![1u8, 2, 3];
let s = &v[0..2];
```

`&[T]` memory consists of two parts

- Length
- Data pointer

## Traits

Rust 面向对象编程：把 Data 定义在 `enum` type 或者 `struct` type, 把 Behavior 定义在 `trait` .

在 `trait` 中声明方法，然后再分别为不同的数据类型实现这个 `trait` .

在 `trait` 中定义的方法可以有默认实现。

可以通过 `where` 简化 trait bound，以避免使用泛型时函数签名难以阅读。

```rust
fn some_function<T: Display + Clone, U: Clone + Debug>(t: T, u: U) -> i32 {
```

可以用 `where` 简化为

```rust
fn some_function<T, U>(t: T, u: U) -> i32
    where T: Display + Clone,
          U: Clone + Debug
{
```

一个例子：

```rust
pub struct Person {
    name: String
}

pub struct Cat {
    name: String
}

pub struct Rabbit {
    name: String
}

pub trait Eat {
    fn eat_dinner(&self) {
        println!("I eat from a dish.")
    }
}

impl Eat for Person {
    fn eat_dinner(&self) {
        println!("I eat from a plate.")
    }
}

impl Eat for Cat {
    fn eat_dinner(&self) {
        println!("I eat from a cat bowl.")
    }
}

impl Eat for Rabbit {}

fn main() {
    let person = Person {
        name: String::from("Nell")
    };
    person.eat_dinner();

    let cat = Cat {
        name: String::from("Zane")
    };
    cat.eat_dinner();

    let rabbit = Rabbit {
        name: String::from("Leia")
    };
    rabbit.eat_dinner();
}
```

运行结果如下：

```rust
I eat from a plate.
I eat from a cat bowl.
I eat from a dish.
```

另一个例子：

```rust
struct Film {
    title: String,
    director: String,
    studio: String
}

struct Book {
    title: String,
    author: String,
    publisher: String
}

trait Catalog {
    fn describe(&self) {
        println!("We need more information about this type of media");
    }
}

impl Catalog for Film {
    fn describe(&self) {
        println!(
            "{} was directed by {} through {} studios",
            self.title,
            self.director,
            self.studio
        )
    }
}

impl Catalog for Book {
    fn describe(&self) {
        println!(
            "{} was written by {} and published by {}",
            self.title,
            self.author,
            self.publisher
        )
    }
}

struct Album {
    title: String,
    artist: String,
    label: String
}

impl Catalog for Album {}

fn main() {
    let capt_marvel = Film {
        title: String::from("Captain Marvel"),
        director: String::from("Anna Boden and Ryan Fleck"),
        studio: String::from("Marvel")
    };
    capt_marvel.describe();

    let elantris = Book {
        title: String::from("Elantris"),
        author: String::from("Brandon Sanderson"),
        publisher: String::from("Tor Books")
    };
    elantris.describe();

    let let_it_be = Album {
        title: String::from("Let it be"),
        artist: String::from("Beatles"),
        label: String::from("Apple")
    };
    let_it_be.describe();
}
```

运行结果如下：

```rust
Captain Marvel was directed by Anna Boden and Ryan Fleck through Marvel studios
Elantris was written by Brandon Sanderson and published by Tor Books
We need more information about this type of media
```

## crate 和 module

练手代码可参考 [这里](https://github.com/sakurafisch/restaurant-demo)

- **Packages:** A Cargo feature that lets you build, test, and share crates
- **Crates:** A tree of modules that produces a library or executable
- **Modules** and **use:** Let you control the organization, scope, and privacy of paths
- **Paths:** A way of naming an item, such as a struct, function, or module

可以使用绝对路径或相对路径来寻找方法。

### pub

当在 struct 前标注 pub，这个 struct 内的字段依然是 private 的，还要单独对需要公开的字段标注 pub。

当在 enum 前标注 pub，这个 enum 内的字段全部也为 pub。

### use

使用 `use` 引入 `crate`。

可以使用 `pub use` 重导出名称。（我对此仍不完全理解。）

可以嵌套路径来消除大量的 `use` 行

```rust
// use std::cmp::Ordering;
// use std::io;
use std::{cmp::Ordering, io};
```

```rust
// use std::io;
// use std::io::Write;
use std::io::{self, Write};
```

如果希望将一个路径下 **所有** 公有项引入作用域，可以指定路径后跟 `*`，glob 运算符：

```rust
use std::collections::*;
```

### Using External Packages

可在 https://crates.io/ 查找要用的 crate.

To use `rand` in our project, we added this line to *Cargo.toml*:

```rust
rand = "0.8.3"
```

Then we can use rand crate in src/main.rs

```rust
use rand::Rng;

fn main() {
    let secret_number = rand::thread_rng().gen_range(1..101);
}
```

