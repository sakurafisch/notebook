# Android

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

## DP尺寸单位

- `dp`在不同密度的屏幕中实际显示比例将保持一致，1dp相当于160dpi屏幕中的1px。

- `dp`一般用于设置控件宽高和图片宽高。


| 屏幕大小 | 1dp | 对应px |
| :-----| ----: | :----: |
| 120dpi | 1dp | 0.75px |
| 160dpi | 1dp | 1px |
| 240dpi | 1dp | 1.5px |
| 320dpi | 1dp | 2px |

## SP尺寸单位

- `sp`是`Scale-independent Pixels`的缩写。

- 字体大小的单位要使用`sp`。

- 它随用户对系统字体大小的设置进行比例缩放。

## 获取屏幕尺寸

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

## DP与PX间的转换

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