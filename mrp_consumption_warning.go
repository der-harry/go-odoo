package odoo

// MrpConsumptionWarning represents mrp.consumption.warning model.
type MrpConsumptionWarning struct {
	Consumption                  *Selection `xmlrpc:"consumption,omitempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omitempty"`
	Id                           *Int       `xmlrpc:"id,omitempty"`
	MrpConsumptionWarningLineIds *Relation  `xmlrpc:"mrp_consumption_warning_line_ids,omitempty"`
	MrpProductionCount           *Int       `xmlrpc:"mrp_production_count,omitempty"`
	MrpProductionIds             *Relation  `xmlrpc:"mrp_production_ids,omitempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MrpConsumptionWarnings represents array of mrp.consumption.warning model.
type MrpConsumptionWarnings []MrpConsumptionWarning

// MrpConsumptionWarningModel is the odoo model name.
const MrpConsumptionWarningModel = "mrp.consumption.warning"

// Many2One convert MrpConsumptionWarning to *Many2One.
func (mcw *MrpConsumptionWarning) Many2One() *Many2One {
	return NewMany2One(mcw.Id.Get(), "")
}

// CreateMrpConsumptionWarning creates a new mrp.consumption.warning model and returns its id.
func (c *Client) CreateMrpConsumptionWarning(mcw *MrpConsumptionWarning) (int64, error) {
	ids, err := c.CreateMrpConsumptionWarnings([]*MrpConsumptionWarning{mcw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpConsumptionWarning creates a new mrp.consumption.warning model and returns its id.
func (c *Client) CreateMrpConsumptionWarnings(mcws []*MrpConsumptionWarning) ([]int64, error) {
	var vv []interface{}
	for _, v := range mcws {
		vv = append(vv, v)
	}
	return c.Create(MrpConsumptionWarningModel, vv, nil)
}

// UpdateMrpConsumptionWarning updates an existing mrp.consumption.warning record.
func (c *Client) UpdateMrpConsumptionWarning(mcw *MrpConsumptionWarning) error {
	return c.UpdateMrpConsumptionWarnings([]int64{mcw.Id.Get()}, mcw)
}

// UpdateMrpConsumptionWarnings updates existing mrp.consumption.warning records.
// All records (represented by ids) will be updated by mcw values.
func (c *Client) UpdateMrpConsumptionWarnings(ids []int64, mcw *MrpConsumptionWarning) error {
	return c.Update(MrpConsumptionWarningModel, ids, mcw, nil)
}

// DeleteMrpConsumptionWarning deletes an existing mrp.consumption.warning record.
func (c *Client) DeleteMrpConsumptionWarning(id int64) error {
	return c.DeleteMrpConsumptionWarnings([]int64{id})
}

// DeleteMrpConsumptionWarnings deletes existing mrp.consumption.warning records.
func (c *Client) DeleteMrpConsumptionWarnings(ids []int64) error {
	return c.Delete(MrpConsumptionWarningModel, ids)
}

// GetMrpConsumptionWarning gets mrp.consumption.warning existing record.
func (c *Client) GetMrpConsumptionWarning(id int64) (*MrpConsumptionWarning, error) {
	mcws, err := c.GetMrpConsumptionWarnings([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mcws)[0]), nil
}

// GetMrpConsumptionWarnings gets mrp.consumption.warning existing records.
func (c *Client) GetMrpConsumptionWarnings(ids []int64) (*MrpConsumptionWarnings, error) {
	mcws := &MrpConsumptionWarnings{}
	if err := c.Read(MrpConsumptionWarningModel, ids, nil, mcws); err != nil {
		return nil, err
	}
	return mcws, nil
}

// FindMrpConsumptionWarning finds mrp.consumption.warning record by querying it with criteria.
func (c *Client) FindMrpConsumptionWarning(criteria *Criteria) (*MrpConsumptionWarning, error) {
	mcws := &MrpConsumptionWarnings{}
	if err := c.SearchRead(MrpConsumptionWarningModel, criteria, NewOptions().Limit(1), mcws); err != nil {
		return nil, err
	}
	return &((*mcws)[0]), nil
}

// FindMrpConsumptionWarnings finds mrp.consumption.warning records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpConsumptionWarnings(criteria *Criteria, options *Options) (*MrpConsumptionWarnings, error) {
	mcws := &MrpConsumptionWarnings{}
	if err := c.SearchRead(MrpConsumptionWarningModel, criteria, options, mcws); err != nil {
		return nil, err
	}
	return mcws, nil
}

// FindMrpConsumptionWarningIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpConsumptionWarningIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpConsumptionWarningModel, criteria, options)
}

// FindMrpConsumptionWarningId finds record id by querying it with criteria.
func (c *Client) FindMrpConsumptionWarningId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpConsumptionWarningModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
