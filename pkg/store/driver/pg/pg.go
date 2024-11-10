package pg

import (
	"database/sql"

	"github.com/frolosofsky/aroundhome/pkg/model"
	"github.com/lib/pq"
)

type Driver struct {
	db *sql.DB
}

func NewDriver(conn string) (*Driver, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	return &Driver{db}, nil
}

func (d *Driver) Health() error {
	var res int
	return d.db.QueryRow("select 1").Scan(&res)
}

func (d *Driver) MatchPartners(material string, pos model.Position) ([]model.PartnerMatchResult, error) {
	q := `
select
  p.id,
  p.name,
  p.rating,
  ST_DistanceSphere(p.geo::geometry, ST_MakePoint($1, $2))::int as dist,
  ST_X(p.geo::geometry), ST_Y(p.geo::geometry)
from partner p
join partner_skill s on s.partner_id = p.id
where
  s.code = $3 and
  ST_DWithin(p.geo, ST_MakePoint($1, $2), p.radius)
order by
  p.rating desc,
  dist asc`

	rows, err := d.db.Query(q, pos.Longitude, pos.Latitude, material)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := []model.PartnerMatchResult{}

	for rows.Next() {
		p := model.PartnerMatchResult{}
		if err := rows.Scan(&p.Id, &p.Name, &p.Rating, &p.Distance, &p.Position.Longitude, &p.Position.Latitude); err != nil {
			return nil, err
		}
		res = append(res, p)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return res, nil
}

func (d *Driver) GetPartner(id string) (*model.Partner, error) {
	q := `
select
  p.id,
  p.name,
  p.radius,
  p.rating,
  ST_X(p.geo::geometry), ST_Y(p.geo::geometry),
  array_agg(s.code order by s.code)
from partner p
join partner_skill s on s.partner_id = p.id
where p.id = $1
group by p.id`

	rows, err := d.db.Query(q, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}

	var p model.Partner
	if err := rows.Scan(&p.Id, &p.Name, &p.Radius, &p.Rating, &p.Position.Longitude, &p.Position.Latitude, pq.Array(&p.Skills)); err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return &p, nil
}
