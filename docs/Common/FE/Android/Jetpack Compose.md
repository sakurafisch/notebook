# Jetpack Compose

[参考文档](https://developer.android.com/jetpack/compose/tutorial)

[概览](https://developer.android.com/jetpack/androidx/versions?hl=zh-cn)

[Android Jetpack 使用入门](https://developer.android.com/jetpack/docs/getting-started?hl=zh-cn)

[Compose Samples Repository](https://github.com/android/compose-samples)

[不妨看看别人的笔记](https://blog.csdn.net/qq_37704124/article/details/100568243)

Jetpack Compose is built around composable functions. To create a composable function, just add the `@Composable` annotation to the function name.

Jetpack Compose 除了支持 Column、Row、Flex 等布局方式以外，同样支持 ConstrainLayout 约束布局方式，在不久的将来会更新出来。

Jetpack Compose 内置了多种 Material Design 默认样式，例如 Text 默认 h1、h2、title、body 的字体效果、Button 默认的圆角、线框、无边框显示效果等等，同样 Jetpack Compose 支持 DarkTheme（暗色主题）的配置。

Jetpack Compose 采取单一向下数据流和单一向上事件流的方式构建 UI。简单来说，就是由父组件向子组件传递数据，子组件通过数据构建 UI，当子组件发送交互事件时，通过Lambda 方法将行为的发生交与父组件处理，父组件处理后修改数据，再通过单一向下数据流的原则通知子组件变化。

## 数据驱动视图

这是 Jetpack MVVM 推崇的一个重要原则。其基本数据流如下所示 ：

- 数据层 Repository 负责从不同数据源获取和整合数据，基本负责所有的业务逻辑。
- ViewModel 持有 Repository，获取数据并驱动 View 层更新。
- View 持有 ViewModel，观察 LiveData 携带的数据，数据驱动 UI。

## ViewModel

```java
public class MyViewModel extends ViewModel {
    private MutableLiveData<Integer> number;
    
    public MutableLiveData<Integer> getNumber() {
        if (number == null) {
            number = new MutableLiveData<>();
            number.setValue(0);
        }
        return number;
    }
    public void add() {
        number.setValue(number.getValue() + 1);
    }
}
```

```java
MyViewModel myViewModel;
// 监听事件
myViewModel = ViewModelProviders.of(this).get(MyViewModel.class);
myViewModel.getNumber().observe(this, new Observer<Integer>() {
   @Override
   public void onChanged(Integer integer) {
       textView.setText(String.valueOf(integer));
   }
});
```

## 在Java代码中使用DataBinding

```gradle
dataBinding {
	enabled true
}
```



```java
MyViewModel myViewModel;
ActivityMainBinding binding;

binding = DataBindingUtil.setContentView(this, R.layout.activity_main);
myViewModel = ViewModelProviders.of(this).get(MyViewModel.class);
myViewModel.getNumber().observe(this, new Observer<Integer>() {
    @Override
    public void onChanged(Integer integer) {
        binding.textView.setText(String.valueOf(integer));
    }
});

binding.button.setOnClickListener(new View.OnClickListener() {
    @Override
    public void onClick(View view) {
        myViewModel.add();
    }
});
```

## 在XML中使用DataBinding

此处省略XML代码。

```java
MyViewModel myViewModel;
ActivityMainBinding binding;

binding = DataBindingUtil.setContentView(this, R.layout.activity_main);
myViewModel = ViewModelProviders.of(this).get(MyViewModel.class);
binding.setData(myViewModel);
binding.setLiftcycleOwner(this);
```

