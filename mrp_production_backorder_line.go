package odoo

// MrpProductionBackorderLine represents mrp.production.backorder.line model.
type MrpProductionBackorderLine struct {
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
	MrpProductionBackorderId *Many2One `xmlrpc:"mrp_production_backorder_id,omitempty"`
	MrpProductionId          *Many2One `xmlrpc:"mrp_production_id,omitempty"`
	ToBackorder              *Bool     `xmlrpc:"to_backorder,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MrpProductionBackorderLines represents array of mrp.production.backorder.line model.
type MrpProductionBackorderLines []MrpProductionBackorderLine

// MrpProductionBackorderLineModel is the odoo model name.
const MrpProductionBackorderLineModel = "mrp.production.backorder.line"

// Many2One convert MrpProductionBackorderLine to *Many2One.
func (mpbl *MrpProductionBackorderLine) Many2One() *Many2One {
	return NewMany2One(mpbl.Id.Get(), "")
}

// CreateMrpProductionBackorderLine creates a new mrp.production.backorder.line model and returns its id.
func (c *Client) CreateMrpProductionBackorderLine(mpbl *MrpProductionBackorderLine) (int64, error) {
	ids, err := c.CreateMrpProductionBackorderLines([]*MrpProductionBackorderLine{mpbl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpProductionBackorderLine creates a new mrp.production.backorder.line model and returns its id.
func (c *Client) CreateMrpProductionBackorderLines(mpbls []*MrpProductionBackorderLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range mpbls {
		vv = append(vv, v)
	}
	return c.Create(MrpProductionBackorderLineModel, vv, nil)
}

// UpdateMrpProductionBackorderLine updates an existing mrp.production.backorder.line record.
func (c *Client) UpdateMrpProductionBackorderLine(mpbl *MrpProductionBackorderLine) error {
	return c.UpdateMrpProductionBackorderLines([]int64{mpbl.Id.Get()}, mpbl)
}

// UpdateMrpProductionBackorderLines updates existing mrp.production.backorder.line records.
// All records (represented by ids) will be updated by mpbl values.
func (c *Client) UpdateMrpProductionBackorderLines(ids []int64, mpbl *MrpProductionBackorderLine) error {
	return c.Update(MrpProductionBackorderLineModel, ids, mpbl, nil)
}

// DeleteMrpProductionBackorderLine deletes an existing mrp.production.backorder.line record.
func (c *Client) DeleteMrpProductionBackorderLine(id int64) error {
	return c.DeleteMrpProductionBackorderLines([]int64{id})
}

// DeleteMrpProductionBackorderLines deletes existing mrp.production.backorder.line records.
func (c *Client) DeleteMrpProductionBackorderLines(ids []int64) error {
	return c.Delete(MrpProductionBackorderLineModel, ids)
}

// GetMrpProductionBackorderLine gets mrp.production.backorder.line existing record.
func (c *Client) GetMrpProductionBackorderLine(id int64) (*MrpProductionBackorderLine, error) {
	mpbls, err := c.GetMrpProductionBackorderLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mpbls)[0]), nil
}

// GetMrpProductionBackorderLines gets mrp.production.backorder.line existing records.
func (c *Client) GetMrpProductionBackorderLines(ids []int64) (*MrpProductionBackorderLines, error) {
	mpbls := &MrpProductionBackorderLines{}
	if err := c.Read(MrpProductionBackorderLineModel, ids, nil, mpbls); err != nil {
		return nil, err
	}
	return mpbls, nil
}

// FindMrpProductionBackorderLine finds mrp.production.backorder.line record by querying it with criteria.
func (c *Client) FindMrpProductionBackorderLine(criteria *Criteria) (*MrpProductionBackorderLine, error) {
	mpbls := &MrpProductionBackorderLines{}
	if err := c.SearchRead(MrpProductionBackorderLineModel, criteria, NewOptions().Limit(1), mpbls); err != nil {
		return nil, err
	}
	return &((*mpbls)[0]), nil
}

// FindMrpProductionBackorderLines finds mrp.production.backorder.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpProductionBackorderLines(criteria *Criteria, options *Options) (*MrpProductionBackorderLines, error) {
	mpbls := &MrpProductionBackorderLines{}
	if err := c.SearchRead(MrpProductionBackorderLineModel, criteria, options, mpbls); err != nil {
		return nil, err
	}
	return mpbls, nil
}

// FindMrpProductionBackorderLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpProductionBackorderLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpProductionBackorderLineModel, criteria, options)
}

// FindMrpProductionBackorderLineId finds record id by querying it with criteria.
func (c *Client) FindMrpProductionBackorderLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpProductionBackorderLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
