# Android

建议按照[应用架构指南](https://developer.android.google.cn/jetpack/docs/guide)优雅地开发。

[安卓版本列表](https://zh.wikipedia.org/wiki/Android%E7%89%88%E6%9C%AC%E5%88%97%E8%A1%A8)

[布局](https://developer.android.com/guide/topics/ui/declaring-layout?hl=zh-cn#kotlin)

[约束布局 training](https://developer.android.com/training/constraint-layout)

[约束布局 reference](https://developer.android.com/reference/android/support/constraint/ConstraintLayout)

[Material Design](https://material.io/develop/android/)

[Fluent Design](https://www.microsoft.com/design/fluent/#/android)

## 设置 JDK 版本

参考 [设置JDK版本](https://developer.android.com/studio/intro/studio-config#jdk)

Android Studio 2.2 及更高版本捆绑提供了最新版本的 OpenJDK，这是我们建议用于 Android 项目的 JDK 版本。要使用捆绑的 JDK，请执行以下操作：

1. 在 Android Studio 中打开您的项目，然后在菜单栏中依次选择 **File > Project Structure**。
2. 在 **SDK Location** 页面中的 **JDK location** 下方，选中 **Use embedded JDK** 复选框。
3. 点击 **OK**。

默认情况下，用于编译项目的 Java 语言版本基于项目的 [`compileSdkVersion`](http://google.github.io/android-gradle-dsl/current/com.android.build.gradle.BaseExtension.html#com.android.build.gradle.BaseExtension:compileSdkVersion)（因为不同版本的 Android 支持不同版本的 Java）。如有必要，您可以通过将以下[`CompileOptions {}`](http://google.github.io/android-gradle-dsl/current/com.android.build.gradle.internal.CompileOptions.html)代码块添加到`build.gradle` 文件来替换此默认 Java 版本：

```groovy
android {
        compileOptions {
            sourceCompatibility JavaVersion.VERSION\_1\_6
            targetCompatibility JavaVersion.VERSION\_1\_6
        }
    }
```

## 支持不同的像素密度

[参考文档](https://developer.android.com/training/multiscreen/screendensities)

### DP尺寸单位

- `dp`在不同密度的屏幕中实际显示比例将保持一致，1dp相当于160dpi屏幕中的1px。

- `dp`一般用于设置控件宽高和图片宽高。


| 屏幕大小 | 1dp | 对应px |
| :-----| ----: | :----: |
| 120dpi | 1dp | 0.75px |
| 160dpi | 1dp | 1px |
| 240dpi | 1dp | 1.5px |
| 320dpi | 1dp | 2px |

### SP尺寸单位

- `sp`是`Scale-independent Pixels`的缩写。

- 字体大小的单位要使用`sp`。

- 它随用户对系统字体大小的设置进行比例缩放。

### 获取屏幕尺寸

```java
// 获取系统管理器
WindowManager manager = (WindowManager) getContext().getSystemService(Context.WINDOW_SERVICE);
// 获取显示器
Display defaultDisplay = manager.getDefaultDisplay();
// 获取屏幕宽度
screenWidth = defaultDisplay.getWidth();
// 获取屏幕高度
screenHeight - defaultDisplay.getHeight();
```

### DP与PX间的转换

公式：px = dp * (dpi / 160)

```java
// 根据手机分辨率从 dp 单位转成 px
public static int dip2px(Context context, float dpValue) {
    final float scale = context.getResources().getDisplayMetrics().density; // 获取屏幕像素密度
    return (int)(dpValue * scale + 0.5f);
}
```

```java
// 根据手机分辨率从 px 单位转成 dp
public static int px2dip(Context context, float pxValue) {
    final float scale = context.getResources().getDisplayMetrics().density;
    return (int)(pxValue / scale + 0.5f);
}
```

## 屏幕适配

### 常用方案

[支持不同的屏幕尺寸](https://developer.android.com/training/multiscreen/screensizes)

[创建布局变体](https://developer.android.com/studio/write/layout-editor#create-variant)

[使用 Resource Manager 管理应用的界面资源](https://developer.android.com/studio/write/resource-manager)

### 把屏幕锁定为竖直状态

在 AndroidManifest.xml 中为相应的 Activity 添加属性 android:screenOrientation="portrait" 即可。

例如：

```xml
<activity android:name=".MainActivity" android:screenOrientation="portrait"></activity>
```

### 为横屏单独创建布局

在 Android Studio 中点击 Create Landscape Variation 即可为横屏创建布局。

转屏后保存数据，举个例子：

```java
@Override
protected void onSaveInstanceState(Bundle outState) {
    super.onSaveInstanceState(outState);
    outState.putString("KEY", textView.getText().toString());
}
```

```java
if (saveInstanceState != null) {
    Stirng str = savedInstanceState.getString("KEY");
    textView.setText(str);
}
```

Activity 的优先级可以分为三级：

- 前台 Activity：可见并可交互，优先级最高
- 可见非前台 Activity：当 Activity 被弹窗遮盖时
- 后台 Activity：已经暂停的 Activity，优先级最低，系统内存不足时优先回收此类 Activity

如果不想在系统配置发生变化时，销毁并重新创建 Activity，可以在 `AndroidManifest.xml` 中配置该 Activity 的属性,从 Android 3.2（API 级别 13）开始，当设备在纵向和横向之间切换时，**屏幕尺寸**也会发生变化。因此，在开发针对 API 级别 13 或更高版本的应用时，若要避免由于设备方向改变而导致运行时重启，则除了 "orientation" 值以外，您还必须添加 "screenSize" 值：

```xml
<activity android:name="Activity"
            android:configChanges="orientation|screenSize">
</activity>
```



## TextView文本视图

TextView是不可编辑的文本视图。

常用属性

- android:gravity 在x轴和y轴的显示方向
- android:hint 内容为空时显示的文本
- android:typeface 字体类型
- android:ellipsize 显示内容超出长度时如何省略[none|start|middle|end|marquee]

```xml
<TextView 
          android:layout_width="match_parent"
          android:layout_height="wrap_content"
          android:gravity="center_horizontal"
          android:text="要显示的文本"
          android:hint="内容为空时显示的文本"
          android:typeface="monospace"
          android:textColor="#ff0000"
          android:textSize="20sp"
          />
```

## EditText可编辑输入框

EditText是文本输入框，它是TextView的子类。

常用属性

- android:lines 固定的行数
- android:maxLines 最大行数
- android:minLines 最小行数
- android:password 是否显示为密码
- android:phoneNumber 内容只能是电话号码
- android:scrollHorizontally 是否可以水平滚动
- android:singleLine 单行模式
- android:maxLength 最大长度

```xml
<EditText 
          android:layout_width="match_parent"
          android:layout_height="wrap_content"
          android:hint="电话号码"
          android:inputType="phone"
          />
```

```xml
<EditText 
          android:layout_width="match_parent"
          android:layout_height="wrap_content"
          android:hint="数字密码"
          android:inputType="numberPassword"
          />
```

## ImageView显示图片的控件

常用属性

- android:src 要显示的图片
- android:maxHeight 最大高度
- android:maxWidth 最大宽度

常用方法

- setImageResource(int resid) 设置图片id
- setImageURI(Uri uri) 设置图片为指定的URI
- setImageBitmap(Bitmap bitmap) 设置Bitmap对象
- setImageDrawable(Drawable drawable) 设置为指定Drawable

```xml
<ImageView 
           android:layout_width="240dp"
           android:layout_height="240dp"
           android:layout_gravity="center_horizontal"
           android:scaleType="fitCenter"
           android:src="@drawable/imgname"
           />
```

## Button按钮

`Button`是`TextView`的子类，可以添加`View.OnClickListener`

```XML
<Button 
        android:id="@+id/btn1"
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:text="显示的名称"
        android:onClick="clickButton"
        />
```

```java
public void clickButton(View view) {
    Toast.makeText(MainActivity.this, "我被点击了", Toast.LENGTH_SHORT).show();
}
```

```java
protected void onCreate(Bunder saveInstanceState){
    super.onCreate(saveInstanceState);
    setContentView(R.layout.activity_main);
    Button btn1 = (Button)findViewById(R.id.btn1);
    btn1.setOnClickListener(new View.OnClickListener(){
       @Override
       public void onClick(View view){
           Toast.makeText(MainActivity.this, "我被点击了", Toast.LENGTH_SHORT).show();
       }
    });
}

```

### ToggleButton状态按钮

`ToggleButton`是`Button`的子类。

- android:textOff 按钮未被选中时显示的文本
- android:textOn 按钮被选中时显示的文本
- 可添加的监听 OnCheckedChangeListener

### RadioGroup

`RadioGroup`是`LinearLayout`的子类。

`RadioButton`是`Button`的子类。

`Radio`注册监听：

- `OnCheckedChangeListener`
- 可获得被选中的`RadioButton`的id

`RadioButton`注册监听：

- `OnCheckedChangeListener`

```xml
<RadioGroup 
            android:layout_width="wrap_content"
            android:layout_height="wrap_content">
            
            <RadioButton 
						android:layout_width="wrap_content"
						android:layout_height="wrap_content"
						android:text="男"
						/>

            <RadioButton 
						android:layout_width="wrap_content"
						android:layout_height="wrap_content"
						android:text="女"
						/>
</RadioGroup>
```

## ProgressBar

```xml
<!--圆形进度条-->
<ProgressBar
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            style="@style/Base.Widget.AppCompat.ProgressBar"
            />
```

```xml
<!--横向进度条-->
<ProgressBar 
             android:id=@+id/progressA
             android:layout_width="match_parent"
             android:layout_height="wrap_content"
             android:max="20"
             android:progress="10"
             style="@style/Base.Widget.AppCompat.ProgressBar.Horizontal"
             />
```

```java
protected void onCreate(Bunder saveInstanceState){
    super.onCreate(saveInstanceState);
    setContentView(R.layout.activity_main);
    ProgressBar pro = (ProgressBar)findViewById(R.id.progress);
    // 设置进度条的最大进度
    pro.setMax(100); // 分为 100 等份
    // 设置进度条的当前进度
    pro.setMax(50); 
}
```

```xml
<SeekBar 
         android:id=@+id/seekbar
         android:layout_width="match_parent"
         android:layout_height="wrap_content"
         android:max="100"
         android:progress="0"
         style="@+id/Base.Widget.AppCompat.ProgressBar.Horizontal"
         />
```

```java
protected void onCreate(Bunder saveInstanceState){
    super.onCreate(saveInstanceState);
    setContentView(R.layout.activity_main);
    SeekBar seekBar = (SeekBar)findViewById(R.id.seekbar);
    seekbar.setOnSeekBarChangeListener(new SeekBar.OnSeekBarChangeListener(){
        @Override
        public void onProgressChanged(SeekBar seekBar, int i, boolean b){
            seekBar.setProgress(i);
        }
        
        @Override
        public void onStartTrackingTouch(SeekBar seekBar){
            
        }
        
        @Override
        public void onStopTrackingTouch(SeekBar seekBar){
            
        }
    });
}
```

```xml
<!--星形进度条-->
<RatingBar 
           android:layout_width="wrap_content"
           android:layout_height="wrap_content"
           />
```

## CheckBox多选按钮

`CheckBox`是`Button`的子类。

XML属性

- android:checked 初始时是否选中

注册监听事件

- OnCheckedChangeListener 选中状态改变时触发监听器

## 时钟组件

### AnalogClock模拟时钟

### DigitalClock数字时钟

`DigitalClock`是`TextVIew`的子类。

## 日期时间选择组件

### DatePicker日期选择器

XML属性

- android:startYear 起始年份
- android:endYear 终止年份

常用方法

- getDayOfMonth() 获取天数
- getMonth() 获取月份
- getYear() 获取年份
- setEnabled(boolean enabled) 控制控件是否可用

更新控件属性值

- updateDate(int year, int month, int day)

初始化属性，设置监听

- init(int year, int month, int day, DatePicker.OnDateChangedListener listener)

### TimePicker时间选择器

常用方法

- getCurrentHour() 获得当前小时
- getCurrentMinute() 获得当前分钟
- is24HourView() 是否是24小时制
- setIs24HourView(boolean b) 设置24小时制
- setCurrentHour(int h) 设置小时
- setCurrentMinute(int m) 设置分钟
- setOnTimeChangedListener(TimePicker.OnTimeChangedListener listener) 监听

```xml
<TimePicker 
            android:id="@+id/timePicker"
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            />
```

```java
TimePicker timepicker;
protected void onCreate(Bunder saveInstanceState){
    super.onCreate(saveInstanceState);
    setContentView(R.layout.activity_main);
    initView();
    timepicker.setHour(10);
    timepicker.setOnTimeChangedListener(new TimePicker.OnTimeChangedListener(){       
        @Override
        public void onTimeChanged(TimePicker view, int hourOfDay, int minute){
            Toast.makeText(TimeActivity.this, "" + minute, Toast.LENGTH_SHORT).show();
        }
    });

    final Handler handler = handleMessage(msg) -> {
        minute++;
        if(minute>=60){
            minute=0;
            hour++;
        }
        timepicker.setMinute(minute);
        timepicker.setHour(hour);
    };
    
    // 时间和定时器的使用
    java.util.Timer timer = new java.util.Timer(true);
    TimerTask timerTask = () -> {handler.sendEmptyMessage(0);};
    timer.schedule(task, 0, 1000);
}

private void initView(){
    timepicker = (TimePicker)findViewById(R.id.timepicker);
}
```

## Adapter适配器

- 数据与UI组件分离
- 相当于MVC架构中的C

### SimpleAdapter

```java
SimpleAdapter adapter = new SimpleAdapter(this, data, R.layout.simpleadapter_item, new String[]{"KEY01", "KEY02", "KEY03"}, new int[]{R.id.textView1, R.id.textView2, R.id.textView3});
```

### ArrayAdapter

```java
ArrayAdapter arrayAdapter = new ArrayAdapter(this, R.layout,simpleadapter_item, R.id.textView1, new String[]{"张三", "李四", "杜甫"})
```

## ListView

```java
ListView listView = new ListView();
listView.setAdapter(adapter);
```

## GridView

## Intent

```java
Intent intent = new Intent();
Bundle bundle = new Bundle();
bundle.putString("name", "tom");
intent.putExtras(bundle);
```

```java
Intent intent = new Intent();
intent.putExtras("name", "tom");
```

putExtra()重载方法

- putExtra(String name, String value)
- putExtra(String name, long value)

```java
// putExtra()方法内部原理
putExtra(String name, String value) {
    if (mExtras == null) mExtras = new Bundle();
    mExtras.putString(name, value);
}
```

### Activity数据回传

```java
// 获取数据的启动方式
StartActivityForResult(intent, requestCode);
```

```java
// 原 Activity 获取新 Activity 传回的数据
onActivityResult(requestCode, resultCode, intent);
```

- request 请求码
- resultCode 结果码
- intent 携带返回数据的intent。
- intent.getExtras().getString("key");

新启动的 Activity 在关闭的时候向前面的Activity设置数据

- setResult(resultCode, intent);

```java
Intent intent = new Intent();
intent.putExtra("string", "字符串数据");
intent.putExtra("int", 110);
setResult(1, intent); // 设置回传数据
// finish();
```

```java
// 获取传递过来的数据
Intent intent= getIntent();
String string = intent.getStringExtra("string");
int anInt = intent.getIntentExtra("int", 0); // 若没有传递，默认值为0
```

```java
// 获取回传数据
// 写在onActivityResult方法中
super.onActivityResult(resultCode, resultCode, data);
rec = data.getStringExtra("string");
// rec = data.getExtras().getString("string");
```

### Component

Component属性明确指定Intent的目标组件的类名称

```java
Intent intent = new Intent();
intent.setComponent(new ComponentName("com.example.otherapp",
                     "com.example.otherapp.MainActivity2"));
startActivity(intent);
```

### 拨打电话

```java
Intent intent = new Intent(Intent.ACTION.DIAL);
intent.setData(Uri.parse("tel:15205205205"));
startActivity(intent);
```

### 发送短信

```java
Intent intent = new Intent(Intent.ACTION_SENDTO);
intent.setData(Uri.parse("smsto:15205205205"));
intent.putExtra("sms_body", "你好");
startActivity(intent);
```

### 播放MP3文件

```java
Intent intent = new Intent();
intent.setAction(Intent.ACTION_VIEW);
Uri uri = Uri.parse("" + "file:///storage/emulated/0/Music/我爱你.mp3");
intent.setDataAndType(uri, "audio/mp3");
startActivity(intent);
```

## 广播

### 广播发送者

- Context.sendBroadcast() 发送普通广播，订阅者都有机会获得并进行处理。
- Context.sendOrderedBroadcast() 发送有序广播。
- BroadcastReceiver.abortBroadcast() 终止广播

### 广播接收者

订阅Intent后，异步接收广播Intent。

```java
// 动态注册
IntentFilter filter = new IntentFilter("ACTION");
registerReceiver(reveiver, filter);
```

```xml
<!--静态注册-->
<!--AndroidManifest.xml-->
<action android:name="ACTION" />
```

常见系统广播Action

- android.intent.action.BATTERY_CHANGED 电池电量改变
- android.intent.action.PHONE_STATE 通话状态改变，比如有电话接入
- android.intent.action.BOOT_COMPLETED 系统启动完毕
- android.intent.action.DATE_CHANGED 日期改变
- android.provider.Telephony.SMS_RECEIVED 收到短信

## Handler实现线程间通信

handler用于实现Activity与Thread/runnable之间的通信，它运行于主UI线程，与子线程通过Message传递数据。

主线程中接收消息

- 创建handler对象，实现handleMessage(msg)方法。

子线程中发送消息

- handler.sendEmptyMessage(key) 
- handler.obtainMessage(key, obj).sendToTarget()

## Server服务

服务没有图形界面，一直在后台运行。

#### 使用步骤：

- 继承Server类

- 在AndroidManifest.xml对服务进行配置

```xml
<service android:name=".SMSService"/>
```

- 调用startService()或bindService()方法启动服务。

#### startService()启动特点

调用者与服务之间没有关联，即使调用者退出了，服务仍然运行。只能调用Context.stopService()方法结束服务。

##### 生命周期

`onCreate` -> `onStartCommand` -> `onDestory`

##### Note

- 如果服务已经开启，不会再调用onCreate()

- 服务只能被停止一次

#### bindService()启动特点

bind方式开启服务，绑定服务，调用者挂了，服务也会跟着挂掉。绑定者可以与服务进行交互。

##### 生命周期

`onCreate` -> `onBind` -> `onUnBind` -> `onDestory`

##### Note

- 绑定服务不会调用 `onStart()` 或者 `onStartCommond`方法。

- 回调方法onServiceConnected传递一个lBinder对象，可以此调用服务方法。

## XML数据操作

### XmlSerializer

[看文档吧](https://developer.android.com/reference/org/xmlpull/v1/XmlSerializer)

### Pull

```java
// 获取解析器对象
XmlPullParser pullParser = Xml.newPullParser();
// 绑定文件
pullParser.setInput(is, "utf-8");
// 基于事件类型解析
int eventType = pullParser.getEventType();
// 获取对应标签的内容
pullParser.nextText();
// 移动到下一个标签
pullParser.next();
```

## SharedPreferences

### Mode常用模式

- MODE_PRIVATE 默认模式，文件只能被本程序访问
- MODE_WORLD_READABLE 允许所有程序读取文件
- MODE_WORLD_WRITEABLE 允许所有程序改写文件

### 存储数据

SharedPreferences是轻量级的存储类，类似于Properties类。

存储数据

```java
SharedPreferences sharedPreferences = getSharedPreferences("filename", MODE_PRIVATE);
// 获取编辑器
SharePreferences.Editor editor = sharedPreferences.edit();
// 设置String类型的值
editor.putString("key1", "value");
// 存储Int类型的值
editor.putInt("key2", 12);
// 提交事物才会生效
editor.commit();
```

数据以XML文件形式存储

存储路径为

```
/data/data/<package_name>/filename
```

存储内容为

```xml
<?xml version='1.0' encoding='utf-8' standalone='yes' ?>
<map>
	<string name="key1">value</string>
    <int name="key2" value="12" />
</map>
```

### 读取数据

```java
// 获取到 SharedPreferences 对象
SharedPreferences sharedPreferences = this.getPreferences(MODE_PRIVATE);
// 获取 String 类型的数据
String name = sharedPreferences.getString("key1", "default_value");
// 获取 int 类型的数据
int age = sharedPreferences.getInt("age", 0);
```

### 访问其他应用的SharedPreferences

访问前提

- 写入方式为：MODE_WORLD_READABLE

使用方式

```java
try {
    // 获取对应应用的 Context
    Context = context = createPackageContext("包名", CONTEXT_IGNORE_SECURITY);
} catch (PackageManager.NameNotFoundException e) {
    e.printStackTrace();
}
SharedPreferences sharePreferences = other.getSharedPreferences("xml文件名")
```

## SD Card存储配置

在电脑可模拟创建SD卡

- cd到Android SDK安装路径的tools目录，输入指令：

```cmd
mksdcard 2048M D:\sdcard.img
```

SD卡操作权限配置

```xml
<!--在SD Card中创建与删除文件权限-->
<uses-permission android:name="android.permission.MOUNT_UNMOUNT_FILESYSTEMS" />
<!--往SD Card写入数据权限-->
<uses-permission android:name="android.permission.WRITE_EXTERNAL_STORAGE" />
```

SD Card读写操作

```java
// 获取SD卡的状态
Environment.getExternalStorageState(); 
// 判断SD卡是否存在,如果存在并可读写，则返回 MEDIA_MOUNTED
Environment.getExternalStorageState().equals(Environment.MEDIA_MOUNTED);
// 获取SD卡目录
Environment.getExternalStorageDirectory();
// 或
File saveFile = new File("/sdcard/abc.txt");
```

## SQLite

可以使用 [SQLiteOpenHelper](https://developer.android.com/reference/android/database/sqlite/SQLiteOpenHelper) 类来创建和更新SQLite数据库。

首先需要写一个类继承 SQLiteOpenHelper 类，然后参考下文：

### 创建数据库

```java
// 创建数据库表格
@Override
public void onCreate(SQLiteDatabase sqLiteDatabase) {
    sqLiteDatabase.execSQL("create table student(" + "_id Integer primary key autoincrement" + ", name varchar(20), age Integer)")
}
```

### 更新数据库版本

```java
// 更新数据库版本
@Override
public void onUpgrade(SQLiteDatabase sqLiteDatabase, int i, int j) {
}
```

### 插入数据

- insert(String 表名, String 非空列名, ContentValues values);

```java
// 获取数据库
SQLiteDatabase sqLiteDatabase = help.getWritableDatabase();
// 封装数据
ContentValues contentValues = new ContentValues();
contentValues.put("name", name);
contentValues.put("age", age);
// 插入操作
long insert =sqLiteDatabase.insert(table, null, contentValues);
// 关闭数据库
sqLiteDatabase.close()
```

### 删除操作

- delete(String 表名, String where, String[] args);

```java
SQLiteDatabase sqLiteDatabase = help.getWritableDatabase();
// 删除数据
int delete = sqLiteDatabase.delete(table, "name=?", new String[]{name});
sqLiteDatabase.close
```

### 修改数据

```java
SQLiteDatabase sqLiteDatabase = help.getWritableDatabase();
// 封装数据
ContentValues contentValues = new ContentValues();
values.put("age", age);
// 修改数据
int update = sqLiteDatabase.update(table, values, "name=?", new String[]{name});
sqLiteDatabase.close
```

### 查询操作

```java
SQLiteDatabase sqLiteDatabase = help.getReadableDatabase();
Cursor cursor = sqLiteDatabase.query(table, new String[]{"_id", "name", "age"}, null, null, null, null, null);
while (cursor.moveToNext()) { // 判断是否存在数据
    User user = new User();
    int id = cursor.getInt(0);
    user.setId(id);
    lists.add(user);
}
cursor.close();
sqLiteDatabase.close;
return lists;
```

## ContentProvider

四大基本组件之一，主要用于实现不同应用间的数据共享。

### 主要属性 URI

content://hx.android.text.myprovider/tablename/#

- hx.android.text.myprovider 主机名
- tablename 路径，即要操作的表名
- \# 需要获取的记录的ID

### ContentResolver

用于获取ContentProvider数据

```java
// 获取ContentResolver对象
ContentResolver contentResolver = getContentResolver();
```

常用方法

- query(Uri uri, String[] cols, String where, String[] args, String orderBy);
- insert(Uri uri, ContentValues contentValues);
- update(Uri uri, ContentValues contentValues, String where, String[] args);
- delete(Uri uri, String where, String[] args);

## 网络操作

### HttpUrlConnection

java标准类，没有封装，方便扩展。

```java
URL url = new URL("http://winnerwinter.com?username=yourusername");
HttpUrlConnection httpUrlConnection = (HttpUrlConnection)url.openConnection();
httpUrlConnection.setRequestMethod("GET");
if(httpUrlConnection.getResponseCode() == 200){
    httpUrlConnection.getInputStream();
}
```

### HttpClient（已弃用）

HttpClient封装了http的header、参数、body、response等。

于 Android 6.0 弃用。

## JSON数据解析

### JSONObject对象解析

```java
JSONObject jsonObject = new JSONObject(str);
jsonObject.getString("name");
jsonObject.getInt("age");
jsonObject.getJSONObjecj("dept");
```

### JSONArray数组解析

```java
JSONArray jsonArray = object.getJSONArray("person");
JSONObject jsonObject = array.getJSONObject(i);
```

### 第三方JSON解析框架

- 阿里巴巴的Fastjson
- 谷歌的Gson
- jackson

## 加载网络图片

### 网络获取输入流转换为图片

```java
InputStream inputStream = httpUrlConnection.getInputStream();
Bitmap bitmap = BitmapFactory.decodeStream(inputStream);
```

### 三级缓存策略

内存 - 文件 - 网络

当根据URL向网络拉取图片的时候，先从内存中找，如果内存中没有，再从缓存文件中找，如果缓存中也没有，再从网络上通过HTTP请求拉取图片。

### 大图和多图加载处理

加载大图片出现的问题

- 在加载较大的图片到内存时，程序可能挂掉并报COM的错误。

问题产生的原因

- Android系统App的每个进程或者每个虚拟机有最大内存限制，如果申请的内存资源超过这个限制，系统就会抛出COM错误。

大图COM问题的解决方式

- 压缩图片

```java
// 加载到内存前
// 先算出该bitmap的大小
// 然后通过适当调节采样频率使得加载的图片刚好
Options options = new Options();
options.inJustDecodeBounds = true(false);
BitmapFactory.decodeFile(pathName, options)
```

- 采用低内存占用量的编码方式

  Bitmap.Config.ARGB_4444 比 Bitmap.Config.ARGB_8888更省内存。

- 采用第三方框架，比如 picasso、imageloader、glide、frsco

## Frame帧动画

Frame动画是一系列图片按照一定的顺序展示的过程，和放电影的机制很相似。

实现方式

- 在XML中配置
- Java代码实现

## 补间动画

补间动画无需逐一定义每一帧，只需定义开始、结束的帧和指定动画持续时间。

补间动画有四种，它们均为Animation抽象类的子类

- AlpahAnimation（透明度，0~1）

```java
AlphaAnimation alphaAnimation = new AlphaAnimation(0, 1); // 参数为透明度始末数值
alphaAnimation.setDuration(1000);  // 播放时间
alphaAnimation.setFillAfter(true); // 停留在最后一帧
image.startAnimation(alphaAnimation); // 启动动画
```

- ScaleAnimation（大小缩放，X、Y轴缩放，还包括缩放中心pivotX、pivotY）

```java
ScaleAnimation scaleAnimation = new ScaleAnimation(1f, 2, 1f, 1.5f); // x轴变为原来的2倍数，y轴变为原来的1.5倍
scaleAnimation.setDuration(1000); // 持续1s
scaleAnimation.setFillAfter(true); // 停留在最后一帧
image.startAnimation(scaleAnimation); // 启动动画
```

- TranslationAnimation（位移，X、Y轴位移）

```java
TranslateAnimation translateAnimation = new TranslateAnimation(0, 100, 0, 0); // 参数依次为x轴起点终点、y轴起点终点
translateAnimation.setDuration(1000); // 持续时间
translateAnimation.setFillAfter(true); // 停留在最后一帧

```

- RotateAnimation（旋转，包括缩放中心pivotX、pivotY）

```java
RotateAnimation rotateAnimation = new RotateAnimation(0, -60); // 参数依存为开始位置和旋转度数
rotateAnimation.setDuration(1000);
rotateAnimation.setFillAfter(false);
image.startAnimation(rotateAnimation);
```

## 属性动画

### 继承关系

- Animator -> ValueAnimator -> ObjectAnimator

### 使用步骤

1. 调用ObjectAnimator的静态工厂方法创建动画(ofInt、ofFloat、ofObject)
2. 调用SetXxx()设置动画持续时间、插值方式、重复次数等。
3. 动画的监听事件
4. 调用Animator对象的start()方法启动动画

```java
// 透明度改变动画
ObjectAnimator.ofFloat(image, "alpha", 0, 1).start();
```

```java
// 平移动画
ObjectAnimator.ofFloat(image, "translationX", 0, 100).start();
```

```java
// 旋转动画
ObjectAnimator.ofFloat(image, "scaleX", 1, 2).start();
```

```java
// 背景变化动画
ValueAnimator valueAnimator = ObjectAnimator.ofInt(image, "backgroundColor", 0x979b9a, 0xF9AC07);
valueAnimator.setEvaluator(new ArgbEvaluator());
valueAnimator.setRepeatCount(ValueAnimator.INFINITE);
valueAnimator.setRepeatMode(ValueAnimator.REVERSE);
valueAnimator.start();
// 动画监听器
valueAnimator.addListener(new Animator.AnimatorListener(){
    @Override
    public void onAnimationStart(Animator animator) {
        
    }
    // 还可以重载更多方法
});
```

## MediaPlayer

MediaPlayer 包含了 Audio 和 Video 的播放功能，在 Android 的界面上，Music 和 Video 两个应用程序都是调用 MediaPlayer 来实现的。

### 使用步骤

1. 创建 MediaPlayer 对象
2. 调用 setDataSource() 方法设置音视频文件路径
3. 调用 prepare() 方法使 MediaPlayer 进入准备状态
4. 调用 start() 方法播放音视频

### Media要播放的媒体文件来源

- 用户在应用中事先自带的 resource 资源

```java
MediaPlayer.create(this, R.raw.test);
```

- 存储在SD卡或者其他文件路径下的媒体文件

```java
mediaPlayer.setDataSource("/sdcard/path/to/media/test.mp3");
```

- 网络上的媒体文件

```java
mediaPlayer.setDataSource("url");
```

## 弹框提示组件

### Dialog普通对话框

```java
public void normalDialog(View view){
    AlertDialog.Builder builder = new AlertDialog.builder(this);
    mNormalDialog = builder.setTitle("我是普通对话框")
        .setPositiveButton("确认", new DialogInterface.OnClickListener(){
            @Override
            public void onClick(DialogInterface dialogInterface, int i){
                Toast.makeText(AlertDialogActivity.this, "点击确认", Toast.LENGTH_SHORT).show();
            }
        })
        .setNegativeButton("取消", new DialogInterface.OnClickListener(){
            @Override
            public void onClick(DialogInterface dialogInterface, int i){
                Toast.makeText(AlertDialogActivity.this, "点击取消", Toast.LENGTH_SHORT).show();
            }
        }).create();
    mNormalDialog.show();
}
```

### Dialog列表对话框

```java
public void lieBiaDialog(View view) {
    AlertDialog.Builder = new AlertDialog.builder(this);
    mLieBiaoDialog = builder.setTitle("列表对话框")
        .setItems(items, new OnClick(DialogInterface dialogInterface, int i){
            @Override
            public void onClick(DialogInterface dialogInterface, int i){
                Toast.makeText(AlertDialogActivity.this, "点击了" + i + "次", Toast.LENGTH_SHORT).show();
            }
        }).create();
    mLieBiaoDialog.show();
}
```

### Dialog单选对话框

```java
public void singleDialog(View view) {
    AlertDialog.Builder builder = new AlertDialog.Builder(this,);
    mSingleDialog = builder.setTitle("单选对话框");
    builder.setTitle("单选对话框")
        .setSingleChoiceItems(items, 1, new DialogInterface.OnClickListener(){
            @Override
            public void onClick(DialogInterface dialogInterface, int i) {
                Toast.makeText(AlertDialogActivity.this, "点击了" + item[which] + "个", Toast.LENGTH_SHORT).show();
            }
        })
        .setPositiveButton("确认", new DialogInterface.OnClickListener(){
            @Override
            public void onClick(DialogInterface dialogInterface, int i) {
                mSingleDialog.dismiss();
            }
        })
        .create();
    mSingleDialog.show();
}
```

### Dialog复选对话框

```java
public void multiDialog{
	AlertDialog.Builder builder = new AlertDialog.Builder(this,);
	mCheckDialog = build.setTitle("多选对话框")
    	.setMultiChoiceItems(items, checkItems, new DialogInterface.OnMultiChoiceClickListener(){
        	@Override
        	public void onClick(DialogInterface dialogInterface, int i){
            	// do something
        	}
    	}).create();
	mCheckDialog.show();
}
```

### Dialog自定义对话框

```java
// do it yourself
```

### Notification通知提示栏

Notification是在应用的常规界面之外展示消息，当app让系统发送一个消息的时候，消息首先以图表的形式显示在通知栏。

Notification的使用

- 系统默认通知
- 自定义通知

### Popwindow弹出菜单

PopupWindow的作用与Notification类似，但它可以指定显示的位置。

必须设置：

- View contentView
- int width
- int height

常用方法：

- showAsDropDown(View anchor)
- showAsDropDown(View anchor, int xoff, int yoff)
- showAsLocation(View parent, int gravity, int x, int y)

## WebView

- 加载网址

```java
webview.loadURL(https://cn.bing.com);
```

- 加载本地的HTML文件

```java
webview.loadURL("file:///android_asset/file_name.html");
```

- webview与js+html交互