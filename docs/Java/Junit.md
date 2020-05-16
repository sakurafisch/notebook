# Junit

举个🌰

```java
public class Calculater {
	public int add(int a, int b) {
    	return a + b;
	}
}
```

```java
public class CalculaterTest{
    @Before
    public void init() {
        
    }
    
    @After
    public void close() {
        
    }
    
	@Test
	public void testAdd() {
    	Calculater calculater = new Calculater();
    	int result = calculater.add(1, 2);
    	Assert.assertEquals(3, result);
	}
}
```

