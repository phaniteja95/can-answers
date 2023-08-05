#!/usr/bin/env bash

set -eux

#install qemu
sudo apt install qemu

#Fetch the linux image
if [ ! -f bzImage-5.17.3 ]
then
	wget https://l4re.org/download/Linux-kernel/x86-64/bzImage-5.17.3
fi

#Compile our init program and put it into the initramfs directory
gcc hello_world.c -o hello_world
sudo cp hello_world /usr/share/initramfs-tools/init

#Create ramfs
mkinitramfs -o ramdisk.img

#Boot linux
qemu-system-x86_64 -kernel bzImage-5.17.3 -nographic -append "console=ttyS0" -initrd ramdisk.img -m 1024

