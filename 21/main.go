// Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

// WifiDevice - устройство с интерфейсом предоставления Wi-Fi соединения
type WifiDevice struct{}

// provideWifiConnection - метод предоставления Wi-Fi соединения
func (device WifiDevice) provideWifiConnection() {
	fmt.Println("Предоставляем Wi-Fi соединение.")
}

// EthernetAdapter - адаптер для адаптации интерфейса устройства с Wi-Fi под Ethernet
type EthernetAdapter struct {
	Device *WifiDevice
}

// requestEthernetConnection - метод адаптированного устройства для запроса Ethernet соединения
func (adapter EthernetAdapter) requestEthernetConnection() {
	fmt.Println("Используем адаптер для запроса Ethernet соединения.")
	adapter.Device.provideWifiConnection()
}

// Client - клиентский код, ожидающий использование Ethernet соединения
type Client struct{}

// UseEthernetConnection - клиентский код использует устройство с Ethernet соединением
func (client Client) UseEthernetConnection(adapter EthernetAdapter) {
	adapter.requestEthernetConnection()
}

func main() {
	client := Client{}
	wifiDevice := WifiDevice{}
	adapter := EthernetAdapter{Device: &wifiDevice}

	client.UseEthernetConnection(adapter)
}
