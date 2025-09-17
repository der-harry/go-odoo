package odoo

// MrpWorkcenterProductivity represents mrp.workcenter.productivity model.
type MrpWorkcenterProductivity struct {
	AccountMoveLineId *Many2One  `xmlrpc:"account_move_line_id,omitempty"`
	CompanyId         *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateEnd           *Time      `xmlrpc:"date_end,omitempty"`
	DateStart         *Time      `xmlrpc:"date_start,omitempty"`
	Description       *String    `xmlrpc:"description,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	Duration          *Float     `xmlrpc:"duration,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	LossId            *Many2One  `xmlrpc:"loss_id,omitempty"`
	LossType          *Selection `xmlrpc:"loss_type,omitempty"`
	ProductionId      *Many2One  `xmlrpc:"production_id,omitempty"`
	UserId            *Many2One  `xmlrpc:"user_id,omitempty"`
	WorkcenterId      *Many2One  `xmlrpc:"workcenter_id,omitempty"`
	WorkorderId       *Many2One  `xmlrpc:"workorder_id,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MrpWorkcenterProductivitys represents array of mrp.workcenter.productivity model.
type MrpWorkcenterProductivitys []MrpWorkcenterProductivity

// MrpWorkcenterProductivityModel is the odoo model name.
const MrpWorkcenterProductivityModel = "mrp.workcenter.productivity"

// Many2One convert MrpWorkcenterProductivity to *Many2One.
func (mwp *MrpWorkcenterProductivity) Many2One() *Many2One {
	return NewMany2One(mwp.Id.Get(), "")
}

// CreateMrpWorkcenterProductivity creates a new mrp.workcenter.productivity model and returns its id.
func (c *Client) CreateMrpWorkcenterProductivity(mwp *MrpWorkcenterProductivity) (int64, error) {
	ids, err := c.CreateMrpWorkcenterProductivitys([]*MrpWorkcenterProductivity{mwp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpWorkcenterProductivity creates a new mrp.workcenter.productivity model and returns its id.
func (c *Client) CreateMrpWorkcenterProductivitys(mwps []*MrpWorkcenterProductivity) ([]int64, error) {
	var vv []interface{}
	for _, v := range mwps {
		vv = append(vv, v)
	}
	return c.Create(MrpWorkcenterProductivityModel, vv, nil)
}

// UpdateMrpWorkcenterProductivity updates an existing mrp.workcenter.productivity record.
func (c *Client) UpdateMrpWorkcenterProductivity(mwp *MrpWorkcenterProductivity) error {
	return c.UpdateMrpWorkcenterProductivitys([]int64{mwp.Id.Get()}, mwp)
}

// UpdateMrpWorkcenterProductivitys updates existing mrp.workcenter.productivity records.
// All records (represented by ids) will be updated by mwp values.
func (c *Client) UpdateMrpWorkcenterProductivitys(ids []int64, mwp *MrpWorkcenterProductivity) error {
	return c.Update(MrpWorkcenterProductivityModel, ids, mwp, nil)
}

// DeleteMrpWorkcenterProductivity deletes an existing mrp.workcenter.productivity record.
func (c *Client) DeleteMrpWorkcenterProductivity(id int64) error {
	return c.DeleteMrpWorkcenterProductivitys([]int64{id})
}

// DeleteMrpWorkcenterProductivitys deletes existing mrp.workcenter.productivity records.
func (c *Client) DeleteMrpWorkcenterProductivitys(ids []int64) error {
	return c.Delete(MrpWorkcenterProductivityModel, ids)
}

// GetMrpWorkcenterProductivity gets mrp.workcenter.productivity existing record.
func (c *Client) GetMrpWorkcenterProductivity(id int64) (*MrpWorkcenterProductivity, error) {
	mwps, err := c.GetMrpWorkcenterProductivitys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mwps)[0]), nil
}

// GetMrpWorkcenterProductivitys gets mrp.workcenter.productivity existing records.
func (c *Client) GetMrpWorkcenterProductivitys(ids []int64) (*MrpWorkcenterProductivitys, error) {
	mwps := &MrpWorkcenterProductivitys{}
	if err := c.Read(MrpWorkcenterProductivityModel, ids, nil, mwps); err != nil {
		return nil, err
	}
	return mwps, nil
}

// FindMrpWorkcenterProductivity finds mrp.workcenter.productivity record by querying it with criteria.
func (c *Client) FindMrpWorkcenterProductivity(criteria *Criteria) (*MrpWorkcenterProductivity, error) {
	mwps := &MrpWorkcenterProductivitys{}
	if err := c.SearchRead(MrpWorkcenterProductivityModel, criteria, NewOptions().Limit(1), mwps); err != nil {
		return nil, err
	}
	return &((*mwps)[0]), nil
}

// FindMrpWorkcenterProductivitys finds mrp.workcenter.productivity records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpWorkcenterProductivitys(criteria *Criteria, options *Options) (*MrpWorkcenterProductivitys, error) {
	mwps := &MrpWorkcenterProductivitys{}
	if err := c.SearchRead(MrpWorkcenterProductivityModel, criteria, options, mwps); err != nil {
		return nil, err
	}
	return mwps, nil
}

// FindMrpWorkcenterProductivityIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpWorkcenterProductivityIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpWorkcenterProductivityModel, criteria, options)
}

// FindMrpWorkcenterProductivityId finds record id by querying it with criteria.
func (c *Client) FindMrpWorkcenterProductivityId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpWorkcenterProductivityModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
