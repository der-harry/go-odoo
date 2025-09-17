package odoo

// MrpProductionSplitLine represents mrp.production.split.line model.
type MrpProductionSplitLine struct {
	CreateDate           *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omitempty"`
	Date                 *Time     `xmlrpc:"date,omitempty"`
	DisplayName          *String   `xmlrpc:"display_name,omitempty"`
	Id                   *Int      `xmlrpc:"id,omitempty"`
	MrpProductionSplitId *Many2One `xmlrpc:"mrp_production_split_id,omitempty"`
	Quantity             *Float    `xmlrpc:"quantity,omitempty"`
	UserId               *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MrpProductionSplitLines represents array of mrp.production.split.line model.
type MrpProductionSplitLines []MrpProductionSplitLine

// MrpProductionSplitLineModel is the odoo model name.
const MrpProductionSplitLineModel = "mrp.production.split.line"

// Many2One convert MrpProductionSplitLine to *Many2One.
func (mpsl *MrpProductionSplitLine) Many2One() *Many2One {
	return NewMany2One(mpsl.Id.Get(), "")
}

// CreateMrpProductionSplitLine creates a new mrp.production.split.line model and returns its id.
func (c *Client) CreateMrpProductionSplitLine(mpsl *MrpProductionSplitLine) (int64, error) {
	ids, err := c.CreateMrpProductionSplitLines([]*MrpProductionSplitLine{mpsl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpProductionSplitLine creates a new mrp.production.split.line model and returns its id.
func (c *Client) CreateMrpProductionSplitLines(mpsls []*MrpProductionSplitLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range mpsls {
		vv = append(vv, v)
	}
	return c.Create(MrpProductionSplitLineModel, vv, nil)
}

// UpdateMrpProductionSplitLine updates an existing mrp.production.split.line record.
func (c *Client) UpdateMrpProductionSplitLine(mpsl *MrpProductionSplitLine) error {
	return c.UpdateMrpProductionSplitLines([]int64{mpsl.Id.Get()}, mpsl)
}

// UpdateMrpProductionSplitLines updates existing mrp.production.split.line records.
// All records (represented by ids) will be updated by mpsl values.
func (c *Client) UpdateMrpProductionSplitLines(ids []int64, mpsl *MrpProductionSplitLine) error {
	return c.Update(MrpProductionSplitLineModel, ids, mpsl, nil)
}

// DeleteMrpProductionSplitLine deletes an existing mrp.production.split.line record.
func (c *Client) DeleteMrpProductionSplitLine(id int64) error {
	return c.DeleteMrpProductionSplitLines([]int64{id})
}

// DeleteMrpProductionSplitLines deletes existing mrp.production.split.line records.
func (c *Client) DeleteMrpProductionSplitLines(ids []int64) error {
	return c.Delete(MrpProductionSplitLineModel, ids)
}

// GetMrpProductionSplitLine gets mrp.production.split.line existing record.
func (c *Client) GetMrpProductionSplitLine(id int64) (*MrpProductionSplitLine, error) {
	mpsls, err := c.GetMrpProductionSplitLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mpsls)[0]), nil
}

// GetMrpProductionSplitLines gets mrp.production.split.line existing records.
func (c *Client) GetMrpProductionSplitLines(ids []int64) (*MrpProductionSplitLines, error) {
	mpsls := &MrpProductionSplitLines{}
	if err := c.Read(MrpProductionSplitLineModel, ids, nil, mpsls); err != nil {
		return nil, err
	}
	return mpsls, nil
}

// FindMrpProductionSplitLine finds mrp.production.split.line record by querying it with criteria.
func (c *Client) FindMrpProductionSplitLine(criteria *Criteria) (*MrpProductionSplitLine, error) {
	mpsls := &MrpProductionSplitLines{}
	if err := c.SearchRead(MrpProductionSplitLineModel, criteria, NewOptions().Limit(1), mpsls); err != nil {
		return nil, err
	}
	return &((*mpsls)[0]), nil
}

// FindMrpProductionSplitLines finds mrp.production.split.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpProductionSplitLines(criteria *Criteria, options *Options) (*MrpProductionSplitLines, error) {
	mpsls := &MrpProductionSplitLines{}
	if err := c.SearchRead(MrpProductionSplitLineModel, criteria, options, mpsls); err != nil {
		return nil, err
	}
	return mpsls, nil
}

// FindMrpProductionSplitLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpProductionSplitLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpProductionSplitLineModel, criteria, options)
}

// FindMrpProductionSplitLineId finds record id by querying it with criteria.
func (c *Client) FindMrpProductionSplitLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpProductionSplitLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
