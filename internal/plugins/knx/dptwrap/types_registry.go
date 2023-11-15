package dptwrap

import (
	"fmt"
	"reflect"
	"sync"
)

type IDpt interface {
	Name() string

	Pack() []byte

	Unpack(data []byte) error

	Unit() string

	String() string

	Set(i interface{})

	Get() interface{}
}

var (
	types = [...]IDpt{
		//1.x
		new(DPT_1001_Wrap),
		new(DPT_1002_Wrap),
		new(DPT_1003_Wrap),
		new(DPT_1004_Wrap),
		new(DPT_1005_Wrap),
		new(DPT_1006_Wrap),
		new(DPT_1007_Wrap),
		new(DPT_1008_Wrap),
		new(DPT_1009_Wrap),
		new(DPT_1010_Wrap),
		new(DPT_1011_Wrap),
		new(DPT_1012_Wrap),
		new(DPT_1013_Wrap),
		new(DPT_1014_Wrap),
		new(DPT_1015_Wrap),
		new(DPT_1016_Wrap),
		new(DPT_1017_Wrap),
		new(DPT_1018_Wrap),
		new(DPT_1019_Wrap),
		new(DPT_1021_Wrap),
		new(DPT_1022_Wrap),
		new(DPT_1023_Wrap),
		new(DPT_1024_Wrap),
		new(DPT_1100_Wrap),
		//5.x
		new(DPT_5001_Wrap),
		new(DPT_5003_Wrap),
		new(DPT_5004_Wrap),
		new(DPT_5005_Wrap),
		//6.x
		new(DPT_6010_Wrap),
		//7.x
		new(DPT_7001_Wrap),
		new(DPT_7002_Wrap),
		new(DPT_7003_Wrap),
		new(DPT_7004_Wrap),
		new(DPT_7005_Wrap),
		new(DPT_7006_Wrap),
		new(DPT_7007_Wrap),
		new(DPT_7010_Wrap),
		new(DPT_7011_Wrap),
		new(DPT_7012_Wrap),
		new(DPT_7013_Wrap),
		//9.x
		new(DPT_9001_Wrap),
		new(DPT_9002_Wrap),
		new(DPT_9003_Wrap),
		new(DPT_9004_Wrap),
		new(DPT_9005_Wrap),
		new(DPT_9006_Wrap),
		new(DPT_9007_Wrap),
		new(DPT_9008_Wrap),
		new(DPT_9010_Wrap),
		new(DPT_9011_Wrap),
		new(DPT_9020_Wrap),
		new(DPT_9021_Wrap),
		new(DPT_9022_Wrap),
		new(DPT_9024_Wrap),
		new(DPT_9025_Wrap),
		new(DPT_9028_Wrap),
		//13.x
		new(DPT_13001_Wrap),
		new(DPT_13002_Wrap),
		new(DPT_13010_Wrap),
		new(DPT_13011_Wrap),
		new(DPT_13013_Wrap),
		new(DPT_13014_Wrap),
		new(DPT_13016_Wrap),
		//14.x
		new(DPT_14033_Wrap),
		new(DPT_14056_Wrap),
		new(DPT_14058_Wrap),
		new(DPT_14059_Wrap),
		new(DPT_14060_Wrap),
		new(DPT_14065_Wrap),
		//17.x
		new(DPT_17001_Wrap),
		//18.x
		new(DPT_18001_Wrap),
	}

	once     sync.Once
	registry map[string]reflect.Type
)

// Init function used to add all types
func setup() {
	// Singleton, can only run once
	once.Do(func() {
		// Register the types
		registry = make(map[string]reflect.Type)
		for _, d := range types {
			// Determine the name of the datatype
			d_type := reflect.TypeOf(d).Elem()
			name := d_type.Name()

			// Convert the name into KNX yy.xxx (e.g. DPT_1001_wrap --> 1.001) format
			name = name[4:len(name)-8] + "." + name[len(name)-8:len(name)-5]

			// Register the type
			registry[name] = d_type
		}
	})
}

func ListSelectEditor() []string {
	setup()
	keys := make([]string, len(types))
	i := 0
	for _, d := range types {
		// Determine the name of the datatype
		d_type := reflect.TypeOf(d).Elem()
		name := d_type.Name()
		// Convert the name into KNX yy.xxx (e.g. DPT_1001_wrap --> 1.001) format
		name = name[4:len(name)-8] + "." + name[len(name)-8:len(name)-5]

		dpt := reflect.New(d_type).Interface().(IDpt)
		keys[i] = fmt.Sprintf("{\"key\":\"%s\",\"value\":\"%s\"}", name+" "+dpt.Name(), name)
		i++
	}

	return keys
}

func ListSupportedTypes() []string {
	// Setup the registry
	setup()

	// Initialize the key-list
	keys := make([]string, len(registry))

	// Fill the key-list
	i := 0
	for k := range registry {
		keys[i] = k
		i++
	}

	return keys
}

func Produce(name string) (d IDpt, ok bool) {
	// Setup the registry
	setup()

	// Lookup the given type and create a new instance of that type
	x, ok := registry[name]
	if ok {
		d = reflect.New(x).Interface().(IDpt)
	}
	return d, ok
}
