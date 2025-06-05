+++
draft = true
title = "Installing NixOS on the Steam Deck"
date = 2024-01-07
+++

# Why

I like my Steam Deck. It's a portable and powerful device to play video games with. So powerful, that I sold my desktop computer and only use my Steam Deck to game with. To really make use of the Steam Deck, I want a powerful package manager and a operating system I am somewhat comfortable with.  
Also I just wanted an execuse to tinker :)

# Installing NixOS

## What do we need?

- Steam Deck
- MicroSD card
- A keyboard to connect to the Steam Deck

## The boring part

First we need to download NixOS. Just navigate to [this](https://nixos.org/download) page and download the iso.  
When the iso is downloaded, we can flash it to the microSD card. To do so, we can utilize the handy cp utility. Just copy the iso to the microSD card. As I used a usb adapter, I had to do the following:
```bash
sudo cp nixos-minimal-23.11.2596.c1be43e8e837-x86_64-linux.iso "/dev/disk/by-id/usb-Generic_MassStorageClass_000000001536-0:0"
```
To boot from the microSD card, hold down the "Volume Down" key and press the power button once, **do not hold it.**

## Configuring WiFi

Unfortunatly NetworkManager isn't preinstalled and configured on the NixOS installer, therefore we have to use wpa_supplicant.  
First, enable wpa_supplicant with `sudo systemctl start wpa_supplicant`, then start `wpa_cli`.

```bash
> add_network
0
> set_network 0 ssid "WIFI-NAME"
> set_network 0 psk "PASSWORD"
> enable_network 0
> save_config
> quit
```
## SSH

SSH is enabled by default. You just have to give the `nixos` a password with `passwd`.
```
passwd
Changing password for nixos.
New password: 123
Retype new password: 123
passwd: password updated successfully
```
Now, use `ip addr` to get the ip adress from the Steam Deck and connect with ssh.
```bash
ssh nixos@ip-adress
```

## Just follow the official manual :)

Now you can just follow the [official manual](https://nixos.org/manual/nixos/stable/#sec-installation-manual-partitioning) from NixOS
