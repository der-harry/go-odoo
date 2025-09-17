package odoo

// MrpAccountWipAccountingLine represents mrp.account.wip.accounting.line model.
type MrpAccountWipAccountingLine struct {
	AccountId       *Many2One `xmlrpc:"account_id,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	Credit          *Float    `xmlrpc:"credit,omitempty"`
	CurrencyId      *Many2One `xmlrpc:"currency_id,omitempty"`
	Debit           *Float    `xmlrpc:"debit,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	Label           *String   `xmlrpc:"label,omitempty"`
	WipAccountingId *Many2One `xmlrpc:"wip_accounting_id,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MrpAccountWipAccountingLines represents array of mrp.account.wip.accounting.line model.
type MrpAccountWipAccountingLines []MrpAccountWipAccountingLine

// MrpAccountWipAccountingLineModel is the odoo model name.
const MrpAccountWipAccountingLineModel = "mrp.account.wip.accounting.line"

// Many2One convert MrpAccountWipAccountingLine to *Many2One.
func (mawal *MrpAccountWipAccountingLine) Many2One() *Many2One {
	return NewMany2One(mawal.Id.Get(), "")
}

// CreateMrpAccountWipAccountingLine creates a new mrp.account.wip.accounting.line model and returns its id.
func (c *Client) CreateMrpAccountWipAccountingLine(mawal *MrpAccountWipAccountingLine) (int64, error) {
	ids, err := c.CreateMrpAccountWipAccountingLines([]*MrpAccountWipAccountingLine{mawal})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpAccountWipAccountingLine creates a new mrp.account.wip.accounting.line model and returns its id.
func (c *Client) CreateMrpAccountWipAccountingLines(mawals []*MrpAccountWipAccountingLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range mawals {
		vv = append(vv, v)
	}
	return c.Create(MrpAccountWipAccountingLineModel, vv, nil)
}

// UpdateMrpAccountWipAccountingLine updates an existing mrp.account.wip.accounting.line record.
func (c *Client) UpdateMrpAccountWipAccountingLine(mawal *MrpAccountWipAccountingLine) error {
	return c.UpdateMrpAccountWipAccountingLines([]int64{mawal.Id.Get()}, mawal)
}

// UpdateMrpAccountWipAccountingLines updates existing mrp.account.wip.accounting.line records.
// All records (represented by ids) will be updated by mawal values.
func (c *Client) UpdateMrpAccountWipAccountingLines(ids []int64, mawal *MrpAccountWipAccountingLine) error {
	return c.Update(MrpAccountWipAccountingLineModel, ids, mawal, nil)
}

// DeleteMrpAccountWipAccountingLine deletes an existing mrp.account.wip.accounting.line record.
func (c *Client) DeleteMrpAccountWipAccountingLine(id int64) error {
	return c.DeleteMrpAccountWipAccountingLines([]int64{id})
}

// DeleteMrpAccountWipAccountingLines deletes existing mrp.account.wip.accounting.line records.
func (c *Client) DeleteMrpAccountWipAccountingLines(ids []int64) error {
	return c.Delete(MrpAccountWipAccountingLineModel, ids)
}

// GetMrpAccountWipAccountingLine gets mrp.account.wip.accounting.line existing record.
func (c *Client) GetMrpAccountWipAccountingLine(id int64) (*MrpAccountWipAccountingLine, error) {
	mawals, err := c.GetMrpAccountWipAccountingLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mawals)[0]), nil
}

// GetMrpAccountWipAccountingLines gets mrp.account.wip.accounting.line existing records.
func (c *Client) GetMrpAccountWipAccountingLines(ids []int64) (*MrpAccountWipAccountingLines, error) {
	mawals := &MrpAccountWipAccountingLines{}
	if err := c.Read(MrpAccountWipAccountingLineModel, ids, nil, mawals); err != nil {
		return nil, err
	}
	return mawals, nil
}

// FindMrpAccountWipAccountingLine finds mrp.account.wip.accounting.line record by querying it with criteria.
func (c *Client) FindMrpAccountWipAccountingLine(criteria *Criteria) (*MrpAccountWipAccountingLine, error) {
	mawals := &MrpAccountWipAccountingLines{}
	if err := c.SearchRead(MrpAccountWipAccountingLineModel, criteria, NewOptions().Limit(1), mawals); err != nil {
		return nil, err
	}
	return &((*mawals)[0]), nil
}

// FindMrpAccountWipAccountingLines finds mrp.account.wip.accounting.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpAccountWipAccountingLines(criteria *Criteria, options *Options) (*MrpAccountWipAccountingLines, error) {
	mawals := &MrpAccountWipAccountingLines{}
	if err := c.SearchRead(MrpAccountWipAccountingLineModel, criteria, options, mawals); err != nil {
		return nil, err
	}
	return mawals, nil
}

// FindMrpAccountWipAccountingLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpAccountWipAccountingLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpAccountWipAccountingLineModel, criteria, options)
}

// FindMrpAccountWipAccountingLineId finds record id by querying it with criteria.
func (c *Client) FindMrpAccountWipAccountingLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpAccountWipAccountingLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
