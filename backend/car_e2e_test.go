package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
)

const FORMAT = "2006-Jan-02"

func seedTestCarCx5() {
	modelChangeFull, _ := time.Parse(FORMAT, "2016-Dec-15")
	modelChangeLast, _ := time.Parse(FORMAT, "2018-Jan-01")
	DB.Create(&Car{
		MakerName:       "マツダ",
		ModelName:       "CX-5",
		GradeName:       "25S Proactive",
		ModelCode:       "6BA-KF5P",
		Price:           null.NewInt(3140500, true),
		Url:             null.NewString("https://www.mazda.co.jp/cars/cx-5/", true),
		ImageUrl:        null.NewString("https://upload.wikimedia.org/wikipedia/commons/8/85/2017_Mazda_CX-5_%28KF%29_Maxx_2WD_wagon_%282018-11-02%29_01.jpg", true),
		ModelChangeFull: null.NewTime(modelChangeFull, true),
		ModelChangeLast: null.NewTime(modelChangeLast, true),
		Body: Body{
			Type:             null.NewString((string)(SUV), true),
			Length:           null.NewInt(4545, true),
			Width:            null.NewInt(1840, true),
			Height:           null.NewInt(1690, true),
			WheelBase:        null.NewInt(2700, true),
			TreadFront:       null.NewInt(1595, true),
			TreadRear:        null.NewInt(1595, true),
			MinRoadClearance: null.NewInt(210, true),
			Weight:           null.NewInt(1620, true),
			Doors:            null.NewInt(4, true),
		},
		Interior: Interior{
			Length: null.NewInt(1890, true),
			Width:  null.NewInt(1540, true),
			Height: null.NewInt(1265, true),
			// LuggageCap: null.NewInt(0, false),
			RidingCap: null.NewInt(5, true),
		},
		Performance: Performance{
			MinTurningRadius: null.NewFloat(5.5, true),
			FcrWltc:          null.NewFloat(13.0, true),
			FcrWltcL:         null.NewFloat(10.2, true),
			FcrWltcM:         null.NewFloat(13.4, true),
			FcrWltcH:         null.NewFloat(14.7, true),
			//FcrWltcExh:       null.NewFloat(0, false),
			FcrJc08: null.NewFloat(14.2, true),
			//MpcWltc:          null.NewFloat(0, false),
			//EcrWltc:          null.NewFloat(0, false),
			//EcrWltcL:         null.NewFloat(0, false),
			//EcrWltcM:         null.NewFloat(0, false),
			//EcrWltcH:         null.NewFloat(0, false),
			//EcrWltcExh:       null.NewFloat(0, false),
			//EcrJc08:          null.NewFloat(0, false),
			//MpcJc08:          null.NewFloat(0, false),
		},
		PowerTrain:  null.NewString((string)(ICE), true),
		DriveSystem: null.NewString((string)(AWD), true),
		Engine: Engine{
			Code:              null.NewString("PY-RPS", true),
			Type:              null.NewString("水冷直列4気筒DOHC16バルブ", true),
			Cylinders:         null.NewInt(4, true),
			CylinderLayout:    null.NewString((string)(I), true),
			ValveSystem:       null.NewString((string)(DOHC), true),
			Displacement:      null.NewFloat(2.488, true),
			Bore:              null.NewFloat(89.0, true),
			Stroke:            null.NewFloat(100.0, true),
			CompressionRatio:  null.NewFloat(13.0, true),
			MaxOutput:         null.NewFloat(138, true),
			MaxOutputLowerRpm: null.NewFloat(6000, true),
			MaxOutputUpperRpm: null.NewFloat(6000, true),
			MaxTorque:         null.NewFloat(250, true),
			MaxTorqueLowerRpm: null.NewFloat(4000, true),
			MaxTorqueUpperRpm: null.NewFloat(4000, true),
			FuelSystem:        null.NewString("DI", true),
			FuelType:          null.NewString("無鉛レギュラーガソリン", true),
			FuelTankCap:       null.NewInt(58, true),
		},
		// MotorX: Motor{},
		// MotorY: Motor{},
		// Battery: Battery{},
		Steering:        null.NewString("ラック&ピニオン式", true),
		SuspensionFront: null.NewString("マクファーソンストラット式", true),
		SuspensionRear:  null.NewString("マルチリンク式", true),
		BrakeFront:      null.NewString("ベンチレーテッドディスク", true),
		BrakeRear:       null.NewString("ディスク", true),
		TireFront: Tire{
			SectionWidth:  null.NewInt(225, true),
			AspectRatio:   null.NewInt(55, true),
			WheelDiameter: null.NewInt(19, true),
		},
		TireRear: Tire{
			SectionWidth:  null.NewInt(225, true),
			AspectRatio:   null.NewInt(55, true),
			WheelDiameter: null.NewInt(19, true),
		},
		Transmission: Transmission{
			Type:   null.NewString((string)(AT), true),
			Gears:  null.NewInt(6, true),
			Ratio1: null.NewFloat(3.552, true),
			Ratio2: null.NewFloat(2.022, true),
			Ratio3: null.NewFloat(1.452, true),
			Ratio4: null.NewFloat(1.000, true),
			Ratio5: null.NewFloat(0.708, true),
			Ratio6: null.NewFloat(0.599, true),
			//Ratio7: ,
			//Ratio8: ,
			//Ratio9: ,
			//Ratio10: ,
			RatioRear:           null.NewFloat(3.893, true),
			ReductionRatioFront: null.NewFloat(4.624, true),
			ReductionRatioRear:  null.NewFloat(2.928, true),
		},
		FuelEfficiency: null.NewString("ミラーサイクルエンジン アイドリングストップ機構 筒内直接噴射 可変バルブタイミング 気筒休止 充電制御 ロックアップ機構付トルクコンバーター 電動パワーステアリング", true),
	})
}

func seedTestCarCorollaTouring() {
	modelChangeFull, _ := time.Parse(FORMAT, "2019-Sep-17")
	modelChangeLast, _ := time.Parse(FORMAT, "2021-Nov-15")
	DB.Create(&Car{
		MakerName:       "トヨタ",
		ModelName:       "カローラツーリング",
		GradeName:       "HYBRID G-X E-Four",
		ModelCode:       "6AA-ZWE214W-AWXNB",
		Price:           null.NewInt(2678500, true),
		Url:             null.NewString("https://toyota.jp/corollatouring/", true),
		ImageUrl:        null.NewString("https://upload.wikimedia.org/wikipedia/commons/8/8a/Toyota_COROLLA_TOURING_HYBRID_W%C3%97B_2WD_%286AA-ZWE211W-AWXSB%29_front.jpg", true),
		ModelChangeFull: null.NewTime(modelChangeFull, true),
		ModelChangeLast: null.NewTime(modelChangeLast, true),
		Body: Body{
			Type:             null.NewString((string)(STATION_WAGON), true),
			Length:           null.NewInt(4495, true),
			Width:            null.NewInt(1745, true),
			Height:           null.NewInt(1460, true),
			WheelBase:        null.NewInt(2640, true),
			TreadFront:       null.NewInt(1530, true),
			TreadRear:        null.NewInt(1540, true),
			MinRoadClearance: null.NewInt(130, true),
			Weight:           null.NewInt(1410, true),
			Doors:            null.NewInt(4, true),
		},
		Interior: Interior{
			Length: null.NewInt(1790, true),
			Width:  null.NewInt(1510, true),
			Height: null.NewInt(1160, true),
			// LuggageCap: null.NewInt(0, false),
			RidingCap: null.NewInt(5, true),
		},
		Performance: Performance{
			MinTurningRadius: null.NewFloat(5.0, true),
			FcrWltc:          null.NewFloat(26.8, true),
			FcrWltcL:         null.NewFloat(25.1, true),
			FcrWltcM:         null.NewFloat(28.1, true),
			FcrWltcH:         null.NewFloat(26.8, true),
			//FcrWltcExh:       null.NewFloat(0, false),
			FcrJc08: null.NewFloat(31.0, true),
			//MpcWltc:          null.NewFloat(0, false),
			//EcrWltc:          null.NewFloat(0, false),
			//EcrWltcL:         null.NewFloat(0, false),
			//EcrWltcM:         null.NewFloat(0, false),
			//EcrWltcH:         null.NewFloat(0, false),
			//EcrWltcExh:       null.NewFloat(0, false),
			//EcrJc08:          null.NewFloat(0, false),
			//MpcJc08:          null.NewFloat(0, false),
		},
		PowerTrain:  null.NewString((string)(StrHV), true),
		DriveSystem: null.NewString((string)(AWD), true),
		Engine: Engine{
			Code:           null.NewString("2ZR-FXE", true),
			Type:           null.NewString("直列4気筒 DOHC 16バルブ VVT-i ミラーサイクル", true),
			Cylinders:      null.NewInt(4, true),
			CylinderLayout: null.NewString((string)(I), true),
			ValveSystem:    null.NewString((string)(DOHC), true),
			Displacement:   null.NewFloat(1.797, true),
			Bore:           null.NewFloat(80.5, true),
			Stroke:         null.NewFloat(88.3, true),
			//CompRatio:          null.NewFloat(0, false),
			MaxOutput:         null.NewFloat(72, true),
			MaxOutputLowerRpm: null.NewFloat(5200, true),
			MaxOutputUpperRpm: null.NewFloat(5200, true),
			MaxTorque:         null.NewFloat(142, true),
			MaxTorqueLowerRpm: null.NewFloat(3600, true),
			MaxTorqueUpperRpm: null.NewFloat(3600, true),
			FuelSystem:        null.NewString("電子制御式燃料噴射装置(EFI)", true),
			FuelType:          null.NewString("無鉛レギュラーガソリン", true),
			FuelTankCap:       null.NewInt(43, true),
		},
		MotorX: Motor{
			Code:    null.NewString("1NM", true),
			Type:    null.NewString("交流同期電動機", true),
			Purpose: null.NewString((string)(TRACTION_FRONT), true),
			//RatedOutput: ,
			MaxOutput: null.NewFloat(53, true),
			//MaxOutputLowerRpm: ,
			//MaxOutputUpperRpm: ,
			MaxTorque: null.NewFloat(163, true),
			//MaxTorqueLowerRpm: ,
			//MaxTorqueUpperRpm: ,
		},
		MotorY: Motor{
			Code:    null.NewString("1MM", true),
			Type:    null.NewString("交流同期電動機", true),
			Purpose: null.NewString((string)(TRACTION_REAR), true),
			//RatedOutput: ,
			MaxOutput: null.NewFloat(5.3, true),
			//MaxOutputLowerRpm: ,
			//MaxOutputUpperRpm: ,
			MaxTorque: null.NewFloat(55, true),
			//MaxTorqueLowerRpm: ,
			//MaxTorqueUpperRpm: ,
		},
		Battery: Battery{
			Type: null.NewString("ニッケル水素電池", true),
			//Quantity: ,
			//Voltage: ,
			Capacity: null.NewFloat(6.5, true),
		},
		//Steering:        null.NewString( "ラック&ピニオン式", true),
		SuspensionFront: null.NewString("マクファーソンストラット式コイルスプリング", true),
		SuspensionRear:  null.NewString("ダブルウィッシュボーン式コイルスプリング", true),
		BrakeFront:      null.NewString("ベンチレーテッドディスク", true),
		BrakeRear:       null.NewString("ディスク", true),
		TireFront: Tire{
			SectionWidth:  null.NewInt(195, true),
			AspectRatio:   null.NewInt(65, true),
			WheelDiameter: null.NewInt(15, true),
		},
		TireRear: Tire{
			SectionWidth:  null.NewInt(195, true),
			AspectRatio:   null.NewInt(65, true),
			WheelDiameter: null.NewInt(15, true),
		},
		Transmission: Transmission{
			Type: null.NewString((string)(PG), true),
			//Gears: null.NewInt(6, true),
			//Ratio1: null.NewFloat(3.552, true),
			//Ratio2: null.NewFloat(2.022, true),
			//Ratio3: null.NewFloat(1.452, true),
			//Ratio4: null.NewFloat(1.000, true),
			//Ratio5: null.NewFloat(0.708, true),
			//Ratio6: null.NewFloat(0.599, true),
			//Ratio7: ,
			//Ratio8: ,
			//Ratio9: ,
			//Ratio10: ,
			//RatioRear:           null.NewFloat(3.893, true),
			ReductionRatioFront: null.NewFloat(2.834, true),
			ReductionRatioRear:  null.NewFloat(10.487, true),
		},
		FuelEfficiency: null.NewString("ハイブリッドシステム アイドリングストップ装置 可変バルブタイミング 電動パワーステアリング 充電制御 電気式無段変速機", true),
	})
}

func seedTestCarNsx() {
	modelChangeFull, _ := time.Parse(FORMAT, "2017-Feb-27")
	modelChangeLast, _ := time.Parse(FORMAT, "2021-Aug-30")
	DB.Create(&Car{
		MakerName:       "ホンダ",
		ModelName:       "NSX",
		GradeName:       "Type S",
		ModelCode:       "5AA-NC1",
		Price:           null.NewInt(27940000, true),
		Url:             null.NewString("https://www.honda.co.jp/NSX/types/", true),
		ImageUrl:        null.NewString("https://upload.wikimedia.org/wikipedia/commons/e/ea/2019_Honda_NSX_3.5_CAA-NC1_%2820190722%29_01.jpg", true),
		ModelChangeFull: null.NewTime(modelChangeFull, true),
		ModelChangeLast: null.NewTime(modelChangeLast, true),
		Body: Body{
			Type:             null.NewString((string)(COUPE), true),
			Length:           null.NewInt(4535, true),
			Width:            null.NewInt(1940, true),
			Height:           null.NewInt(1215, true),
			WheelBase:        null.NewInt(2630, true),
			TreadFront:       null.NewInt(1665, true),
			TreadRear:        null.NewInt(1635, true),
			MinRoadClearance: null.NewInt(110, true),
			Weight:           null.NewInt(1790, true),
			Doors:            null.NewInt(2, true),
		},
		Interior: Interior{
			// Length: null.NewInt(1790, true),
			// Width:  null.NewInt(1510, true),
			// Height: null.NewInt(1160, true),
			// LuggageCap: null.NewInt(0, false),
			RidingCap: null.NewInt(2, true),
		},
		Performance: Performance{
			MinTurningRadius: null.NewFloat(5.9, true),
			FcrWltc:          null.NewFloat(10.6, true),
			FcrWltcL:         null.NewFloat(7.8, true),
			FcrWltcM:         null.NewFloat(12.1, true),
			FcrWltcH:         null.NewFloat(11.4, true),
			//FcrWltcExh:       null.NewFloat(0, false),
			// FcrJc08: null.NewFloat(31.0, true),
			//MpcWltc:          null.NewFloat(0, false),
			//EcrWltc:          null.NewFloat(0, false),
			//EcrWltcL:         null.NewFloat(0, false),
			//EcrWltcM:         null.NewFloat(0, false),
			//EcrWltcH:         null.NewFloat(0, false),
			//EcrWltcExh:       null.NewFloat(0, false),
			//EcrJc08:          null.NewFloat(0, false),
			//MpcJc08:          null.NewFloat(0, false),
		},
		PowerTrain:  null.NewString((string)(MldHV), true),
		DriveSystem: null.NewString((string)(AWD), true),
		Engine: Engine{
			Code:              null.NewString("JNC", true),
			Type:              null.NewString("水冷V型6気筒縦置", true),
			Cylinders:         null.NewInt(6, true),
			CylinderLayout:    null.NewString((string)(V), true),
			ValveSystem:       null.NewString((string)(DOHC), true),
			Displacement:      null.NewFloat(3.492, true),
			Bore:              null.NewFloat(91.0, true),
			Stroke:            null.NewFloat(89.5, true),
			CompressionRatio:  null.NewFloat(10.0, false),
			MaxOutput:         null.NewFloat(389, true),
			MaxOutputLowerRpm: null.NewFloat(6500, true),
			MaxOutputUpperRpm: null.NewFloat(6850, true),
			MaxTorque:         null.NewFloat(600, true),
			MaxTorqueLowerRpm: null.NewFloat(2300, true),
			MaxTorqueUpperRpm: null.NewFloat(6000, true),
			FuelSystem:        null.NewString("電子制御燃料噴射式(ホンダ PGM-FI)", true),
			FuelType:          null.NewString("無鉛プレミアムガソリン", true),
			FuelTankCap:       null.NewInt(59, true),
		},
		MotorX: Motor{
			Code:    null.NewString("H3", true),
			Type:    null.NewString("交流同期電動機", true),
			Purpose: null.NewString((string)(TRACTION_FRONT), true),
			//RatedOutput: ,
			MaxOutput:         null.NewFloat(27, true),
			MaxOutputLowerRpm: null.NewFloat(4000, true),
			MaxOutputUpperRpm: null.NewFloat(4000, true),
			MaxTorque:         null.NewFloat(73, true),
			MaxTorqueLowerRpm: null.NewFloat(0, true),
			MaxTorqueUpperRpm: null.NewFloat(2000, true),
		},
		MotorY: Motor{
			Code:    null.NewString("H2", true),
			Type:    null.NewString("交流同期電動機", true),
			Purpose: null.NewString((string)(TRACTION_REAR), true),
			//RatedOutput: ,
			MaxOutput:         null.NewFloat(35, true),
			MaxOutputLowerRpm: null.NewFloat(3000, true),
			MaxOutputUpperRpm: null.NewFloat(3000, true),
			MaxTorque:         null.NewFloat(148, true),
			MaxTorqueLowerRpm: null.NewFloat(500, true),
			MaxTorqueUpperRpm: null.NewFloat(2000, true),
		},
		Battery: Battery{
			Type:     null.NewString("ニッケル水素電池", true),
			Quantity: null.NewInt(72, true),
			//Voltage: ,
			// Capacity: null.NewFloat(6.5, true),
		},
		Steering:        null.NewString("ラック&ピニオン式(電動パワーステアリング仕様)", true),
		SuspensionFront: null.NewString("ダブルウィッシュボーン式", true),
		SuspensionRear:  null.NewString("ウィッシュボーン式", true),
		BrakeFront:      null.NewString("油圧式ベンチレーテッドディスク", true),
		BrakeRear:       null.NewString("油圧式ベンチレーテッドディスク", true),
		TireFront: Tire{
			SectionWidth:  null.NewInt(245, true),
			AspectRatio:   null.NewInt(35, true),
			WheelDiameter: null.NewInt(19, true),
		},
		TireRear: Tire{
			SectionWidth:  null.NewInt(305, true),
			AspectRatio:   null.NewInt(30, true),
			WheelDiameter: null.NewInt(20, true),
		},
		Transmission: Transmission{
			Type:   null.NewString((string)(DCT), true),
			Gears:  null.NewInt(9, true),
			Ratio1: null.NewFloat(3.838, true),
			Ratio2: null.NewFloat(2.433, true),
			Ratio3: null.NewFloat(1.777, true),
			Ratio4: null.NewFloat(1.427, true),
			Ratio5: null.NewFloat(1.211, true),
			Ratio6: null.NewFloat(1.038, true),
			Ratio7: null.NewFloat(0.880, true),
			Ratio8: null.NewFloat(0.747, true),
			Ratio9: null.NewFloat(0.633, true),
			//Ratio10: ,
			RatioRear:           null.NewFloat(2.394, true),
			ReductionRatioFront: null.NewFloat(10.382, true),
			ReductionRatioRear:  null.NewFloat(3.583, true),
		},
		FuelEfficiency: null.NewString("ハイブリッドシステム 直噴エンジン 可変バルブタイミング アイドリングストップ装置 電動パワーステアリング", true),
	})
}

func seedTestCarHondaE() {
	modelChangeFull, _ := time.Parse(FORMAT, "2020-Aug-27")
	modelChangeLast, _ := time.Parse(FORMAT, "2020-Aug-27")
	DB.Create(&Car{
		MakerName:       "ホンダ",
		ModelName:       "Honda e",
		GradeName:       "Honda e Advance",
		ModelCode:       "ZAA-ZC7",
		Price:           null.NewInt(4950000, true),
		Url:             null.NewString("https://www.honda.co.jp/honda-e/", true),
		ImageUrl:        null.NewString("https://upload.wikimedia.org/wikipedia/commons/9/9e/Honda_e_Advance_%28ZAA-ZC7%29_front.jpg", true),
		ModelChangeFull: null.NewTime(modelChangeFull, true),
		ModelChangeLast: null.NewTime(modelChangeLast, true),
		Body: Body{
			Type:             null.NewString((string)(HATCHBACK), true),
			Length:           null.NewInt(3895, true),
			Width:            null.NewInt(1750, true),
			Height:           null.NewInt(1510, true),
			WheelBase:        null.NewInt(2530, true),
			TreadFront:       null.NewInt(1510, true),
			TreadRear:        null.NewInt(1505, true),
			MinRoadClearance: null.NewInt(145, true),
			Weight:           null.NewInt(1540, true),
			Doors:            null.NewInt(4, true),
		},
		Interior: Interior{
			Length: null.NewInt(1845, true),
			Width:  null.NewInt(1385, true),
			Height: null.NewInt(1120, true),
			// LuggageCap: null.NewInt(0, false),
			RidingCap: null.NewInt(4, true),
		},
		Performance: Performance{
			MinTurningRadius: null.NewFloat(4.3, true),
			//FcrWltc:          null.NewFloat(26.8, true),
			//FcrWltcL:         null.NewFloat(25.1, true),
			//FcrWltcM:         null.NewFloat(28.1, true),
			//FcrWltcH:         null.NewFloat(26.8, true),
			//FcrWltcExh:       null.NewFloat(0, false),
			//FcrJc08: null.NewFloat(31.0, true),
			MpcWltc:  null.NewFloat(259, false),
			EcrWltc:  null.NewFloat(138, false),
			EcrWltcL: null.NewFloat(116, false),
			EcrWltcM: null.NewFloat(130, false),
			EcrWltcH: null.NewFloat(149, false),
			//EcrWltcExh:       null.NewFloat(0, false),
			EcrJc08: null.NewFloat(135, false),
			MpcJc08: null.NewFloat(274, false),
		},
		PowerTrain:  null.NewString((string)(BEV), true),
		DriveSystem: null.NewString((string)(RR), true),
		// Engine: Engine{
		// 	Code:           null.NewString( "2ZR-FXE", true),
		// 	Type:           null.NewString( "直列4気筒 DOHC 16バルブ VVT-i ミラーサイクル", true),
		// 	Cylinders:      null.NewInt(4, true),
		// 	CylinderLayout: null.NewString( (string)(I), true),
		// 	ValveSystem:    null.NewString( (string)(DOHC), true),
		// 	Displacement:   null.NewFloat(1.797, true),
		// 	Bore:           null.NewFloat(80.5, true),
		// 	Stroke:         null.NewFloat(88.3, true),
		// 	CompRatio:          null.NewFloat(0, false),
		// 	MaxOutput:          null.NewFloat(72, true),
		// 	MaxOutputLowerRpm:  null.NewFloat(5200, true),
		// 	MaxOutputUpperRpm: null.NewFloat(5200, true),
		// 	MaxTorque:          null.NewFloat(142, true),
		// 	MaxTorqueLowerRpm:  null.NewFloat(3600, true),
		// 	MaxTorqueUpperRpm: null.NewFloat(3600, true),
		// 	FuelSystem:         null.NewString( "電子制御式燃料噴射装置(EFI)", true),
		// 	FuelType:           null.NewString( "無鉛レギュラーガソリン", true),
		// 	FuelTankCap:        null.NewInt(43, true),
		// },
		MotorX: Motor{
			Code:              null.NewString("MCF5", true),
			Type:              null.NewString("交流同期電動機", true),
			Purpose:           null.NewString((string)(TRACTION_REAR), true),
			RatedOutput:       null.NewFloat(60, true),
			MaxOutput:         null.NewFloat(113, true),
			MaxOutputLowerRpm: null.NewFloat(3497, true),
			MaxOutputUpperRpm: null.NewFloat(10000, true),
			MaxTorque:         null.NewFloat(315, true),
			MaxTorqueLowerRpm: null.NewFloat(0, true),
			MaxTorqueUpperRpm: null.NewFloat(2000, true),
		},
		// MotorY: Motor{
		// 	Code:    null.NewString( "1MM", true),
		// 	Type:    null.NewString( "交流同期電動機", true),
		// 	Purpose: null.NewString( (string)(TRACTION_REAR), true),
		// 	RatedOutput: ,
		// 	MaxOutput: null.NewFloat(5.3, true),
		// 	MaxOutputLowerRpm: ,
		// 	MaxOutputUpperRpm: ,
		// 	MaxTorque: null.NewFloat(55, true),
		// 	MaxTorqueLowerRpm: ,
		// 	MaxTorqueUpperRpm: ,
		// },
		Battery: Battery{
			Type:     null.NewString("リチウムイオン電池", true),
			Quantity: null.NewInt(193, true),
			Voltage:  null.NewFloat(3.7, true),
			Capacity: null.NewFloat(50.0, true),
		},
		Steering:        null.NewString("ラック&ピニオン式", true),
		SuspensionFront: null.NewString("マクファーソン式", true),
		SuspensionRear:  null.NewString("マクファーソン式", true),
		BrakeFront:      null.NewString("油圧式ベンチレーテッドディスク", true),
		BrakeRear:       null.NewString("油圧式ディスク", true),
		TireFront: Tire{
			SectionWidth:  null.NewInt(205, true),
			AspectRatio:   null.NewInt(45, true),
			WheelDiameter: null.NewInt(17, true),
		},
		TireRear: Tire{
			SectionWidth:  null.NewInt(225, true),
			AspectRatio:   null.NewInt(45, true),
			WheelDiameter: null.NewInt(17, true),
		},
		Transmission: Transmission{
			// 	Type: null.NewString( (string)(PG), true),
			// 	Gears: null.NewInt(6, true),
			// 	Ratio1: null.NewFloat(3.552, true),
			// 	Ratio2: null.NewFloat(2.022, true),
			// 	Ratio3: null.NewFloat(1.452, true),
			// 	Ratio4: null.NewFloat(1.000, true),
			// 	Ratio5: null.NewFloat(0.708, true),
			// 	Ratio6: null.NewFloat(0.599, true),
			// 	Ratio7: ,
			// 	Ratio8: ,
			// 	Ratio9: ,
			// 	Ratio10: ,
			// 	RatioRear:           null.NewFloat(3.893, true),
			ReductionRatioFront: null.NewFloat(9.545, true),
			// 	ReductionRatioRear:  null.NewFloat(10.487, true),
		},
		FuelEfficiency: null.NewString("電動パワーステアリング", true),
	})

}

func seedTestCarNote() {
	modelChangeFull, _ := time.Parse(FORMAT, "2020-Nov-24")
	modelChangeLast, _ := time.Parse(FORMAT, "2021-Nov-04")
	DB.Create(&Car{
		MakerName:       "日産",
		ModelName:       "ノート",
		GradeName:       "X FOUR",
		ModelCode:       "6AA-SNE13",
		Price:           null.NewInt(2445300, true),
		Url:             null.NewString("https://www3.nissan.co.jp/vehicles/new/note.html", true),
		ImageUrl:        null.NewString("https://upload.wikimedia.org/wikipedia/commons/0/0a/Nissan_Note_e-POWER_%28E13%29%2C_2021%2C_front-left.jpg", true),
		ModelChangeFull: null.NewTime(modelChangeFull, true),
		ModelChangeLast: null.NewTime(modelChangeLast, true),
		Body: Body{
			Type:             null.NewString((string)(HATCHBACK), true),
			Length:           null.NewInt(4045, true),
			Width:            null.NewInt(1695, true),
			Height:           null.NewInt(1520, true),
			WheelBase:        null.NewInt(2580, true),
			TreadFront:       null.NewInt(1490, true),
			TreadRear:        null.NewInt(1490, true),
			MinRoadClearance: null.NewInt(125, true),
			Weight:           null.NewInt(1340, true),
			Doors:            null.NewInt(4, true),
		},
		Interior: Interior{
			Length: null.NewInt(2030, true),
			Width:  null.NewInt(1445, true),
			Height: null.NewInt(1240, true),
			// LuggageCap: null.NewInt(0, false),
			RidingCap: null.NewInt(5, true),
		},
		Performance: Performance{
			MinTurningRadius: null.NewFloat(4.9, true),
			FcrWltc:          null.NewFloat(23.8, true),
			FcrWltcL:         null.NewFloat(23.1, true),
			FcrWltcM:         null.NewFloat(25.8, true),
			FcrWltcH:         null.NewFloat(22.9, true),
			//FcrWltcExh:       null.NewFloat(0, false),
			FcrJc08: null.NewFloat(28.2, true),
			//MpcWltc:          null.NewFloat(0, false),
			//EcrWltc:          null.NewFloat(0, false),
			//EcrWltcL:         null.NewFloat(0, false),
			//EcrWltcM:         null.NewFloat(0, false),
			//EcrWltcH:         null.NewFloat(0, false),
			//EcrWltcExh:       null.NewFloat(0, false),
			//EcrJc08:          null.NewFloat(0, false),
			//MpcJc08:          null.NewFloat(0, false),
		},
		PowerTrain:  null.NewString((string)(SerHV), true),
		DriveSystem: null.NewString((string)(AWD), true),
		Engine: Engine{
			Code:              null.NewString("HR12DE", true),
			Type:              null.NewString("DOHC水冷直列3気筒", true),
			Cylinders:         null.NewInt(3, true),
			CylinderLayout:    null.NewString((string)(I), true),
			ValveSystem:       null.NewString((string)(DOHC), true),
			Displacement:      null.NewFloat(1.198, true),
			Bore:              null.NewFloat(78.0, true),
			Stroke:            null.NewFloat(83.6, true),
			CompressionRatio:  null.NewFloat(12.0, false),
			MaxOutput:         null.NewFloat(60, true),
			MaxOutputLowerRpm: null.NewFloat(6000, true),
			MaxOutputUpperRpm: null.NewFloat(6000, true),
			MaxTorque:         null.NewFloat(103, true),
			MaxTorqueLowerRpm: null.NewFloat(4800, true),
			MaxTorqueUpperRpm: null.NewFloat(4800, true),
			FuelSystem:        null.NewString("ニッサンEGI(ECCS)電子制御燃料噴射装置", true),
			FuelType:          null.NewString("無鉛レギュラーガソリン", true),
			FuelTankCap:       null.NewInt(36, true),
		},
		MotorX: Motor{
			Code:    null.NewString("EM47", true),
			Type:    null.NewString("交流同期電動機", true),
			Purpose: null.NewString((string)(GENERATOR), true),
			//RatedOutput: ,
			MaxOutput:         null.NewFloat(85, true),
			MaxOutputLowerRpm: null.NewFloat(2900, true),
			MaxOutputUpperRpm: null.NewFloat(10341, true),
			MaxTorque:         null.NewFloat(280, true),
			MaxTorqueLowerRpm: null.NewFloat(0, true),
			MaxTorqueUpperRpm: null.NewFloat(2900, true),
		},
		MotorY: Motor{
			Code:    null.NewString("MM48", true),
			Type:    null.NewString("交流同期電動機", true),
			Purpose: null.NewString((string)(TRACTION_REAR), true),
			// RatedOutput: ,
			MaxOutput:         null.NewFloat(50, true),
			MaxOutputLowerRpm: null.NewFloat(4775, true),
			MaxOutputUpperRpm: null.NewFloat(10024, true),
			MaxTorque:         null.NewFloat(100, true),
			MaxTorqueLowerRpm: null.NewFloat(0, true),
			MaxTorqueUpperRpm: null.NewFloat(4775, true),
		},
		Battery: Battery{
			Type: null.NewString("リチウムイオン電池", true),
			//Quantity: ,
			//Voltage: ,
			// Capacity: null.NewFloat(6.5, true),
		},
		Steering:        null.NewString("ラック&ピニオン式", true),
		SuspensionFront: null.NewString("独立懸架ストラット式", true),
		SuspensionRear:  null.NewString("トーションビーム式", true),
		BrakeFront:      null.NewString("ベンチレーテッドディスク式", true),
		BrakeRear:       null.NewString("リーディングトレーリング式", true),
		TireFront: Tire{
			SectionWidth:  null.NewInt(185, true),
			AspectRatio:   null.NewInt(60, true),
			WheelDiameter: null.NewInt(16, true),
		},
		TireRear: Tire{
			SectionWidth:  null.NewInt(185, true),
			AspectRatio:   null.NewInt(60, true),
			WheelDiameter: null.NewInt(16, true),
		},
		Transmission: Transmission{
			// Type: null.NewString( (string)(PG), true),
			//Gears: null.NewInt(6, true),
			//Ratio1: null.NewFloat(3.552, true),
			//Ratio2: null.NewFloat(2.022, true),
			//Ratio3: null.NewFloat(1.452, true),
			//Ratio4: null.NewFloat(1.000, true),
			//Ratio5: null.NewFloat(0.708, true),
			//Ratio6: null.NewFloat(0.599, true),
			//Ratio7: ,
			//Ratio8: ,
			//Ratio9: ,
			//Ratio10: ,
			//RatioRear:           null.NewFloat(3.893, true),
			ReductionRatioFront: null.NewFloat(7.388, true),
			ReductionRatioRear:  null.NewFloat(7.282, true),
		},
		FuelEfficiency: null.NewString("ハイブリッドシステム アイドリングストップ装置 可変バルブタイミング ミラーサイクル 電動パワーステアリング", true),
	})
}

func seedTestCarThree() {
	modelChangeFull, _ := time.Parse(FORMAT, "2019-Sep-26")
	modelChangeLast, _ := time.Parse(FORMAT, "2019-Sep-26")
	DB.Create(&Car{
		MakerName: "BMW",
		ModelName: "3シリーズツーリング",
		GradeName: "320d xDriveツーリング Standard",
		ModelCode: "3DA-6L20",
		Price:     null.NewInt(6340000, true),
		Url:       null.NewString("https://www.bmw.co.jp/ja/all-models/3-series/touring/2019/bmw-3-series-touring-inspire.html", true),
		// ImageUrl:        null.NewString( "", true),
		ModelChangeFull: null.NewTime(modelChangeFull, true),
		ModelChangeLast: null.NewTime(modelChangeLast, true),
		Body: Body{
			Type:             null.NewString((string)(STATION_WAGON), true),
			Length:           null.NewInt(4715, true),
			Width:            null.NewInt(1825, true),
			Height:           null.NewInt(1475, true),
			WheelBase:        null.NewInt(2850, true),
			TreadFront:       null.NewInt(1575, true),
			TreadRear:        null.NewInt(1590, true),
			MinRoadClearance: null.NewInt(135, true),
			Weight:           null.NewInt(1730, true),
			Doors:            null.NewInt(4, true),
		},
		Interior: Interior{
			// Length: null.NewInt(1890, true),
			// Width:  null.NewInt(1540, true),
			// Height: null.NewInt(1265, true),
			LuggageCap: null.NewInt(500, false),
			RidingCap:  null.NewInt(5, true),
		},
		Performance: Performance{
			MinTurningRadius: null.NewFloat(5.7, true),
			FcrWltc:          null.NewFloat(15.6, true),
			FcrWltcL:         null.NewFloat(12.6, true),
			FcrWltcM:         null.NewFloat(14.9, true),
			FcrWltcH:         null.NewFloat(18.0, true),
			//FcrWltcExh:       null.NewFloat(0, false),
			FcrJc08: null.NewFloat(19.6, true),
			//MpcWltc:          null.NewFloat(0, false),
			//EcrWltc:          null.NewFloat(0, false),
			//EcrWltcL:         null.NewFloat(0, false),
			//EcrWltcM:         null.NewFloat(0, false),
			//EcrWltcH:         null.NewFloat(0, false),
			//EcrWltcExh:       null.NewFloat(0, false),
			//EcrJc08:          null.NewFloat(0, false),
			//MpcJc08:          null.NewFloat(0, false),
		},
		PowerTrain:  null.NewString((string)(ICE), true),
		DriveSystem: null.NewString((string)(AWD), true),
		Engine: Engine{
			Code:           null.NewString("B47D20B", true),
			Type:           null.NewString("直列4気筒DOHCディーゼル", true),
			Cylinders:      null.NewInt(4, true),
			CylinderLayout: null.NewString((string)(I), true),
			ValveSystem:    null.NewString((string)(DOHC), true),
			Displacement:   null.NewFloat(1.995, true),
			// Bore:               null.NewFloat(140, true),
			// Stroke:             null.NewFloat(100.0, true),
			// CompRatio:          null.NewFloat(13.0, true),
			MaxOutput:         null.NewFloat(140, true),
			MaxOutputLowerRpm: null.NewFloat(4000, true),
			MaxOutputUpperRpm: null.NewFloat(4000, true),
			MaxTorque:         null.NewFloat(400, true),
			MaxTorqueLowerRpm: null.NewFloat(1750, true),
			MaxTorqueUpperRpm: null.NewFloat(2500, true),
			FuelSystem:        null.NewString("デジタル・ディーゼル・エレクトロニクス(DDE/電子燃料噴射装置)", true),
			FuelType:          null.NewString("軽油", true),
			FuelTankCap:       null.NewInt(59, true),
		},
		// MotorX: Motor{},
		// MotorY: Motor{},
		// Battery: Battery{},
		Steering:        null.NewString("ラック&ピニオン式、単速感応式パワー・ステアリング", true),
		SuspensionFront: null.NewString("ダブル・ジョイント・スプリング・ストラット式、コイルスプリング", true),
		SuspensionRear:  null.NewString("5リンク式、コイル・スプリング", true),
		BrakeFront:      null.NewString("ベンチレーテッドディスク", true),
		BrakeRear:       null.NewString("ベンチレーテッドディスク", true),
		TireFront: Tire{
			SectionWidth:  null.NewInt(225, true),
			AspectRatio:   null.NewInt(50, true),
			WheelDiameter: null.NewInt(17, true),
		},
		TireRear: Tire{
			SectionWidth:  null.NewInt(225, true),
			AspectRatio:   null.NewInt(50, true),
			WheelDiameter: null.NewInt(17, true),
		},
		Transmission: Transmission{
			Type:   null.NewString((string)(AT), true),
			Gears:  null.NewInt(8, true),
			Ratio1: null.NewFloat(5.250, true),
			Ratio2: null.NewFloat(3.360, true),
			Ratio3: null.NewFloat(2.172, true),
			Ratio4: null.NewFloat(1.720, true),
			Ratio5: null.NewFloat(1.316, true),
			Ratio6: null.NewFloat(1.000, true),
			Ratio7: null.NewFloat(0.822, true),
			Ratio8: null.NewFloat(0.640, true),
			//Ratio9: ,
			//Ratio10: ,
			RatioRear:           null.NewFloat(3.712, true),
			ReductionRatioFront: null.NewFloat(2.813, true),
			// ReductionRatioRear:  null.NewFloat(2.813, true),
		},
		FuelEfficiency: null.NewString("筒内直接噴射 電子制御式燃料噴射 高圧噴射(コモンレール・ダイレクト・インジェクション・システム) 過給機(可変ジオメトリー・ターボチャージャー) 充電制御(ブレーキ・エネルギー回生システム) アイドリング・ストップ装置(エンジン・オート・スタート/ストップ) 電動パワーステアリング", true),
	})
}

func TestSearchCarsQueryEmpty(t *testing.T) {
	DB.Exec("TRUNCATE TABLE cars")
	seedTestCarCx5()
	seedTestCarCorollaTouring()
	seedTestCarHondaE()
	seedTestCarNote()
	seedTestCarThree()
	seedTestCarNsx()
	token := login("user")
	// HTTPリクエストの生成
	body := `{}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/cars/search", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestSearchCarsQueryMakerName(t *testing.T) {
	DB.Exec("TRUNCATE TABLE cars")
	seedTestCarCx5()
	seedTestCarCorollaTouring()
	seedTestCarHondaE()
	seedTestCarNote()
	seedTestCarThree()
	seedTestCarNsx()
	token := login("user")
	// HTTPリクエストの生成
	body := `{"maker_name": "マツダ"}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/cars/search", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
	var cars []Car
	json.Unmarshal(recorder.Body.Bytes(), &cars)
	assert.Equal(t, 1, len(cars))
}

func TestSearchCarsQueryModelName(t *testing.T) {
	DB.Exec("TRUNCATE TABLE cars")
	seedTestCarCx5()
	seedTestCarCorollaTouring()
	seedTestCarHondaE()
	seedTestCarNote()
	seedTestCarThree()
	seedTestCarNsx()
	token := login("user")
	// HTTPリクエストの生成
	body := `{"model_name": "CX-5"}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/cars/search", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
	var cars []Car
	json.Unmarshal(recorder.Body.Bytes(), &cars)
	assert.Equal(t, 1, len(cars))
}

func TestSearchCarsQueryGradeName(t *testing.T) {
	DB.Exec("TRUNCATE TABLE cars")
	seedTestCarCx5()
	seedTestCarCorollaTouring()
	seedTestCarHondaE()
	seedTestCarNote()
	seedTestCarThree()
	seedTestCarNsx()
	token := login("user")
	// HTTPリクエストの生成
	body := `{"grade_name": "25S Proactive"}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/cars/search", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
	var cars []Car
	json.Unmarshal(recorder.Body.Bytes(), &cars)
	assert.Equal(t, 1, len(cars))
}

func TestSearchCarsQueryModelCode(t *testing.T) {
	DB.Exec("TRUNCATE TABLE cars")
	seedTestCarCx5()
	seedTestCarCorollaTouring()
	seedTestCarHondaE()
	seedTestCarNote()
	seedTestCarThree()
	seedTestCarNsx()
	token := login("user")
	// HTTPリクエストの生成
	body := `{"model_code": "6BA-KF5P"}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/cars/search", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
	var cars []Car
	json.Unmarshal(recorder.Body.Bytes(), &cars)
	assert.Equal(t, 1, len(cars))
}

func TestSearchCarsQueryPriceLower(t *testing.T) {
	DB.Exec("TRUNCATE TABLE cars")
	seedTestCarCx5()
	seedTestCarCorollaTouring()
	seedTestCarHondaE()
	seedTestCarNote()
	seedTestCarThree()
	seedTestCarNsx()
	token := login("user")
	// HTTPリクエストの生成
	body := `{"price_lower": 10000000}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/cars/search", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
	var cars []Car
	json.Unmarshal(recorder.Body.Bytes(), &cars)
	assert.Equal(t, 1, len(cars))
}

func TestSearchCarsQueryPriceUpper(t *testing.T) {
	DB.Exec("TRUNCATE TABLE cars")
	seedTestCarCx5()
	seedTestCarCorollaTouring()
	seedTestCarHondaE()
	seedTestCarNote()
	seedTestCarThree()
	seedTestCarNsx()
	token := login("user")
	// HTTPリクエストの生成
	body := `{"price_upper": 2500000}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/cars/search", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
	var cars []Car
	json.Unmarshal(recorder.Body.Bytes(), &cars)
	assert.Equal(t, 1, len(cars))
}

func TestSearchCarsQueryModelChangeFrom(t *testing.T) {
	DB.Exec("TRUNCATE TABLE cars")
	seedTestCarCx5()
	seedTestCarCorollaTouring()
	seedTestCarHondaE()
	seedTestCarNote()
	seedTestCarThree()
	seedTestCarNsx()
	token := login("user")
	// HTTPリクエストの生成
	body := `{"model_change_from": "2021-11-14"}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/cars/search", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
	var cars []Car
	json.Unmarshal(recorder.Body.Bytes(), &cars)
	assert.Equal(t, 1, len(cars))
}

func TestSearchCarsQueryModelChangeTo(t *testing.T) {
	DB.Exec("TRUNCATE TABLE cars")
	seedTestCarCx5()
	seedTestCarCorollaTouring()
	seedTestCarHondaE()
	seedTestCarNote()
	seedTestCarThree()
	seedTestCarNsx()
	token := login("user")
	// HTTPリクエストの生成
	body := `{"model_change_to": "2016-12-16"}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/cars/search", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
	var cars []Car
	json.Unmarshal(recorder.Body.Bytes(), &cars)
	assert.Equal(t, 1, len(cars))
}

func TestSearchCarsQueryPowerTrain(t *testing.T) {
	DB.Exec("TRUNCATE TABLE cars")
	seedTestCarCx5()
	seedTestCarCorollaTouring()
	seedTestCarHondaE()
	seedTestCarNote()
	seedTestCarThree()
	seedTestCarNsx()
	token := login("user")
	// HTTPリクエストの生成
	body := `{"power_train": ["BEV"]}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/cars/search", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
	var cars []Car
	json.Unmarshal(recorder.Body.Bytes(), &cars)
	assert.Equal(t, 1, len(cars))
}

func TestSearchCarsQueryBodyType(t *testing.T) {
	DB.Exec("TRUNCATE TABLE cars")
	seedTestCarCx5()
	seedTestCarCorollaTouring()
	seedTestCarHondaE()
	seedTestCarNote()
	seedTestCarThree()
	seedTestCarNsx()
	token := login("user")
	// HTTPリクエストの生成
	body := `{"body_type": ["SUV"]}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/cars/search", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
	var cars []Car
	json.Unmarshal(recorder.Body.Bytes(), &cars)
	assert.Equal(t, 1, len(cars))
}

func TestGetCar(t *testing.T) {
	DB.Exec("TRUNCATE TABLE cars")
	seedTestCarCx5()
	token := login("user")
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/cars/1", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestPatchCarSuccessAllColumnUser(t *testing.T) {
	seedTestCarCx5()       // テストデータの準備
	token := login("user") // 認証実行
	// HTTPリクエストの生成
	body := `{
		"maker_name": "マツダ",
		"model_name": "CX-5",
		"grade_name": "25S Proactive",
		"model_code": "6BA-KF5P",
		"price": 3140500,
		"url": "https://www.mazda.co.jp/cars/cx-5/",
		"image_url": "https://upload.wikimedia.org/wikipedia/commons/8/85/2017_Mazda_CX-5_%28KF%29_Maxx_2WD_wagon_%282018-11-02%29_01.jpg",
		"model_change_full": "2016-12-01T00:00:00+09:00",
		"model_change_last": "2018-01-01T00:00:00+09:00",
		"body": {
			"type": "SUV",
			"length": 4545,
			"width": 1840,
			"height": 1690,
			"wheel_base": 2700,
			"tread_front": 1595,
			"tread_rear": 1595,
			"min_road_clearance": 210,
			"body_weight": 1620,
			"doors": 4
		},
		"interior": {
			"length": 1890,
			"width": 1540,
			"height": 1265,
			"luggage_cap": null,
			"riding_cap": 5
		},
		"performance": {
			"min_turning_radius": 5.5,
			"fcr_wltc": 13,
			"fcr_wltc_l": 10.2,
			"fcr_wltc_m": 13.4,
			"fcr_wltc_h": 14.7,
			"fcr_wltc_exh": null,
			"fcr_jc08": 14.2,
			"mpc_wltc": null,
			"ecr_wltc": null,
			"ecr_wltc_l": null,
			"ecr_wltc_m": null,
			"ecr_wltc_h": null,
			"ecr_wltc_exh": null,
			"ecr_jc08": null,
			"mpc_jc08": null
		},
		"power_train": "ICE",
		"drive_system": "AWD",
		"engine": {
			"code": "PY-RPS",
			"type": "水冷直列4気筒DOHC16バルブ",
			"cylinders": 4,
			"cylinder_layout": "I",
			"valve_system": "DOHC",
			"displacement": 2.488,
			"bore": 89,
			"stroke": 100,
			"compression_ratio": 13,
			"max_output": 138,
			"max_output_lower_rpm": 6000,
			"max_output_upper_rpm": 6000,
			"max_torque": 250,
			"max_torque_lower_rpm": 4000,
			"max_torque_upper_rpm": 4000,
			"fuel_system": "DI",
			"fuel_type": "無鉛レギュラーガソリン",
			"fuel_tank_cap": 58
		},
		"motor_x": {
			"code": null,
			"type": null,
			"purpose": null,
			"rated_output": null,
			"max_output": null,
			"max_output_lower_rpm": null,
			"max_output_upper_rpm": null,
			"max_torque": null,
			"max_torque_lower_rpm": null,
			"max_torque_upper_rpm": null
		},
		"motor_y": {
			"code": null,
			"type": null,
			"purpose": null,
			"rated_output": null,
			"max_output": null,
			"max_output_lower_rpm": null,
			"max_output_upper_rpm": null,
			"max_torque": null,
			"max_torque_lower_rpm": null,
			"max_torque_upper_rpm": null
		},
		"battery": {
			"type": null,
			"quantity": null,
			"voltage": null,
			"capacity": null
		},
		"steering": "ラック&ピニオン式",
		"suspension_front": "マクファーソンストラット式",
		"suspension_rear": "マルチリンク式",
		"brake_front": "ベンチレーテッドディスク",
		"brake_rear": "ディスク",
		"tire_front": {
			"section_width": 225,
			"aspect_ratio": 55,
			"wheel_diameter": 19
		},
		"tire_rear": {
			"section_width": 225,
			"aspect_ratio": 55,
			"wheel_diameter": 19
		},
		"transmission": {
			"type": "AT",
			"gears": 6,
			"ratio_1": 3.552,
			"ratio_2": 2.022,
			"ratio_3": 1.452,
			"ratio_4": 1,
			"ratio_5": 0.708,
			"ratio_6": 0.599,
			"ratio_7": null,
			"ratio_8": null,
			"ratio_9": null,
			"ratio_10": null,
			"ratio_rear": 3.893,
			"reduction_ratio_front": 4.624,
			"reduction_ratio_rear": 2.928
		},
		"fuel_efficiency": "ミラーサイクルエンジン アイドリングストップ機構 筒内直接噴射 可変バルブタイミング 気筒休止 充電制御 ロックアップ機構付トルクコンバーター 電動パワーステアリング"
	}`
	httpReq, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/cars/1", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestPatchCarSuccessMakerPartialColumn(t *testing.T) {
	seedTestCarCx5()       // テストデータの準備
	token := login("user") // 認証実行
	// HTTPリクエストの生成
	body := `{
		"maker_name": "マツダ"
	}`
	httpReq, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/cars/1", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestPatchCarSuccessNoColumn(t *testing.T) {
	seedTestCarCx5()       // テストデータの準備
	token := login("user") // 認証実行
	// HTTPリクエストの生成
	body := `{
	}`
	httpReq, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/cars/1", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestPatchCarNoRecord(t *testing.T) {
	seedTestCarCx5()       // テストデータの準備
	token := login("user") // 認証実行
	// HTTPリクエストの生成
	body := `{
	}`
	httpReq, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/cars/123", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 404, recorder.Result().StatusCode)
}

func TestPatchCarNoLogin(t *testing.T) {
	seedTestCarCx5() // テストデータの準備
	// HTTPリクエストの生成
	body := `{
	}`
	httpReq, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/cars/123", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "token"))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 401, recorder.Result().StatusCode)
}

func TestDeleteCarSuccessUser(t *testing.T) {
	seedTestUser()
	DB.Exec("TRUNCATE TABLE cars")
	seedTestCarCx5() // テストデータの準備
	token := login("admin")
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/cars/1", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 204, recorder.Result().StatusCode)
	//
	// httpReq2, err2 := http.NewRequest(http.MethodGet, "http://localhost:8080/cars", nil)
	// httpReq2.Header.Add("Content-Type", "application/json")
	// httpReq2.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	// if err2 != nil {
	// 	panic(err)
	// }
	// recorder2 := ServeAndRequest(httpReq2)
	// // テストケース固有のチェック
	// assert.Equal(t, 200, recorder2.Result().StatusCode)
	// var body []Car

	// json.Unmarshal(recorder2.Body.Bytes(), &body)
	// assert.Equal(t, 0, len(body))
}

func TestDeleteCarNoRecord(t *testing.T) {
	seedTestUser()                 // テストデータの準備
	DB.Exec("TRUNCATE TABLE cars") // 認証実行
	token := login("admin")
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/cars/123", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 404, recorder.Result().StatusCode)
}

func TestDeleteCarNoLogin(t *testing.T) {
	seedTestUser() // テストデータの準備
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/cars/1", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "token"))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 401, recorder.Result().StatusCode)
}

func TestPostCarSuccessUser(t *testing.T) {
	DB.Exec("TRUNCATE TABLE cars")
	seedTestUser()         // テストデータの準備
	token := login("user") // 認証実行
	// HTTPリクエストの生成
	body := `{
		"maker_name": "マツダ",
		"model_name": "CX-5",
		"grade_name": "25S Proactive",
		"model_code": "6BA-KF5P",
		"price": 3140500,
		"url": "https://www.mazda.co.jp/cars/cx-5/",
		"image_url": "https://upload.wikimedia.org/wikipedia/commons/8/85/2017_Mazda_CX-5_%28KF%29_Maxx_2WD_wagon_%282018-11-02%29_01.jpg",
		"model_change_full": "2016-12-01T00:00:00+09:00",
		"model_change_last": "2018-01-01T00:00:00+09:00",
		"body": {
			"type": "SUV",
			"length": 4545,
			"width": 1840,
			"height": 1690,
			"wheel_base": 2700,
			"tread_front": 1595,
			"tread_rear": 1595,
			"min_road_clearance": 210,
			"body_weight": 1620,
			"doors": 4
		},
		"interior": {
			"length": 1890,
			"width": 1540,
			"height": 1265,
			"luggage_cap": null,
			"riding_cap": 5
		},
		"performance": {
			"min_turning_radius": 5.5,
			"fcr_wltc": 13,
			"fcr_wltc_l": 10.2,
			"fcr_wltc_m": 13.4,
			"fcr_wltc_h": 14.7,
			"fcr_wltc_exh": null,
			"fcr_jc08": 14.2,
			"mpc_wltc": null,
			"ecr_wltc": null,
			"ecr_wltc_l": null,
			"ecr_wltc_m": null,
			"ecr_wltc_h": null,
			"ecr_wltc_exh": null,
			"ecr_jc08": null,
			"mpc_jc08": null
		},
		"power_train": "ICE",
		"drive_system": "AWD",
		"engine": {
			"code": "PY-RPS",
			"type": "水冷直列4気筒DOHC16バルブ",
			"cylinders": 4,
			"cylinder_layout": "I",
			"valve_system": "DOHC",
			"displacement": 2.488,
			"bore": 89,
			"stroke": 100,
			"compression_ratio": 13,
			"max_output": 138,
			"max_output_lower_rpm": 6000,
			"max_output_upper_rpm": 6000,
			"max_torque": 250,
			"max_torque_lower_rpm": 4000,
			"max_torque_upper_rpm": 4000,
			"fuel_system": "DI",
			"fuel_type": "無鉛レギュラーガソリン",
			"fuel_tank_cap": 58
		},
		"motor_x": {
			"code": null,
			"type": null,
			"purpose": null,
			"rated_output": null,
			"max_output": null,
			"max_output_lower_rpm": null,
			"max_output_upper_rpm": null,
			"max_torque": null,
			"max_torque_lower_rpm": null,
			"max_torque_upper_rpm": null
		},
		"motor_y": {
			"code": null,
			"type": null,
			"purpose": null,
			"rated_output": null,
			"max_output": null,
			"max_output_lower_rpm": null,
			"max_output_upper_rpm": null,
			"max_torque": null,
			"max_torque_lower_rpm": null,
			"max_torque_upper_rpm": null
		},
		"battery": {
			"type": null,
			"quantity": null,
			"voltage": null,
			"capacity": null
		},
		"steering": "ラック&ピニオン式",
		"suspension_front": "マクファーソンストラット式",
		"suspension_rear": "マルチリンク式",
		"brake_front": "ベンチレーテッドディスク",
		"brake_rear": "ディスク",
		"tire_front": {
			"section_width": 225,
			"aspect_ratio": 55,
			"wheel_diameter": 19
		},
		"tire_rear": {
			"section_width": 225,
			"aspect_ratio": 55,
			"wheel_diameter": 19
		},
		"transmission": {
			"type": "AT",
			"gears": 6,
			"ratio_1": 3.552,
			"ratio_2": 2.022,
			"ratio_3": 1.452,
			"ratio_4": 1,
			"ratio_5": 0.708,
			"ratio_6": 0.599,
			"ratio_7": null,
			"ratio_8": null,
			"ratio_9": null,
			"ratio_10": null,
			"ratio_rear": 3.893,
			"reduction_ratio_front": 4.624,
			"reduction_ratio_rear": 2.928
		},
		"fuel_efficiency": "ミラーサイクルエンジン アイドリングストップ機構 筒内直接噴射 可変バルブタイミング 気筒休止 充電制御 ロックアップ機構付トルクコンバーター 電動パワーステアリング"
	}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/cars", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 201, recorder.Result().StatusCode)
}

func TestPostCarSuccessAdmin(t *testing.T) {
	DB.Exec("TRUNCATE TABLE cars")
	seedTestUser()          // テストデータの準備
	token := login("admin") // 認証実行
	// HTTPリクエストの生成
	body := `{
		"maker_name": "マツダ",
		"model_name": "CX-5",
		"grade_name": "25S Proactive",
		"model_code": "6BA-KF5P",
		"price": 3140500,
		"url": "https://www.mazda.co.jp/cars/cx-5/",
		"image_url": "https://upload.wikimedia.org/wikipedia/commons/8/85/2017_Mazda_CX-5_%28KF%29_Maxx_2WD_wagon_%282018-11-02%29_01.jpg",
		"model_change_full": "2016-12-01T00:00:00+09:00",
		"model_change_last": "2018-01-01T00:00:00+09:00",
		"body": {
			"type": "SUV",
			"length": 4545,
			"width": 1840,
			"height": 1690,
			"wheel_base": 2700,
			"tread_front": 1595,
			"tread_rear": 1595,
			"min_road_clearance": 210,
			"body_weight": 1620,
			"doors": 4
		},
		"interior": {
			"length": 1890,
			"width": 1540,
			"height": 1265,
			"luggage_cap": null,
			"riding_cap": 5
		},
		"performance": {
			"min_turning_radius": 5.5,
			"fcr_wltc": 13,
			"fcr_wltc_l": 10.2,
			"fcr_wltc_m": 13.4,
			"fcr_wltc_h": 14.7,
			"fcr_wltc_exh": null,
			"fcr_jc08": 14.2,
			"mpc_wltc": null,
			"ecr_wltc": null,
			"ecr_wltc_l": null,
			"ecr_wltc_m": null,
			"ecr_wltc_h": null,
			"ecr_wltc_exh": null,
			"ecr_jc08": null,
			"mpc_jc08": null
		},
		"power_train": "ICE",
		"drive_system": "AWD",
		"engine": {
			"code": "PY-RPS",
			"type": "水冷直列4気筒DOHC16バルブ",
			"cylinders": 4,
			"cylinder_layout": "I",
			"valve_system": "DOHC",
			"displacement": 2.488,
			"bore": 89,
			"stroke": 100,
			"compression_ratio": 13,
			"max_output": 138,
			"max_output_lower_rpm": 6000,
			"max_output_upper_rpm": 6000,
			"max_torque": 250,
			"max_torque_lower_rpm": 4000,
			"max_torque_upper_rpm": 4000,
			"fuel_system": "DI",
			"fuel_type": "無鉛レギュラーガソリン",
			"fuel_tank_cap": 58
		},
		"motor_x": {
			"code": null,
			"type": null,
			"purpose": null,
			"rated_output": null,
			"max_output": null,
			"max_output_lower_rpm": null,
			"max_output_upper_rpm": null,
			"max_torque": null,
			"max_torque_lower_rpm": null,
			"max_torque_upper_rpm": null
		},
		"motor_y": {
			"code": null,
			"type": null,
			"purpose": null,
			"rated_output": null,
			"max_output": null,
			"max_output_lower_rpm": null,
			"max_output_upper_rpm": null,
			"max_torque": null,
			"max_torque_lower_rpm": null,
			"max_torque_upper_rpm": null
		},
		"battery": {
			"type": null,
			"quantity": null,
			"voltage": null,
			"capacity": null
		},
		"steering": "ラック&ピニオン式",
		"suspension_front": "マクファーソンストラット式",
		"suspension_rear": "マルチリンク式",
		"brake_front": "ベンチレーテッドディスク",
		"brake_rear": "ディスク",
		"tire_front": {
			"section_width": 225,
			"aspect_ratio": 55,
			"wheel_diameter": 19
		},
		"tire_rear": {
			"section_width": 225,
			"aspect_ratio": 55,
			"wheel_diameter": 19
		},
		"transmission": {
			"type": "AT",
			"gears": 6,
			"ratio_1": 3.552,
			"ratio_2": 2.022,
			"ratio_3": 1.452,
			"ratio_4": 1,
			"ratio_5": 0.708,
			"ratio_6": 0.599,
			"ratio_7": null,
			"ratio_8": null,
			"ratio_9": null,
			"ratio_10": null,
			"ratio_rear": 3.893,
			"reduction_ratio_front": 4.624,
			"reduction_ratio_rear": 2.928
		},
		"fuel_efficiency": "ミラーサイクルエンジン アイドリングストップ機構 筒内直接噴射 可変バルブタイミング 気筒休止 充電制御 ロックアップ機構付トルクコンバーター 電動パワーステアリング"
	}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/cars", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 201, recorder.Result().StatusCode)
}

func TestGetCarsMakers(t *testing.T) {
	DB.Exec("TRUNCATE TABLE cars")
	seedTestCarCx5()
	seedTestCarCorollaTouring()
	seedTestCarHondaE()
	seedTestCarNote()
	seedTestCarThree()
	seedTestCarNsx()
	token := login("user")
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/cars/makers", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
	var makers []string
	json.Unmarshal(recorder.Body.Bytes(), &makers)
	assert.Equal(t, 5, len(makers))
}

func TestGetCarsMakersModels(t *testing.T) {
	DB.Exec("TRUNCATE TABLE cars")
	seedTestCarCx5()
	seedTestCarCorollaTouring()
	seedTestCarHondaE()
	seedTestCarNote()
	seedTestCarThree()
	seedTestCarNsx()
	token := login("user")
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/cars/makers/models?maker_name=マツダ", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
	var models []string
	json.Unmarshal(recorder.Body.Bytes(), &models)
	assert.Equal(t, 1, len(models))
}

func TestGetCarsMakersModelsNothing(t *testing.T) {
	DB.Exec("TRUNCATE TABLE cars")
	seedTestCarCx5()
	seedTestCarCorollaTouring()
	seedTestCarHondaE()
	seedTestCarNote()
	seedTestCarThree()
	seedTestCarNsx()
	token := login("user")
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/cars/makers/models?maker_name=存在しないメーカー", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
	var models []string
	json.Unmarshal(recorder.Body.Bytes(), &models)
	assert.Equal(t, 0, len(models))
}

func TestGetCarsBodyTypes(t *testing.T) {
	DB.Exec("TRUNCATE TABLE cars")
	seedTestCarCx5()
	seedTestCarCorollaTouring()
	seedTestCarHondaE()
	seedTestCarNote()
	seedTestCarThree()
	seedTestCarNsx()
	token := login("user")
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/cars/body_types", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
	var makers []string
	json.Unmarshal(recorder.Body.Bytes(), &makers)
	assert.Equal(t, 4, len(makers))
}
