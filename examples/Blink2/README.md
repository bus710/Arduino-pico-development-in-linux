# Blink

## To clean/build/upload in terminal

```sh
$ make clean && make build && make upload 
```

## To use BTstack

https://bluekitchen-gmbh.com/btstack-arduino/examples/generated/#section:LEPeripheral

Also, don't forget to enable BT stack from the Arduino Board Configuration (IP/Bluetooth Stack: IPv4 + Bluetooth). 
The arduino.json should have the parameter "ipbtstack=ipv4btcble".

If there is a compilation error like below, add **--board-options ipbtstack=ipv4btcble** as arduino-cli compile option.
```
This library needs Bluetooth enabled.  Use the 'Tools->IP/Bluetooth Stack' menu in the IDE to enable it.
```