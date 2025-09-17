package odoo

// MrpProductionBackorder represents mrp.production.backorder model.
type MrpProductionBackorder struct {
	CreateDate                    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName                   *String   `xmlrpc:"display_name,omitempty"`
	Id                            *Int      `xmlrpc:"id,omitempty"`
	MrpProductionBackorderLineIds *Relation `xmlrpc:"mrp_production_backorder_line_ids,omitempty"`
	MrpProductionIds              *Relation `xmlrpc:"mrp_production_ids,omitempty"`
	ShowBackorderLines            *Bool     `xmlrpc:"show_backorder_lines,omitempty"`
	WriteDate                     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MrpProductionBackorders represents array of mrp.production.backorder model.
type MrpProductionBackorders []MrpProductionBackorder

// MrpProductionBackorderModel is the odoo model name.
const MrpProductionBackorderModel = "mrp.production.backorder"

// Many2One convert MrpProductionBackorder to *Many2One.
func (mpb *MrpProductionBackorder) Many2One() *Many2One {
	return NewMany2One(mpb.Id.Get(), "")
}

// CreateMrpProductionBackorder creates a new mrp.production.backorder model and returns its id.
func (c *Client) CreateMrpProductionBackorder(mpb *MrpProductionBackorder) (int64, error) {
	ids, err := c.CreateMrpProductionBackorders([]*MrpProductionBackorder{mpb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpProductionBackorder creates a new mrp.production.backorder model and returns its id.
func (c *Client) CreateMrpProductionBackorders(mpbs []*MrpProductionBackorder) ([]int64, error) {
	var vv []interface{}
	for _, v := range mpbs {
		vv = append(vv, v)
	}
	return c.Create(MrpProductionBackorderModel, vv, nil)
}

// UpdateMrpProductionBackorder updates an existing mrp.production.backorder record.
func (c *Client) UpdateMrpProductionBackorder(mpb *MrpProductionBackorder) error {
	return c.UpdateMrpProductionBackorders([]int64{mpb.Id.Get()}, mpb)
}

// UpdateMrpProductionBackorders updates existing mrp.production.backorder records.
// All records (represented by ids) will be updated by mpb values.
func (c *Client) UpdateMrpProductionBackorders(ids []int64, mpb *MrpProductionBackorder) error {
	return c.Update(MrpProductionBackorderModel, ids, mpb, nil)
}

// DeleteMrpProductionBackorder deletes an existing mrp.production.backorder record.
func (c *Client) DeleteMrpProductionBackorder(id int64) error {
	return c.DeleteMrpProductionBackorders([]int64{id})
}

// DeleteMrpProductionBackorders deletes existing mrp.production.backorder records.
func (c *Client) DeleteMrpProductionBackorders(ids []int64) error {
	return c.Delete(MrpProductionBackorderModel, ids)
}

// GetMrpProductionBackorder gets mrp.production.backorder existing record.
func (c *Client) GetMrpProductionBackorder(id int64) (*MrpProductionBackorder, error) {
	mpbs, err := c.GetMrpProductionBackorders([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mpbs)[0]), nil
}

// GetMrpProductionBackorders gets mrp.production.backorder existing records.
func (c *Client) GetMrpProductionBackorders(ids []int64) (*MrpProductionBackorders, error) {
	mpbs := &MrpProductionBackorders{}
	if err := c.Read(MrpProductionBackorderModel, ids, nil, mpbs); err != nil {
		return nil, err
	}
	return mpbs, nil
}

// FindMrpProductionBackorder finds mrp.production.backorder record by querying it with criteria.
func (c *Client) FindMrpProductionBackorder(criteria *Criteria) (*MrpProductionBackorder, error) {
	mpbs := &MrpProductionBackorders{}
	if err := c.SearchRead(MrpProductionBackorderModel, criteria, NewOptions().Limit(1), mpbs); err != nil {
		return nil, err
	}
	return &((*mpbs)[0]), nil
}

// FindMrpProductionBackorders finds mrp.production.backorder records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpProductionBackorders(criteria *Criteria, options *Options) (*MrpProductionBackorders, error) {
	mpbs := &MrpProductionBackorders{}
	if err := c.SearchRead(MrpProductionBackorderModel, criteria, options, mpbs); err != nil {
		return nil, err
	}
	return mpbs, nil
}

// FindMrpProductionBackorderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpProductionBackorderIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpProductionBackorderModel, criteria, options)
}

// FindMrpProductionBackorderId finds record id by querying it with criteria.
func (c *Client) FindMrpProductionBackorderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpProductionBackorderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
