package main

import (
	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	servStat := false
	systray.SetIcon(iconRed)
	systray.SetTitle("Lamp server manager")
	systray.SetTooltip("Lamp manager tooltip")
	systray.AddSeparator()
	startserv := systray.AddMenuItem("Start Server", "Start LAMP server")
	systray.AddSeparator()
	stopserv := systray.AddMenuItem("Stoping server", "Stoping LAMP server")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quits this app")

	go func() {
		for {
			select {

			case <-startserv.ClickedCh:
				cmdexec("start")
				if servStat == false {
					servStat = true
					systray.SetIcon(iconGreen)
				}
			case <-stopserv.ClickedCh:
				cmdexec("stop")
				if servStat == true {
					servStat = false
					systray.SetIcon(iconRed)
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// On exit force stop the apache2 and mysql server
	cmdexec("stop")
}
