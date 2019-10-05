#!/usr/bin/env sh

INTERFACE=$(ip route | awk '/default/ { print $5 }')
HOST_IP=$(sipcalc ${INTERFACE} | awk 'NR == 4 {print $4}')
MIN_IP=$(sipcalc ${INTERFACE} | awk '/Usable range/ {print $4}')

/usr/sbin/dnsmasq \
  --dhcp-range=${MIN_IP},proxy,255.255.255.0 \
  --enable-tftp --tftp-root=/tftp \
  --dhcp-userclass=set:ipxe,iPXE \
  --pxe-service=tag:#ipxe,x86PC,"PXE chainload to iPXE",undionly.kpxe \
  --pxe-service=tag:ipxe,x86PC,"iPXE",http://${HOST_IP}:8080/ipxe \
  --log-queries \
  --log-dhcp \
  --no-daemon \
  --port=0