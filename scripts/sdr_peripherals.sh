#!/bin/bash

source common.sh

function ad_devices_install() {
	goodecho "[+] Installing AD libs and tools from package manager"
	installfromnet "apt-fast install -y libad9361-dev libiio-utils"
}

function uhd_devices_install() {
	goodecho "[+] Installing UHD's libs and tools from package manager"
	installfromnet "apt-fast install -y libuhd4.1.0 libuhd-dev uhd-host"
	goodecho "[+] Copying rules sets"
	cp /root/rules/uhd-usrp.rules  /etc/udev/rules.d/
	goodecho "[+] Downloading Hardware Driver firmware/FPGA"
    installfromnet "/usr/bin/uhd_images_downloader"
}

function nuand_devices_install() {
	goodecho "[+] Installing Nuand's libs and tools from package manager"
	installfromnet "add-apt-repository ppa:nuandllc/bladerf"
	installfromnet "apt-fast update"
	installfromnet "apt-fast install -y bladerf libbladerf-dev bladerf-firmware-fx3"
	goodecho "[+] Copying rules sets"
	cp /root/rules/88-nuand-bladerf1.rules.in /etc/udev/rules.d/
	cp /root/rules/88-nuand-bladerf2.rules.in /etc/udev/rules.d/
	cp /root/rules/88-nuand-bootloader.rules.in /etc/udev/rules.d/
}

function hackrf_devices_install() {
	goodecho "[+] Installing hackRF's libs and tools from package manager"
	installfromnet "apt-fast install -y hackrf libhackrf-dev"
}

function airspy_devices_install() {
	goodecho "[+] Installing airspy from package manager"
	installfromnet "apt-fast install -y airspy"
	goodecho "[+] Downloading airspy HF+ drivers and libs"
	[ -d /root/thirdparty ] || mkdir /root/thirdparty
	cd /root/thirdparty
	installfromnet "git clone https://github.com/airspy/airspyhf.git"
	cd airspyhf
	mkdir build
	cd build
	goodecho "[+] Building and installing Airspy libs"
	cmake -DCMAKE_INSTALL_PREFIX=/usr ../
	make -j$(nproc); sudo make install
	cd ../..
}

function limesdr_devices_install() {
	goodecho "[+] Installing LimeSDR's libs and tools from package manager"
	installfromnet "apt-fast install -y soapysdr-module-lms7 libsoapysdr-dev"
}

function install_soapy_modules() {
	goodecho "[+] Installing Soapy extra modules"
	installfromnet "apt-fast install -y libsoapysdr-dev soapysdr-module-osmosdr soapysdr-module-rtlsdr soapysdr-module-bladerf soapysdr-module-hackrf soapysdr-module-uhd soapysdr-module-mirisdr soapysdr-module-rfspace soapysdr-module-airspy"
}

function rtlsdr_devices_install() {
	goodecho "[+] Installing RTL-SDR's libs and tools from package manager"
	installfromnet "apt-fast install -y librtlsdr-dev librtlsdr0"
}

function rtlsdrv4_devices_install() {
	goodecho "[+] Installing RTL-SDR v4's libs and tools from package manager"
	apt purge -y ^librtlsdr
	rm -rvf /usr/lib/librtlsdr* /usr/include/rtl-sdr* /usr/local/lib/librtlsdr* /usr/local/include/rtl-sdr* /usr/local/include/rtl_* /usr/local/bin/rtl_*
	installfromnet "apt-fast install -y libusb-1.0-0-dev git cmake pkg-config"
	[ -d /root/thirdparty ] || mkdir /root/thirdparty
	cd /root/thirdparty
	installfromnet "git clone https://github.com/rtlsdrblog/rtl-sdr-blog"
	cd rtl-sdr-blog
	mkdir build
	cd build
	cmake ../ -DINSTALL_UDEV_RULES=ON
	make
	sudo make install
	sudo cp ../rtl-sdr.rules /etc/udev/rules.d/
	sudo ldconfig
	cd /root
	rm -R /root/thirdparty
}