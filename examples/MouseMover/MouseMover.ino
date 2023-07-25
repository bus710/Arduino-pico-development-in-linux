

#include <pico/stdlib.h> // For Pico SDK
#include <Arduino.h>
#include <FreeRTOS.h>
#include <Mouse.h>
// #include <hardware/pio.h>

#include <Adafruit_NeoPixel.h>
// #include <rp2040_pio.h>

#define PIN 25
#define NUMPIXELS 1
Adafruit_NeoPixel pixels(NUMPIXELS, PIN, NEO_GRB + NEO_KHZ800);

int cnt = 0;
bool blink = false;

void setup()
{
    Mouse.begin();
    pixels.begin();
}

void loop()
{
    // Blink the built-in LED once every 10 seconds.
    // blink = !blink;
    // digitalWrite(LED_BUILTIN, blink);

    // pixels.clear();
    delay(9900); // 9.9 sec
    pixels.setPixelColor(0, pixels.Color(0, 4, 0));
    pixels.show();

    delay(100);
    pixels.setPixelColor(0, pixels.Color(0, 0, 0));
    pixels.show();


    // Move cursor once every 240 seconds (about 4 mins)
    cnt += 1;
    if (cnt > 24)
    {
        cnt = 0;
        Mouse.move(1, 2, 0);
    }
}
