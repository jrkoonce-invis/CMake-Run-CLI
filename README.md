# CMake-Run-CLI
A little go program that will compile your C/C++ code with CMake into a build/ folder and runs it with one simple command!

## Purpose
I'm trying to learn vim and get into the habit of going through my entire development workflow on the command line. This means I have to use the CMake CLI, which takes multiple commands to build and test your code every time you make changes. Because of this hassle, and my desire to learn a little bit more about Go, I created a little program which does this for me :)
- You must have a CMakeLists.txt file which specifies the project name
- The programs uses os/exec in the standard go library which helps to run the bash commands
- Thanks to the bufio stdout pipe (also in the standard library), you can recieve output form the cmake project in realtime

**Note**: This is simply just for personal development purposes and doesn't have checks for os. Go projects using os/exec can only be used on Unix/Linux systems
