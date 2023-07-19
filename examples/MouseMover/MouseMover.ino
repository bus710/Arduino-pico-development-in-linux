#include <pico/stdlib.h> // For Pico SDK
#include <Arduino.h>
#include <FreeRTOS.h>
#include <Mouse.h>

int cnt = 0;
bool blink = false;

void setup()
{
    Mouse.begin();
}

void loop()
{
    // Blink the built-in LED once every 2 seconds.
    blink = !blink;
    digitalWrite(LED_BUILTIN, blink);
    delay(1000);

    // Move cursor once every 240 seconds (4 mins)
    cnt += 1;
    if (cnt > 240)
    {
        cnt = 0;
        Mouse.move(1, 2, 0);
    }
}
