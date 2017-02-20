These are the Protocol Buffer definitions for routing between TTNode, TTGate, and TTServe.  They define the highly-compressed "on-air" data format optimized initially for the highly-constrained LoRaWAN transport.

There are two language bindings: golang for the gateway and service, and clang for the device.

Google's protocol buffers tools are used for the gateway and service, and "nanopb" for the device.