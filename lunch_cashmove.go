package odoo

// LunchCashmove represents lunch.cashmove model.
type LunchCashmove struct {
	Amount      *Float    `xmlrpc:"amount,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId  *Many2One `xmlrpc:"currency_id,omitempty"`
	Date        *Time     `xmlrpc:"date,omitempty"`
	Description *String   `xmlrpc:"description,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	UserId      *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// LunchCashmoves represents array of lunch.cashmove model.
type LunchCashmoves []LunchCashmove

// LunchCashmoveModel is the odoo model name.
const LunchCashmoveModel = "lunch.cashmove"

// Many2One convert LunchCashmove to *Many2One.
func (lc *LunchCashmove) Many2One() *Many2One {
	return NewMany2One(lc.Id.Get(), "")
}

// CreateLunchCashmove creates a new lunch.cashmove model and returns its id.
func (c *Client) CreateLunchCashmove(lc *LunchCashmove) (int64, error) {
	ids, err := c.CreateLunchCashmoves([]*LunchCashmove{lc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLunchCashmove creates a new lunch.cashmove model and returns its id.
func (c *Client) CreateLunchCashmoves(lcs []*LunchCashmove) ([]int64, error) {
	var vv []interface{}
	for _, v := range lcs {
		vv = append(vv, v)
	}
	return c.Create(LunchCashmoveModel, vv, nil)
}

// UpdateLunchCashmove updates an existing lunch.cashmove record.
func (c *Client) UpdateLunchCashmove(lc *LunchCashmove) error {
	return c.UpdateLunchCashmoves([]int64{lc.Id.Get()}, lc)
}

// UpdateLunchCashmoves updates existing lunch.cashmove records.
// All records (represented by ids) will be updated by lc values.
func (c *Client) UpdateLunchCashmoves(ids []int64, lc *LunchCashmove) error {
	return c.Update(LunchCashmoveModel, ids, lc, nil)
}

// DeleteLunchCashmove deletes an existing lunch.cashmove record.
func (c *Client) DeleteLunchCashmove(id int64) error {
	return c.DeleteLunchCashmoves([]int64{id})
}

// DeleteLunchCashmoves deletes existing lunch.cashmove records.
func (c *Client) DeleteLunchCashmoves(ids []int64) error {
	return c.Delete(LunchCashmoveModel, ids)
}

// GetLunchCashmove gets lunch.cashmove existing record.
func (c *Client) GetLunchCashmove(id int64) (*LunchCashmove, error) {
	lcs, err := c.GetLunchCashmoves([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lcs)[0]), nil
}

// GetLunchCashmoves gets lunch.cashmove existing records.
func (c *Client) GetLunchCashmoves(ids []int64) (*LunchCashmoves, error) {
	lcs := &LunchCashmoves{}
	if err := c.Read(LunchCashmoveModel, ids, nil, lcs); err != nil {
		return nil, err
	}
	return lcs, nil
}

// FindLunchCashmove finds lunch.cashmove record by querying it with criteria.
func (c *Client) FindLunchCashmove(criteria *Criteria) (*LunchCashmove, error) {
	lcs := &LunchCashmoves{}
	if err := c.SearchRead(LunchCashmoveModel, criteria, NewOptions().Limit(1), lcs); err != nil {
		return nil, err
	}
	return &((*lcs)[0]), nil
}

// FindLunchCashmoves finds lunch.cashmove records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchCashmoves(criteria *Criteria, options *Options) (*LunchCashmoves, error) {
	lcs := &LunchCashmoves{}
	if err := c.SearchRead(LunchCashmoveModel, criteria, options, lcs); err != nil {
		return nil, err
	}
	return lcs, nil
}

// FindLunchCashmoveIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchCashmoveIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(LunchCashmoveModel, criteria, options)
}

// FindLunchCashmoveId finds record id by querying it with criteria.
func (c *Client) FindLunchCashmoveId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LunchCashmoveModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
