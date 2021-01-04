# Flutter

## CLI

```bash
flutter doctor -v # 检查开发环境
flutter create <projectname> # 创建项目 
flutter run # 运行项目
export PUB_HOSTED_URL=https://pub.flutter-io.cn  # 使用国内镜像
export FLUTTER_STORAGE_BASE_URL=https://storage.flutter-io.cn # 使用国内镜像
flutter pub global activate devtools # 全局下载开发者工具
flutter pub global run devtools   # 启动开发者工具服务
```

每次启动你的 app 时，链接都会改变。如果重启 app 后，需要用新的 URL 链接来连接开发者工具。

## 已知问题

当 Flutter 应用执行热重载时，用户的断点会被清除。

## 备忘

在 Dart 语言中使用`_`下划线前缀标识符，会 [强制其变成私有](https://dart.cn/guides/language/language-tour)。

## Hello World

```dart
import 'package:flutter/material.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Welcome to Flutter',
      home: Scaffold(
        appBar: AppBar(
          title: Text('Welcome to Flutter'),
        ),
        body: Center(
          child: Text('Hello World'),
        ),
      ),
    );
  }
}
```

## 更简洁的Hello World

```dart
import 'package:flutter/material.dart';

void main() {
  runApp(
    new Center(
      child: new Text(
        'Hello, world!',
        textDirection: TextDirection.ltr,
      ),
    ),
  );
}
```

## 常用Widget

- [`Text`](https://docs.flutter.io/flutter/widgets/Text-class.html)：该 widget 可让创建一个带格式的文本。
- [`Row`](https://docs.flutter.io/flutter/widgets/Row-class.html)、 [`Column`](https://docs.flutter.io/flutter/widgets/Column-class.html)： 这些具有弹性空间的布局类Widget可让您在水平（Row）和垂直（Column）方向上创建灵活的布局。其设计是基于web开发中的Flexbox布局模型。
- [`Stack`](https://docs.flutter.io/flutter/widgets/Stack-class.html)： 取代线性布局 (和Android中的LinearLayout相似)，[`Stack`](https://docs.flutter.io/flutter/widgets/Stack-class.html)允许子 widget 堆叠， 你可以使用 [`Positioned`](https://docs.flutter.io/flutter/widgets/Positioned-class.html) 来定位他们相对于`Stack`的上下左右四条边的位置。Stacks是基于Web开发中的绝度定位（absolute positioning )布局模型设计的。
- [`Container`](https://docs.flutter.io/flutter/widgets/Container-class.html)： [`Container`](https://docs.flutter.io/flutter/widgets/Container-class.html) 可让您创建矩形视觉元素。container 可以装饰为一个[`BoxDecoration`](https://docs.flutter.io/flutter/painting/BoxDecoration-class.html), 如 background、一个边框、或者一个阴影。 [`Container`](https://docs.flutter.io/flutter/widgets/Container-class.html) 也可以具有边距（margins）、填充(padding)和应用于其大小的约束(constraints)。另外， [`Container`](https://docs.flutter.io/flutter/widgets/Container-class.html)可以使用矩阵在三维空间中对其进行变换。

## 工程目录结构

```
┬
└ projectname
  ┬
  ├ android      - 包含 Android 相关文件。
  ├ build        - 存储 iOS 和 Android 构建文件。
  ├ ios          - 包含 iOS 相关文件。
  ├ lib          - 包含外部可访问 Dart 源文件。
    ┬
    └ src        - 包含附加源文件。
    └ main.dart  - Flutter 程序入口和新应用程序的起点。当你创建 Flutter 工程的时候会自动生成这些文件。你从这里开始写 Dart 代码
  ├ test         - 包含自动测试文件。
  └ pubspec.yaml - 包含 Flutter 应用程序的元数据。这个文件相当于 React Native 里的 package.json 文件。
```

## 把资源文件放到哪并且如何调用

一个 Flutter 资源就是打包到你应用程序里的一个文件并且在程序运行的时候可以访问。 Flutter 应用程序可以包含下述几种资源类型：

- 静态数据比如 JSON 文件
- 配置文件
- 图标和图片 (JPEG, PNG, GIF, Animated GIF, WebP, Animated WebP, BMP, and WBMP)

Flutter 使用 `pubspec.yaml` 文件来确定应用程序中的资源。该文件在工程的根目录。

```yaml
flutter:
  assets:
    - assets/my_icon.png
    - assets/background.png
```

## 如何在网络中加载图片

在 React Native，你可以在 `Image` 的 `source` 属性中设置 `uri` 和所需的尺寸。

在 Flutter 中，使用 `Image.network` 构造函数来实现通过地址加载图片的操作。

```dart
// Flutter
body: Image.network(
          'https://flutter.io/images/owl.jpg',
```

## 如何安装依赖包和包插件

在 Flutter 中，安装代码包需要按照如下的步骤：

- 在 `pubspec.yaml` 的 dependencies 区域添加包名和版本。下面的例子向大家展示了如何将 `google_sign_in` 的 Dart package 添加到 `pubspec.yaml` 中。一定要检查一下 YAML 文件中的空格。因为 **空格很重要**!

```yaml
dependencies:
  flutter:
    sdk: flutter
  google_sign_in: ^3.0.3
```

- 在命令行中输入 `flutter packages get` 来安装代码包。如果使用 IDE，它自己会运行 `flutter packages get`，或者它会提示你是不是要运行该命令。

- 向下面代码一样在程序中引用代码包：

```dart
import 'package:flutter/cupertino.dart';
```

如果想了解更多相关信息，请参考 [在 Flutter 里使用 Packages](https://flutter.cn/docs/development/packages-and-plugins/using-packages) 和 [Flutter Packages 的开发和提交](https://flutter.cn/docs/development/packages-and-plugins/developing-packages)。

你可以找到很多 Flutter 开发者分享的代码包，就在 [[Flutter packages](https://pub.flutter-io.cn/flutter/) 的 [pub.dev](https://pub.flutter-io.cn/).

## 存储在应用程序中全局有效的键值对

在 Flutter 中，使用 [`shared_preferences`](https://github.com/flutter/plugins/tree/master/packages/shared_preferences) 插件来存储和访问应用程序内全局有效的键值对数据。

`shared_preferences` 插件封装了 iOS 中的 `NSUserDefaults` 和 Android 中的 `SharedPreferences` 来实现简单数据的持续存储。如果要使用该插件，可以在 `pubspec.yaml` 中添加依赖 `shared_preferences`，然后在 Dart 文件中引用包即可。

```yaml
# pubspec.yaml 
dependencies:
  flutter:
    sdk: flutter
  shared_preferences: ^0.4.3
```

```dart
// Dart
import 'package:shared_preferences/shared_preferences.dart';
```

要实现持久数据存储，使用 `SharedPreferences` 类提供的 setter 方法即可。 Setter 方法适用于多种原始类型数据，比如 `setInt`, `setBool`, 和 `setString`。要读取数据，使用 `SharedPreferences` 类中相应的 getter 方法。每一个 setter 方法都有对应的 getter 方法，比如，`getInt`, `getBool`, 和 `getString`。

```dart
// Dart
SharedPreferences prefs = await SharedPreferences.getInstance();
_counter = prefs.getInt('counter');
prefs.setInt('counter', ++_counter);
setState(() {
  _counter = _counter;
});
```

## 使用 SQLite

在 Flutter 中，使用 [SQFlite](https://pub.flutter-io.cn/packages/sqflite) 插件实现此功能。

## 设置推送通知

在 Flutter 中，使用 [Firebase_Messaging](https://github.com/flutter/plugins/tree/master/packages/firebase_messaging) 插件实现此功能。请查阅 [`firebase_messaging`](https://pub.flutter-io.cn/packages/firebase_messaging) 插件文档。