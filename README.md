This library is a fork of https://github.com/imc-trading/ifwatch/tree/master/module

The purpose of this library is to decode the programmings of QSFP / QSFP+ / QSFP28 / SFP / SFP+ / SFP28 transceivers according to SFF8636 and SFF8079 standards (although it is currently mainly focused on QSFP / SFF8636)

It is work in progress and there are absolutely no guarantees regarding API stability.

# How to use
The command line interface is essentially `<binaryname> <8636 | 8636_upper00h | 8079> <file>`

- `binaryname` is the name of the binary or `go run .`
- `<8636 | 8636_upper00h | 8079>` are the operation modes
    - `8636`: Decodes a full 256 byte SFF8636 transceiver dump (page 00)
    - `8636_upper00h`: Decodes the upper 128 bytes of a SFF8636 transceiver dump (e.g. for FS.com transceiver programmings) - That means 128 bytes of zeros are prepended to the file
    - `8079`: Decodes a full 256byte SFF8079 transceiver dump

