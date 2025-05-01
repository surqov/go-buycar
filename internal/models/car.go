package models

type (
	BodyType          string
	StateType         int
	WheelSide         bool
	SteeringWheelSide bool
	DriveType         string
	AddonsList        []string
	EquipmentList     []string
	OptionsList       []string
	SeatMaterial      string
)

type Car struct {
	ID                 uint64
	Images             []string
	Status             bool
	State              StateType
	VendorID           int
	ModelName          string
	Generation         int
	Year               int
	OwnersCount        int
	DriveType          DriveType
	WheelsSize         WheelsInfo
	InteriorMaterial   InteriorMaterial
	MileageKm          int
	DoorsNumber        int
	SeatsNumber        int
	BodyType           BodyType
	Color              Color
	City               string
	EngineType         EngineType
	Engine             EngineInfo
	Equipment          EquipmentList
	Transmission       TransmissionInfo
	SteeringWheel      SteeringWheelSide
	Condition          string
	PTS                string
	CustomsCleared     bool
	ExchangePossible   bool
	VIN                string
	TestDriveAvailable bool
	Dimensions         Dimensions
	DrivingParams      DrivingParams
	Options            OptionsList
	Addons             *AddonsList
}

type Dimensions struct {
	WidthCm      int // ширина
	LengthCm     int // длина
	HeightCm     int // высота
	ClearanceCm  int // клиренс
	WheelbaseCm  int // колёсная база
	TrackWidthCm int // колея
	MassKg       int // масса
}

type DrivingParams struct {
	MaxSpeedKph            int     // Максимальная скорость
	AccelerationTo100      float64 // Разгон до 100 км/ч
	FuelConsumptionCity    float64
	FuelConsumptionHighway float64
	FuelConsumptionMixed   float64
	FuelType               string
	EcoClass               string
	CO2Emissions           int // грамм на км
}

type WheelSize struct {
	Width    float64 // в дюймах, например 8.5
	Diameter int     // диаметр в дюймах, например 19
	ET       int     // вылет диска (ET), например 32
}

type TireSize struct {
	Width_mm      int // мм (e.g. 255)
	AspectRatio   int // % (e.g. 45)
	Diameter_inch int // дюймы (e.g. 18)
}

type BoltPattern struct {
	PCD        string // пример: "5x112"
	CenterBore string // пример: "66.6" (мм)
}

type WheelsInfo struct {
	TireSizes   []TireSize
	WheelSizes  []WheesSize
	BoltPattern BoltPattern
}

// TODO::
type SuspensionAndBrakes struct{}

const (
	Sedan     BodyType = "Sedan"
	Hatchback BodyType = "Hatchback"
	Liftback  BodyType = "Liftback"
	SUV       BodyType = "SUV"
	Wagon     BodyType = "Wagon"
	Coupe     BodyType = "Coupe"
	Minivan   BodyType = "Minivan"
	Pickup    BodyType = "Pickup"
	Limousine BodyType = "Limousine"
	Van       BodyType = "Van"
	Cabriolet BodyType = "Cabriolet"
)

const (
	New StateType = iota
	Used
	NeedRepair
  ForSpares
)

const (
	Left  WheelSide = true
	Right WheelSide = false
)

const (
	DriveTypeFWD DriveType = "FWD" // передний привод (Front-Wheel Drive)
	DriveTypeRWD DriveType = "RWD" // задний привод (Rear-Wheel Drive)
	DriveTypeAWD DriveType = "AWD" // полный привод (All-Wheel Drive)
	DriveType4WD DriveType = "4WD" // жёстко подключаемый полный привод (Four-Wheel Drive)
)

// TODO::
type TransmissionInfo struct{}

// TODO::
const (
	Leather    SeatMaterial = "leather"
	Alcantara  SeatMaterial = "alcantara"
	Fabric     SeatMaterial = "fabric"
	EcoLeather SeatMaterial = "eco_leather"
	Vinyl      SeatMaterial = "vinyl"
)
