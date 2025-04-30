package models

type BodyType int
type StateType int
type WheelSide bool

type Dimensions struct {
	WidthCm     int // ширина
	LengthCm    int // длина
	HeightCm    int // высота
	ClearanceCm int // клиренс
	WheelbaseCm int // колёсная база
	TrackWidthCm int // колея
	MassKg      int  // масса
}

type DrivingParams struct {
	MaxSpeedKph      int     // Максимальная скорость
	AccelerationTo100 float64 // Разгон до 100 км/ч
	FuelConsumptionCity float64
	FuelConsumptionHighway float64
	FuelConsumptionMixed float64
	FuelType          string
	EcoClass          string
	CO2Emissions      int // грамм на км
}

// Можно будет вынести отдельно
// type Wheels struct {}
// EquipmentInfo?
// type SuspensionAndBrakes struct {}

const (
	Liftback BodyType = iota
	Sedan
	SUV
	Hatchback
)

const (
	New StateType = iota
	Used
)

const (
	Right WheelSide = true
	Left  WheelSide = false
)

type Car struct {
	ID                 uint64
	Status             bool
	State              StateType
	VendorID           int
	ModelName          string
	Year               int
	Generation         int
	Trim               string
	DriveType          LocalizedString
	WheelSize          Wheels // убери models::, если внутри одного пакета
	OwnersCount        int
	MileageKm          int
	BodyType           BodyType
	Color              Color
	City               string
	Engine             EngineInfo
	Equipment          EquipmentInfo
	Transmission       LocalizedString
	SteeringWheel      WheelSide
	Condition          LocalizedString
	PTS                LocalizedString
	CustomsCleared     bool
	ExchangePossible   bool
	VIN                string
	TestDriveAvailable bool
	Addons             []string
	Dimensions         Dimensions
	DrivingParams      DrivingParams
}
