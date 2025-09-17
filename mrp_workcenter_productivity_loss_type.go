package odoo

// MrpWorkcenterProductivityLossType represents mrp.workcenter.productivity.loss.type model.
type MrpWorkcenterProductivityLossType struct {
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	LossType    *Selection `xmlrpc:"loss_type,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MrpWorkcenterProductivityLossTypes represents array of mrp.workcenter.productivity.loss.type model.
type MrpWorkcenterProductivityLossTypes []MrpWorkcenterProductivityLossType

// MrpWorkcenterProductivityLossTypeModel is the odoo model name.
const MrpWorkcenterProductivityLossTypeModel = "mrp.workcenter.productivity.loss.type"

// Many2One convert MrpWorkcenterProductivityLossType to *Many2One.
func (mwplt *MrpWorkcenterProductivityLossType) Many2One() *Many2One {
	return NewMany2One(mwplt.Id.Get(), "")
}

// CreateMrpWorkcenterProductivityLossType creates a new mrp.workcenter.productivity.loss.type model and returns its id.
func (c *Client) CreateMrpWorkcenterProductivityLossType(mwplt *MrpWorkcenterProductivityLossType) (int64, error) {
	ids, err := c.CreateMrpWorkcenterProductivityLossTypes([]*MrpWorkcenterProductivityLossType{mwplt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpWorkcenterProductivityLossType creates a new mrp.workcenter.productivity.loss.type model and returns its id.
func (c *Client) CreateMrpWorkcenterProductivityLossTypes(mwplts []*MrpWorkcenterProductivityLossType) ([]int64, error) {
	var vv []interface{}
	for _, v := range mwplts {
		vv = append(vv, v)
	}
	return c.Create(MrpWorkcenterProductivityLossTypeModel, vv, nil)
}

// UpdateMrpWorkcenterProductivityLossType updates an existing mrp.workcenter.productivity.loss.type record.
func (c *Client) UpdateMrpWorkcenterProductivityLossType(mwplt *MrpWorkcenterProductivityLossType) error {
	return c.UpdateMrpWorkcenterProductivityLossTypes([]int64{mwplt.Id.Get()}, mwplt)
}

// UpdateMrpWorkcenterProductivityLossTypes updates existing mrp.workcenter.productivity.loss.type records.
// All records (represented by ids) will be updated by mwplt values.
func (c *Client) UpdateMrpWorkcenterProductivityLossTypes(ids []int64, mwplt *MrpWorkcenterProductivityLossType) error {
	return c.Update(MrpWorkcenterProductivityLossTypeModel, ids, mwplt, nil)
}

// DeleteMrpWorkcenterProductivityLossType deletes an existing mrp.workcenter.productivity.loss.type record.
func (c *Client) DeleteMrpWorkcenterProductivityLossType(id int64) error {
	return c.DeleteMrpWorkcenterProductivityLossTypes([]int64{id})
}

// DeleteMrpWorkcenterProductivityLossTypes deletes existing mrp.workcenter.productivity.loss.type records.
func (c *Client) DeleteMrpWorkcenterProductivityLossTypes(ids []int64) error {
	return c.Delete(MrpWorkcenterProductivityLossTypeModel, ids)
}

// GetMrpWorkcenterProductivityLossType gets mrp.workcenter.productivity.loss.type existing record.
func (c *Client) GetMrpWorkcenterProductivityLossType(id int64) (*MrpWorkcenterProductivityLossType, error) {
	mwplts, err := c.GetMrpWorkcenterProductivityLossTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mwplts)[0]), nil
}

// GetMrpWorkcenterProductivityLossTypes gets mrp.workcenter.productivity.loss.type existing records.
func (c *Client) GetMrpWorkcenterProductivityLossTypes(ids []int64) (*MrpWorkcenterProductivityLossTypes, error) {
	mwplts := &MrpWorkcenterProductivityLossTypes{}
	if err := c.Read(MrpWorkcenterProductivityLossTypeModel, ids, nil, mwplts); err != nil {
		return nil, err
	}
	return mwplts, nil
}

// FindMrpWorkcenterProductivityLossType finds mrp.workcenter.productivity.loss.type record by querying it with criteria.
func (c *Client) FindMrpWorkcenterProductivityLossType(criteria *Criteria) (*MrpWorkcenterProductivityLossType, error) {
	mwplts := &MrpWorkcenterProductivityLossTypes{}
	if err := c.SearchRead(MrpWorkcenterProductivityLossTypeModel, criteria, NewOptions().Limit(1), mwplts); err != nil {
		return nil, err
	}
	return &((*mwplts)[0]), nil
}

// FindMrpWorkcenterProductivityLossTypes finds mrp.workcenter.productivity.loss.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpWorkcenterProductivityLossTypes(criteria *Criteria, options *Options) (*MrpWorkcenterProductivityLossTypes, error) {
	mwplts := &MrpWorkcenterProductivityLossTypes{}
	if err := c.SearchRead(MrpWorkcenterProductivityLossTypeModel, criteria, options, mwplts); err != nil {
		return nil, err
	}
	return mwplts, nil
}

// FindMrpWorkcenterProductivityLossTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpWorkcenterProductivityLossTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpWorkcenterProductivityLossTypeModel, criteria, options)
}

// FindMrpWorkcenterProductivityLossTypeId finds record id by querying it with criteria.
func (c *Client) FindMrpWorkcenterProductivityLossTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpWorkcenterProductivityLossTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
