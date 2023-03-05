# Arduino-pico-development-in-linux

Raspberry Pi RP2040 is an ARM Cortex M0 based small microcontroller. It has 2 cores so robust to handle multitasking. It has PIO capability so the digital I/O has high accuracy for low level protocols such as SPI/I2C/UART. The Pico W board can talk to other devices without wire so easy to start IoT projects.

This repo is written with a particular combination in mind:
- Debian/Ubuntu Linux as the host
- Raspberry Pi Pico (+W) as the target
- Arduino-pico as the main SDK
- Arduino-cli as the tool to manage Arduino based projects
- Visual Studio Code as the editor

So this doc can be helpful for someone who would like to start with this setup.

<br/>

## Install arduino-cli

Arduino CLI is a command line tool to manage Arduino components (board profiles, libraries, compilers, etc). The detail can be found from the official document (https://arduino.github.io/arduino-cli/0.31/installation/).

It can be installed by:
```sh
$ cd ~
$ mkdir -p Arduino/cli
$ curl -fsSL https://raw.githubusercontent.com/arduino/arduino-cli/master/install.sh | sh
```

The executable binary will be installed under $HOME/Arduino/cli/bin, so add the path to PATH variable in bashrc or zshrc. Then restart terminal.
```sh
export PATH=$PATH:$HOME/Arduino/cli/bin
```

<br/>

## Install arduino-pico via arduino-cli

The arduino-cli command can be used to install basic configuration:
```sh
# Add arduino-pico index to Arduino board manager
$ arduino-cli config init --additional-urls \
    https://github.com/earlephilhower/arduino-pico/releases/download/global/package_rp2040_index.json
# Update Arduino core cache
$ arduino-cli core update-index
# List all the available boards - rpipicow should be in the list
$ arduino-cli board listall
...
Raspberry Pi Pico                   rp2040:rp2040:rpipico                      
Raspberry Pi Pico                   arduino:mbed_rp2040:pico                   
Raspberry Pi Pico W                 rp2040:rp2040:rpipicow 
...
```

<br/>

## Configure upload permission

Although pico boards support UF2 (so firmware update can be done by drag and drop), it can be little cumbersome during development phase. Arduino-cli can put pico boards into bootloader mode when reset (so no need to press the boot button on the board). To allow that, the udev configuration should be applied prior to any sketch uploading.
```sh
$ arduino-cli core update-index
$ arduino-cli core install arduino:mbed_rp2040 
$ cd ~/.arduino15/packages/arduino/hardware/mbed_rp2040/${VERSION}
$ sudo ./post_install.sh
```

<br/>

## Vscode and extensions

Since Vscode is popular, no need to introduce how to install it.
Still few extensions should be installed:
- ms-vscode.cpptools
- ms-vscode.cpptools-extension-pack
- vsciot-vscode.vscode-arduino

<br/>

## Create a new sketch

In terminal, a new Arduino project can be created.
The command below will create a new directory with a sketch file.
```sh
$ arduino-cli sketch new Blink
$ cd Blink
$ cat Blink.ino
void setup() {
}
void loop() {
}
```

<br/>

## Configure the sketch in Vscode

In the Blink directory, open Vscode:
```sh
$ code
```

<br/>

## Vscode-arduino configuration

Once Vscode is opened, let's double check if pico package is installed.
Open the command palette with **Control + Shift + p** and choose **Arduino: Board Manager**. Press the **Refresh Package Indexes** and search for pico if **Raspberry Pi Pico/RP2040** is installed. If not, please press the install button.

<br/>

## Project specific configuration

Now, project specific items should be configured.


<br/>

### 1. The sketch

The sketch file can be updated - it is just an LED blinking example:
```c
#include <Arduino.h>

void setup()
{
    pinMode(LED_BUILTIN, OUTPUT);
}

void loop()
{
    digitalWrite(LED_BUILTIN, HIGH);
    delay(300);
    digitalWrite(LED_BUILTIN, LOW);
    delay(300);
}
``` 

<br/>

### 2. Board type configuration and arduino.json

At this point, we cannot even compile the sketch because the **board type** is not specified yet.
Open the command palette with **Control + Shift + p** and choose **Arduino: Board Config**. There will be **Select your board** dropdown - choose the pico or pico w in the list. The .vscode directory and arduino.json file will be automatically created.

The arduino.json file can be updated with these lines to specify the name of this sketch as well as output directoy:
```json
{   
    ...
    "sketch": "Blink.ino",
    "output": "./build",
    ...
}
```

<br/>

### 3. Sketch verification and c_cpp_properties.json

Now, let's verify the sketch. Open the command palette with **Control + Shift + p** and choose **Arduino: Verify**. This command will do 2 things:
1. it compiles the sketch and creates all the results under the build directory. 
2. it creates a new file **c_cpp_properties.json** for this C/C++ based project. The **includePath** should be updated since it doesn't point the correct header file paths.

The includePath should look like the one in **Blink/.vscode/c_cpp_properties.json**, but the paths can be little different depend on the Arduino-pico version and the actual path in file system. There is this [another doc](README2.md#how-to-get-the-correct-header-paths), which explains how to get the correct header paths.

<br/>

### 4. Surpress Intellisense with settings.json

One thing shdouldn't be forgotten is, even if the includePath gets manually updated, Intellisense can overwrite the includePath. To prevent, add a new file **settings.json** under the .vscode directory with those lines.
```json
{
    "arduino.disableIntelliSenseAutoGen": true,
    // "C_Cpp.errorSquiggles": "disabled" // Add this if redlines are too annoying.
}
```

<br/>

## Upload sketch 

Lastly, let's upload the sketch to the board. As it comes with UF2 as default, Arduino-cli cannot reset it by itself for the first time. 

To get the board's bootloader activated: 
1. simply unplug the board from host (if it was already plugged in).
2. press and hold the boot button on the board during plugging it back. 
3. press **Control + Alt + u**. 
4. A popup will say **Serial port should be selected** (or so). Then pick the right one (like /dev/ttyACM0) and try again **Control + Alt + u**.

Make sure the short cut! It is **Control + Alt + u** (not Control + Shift + p).

Once the uploading is done, the built in LED should blink.

