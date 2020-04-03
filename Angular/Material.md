# Material

[参考文档](https://material.angular.io/)

## 安装

```bash
ng add @angular/material
```

该`ng add`命令将安装Angular Material，[Component Dev Kit（CDK）](https://material.angular.io/cdk/categories)和[Angular Animations，](https://angular.io/guide/animations)并询问您以下问题以确定要包括的功能：

1. 选择一个预先构建的主题名称，或为一个自定义主题选择“custom”：

   您可以从[prebuilt material design themes](https://material.angular.io/guide/theming#using-a-pre-built-theme)进行选择，也可以设置可扩展的[自定义主题](https://material.angular.io/guide/theming#defining-a-custom-theme)。

2. 设置Angular Material的浏览器动画：

   导入[`BrowserAnimationsModule`](https://angular.io/api/platform-browser/animations/BrowserAnimationsModule)到您的应用程序中，以启用Angular的[动画系统](https://angular.io/guide/animations)。减少此设置将禁用大多数Angular Material的动画。

该`ng add`命令将额外执行以下配置：

- 将项目依赖项添加到 `package.json`
- 将Roboto字体添加到您的 `index.html`
- 将Material Design图标字体添加到您的 `index.html`
- 将一些全局CSS样式添加到：
  - 删除 `body`的margins
  - 设置`height: 100%`上`html`和`body`
  - 将Roboto设置为默认应用程序字体

## 使用预定义主题

可以从 `@angular/material/prebuilt-themes` 直接把主题文件包含到应用中。

有效的预定义主题有：

- `deeppurple-amber.css`
- `indigo-pink.css`
- `pink-bluegrey.css`
- `purple-green.css`

如果你正在使用 Angular CLI，那么只要在 `styles.css` 文件中添加一行就可以了：

```css
@import '@angular/material/prebuilt-themes/deeppurple-amber.css';
```

此外，还可以直接引用这个文件。类似这样：

```html
<link href="node_modules/@angular/material/prebuilt-themes/indigo-pink.css" rel="stylesheet">
```

实际的路径取决于服务器设置。

最后，如果您的应用程序的内容**未**放置在`mat-sidenav-container`元素内，则需要将`mat-app-background`类添加到wrapper元素（例如`body`）。这样可以确保将正确的主题背景应用于您的页面。

## 显示组件

让我们在您的应用程序中显示一个滑块组件，并验证一切正常。

在 app.module.ts 文件中 import 所需的 [schematic](https://material.angular.io/guide/schematics) ，比如 MatSliderModule

```ts
// app.module.ts
import { MatSliderModule } from '@angular/material/slider';
…
@NgModule ({....
  imports: [...,
  MatSliderModule,
…]
})
```

在app.component.html添加 \<mat-slider\> 标签

```tsx
<mat-slider min="1" max="100" step="1" value="1"></mat-slider>
```

## 在文档中查找组件

可以对照着文档编程

[组件](https://material.angular.cn/components/categories)

[Components](https://material.angular.io/components/categories)

## Component schematics

| 名称           |                             说明                             |
| :------------- | :----------------------------------------------------------: |
| `address-form` | 一个表单组，它使用 Material Design 的多个表单控件接收一个送货地址(shipping address) |
| `navigation`   | 创建一个带有响应式 Material Design 侧边栏组件和一个用于显示应用名称的工具栏组件 |
| `dashboard`    |     带有多个 Material Design 卡片和菜单的网格式布局组件      |
| `table`        |  生成一个带有 Material Design 并支持排序和分页的数据表组件   |
| `tree`         |        该组件以 `` 组件来展示一个可交互的文件夹式结构        |



| 名称        |                       说明Description                        |
| :---------- | :----------------------------------------------------------: |
| `drag-drop` | 该组件使用 `@angular/cdk/drag-drop` 指令来创建一个可交互的 to-do 列表 |

### Address form schematic

运行 `address-form` 原理图会生成一个新的 Angular 组件，它可用于快速开始一个 Material Design 表单组，其中包括：

- 一些 Material Design 表单字段
- 一些 Material Design 单选控件
- 一些 Material Design 按钮

```bash
ng generate @angular/material:address-form <component-name>
```

### Navigation schematic

`navigation` schematic 将会创建一个包括应用名的工具栏和一个能自适应 Material 的断点的侧边栏。

```bash
ng generate @angular/material:nav <component-name>
```

### Table schematic

表格schematic 将创建一个组件，它可以渲染出一个预置了可排序、可分页数据源的 Angular Material `table`。

```bash
ng generate @angular/material:table <component-name>
```

### Dashboard schematic

`dashboard` schematic 将创建一个新组件，它包含一个由 Material Design 卡片组成的动态网格列表。

```bash
ng generate @angular/material:dashboard <component-name>
```

### Tree schematic

`tree` schematic 可用于快速生成一个 Angular 组件，它使用 Angular Material 的 `mat-tree` 组件来展示一个嵌套的文件夹式结构。

```bash
ng generate @angular/material:tree <component-name>
```

### Drag and Drop schematic

`drag-drop` schematic 是由 `@angular/cdk` 提供的，它可用来生成带有 CDK 拖放指令的组件。

```bash
ng generate @angular/cdk:drag-drop <component-name>
```

