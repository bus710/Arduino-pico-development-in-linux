[Back to README.md](README.md#3-sketch-verification-and-c_cpp_propertiesjson)


As an example, let's see this line - **~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/cores/rp2040**.

1. There is this long prefix: ~/.arduino15/packages/rp2040/hardware/rp2040/
2. Then the version follows: 3.0.0
3. This is the actual path: /cores/rp2040

To manually update the **includePath**, 
1. delete all the auto-generated paths in the includePath.
2. find the actual path from **~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/lib/platform_inc.txt**. The problem is, each line has this **-iwithprefixbefore** prefix. This prefix should be replaced with the other prefix and version in the includePath of the c_cpp_properties.json file. 




- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/cores/rp2040",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/cores/rp2040/api",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/cores/rp2040/api/deprecated-avr-comp",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/include/pico_base",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/variants/rpipico",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/lib/tinyusb/src",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/boards/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/common/pico_base/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/common/pico_binary_info/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/common/pico_bit_ops/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/common/pico_divider/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/common/pico_stdlib/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/common/pico_sync/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/common/pico_time/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/common/pico_usb_reset_interface/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/common/pico_util/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2040/hardware_regs/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2040/hardware_structs/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/cmsis/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/cmsis/stub/CMSIS/Core/Include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/cmsis/stub/CMSIS/Device/RaspberryPi/RP2040/Include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_adc/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_base/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_claim/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_clocks/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_divider/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_dma/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_exception/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_flash/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_gpio/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_i2c/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_interp/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_irq/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_rtc/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_pio/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_pll/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_pwm/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_resets/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_spi/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_sync/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_timer/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_uart/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_vreg/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_watchdog/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/hardware_xosc/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_async_context/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_bootrom/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_btstack/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_cyw43_arch/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_cyw43_driver/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_double/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_float/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_int64_ops/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_lwip/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_multicore/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_platform/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_printf/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_runtime/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_rand/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_stdio/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_stdio_uart/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/src/rp2_common/pico_unique_id/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/lib/cyw43-driver/src",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/lib/lwip/src/include",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/lib/btstack/src",
- "~/.arduino15/packages/rp2040/hardware/rp2040/3.0.0/pico-sdk/lib/btstack/platform/embedded"


[Back to README.md](README.md#3-sketch-verification-and-c_cpp_propertiesjson)