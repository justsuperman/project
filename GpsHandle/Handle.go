package GpsHandle

import (
	"strconv"
	"strings"
	"time"
)

type GPRMC struct {
	Time        time.Time
	Status      byte
	Latitde     float64
	SN          byte
	Longitude   float64
	EW          byte
	Speed       float64
	Direction   float64
	Declination float64
	Dec_EW      byte
	Mode        string
}

func (gps *GPRMC) DecodeData(data []byte) error {
	gps.Status = 'V'
	var err error
	gprsData := strings.Split(string(data), ",")
	if gprsData[0] == "$GPRMC" && gprsData[2] == "A" {
		gps.Status = 'A'
		gps.Latitde, err = strconv.ParseFloat(gprsData[3], 64)
		if err != nil {
			return err
		}
		gps.SN = []byte(gprsData[4])[0]
		gps.Longitude, err = strconv.ParseFloat(gprsData[5], 64)
		if err != nil {
			return err
		}
		gps.EW = []byte(gprsData[6])[0]
		gps.Speed, err = strconv.ParseFloat(gprsData[7], 64)
		if err != nil {
			return err
		}
		gps.Direction, err = strconv.ParseFloat(gprsData[8], 64)
		if err != nil {
			return err
		}
		times := gprsData[1][0:6]
		dates := gprsData[9]
		gps.Time, err = time.Parse("020106150405", dates+times)
		if err != nil {
			return err
		}
		gps.Mode = gprsData[12]
	} else {
		return err
	}
	return nil
}

func (data GPRMC) RTD() (float64, float64) {
	LatitdeH := int(data.Latitde / 100)
	LatitdeM := int(data.Latitde - float64(LatitdeH)*100)
	LatitdeS := (data.Latitde - float64(100*LatitdeH+LatitdeM)) * 60

	LongitudeH := int(data.Longitude / 100)
	LongitudeM := int(data.Longitude - float64(LongitudeH)*100)
	LongitudeS := (data.Longitude - float64(100*LongitudeH+LongitudeM)) * 60

	Latitde := float64(LatitdeH) + float64(LatitdeM)/60 + LatitdeS/3600
	Longitude := float64(LongitudeH) + float64(LongitudeM)/60 + LongitudeS/3600
	return Latitde, Longitude
}
