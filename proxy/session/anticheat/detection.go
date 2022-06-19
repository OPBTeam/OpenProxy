package anticheat

type Detection struct { // bson is to make it look better when storing in MongoDB :)
	Type      DetectionType          `bson:"Type"`
	Arguments map[string]interface{} `bson:"Arguments"`
}

type DetectionType string

const (
	DetectionTypeFly         = "FLY"
	DetectionTypeSpeed       = "SPEED"
	DetectionTypeKillaura    = "KILLAURA"
	DetectionTypeTimer       = "TIMER"
	DetectionTypeAutoclicker = "AUTOCLICKER"
	DetectionNukerA          = "NUKERA"
)
