package conf

import (
	"net/url"
	"sync"

	"github.com/pelletier/go-toml"

	"github.com/guillaumebchd/styx/pkg/model"
)

// GetSites gets all the sites that the rvp can reach from the configuration file
func GetSites(config *toml.Tree) (model.Destination, map[string]*model.Site) {

	siteTree := config.Get("sites").([]*toml.Tree)
	def := config.Get("Default").(*toml.Tree)

	sites := make(map[string]*model.Site)

	for index := range siteTree {

		name := siteTree[index].Get("name").(string)
		entrypoint := siteTree[index].Get("entrypoint").(string)
		destinations := siteTree[index].GetArray("addresses").([]string)
		alives := siteTree[index].GetArray("alives").([]bool)
		weights := siteTree[index].GetArray("weights").([]int64)

		var dest []*model.Destination

		for i := 0; i < len(destinations); i++ {
			var m sync.RWMutex
			url, _ := url.Parse(destinations[i])

			destination := model.Destination{
				URL:    url,
				Alive:  alives[i],
				Weight: weights[i],
				Mux:    &m,
			}

			dest = append(dest, &destination)
		}

		destPool := model.DestinationsPool{
			Destinations: dest,
			Current:      0,
			Total_weight: Somme(weights),
		}

		site := model.Site{
			Name:             name,
			Entrypoint:       entrypoint,
			DestinationsPool: destPool,
		}

		sites[siteTree[index].Get("entrypoint").(string)] = &site
	}

	var m sync.RWMutex
	def_url, _ := url.Parse(def.Get("default_route").(string))

	defaut := model.Destination{
		URL:    def_url,
		Alive:  def.Get("alive").(bool),
		Weight: def.Get("weight").(int64),
		Mux:    &m,
	}

	return defaut, sites
}

func Somme(l []int64) (somme int) {
	for v := range l {
		somme += v
	}
	return
}
