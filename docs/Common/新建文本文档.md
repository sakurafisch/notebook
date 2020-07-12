## Package an application into a JAR﻿

When the code is compiled and ready, you can package your application in a Java archive (JAR) to share it with other developers. A built Java archive is called an *artifact*.

### Create an artifact configuration for the JAR﻿

1. From the main menu, select **File | Project Structure** Ctrl+Shift+Alt+S and click **Artifacts**.

2. Click ![the Add button](https://resources.jetbrains.com/help/img/idea/2020.1/icons.general.add@2x.png), point to **JAR**, and select **From modules with dependencies**.

3. To the right of the **Main Class** field, click ![the Browse button](https://resources.jetbrains.com/help/img/idea/2020.1/icons.general.openDisk@2x.png) and select the main class in the dialog that opens (for example, **HelloWorld (com.example.helloworld)**).

   IntelliJ IDEA creates the artifact configuration and shows its settings in the right-hand part of the **Project Structure** dialog.

4. Apply the changes and close the dialog.

Gif

![https://resources.jetbrains.com/help/img/idea/2020.1/manifest.png](.\.img\manifest.png)

### Build the JAR artifact﻿

1. From the main menu, select **Build | Build Artifacts**.

2. Point to the created .jar (**HelloWorld:jar**) and select **Build**.

   If you now look at the out/artifacts folder, you'll find your .jar file there.

![The JAR artifact is built](.\.img\jt-jar-built.png)