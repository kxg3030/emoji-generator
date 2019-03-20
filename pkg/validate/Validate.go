package validate

import "emoji/pkg/config"

func ExtIsIllegal(ext string)bool  {
	var contain = false
	for _,val := range config.Extension{
		if val == ext{
			contain = true
		}
	}
	return contain
}
