package odoo

// MrpProductionSplit represents mrp.production.split model.
type MrpProductionSplit struct {
	Counter                   *Int      `xmlrpc:"counter,omitempty"`
	CreateDate                *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName               *String   `xmlrpc:"display_name,omitempty"`
	Id                        *Int      `xmlrpc:"id,omitempty"`
	ProductId                 *Many2One `xmlrpc:"product_id,omitempty"`
	ProductQty                *Float    `xmlrpc:"product_qty,omitempty"`
	ProductUomId              *Many2One `xmlrpc:"product_uom_id,omitempty"`
	ProductionCapacity        *Float    `xmlrpc:"production_capacity,omitempty"`
	ProductionDetailedValsIds *Relation `xmlrpc:"production_detailed_vals_ids,omitempty"`
	ProductionId              *Many2One `xmlrpc:"production_id,omitempty"`
	ProductionSplitMultiId    *Many2One `xmlrpc:"production_split_multi_id,omitempty"`
	ValidDetails              *Bool     `xmlrpc:"valid_details,omitempty"`
	WriteDate                 *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MrpProductionSplits represents array of mrp.production.split model.
type MrpProductionSplits []MrpProductionSplit

// MrpProductionSplitModel is the odoo model name.
const MrpProductionSplitModel = "mrp.production.split"

// Many2One convert MrpProductionSplit to *Many2One.
func (mps *MrpProductionSplit) Many2One() *Many2One {
	return NewMany2One(mps.Id.Get(), "")
}

// CreateMrpProductionSplit creates a new mrp.production.split model and returns its id.
func (c *Client) CreateMrpProductionSplit(mps *MrpProductionSplit) (int64, error) {
	ids, err := c.CreateMrpProductionSplits([]*MrpProductionSplit{mps})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpProductionSplit creates a new mrp.production.split model and returns its id.
func (c *Client) CreateMrpProductionSplits(mpss []*MrpProductionSplit) ([]int64, error) {
	var vv []interface{}
	for _, v := range mpss {
		vv = append(vv, v)
	}
	return c.Create(MrpProductionSplitModel, vv, nil)
}

// UpdateMrpProductionSplit updates an existing mrp.production.split record.
func (c *Client) UpdateMrpProductionSplit(mps *MrpProductionSplit) error {
	return c.UpdateMrpProductionSplits([]int64{mps.Id.Get()}, mps)
}

// UpdateMrpProductionSplits updates existing mrp.production.split records.
// All records (represented by ids) will be updated by mps values.
func (c *Client) UpdateMrpProductionSplits(ids []int64, mps *MrpProductionSplit) error {
	return c.Update(MrpProductionSplitModel, ids, mps, nil)
}

// DeleteMrpProductionSplit deletes an existing mrp.production.split record.
func (c *Client) DeleteMrpProductionSplit(id int64) error {
	return c.DeleteMrpProductionSplits([]int64{id})
}

// DeleteMrpProductionSplits deletes existing mrp.production.split records.
func (c *Client) DeleteMrpProductionSplits(ids []int64) error {
	return c.Delete(MrpProductionSplitModel, ids)
}

// GetMrpProductionSplit gets mrp.production.split existing record.
func (c *Client) GetMrpProductionSplit(id int64) (*MrpProductionSplit, error) {
	mpss, err := c.GetMrpProductionSplits([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mpss)[0]), nil
}

// GetMrpProductionSplits gets mrp.production.split existing records.
func (c *Client) GetMrpProductionSplits(ids []int64) (*MrpProductionSplits, error) {
	mpss := &MrpProductionSplits{}
	if err := c.Read(MrpProductionSplitModel, ids, nil, mpss); err != nil {
		return nil, err
	}
	return mpss, nil
}

// FindMrpProductionSplit finds mrp.production.split record by querying it with criteria.
func (c *Client) FindMrpProductionSplit(criteria *Criteria) (*MrpProductionSplit, error) {
	mpss := &MrpProductionSplits{}
	if err := c.SearchRead(MrpProductionSplitModel, criteria, NewOptions().Limit(1), mpss); err != nil {
		return nil, err
	}
	return &((*mpss)[0]), nil
}

// FindMrpProductionSplits finds mrp.production.split records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpProductionSplits(criteria *Criteria, options *Options) (*MrpProductionSplits, error) {
	mpss := &MrpProductionSplits{}
	if err := c.SearchRead(MrpProductionSplitModel, criteria, options, mpss); err != nil {
		return nil, err
	}
	return mpss, nil
}

// FindMrpProductionSplitIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpProductionSplitIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpProductionSplitModel, criteria, options)
}

// FindMrpProductionSplitId finds record id by querying it with criteria.
func (c *Client) FindMrpProductionSplitId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpProductionSplitModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
