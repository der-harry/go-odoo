package odoo

// MrpWorkcenterProductivityLoss represents mrp.workcenter.productivity.loss model.
type MrpWorkcenterProductivityLoss struct {
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	LossId      *Many2One  `xmlrpc:"loss_id,omitempty"`
	LossType    *Selection `xmlrpc:"loss_type,omitempty"`
	Manual      *Bool      `xmlrpc:"manual,omitempty"`
	Name        *String    `xmlrpc:"name,omitempty"`
	Sequence    *Int       `xmlrpc:"sequence,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MrpWorkcenterProductivityLosss represents array of mrp.workcenter.productivity.loss model.
type MrpWorkcenterProductivityLosss []MrpWorkcenterProductivityLoss

// MrpWorkcenterProductivityLossModel is the odoo model name.
const MrpWorkcenterProductivityLossModel = "mrp.workcenter.productivity.loss"

// Many2One convert MrpWorkcenterProductivityLoss to *Many2One.
func (mwpl *MrpWorkcenterProductivityLoss) Many2One() *Many2One {
	return NewMany2One(mwpl.Id.Get(), "")
}

// CreateMrpWorkcenterProductivityLoss creates a new mrp.workcenter.productivity.loss model and returns its id.
func (c *Client) CreateMrpWorkcenterProductivityLoss(mwpl *MrpWorkcenterProductivityLoss) (int64, error) {
	ids, err := c.CreateMrpWorkcenterProductivityLosss([]*MrpWorkcenterProductivityLoss{mwpl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpWorkcenterProductivityLoss creates a new mrp.workcenter.productivity.loss model and returns its id.
func (c *Client) CreateMrpWorkcenterProductivityLosss(mwpls []*MrpWorkcenterProductivityLoss) ([]int64, error) {
	var vv []interface{}
	for _, v := range mwpls {
		vv = append(vv, v)
	}
	return c.Create(MrpWorkcenterProductivityLossModel, vv, nil)
}

// UpdateMrpWorkcenterProductivityLoss updates an existing mrp.workcenter.productivity.loss record.
func (c *Client) UpdateMrpWorkcenterProductivityLoss(mwpl *MrpWorkcenterProductivityLoss) error {
	return c.UpdateMrpWorkcenterProductivityLosss([]int64{mwpl.Id.Get()}, mwpl)
}

// UpdateMrpWorkcenterProductivityLosss updates existing mrp.workcenter.productivity.loss records.
// All records (represented by ids) will be updated by mwpl values.
func (c *Client) UpdateMrpWorkcenterProductivityLosss(ids []int64, mwpl *MrpWorkcenterProductivityLoss) error {
	return c.Update(MrpWorkcenterProductivityLossModel, ids, mwpl, nil)
}

// DeleteMrpWorkcenterProductivityLoss deletes an existing mrp.workcenter.productivity.loss record.
func (c *Client) DeleteMrpWorkcenterProductivityLoss(id int64) error {
	return c.DeleteMrpWorkcenterProductivityLosss([]int64{id})
}

// DeleteMrpWorkcenterProductivityLosss deletes existing mrp.workcenter.productivity.loss records.
func (c *Client) DeleteMrpWorkcenterProductivityLosss(ids []int64) error {
	return c.Delete(MrpWorkcenterProductivityLossModel, ids)
}

// GetMrpWorkcenterProductivityLoss gets mrp.workcenter.productivity.loss existing record.
func (c *Client) GetMrpWorkcenterProductivityLoss(id int64) (*MrpWorkcenterProductivityLoss, error) {
	mwpls, err := c.GetMrpWorkcenterProductivityLosss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mwpls)[0]), nil
}

// GetMrpWorkcenterProductivityLosss gets mrp.workcenter.productivity.loss existing records.
func (c *Client) GetMrpWorkcenterProductivityLosss(ids []int64) (*MrpWorkcenterProductivityLosss, error) {
	mwpls := &MrpWorkcenterProductivityLosss{}
	if err := c.Read(MrpWorkcenterProductivityLossModel, ids, nil, mwpls); err != nil {
		return nil, err
	}
	return mwpls, nil
}

// FindMrpWorkcenterProductivityLoss finds mrp.workcenter.productivity.loss record by querying it with criteria.
func (c *Client) FindMrpWorkcenterProductivityLoss(criteria *Criteria) (*MrpWorkcenterProductivityLoss, error) {
	mwpls := &MrpWorkcenterProductivityLosss{}
	if err := c.SearchRead(MrpWorkcenterProductivityLossModel, criteria, NewOptions().Limit(1), mwpls); err != nil {
		return nil, err
	}
	return &((*mwpls)[0]), nil
}

// FindMrpWorkcenterProductivityLosss finds mrp.workcenter.productivity.loss records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpWorkcenterProductivityLosss(criteria *Criteria, options *Options) (*MrpWorkcenterProductivityLosss, error) {
	mwpls := &MrpWorkcenterProductivityLosss{}
	if err := c.SearchRead(MrpWorkcenterProductivityLossModel, criteria, options, mwpls); err != nil {
		return nil, err
	}
	return mwpls, nil
}

// FindMrpWorkcenterProductivityLossIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpWorkcenterProductivityLossIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpWorkcenterProductivityLossModel, criteria, options)
}

// FindMrpWorkcenterProductivityLossId finds record id by querying it with criteria.
func (c *Client) FindMrpWorkcenterProductivityLossId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpWorkcenterProductivityLossModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
