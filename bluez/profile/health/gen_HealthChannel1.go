// Code generated DO NOT EDIT

package health



import (
   "sync"
   "github.com/phedoreanu/go-bluetooth/bluez"
   "github.com/phedoreanu/go-bluetooth/util"
   "github.com/phedoreanu/go-bluetooth/props"
   "github.com/godbus/dbus"
)

var HealthChannel1Interface = "org.bluez.HealthChannel1"


// NewHealthChannel1 create a new instance of HealthChannel1
//
// Args:
// - objectPath: [variable prefix]/{hci0,hci1,...}/dev_XX_XX_XX_XX_XX_XX/chanZZZ
func NewHealthChannel1(objectPath dbus.ObjectPath) (*HealthChannel1, error) {
	a := new(HealthChannel1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez",
			Iface: HealthChannel1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(HealthChannel1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}


/*
HealthChannel1 HealthChannel hierarchy

*/
type HealthChannel1 struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	objectManagerSignal chan *dbus.Signal
	objectManager       *bluez.ObjectManager
	Properties 				*HealthChannel1Properties
	watchPropertiesChannel chan *dbus.Signal
}

// HealthChannel1Properties contains the exposed properties of an interface
type HealthChannel1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

	/*
	Type The quality of service of the data channel. ("reliable"
			or "streaming")
	*/
	Type string

	/*
	Device Identifies the Remote Device that is connected with.
			Maps with a HealthDevice object.
	*/
	Device dbus.ObjectPath

	/*
	Application Identifies the HealthApplication to which this channel
			is related to (which indirectly defines its role and
			data type).
	*/
	Application dbus.ObjectPath

}

//Lock access to properties
func (p *HealthChannel1Properties) Lock() {
	p.lock.Lock()
}

//Unlock access to properties
func (p *HealthChannel1Properties) Unlock() {
	p.lock.Unlock()
}






// GetType get Type value
func (a *HealthChannel1) GetType() (string, error) {
	v, err := a.GetProperty("Type")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}






// GetDevice get Device value
func (a *HealthChannel1) GetDevice() (dbus.ObjectPath, error) {
	v, err := a.GetProperty("Device")
	if err != nil {
		return dbus.ObjectPath(""), err
	}
	return v.Value().(dbus.ObjectPath), nil
}






// GetApplication get Application value
func (a *HealthChannel1) GetApplication() (dbus.ObjectPath, error) {
	v, err := a.GetProperty("Application")
	if err != nil {
		return dbus.ObjectPath(""), err
	}
	return v.Value().(dbus.ObjectPath), nil
}



// Close the connection
func (a *HealthChannel1) Close() {
	
	a.unregisterPropertiesSignal()
	
	a.client.Disconnect()
}

// Path return HealthChannel1 object path
func (a *HealthChannel1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return HealthChannel1 dbus client
func (a *HealthChannel1) Client() *bluez.Client {
	return a.client
}

// Interface return HealthChannel1 interface
func (a *HealthChannel1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *HealthChannel1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

	if a.objectManagerSignal == nil {
		if a.objectManager == nil {
			om, err := bluez.GetObjectManager()
			if err != nil {
				return nil, nil, err
			}
			a.objectManager = om
		}

		s, err := a.objectManager.Register()
		if err != nil {
			return nil, nil, err
		}
		a.objectManagerSignal = s
	}

	cancel := func() {
		if a.objectManagerSignal == nil {
			return
		}
		a.objectManagerSignal <- nil
		a.objectManager.Unregister(a.objectManagerSignal)
		a.objectManagerSignal = nil
	}

	return a.objectManagerSignal, cancel, nil
}


// ToMap convert a HealthChannel1Properties to map
func (a *HealthChannel1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an HealthChannel1Properties
func (a *HealthChannel1Properties) FromMap(props map[string]interface{}) (*HealthChannel1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an HealthChannel1Properties
func (a *HealthChannel1Properties) FromDBusMap(props map[string]dbus.Variant) (*HealthChannel1Properties, error) {
	s := new(HealthChannel1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *HealthChannel1) ToProps() bluez.Properties {
	return a.Properties
}

// GetWatchPropertiesChannel return the dbus channel to receive properties interface
func (a *HealthChannel1) GetWatchPropertiesChannel() chan *dbus.Signal {
	return a.watchPropertiesChannel
}

// SetWatchPropertiesChannel set the dbus channel to receive properties interface
func (a *HealthChannel1) SetWatchPropertiesChannel(c chan *dbus.Signal) {
	a.watchPropertiesChannel = c
}

// GetProperties load all available properties
func (a *HealthChannel1) GetProperties() (*HealthChannel1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *HealthChannel1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *HealthChannel1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *HealthChannel1) GetPropertiesSignal() (chan *dbus.Signal, error) {

	if a.propertiesSignal == nil {
		s, err := a.client.Register(a.client.Config.Path, bluez.PropertiesInterface)
		if err != nil {
			return nil, err
		}
		a.propertiesSignal = s
	}

	return a.propertiesSignal, nil
}

// Unregister for changes signalling
func (a *HealthChannel1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *HealthChannel1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *HealthChannel1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}




/*
Acquire 
			Returns the file descriptor for this data channel. If
			the data channel is not connected it will also
			reconnect.

			Possible Errors: org.bluez.Error.NotConnected
					 org.bluez.Error.NotAllowed


*/
func (a *HealthChannel1) Acquire() (dbus.UnixFD, error) {
	
	var val0 dbus.UnixFD
	err := a.client.Call("Acquire", 0, ).Store(&val0)
	return val0, err	
}

/*
Release 

*/
func (a *HealthChannel1) Release() error {
	
	return a.client.Call("Release", 0, ).Store()
	
}

/*
close 
			Possible Errors: org.bluez.Error.NotAcquired
					 org.bluez.Error.NotAllowed


*/
func (a *HealthChannel1) close() error {
	
	return a.client.Call("close", 0, ).Store()
	
}

