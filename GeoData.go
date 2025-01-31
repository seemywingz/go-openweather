package openweather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type GeoData []struct {
	Name       string `json:"name"`
	LocalNames struct {
		Ms          string `json:"ms"`
		Gu          string `json:"gu"`
		Is          string `json:"is"`
		Wa          string `json:"wa"`
		Mg          string `json:"mg"`
		Gl          string `json:"gl"`
		Om          string `json:"om"`
		Ku          string `json:"ku"`
		Tw          string `json:"tw"`
		Mk          string `json:"mk"`
		Ee          string `json:"ee"`
		Fj          string `json:"fj"`
		Gd          string `json:"gd"`
		Ky          string `json:"ky"`
		Yo          string `json:"yo"`
		Zu          string `json:"zu"`
		Bg          string `json:"bg"`
		Tk          string `json:"tk"`
		Co          string `json:"co"`
		Sh          string `json:"sh"`
		De          string `json:"de"`
		Kl          string `json:"kl"`
		Bi          string `json:"bi"`
		Km          string `json:"km"`
		Lt          string `json:"lt"`
		Fi          string `json:"fi"`
		Fy          string `json:"fy"`
		Ba          string `json:"ba"`
		Sc          string `json:"sc"`
		FeatureName string `json:"feature_name"`
		Ja          string `json:"ja"`
		Am          string `json:"am"`
		Sk          string `json:"sk"`
		Mr          string `json:"mr"`
		Es          string `json:"es"`
		Sq          string `json:"sq"`
		Te          string `json:"te"`
		Br          string `json:"br"`
		Uz          string `json:"uz"`
		Da          string `json:"da"`
		Sw          string `json:"sw"`
		Fa          string `json:"fa"`
		Sr          string `json:"sr"`
		Cu          string `json:"cu"`
		Ln          string `json:"ln"`
		Na          string `json:"na"`
		Wo          string `json:"wo"`
		Ig          string `json:"ig"`
		To          string `json:"to"`
		Ta          string `json:"ta"`
		Mt          string `json:"mt"`
		Ar          string `json:"ar"`
		Su          string `json:"su"`
		Ab          string `json:"ab"`
		Ps          string `json:"ps"`
		Bm          string `json:"bm"`
		Mi          string `json:"mi"`
		Kn          string `json:"kn"`
		Kv          string `json:"kv"`
		Os          string `json:"os"`
		Bn          string `json:"bn"`
		Li          string `json:"li"`
		Vi          string `json:"vi"`
		Zh          string `json:"zh"`
		Eo          string `json:"eo"`
		Ha          string `json:"ha"`
		Tt          string `json:"tt"`
		Lb          string `json:"lb"`
		Ce          string `json:"ce"`
		Hu          string `json:"hu"`
		It          string `json:"it"`
		Tl          string `json:"tl"`
		Pl          string `json:"pl"`
		Sm          string `json:"sm"`
		En          string `json:"en"`
		Vo          string `json:"vo"`
		El          string `json:"el"`
		Sn          string `json:"sn"`
		Fr          string `json:"fr"`
		Cs          string `json:"cs"`
		Io          string `json:"io"`
		Hi          string `json:"hi"`
		Et          string `json:"et"`
		Pa          string `json:"pa"`
		Av          string `json:"av"`
		Ko          string `json:"ko"`
		Bh          string `json:"bh"`
		Yi          string `json:"yi"`
		Sa          string `json:"sa"`
		Sl          string `json:"sl"`
		Hr          string `json:"hr"`
		Si          string `json:"si"`
		So          string `json:"so"`
		Gn          string `json:"gn"`
		Ay          string `json:"ay"`
		Se          string `json:"se"`
		Sd          string `json:"sd"`
		Af          string `json:"af"`
		Ga          string `json:"ga"`
		Or          string `json:"or"`
		Ia          string `json:"ia"`
		Ie          string `json:"ie"`
		Ug          string `json:"ug"`
		Nl          string `json:"nl"`
		Gv          string `json:"gv"`
		Qu          string `json:"qu"`
		Be          string `json:"be"`
		An          string `json:"an"`
		Fo          string `json:"fo"`
		Hy          string `json:"hy"`
		Nv          string `json:"nv"`
		Bo          string `json:"bo"`
		ASCII       string `json:"ascii"`
		ID          string `json:"id"`
		Lv          string `json:"lv"`
		Ca          string `json:"ca"`
		No          string `json:"no"`
		Nn          string `json:"nn"`
		Ml          string `json:"ml"`
		My          string `json:"my"`
		Ne          string `json:"ne"`
		He          string `json:"he"`
		Cy          string `json:"cy"`
		Lo          string `json:"lo"`
		Jv          string `json:"jv"`
		Sv          string `json:"sv"`
		Mn          string `json:"mn"`
		Tg          string `json:"tg"`
		Kw          string `json:"kw"`
		Cv          string `json:"cv"`
		Az          string `json:"az"`
		Oc          string `json:"oc"`
		Th          string `json:"th"`
		Ru          string `json:"ru"`
		Ny          string `json:"ny"`
		Bs          string `json:"bs"`
		St          string `json:"st"`
		Ro          string `json:"ro"`
		Rm          string `json:"rm"`
		Ff          string `json:"ff"`
		Kk          string `json:"kk"`
		Uk          string `json:"uk"`
		Pt          string `json:"pt"`
		Tr          string `json:"tr"`
		Eu          string `json:"eu"`
		Ht          string `json:"ht"`
		Ka          string `json:"ka"`
		Ur          string `json:"ur"`
	} `json:"local_names"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
	State   string  `json:"state"`
}

// GetGeoData: Get GeoData from OpenWeatherMap API
func GetGeoData(location, apiKey string) (GeoData, error) {

	encLocation := url.QueryEscape(location)
	url := "http://api.openweathermap.org/geo/1.0/direct?q=" + encLocation + "&limit=1&appid=" + apiKey

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get geo data: %s", resp.Status)
	}

	var geoData GeoData
	if err := json.NewDecoder(resp.Body).Decode(&geoData); err != nil {
		return nil, err
	}

	return geoData, nil
}
