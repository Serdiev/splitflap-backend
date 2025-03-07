## Run project (check makefile)

To build the binary run
```
sudo make build
```
To run the program in the background run
```
sudo make ex
```

To run it locally on ubuntu ```sudo make run```.
On windows ```sudo make runn```



# Creating a Ubuntu Server 

Use the raspberry imager where you can select the correct ubuntu server LTS 22.04 image.

Install on a SD card, put it in your Raspberry Pi 4.

## Step 1: Setup wifi (if not setup via raspberry imager)

Get wlan name for raspberry:

 ``` 
 ls -l /sys/class/net
 ```

In my case it was named wlan0. Create backup of the network file.
 ``` 
 sudo cp /etc/netplan/50-cloud-init.yaml /etc/netplan/50-cloud-init.original.yaml
 ```

Edit the file
 ``` 
 sudo nano /etc/netplan/50-cloud-init.yaml
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

## Config the keyboard (if not setup via raspberry imager)
```
sudo apt update
sudo apt install console-data
sudo dkpg-reconfigure keyboard-configuration 
```
Then pick whatever keyboard/language you want.

## Update apt
```
sudo apt update
```

## Install make
```
sudo apt install make
```

## Create ssh key
```
  ssh-keygen -t ed25519 -C "your_email@example.com"
  eval "$(ssh-agent -s)"
  ssh-add ~/.ssh/rp3
```

## Install Golang
Make sure armv6l version for raspberry pi.
```
  wget https://go.dev/dl/go1.23.2.linux-armv6l.tar.gz
  sudo tar -C /usr/local -xzf go1.23.2.linux-armv6l.tar.gz
  sudo snap install golangci-lint --classic
```

## Random commands:
SSH into server

```
ssh ubuntu@192.168.1.86
```

## Install MQTT (personally not using it)

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

## To run demo.py (scottbez1 repo)
https://github.com/scottbez1/splitflap

 ``` 
 pip install -r requirements.txt
 python3 demo.py
 ```