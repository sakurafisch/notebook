# Angular

## CLI

```bash
npm install -g @angular/cli  # 全局安装 Angular CLI
ng new <app-name>  # 创建应用
cd <app-name>  # 进入应用工作区目录

ng serve --open  #启动
```

`ng serve` 命令会启动开发服务器、监视文件，并在这些文件发生更改时重建应用。

### ng generate

[参考文档：CLI Generate](https://angular.io/cli/generate)

```bash
ng generate <schematic> [options]
# 比如 
ng generate module <module-name>  # 创建组件
ng generate directive <directive-name>  # 创建指令
ng generate pipe <pipe-name>
```

### ng build

[参考文档：CLI Build](https://angular.io/cli/build)

[参考文档：Deployment](https://angular.io/guide/deployment)

```bash
ng build --prod=true --aot=true
```

## NgModule

[参考文档](https://angular.cn/guide/architecture-modules)

模块由一块代码组成，可用于执行一个简单的任务。另外，这项技术还能让你获得*惰性加载*（也就是按需加载模块）的优点，以尽可能减小启动时需要加载的代码体积。

Angular 应用是由模块化的，它有自己的模块系统：NgModules。

每个 Angular 应该至少要有一个模块(根模块)，一般可以命名为：AppModule。

Angular 模块是一个带有 @NgModule 装饰器的类，它接收一个用来描述模块属性的元数据对象。

几个重要的属性如下：

- declarations  - 视图类属于这个模块。 Angular 有三种类型的视图类： 组件 、 指令 和 管道 。
- exports - declaration 的子集，可用于其它模块中的组件模板 。
- imports - 本模块组件模板中需要由其它导出类的模块。
- providers - 服务的创建者。本模块把它们加入全局的服务表中，让它们在应用中的任何部分都可被访问到。
- bootstrap - 应用的主视图，称为根组件，它是所有其它应用视图的宿主。只有根模块需要设置 bootstrap 属性中。

举个 🌰，一个最简单的根模块:

```ts
// app/app.module.ts
import { NgModule }      from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
@NgModule({
  imports:      [ BrowserModule ],
  providers:    [ Logger ],
  declarations: [ AppComponent ],
  exports:      [ AppComponent ],
  bootstrap:    [ AppComponent ]
}) export class AppModule { }
```

在 app/main.ts 引导 AppModule ：

```ts
// app/main.ts
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';
import { AppModule } from './app.module';
 
platformBrowserDynamic().bootstrapModule(AppModule);
```

## Component

组件是一个模板的控制类用于处理应用和逻辑页面的视图部分。

组件是构成 Angular 应用的基础和核心，可用于整个应用程序中。

组件知道如何渲染自己及配置依赖注入。

组件通过一些由属性和方法组成的 API 与视图交互。

创建 Angular 组件的方法有三步：

- 从 @angular/core 中引入 Component 修饰器
- 建立一个普通的类，并用 @Component 修饰它
- 在 @Component 中，设置 selector ，以及 template 。

## Template

使用模板定义组件的视图

举个 🌰

```tsx
<div> 网站地址 : {{site}} </div>
```

在Angular中，默认使用的是双大括号作为插值语法，大括号中间的值通常是一个组件属性的变量名。

## 元数据(Metadata)

元数据告诉 Angular 如何处理一个类。

考虑以下情况我们有一个组件叫作 Component ，它是一个类，直到我们告诉 Angular 这是一个组件为止。

你可以把元数据附加到这个类上来告诉 Angular Component 是一个组件。

在 TypeScript 中，我们用 装饰器 (decorator) 来附加元数据。

举个 🌰

```ts
@Component({
   selector : 'mylist',
   template : '<h2>菜鸟教程</h2>'
   directives : [ComponentDetails]
})
export class ListComponent{...}
```

@Component 装饰器能接受一个配置对象，并把紧随其后的类标记成了组件类。

Angular 会基于这些信息创建和展示组件及其视图。

@Component 中的配置项说明：

- selector - 一个 css 选择器，在 父级 HTML 中寻找一个 <mylist> 标签，然后创建该组件，并插入此标签中。
- templateUrl - 组件 HTML 模板的地址。
- directives - 一个数组，包含此模板需要依赖的组件或指令。
- providers - 一个数组，包含组件所依赖的服务所需要的依赖注入提供者。

## 数据绑定(Data binding)

数据绑定为应用程序提供了一种简单而一致的方法来显示数据以及数据交互，它是管理应用程序里面数值的一种机制。

通过这种机制，可以从HTML里面取值和赋值，使得数据的读写，数据的持久化操作变得更加简单快捷。

数据绑定的语法有四种形式。每种形式都有一个方向——从 DOM 来、到 DOM 去、双向。

### 插值

在 HTML 标签中显示组件值。

```tsx
<h3>
{{title}}
<img src="{{ImageUrl}}">
</h3>
```

### 属性绑定

 把元素的属性设置为组件中属性的值。

```tsx
<img [src]="userImageUrl">
```

### 事件绑定

在组件方法名被点击时触发。

```tsx
<button (click)="onSave()">保存</button>
```

### 双向绑定

使用Angular里的NgModel指令可以更便捷的进行双向绑定。

```tsx
<input [value]="currentUser.firstName"
       (input)="currentUser.firstName=$event.target.value" >
```

## 指令（Directives）

Angular模板是动态的 。当 Angular 渲染它们时，它会根据指令对 DOM 进行修改。

指令是一个带有"指令元数据"的类。在 TypeScript 中，要通过 @Directive 装饰器把元数据附加到类上。

在Angular中包含以下三种类型的指令：

- 属性指令：以元素的属性形式来使用的指令。
- 结构指令：用来改变DOM树的结构
- 组件：作为指令的一个重要子类，组件本质上可以看作是一个带有模板的指令。

举个 🌰

```tsx
<li *ngFor="let site of sites"></li>
<site-detail *ngIf="selectedSite"></site-detail>
```

*ngFor 为 sites 列表中的每个项生成一个 <li> 标签。

*ngIf 表示只有在选择的项存在时，才会包含 SiteDetail 组件。

## 服务(Services)

[参考文档](https://angular.cn/guide/architecture-services)

服务是独立模块，它封装了某一特定功能，可以通过注入的方式供他人使用。

对于与特定视图无关并希望跨组件共享的数据或逻辑，可以创建服务类。 服务类的定义通常紧跟在 `@Injectable()`装饰器之后。该装饰器提供的元数据可以让你的服务作为依赖被注入到客户组件中。

服务分为很多种，包括：值、函数，以及应用所需的特性。

例如，多个组件中出现了重复代码时，把重复代码提取到服务中实现代码复用。

以下是几种常见的服务：

- 日志服务
- 数据服务
- 消息总线
- 税款计算器
- 应用程序配置

以下实例是一个日志服务，用于把日志记录到浏览器的控制台：

```ts
export class Logger {
  log(msg: any)   { console.log(msg); }
  error(msg: any) { console.error(msg); }
  warn(msg: any)  { console.warn(msg); }
}
```

## 依赖注入

[参考文档](https://angular.cn/guide/architecture-services#dependency-injection-di)

依赖注入（或 DI）让你可以保持组件类的精简和高效。有了 DI，组件就不用从服务器获取数据、验证用户输入或直接把日志写到控制台，而是会把这些任务委托给服务。

Angular 能通过查看构造函数的参数类型，来得知组件需要哪些服务。

```ts
constructor(private service: SiteService) { }
```

当 Angular 创建组件时，会首先为组件所需的服务找一个注入器（ Injector ） 。

注入器是一个维护服务实例的容器，存放着以前创建的实例。

如果容器中还没有所请求的服务实例，注入器就会创建一个服务实例，并且添加到容器中，然后把这个服务返回给 Angular 。

当所有的服务都被解析完并返回时， Angular 会以这些服务为参数去调用组件的构造函数。 这就是依赖注入 。

# HttpClient

[参考文档](https://angular.cn/guide/http)

现代浏览器支持使用两种不同的 API 发起 HTTP 请求：`XMLHttpRequest` 接口和 `fetch()` API。

`@angular/common/http` 中的 `HttpClient` 类为 Angular 应用程序提供了一个简化的 API 来实现 HTTP 客户端功能。它基于浏览器提供的 `XMLHttpRequest` 接口。

## 路由

[参考文档](https://angular.cn/guide/router)

路由器会把类似 URL 的路径映射到视图而不是页面。当用户执行一个动作时（比如点击链接），本应该在浏览器中加载一个新页面，但是路由器拦截了浏览器的这个行为，并显示或隐藏一个视图层次结构。

如果路由器认为当前的应用状态需要某些特定的功能，而定义此功能的模块尚未加载，路由器就会按需`lazy-load(惰性加载)`此模块。

路由器会根据你应用中的导航规则和数据状态来拦截 URL。 当用户点击按钮、选择下拉框或收到其它任何来源的输入时，你可以导航到一个新视图。 路由器会在浏览器的历史日志中记录这个动作，所以前进和后退按钮也能正常工作。

要定义导航规则，你就要把`navigation paths(导航路径)`和你的组件关联起来。 路径（path）使用类似 URL 的语法来和程序数据整合在一起，就像模板语法会把你的视图和程序数据整合起来一样。 然后你就可以用程序逻辑来决定要显示或隐藏哪些视图，以根据你制定的访问规则对用户的输入做出响应。