#include <pico/stdlib.h> // For Pico SDK
#include <Arduino.h>
#include <FreeRTOS.h>
#include <btstack_config.h>
#include <BTstackLib.h>
#include <SPI.h>

void setup()
{
    pinMode(LED_BUILTIN, OUTPUT);

    BTstack.setBLEAdvertisementCallback(advertizementCallback);
}

void loop()
{
    digitalWrite(LED_BUILTIN, HIGH);
    delay(300);
    digitalWrite(LED_BUILTIN, LOW);
    delay(300);
}

void advertizementCallback(BLEAdvertisement *adv){

}