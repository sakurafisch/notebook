# Activity生命周期

## onCreate()

Activity 开始创建，并且只会被调用一次

## onStart()

可见

Activity 创建完成开始启动，此时的 Activity **已经可见**，但是还没有出现在前台，不能交互，可以理解为已经显示出来了，但是还看不到

## onRestart()

可交互

Activity 重新启动，通常当 Activity 从不可见到可见会执行该声明周期函数，比如用户按下 Home 键或打开了一个新的 Activity 后再回到这个 Activity

## onResume()

Activity 可见，Activity 此时在前台活动，onStart() 执行的时候这个 Activity 还是后台运行，直到 onResume() 这个 Activity 才在前台活动

## onPause()

不可交互

Activity 正在停止，这个 Activity 的 onPause() 执行完毕才能执行新的 Activity 的 onResume()，所以不建议在 onPause() 里面做耗时操作，这样会影响打开新的 Activity 的速度

## onStop()

不可见

Activity 即将停止，onPause() 之后就会立即执行该生命周期函数，可以执行稍微重量级的回收工作。理想状态下 onPause() 后快速回到当前 Activity，onStop() 则不会执行

## onDestory()

Activity 即将销毁，此时可以执行一些资源回收等操作

## 常见情况

- Activity 首次启动：onCreate() -> onStart() -> onResume()
- 返回桌面或打开新的 Activity：onPause() -> onStop()
- 重新回到这个 Activity：onRestart() -> onStart() -> onResume()
- 返回键或者调用 finish() 函数：onPause() -> onStop() -> onDestroy()

## 异常情况

分别有两种情况：

- 系统配置发生变化，例如横竖屏切换
- 系统内存不足杀掉 Activity

这种情况下生命周期函数执行顺序就为：

```
s=>start: 横竖屏切换或者Activity被杀死
e=>end: 执行结束
op1=>operation: onPause()
op2=>operation: onSaveInstanceState()
op3=>operation: onStop()
op4=>operation: onDestroy()
op5=>operation: onCreate()
op6=>operation: onStart()
op7=>operation: onRestoreInstanceState()
op8=>operation: onResume()
s->op1->op2->op3->op4->op5->op6->op7->op8->e`
```

由于 Activity 是因为异常原因才结束的，所以在结束的时候会调用 `onSaveInstanceState()` 函数保存状态，该函数的调用顺序和 `onPause()` 函数并没有既定的时序关系，可能是 `onPause()` 之前也可能是之后，Activity 被重新创建后会调用 `onRestoreInstanceState()` 函数中的 `Bundle` 就是之前保存在的 `onSaveInstanceState()` 中的数据，并且此时 `onCreate()` 中的参数也同样是该 `Bundle` 对象。

Activity 异常结束并重新创建的过程中，系统同时把该 Activity 的视图结构也保存了下来，所以同 Activity 一样，所有的 View 也都具有 `onSaveInstanceState()` 和 `onRestoreInstanceState()` 函数。

