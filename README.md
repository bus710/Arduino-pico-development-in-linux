# Arduino-pico-development-in-linux

Raspberry Pi Pico is an ARM Cortex M0 based small microcontroller. It has 2 cores so robust to handle multitasking. It has PIO capability so the digital I/O has high accuracy for low level protocols such as SPI/I2C/UART. 

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
# List all the available boards - picow should be in the list
$ arduino-cli board listall
...
Raspberry Pi Pico                   rp2040:rp2040:rpipico                      
Raspberry Pi Pico                   arduino:mbed_rp2040:pico                   
Raspberry Pi Pico W                 rp2040:rp2040:rpipicow 
...
```

## Configure upload permission

```sh
$ arduino-cli core update-index
$ arduino-cli core install arduino:mbed_rp2040 
$ cd ~/.arduino15/packages/arduino/hardware/mbed_rp2040/${VERSION}
$ sudo ./post_install.sh
```
