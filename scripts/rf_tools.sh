#!/bin/bash

function mirage_soft_install() {
	goodecho "[+] Installing Mirage"
	[ -d /root/thirdparty ] || mkdir /root/thirdparty
	cd /root/thirdparty
	installfromnet "git clone https://github.com/RCayre/mirage"
	cd mirage/
	python3 setup.py install
}

# RFID package

function proxmark3_soft_install() {
	goodecho "[+] Installing proxmark3 dependencies"
	installfromnet "apt-fast install -y --no-install-recommends git ca-certificates build-essential pkg-config libreadline-dev"
	installfromnet "apt-fast install -y  gcc-arm-none-eabi libnewlib-dev qtbase5-dev libbz2-dev liblz4-dev libbluetooth-dev libpython3-dev libssl-dev libgd-dev"
	goodecho "[+] Installing proxmark3"
	[ -d /rftools ] || mkdir /rftools
	cd /rftools
	installfromnet "git clone https://github.com/RfidResearchGroup/proxmark3.git"
	cd proxmark3/
	make clean && make -j$(nproc)
}

function libnfc_soft_install() {
	goodecho "[+] Installing libnfc dependencies"
	installfromnet "apt-fast install -y autoconf libtool libusb-dev libpcsclite-dev build-essential pcsc-tools"
	goodecho "[+] Installing libnfc"
	installfromnet "apt-fast install -y libnfc-dev libnfc-bin"
}

function mfoc_soft_install() {
	goodecho "[+] Installing mfoc"
	installfromnet "apt-fast install -y mfoc"
}

function mfcuk_soft_install() {
	goodecho "[+] Installing mfcuk"
	installfromnet "apt-fast install -y mfcuk"
}

# Wi-Fi Package
function common_nettools() {
	echo apt-fast macchanger/automatically_run  boolean false | debconf-set-selections
	installfromnet "apt-fast install -y -q macchanger"
	echo apt-fast wireshark-common/install-setuid boolean true | debconf-set-selections
	installfromnet "apt-fast install -y -q tshark"
}


function aircrack_soft_install() {
	goodecho "[+] Installing aircrack-ng"
	installfromnet "apt-fast install -y aircrack-ng"
}

function reaver_soft_install() {
	goodecho "[+] Installing reaver"
	installfromnet "apt-fast install -y reaver"
}

function bully_soft_install() {
	goodecho "[+] Installing bully"
	installfromnet "apt-fast install -y bully"
}

function pixiewps_soft_install() {
	goodecho "[+] Installing pixiewps"
	installfromnet "apt-fast install -y pixiewps"
}

function Pyrit_soft_install() {
	goodecho "[+] Installing Pyrit"
	installfromnet "pip3 install pyrit"
}

function eaphammer_soft_install() {
	goodecho "[+] Installing eaphammer"
	[ -d /root/thirdparty ] || mkdir /root/thirdparty
	cd /root/thirdparty
	installfromnet "git clone https://github.com/s0lst1c3/eaphammer.git"
	cd eaphammer/
	./ubuntu-unattended-setup
	./eaphammer --cert-wizard
}

function airgeddon_soft_install() { # TODO: install all dependencies
	goodecho "[+] Installing airgeddon"
	[ -d /rftools ] || mkdir /rftools
	cd /rftools
	installfromnet "git clone https://github.com/v1s1t0r1sh3r3/airgeddon.git"
	cd airgeddon/
}


function wifite2_soft_install () {
	goodecho "[+] Installing wifite2"
	[ -d /rftools ] || mkdir /rftools
	cd /rftools
	installfromnet "git clone https://github.com/derv82/wifite2.git"
	cd wifite2/
}