package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
	"gorm.io/plugin/soft_delete"
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

// type jsonTime struct {
// 	null.Time
// }

// func (j jsonTime) format() string {
// 	fmt.Println("***")
// 	fmt.Println(j)
// 	fmt.Println(j.Time.Time.Format("2006-01-02"))
// 	fmt.Println("***")
// 	if j.Valid {
// 		return j.NullTime.Time.Format("2006-01-02")
// 	} else {
// 		return "null"
// 	}
// 	// return null.NewString(j.NullTime.Time.Format("2006-01-02"), j.Valid)
// 	// if !j.NullTime.Valid {
// 	// 	return nil
// 	// }
// 	// return j.Time.Time.Format("2006-01-02")
// }

// func (j jsonTime) MarshalJSON() ([]byte, error) {
// 	if j.format() == "null" {
// 		return []byte(j.format()), nil
// 	} else {
// 		return []byte(`"` + j.format() + `"`), nil
// 	}
// }

// クルマ
type Car struct {
	// gorm.Model
	Id              uint                  `json:"id" gorm:"primarykey"`
	IsDel           soft_delete.DeletedAt `json:"is_del" gorm:"softDelete:flag"`
	CreatedAt       time.Time             `json:"created_at"`
	UpdatedAt       time.Time             `json:"updated_at"`
	MakerName       string                `json:"maker_name" gorm:"not null"`                                // メーカー名
	ModelName       string                `json:"model_name" gorm:"not null"`                                // モデル名
	GradeName       string                `json:"grade_name" gorm:"not null"`                                // グレード
	ModelCode       string                `json:"model_code" gorm:"not null"`                                // 型式
	Price           null.Int              `json:"price"`                                                     // 小売価格(税込/円)
	Url             null.String           `json:"url"`                                                       // URL
	ImageUrl        null.String           `json:"image_url"`                                                 // イメージURL
	ModelChangeFull null.Time             `json:"model_change_full" gorm:"type:date"`                        // モデルチェンジ時期(フル日本)
	ModelChangeLast null.Time             `json:"model_change_last" gorm:"type:date"`                        // モデルチェンジ時期(最終日本)
	Body            Body                  `json:"body" gorm:"embedded;embeddedPrefix:body_"`                 // 車体
	Interior        Interior              `json:"interior" gorm:"embedded;embeddedPrefix:int_"`              // 車内
	Perf            Perf                  `json:"perf" gorm:"embedded;embeddedPrefix:perf_"`                 // 性能
	PowerTrain      null.String           `json:"power_train"`                                               // パワートレイン(ICE/StrHV/MldHV/SerHV/PHEV/BEV/RexEV/FCEV)
	DriveSystem     null.String           `json:"drive_system"`                                              // 駆動方式(FF/FR/RR/MR/AWD)
	Engine          Engine                `json:"engine" gorm:"embedded;embeddedPrefix:engine_"`             // エンジン
	MotorX          Motor                 `json:"motor_x" gorm:"embedded;embeddedPrefix:motor_x_"`           // 電動機1
	MotorY          Motor                 `json:"motor_y" gorm:"embedded;embeddedPrefix:motor_y_"`           // 電動機2
	Battery         Battery               `json:"battery" gorm:"embedded;embeddedPrefix:battery_"`           // バッテリー
	Steering        null.String           `json:"steering"`                                                  // ステアリング形式
	SuspensionFront null.String           `json:"suspension_front"`                                          // サスペンション形式前
	SuspensionRear  null.String           `json:"suspension_rear"`                                           // サスペンション形式後
	BrakeFront      null.String           `json:"brake_front"`                                               // ブレーキ形式前
	BrakeRear       null.String           `json:"brake_rear"`                                                // ブレーキ形式後
	TireFront       Tire                  `json:"tire_front" gorm:"embedded;embeddedPrefix:tire_front_"`     // タイヤ前
	TireRear        Tire                  `json:"tire_rear" gorm:"embedded;embeddedPrefix:tire_rear_"`       // タイヤ後
	Transmission    Transmission          `json:"transmission" gorm:"embedded;embeddedPrefix:transmission_"` // トランスミッション
	FuelEfficiency  null.String           `json:"fuel_efficiency"`                                           // 燃費向上対策
}

// 車体
type Body struct {
	Length           null.Int `json:"length"`             // 全長(mm)
	Width            null.Int `json:"width"`              // 全幅(mm)
	Height           null.Int `json:"height"`             // 全高(mm)
	WheelBase        null.Int `json:"wheel_base"`         // ホイールベース(mm)
	TreadFront       null.Int `json:"tread_front"`        // トレッド前(mm)
	TreadRear        null.Int `json:"tread_rear"`         // トレッド後(mm)
	MinRoadClearance null.Int `json:"min_road_clearance"` // 最低地上高(mm)
	Weight           null.Int `json:"body_weight"`        // 車両重量(kg)
}

// 車内
type Interior struct {
	Length     null.Int `json:"length"`      // 室内長(mm)
	Width      null.Int `json:"width"`       // 室内幅(mm)
	Height     null.Int `json:"height"`      // 室内高(mm)
	LuggageCap null.Int `json:"luggage_cap"` // ラゲッジルーム容量(L)
	RidingCap  null.Int `json:"riding_cap"`  // 乗車定員(人)
}

// 性能
type Perf struct {
	MinTurningRadius null.Float `json:"min_turning_radius"` // 最小回転半径(m)
	FcrWltc          null.Float `json:"fcr_wltc"`           // 燃料消費率WLTCモード(km/L)
	FcrWltcL         null.Float `json:"fcr_wltc_l"`         // 燃料消費率WLTC市街地モード(km/L)
	FcrWltcM         null.Float `json:"fcr_wltc_m"`         // 燃料消費率WLTC郊外モード(km/L)
	FcrWltcH         null.Float `json:"fcr_wltc_h"`         // 燃料消費率WLTC高速道路モード(km/L)
	FcrWltcExh       null.Float `json:"fcr_wltc_exh"`       // 燃料消費率WLTC高高速道路モード(km/L)
	FcrJc08          null.Float `json:"fcr_jc08"`           // 燃料消費率JC08モード(km/L)
	MpcWltc          null.Float `json:"mpc_wltc"`           // 一充電走行距離WLTCモード(km)
	EcrWltc          null.Float `json:"ecr_wltc"`           // 交流電力消費率WTLCモード(Wh/km)
	EcrWltcL         null.Float `json:"ecr_wltc_l"`         // 交流電力消費率WLTC市街地モード(Wh/km)
	EcrWltcM         null.Float `json:"ecr_wltc_m"`         // 交流電力消費率WLTC郊外モード(Wh/km)
	EcrWltcH         null.Float `json:"ecr_wltc_h"`         // 交流電力消費率WLTC高速道路モード(Wh/km)
	EcrWltcExh       null.Float `json:"ecr_wltc_exh"`       // 交流電力消費率WLTC高高速道路モード(Wh/km)
	EcrJc08          null.Float `json:"ecr_jc08"`           // 交流電力消費率JC08モード(Wh/km)
	MpcJc08          null.Float `json:"mpc_jc08"`           // 一充電走行距離JC08モード(km)
}

// エンジン
type Engine struct {
	Code               null.String `json:"code"`                  // 型式
	Type               null.String `json:"type"`                  // 種類
	Cylinders          null.Int    `json:"cylinders"`             // 気筒数
	CylinderLayout     null.String `json:"cylinder_layout"`       // シリンダーレイアウト(I/V/B/W)
	ValveSystem        null.String `json:"valve_system"`          // バルブ構造(SV/OHV/SOHC/DOHC)
	Displacement       null.Float  `json:"displacement"`          // 総排気量(L)
	Bore               null.Float  `json:"bore"`                  // ボア(mm)
	Stroke             null.Float  `json:"stroke"`                // ストローク(mm)
	CompRatio          null.Float  `json:"comp_ratio"`            // 圧縮比
	MaxOutput          null.Float  `json:"max_output"`            // 最高出力(kW)
	MaxOutputLowerRpm  null.Float  `json:"max_output_lower_rpm"`  // 最高出力回転数(低)(rpm)
	MaxOutputHigherRpm null.Float  `json:"max_output_higher_rpm"` // 最高出力回転数(高)(rpm)
	MaxTorque          null.Float  `json:"max_torque"`            // 最大トルク(Nm)
	MaxTorqueLowerRpm  null.Float  `json:"max_torque_lower_rpm"`  // 最大トルク回転数(低)(rpm)
	MaxTorqueHigherRpm null.Float  `json:"max_torque_higher_rpm"` // 最大トルク回転数(高)(rpm)
	FuelSystem         null.String `json:"fuel_system"`           // 燃料供給装置
	FuelType           null.String `json:"fuel_type"`             // 使用燃料種類(軽油/無鉛レギュラーガソリン/無鉛プレミアムガソリン)
	FuelTankCap        null.Int    `json:"fuel_tank_cap"`         // 燃料タンク容量(L)
}

// 電動機
type Motor struct {
	Code               null.String `json:"code"`                  // 型式
	Type               null.String `json:"type"`                  // 種類
	Purpose            null.String `json:"purpose"`               // 用途(動力前用/動力後用/発電用)
	RatedOutput        null.Float  `json:"rated_output"`          // 定格出力(kW)
	MaxOutput          null.Float  `json:"max_output"`            // 最高出力(kW)
	MaxOutputLowerRpm  null.Float  `json:"max_output_lower_rpm"`  // 最高出力回転数(低)(rpm)
	MaxOutputHigherRpm null.Float  `json:"max_output_higher_rpm"` // 最高出力回転数(高)(rpm)
	MaxTorque          null.Float  `json:"max_torque"`            // 最大トルク(Nm)
	MaxTorqueLowerRpm  null.Float  `json:"max_torque_lower_rpm"`  // 最大トルク回転数(低)(rpm)
	MaxTorqueHigherRpm null.Float  `json:"max_torque_higher_rpm"` // 最大トルク回転数(高)(rpm)
}

// バッテリー
type Battery struct {
	Type     null.String `json:"type"`     // 種類
	Quantity null.Int    `json:"quantity"` // 個数
	Voltage  null.Float  `json:"voltage"`  // 電圧(V)
	Capacity null.Float  `json:"capacity"` // 容量(Ah)
}

// タイヤ
type Tire struct {
	SectionWidth  null.Int `json:"section_width"`  // タイヤ幅(mm)
	AspectRatio   null.Int `json:"aspect_ratio"`   // 扁平率(%)
	WheelDiameter null.Int `json:"wheel_diameter"` // ホイール径(インチ)
}

// トランスミッション
type Transmission struct {
	Type                null.String `json:"type"`                  // AT/DCT/AMT/MT/CVT
	Gears               null.Int    `json:"gears"`                 // 段数
	Ratio1              null.Float  `json:"ratio_1"`               // 変速比1速
	Ratio2              null.Float  `json:"ratio_2"`               // 変速比2速
	Ratio3              null.Float  `json:"ratio_3"`               // 変速比3速
	Ratio4              null.Float  `json:"ratio_4"`               // 変速比4速
	Ratio5              null.Float  `json:"ratio_5"`               // 変速比5速
	Ratio6              null.Float  `json:"ratio_6"`               // 変速比6速
	Ratio7              null.Float  `json:"ratio_7"`               // 変速比7速
	Ratio8              null.Float  `json:"ratio_8"`               // 変速比8速
	Ratio9              null.Float  `json:"ratio_9"`               // 変速比9速
	Ratio10             null.Float  `json:"ratio_10"`              // 変速比10速
	RatioRear           null.Float  `json:"ratio_rear"`            // 変速比後退
	ReductionRatioFront null.Float  `json:"reduction_ratio_front"` // 減速比フロント
	ReductionRatioRear  null.Float  `json:"reduction_ratio_rear"`  // 減速比リア
}

type CarsQuery struct {
	MakerName       string   `json:"maker_name"`
	ModelName       string   `json:"model_name"`        //null.String
	GradeName       string   `json:"grade_name"`        //null.String
	ModelCode       string   `json:"model_code"`        //null.String
	PriceLower      int      `json:"price_lower"`       //null.Int
	PriceHigher     int      `json:"price_higher"`      //null.Int
	ModelChangeFrom string   `json:"model_change_from"` //sql.NullTime
	ModelChangeTo   string   `json:"model_change_to"`   //sql.NullTime
	PowerTrain      []string `json:"power_train"`
}

func SearchCars(c *gin.Context) {
	// HTTPリクエストのペイロードを取得
	var query CarsQuery
	c.BindJSON(&query)
	fmt.Println(query)
	var cars []Car
	d := DB.Debug().Table("cars").Where(
		&Car{
			MakerName: query.MakerName,
			ModelName: query.ModelName,
			GradeName: query.GradeName,
			ModelCode: query.ModelCode,
		},
	)
	if query.PriceLower != 0 {
		d.Where(
			"price > ?", query.PriceLower,
		)
	}
	if query.PriceHigher != 0 {
		d.Where(
			"price < ?", query.PriceHigher,
		)
	}
	if query.ModelChangeFrom != "" {
		d.Where(
			"(model_change_last > ? OR model_change_full > ?)",
			query.ModelChangeFrom,
			query.ModelChangeFrom,
		)
	}
	if query.ModelChangeTo != "" {
		d.Where(
			"(model_change_last < ? OR model_change_full < ?)",
			query.ModelChangeTo,
			query.ModelChangeTo,
		)
	}
	d.Find(&cars)
	fmt.Printf("%+v", cars)
	c.IndentedJSON(http.StatusOK, cars)
}

func GetCar(c *gin.Context) {
	var car Car
	result := DB.Table("cars").First(&car, c.Param("id"))
	if result.RowsAffected != 1 {
		c.IndentedJSON(
			http.StatusNotFound,
			&ErrorResp{Message: result.Error.Error()},
		)
		c.Abort()
		return
	}
	c.IndentedJSON(http.StatusOK, car)
}
