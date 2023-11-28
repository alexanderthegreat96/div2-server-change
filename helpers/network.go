package helpers

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/google/gopacket/pcap"
)

func isProcessRunning(output []byte, processName string) bool {
	return bytesContains(output, []byte(processName))
}

func limitBandwidth(processName string, bandwidthLimit int) {
	// (The previous code to limit bandwidth)

	cmd := exec.Command("netsh", "wlan", "set", "autoconfig", "enabled=no")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error disabling auto configuration:", err)
	}

	cmd = exec.Command("netsh", "interface", "set", "interface", "name=\"Wi-Fi\"", "admin=disable")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error disabling interface:", err)
	}

	cmd = exec.Command("netsh", "interface", "ipv4", "set", "interface", "name=\"Wi-Fi\"", "mtu=800")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error setting MTU:", err)
	}
}

func resetNetworkSettings() {
	cmd := exec.Command("netsh", "wlan", "set", "autoconfig", "enabled=yes")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error enabling auto configuration:", err)
	}

	cmd = exec.Command("netsh", "interface", "set", "interface", "name=\"Wi-Fi\"", "admin=enable")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error enabling interface:", err)
	}

	cmd = exec.Command("netsh", "interface", "ipv4", "set", "interface", "name=\"Wi-Fi\"", "mtu=1500")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error resetting MTU:", err)
	}
}

func bytesContains(haystack, needle []byte) bool {
	return bytes.Index(haystack, needle) != -1
}

func GrabAllDevices() map[string]string {

	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln("Unable to fetch network interfaces")
	}

	deviceList := make(map[string]string)

	for _, device := range devices {
		deviceList[device.Name] = device.Description
	}

	return deviceList
}
