
# Exposing USB ports to WSL2 for use in linux (on windows)

On windows side:

 ``` 
 winget install usbipd
 ```

Linux side:
Following this guide: https://gitlab.com/alelec/wsl-usb-gui

1. Run the following command: 
Check wsl2 version: It should be at least 5.10.60.1 otherwise update wsl2.
 ``` 
 uname -a 
 ```

 Install some things...
 ``` 
 sudo apt install linux-tools-generic hwdata
 ```
 ``` 
 sudo update-alternatives --install /usr/local/bin/usbip usbip /usr/lib/linux-tools/*/usbip 20
 ```

## Exposing the port (windows side)

List which ports you have
 ``` 
 usbipd list
 ```

Attaches given port to wsl2
 ``` 
 usbipd wsl attach --busid={BUS-ID}
 ```

## Verifying port was exposed (linux side)
To check what ports exist:

 ``` 
 lsusb
 ```
 You might have to run the following so it has correct privileges.

 ``` 
 sudo chmod 666 /dev/ttyACM0
 ```


## To run demo.py

 ``` 
 pip install -r requirements.txt
 python3 demo.py
 ```

# Creating a Ubuntu Server 

Download Ubuntu Server via their website, or use the raspberry imager where you can select the correct ubuntu server LTS 22.04 image.

Install on a SD card, put it in your Raspberry Pi 4.

## Step 1: Setup wifi

Get wlan name for raspberry:

 ``` 
 ls -l /sys/class/net
 ```

In my case it was named wlan0. Create backup of the network file.
 ``` 
 sudo cp /etc/netplan/00-installer-config.yaml /etc/netplan/00-install-config.original.yaml
 ```

Edit the file
 ``` 
 sudo nano /etc/netplan/00-installer-config.yaml
 ```

Should look something like this:

```
network:
 version: 2 
 wifis:
   wlan0:
     access-points:
       COVID_19_TEST_MAST:
         password: some_password
     dhcp4: true
```

## Config the keyboard (optional)
```
sudo apt update
sudo apt install console-data
sudo dkpg-reconfigure keyboard-configuration 
```
Then pick whatever keyboard/language you want.

## Random commands:
SSH into server 

```
ssh ubuntu@192.168.1.86
```

## Install MQTT


Install MQTT broker on server:
```
sudo apt install -y mosquitto
```

Create a mosquitto.conf config file in /etc/mosquitto/conf.d folder.

Add the following.
```
allow_anonymous true
listener 1883 0.0.0.0
```

Other commands
```
sudo systemctl start mosquitto
sudo systemctl stop mosquitto
sudo systemctl restart mosquitto
```

