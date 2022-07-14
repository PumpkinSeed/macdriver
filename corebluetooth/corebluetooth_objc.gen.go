package corebluetooth

import (
	core "github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"unsafe"
)

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -lobjc
#define __OBJC2__ 1
#include <objc/message.h>
#include <stdlib.h>


bool corebluetooth_convertObjCBool(BOOL b) {
	if (b) { return true; }
	return false;
}


void* CBATTRequest_type_alloc() {
	return [CBATTRequest
		alloc];
}
void* CBAttribute_type_alloc() {
	return [CBAttribute
		alloc];
}
void* CBCentralManager_type_alloc() {
	return [CBCentralManager
		alloc];
}
void* CBCentral_type_alloc() {
	return [CBCentral
		alloc];
}
void* CBCharacteristic_type_alloc() {
	return [CBCharacteristic
		alloc];
}
void* CBDescriptor_type_alloc() {
	return [CBDescriptor
		alloc];
}
void* CBManager_type_alloc() {
	return [CBManager
		alloc];
}
void* CBMutableCharacteristic_type_alloc() {
	return [CBMutableCharacteristic
		alloc];
}
void* CBMutableDescriptor_type_alloc() {
	return [CBMutableDescriptor
		alloc];
}
void* CBMutableService_type_alloc() {
	return [CBMutableService
		alloc];
}
void* CBPeer_type_alloc() {
	return [CBPeer
		alloc];
}
void* CBPeripheralManager_type_alloc() {
	return [CBPeripheralManager
		alloc];
}
void* CBPeripheral_type_alloc() {
	return [CBPeripheral
		alloc];
}
void* CBService_type_alloc() {
	return [CBService
		alloc];
}
void* CBUUID_type_alloc() {
	return [CBUUID
		alloc];
}
void* CBUUID_type_UUIDWithString_(void* theString) {
	return [CBUUID
		UUIDWithString: theString];
}
void* CBUUID_type_UUIDWithData_(void* theData) {
	return [CBUUID
		UUIDWithData: theData];
}


void* CBATTRequest_inst_init(void *id) {
	return [(CBATTRequest*)id
		init];
}

void* CBATTRequest_inst_central(void *id) {
	return [(CBATTRequest*)id
		central];
}

void* CBATTRequest_inst_characteristic(void *id) {
	return [(CBATTRequest*)id
		characteristic];
}

unsigned long CBATTRequest_inst_offset(void *id) {
	return [(CBATTRequest*)id
		offset];
}

void* CBAttribute_inst_init(void *id) {
	return [(CBAttribute*)id
		init];
}

void* CBAttribute_inst_UUID(void *id) {
	return [(CBAttribute*)id
		UUID];
}

void* CBCentralManager_inst_init(void *id) {
	return [(CBCentralManager*)id
		init];
}

void CBCentralManager_inst_connectPeripheral_options_(void *id, void* peripheral, void* options) {
	[(CBCentralManager*)id
		connectPeripheral: peripheral
		options: options];
}

void CBCentralManager_inst_cancelPeripheralConnection_(void *id, void* peripheral) {
	[(CBCentralManager*)id
		cancelPeripheralConnection: peripheral];
}

void* CBCentralManager_inst_retrieveConnectedPeripheralsWithServices_(void *id, void* serviceUUIDs) {
	return [(CBCentralManager*)id
		retrieveConnectedPeripheralsWithServices: serviceUUIDs];
}

void* CBCentralManager_inst_retrievePeripheralsWithIdentifiers_(void *id, void* identifiers) {
	return [(CBCentralManager*)id
		retrievePeripheralsWithIdentifiers: identifiers];
}

void CBCentralManager_inst_scanForPeripheralsWithServices_options_(void *id, void* serviceUUIDs, void* options) {
	[(CBCentralManager*)id
		scanForPeripheralsWithServices: serviceUUIDs
		options: options];
}

void CBCentralManager_inst_stopScan(void *id) {
	[(CBCentralManager*)id
		stopScan];
}

void CBCentralManager_inst_registerForConnectionEventsWithOptions_(void *id, void* options) {
	[(CBCentralManager*)id
		registerForConnectionEventsWithOptions: options];
}

BOOL CBCentralManager_inst_isScanning(void *id) {
	return [(CBCentralManager*)id
		isScanning];
}

void* CBCentralManager_inst_delegate(void *id) {
	return [(CBCentralManager*)id
		delegate];
}

void CBCentralManager_inst_setDelegate_(void *id, void* value) {
	[(CBCentralManager*)id
		setDelegate: value];
}

void* CBCentral_inst_init(void *id) {
	return [(CBCentral*)id
		init];
}

unsigned long CBCentral_inst_maximumUpdateValueLength(void *id) {
	return [(CBCentral*)id
		maximumUpdateValueLength];
}

void* CBCharacteristic_inst_init(void *id) {
	return [(CBCharacteristic*)id
		init];
}

void* CBCharacteristic_inst_service(void *id) {
	return [(CBCharacteristic*)id
		service];
}

void* CBCharacteristic_inst_value(void *id) {
	return [(CBCharacteristic*)id
		value];
}

void* CBCharacteristic_inst_descriptors(void *id) {
	return [(CBCharacteristic*)id
		descriptors];
}

BOOL CBCharacteristic_inst_isNotifying(void *id) {
	return [(CBCharacteristic*)id
		isNotifying];
}

BOOL CBCharacteristic_inst_isBroadcasted(void *id) {
	return [(CBCharacteristic*)id
		isBroadcasted];
}

void* CBDescriptor_inst_init(void *id) {
	return [(CBDescriptor*)id
		init];
}

void* CBDescriptor_inst_characteristic(void *id) {
	return [(CBDescriptor*)id
		characteristic];
}

void* CBDescriptor_inst_value(void *id) {
	return [(CBDescriptor*)id
		value];
}

void* CBManager_inst_init(void *id) {
	return [(CBManager*)id
		init];
}

void* CBMutableCharacteristic_inst_init(void *id) {
	return [(CBMutableCharacteristic*)id
		init];
}

void* CBMutableCharacteristic_inst_subscribedCentrals(void *id) {
	return [(CBMutableCharacteristic*)id
		subscribedCentrals];
}

void* CBMutableDescriptor_inst_initWithType_value_(void *id, void* UUID, void* value) {
	return [(CBMutableDescriptor*)id
		initWithType: UUID
		value: value];
}

void* CBMutableDescriptor_inst_init(void *id) {
	return [(CBMutableDescriptor*)id
		init];
}

void* CBMutableService_inst_initWithType_primary_(void *id, void* UUID, BOOL isPrimary) {
	return [(CBMutableService*)id
		initWithType: UUID
		primary: isPrimary];
}

void* CBMutableService_inst_init(void *id) {
	return [(CBMutableService*)id
		init];
}

void* CBPeer_inst_init(void *id) {
	return [(CBPeer*)id
		init];
}

void* CBPeripheralManager_inst_init(void *id) {
	return [(CBPeripheralManager*)id
		init];
}

void CBPeripheralManager_inst_addService_(void *id, void* service) {
	[(CBPeripheralManager*)id
		addService: service];
}

void CBPeripheralManager_inst_removeService_(void *id, void* service) {
	[(CBPeripheralManager*)id
		removeService: service];
}

void CBPeripheralManager_inst_removeAllServices(void *id) {
	[(CBPeripheralManager*)id
		removeAllServices];
}

void CBPeripheralManager_inst_startAdvertising_(void *id, void* advertisementData) {
	[(CBPeripheralManager*)id
		startAdvertising: advertisementData];
}

void CBPeripheralManager_inst_stopAdvertising(void *id) {
	[(CBPeripheralManager*)id
		stopAdvertising];
}

BOOL CBPeripheralManager_inst_updateValue_forCharacteristic_onSubscribedCentrals_(void *id, void* value, void* characteristic, void* centrals) {
	return [(CBPeripheralManager*)id
		updateValue: value
		forCharacteristic: characteristic
		onSubscribedCentrals: centrals];
}

void CBPeripheralManager_inst_publishL2CAPChannelWithEncryption_(void *id, BOOL encryptionRequired) {
	[(CBPeripheralManager*)id
		publishL2CAPChannelWithEncryption: encryptionRequired];
}

void* CBPeripheralManager_inst_delegate(void *id) {
	return [(CBPeripheralManager*)id
		delegate];
}

void CBPeripheralManager_inst_setDelegate_(void *id, void* value) {
	[(CBPeripheralManager*)id
		setDelegate: value];
}

BOOL CBPeripheralManager_inst_isAdvertising(void *id) {
	return [(CBPeripheralManager*)id
		isAdvertising];
}

void CBPeripheral_inst_discoverServices_(void *id, void* serviceUUIDs) {
	[(CBPeripheral*)id
		discoverServices: serviceUUIDs];
}

void CBPeripheral_inst_discoverIncludedServices_forService_(void *id, void* includedServiceUUIDs, void* service) {
	[(CBPeripheral*)id
		discoverIncludedServices: includedServiceUUIDs
		forService: service];
}

void CBPeripheral_inst_discoverCharacteristics_forService_(void *id, void* characteristicUUIDs, void* service) {
	[(CBPeripheral*)id
		discoverCharacteristics: characteristicUUIDs
		forService: service];
}

void CBPeripheral_inst_discoverDescriptorsForCharacteristic_(void *id, void* characteristic) {
	[(CBPeripheral*)id
		discoverDescriptorsForCharacteristic: characteristic];
}

void CBPeripheral_inst_readValueForCharacteristic_(void *id, void* characteristic) {
	[(CBPeripheral*)id
		readValueForCharacteristic: characteristic];
}

void CBPeripheral_inst_readValueForDescriptor_(void *id, void* descriptor) {
	[(CBPeripheral*)id
		readValueForDescriptor: descriptor];
}

void CBPeripheral_inst_writeValue_forDescriptor_(void *id, void* data, void* descriptor) {
	[(CBPeripheral*)id
		writeValue: data
		forDescriptor: descriptor];
}

void CBPeripheral_inst_setNotifyValue_forCharacteristic_(void *id, BOOL enabled, void* characteristic) {
	[(CBPeripheral*)id
		setNotifyValue: enabled
		forCharacteristic: characteristic];
}

void CBPeripheral_inst_readRSSI(void *id) {
	[(CBPeripheral*)id
		readRSSI];
}

void* CBPeripheral_inst_init(void *id) {
	return [(CBPeripheral*)id
		init];
}

void* CBPeripheral_inst_name(void *id) {
	return [(CBPeripheral*)id
		name];
}

void* CBPeripheral_inst_delegate(void *id) {
	return [(CBPeripheral*)id
		delegate];
}

void CBPeripheral_inst_setDelegate_(void *id, void* value) {
	[(CBPeripheral*)id
		setDelegate: value];
}

void* CBPeripheral_inst_services(void *id) {
	return [(CBPeripheral*)id
		services];
}

BOOL CBPeripheral_inst_canSendWriteWithoutResponse(void *id) {
	return [(CBPeripheral*)id
		canSendWriteWithoutResponse];
}

void* CBPeripheral_inst_RSSI(void *id) {
	return [(CBPeripheral*)id
		RSSI];
}

BOOL CBPeripheral_inst_ancsAuthorized(void *id) {
	return [(CBPeripheral*)id
		ancsAuthorized];
}

void* CBService_inst_init(void *id) {
	return [(CBService*)id
		init];
}

void* CBService_inst_peripheral(void *id) {
	return [(CBService*)id
		peripheral];
}

BOOL CBService_inst_isPrimary(void *id) {
	return [(CBService*)id
		isPrimary];
}

void* CBService_inst_characteristics(void *id) {
	return [(CBService*)id
		characteristics];
}

void* CBService_inst_includedServices(void *id) {
	return [(CBService*)id
		includedServices];
}

void* CBUUID_inst_init(void *id) {
	return [(CBUUID*)id
		init];
}

void* CBUUID_inst_data(void *id) {
	return [(CBUUID*)id
		data];
}

void* CBUUID_inst_UUIDString(void *id) {
	return [(CBUUID*)id
		UUIDString];
}

*/
import "C"

func convertObjCBoolToGo(b C.BOOL) bool {
	// NOTE: the prefix here is used to namespace these since the linker will
	// otherwise report a "duplicate symbol" because the C functions have the
	// same name.
	return bool(C.corebluetooth_convertObjCBool(b))
}

func convertToObjCBool(b bool) C.BOOL {
	if b {
		return 1
	}
	return 0
}

func CBATTRequest_alloc() (
	r0 CBATTRequest,
) {
	ret := C.CBATTRequest_type_alloc()
	r0 = CBATTRequest_fromPointer(ret)
	return
}

func CBAttribute_alloc() (
	r0 CBAttribute,
) {
	ret := C.CBAttribute_type_alloc()
	r0 = CBAttribute_fromPointer(ret)
	return
}

func CBCentralManager_alloc() (
	r0 CBCentralManager,
) {
	ret := C.CBCentralManager_type_alloc()
	r0 = CBCentralManager_fromPointer(ret)
	return
}

func CBCentral_alloc() (
	r0 CBCentral,
) {
	ret := C.CBCentral_type_alloc()
	r0 = CBCentral_fromPointer(ret)
	return
}

func CBCharacteristic_alloc() (
	r0 CBCharacteristic,
) {
	ret := C.CBCharacteristic_type_alloc()
	r0 = CBCharacteristic_fromPointer(ret)
	return
}

func CBDescriptor_alloc() (
	r0 CBDescriptor,
) {
	ret := C.CBDescriptor_type_alloc()
	r0 = CBDescriptor_fromPointer(ret)
	return
}

func CBManager_alloc() (
	r0 CBManager,
) {
	ret := C.CBManager_type_alloc()
	r0 = CBManager_fromPointer(ret)
	return
}

func CBMutableCharacteristic_alloc() (
	r0 CBMutableCharacteristic,
) {
	ret := C.CBMutableCharacteristic_type_alloc()
	r0 = CBMutableCharacteristic_fromPointer(ret)
	return
}

func CBMutableDescriptor_alloc() (
	r0 CBMutableDescriptor,
) {
	ret := C.CBMutableDescriptor_type_alloc()
	r0 = CBMutableDescriptor_fromPointer(ret)
	return
}

func CBMutableService_alloc() (
	r0 CBMutableService,
) {
	ret := C.CBMutableService_type_alloc()
	r0 = CBMutableService_fromPointer(ret)
	return
}

func CBPeer_alloc() (
	r0 CBPeer,
) {
	ret := C.CBPeer_type_alloc()
	r0 = CBPeer_fromPointer(ret)
	return
}

func CBPeripheralManager_alloc() (
	r0 CBPeripheralManager,
) {
	ret := C.CBPeripheralManager_type_alloc()
	r0 = CBPeripheralManager_fromPointer(ret)
	return
}

func CBPeripheral_alloc() (
	r0 CBPeripheral,
) {
	ret := C.CBPeripheral_type_alloc()
	r0 = CBPeripheral_fromPointer(ret)
	return
}

func CBService_alloc() (
	r0 CBService,
) {
	ret := C.CBService_type_alloc()
	r0 = CBService_fromPointer(ret)
	return
}

func CBUUID_alloc() (
	r0 CBUUID,
) {
	ret := C.CBUUID_type_alloc()
	r0 = CBUUID_fromPointer(ret)
	return
}

func CBUUID_UUIDWithString_(
	theString core.NSStringRef,
) (
	r0 CBUUID,
) {
	ret := C.CBUUID_type_UUIDWithString_(
		objc.RefPointer(theString),
	)
	r0 = CBUUID_fromPointer(ret)
	return
}

func CBUUID_UUIDWithData_(
	theData core.NSDataRef,
) (
	r0 CBUUID,
) {
	ret := C.CBUUID_type_UUIDWithData_(
		objc.RefPointer(theData),
	)
	r0 = CBUUID_fromPointer(ret)
	return
}

type CBATTRequestRef interface {
	Pointer() uintptr
	Init_asCBATTRequest() CBATTRequest
}

type gen_CBATTRequest struct {
	objc.Object
}

func CBATTRequest_fromPointer(ptr unsafe.Pointer) CBATTRequest {
	return CBATTRequest{gen_CBATTRequest{
		objc.Object_fromPointer(ptr),
	}}
}

func CBATTRequest_fromRef(ref objc.Ref) CBATTRequest {
	return CBATTRequest_fromPointer(unsafe.Pointer(ref.Pointer()))
}

func (x gen_CBATTRequest) Init_asCBATTRequest() (
	r0 CBATTRequest,
) {
	ret := C.CBATTRequest_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBATTRequest_fromPointer(ret)
	return
}

func (x gen_CBATTRequest) Central() (
	r0 CBCentral,
) {
	ret := C.CBATTRequest_inst_central(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBCentral_fromPointer(ret)
	return
}

func (x gen_CBATTRequest) Characteristic() (
	r0 CBCharacteristic,
) {
	ret := C.CBATTRequest_inst_characteristic(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBCharacteristic_fromPointer(ret)
	return
}

func (x gen_CBATTRequest) Offset() (
	r0 core.NSUInteger,
) {
	ret := C.CBATTRequest_inst_offset(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSUInteger(ret)
	return
}

type CBAttributeRef interface {
	Pointer() uintptr
	Init_asCBAttribute() CBAttribute
}

type gen_CBAttribute struct {
	objc.Object
}

func CBAttribute_fromPointer(ptr unsafe.Pointer) CBAttribute {
	return CBAttribute{gen_CBAttribute{
		objc.Object_fromPointer(ptr),
	}}
}

func CBAttribute_fromRef(ref objc.Ref) CBAttribute {
	return CBAttribute_fromPointer(unsafe.Pointer(ref.Pointer()))
}

func (x gen_CBAttribute) Init_asCBAttribute() (
	r0 CBAttribute,
) {
	ret := C.CBAttribute_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBAttribute_fromPointer(ret)
	return
}

func (x gen_CBAttribute) UUID() (
	r0 CBUUID,
) {
	ret := C.CBAttribute_inst_UUID(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBUUID_fromPointer(ret)
	return
}

type CBCentralManagerRef interface {
	Pointer() uintptr
	Init_asCBCentralManager() CBCentralManager
}

type gen_CBCentralManager struct {
	CBManager
}

func CBCentralManager_fromPointer(ptr unsafe.Pointer) CBCentralManager {
	return CBCentralManager{gen_CBCentralManager{
		CBManager_fromPointer(ptr),
	}}
}

func CBCentralManager_fromRef(ref objc.Ref) CBCentralManager {
	return CBCentralManager_fromPointer(unsafe.Pointer(ref.Pointer()))
}

func (x gen_CBCentralManager) Init_asCBCentralManager() (
	r0 CBCentralManager,
) {
	ret := C.CBCentralManager_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBCentralManager_fromPointer(ret)
	return
}

func (x gen_CBCentralManager) ConnectPeripheral_options_(
	peripheral CBPeripheralRef,
	options core.NSDictionaryRef,
) {
	C.CBCentralManager_inst_connectPeripheral_options_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(peripheral),
		objc.RefPointer(options),
	)
	return
}

func (x gen_CBCentralManager) CancelPeripheralConnection_(
	peripheral CBPeripheralRef,
) {
	C.CBCentralManager_inst_cancelPeripheralConnection_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(peripheral),
	)
	return
}

func (x gen_CBCentralManager) RetrieveConnectedPeripheralsWithServices_(
	serviceUUIDs core.NSArrayRef,
) (
	r0 core.NSArray,
) {
	ret := C.CBCentralManager_inst_retrieveConnectedPeripheralsWithServices_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(serviceUUIDs),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_CBCentralManager) RetrievePeripheralsWithIdentifiers_(
	identifiers core.NSArrayRef,
) (
	r0 core.NSArray,
) {
	ret := C.CBCentralManager_inst_retrievePeripheralsWithIdentifiers_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(identifiers),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_CBCentralManager) ScanForPeripheralsWithServices_options_(
	serviceUUIDs core.NSArrayRef,
	options core.NSDictionaryRef,
) {
	C.CBCentralManager_inst_scanForPeripheralsWithServices_options_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(serviceUUIDs),
		objc.RefPointer(options),
	)
	return
}

func (x gen_CBCentralManager) StopScan() {
	C.CBCentralManager_inst_stopScan(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_CBCentralManager) RegisterForConnectionEventsWithOptions_(
	options core.NSDictionaryRef,
) {
	C.CBCentralManager_inst_registerForConnectionEventsWithOptions_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(options),
	)
	return
}

func (x gen_CBCentralManager) IsScanning() (
	r0 bool,
) {
	ret := C.CBCentralManager_inst_isScanning(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_CBCentralManager) Delegate() (
	r0 objc.Object,
) {
	ret := C.CBCentralManager_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_CBCentralManager) SetDelegate_(
	value objc.Ref,
) {
	C.CBCentralManager_inst_setDelegate_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

type CBCentralRef interface {
	Pointer() uintptr
	Init_asCBCentral() CBCentral
}

type gen_CBCentral struct {
	CBPeer
}

func CBCentral_fromPointer(ptr unsafe.Pointer) CBCentral {
	return CBCentral{gen_CBCentral{
		CBPeer_fromPointer(ptr),
	}}
}

func CBCentral_fromRef(ref objc.Ref) CBCentral {
	return CBCentral_fromPointer(unsafe.Pointer(ref.Pointer()))
}

func (x gen_CBCentral) Init_asCBCentral() (
	r0 CBCentral,
) {
	ret := C.CBCentral_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBCentral_fromPointer(ret)
	return
}

func (x gen_CBCentral) MaximumUpdateValueLength() (
	r0 core.NSUInteger,
) {
	ret := C.CBCentral_inst_maximumUpdateValueLength(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSUInteger(ret)
	return
}

type CBCharacteristicRef interface {
	Pointer() uintptr
	Init_asCBCharacteristic() CBCharacteristic
}

type gen_CBCharacteristic struct {
	CBAttribute
}

func CBCharacteristic_fromPointer(ptr unsafe.Pointer) CBCharacteristic {
	return CBCharacteristic{gen_CBCharacteristic{
		CBAttribute_fromPointer(ptr),
	}}
}

func CBCharacteristic_fromRef(ref objc.Ref) CBCharacteristic {
	return CBCharacteristic_fromPointer(unsafe.Pointer(ref.Pointer()))
}

func (x gen_CBCharacteristic) Init_asCBCharacteristic() (
	r0 CBCharacteristic,
) {
	ret := C.CBCharacteristic_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBCharacteristic_fromPointer(ret)
	return
}

func (x gen_CBCharacteristic) Service() (
	r0 CBService,
) {
	ret := C.CBCharacteristic_inst_service(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBService_fromPointer(ret)
	return
}

func (x gen_CBCharacteristic) Value() (
	r0 core.NSData,
) {
	ret := C.CBCharacteristic_inst_value(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSData_fromPointer(ret)
	return
}

func (x gen_CBCharacteristic) Descriptors() (
	r0 core.NSArray,
) {
	ret := C.CBCharacteristic_inst_descriptors(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_CBCharacteristic) IsNotifying() (
	r0 bool,
) {
	ret := C.CBCharacteristic_inst_isNotifying(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_CBCharacteristic) IsBroadcasted() (
	r0 bool,
) {
	ret := C.CBCharacteristic_inst_isBroadcasted(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

type CBDescriptorRef interface {
	Pointer() uintptr
	Init_asCBDescriptor() CBDescriptor
}

type gen_CBDescriptor struct {
	CBAttribute
}

func CBDescriptor_fromPointer(ptr unsafe.Pointer) CBDescriptor {
	return CBDescriptor{gen_CBDescriptor{
		CBAttribute_fromPointer(ptr),
	}}
}

func CBDescriptor_fromRef(ref objc.Ref) CBDescriptor {
	return CBDescriptor_fromPointer(unsafe.Pointer(ref.Pointer()))
}

func (x gen_CBDescriptor) Init_asCBDescriptor() (
	r0 CBDescriptor,
) {
	ret := C.CBDescriptor_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBDescriptor_fromPointer(ret)
	return
}

func (x gen_CBDescriptor) Characteristic() (
	r0 CBCharacteristic,
) {
	ret := C.CBDescriptor_inst_characteristic(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBCharacteristic_fromPointer(ret)
	return
}

func (x gen_CBDescriptor) Value() (
	r0 objc.Object,
) {
	ret := C.CBDescriptor_inst_value(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

type CBManagerRef interface {
	Pointer() uintptr
	Init_asCBManager() CBManager
}

type gen_CBManager struct {
	objc.Object
}

func CBManager_fromPointer(ptr unsafe.Pointer) CBManager {
	return CBManager{gen_CBManager{
		objc.Object_fromPointer(ptr),
	}}
}

func CBManager_fromRef(ref objc.Ref) CBManager {
	return CBManager_fromPointer(unsafe.Pointer(ref.Pointer()))
}

func (x gen_CBManager) Init_asCBManager() (
	r0 CBManager,
) {
	ret := C.CBManager_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBManager_fromPointer(ret)
	return
}

type CBMutableCharacteristicRef interface {
	Pointer() uintptr
	Init_asCBMutableCharacteristic() CBMutableCharacteristic
}

type gen_CBMutableCharacteristic struct {
	CBCharacteristic
}

func CBMutableCharacteristic_fromPointer(ptr unsafe.Pointer) CBMutableCharacteristic {
	return CBMutableCharacteristic{gen_CBMutableCharacteristic{
		CBCharacteristic_fromPointer(ptr),
	}}
}

func CBMutableCharacteristic_fromRef(ref objc.Ref) CBMutableCharacteristic {
	return CBMutableCharacteristic_fromPointer(unsafe.Pointer(ref.Pointer()))
}

func (x gen_CBMutableCharacteristic) Init_asCBMutableCharacteristic() (
	r0 CBMutableCharacteristic,
) {
	ret := C.CBMutableCharacteristic_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBMutableCharacteristic_fromPointer(ret)
	return
}

func (x gen_CBMutableCharacteristic) SubscribedCentrals() (
	r0 core.NSArray,
) {
	ret := C.CBMutableCharacteristic_inst_subscribedCentrals(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

type CBMutableDescriptorRef interface {
	Pointer() uintptr
	Init_asCBMutableDescriptor() CBMutableDescriptor
}

type gen_CBMutableDescriptor struct {
	CBDescriptor
}

func CBMutableDescriptor_fromPointer(ptr unsafe.Pointer) CBMutableDescriptor {
	return CBMutableDescriptor{gen_CBMutableDescriptor{
		CBDescriptor_fromPointer(ptr),
	}}
}

func CBMutableDescriptor_fromRef(ref objc.Ref) CBMutableDescriptor {
	return CBMutableDescriptor_fromPointer(unsafe.Pointer(ref.Pointer()))
}

func (x gen_CBMutableDescriptor) InitWithType_value__asCBMutableDescriptor(
	UUID CBUUIDRef,
	value objc.Ref,
) (
	r0 CBMutableDescriptor,
) {
	ret := C.CBMutableDescriptor_inst_initWithType_value_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(UUID),
		objc.RefPointer(value),
	)
	r0 = CBMutableDescriptor_fromPointer(ret)
	return
}

func (x gen_CBMutableDescriptor) Init_asCBMutableDescriptor() (
	r0 CBMutableDescriptor,
) {
	ret := C.CBMutableDescriptor_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBMutableDescriptor_fromPointer(ret)
	return
}

type CBMutableServiceRef interface {
	Pointer() uintptr
	Init_asCBMutableService() CBMutableService
}

type gen_CBMutableService struct {
	CBService
}

func CBMutableService_fromPointer(ptr unsafe.Pointer) CBMutableService {
	return CBMutableService{gen_CBMutableService{
		CBService_fromPointer(ptr),
	}}
}

func CBMutableService_fromRef(ref objc.Ref) CBMutableService {
	return CBMutableService_fromPointer(unsafe.Pointer(ref.Pointer()))
}

func (x gen_CBMutableService) InitWithType_primary__asCBMutableService(
	UUID CBUUIDRef,
	isPrimary bool,
) (
	r0 CBMutableService,
) {
	ret := C.CBMutableService_inst_initWithType_primary_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(UUID),
		convertToObjCBool(isPrimary),
	)
	r0 = CBMutableService_fromPointer(ret)
	return
}

func (x gen_CBMutableService) Init_asCBMutableService() (
	r0 CBMutableService,
) {
	ret := C.CBMutableService_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBMutableService_fromPointer(ret)
	return
}

type CBPeerRef interface {
	Pointer() uintptr
	Init_asCBPeer() CBPeer
}

type gen_CBPeer struct {
	objc.Object
}

func CBPeer_fromPointer(ptr unsafe.Pointer) CBPeer {
	return CBPeer{gen_CBPeer{
		objc.Object_fromPointer(ptr),
	}}
}

func CBPeer_fromRef(ref objc.Ref) CBPeer {
	return CBPeer_fromPointer(unsafe.Pointer(ref.Pointer()))
}

func (x gen_CBPeer) Init_asCBPeer() (
	r0 CBPeer,
) {
	ret := C.CBPeer_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBPeer_fromPointer(ret)
	return
}

type CBPeripheralManagerRef interface {
	Pointer() uintptr
	Init_asCBPeripheralManager() CBPeripheralManager
}

type gen_CBPeripheralManager struct {
	CBManager
}

func CBPeripheralManager_fromPointer(ptr unsafe.Pointer) CBPeripheralManager {
	return CBPeripheralManager{gen_CBPeripheralManager{
		CBManager_fromPointer(ptr),
	}}
}

func CBPeripheralManager_fromRef(ref objc.Ref) CBPeripheralManager {
	return CBPeripheralManager_fromPointer(unsafe.Pointer(ref.Pointer()))
}

func (x gen_CBPeripheralManager) Init_asCBPeripheralManager() (
	r0 CBPeripheralManager,
) {
	ret := C.CBPeripheralManager_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBPeripheralManager_fromPointer(ret)
	return
}

func (x gen_CBPeripheralManager) AddService_(
	service CBMutableServiceRef,
) {
	C.CBPeripheralManager_inst_addService_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(service),
	)
	return
}

func (x gen_CBPeripheralManager) RemoveService_(
	service CBMutableServiceRef,
) {
	C.CBPeripheralManager_inst_removeService_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(service),
	)
	return
}

func (x gen_CBPeripheralManager) RemoveAllServices() {
	C.CBPeripheralManager_inst_removeAllServices(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_CBPeripheralManager) StartAdvertising_(
	advertisementData core.NSDictionaryRef,
) {
	C.CBPeripheralManager_inst_startAdvertising_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(advertisementData),
	)
	return
}

func (x gen_CBPeripheralManager) StopAdvertising() {
	C.CBPeripheralManager_inst_stopAdvertising(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_CBPeripheralManager) UpdateValue_forCharacteristic_onSubscribedCentrals_(
	value core.NSDataRef,
	characteristic CBMutableCharacteristicRef,
	centrals core.NSArrayRef,
) (
	r0 bool,
) {
	ret := C.CBPeripheralManager_inst_updateValue_forCharacteristic_onSubscribedCentrals_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
		objc.RefPointer(characteristic),
		objc.RefPointer(centrals),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_CBPeripheralManager) PublishL2CAPChannelWithEncryption_(
	encryptionRequired bool,
) {
	C.CBPeripheralManager_inst_publishL2CAPChannelWithEncryption_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(encryptionRequired),
	)
	return
}

func (x gen_CBPeripheralManager) Delegate() (
	r0 objc.Object,
) {
	ret := C.CBPeripheralManager_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_CBPeripheralManager) SetDelegate_(
	value objc.Ref,
) {
	C.CBPeripheralManager_inst_setDelegate_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_CBPeripheralManager) IsAdvertising() (
	r0 bool,
) {
	ret := C.CBPeripheralManager_inst_isAdvertising(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

type CBPeripheralRef interface {
	Pointer() uintptr
	Init_asCBPeripheral() CBPeripheral
}

type gen_CBPeripheral struct {
	CBPeer
}

func CBPeripheral_fromPointer(ptr unsafe.Pointer) CBPeripheral {
	return CBPeripheral{gen_CBPeripheral{
		CBPeer_fromPointer(ptr),
	}}
}

func CBPeripheral_fromRef(ref objc.Ref) CBPeripheral {
	return CBPeripheral_fromPointer(unsafe.Pointer(ref.Pointer()))
}

func (x gen_CBPeripheral) DiscoverServices_(
	serviceUUIDs core.NSArrayRef,
) {
	C.CBPeripheral_inst_discoverServices_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(serviceUUIDs),
	)
	return
}

func (x gen_CBPeripheral) DiscoverIncludedServices_forService_(
	includedServiceUUIDs core.NSArrayRef,
	service CBServiceRef,
) {
	C.CBPeripheral_inst_discoverIncludedServices_forService_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(includedServiceUUIDs),
		objc.RefPointer(service),
	)
	return
}

func (x gen_CBPeripheral) DiscoverCharacteristics_forService_(
	characteristicUUIDs core.NSArrayRef,
	service CBServiceRef,
) {
	C.CBPeripheral_inst_discoverCharacteristics_forService_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(characteristicUUIDs),
		objc.RefPointer(service),
	)
	return
}

func (x gen_CBPeripheral) DiscoverDescriptorsForCharacteristic_(
	characteristic CBCharacteristicRef,
) {
	C.CBPeripheral_inst_discoverDescriptorsForCharacteristic_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(characteristic),
	)
	return
}

func (x gen_CBPeripheral) ReadValueForCharacteristic_(
	characteristic CBCharacteristicRef,
) {
	C.CBPeripheral_inst_readValueForCharacteristic_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(characteristic),
	)
	return
}

func (x gen_CBPeripheral) ReadValueForDescriptor_(
	descriptor CBDescriptorRef,
) {
	C.CBPeripheral_inst_readValueForDescriptor_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(descriptor),
	)
	return
}

func (x gen_CBPeripheral) WriteValue_forDescriptor_(
	data core.NSDataRef,
	descriptor CBDescriptorRef,
) {
	C.CBPeripheral_inst_writeValue_forDescriptor_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(descriptor),
	)
	return
}

func (x gen_CBPeripheral) SetNotifyValue_forCharacteristic_(
	enabled bool,
	characteristic CBCharacteristicRef,
) {
	C.CBPeripheral_inst_setNotifyValue_forCharacteristic_(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(enabled),
		objc.RefPointer(characteristic),
	)
	return
}

func (x gen_CBPeripheral) ReadRSSI() {
	C.CBPeripheral_inst_readRSSI(
		unsafe.Pointer(x.Pointer()),
	)
	return
}

func (x gen_CBPeripheral) Init_asCBPeripheral() (
	r0 CBPeripheral,
) {
	ret := C.CBPeripheral_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBPeripheral_fromPointer(ret)
	return
}

func (x gen_CBPeripheral) Name() (
	r0 core.NSString,
) {
	ret := C.CBPeripheral_inst_name(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}

func (x gen_CBPeripheral) Delegate() (
	r0 objc.Object,
) {
	ret := C.CBPeripheral_inst_delegate(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = objc.Object_fromPointer(ret)
	return
}

func (x gen_CBPeripheral) SetDelegate_(
	value objc.Ref,
) {
	C.CBPeripheral_inst_setDelegate_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)
	return
}

func (x gen_CBPeripheral) Services() (
	r0 core.NSArray,
) {
	ret := C.CBPeripheral_inst_services(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_CBPeripheral) CanSendWriteWithoutResponse() (
	r0 bool,
) {
	ret := C.CBPeripheral_inst_canSendWriteWithoutResponse(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_CBPeripheral) RSSI() (
	r0 core.NSNumber,
) {
	ret := C.CBPeripheral_inst_RSSI(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSNumber_fromPointer(ret)
	return
}

func (x gen_CBPeripheral) AncsAuthorized() (
	r0 bool,
) {
	ret := C.CBPeripheral_inst_ancsAuthorized(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

type CBServiceRef interface {
	Pointer() uintptr
	Init_asCBService() CBService
}

type gen_CBService struct {
	CBAttribute
}

func CBService_fromPointer(ptr unsafe.Pointer) CBService {
	return CBService{gen_CBService{
		CBAttribute_fromPointer(ptr),
	}}
}

func CBService_fromRef(ref objc.Ref) CBService {
	return CBService_fromPointer(unsafe.Pointer(ref.Pointer()))
}

func (x gen_CBService) Init_asCBService() (
	r0 CBService,
) {
	ret := C.CBService_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBService_fromPointer(ret)
	return
}

func (x gen_CBService) Peripheral() (
	r0 CBPeripheral,
) {
	ret := C.CBService_inst_peripheral(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBPeripheral_fromPointer(ret)
	return
}

func (x gen_CBService) IsPrimary() (
	r0 bool,
) {
	ret := C.CBService_inst_isPrimary(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = convertObjCBoolToGo(ret)
	return
}

func (x gen_CBService) Characteristics() (
	r0 core.NSArray,
) {
	ret := C.CBService_inst_characteristics(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

func (x gen_CBService) IncludedServices() (
	r0 core.NSArray,
) {
	ret := C.CBService_inst_includedServices(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSArray_fromPointer(ret)
	return
}

type CBUUIDRef interface {
	Pointer() uintptr
	Init_asCBUUID() CBUUID
}

type gen_CBUUID struct {
	objc.Object
}

func CBUUID_fromPointer(ptr unsafe.Pointer) CBUUID {
	return CBUUID{gen_CBUUID{
		objc.Object_fromPointer(ptr),
	}}
}

func CBUUID_fromRef(ref objc.Ref) CBUUID {
	return CBUUID_fromPointer(unsafe.Pointer(ref.Pointer()))
}

func (x gen_CBUUID) Init_asCBUUID() (
	r0 CBUUID,
) {
	ret := C.CBUUID_inst_init(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = CBUUID_fromPointer(ret)
	return
}

func (x gen_CBUUID) Data() (
	r0 core.NSData,
) {
	ret := C.CBUUID_inst_data(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSData_fromPointer(ret)
	return
}

func (x gen_CBUUID) UUIDString() (
	r0 core.NSString,
) {
	ret := C.CBUUID_inst_UUIDString(
		unsafe.Pointer(x.Pointer()),
	)
	r0 = core.NSString_fromPointer(ret)
	return
}
