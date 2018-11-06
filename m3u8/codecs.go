package m3u8

import "strings"

var (
	AudioCodecMap = map[string]string{
		"aac-lc": "mp4a.40.2",
		"he-aac": "mp4a.40.5",
		"mp3":    "mp4a.40.34",
	}

	BaselineCodecMap = map[string]string{
		"3.0": "avc1.66.30",
		"3.1": "avc1.42001f",
	}

	MainCodecMap = map[string]string{
		"3.0": "avc1.77.30",
		"3.1": "avc1.4d001f",
		"4.0": "avc1.4d0028",
		"4.1": "avc1.4d0029",
	}

	HighCodecMap = map[string]string{
		"3.0": "avc1.64001e",
		"3.1": "avc1.64001f",
		"3.2": "avc1.640020",
		"4.0": "avc1.640028",
		"4.1": "avc1.640029",
		"4.2": "avc1.64002a",
		"5.0": "avc1.640032",
		"5.1": "avc1.640033",
		"5.2": "avc1.640034",
	}
)

func audioCodec(codec *string) *string {
	if codec == nil {
		return nil
	}

	key := strings.ToLower(*codec)
	value, ok := AudioCodecMap[key]

	if !ok {
		return nil
	}

	return &value
}

func videoCodec(profile *string, level *string) *string {
	if profile == nil || level == nil {
		return nil
	}

	var value string
	var ok bool

	switch *profile {
	case "baseline":
		value, ok = BaselineCodecMap[*level]
	case "main":
		value, ok = MainCodecMap[*level]
	case "high":
		value, ok = HighCodecMap[*level]
	}

	if !ok {
		return nil
	}

	return &value
}
