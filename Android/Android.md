# Android

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

