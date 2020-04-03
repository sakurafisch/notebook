# TypeScript

## 语法特性

- 类 Classes
- 接口 Interfaces
- 模块 Modules 
- 类型注解 Type annotations
- 编译时类型检查 Compile time type checking 
- Arrow 函数 (类似 C# 的 Lambda 表达式)

## 安装和编译

```bash
$ npm install -g typescript # 安装
$ tsc filename.ts # 编译
```

默认情况下编译器以ECMAScript 3（ES3）为目标但ES5也是受支持的一个选项。

## 类型注解

```ts
function Add(left: number, right: number): number {
    return left + right;
}
```

对于基本类型的注解是number, bool和string。而弱或动态类型的结构则是any类型。

当类型没有给出时，TypeScript编译器利用类型推断以推断类型。如果由于缺乏声明，没有类型可以被推断出，那么它就会默认为是动态的any类型。

举个 🌰

```ts
function area(shape: string, width: number, height: number) {
    var area = width * height;
    return "I'm a " + shape + " with an area of " + area + " cm squared.";
}
 
document.body.innerHTML = area("rectangle", 30, 15);
```

## Interface

举个 🌰

```ts
interface Shape {
    name: string;
    width: number;
    height: number;
    color?: string;
}
 
function area(shape : Shape) {
    var area = shape.width * shape.height;
    return "I'm " + shape.name + " with area " + area + " cm squared";
}
 
console.log( area( {name: "rectangle", width: 30, height: 15} ) );
console.log( area( {name: "square", width: 30, height: 30, color: "blue"} ) );
```

## 类

举个 🌰

```ts
class Shape {
 
    area: number;
    private color: string;
 
    constructor ( name: string, width: number, height: number ) {
        this.area = width * height;
        this.color = "pink";
    };
 
    shoutout() {
        return "I'm " + this.color + " " + this.name +  " with an area of " + this.area + " cm squared.";
    }
}
 
var square = new Shape("square", 30, 30);

console.log( square.shoutout() );
console.log( 'Area of Shape: ' + square.area );
console.log( 'Color of Shape: ' + square.color );
```

还可以添加 public 和 private 访问修饰符。public 成员可以在任何地方访问， private 成员只允许在类中访问。

## extends 继承

举个 🌰，这里的 Shape3D 类继承上文中的 Shape 类

```ts
class Shape3D extends Shape {
 
    volume: number;
 
    constructor ( public name: string, width: number, height: number, length: number ) {
        super( name, width, height );
        this.volume = length * this.area;
    };
 
    shoutout() {
        return "I'm " + this.name +  " with a volume of " + this.volume + " cm cube.";
    }
 
    superShout() {
        return super.shoutout();
    }
}
 
var cube = new Shape3D("cube", 30, 30, 30);
console.log( cube.shoutout() );
console.log( cube.superShout() );
```