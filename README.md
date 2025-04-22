# GoPro live-streaming without Internet connectivity

[![Go Reference](https://pkg.go.dev/badge/github.com/MatthiasKunnen/goprolocal.svg)](https://pkg.go.dev/github.com/MatthiasKunnen/goprolocal)

This program responds to a GoPro's internet connectivity checks so that it successfully connects to
a WiFi network without internet access. This allows live-streaming to a local server. 

## What does it do?

When attempting to start a livestream from the QUIK GoPro Android app, and selecting the Wi-Fi
network, the GoPro will connect to that network and perform an HTTP request to api.gopro.com to
determine network connectivity.
However, sometimes, you want to stream to a local server without internet access which is prevented
by this connectivity check.

The `goprolocal` binary is a simple HTTP server that responds positively to the connectivity
checks.
This results in livestreams succeeding with; no internet access, in case api.gopro.com is ever
offline, or worse, gets discontinued.

## Install
`go install github.com/MatthiasKunnen/goprolocal/cmd/goprolocal`

## How to use
You will need to control the WiFi router that the GoPro connects to.

1. Set up DHCP pointing to a DNS under your control.
1. Run `goprolocal`.
1. Set up the DNS to resolve `A api.gopro.com` to the IP of the server running `goprolocal`.

## What does the GoPro communicate?
A pcapng capture of a GoPro 13 Black can be found in the zstd compressed
[`gopro.pcapng.zstd`](./gopro.pcapng.zstd) file.
This file can be inspected using Wireshark.
