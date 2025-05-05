package model

type ElectricEngine struct {
	PowerKw               *int            `json:"power_kw,omitempty"`             // мощность, кВт
	TorqueNm              *int            `json:"torque_nm,omitempty"`            // крутящий момент, Н·м
	BatteryCapacityKWh    *float64        `json:"battery_capacity_kwh,omitempty"` // ёмкость батареи, кВт·ч
	Voltage               *int            `json:"voltage,omitempty"`              // напряжение системы
	FullChargingTimeHours *float64        `json:"charging_time_hours,omitempty"`  // время зарядки
	FastCharging          *bool           `json:"fast_charging,omitempty"`        // поддержка быстрой зарядки
	OneChargeResource     *int            `json:"wltp_range_km,omitempty"`        // запас хода (WLTP), км
	MotorLocation         *string `json:"motor_location,omitempty"`       // расположение электромотора (на передней/задней оси)
}

type FuelEngine struct {
	VolumeLiters      *float64        `json:"volume_liters,omitempty"`       // объём двигателя, л
	PowerHp           *int            `json:"power_hp,omitempty"`            // мощность, л.с.
	TorqueNm          *int            `json:"torque_nm,omitempty"`           // крутящий момент, Н·м
	Cylinders         *int            `json:"cylinders,omitempty"`           // количество цилиндров
	CylinderLayout    *string `json:"cylinder_layout,omitempty"`     // расположение цилиндров (рядное, V-образное и т.д.)
	ValvesPerCylinder *int            `json:"valves_per_cylinder,omitempty"` // клапанов на цилиндр
	CompressionRatio  *float64        `json:"compression_ratio,omitempty"`   // степень сжатия
	FuelType          *string `json:"fuel_type,omitempty"`           // тип топлива: бензин, дизель
	InjectionType     *string `json:"injection_type,omitempty"`      // тип впрыска
	Turbocharged      *bool           `json:"turbocharged,omitempty"`        // турбонаддув
	Supercharged      *bool           `json:"supercharged,omitempty"`        // компрессор
	EmissionStandard  *string `json:"emission_standard,omitempty"`   // экостандарт: Euro-5 и т.д.
	CoolingType       *string `json:"cooling_type,omitempty"`        // тип охлаждения
	EngineCode        *string         `json:"engine_code,omitempty"`         // код двигателя
}

type HybridEngine struct {
	FuelEnginePart      *FuelEngine     `json:"fuel_engine_part,omitempty"`     // ДВС часть
	ElectricMotorPart   *ElectricEngine  `json:"electric_motor_part,omitempty"`  // Электрическая часть
	CombinedPowerHp     *int            `json:"combined_power_hp,omitempty"`    // суммарная мощность, л.с.
	CombinedTorqueNm    *int            `json:"combined_torque_nm,omitempty"`   // суммарный момент, Н·м
	HybridType          *string `json:"hybrid_type,omitempty"`          // тип: mild, full, plug-in
	RegenerativeBraking *bool           `json:"regenerative_braking,omitempty"` // рекуперация
}


