package gart

type Palette []HSLColor

var Palettes = map[string]Palette{
	"pleasant": {
		{0.500, 0.33, 0.70},
		{0.436, 0.23, 0.61},
		{0.286, 0.28, 0.66},
		{0.119, 0.70, 0.83},
		{0.072, 0.33, 0.70},
	},
	"vangogh_cafeterrace": {
		{0.1361, 0.7200, 0.5700},
		{0.6000, 0.4400, 0.3800},
		{0.3306, 0.1900, 0.2500},
		{0.0889, 0.7600, 0.5700},
		{0.3361, 0.1900, 0.2200},
		{0.6056, 0.3900, 0.4200},
		{0.1806, 0.5700, 0.8100},
	},
	"vangogh_roadwithcypressandstar": {
		{0.0417, 0.6800, 0.4100},
		{0.1611, 0.9600, 0.8200},
		{0.6333, 0.6100, 0.4300},
		{0.7333, 0.7300, 0.0600},
		{0.7528, 0.9900, 0.0500},
		{0.5833, 0.5300, 0.5600},
		{0.6389, 0.6700, 0.0800},
	},
	"port": {
		{0.0389, 0.8700, 0.5200},
		{0.5944, 1.0000, 0.5000},
		{0.5694, 0.9000, 0.2700},
		{0.4889, 0.3200, 0.2500},
		{0.0472, 0.1300, 0.4100},
		{0.0778, 0.4700, 0.6500},
	},
	"vangogh_potatoeaters": {
		{0.2361, 0.3100, 0.1400},
		{0.1222, 0.6900, 0.2900},
		{0.1139, 0.7000, 0.4300},
		{0.2528, 0.3700, 0.0700},
		{0.1500, 0.2800, 0.4700},
		{0.1361, 0.4500, 0.1400},
		{0.2861, 0.2200, 0.0800},
	},
	"grayscale": {
		{0.0, 0.0, 0.0},
		{0.0, 0.0, 0.1},
		{0.0, 0.0, 0.2},
		{0.0, 0.0, 0.3},
		{0.0, 0.0, 0.4},
		{0.0, 0.0, 0.5},
		{0.0, 0.0, 0.6},
		{0.0, 0.0, 0.7},
	},
	"desert1": {
		{0.6111, 0.2700, 0.3000},
		{0.0833, 0.1600, 0.4900},
		{0.0778, 0.5000, 0.7000},
		{0.0806, 0.5300, 0.8700},
		{0.0861, 0.3800, 0.7600},
		{0.5833, 0.5100, 0.9400},
		{0.9750, 0.0400, 0.6100},
	},
}
