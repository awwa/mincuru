package main

import (
	"database/sql"

	"gorm.io/gorm"
)

type PowerTrain string

const (
	ICE   = PowerTrain("ICE")   // エンジン車
	StrHV = PowerTrain("StrHV") // ストロングハイブリッド
	MldHV = PowerTrain("MldHV") // マイルドハイブリッド
	SerHV = PowerTrain("SerHV") // シリーズハイブリッド
	PHEV  = PowerTrain("PHEV")  // プラグインハイブリッド
	BEV   = PowerTrain("BEV")   // バッテリーEV
	RexEV = PowerTrain("RexEV") // レンジエクステンダーEV
	FCEV  = PowerTrain("FCEV")  // 燃料電池車
)

type DriveSystem string

const (
	FF  = DriveSystem("FF")
	FR  = DriveSystem("FR")
	RR  = DriveSystem("RR")
	MR  = DriveSystem("MR")
	AWD = DriveSystem("AWD")
)

type CylinderLayout string

const (
	V = CylinderLayout("V") // V型
	I = CylinderLayout("I") // 直列
	B = CylinderLayout("B") // 水平対向
	W = CylinderLayout("W") // W型
)

type ValveSystem string

const (
	SV   = ValveSystem("SV")
	OHV  = ValveSystem("OHV")
	SOHC = ValveSystem("SOHC")
	DOHC = ValveSystem("DOHC")
)

type TransmissionType string

const (
	AT  = TransmissionType("AT")
	DCT = TransmissionType("DCT")
	AMT = TransmissionType("AMT")
	MT  = TransmissionType("MT")
	CVT = TransmissionType("CVT")
	PG  = TransmissionType("電気式無段変速機")
)

type MotorPurpose string

const (
	TRACTION_FRONT = MotorPurpose("走行用前")
	TRACTION_REAR  = MotorPurpose("走行用後")
	GENERATOR      = MotorPurpose("発電用")
)

// クルマ
type Car struct {
	gorm.Model
	MakerName       string         `json:"maker_name" gorm:"not null"`            // メーカー名
	ModelName       string         `json:"model_name" gorm:"not null"`            // モデル名
	GradeName       string         `json:"grade_name" gorm:"not null"`            // グレード
	ModelCode       sql.NullString `json:"model_code"`                            // 型式
	Price           sql.NullInt64  `json:"price"`                                 // 小売価格(税込/円)
	Url             sql.NullString `json:"url"`                                   // URL
	ImageUrl        sql.NullString `json:"image_url"`                             // イメージURL
	ModelChangeFull sql.NullTime   `json:"model_change_full" gorm:"type:date"`    // モデルチェンジ時期(フル日本)
	ModelChangeLast sql.NullTime   `json:"model_change_last" gorm:"type:date"`    // モデルチェンジ時期(最終日本)
	Body            Body           `gorm:"embedded;embeddedPrefix:body_"`         // 車体
	Interior        Interior       `gorm:"embedded;embeddedPrefix:int_"`          // 車内
	Perf            Perf           `gorm:"embedded;embeddedPrefix:perf_"`         // 性能
	PowerTrain      sql.NullString `json:"power_train"`                           // パワートレイン(ICE/StrHV/MldHV/SerHV/PHEV/BEV/RexEV/FCEV)
	DriveSystem     sql.NullString `json:"drive_system"`                          // 駆動方式(FF/FR/RR/MR/AWD)
	Engine          Engine         `gorm:"embedded;embeddedPrefix:engine_"`       // エンジン
	MotorX          Motor          `gorm:"embedded;embeddedPrefix:motor_x_"`      // 電動機1
	MotorY          Motor          `gorm:"embedded;embeddedPrefix:motor_y_"`      // 電動機2
	Battery         Battery        `gorm:"embedded;embeddedPrefix:battery_"`      // バッテリー
	Steering        sql.NullString `json:"steering"`                              // ステアリング形式
	SuspensionFront sql.NullString `json:"suspension_front"`                      // サスペンション形式前
	SuspensionRear  sql.NullString `json:"suspension_rear"`                       // サスペンション形式後
	BrakeFront      sql.NullString `json:"brake_front"`                           // ブレーキ形式前
	BrakeRear       sql.NullString `json:"brake_rear"`                            // ブレーキ形式後
	TireFront       Tire           `gorm:"embedded;embeddedPrefix:tire_front_"`   // タイヤ前
	TireRear        Tire           `gorm:"embedded;embeddedPrefix:tire_rear_"`    // タイヤ後
	Transmission    Transmission   `gorm:"embedded;embeddedPrefix:transmission_"` // トランスミッション
	FuelEfficiency  sql.NullString `json:"fuel_efficiency"`                       // 燃費向上対策
}

// 車体
type Body struct {
	Length           sql.NullInt64 `json:"length"`             // 全長(mm)
	Width            sql.NullInt64 `json:"width"`              // 全幅(mm)
	Height           sql.NullInt64 `json:"height"`             // 全高(mm)
	WheelBase        sql.NullInt64 `json:"wheel_base"`         // ホイールベース(mm)
	TreadFront       sql.NullInt64 `json:"tread_front"`        // トレッド前(mm)
	TreadRear        sql.NullInt64 `json:"tread_rear"`         // トレッド後(mm)
	MinRoadClearance sql.NullInt64 `json:"min_road_clearance"` // 最低地上高(mm)
	Weight           sql.NullInt64 `json:"body_weight"`        // 車両重量(kg)
}

// 車内
type Interior struct {
	Length     sql.NullInt64 `json:"length"`      // 室内長(mm)
	Width      sql.NullInt64 `json:"width"`       // 室内幅(mm)
	Height     sql.NullInt64 `json:"height"`      // 室内高(mm)
	LuggageCap sql.NullInt64 `json:"luggage_cap"` // ラゲッジルーム容量(L)
	RidingCap  sql.NullInt64 `json:"riding_cap"`  // 乗車定員(人)
}

// 性能
type Perf struct {
	MinTurningRadius sql.NullFloat64 `json:"min_turning_radius"` // 最小回転半径(m)
	FcrWltc          sql.NullFloat64 `json:"fcr_wltc"`           // 燃料消費率WLTCモード(km/L)
	FcrWltcL         sql.NullFloat64 `json:"fcr_wltc_l"`         // 燃料消費率WLTC市街地モード(km/L)
	FcrWltcM         sql.NullFloat64 `json:"fcr_wltc_m"`         // 燃料消費率WLTC郊外モード(km/L)
	FcrWltcH         sql.NullFloat64 `json:"fcr_wltc_h"`         // 燃料消費率WLTC高速道路モード(km/L)
	FcrWltcExh       sql.NullFloat64 `json:"fcr_wltc_exh"`       // 燃料消費率WLTC高高速道路モード(km/L)
	FcrJc08          sql.NullFloat64 `json:"fcr_jc08"`           // 燃料消費率JC08モード(km/L)
	MpcWltc          sql.NullFloat64 `json:"mpc_wltc"`           // 一充電走行距離WLTCモード(km)
	EcrWltc          sql.NullFloat64 `json:"ecr_wltc"`           // 交流電力消費率WTLCモード(Wh/km)
	EcrWltcL         sql.NullFloat64 `json:"ecr_wltc_l"`         // 交流電力消費率WLTC市街地モード(Wh/km)
	EcrWltcM         sql.NullFloat64 `json:"ecr_wltc_m"`         // 交流電力消費率WLTC郊外モード(Wh/km)
	EcrWltcH         sql.NullFloat64 `json:"ecr_wltc_h"`         // 交流電力消費率WLTC高速道路モード(Wh/km)
	EcrWltcExh       sql.NullFloat64 `json:"ecr_wltc_exh"`       // 交流電力消費率WLTC高高速道路モード(Wh/km)
	EcrJc08          sql.NullFloat64 `json:"ecr_jc08"`           // 交流電力消費率JC08モード(Wh/km)
	MpcJc08          sql.NullFloat64 `json:"mpc_jc08"`           // 一充電走行距離JC08モード(km)
}

// エンジン
type Engine struct {
	Code               sql.NullString  `json:"code"`                  // 型式
	Type               sql.NullString  `json:"type"`                  // 種類
	Cylinders          sql.NullInt64   `json:"cylinders"`             // 気筒数
	CylinderLayout     sql.NullString  `json:"cylinder_layout"`       // シリンダーレイアウト(I/V/B/W)
	ValveSystem        sql.NullString  `json:"valve_system"`          // バルブ構造(SV/OHV/SOHC/DOHC)
	Displacement       sql.NullFloat64 `json:"displacement"`          // 総排気量(L)
	Bore               sql.NullFloat64 `json:"bore"`                  // ボア(mm)
	Stroke             sql.NullFloat64 `json:"stroke"`                // ストローク(mm)
	CompRatio          sql.NullFloat64 `json:"compression_ratio"`     // 圧縮比
	MaxOutput          sql.NullFloat64 `json:"max_output"`            // 最高出力(kW)
	MaxOutputLowerRpm  sql.NullFloat64 `json:"max_output_lower_rpm"`  // 最高出力回転数(低)(rpm)
	MaxOutputHigherRpm sql.NullFloat64 `json:"max_output_higher_rpm"` // 最高出力回転数(高)(rpm)
	MaxTorque          sql.NullFloat64 `json:"max_torque"`            // 最大トルク(Nm)
	MaxTorqueLowerRpm  sql.NullFloat64 `json:"max_torque_lower_rpm"`  // 最大トルク回転数(低)(rpm)
	MaxTorqueHigherRpm sql.NullFloat64 `json:"max_torque_higher_rpm"` // 最大トルク回転数(高)(rpm)
	FuelSystem         sql.NullString  `json:"fuel_system"`           // 燃料供給装置
	FuelType           sql.NullString  `json:"fuel_type"`             // 使用燃料種類(軽油/無鉛レギュラーガソリン/無鉛プレミアムガソリン)
	FuelTankCap        sql.NullInt64   `json:"fuel_tank_cap"`         // 燃料タンク容量(L)
}

// 電動機
type Motor struct {
	Code               sql.NullString  `json:"code"`                  // 型式
	Type               sql.NullString  `json:"type"`                  // 種類
	Purpose            sql.NullString  `json:"purpose"`               // 用途(動力前用/動力後用/発電用)
	RatedOutput        sql.NullFloat64 `json:"rated_output"`          // 定格出力(kW)
	MaxOutput          sql.NullFloat64 `json:"max_output"`            // 最高出力(kW)
	MaxOutputLowerRpm  sql.NullFloat64 `json:"max_output_lower_rpm"`  // 最高出力回転数(低)(rpm)
	MaxOutputHigherRpm sql.NullFloat64 `json:"max_output_higher_rpm"` // 最高出力回転数(高)(rpm)
	MaxTorque          sql.NullFloat64 `json:"max_torque"`            // 最大トルク(Nm)
	MaxTorqueLowerRpm  sql.NullFloat64 `json:"max_torque_lower_rpm"`  // 最大トルク回転数(低)(rpm)
	MaxTorqueHigherRpm sql.NullFloat64 `json:"max_torque_higher_rpm"` // 最大トルク回転数(高)(rpm)
}

// バッテリー
type Battery struct {
	Type     sql.NullString  `json:"type"`     // 種類
	Quantity sql.NullInt64   `json:"quantity"` // 個数
	Voltage  sql.NullFloat64 `json:"voltage"`  // 電圧(V)
	Capacity sql.NullFloat64 `json:"capacity"` // 容量(Ah)
}

// タイヤ
type Tire struct {
	SectionWidth  sql.NullInt64 `json:"section_width"`  // タイヤ幅(mm)
	AspectRatio   sql.NullInt64 `json:"aspect_ratio"`   // 扁平率(%)
	WheelDiameter sql.NullInt64 `json:"wheel_diameter"` // ホイール径(インチ)
}

// トランスミッション
type Transmission struct {
	Type                sql.NullString  `json:"type"`                  // AT/DCT/AMT/MT/CVT
	Gears               sql.NullInt64   `json:"gears"`                 // 段数
	Ratio1              sql.NullFloat64 `json:"ratio_1"`               // 変速比1速
	Ratio2              sql.NullFloat64 `json:"ratio_2"`               // 変速比2速
	Ratio3              sql.NullFloat64 `json:"ratio_3"`               // 変速比3速
	Ratio4              sql.NullFloat64 `json:"ratio_4"`               // 変速比4速
	Ratio5              sql.NullFloat64 `json:"ratio_5"`               // 変速比5速
	Ratio6              sql.NullFloat64 `json:"ratio_6"`               // 変速比6速
	Ratio7              sql.NullFloat64 `json:"ratio_7"`               // 変速比7速
	Ratio8              sql.NullFloat64 `json:"ratio_8"`               // 変速比8速
	Ratio9              sql.NullFloat64 `json:"ratio_9"`               // 変速比9速
	Ratio10             sql.NullFloat64 `json:"ratio_10"`              // 変速比10速
	RatioRear           sql.NullFloat64 `json:"ratio_rear"`            // 変速比後退
	ReductionRatioFront sql.NullFloat64 `json:"reduction_ratio_front"` // 減速比フロント
	ReductionRatioRear  sql.NullFloat64 `json:"reduction_ratio_rear"`  // 減速比リア
}
