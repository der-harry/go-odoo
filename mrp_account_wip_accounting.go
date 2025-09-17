package odoo

// MrpAccountWipAccounting represents mrp.account.wip.accounting model.
type MrpAccountWipAccounting struct {
	CreateDate   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omitempty"`
	Date         *Time     `xmlrpc:"date,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty"`
	JournalId    *Many2One `xmlrpc:"journal_id,omitempty"`
	LineIds      *Relation `xmlrpc:"line_ids,omitempty"`
	MoIds        *Relation `xmlrpc:"mo_ids,omitempty"`
	Reference    *String   `xmlrpc:"reference,omitempty"`
	ReversalDate *Time     `xmlrpc:"reversal_date,omitempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MrpAccountWipAccountings represents array of mrp.account.wip.accounting model.
type MrpAccountWipAccountings []MrpAccountWipAccounting

// MrpAccountWipAccountingModel is the odoo model name.
const MrpAccountWipAccountingModel = "mrp.account.wip.accounting"

// Many2One convert MrpAccountWipAccounting to *Many2One.
func (mawa *MrpAccountWipAccounting) Many2One() *Many2One {
	return NewMany2One(mawa.Id.Get(), "")
}

// CreateMrpAccountWipAccounting creates a new mrp.account.wip.accounting model and returns its id.
func (c *Client) CreateMrpAccountWipAccounting(mawa *MrpAccountWipAccounting) (int64, error) {
	ids, err := c.CreateMrpAccountWipAccountings([]*MrpAccountWipAccounting{mawa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpAccountWipAccounting creates a new mrp.account.wip.accounting model and returns its id.
func (c *Client) CreateMrpAccountWipAccountings(mawas []*MrpAccountWipAccounting) ([]int64, error) {
	var vv []interface{}
	for _, v := range mawas {
		vv = append(vv, v)
	}
	return c.Create(MrpAccountWipAccountingModel, vv, nil)
}

// UpdateMrpAccountWipAccounting updates an existing mrp.account.wip.accounting record.
func (c *Client) UpdateMrpAccountWipAccounting(mawa *MrpAccountWipAccounting) error {
	return c.UpdateMrpAccountWipAccountings([]int64{mawa.Id.Get()}, mawa)
}

// UpdateMrpAccountWipAccountings updates existing mrp.account.wip.accounting records.
// All records (represented by ids) will be updated by mawa values.
func (c *Client) UpdateMrpAccountWipAccountings(ids []int64, mawa *MrpAccountWipAccounting) error {
	return c.Update(MrpAccountWipAccountingModel, ids, mawa, nil)
}

// DeleteMrpAccountWipAccounting deletes an existing mrp.account.wip.accounting record.
func (c *Client) DeleteMrpAccountWipAccounting(id int64) error {
	return c.DeleteMrpAccountWipAccountings([]int64{id})
}

// DeleteMrpAccountWipAccountings deletes existing mrp.account.wip.accounting records.
func (c *Client) DeleteMrpAccountWipAccountings(ids []int64) error {
	return c.Delete(MrpAccountWipAccountingModel, ids)
}

// GetMrpAccountWipAccounting gets mrp.account.wip.accounting existing record.
func (c *Client) GetMrpAccountWipAccounting(id int64) (*MrpAccountWipAccounting, error) {
	mawas, err := c.GetMrpAccountWipAccountings([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mawas)[0]), nil
}

// GetMrpAccountWipAccountings gets mrp.account.wip.accounting existing records.
func (c *Client) GetMrpAccountWipAccountings(ids []int64) (*MrpAccountWipAccountings, error) {
	mawas := &MrpAccountWipAccountings{}
	if err := c.Read(MrpAccountWipAccountingModel, ids, nil, mawas); err != nil {
		return nil, err
	}
	return mawas, nil
}

// FindMrpAccountWipAccounting finds mrp.account.wip.accounting record by querying it with criteria.
func (c *Client) FindMrpAccountWipAccounting(criteria *Criteria) (*MrpAccountWipAccounting, error) {
	mawas := &MrpAccountWipAccountings{}
	if err := c.SearchRead(MrpAccountWipAccountingModel, criteria, NewOptions().Limit(1), mawas); err != nil {
		return nil, err
	}
	return &((*mawas)[0]), nil
}

// FindMrpAccountWipAccountings finds mrp.account.wip.accounting records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpAccountWipAccountings(criteria *Criteria, options *Options) (*MrpAccountWipAccountings, error) {
	mawas := &MrpAccountWipAccountings{}
	if err := c.SearchRead(MrpAccountWipAccountingModel, criteria, options, mawas); err != nil {
		return nil, err
	}
	return mawas, nil
}

// FindMrpAccountWipAccountingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpAccountWipAccountingIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpAccountWipAccountingModel, criteria, options)
}

// FindMrpAccountWipAccountingId finds record id by querying it with criteria.
func (c *Client) FindMrpAccountWipAccountingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpAccountWipAccountingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
