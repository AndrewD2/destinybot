package main

import (
	"fmt"
	"strings"
)

// Card is structure of the standard card in Destiny as defined by swdestinydb
type Card struct {
	// (+)#MD$ = Melee Damage, (+)#RD$ = Ranged Damage, #F = Focus, #Dc$ = Discard, (+)#R = Resource, #Dr = Disrupt, #Sh$ = Shield, Sp# = Special
	Sides           [6]string `json:"sides"`
	SetCode         string    `json:"set_code"`
	SetName         string    `json:"set_name"`
	TypeCode        string    `json:"type_code"`
	TypeName        string    `json:"type_name"`
	SubtypeCode     string    `json:"subtype_code"`
	SubtypeName     string    `json:"subtype_name"`
	FactionCode     string    `json:"faction_code"`
	FactionName     string    `json:"faction_name"`
	AffiliationCode string    `json:"affliation_code"`
	AffiliationName string    `json:"affiliation_name"`
	RarityCode      string    `json:"rarity_code"`
	RarityName      string    `json:"rarity_name"`
	Position        int       `json:"position"`
	Code            string    `json:"code"`
	Name            string    `json:"name"`
	Subtitle        string    `json:"subtitle"`
	Cost            int       `json:"cost"`
	Health          uint      `json:"health"`
	Points          string    `json:"points"`
	Text            string    `json:"text"`
	DeckLimit       uint      `json:"deck_limit"`
	Flavor          string    `json:"flavor"`
	Illustrator     string    `json:"illustrator"`
	IsUnique        bool      `json:"is_unique"`
	HasDie          bool      `json:"has_die"`
	URL             string    `json:"url"`
	ImageSrc        string    `json:"imagesrc"`
	Label           string    `json:"label"`
	Cp              int       `json:"cp"`
}

func (c Card) String() string {

	r := strings.NewReplacer(boldStart, boldReplace, boldEnd, boldReplace, italicStart, italicReplace, italicEnd, italicReplace)
	removeRedundancy := strings.NewReplacer("````", " ", "**", " ")
	var str string
	if c.IsUnique {
		switch {
		// case strings.Contains(strings.ToLower(c.Name), "luke"):
		// 	str = ":luke: "
		default:
			str = "``•`` "

		}
	}
	str = fmt.Sprintf("%v``%v``\n",
		str, c.Name)
	if c.Subtitle != "" {
		str = fmt.Sprintf("%v%v\n",
			str, c.Subtitle)
	}
	str = fmt.Sprintf("%v\n", str)
	str = fmt.Sprintf("%v%v. %v. %v.\n", str, c.AffiliationName, c.FactionName, c.RarityName)
	if c.TypeName == "Character" {
		str = fmt.Sprintf("%v``%v.``Points: %v. Health: %v.\n", str, c.TypeName, c.Points, c.Health)
	}
	if c.TypeName == "Upgrade" {
		str = fmt.Sprintf("%v``%v - %v.``Cost: %v.\n", str, c.TypeName, c.SubtypeName, c.Cost)
	}
	if c.TypeName == "Support" || c.TypeName == "Event" {
		str = fmt.Sprintf("%v``%v.``Cost: %v\n", str, c.TypeName, c.Cost)
	}
	if c.HasDie {
		str = fmt.Sprintf("%v%v\n", str, c.Sides)
	}
	str = fmt.Sprintf("%v%v\n%v #%d.", str, c.Text, c.SetName, c.Position)

	return removeRedundancy.Replace(r.Replace(str))

}
