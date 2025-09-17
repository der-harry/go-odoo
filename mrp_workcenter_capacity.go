package odoo

// MrpWorkcenterCapacity represents mrp.workcenter.capacity model.
type MrpWorkcenterCapacity struct {
	Capacity     *Float    `xmlrpc:"capacity,omitempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty"`
	ProductId    *Many2One `xmlrpc:"product_id,omitempty"`
	ProductUomId *Many2One `xmlrpc:"product_uom_id,omitempty"`
	TimeStart    *Float    `xmlrpc:"time_start,omitempty"`
	TimeStop     *Float    `xmlrpc:"time_stop,omitempty"`
	WorkcenterId *Many2One `xmlrpc:"workcenter_id,omitempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MrpWorkcenterCapacitys represents array of mrp.workcenter.capacity model.
type MrpWorkcenterCapacitys []MrpWorkcenterCapacity

// MrpWorkcenterCapacityModel is the odoo model name.
const MrpWorkcenterCapacityModel = "mrp.workcenter.capacity"

// Many2One convert MrpWorkcenterCapacity to *Many2One.
func (mwc *MrpWorkcenterCapacity) Many2One() *Many2One {
	return NewMany2One(mwc.Id.Get(), "")
}

// CreateMrpWorkcenterCapacity creates a new mrp.workcenter.capacity model and returns its id.
func (c *Client) CreateMrpWorkcenterCapacity(mwc *MrpWorkcenterCapacity) (int64, error) {
	ids, err := c.CreateMrpWorkcenterCapacitys([]*MrpWorkcenterCapacity{mwc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpWorkcenterCapacity creates a new mrp.workcenter.capacity model and returns its id.
func (c *Client) CreateMrpWorkcenterCapacitys(mwcs []*MrpWorkcenterCapacity) ([]int64, error) {
	var vv []interface{}
	for _, v := range mwcs {
		vv = append(vv, v)
	}
	return c.Create(MrpWorkcenterCapacityModel, vv, nil)
}

// UpdateMrpWorkcenterCapacity updates an existing mrp.workcenter.capacity record.
func (c *Client) UpdateMrpWorkcenterCapacity(mwc *MrpWorkcenterCapacity) error {
	return c.UpdateMrpWorkcenterCapacitys([]int64{mwc.Id.Get()}, mwc)
}

// UpdateMrpWorkcenterCapacitys updates existing mrp.workcenter.capacity records.
// All records (represented by ids) will be updated by mwc values.
func (c *Client) UpdateMrpWorkcenterCapacitys(ids []int64, mwc *MrpWorkcenterCapacity) error {
	return c.Update(MrpWorkcenterCapacityModel, ids, mwc, nil)
}

// DeleteMrpWorkcenterCapacity deletes an existing mrp.workcenter.capacity record.
func (c *Client) DeleteMrpWorkcenterCapacity(id int64) error {
	return c.DeleteMrpWorkcenterCapacitys([]int64{id})
}

// DeleteMrpWorkcenterCapacitys deletes existing mrp.workcenter.capacity records.
func (c *Client) DeleteMrpWorkcenterCapacitys(ids []int64) error {
	return c.Delete(MrpWorkcenterCapacityModel, ids)
}

// GetMrpWorkcenterCapacity gets mrp.workcenter.capacity existing record.
func (c *Client) GetMrpWorkcenterCapacity(id int64) (*MrpWorkcenterCapacity, error) {
	mwcs, err := c.GetMrpWorkcenterCapacitys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mwcs)[0]), nil
}

// GetMrpWorkcenterCapacitys gets mrp.workcenter.capacity existing records.
func (c *Client) GetMrpWorkcenterCapacitys(ids []int64) (*MrpWorkcenterCapacitys, error) {
	mwcs := &MrpWorkcenterCapacitys{}
	if err := c.Read(MrpWorkcenterCapacityModel, ids, nil, mwcs); err != nil {
		return nil, err
	}
	return mwcs, nil
}

// FindMrpWorkcenterCapacity finds mrp.workcenter.capacity record by querying it with criteria.
func (c *Client) FindMrpWorkcenterCapacity(criteria *Criteria) (*MrpWorkcenterCapacity, error) {
	mwcs := &MrpWorkcenterCapacitys{}
	if err := c.SearchRead(MrpWorkcenterCapacityModel, criteria, NewOptions().Limit(1), mwcs); err != nil {
		return nil, err
	}
	return &((*mwcs)[0]), nil
}

// FindMrpWorkcenterCapacitys finds mrp.workcenter.capacity records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpWorkcenterCapacitys(criteria *Criteria, options *Options) (*MrpWorkcenterCapacitys, error) {
	mwcs := &MrpWorkcenterCapacitys{}
	if err := c.SearchRead(MrpWorkcenterCapacityModel, criteria, options, mwcs); err != nil {
		return nil, err
	}
	return mwcs, nil
}

// FindMrpWorkcenterCapacityIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpWorkcenterCapacityIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpWorkcenterCapacityModel, criteria, options)
}

// FindMrpWorkcenterCapacityId finds record id by querying it with criteria.
func (c *Client) FindMrpWorkcenterCapacityId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpWorkcenterCapacityModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
