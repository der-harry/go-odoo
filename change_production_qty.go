package odoo

// ChangeProductionQty represents change.production.qty model.
type ChangeProductionQty struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	MoId        *Many2One `xmlrpc:"mo_id,omitempty"`
	ProductQty  *Float    `xmlrpc:"product_qty,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ChangeProductionQtys represents array of change.production.qty model.
type ChangeProductionQtys []ChangeProductionQty

// ChangeProductionQtyModel is the odoo model name.
const ChangeProductionQtyModel = "change.production.qty"

// Many2One convert ChangeProductionQty to *Many2One.
func (cpq *ChangeProductionQty) Many2One() *Many2One {
	return NewMany2One(cpq.Id.Get(), "")
}

// CreateChangeProductionQty creates a new change.production.qty model and returns its id.
func (c *Client) CreateChangeProductionQty(cpq *ChangeProductionQty) (int64, error) {
	ids, err := c.CreateChangeProductionQtys([]*ChangeProductionQty{cpq})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateChangeProductionQty creates a new change.production.qty model and returns its id.
func (c *Client) CreateChangeProductionQtys(cpqs []*ChangeProductionQty) ([]int64, error) {
	var vv []interface{}
	for _, v := range cpqs {
		vv = append(vv, v)
	}
	return c.Create(ChangeProductionQtyModel, vv, nil)
}

// UpdateChangeProductionQty updates an existing change.production.qty record.
func (c *Client) UpdateChangeProductionQty(cpq *ChangeProductionQty) error {
	return c.UpdateChangeProductionQtys([]int64{cpq.Id.Get()}, cpq)
}

// UpdateChangeProductionQtys updates existing change.production.qty records.
// All records (represented by ids) will be updated by cpq values.
func (c *Client) UpdateChangeProductionQtys(ids []int64, cpq *ChangeProductionQty) error {
	return c.Update(ChangeProductionQtyModel, ids, cpq, nil)
}

// DeleteChangeProductionQty deletes an existing change.production.qty record.
func (c *Client) DeleteChangeProductionQty(id int64) error {
	return c.DeleteChangeProductionQtys([]int64{id})
}

// DeleteChangeProductionQtys deletes existing change.production.qty records.
func (c *Client) DeleteChangeProductionQtys(ids []int64) error {
	return c.Delete(ChangeProductionQtyModel, ids)
}

// GetChangeProductionQty gets change.production.qty existing record.
func (c *Client) GetChangeProductionQty(id int64) (*ChangeProductionQty, error) {
	cpqs, err := c.GetChangeProductionQtys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cpqs)[0]), nil
}

// GetChangeProductionQtys gets change.production.qty existing records.
func (c *Client) GetChangeProductionQtys(ids []int64) (*ChangeProductionQtys, error) {
	cpqs := &ChangeProductionQtys{}
	if err := c.Read(ChangeProductionQtyModel, ids, nil, cpqs); err != nil {
		return nil, err
	}
	return cpqs, nil
}

// FindChangeProductionQty finds change.production.qty record by querying it with criteria.
func (c *Client) FindChangeProductionQty(criteria *Criteria) (*ChangeProductionQty, error) {
	cpqs := &ChangeProductionQtys{}
	if err := c.SearchRead(ChangeProductionQtyModel, criteria, NewOptions().Limit(1), cpqs); err != nil {
		return nil, err
	}
	return &((*cpqs)[0]), nil
}

// FindChangeProductionQtys finds change.production.qty records by querying it
// and filtering it with criteria and options.
func (c *Client) FindChangeProductionQtys(criteria *Criteria, options *Options) (*ChangeProductionQtys, error) {
	cpqs := &ChangeProductionQtys{}
	if err := c.SearchRead(ChangeProductionQtyModel, criteria, options, cpqs); err != nil {
		return nil, err
	}
	return cpqs, nil
}

// FindChangeProductionQtyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindChangeProductionQtyIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ChangeProductionQtyModel, criteria, options)
}

// FindChangeProductionQtyId finds record id by querying it with criteria.
func (c *Client) FindChangeProductionQtyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ChangeProductionQtyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
