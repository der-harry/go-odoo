package odoo

// MrpProductionSplitMulti represents mrp.production.split.multi model.
type MrpProductionSplitMulti struct {
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	ProductionIds *Relation `xmlrpc:"production_ids,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MrpProductionSplitMultis represents array of mrp.production.split.multi model.
type MrpProductionSplitMultis []MrpProductionSplitMulti

// MrpProductionSplitMultiModel is the odoo model name.
const MrpProductionSplitMultiModel = "mrp.production.split.multi"

// Many2One convert MrpProductionSplitMulti to *Many2One.
func (mpsm *MrpProductionSplitMulti) Many2One() *Many2One {
	return NewMany2One(mpsm.Id.Get(), "")
}

// CreateMrpProductionSplitMulti creates a new mrp.production.split.multi model and returns its id.
func (c *Client) CreateMrpProductionSplitMulti(mpsm *MrpProductionSplitMulti) (int64, error) {
	ids, err := c.CreateMrpProductionSplitMultis([]*MrpProductionSplitMulti{mpsm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpProductionSplitMulti creates a new mrp.production.split.multi model and returns its id.
func (c *Client) CreateMrpProductionSplitMultis(mpsms []*MrpProductionSplitMulti) ([]int64, error) {
	var vv []interface{}
	for _, v := range mpsms {
		vv = append(vv, v)
	}
	return c.Create(MrpProductionSplitMultiModel, vv, nil)
}

// UpdateMrpProductionSplitMulti updates an existing mrp.production.split.multi record.
func (c *Client) UpdateMrpProductionSplitMulti(mpsm *MrpProductionSplitMulti) error {
	return c.UpdateMrpProductionSplitMultis([]int64{mpsm.Id.Get()}, mpsm)
}

// UpdateMrpProductionSplitMultis updates existing mrp.production.split.multi records.
// All records (represented by ids) will be updated by mpsm values.
func (c *Client) UpdateMrpProductionSplitMultis(ids []int64, mpsm *MrpProductionSplitMulti) error {
	return c.Update(MrpProductionSplitMultiModel, ids, mpsm, nil)
}

// DeleteMrpProductionSplitMulti deletes an existing mrp.production.split.multi record.
func (c *Client) DeleteMrpProductionSplitMulti(id int64) error {
	return c.DeleteMrpProductionSplitMultis([]int64{id})
}

// DeleteMrpProductionSplitMultis deletes existing mrp.production.split.multi records.
func (c *Client) DeleteMrpProductionSplitMultis(ids []int64) error {
	return c.Delete(MrpProductionSplitMultiModel, ids)
}

// GetMrpProductionSplitMulti gets mrp.production.split.multi existing record.
func (c *Client) GetMrpProductionSplitMulti(id int64) (*MrpProductionSplitMulti, error) {
	mpsms, err := c.GetMrpProductionSplitMultis([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mpsms)[0]), nil
}

// GetMrpProductionSplitMultis gets mrp.production.split.multi existing records.
func (c *Client) GetMrpProductionSplitMultis(ids []int64) (*MrpProductionSplitMultis, error) {
	mpsms := &MrpProductionSplitMultis{}
	if err := c.Read(MrpProductionSplitMultiModel, ids, nil, mpsms); err != nil {
		return nil, err
	}
	return mpsms, nil
}

// FindMrpProductionSplitMulti finds mrp.production.split.multi record by querying it with criteria.
func (c *Client) FindMrpProductionSplitMulti(criteria *Criteria) (*MrpProductionSplitMulti, error) {
	mpsms := &MrpProductionSplitMultis{}
	if err := c.SearchRead(MrpProductionSplitMultiModel, criteria, NewOptions().Limit(1), mpsms); err != nil {
		return nil, err
	}
	return &((*mpsms)[0]), nil
}

// FindMrpProductionSplitMultis finds mrp.production.split.multi records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpProductionSplitMultis(criteria *Criteria, options *Options) (*MrpProductionSplitMultis, error) {
	mpsms := &MrpProductionSplitMultis{}
	if err := c.SearchRead(MrpProductionSplitMultiModel, criteria, options, mpsms); err != nil {
		return nil, err
	}
	return mpsms, nil
}

// FindMrpProductionSplitMultiIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpProductionSplitMultiIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpProductionSplitMultiModel, criteria, options)
}

// FindMrpProductionSplitMultiId finds record id by querying it with criteria.
func (c *Client) FindMrpProductionSplitMultiId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpProductionSplitMultiModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
