package main

import (
	"fmt"
	"github.com/progrium/macdriver/corebluetooth"
	"github.com/progrium/macdriver/objc"
)

func main() {
	centralManagerDelegate := objc.NewClass("CBCentralManagerDelegate", "NSObject")
	centralManagerDelegate.AddMethod("centralManager:didDiscoverPeripheral:advertisementData:RSSI:", func(
		central objc.Object,
		peripheral objc.Object,
		advertisementData objc.Object,
		rssi objc.Object,
	) {
		fmt.Println(central)
	})
	objc.RegisterClass(centralManagerDelegate)

	centralManagerDelegateInstance := objc.Get("CBCentralManagerDelegate").Alloc().Init()

	centralManager := corebluetooth.CBCentralManager{}
	centralManager.SetDelegate_(centralManagerDelegateInstance)

	centralManager.ScanForPeripheralsWithServices_options_(nil, nil)
}
