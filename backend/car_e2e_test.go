package main

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
)

func seedTestCarCx5() {
	DB.Create(&Car{
		MakerName:       "マツダ",
		ModelName:       "CX-5",
		GradeName:       "25S Proactive",
		ModelCode:       "6BA-KF5P",
		Price:           null.NewInt(3140500, true),
		Url:             null.NewString("https://www.mazda.co.jp/cars/cx-5/", true),
		ImageUrl:        null.NewString("https://upload.wikimedia.org/wikipedia/commons/8/85/2017_Mazda_CX-5_%28KF%29_Maxx_2WD_wagon_%282018-11-02%29_01.jpg", true),
		ModelChangeFull: null.NewTime(time.Date(2017, time.February, 1, 0, 0, 0, 0, time.Local), true),
		//ModelChangeFull: null.NewTime(time.Date(2017, time.February, 1, 0, 0, 0, 0, time.Local), false),
		ModelChangeLast: null.NewTime(time.Date(2018, time.January, 1, 0, 0, 0, 0, time.Local), true),
		Body: Body{
			Length:           null.NewInt(4545, true),
			Width:            null.NewInt(1840, true),
			Height:           null.NewInt(1690, true),
			WheelBase:        null.NewInt(2700, true),
			TreadFront:       null.NewInt(1595, true),
			TreadRear:        null.NewInt(1595, true),
			MinRoadClearance: null.NewInt(210, true),
			Weight:           null.NewInt(1620, true),
		},
		Interior: Interior{
			Length: null.NewInt(1890, true),
			Width:  null.NewInt(1540, true),
			Height: null.NewInt(1265, true),
			// LuggageCap: null.NewInt(0, false),
			RidingCap: null.NewInt(5, true),
		},
		Perf: Perf{
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
			Code:               null.NewString("PY-RPS", true),
			Type:               null.NewString("水冷直列4気筒DOHC16バルブ", true),
			Cylinders:          null.NewInt(4, true),
			CylinderLayout:     null.NewString((string)(I), true),
			ValveSystem:        null.NewString((string)(DOHC), true),
			Displacement:       null.NewFloat(2.488, true),
			Bore:               null.NewFloat(89.0, true),
			Stroke:             null.NewFloat(100.0, true),
			CompRatio:          null.NewFloat(13.0, true),
			MaxOutput:          null.NewFloat(138, true),
			MaxOutputLowerRpm:  null.NewFloat(6000, true),
			MaxOutputHigherRpm: null.NewFloat(6000, true),
			MaxTorque:          null.NewFloat(250, true),
			MaxTorqueLowerRpm:  null.NewFloat(4000, true),
			MaxTorqueHigherRpm: null.NewFloat(4000, true),
			FuelSystem:         null.NewString("DI", true),
			FuelType:           null.NewString("無鉛レギュラーガソリン", true),
			FuelTankCap:        null.NewInt(58, true),
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
	DB.Create(&Car{
		MakerName:       "トヨタ",
		ModelName:       "カローラツーリング",
		GradeName:       "HYBRID G-X E-Four",
		ModelCode:       "6AA-ZWE214W-AWXNB",
		Price:           null.NewInt(2678500, true),
		Url:             null.NewString("https://toyota.jp/corollatouring/", true),
		ImageUrl:        null.NewString("https://upload.wikimedia.org/wikipedia/commons/8/8a/Toyota_COROLLA_TOURING_HYBRID_W%C3%97B_2WD_%286AA-ZWE211W-AWXSB%29_front.jpg", true),
		ModelChangeFull: null.NewTime(time.Date(2019, time.September, 17, 0, 0, 0, 0, time.Local), true),
		ModelChangeLast: null.NewTime(time.Date(2021, time.November, 15, 0, 0, 0, 0, time.Local), true),
		Body: Body{
			Length:           null.NewInt(4495, true),
			Width:            null.NewInt(1745, true),
			Height:           null.NewInt(1460, true),
			WheelBase:        null.NewInt(2640, true),
			TreadFront:       null.NewInt(1530, true),
			TreadRear:        null.NewInt(1540, true),
			MinRoadClearance: null.NewInt(130, true),
			Weight:           null.NewInt(1410, true),
		},
		Interior: Interior{
			Length: null.NewInt(1790, true),
			Width:  null.NewInt(1510, true),
			Height: null.NewInt(1160, true),
			// LuggageCap: null.NewInt(0, false),
			RidingCap: null.NewInt(5, true),
		},
		Perf: Perf{
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
			MaxOutput:          null.NewFloat(72, true),
			MaxOutputLowerRpm:  null.NewFloat(5200, true),
			MaxOutputHigherRpm: null.NewFloat(5200, true),
			MaxTorque:          null.NewFloat(142, true),
			MaxTorqueLowerRpm:  null.NewFloat(3600, true),
			MaxTorqueHigherRpm: null.NewFloat(3600, true),
			FuelSystem:         null.NewString("電子制御式燃料噴射装置(EFI)", true),
			FuelType:           null.NewString("無鉛レギュラーガソリン", true),
			FuelTankCap:        null.NewInt(43, true),
		},
		MotorX: Motor{
			Code:    null.NewString("1NM", true),
			Type:    null.NewString("交流同期電動機", true),
			Purpose: null.NewString((string)(TRACTION_FRONT), true),
			//RatedOutput: ,
			MaxOutput: null.NewFloat(53, true),
			//MaxOutputLowerRpm: ,
			//MaxOutputHigherRpm: ,
			MaxTorque: null.NewFloat(163, true),
			//MaxTorqueLowerRpm: ,
			//MaxTorqueHigherRpm: ,
		},
		MotorY: Motor{
			Code:    null.NewString("1MM", true),
			Type:    null.NewString("交流同期電動機", true),
			Purpose: null.NewString((string)(TRACTION_REAR), true),
			//RatedOutput: ,
			MaxOutput: null.NewFloat(5.3, true),
			//MaxOutputLowerRpm: ,
			//MaxOutputHigherRpm: ,
			MaxTorque: null.NewFloat(55, true),
			//MaxTorqueLowerRpm: ,
			//MaxTorqueHigherRpm: ,
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
	DB.Create(&Car{
		MakerName:       "ホンダ",
		ModelName:       "NSX",
		GradeName:       "Type S",
		ModelCode:       "5AA-NC1",
		Price:           null.NewInt(27940000, true),
		Url:             null.NewString("https://www.honda.co.jp/NSX/types/", true),
		ImageUrl:        null.NewString("https://upload.wikimedia.org/wikipedia/commons/e/ea/2019_Honda_NSX_3.5_CAA-NC1_%2820190722%29_01.jpg", true),
		ModelChangeFull: null.NewTime(time.Date(2017, time.February, 27, 0, 0, 0, 0, time.Local), true),
		ModelChangeLast: null.NewTime(time.Date(2021, time.August, 30, 0, 0, 0, 0, time.Local), true),
		Body: Body{
			Length:           null.NewInt(4535, true),
			Width:            null.NewInt(1940, true),
			Height:           null.NewInt(1215, true),
			WheelBase:        null.NewInt(2630, true),
			TreadFront:       null.NewInt(1665, true),
			TreadRear:        null.NewInt(1635, true),
			MinRoadClearance: null.NewInt(110, true),
			Weight:           null.NewInt(1790, true),
		},
		Interior: Interior{
			// Length: null.NewInt(1790, true),
			// Width:  null.NewInt(1510, true),
			// Height: null.NewInt(1160, true),
			// LuggageCap: null.NewInt(0, false),
			RidingCap: null.NewInt(2, true),
		},
		Perf: Perf{
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
			Code:               null.NewString("JNC", true),
			Type:               null.NewString("水冷V型6気筒縦置", true),
			Cylinders:          null.NewInt(6, true),
			CylinderLayout:     null.NewString((string)(V), true),
			ValveSystem:        null.NewString((string)(DOHC), true),
			Displacement:       null.NewFloat(3.492, true),
			Bore:               null.NewFloat(91.0, true),
			Stroke:             null.NewFloat(89.5, true),
			CompRatio:          null.NewFloat(10.0, false),
			MaxOutput:          null.NewFloat(389, true),
			MaxOutputLowerRpm:  null.NewFloat(6500, true),
			MaxOutputHigherRpm: null.NewFloat(6850, true),
			MaxTorque:          null.NewFloat(600, true),
			MaxTorqueLowerRpm:  null.NewFloat(2300, true),
			MaxTorqueHigherRpm: null.NewFloat(6000, true),
			FuelSystem:         null.NewString("電子制御燃料噴射式(ホンダ PGM-FI)", true),
			FuelType:           null.NewString("無鉛プレミアムガソリン", true),
			FuelTankCap:        null.NewInt(59, true),
		},
		MotorX: Motor{
			Code:    null.NewString("H3", true),
			Type:    null.NewString("交流同期電動機", true),
			Purpose: null.NewString((string)(TRACTION_FRONT), true),
			//RatedOutput: ,
			MaxOutput:          null.NewFloat(27, true),
			MaxOutputLowerRpm:  null.NewFloat(4000, true),
			MaxOutputHigherRpm: null.NewFloat(4000, true),
			MaxTorque:          null.NewFloat(73, true),
			MaxTorqueLowerRpm:  null.NewFloat(0, true),
			MaxTorqueHigherRpm: null.NewFloat(2000, true),
		},
		MotorY: Motor{
			Code:    null.NewString("H2", true),
			Type:    null.NewString("交流同期電動機", true),
			Purpose: null.NewString((string)(TRACTION_REAR), true),
			//RatedOutput: ,
			MaxOutput:          null.NewFloat(35, true),
			MaxOutputLowerRpm:  null.NewFloat(3000, true),
			MaxOutputHigherRpm: null.NewFloat(3000, true),
			MaxTorque:          null.NewFloat(148, true),
			MaxTorqueLowerRpm:  null.NewFloat(500, true),
			MaxTorqueHigherRpm: null.NewFloat(2000, true),
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
	DB.Create(&Car{
		MakerName:       "ホンダ",
		ModelName:       "Honda e",
		GradeName:       "Honda e Advance",
		ModelCode:       "ZAA-ZC7",
		Price:           null.NewInt(4950000, true),
		Url:             null.NewString("https://www.honda.co.jp/honda-e/", true),
		ImageUrl:        null.NewString("https://upload.wikimedia.org/wikipedia/commons/9/9e/Honda_e_Advance_%28ZAA-ZC7%29_front.jpg", true),
		ModelChangeFull: null.NewTime(time.Date(2020, time.August, 27, 0, 0, 0, 0, time.Local), true),
		ModelChangeLast: null.NewTime(time.Date(2020, time.August, 27, 0, 0, 0, 0, time.Local), true),
		Body: Body{
			Length:           null.NewInt(3895, true),
			Width:            null.NewInt(1750, true),
			Height:           null.NewInt(1510, true),
			WheelBase:        null.NewInt(2530, true),
			TreadFront:       null.NewInt(1510, true),
			TreadRear:        null.NewInt(1505, true),
			MinRoadClearance: null.NewInt(145, true),
			Weight:           null.NewInt(1540, true),
		},
		Interior: Interior{
			Length: null.NewInt(1845, true),
			Width:  null.NewInt(1385, true),
			Height: null.NewInt(1120, true),
			// LuggageCap: null.NewInt(0, false),
			RidingCap: null.NewInt(4, true),
		},
		Perf: Perf{
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
		// 	MaxOutputHigherRpm: null.NewFloat(5200, true),
		// 	MaxTorque:          null.NewFloat(142, true),
		// 	MaxTorqueLowerRpm:  null.NewFloat(3600, true),
		// 	MaxTorqueHigherRpm: null.NewFloat(3600, true),
		// 	FuelSystem:         null.NewString( "電子制御式燃料噴射装置(EFI)", true),
		// 	FuelType:           null.NewString( "無鉛レギュラーガソリン", true),
		// 	FuelTankCap:        null.NewInt(43, true),
		// },
		MotorX: Motor{
			Code:               null.NewString("MCF5", true),
			Type:               null.NewString("交流同期電動機", true),
			Purpose:            null.NewString((string)(TRACTION_REAR), true),
			RatedOutput:        null.NewFloat(60, true),
			MaxOutput:          null.NewFloat(113, true),
			MaxOutputLowerRpm:  null.NewFloat(3497, true),
			MaxOutputHigherRpm: null.NewFloat(10000, true),
			MaxTorque:          null.NewFloat(315, true),
			MaxTorqueLowerRpm:  null.NewFloat(0, true),
			MaxTorqueHigherRpm: null.NewFloat(2000, true),
		},
		// MotorY: Motor{
		// 	Code:    null.NewString( "1MM", true),
		// 	Type:    null.NewString( "交流同期電動機", true),
		// 	Purpose: null.NewString( (string)(TRACTION_REAR), true),
		// 	RatedOutput: ,
		// 	MaxOutput: null.NewFloat(5.3, true),
		// 	MaxOutputLowerRpm: ,
		// 	MaxOutputHigherRpm: ,
		// 	MaxTorque: null.NewFloat(55, true),
		// 	MaxTorqueLowerRpm: ,
		// 	MaxTorqueHigherRpm: ,
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
	DB.Create(&Car{
		MakerName:       "日産",
		ModelName:       "ノート",
		GradeName:       "X FOUR",
		ModelCode:       "6AA-SNE13",
		Price:           null.NewInt(2445300, true),
		Url:             null.NewString("https://www3.nissan.co.jp/vehicles/new/note.html", true),
		ImageUrl:        null.NewString("https://upload.wikimedia.org/wikipedia/commons/0/0a/Nissan_Note_e-POWER_%28E13%29%2C_2021%2C_front-left.jpg", true),
		ModelChangeFull: null.NewTime(time.Date(2020, time.November, 24, 0, 0, 0, 0, time.Local), true),
		ModelChangeLast: null.NewTime(time.Date(2021, time.November, 4, 0, 0, 0, 0, time.Local), true),
		Body: Body{
			Length:           null.NewInt(4045, true),
			Width:            null.NewInt(1695, true),
			Height:           null.NewInt(1520, true),
			WheelBase:        null.NewInt(2580, true),
			TreadFront:       null.NewInt(1490, true),
			TreadRear:        null.NewInt(1490, true),
			MinRoadClearance: null.NewInt(125, true),
			Weight:           null.NewInt(1340, true),
		},
		Interior: Interior{
			Length: null.NewInt(2030, true),
			Width:  null.NewInt(1445, true),
			Height: null.NewInt(1240, true),
			// LuggageCap: null.NewInt(0, false),
			RidingCap: null.NewInt(5, true),
		},
		Perf: Perf{
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
			Code:               null.NewString("HR12DE", true),
			Type:               null.NewString("DOHC水冷直列3気筒", true),
			Cylinders:          null.NewInt(3, true),
			CylinderLayout:     null.NewString((string)(I), true),
			ValveSystem:        null.NewString((string)(DOHC), true),
			Displacement:       null.NewFloat(1.198, true),
			Bore:               null.NewFloat(78.0, true),
			Stroke:             null.NewFloat(83.6, true),
			CompRatio:          null.NewFloat(12.0, false),
			MaxOutput:          null.NewFloat(60, true),
			MaxOutputLowerRpm:  null.NewFloat(6000, true),
			MaxOutputHigherRpm: null.NewFloat(6000, true),
			MaxTorque:          null.NewFloat(103, true),
			MaxTorqueLowerRpm:  null.NewFloat(4800, true),
			MaxTorqueHigherRpm: null.NewFloat(4800, true),
			FuelSystem:         null.NewString("ニッサンEGI(ECCS)電子制御燃料噴射装置", true),
			FuelType:           null.NewString("無鉛レギュラーガソリン", true),
			FuelTankCap:        null.NewInt(36, true),
		},
		MotorX: Motor{
			Code:    null.NewString("EM47", true),
			Type:    null.NewString("交流同期電動機", true),
			Purpose: null.NewString((string)(GENERATOR), true),
			//RatedOutput: ,
			MaxOutput:          null.NewFloat(85, true),
			MaxOutputLowerRpm:  null.NewFloat(2900, true),
			MaxOutputHigherRpm: null.NewFloat(10341, true),
			MaxTorque:          null.NewFloat(280, true),
			MaxTorqueLowerRpm:  null.NewFloat(0, true),
			MaxTorqueHigherRpm: null.NewFloat(2900, true),
		},
		MotorY: Motor{
			Code:    null.NewString("MM48", true),
			Type:    null.NewString("交流同期電動機", true),
			Purpose: null.NewString((string)(TRACTION_REAR), true),
			// RatedOutput: ,
			MaxOutput:          null.NewFloat(50, true),
			MaxOutputLowerRpm:  null.NewFloat(4775, true),
			MaxOutputHigherRpm: null.NewFloat(10024, true),
			MaxTorque:          null.NewFloat(100, true),
			MaxTorqueLowerRpm:  null.NewFloat(0, true),
			MaxTorqueHigherRpm: null.NewFloat(4775, true),
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
	DB.Create(&Car{
		MakerName: "BMW",
		ModelName: "3シリーズツーリング",
		GradeName: "320d xDriveツーリング Standard",
		ModelCode: "3DA-6L20",
		Price:     null.NewInt(6340000, true),
		Url:       null.NewString("https://www.bmw.co.jp/ja/all-models/3-series/touring/2019/bmw-3-series-touring-inspire.html", true),
		// ImageUrl:        null.NewString( "", true),
		ModelChangeFull: null.NewTime(time.Date(2019, time.September, 26, 0, 0, 0, 0, time.Local), true),
		ModelChangeLast: null.NewTime(time.Date(2019, time.September, 26, 0, 0, 0, 0, time.Local), true),
		Body: Body{
			Length:           null.NewInt(4715, true),
			Width:            null.NewInt(1825, true),
			Height:           null.NewInt(1475, true),
			WheelBase:        null.NewInt(2850, true),
			TreadFront:       null.NewInt(1575, true),
			TreadRear:        null.NewInt(1590, true),
			MinRoadClearance: null.NewInt(135, true),
			Weight:           null.NewInt(1730, true),
		},
		Interior: Interior{
			// Length: null.NewInt(1890, true),
			// Width:  null.NewInt(1540, true),
			// Height: null.NewInt(1265, true),
			LuggageCap: null.NewInt(500, false),
			RidingCap:  null.NewInt(5, true),
		},
		Perf: Perf{
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
			MaxOutput:          null.NewFloat(140, true),
			MaxOutputLowerRpm:  null.NewFloat(4000, true),
			MaxOutputHigherRpm: null.NewFloat(4000, true),
			MaxTorque:          null.NewFloat(400, true),
			MaxTorqueLowerRpm:  null.NewFloat(1750, true),
			MaxTorqueHigherRpm: null.NewFloat(2500, true),
			FuelSystem:         null.NewString("デジタル・ディーゼル・エレクトロニクス(DDE/電子燃料噴射装置)", true),
			FuelType:           null.NewString("軽油", true),
			FuelTankCap:        null.NewInt(59, true),
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

func TestSearchCars(t *testing.T) {
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
