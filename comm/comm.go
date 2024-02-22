package comm

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"vprenner.com/protos"
)

type TelescopeData struct {
	Declination    float64
	RightAscension float64
}

type TelescopeLocation struct {
	Latitude  float64
	Longitude float64
	Offset    float64
	LocalTime time.Time
	InitDec   bool
	InitRa    bool
}

var tl = TelescopeLocation{
	Latitude:  0,
	Longitude: 0,
	Offset:    0,
	LocalTime: time.Now(),
	InitDec:   false,
	InitRa:    false,
}

var td = TelescopeData{
	Declination:    0,
	RightAscension: 0,
}

var targetTd = TelescopeData{
	Declination:    0,
	RightAscension: 0,
}

var dc = DwarfConn{}

func init() {

	go func() {
		fmt.Println("Connecting to Dwarf...")
		death := dc.connectToDwarf()
		for {
			select {
			case d := <-death:
				fmt.Println(d)
				if d {
					// restart
					time.Sleep(time.Second)
					fmt.Println("Reconnecting to Dwarf...")
					death = dc.connectToDwarf()
				} else {
					os.Exit(0)
					return
				}
			}
		}
	}()
}

func Send(msg string) string {

	if strings.HasPrefix(msg, ":St") {
		degrees, err1 := strconv.ParseFloat(msg[4:6], 64)
		minutes, err2 := strconv.ParseFloat(msg[7:9], 64)
		if err1 != nil || err2 != nil {
			return "0"
		} else {
			sign := 1.0
			if msg[3] == '-' {
				sign = -1.0
			}
			tl.Latitude = sign * (degrees + (minutes / 60.0))
			return "1"
		}
	} else if strings.HasPrefix(msg, ":Sg") {
		degrees, err1 := strconv.ParseFloat(msg[3:6], 64)
		minutes, err2 := strconv.ParseFloat(msg[7:9], 64)
		if err1 != nil || err2 != nil {
			return "0"
		} else {
			sign := 1.0
			if msg[3] == '-' {
				sign = -1.0
			}
			tl.Longitude = sign * (degrees + (minutes / 60.0))
			return "1"
		}
		/*
			Message incoming: ':Sr01:35:12'
			Message incoming: ':Sd+30*47:01'
		*/
	} else if strings.HasPrefix(msg, ":Sr") {
		hours, err1 := strconv.ParseFloat(msg[3:5], 64)
		minutes, err2 := strconv.ParseFloat(msg[6:8], 64)
		seconds, err3 := strconv.ParseFloat(msg[9:11], 64)
		if err1 != nil || err2 != nil || err3 != nil {
			return "0"
		} else {
			targetTd.RightAscension = hours + ((minutes - 2) / 60) + (seconds / 3600)
			return "1"
		}
	} else if strings.HasPrefix(msg, ":Sd") {
		sign := msg[3:4]
		degrees, err1 := strconv.ParseFloat(msg[4:6], 64)
		minutes, err2 := strconv.ParseFloat(msg[7:9], 64)
		seconds, err3 := strconv.ParseFloat(msg[10:12], 64)
		if err1 != nil || err2 != nil || err3 != nil {
			return "0"
		} else {
			mSign := 1.0
			if sign == "-" {
				mSign = -1.0
			}
			targetTd.Declination = mSign * (degrees + (minutes / 60) + (seconds / 3600))
			return "1"
		}
	} else if strings.HasPrefix(msg, ":SG") {
		offset, err := strconv.ParseFloat(msg[3:], 64)
		if err != nil {
			return "0"
		} else {
			tl.Offset = offset
			return "1"
		}
	} else if strings.HasPrefix(msg, ":SC") {

		lt, err := time.Parse("01/02/06", msg[3:])
		if err != nil {
			return "0"
		} else {
			hour, minute, second := tl.LocalTime.Clock()
			year, month, day := lt.Date()
			tl.LocalTime = time.Date(year, month, day, hour, minute, second, 0, tl.LocalTime.Location())
			return "1"
		}
	} else if strings.HasPrefix(msg, ":SL") {

		timePart, err := time.Parse("15:04:05", msg[3:])
		if err != nil {
			return "0"
		} else {
			hour, minute, second := timePart.Clock()
			year, month, day := tl.LocalTime.Date()
			tl.LocalTime = time.Date(year, month, day, hour, minute, second, 0, tl.LocalTime.Location())
			return "1"
		}
	} else if strings.HasPrefix(msg, ":RM") || strings.HasPrefix(msg, ":RS") || strings.HasPrefix(msg, ":RC") || strings.HasPrefix(msg, ":RG") {
		tl.InitRa = false
		tl.InitDec = false
		td.RightAscension = 0
		td.Declination = 90

		//Dwarf2Client.init(Longitude, Latitude, _localTime.AddHours(_utcOffsetHours).ToString("yyyy-MM-dd HH:mm:ss"), "DWARF_GOTO_LX200"+DateTime.UtcNow.ToString("yyyyMMddHHmmss"))

		return "1"
		//return result==0?"1":"0";
	}

	switch msg {
	case ":MS":
		result := 0
		td.Declination = targetTd.Declination
		td.RightAscension = targetTd.RightAscension

		msgToSend := protos.ReqGotoDSO{Ra: td.RightAscension, Dec: td.Declination, TargetName: "LX200"}
		fmt.Println("Sending to Dwarf: ", msgToSend.String())
		byteMsg, _ := proto.Marshal(&msgToSend)
		dc.sendToDwarf <- byteMsg

		response := <-dc.receiveFromDwarf
		unmarshalledResponse := &protos.ComResponse{}
		_ = proto.Unmarshal(response, unmarshalledResponse)

		fmt.Println("Received from Dwarf: ", unmarshalledResponse)

		return strconv.Itoa(result)
	case ":GR":
		if !tl.InitRa {
			td.RightAscension = 0.0
			tl.InitRa = true
		}
		ra := math.Abs(td.RightAscension)
		hours := int(math.Floor(ra))
		minutes := int(math.Floor((ra - float64(hours)) * 60))
		seconds := int(math.Round((((ra - float64(hours)) * 60) - float64(minutes)) * 60))
		sign := "+"
		if ra < 0 {
			sign = "-"
		}
		res := fmt.Sprintf("%s%02d:%02d.%1d#", sign, hours, minutes, seconds)
		return res
	case ":GD":
		if !tl.InitDec {
			td.Declination = 0.0
			tl.InitDec = true
		}
		dec := math.Abs(td.Declination)
		degrees := int(math.Floor(dec))
		minutes := int(math.Floor((dec - float64(degrees)) * 60))
		// seconds := math.Round((((dec - degrees) * 60) - minutes) * 60)
		sign := "+"
		if td.Declination < 0 {
			sign = "-"
		}
		res := fmt.Sprintf("%s%02d*%02d#", sign, degrees, minutes)
		return res

	case ":CM":
		td.Declination = targetTd.Declination
		td.RightAscension = targetTd.RightAscension
		return "Coordinates matched.#"
	case ":GVP":
		return "LX200 Custom#"
	case ":GVN":
		return "1.0.0#"
	case ":GVD":
		return "2023-05-05#"
	case ":GVZ":
		return "12:34:56#"
	default:
		return "1"
	}
}
