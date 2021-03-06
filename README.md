<p align="center">
  <img alt="Cloudflare DDNS Agent logo" src="./readme/header.png" width="auto"></img>
</p>

<p align="center">
    <a href="https://goreportcard.com/report/github.com/zerjioang/ddns-cloudflare">
    	<img alt="Software License" src="https://goreportcard.com/badge/github.com/zerjioang/ddns-cloudflare"/>
    </a>
    <a href="https://github.com/zerjioang/ddns-cloudflare/blob/master/LICENSE">
    	<img alt="Software License" src="http://img.shields.io/:license-GPL3-brightgreen.svg?style=flat-square"/>
    </a>
    <a href="https://travis-ci.org/zerjioang/ddns-cloudflare">
      <img alt="Build Status" src="https://github.com/zerjioang/ddns-cloudflare/actions/workflows/go.yml/badge.svg">
    </a>
	<a href="https://pkg.go.dev/github.com/zerjioang/ddns-cloudflare">
    	<img alt="Package documentation" src="http://img.shields.io/:Go-docs-blue.svg?style=flat-square"/>
    </a>
</p>

## Introduction

DDNS is an agent that continuously monitors changes in the public Internet IP address of a connection for a host and updates Cloudflare's DNS when a change is detected.

An example setup illustrated below has a Raspberry Pi running a webserver with DDNS on a home Internet connection. The website's domain name (let's use example.com) resolves to Cloudflare's DNS servers. When a web client access example.com, the HTTP/S connection terminates on Cloudflare's infrastructure. Cloudflare then proxies the connection to the home Internet connection where the home router port forwards to the Raspberry Pi. At any time if the home Internet connection IP address changes, DDNS sends an update informing Cloudflare of the new home IP address for it to proxy to.


## Goal

It updates your device (raspberry pi, beaglebone, linux, container, vm, etc) IP address on a Dynamic
DNS provider like Cloudflare so that you can forget about device IP and use it a DNS name.

## Requirements

* A Cloudflare account.
* A owned domain name registered in Cloudflare.
* A Token created using Cloudflare dashboard to allow this application to edit your DNS Zone info.

## Usage

Some usage examples:

```bash
X_CF_AGENT_TOKEN=yourtoken \
X_CF_AGENT_ZONE=example.com \
X_CF_AGENT_DNS_A_RECORD=rpi \
ddns-cloudflare update
```

```bash
export X_CF_AGENT_TOKEN=yourtoken
export X_CF_AGENT_ZONE=example.com
export X_CF_AGENT_DNS_A_RECORD=rpi

ddns-cloudflare update
```

```bash
echo 'export X_CF_AGENT_TOKEN=yourtoken' >> $HOME/.bashrc
echo 'export X_CF_AGENT_ZONE=example.com' >> $HOME/.bashrc
echo 'export X_CF_AGENT_DNS_A_RECORD=rpi' >> $HOME/.bashrc

ddns-cloudflare update
```

Previous command will try to update the DNS A Record of `rpi.example.com` with the host actual public IP

### Customization of DDNS Cloudflare Agent

* **X_CF_AGENT_TOKEN**: token provided by Cloudflare in order to authenticate API calls
* **X_CF_AGENT_ZONE**: name of the DNS zone you want to edit
* **X_CF_AGENT_DNS_A_RECORD**: name of the DNS Record name you want to edit.

## Start modes

### Option 1: Using systemd autostart

Edit the ddns-cloudflare.service file and move it to `/etc/systemd/system/`, then run:

```bash
sudo systemctl daemon-reload
sudo systemctl enable ddns-cloudflare
sudo systemctl start ddns-cloudflare
```

### Option 2: Using cron

Access to your cron file with `crontab -e` and add the following content. For example, in order to update the device IP eery 10 minutes you need to add

```bash
# run a Cloudflare DDNS updater each 10 minutes
*/10 * * * * /opt/ddns-cloudflare update
```

You can verify the content of the with `crontab -l`. Your new line in the crontab file should look like

```bash
# Edit this file to introduce tasks to be run by cron.
# 
# Each task to run has to be defined through a single line
# indicating with different fields when the task will be run
# and what command to run for the task
# 
# To define the time you can provide concrete values for
# minute (m), hour (h), day of month (dom), month (mon),
# and day of week (dow) or use '*' in these fields (for 'any').
# 
# Notice that tasks will be started based on the cron's system
# daemon's notion of time and timezones.
# 
# Output of the crontab jobs (including errors) is sent through
# email to the user the crontab file belongs to (unless redirected).
# 
# For example, you can run a backup of all your user accounts
# at 5 a.m every week with:
# 0 5 * * 1 tar -zcf /var/backups/home.tgz /home/
# 
# For more information see the manual pages of crontab(5) and cron(8)
# 
# m h  dom mon dow   command

# run a Cloudflare DDNS updater each 10 minutes
*/10 * * * * /opt/ddns-cloudflare update
```

I recommend you to use https://crontab.guru/ if you want to easily configure the crontab timing

## Troubleshooting

### API Token must not be empty

```bash
2021/10/20 20:27:09 Updating device IP. Please wait...
2021/10/20 20:27:09 Requesting IP check for:  .
2021/10/20 20:27:09 Reading current device IP. Please wait...
2021/10/20 20:27:09 Readed IP:  X.Y.Z.A
2021/10/20 20:27:09 Connecting with Cloudflare services...
2021/10/20 20:27:09 cloudflare ddns updater took 101.71621ms
invalid credentials: API Token must not be empty
```

Make sure you have successfully set the `X_CF_AGENT_TOKEN` environment variable with a valid Cloudflare token.

### 

```bash
2021/10/20 20:28:56 Updating device IP. Please wait...
2021/10/20 20:28:56 Requesting IP check for:  example.com.rpi
2021/10/20 20:28:56 Reading current device IP. Please wait...
2021/10/20 20:28:57 Readed IP:  X.Y.Z.A
2021/10/20 20:28:57 Connecting with Cloudflare services...
2021/10/20 20:28:58 cloudflare ddns updater took 1.070368081s

ListZonesContext command failed: HTTP status 400: Invalid request headers (6003)
```
Make sure you have successfully set the `X_CF_AGENT_TOKEN` and the token has the ability to manage your DNS Zone data.
If you set the wrong scope of the token, you have to create a new one.

## License

<p align="center">
  <img alt="GPLv3 logo" src="https://upload.wikimedia.org/wikipedia/commons/thumb/9/93/GPLv3_Logo.svg/1280px-GPLv3_Logo.svg.png" width="100px"></img>
</p>

All rights reserved to project author(s)

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

 * Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
 * Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
 * Uses GPL license described below

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

See [**LICENSE**](LICENSE) file for full license details.