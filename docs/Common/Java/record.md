# record

jdk14 preview, jdk16 release.

```java
import java.io.Serializable;
import java.util.Objects;

public record CardRecord(int year, String make, String model) implements Serializable {
    public CardRecord {
        Objects.requireNonNull(make);
    }
}
```

