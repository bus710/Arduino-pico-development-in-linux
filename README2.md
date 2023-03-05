[Back to README.md](README.md#3-sketch-verification-and-c_cpp_propertiesjson)

# How to get the correct header paths

## Manual way

This sub section introduces how to get the correct header paths. Whenever there is a new Arduino-pico version, this should be done for the includePath in c_cpp_properties.json for a sketch.

First, make sure where the board package is. For most of Linux distros, it should be under the user's home directory like this:
- $HOME/.arduino15/packages/rp2040/hardware/rp2040

Then there should be directories per version. For version 3.0.0, the path should look like this:
- $HOME/.arduino15/packages/rp2040/hardware/rp2040/3.0.0

Also there is a file **platform_inc.txt**. Please find and copy the file as it should be edited for the includePath. 
The full path should look like this:
- $HOME/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/lib/platform_inc.txt

If open the file, the contents should look like this (total 63 lines):
```
-iwithprefixbefore/cores/rp2040/api/deprecated-avr-comp/
-iwithprefixbefore/include/pico_base/
-iwithprefixbefore/pico-sdk/lib/tinyusb/src/
-iwithprefixbefore/pico-sdk/src/boards/include
-iwithprefixbefore/pico-sdk/src/common/pico_base/include
...
-iwithprefixbefore/pico-sdk/lib/btstack/platform/embedded
```

The **-iwithprefixbefore** should be replaced with the actual path + version.
For example, the first item should be updated like the second item:
- -iwithprefixbefore/cores/rp2040/api/deprecated-avr-comp/
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/cores/rp2040/api/deprecated-avr-comp",

The prefix is replaced. Also double quotes and comma are added to head and tail to be a json format string. This should be done for all 63 lines and pasted into the c_cpp_properties.json as includePath.

Also, those 3 lines should be added to avoid redlines (should be edited for the path and version, too!):
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/cores/rp2040",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/variants/rpipico",

## Little better way


[Back to README.md](README.md#3-sketch-verification-and-c_cpp_propertiesjson)
