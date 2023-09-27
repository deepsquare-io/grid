// Copyright (C) 2023 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package speedtest

import "encoding/json"

func UnmarshalResult(data []byte) (Result, error) {
	var r Result
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Result) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Result struct {
	Type       string      `json:"type"`
	Timestamp  string      `json:"timestamp"`
	Ping       Ping        `json:"ping"`
	Download   Load        `json:"download"`
	Upload     Load        `json:"upload"`
	PacketLoss uint64      `json:"packetLoss"`
	ISP        string      `json:"isp"`
	Interface  Interface   `json:"interface"`
	Server     Server      `json:"server"`
	Result     ResultClass `json:"result"`
}

type Load struct {
	Bandwidth uint64 `json:"bandwidth"`
	Bytes     uint64 `json:"bytes"`
	Elapsed   uint64 `json:"elapsed"`
	Latency   Ping   `json:"latency"`
}

type Ping struct {
	Iqm     *float64 `json:"iqm,omitempty"`
	Low     float64  `json:"low"`
	High    float64  `json:"high"`
	Jitter  float64  `json:"jitter"`
	Latency *float64 `json:"latency,omitempty"`
}

type Interface struct {
	InternalIP string `json:"internalIp"`
	Name       string `json:"name"`
	MACAddr    string `json:"macAddr"`
	IsVPN      bool   `json:"isVpn"`
	ExternalIP string `json:"externalIp"`
}

type ResultClass struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	Persisted bool   `json:"persisted"`
}

type Server struct {
	ID       uint64 `json:"id"`
	Host     string `json:"host"`
	Port     uint64 `json:"port"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Country  string `json:"country"`
	IP       string `json:"ip"`
}
