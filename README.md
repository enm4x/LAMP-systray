<p align="center">
<img width="400" height="250" src="https://raw.githubusercontent.com/maxdvlg/LAMP-systray/master/docs/icon.png">
</p>

# LAMP systray  ![Build Status](https://travis-ci.com/maxdvlg/LAMP-systray.svg?branch=master)

Simple LAMP Gui in the systray for Ubuntu

## Getting Started

These instructions will show how to install the program or how to build it from sources 

## Prerequisites

* System : Ubuntu / Popos ( working on 19.10 / 18.04 )
* LAMP : [Installation guide](https://doc.ubuntu-fr.org/lamp)
* Dependencies : 

```shell
sudo apt-get install libgtk-3-dev libappindicator3-dev libwebkit2gtk-4.0-dev
```

---

## Install Lamp systray

### download the project 

```shell
wget https://github.com/maxdvlg/LAMP-systray/archive/master.zip
```

### copy the exec folder and rename it 

```shell
cp LAMP-systray/bin /your/path/APPNAME
```

### create the icon for the launcher

we need to move the lamp.desktop file to make it visible for your gnome launcher

```shell
sudo mv /path/APPNAME/lamp.desktop /usr/share/applications/
```

---

## Build from source

### Prerequisites

* Go 1.13+
* Go modules

## Deployment

```shell
cd path/LAMP-systray/
go build -o /execPath/execName main.go icon.go cmdexec.go
```

in case of build issue you should check if go-module is on :

```shell
go env 					#display all the go env variable
export GOMODULE111=on 	#activate go module
source $HOME/.profile	#update path for your terminal emulator 
```

## Built With

* GO [1.13.x] - main code
  * [Systray](https://github.com/getlantern/systray) - Go package for systray integration 

## Authors

* **WVX** : *Initial work* - [maxdvlg](https://github.com/maxdvlg)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Roadmap

* Test cross platform availability
* disable multiple launch at the same time
* Adding output status option for the apache and mysql server

## Acknowledgments

* if you want to participate feel free to fork or to propose pull request