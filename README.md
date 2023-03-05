# Arduino-pico-development-in-linux

Raspberry Pi RP2040 is an ARM Cortex M0 based small microcontroller. It has 2 cores so robust to handle multitasking. It has PIO capability so the digital I/O has high accuracy for low level protocols such as SPI/I2C/UART. The Pico W board has good network libraries so easy to start IoT projects.

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

## Configure upload permission

Although pico boards supports UF2 (so firmware update can be done by drag and drop), it can be little cumbersome during development phase. Arduino-cli can put pico boards into bootloader mode when reset (so no need to press the boot button on the board), but the udev configuration should be applied prior to any sketch uploading.
```sh
$ arduino-cli core update-index
$ arduino-cli core install arduino:mbed_rp2040 
$ cd ~/.arduino15/packages/arduino/hardware/mbed_rp2040/${VERSION}
$ sudo ./post_install.sh
```

## Vscode and extensions

Since Vscode is popular, no need to introduce how to install it.

Few extensions should be installed:
- ms-vscode.cpptools
- ms-vscode.cpptools-extension-pack
- vsciot-vscode.vscode-arduino

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

## Configure the sketch in Vscode

In the Blink directory, open Vscode:
```sh
$ code
```

### Vscode-arduino configuration

Once Vscode is opened, let's double check if pico package is installed.
Open the command palette with **Control + Shift + p** and choose **Arduino: Board Manager**. Press the **Refresh Package Indexes** and search for pico if **Raspberry Pi Pico/RP2040** is installed. If not, please press the install button.

### Project specific configuration

Now, project specific items should be configured.

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

Now, let's verify the sketch. Open the command palette with **Control + Shift + p** and choose **Arduino: Verify**. This command will do 2 things:
1. it compiles the sketch and creates all the results under the build directory. 
2. it creates a new file **c_cpp_properties.json** for this C/C++ based project. The **includePath** should be updated since it doesn't point the correct header file paths.

The includePath should look like this, but the path can be little different as Arduino-pico gets updated. 

<details>
  <summary><b>The includePath - 66 lines</b></summary>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/cores/rp2040", </li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/cores/rp2040/api",
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/cores/rp2040/api/deprecated-avr-comp",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/include/pico_base",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/variants/rpipico",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/lib/tinyusb/src",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/boards/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/common/pico_base/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/common/pico_binary_info/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/common/pico_bit_ops/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/common/pico_divider/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/common/pico_stdlib/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/common/pico_sync/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/common/pico_time/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/common/pico_usb_reset_interface/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/common/pico_util/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2040/hardware_regs/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2040/hardware_structs/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/cmsis/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/cmsis/stub/CMSIS/Core/Include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/cmsis/stub/CMSIS/Device/RaspberryPi/RP2040/Include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_adc/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_base/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_claim/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_clocks/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_divider/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_dma/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_exception/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_flash/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_gpio/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_i2c/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_interp/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_irq/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_rtc/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_pio/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_pll/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_pwm/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_resets/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_spi/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_sync/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_timer/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_uart/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_vreg/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_watchdog/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_xosc/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_async_context/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_bootrom/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_btstack/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_cyw43_arch/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_cyw43_driver/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_double/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_float/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_int64_ops/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_lwip/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_multicore/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_platform/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_printf/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_runtime/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_rand/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_stdio/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_stdio_uart/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_unique_id/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/lib/cyw43-driver/src",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/lib/lwip/src/include",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/lib/btstack/src",</li>
    <li>"~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/lib/btstack/platform/embedded"</li>
</details>

</br>

As an example, let's see this line - **~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/cores/rp2040**.

1. There is this long prefix: ~/.arduino15/packages/rp2040/hardware/rp2040/
2. Then the version follows: 3.0.0
3. This is the actual path: /cores/rp2040

To manually update the **includePath**, 
1. delete all the auto-generated paths in the includePath.
2. find the actual path from **~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/lib/platform_inc.txt**. The problem is, each line has this **-iwithprefixbefore** prefix. This prefix should be replaced with the other prefix and version in the includePath of the c_cpp_properties.json file. 

<!-- TODO: Create a go project to automate this -->

One thing shdouldn't be forgotten is, Intellisense can overwrite this hard work on the includePath. To prevent it to be re-generated, add a new file **settings.json** under the .vscode directory.
```json
{
    "arduino.disableIntelliSenseAutoGen": true,
    // "C_Cpp.errorSquiggles": "disabled" // Add this if redlines are too annoying.
}
```

## Upload sketch 

Lastly, let's upload the sketch to the board. As it comes with UF2 as default, Arduino-cli cannot reset it by itself for the first time. To get the board's bootloader activated, simply unplug the board from host (if it was plugged in). Then press and hold the boot button on the board during plugging it back. Press **Control + Alt + u**. A popup will say **Serial port should be selected** (or so). Then pick the right one (like /dev/ttyACM0) and try again **Control + Alt + u**. 

Make sure the short cut! It is **Control + Alt + u** (not Control + Shift + p).