# Macrokiosk SMS Gateway API SDK

Use this library to integrate with Macrokiosk SMS Gateway API.

## Get Started

1. Install into existing go module

```bash
> go get github.com/berrypay/macrokiosk
```

## Developer Note

Currently this library support automatic encoding of ASCII and UCS2 unicode for specifying the text parameter.
UDH Formatted message support are not yet available. Developer are to pre-encoded the text parameter beforehand for calling the related functions for sending UDH encoded message.

## Status

This library is still in active development. Please check this repo for future updates.
