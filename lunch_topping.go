package odoo

// LunchTopping represents lunch.topping model.
type LunchTopping struct {
	CompanyId       *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId      *Many2One `xmlrpc:"currency_id,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	Name            *String   `xmlrpc:"name,omitempty"`
	Price           *Float    `xmlrpc:"price,omitempty"`
	SupplierId      *Many2One `xmlrpc:"supplier_id,omitempty"`
	ToppingCategory *Int      `xmlrpc:"topping_category,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// LunchToppings represents array of lunch.topping model.
type LunchToppings []LunchTopping

// LunchToppingModel is the odoo model name.
const LunchToppingModel = "lunch.topping"

// Many2One convert LunchTopping to *Many2One.
func (lt *LunchTopping) Many2One() *Many2One {
	return NewMany2One(lt.Id.Get(), "")
}

// CreateLunchTopping creates a new lunch.topping model and returns its id.
func (c *Client) CreateLunchTopping(lt *LunchTopping) (int64, error) {
	ids, err := c.CreateLunchToppings([]*LunchTopping{lt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLunchTopping creates a new lunch.topping model and returns its id.
func (c *Client) CreateLunchToppings(lts []*LunchTopping) ([]int64, error) {
	var vv []interface{}
	for _, v := range lts {
		vv = append(vv, v)
	}
	return c.Create(LunchToppingModel, vv, nil)
}

// UpdateLunchTopping updates an existing lunch.topping record.
func (c *Client) UpdateLunchTopping(lt *LunchTopping) error {
	return c.UpdateLunchToppings([]int64{lt.Id.Get()}, lt)
}

// UpdateLunchToppings updates existing lunch.topping records.
// All records (represented by ids) will be updated by lt values.
func (c *Client) UpdateLunchToppings(ids []int64, lt *LunchTopping) error {
	return c.Update(LunchToppingModel, ids, lt, nil)
}

// DeleteLunchTopping deletes an existing lunch.topping record.
func (c *Client) DeleteLunchTopping(id int64) error {
	return c.DeleteLunchToppings([]int64{id})
}

// DeleteLunchToppings deletes existing lunch.topping records.
func (c *Client) DeleteLunchToppings(ids []int64) error {
	return c.Delete(LunchToppingModel, ids)
}

// GetLunchTopping gets lunch.topping existing record.
func (c *Client) GetLunchTopping(id int64) (*LunchTopping, error) {
	lts, err := c.GetLunchToppings([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lts)[0]), nil
}

// GetLunchToppings gets lunch.topping existing records.
func (c *Client) GetLunchToppings(ids []int64) (*LunchToppings, error) {
	lts := &LunchToppings{}
	if err := c.Read(LunchToppingModel, ids, nil, lts); err != nil {
		return nil, err
	}
	return lts, nil
}

// FindLunchTopping finds lunch.topping record by querying it with criteria.
func (c *Client) FindLunchTopping(criteria *Criteria) (*LunchTopping, error) {
	lts := &LunchToppings{}
	if err := c.SearchRead(LunchToppingModel, criteria, NewOptions().Limit(1), lts); err != nil {
		return nil, err
	}
	return &((*lts)[0]), nil
}

// FindLunchToppings finds lunch.topping records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchToppings(criteria *Criteria, options *Options) (*LunchToppings, error) {
	lts := &LunchToppings{}
	if err := c.SearchRead(LunchToppingModel, criteria, options, lts); err != nil {
		return nil, err
	}
	return lts, nil
}

// FindLunchToppingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchToppingIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(LunchToppingModel, criteria, options)
}

// FindLunchToppingId finds record id by querying it with criteria.
func (c *Client) FindLunchToppingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LunchToppingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
