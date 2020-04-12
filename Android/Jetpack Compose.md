# Jetpack Compose

[参考文档](https://developer.android.com/jetpack/compose/tutorial)

[Compose Samples Repository](https://github.com/android/compose-samples)

Jetpack Compose is built around composable functions. To create a composable function, just add the `@Composable` annotation to the function name.

Jetpack Compose 除了支持 Column、Row、Flex 等布局方式以外，同样支持 ConstrainLayout 约束布局方式，在不久的将来会更新出来。

Jetpack Compose 内置了多种 Material Design 默认样式，例如 Text 默认 h1、h2、title、body 的字体效果、Button 默认的圆角、线框、无边框显示效果等等，同样 Jetpack Compose 支持 DarkTheme（暗色主题）的配置。

Jetpack Compose 采取单一向下数据流和单一向上事件流的方式构建 UI。简单来说，就是由父组件向子组件传递数据，子组件通过数据构建 UI，当子组件发送交互事件时，通过Lambda 方法将行为的发生交与父组件处理，父组件处理后修改数据，再通过单一向下数据流的原则通知子组件变化。