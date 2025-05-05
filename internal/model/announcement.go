package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

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
	CarCategory       string
	Color             string
	InteriorMaterial  string
)

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
	WheelSizes  []WheelSize
	BoltPattern BoltPattern
}

type EngineType struct {
	Electric *ElectricEngine
	Fuel     *FuelEngine
	Hybrid   *HybridEngine
}

// TODO::
// Заполнить структуру
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
	D_FWD DriveType = "FWD" // передний привод (Front-Wheel Drive)
	D_RWD DriveType = "RWD" // задний привод (Rear-Wheel Drive)
	D_AWD DriveType = "AWD" // полный привод (All-Wheel Drive)
	D_4WD DriveType = "4WD" // жёстко подключаемый полный привод (Four-Wheel Drive)
)

// TODO::
// Заполнить структуру
type TransmissionInfo struct{}

// TODO::
// Исправить название метариалов
const (
	Leather    SeatMaterial = "leather"
	Alcantara  SeatMaterial = "alcantara"
	Fabric     SeatMaterial = "fabric"
	EcoLeather SeatMaterial = "eco_leather"
	Vinyl      SeatMaterial = "vinyl"
	Combined   SeatMaterial = "combined"
)

type CarModel struct {
	gorm.Model
	ID          uuid.UUID
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	Category    CarCategory
	State       StateType
	VendorID    int
	ModelName   string
	Generation  int
	DriveType   DriveType
	WheelsSize  WheelsInfo
	DoorsNumber int
	SeatsNumber int
	BodyType    BodyType
	EngineType  EngineType
	// Engine             EngineInfo
	Equipment     EquipmentList
	Transmission  TransmissionInfo
	SteeringWheel SteeringWheelSide
	Dimensions    Dimensions
	DrivingParams DrivingParams
	Options       OptionsList
}

type CarToSell struct {
	gorm.Model
	ID                 uuid.uuid
	SellerID           uint64
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          time.Time
	Images             []string
	Year               int
	MileageKm          int
	OwnersCount        int
	EnsuranceValidTill *time.Time
	InteriorMaterial   InteriorMaterial
	LicensePlateNumber *string
	Color              Color
	CarModel           CarModel
	ExchangePossible   bool
	TestDriveAvailable bool
	VIN                string
	CustomsCleared     bool
	PTS                string // Есть такое вообще?
	Addons             *AddonsList
	City               int
	LastService        *time.Time
	Condition          StateType
}
