# **SETUP AND INSTALLATION** 

___

* **NOTE: You have have to edit the setupEnvironment script to allow the program to compile in your home directory because you need to tell Golang where the /bin directory in the workspace**

1. **To Compile the Program:** 

   ```bash
   ./setupEnvironment
   ```

   What this script does that it tells where the bin folder is, so modify the script to allow it to point to wherever the code is. If you want to currently look at your current go environment variables, open a terminal and type

   ```bash
   go env
   ```

   

   Afterwards it cd's into the src code and runs "go install" to compile the program into the bin directory.

2. **To run the Program:**

   ```bash
   # Go into the /bin directory
   ./RSA
   ```

   

