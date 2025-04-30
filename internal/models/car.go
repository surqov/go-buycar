package models

type Car struct {
  ID               uint64 `json:"id"`
  Status            bool `json:"status"`   
  State             LocalizedString
  VendorID         int
  ModelName         string
  Year             int            `json:"year"`
  Generation       string         `json:"generation"`     
  OwnersCount      int            `json:"owners_count"`
  MileageKm        int            `json:"mileage_km"`      
  BodyType         BodyType       `json:"body_type"`       
  Color            Color            `json:"color"`           
  City              string
  Engine           models::EngineInfo      `json:"engine"`
  Equipment        LocalizedString `json:"equipment"`       
  Transmission     LocalizedString `json:"transmission"`    
  DriveType        LocalizedString `json:"drive_type"`      
  SteeringWheel    LocalizedString `json:"steering_wheel"`  
  Condition        LocalizedString `json:"condition"`       
  PTS              LocalizedString `json:"pts"`             
  CustomsCleared   bool           `json:"customs_cleared"` 
  ExchangePossible bool           `json:"exchange_possible"`
  VIN              string         `json:"vin"`
  TestDrive        bool
}

const (
  New StateType
  Used StateType
)

const (
    Sedan
    
    Liftback BodyType = 
    Sedan    BodyType = "седан"
    SUV      BodyType = "внедорожник"
    Hatchback BodyType = "хэтчбек"
    
)

type Dimensions struct {
  width_cm int
  length_cm int
  height_cm int
}


