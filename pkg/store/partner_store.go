package store

import "github.com/frolosofsky/aroundhome/pkg/model"

type PartnerStore interface {
	MatchPartners(material string, pos model.Position) ([]model.PartnerMatchResult, error)
}
