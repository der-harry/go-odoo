package odoo

// LunchProductCategory represents lunch.product.category model.
type LunchProductCategory struct {
	Active       *Bool     `xmlrpc:"active,omitempty"`
	CompanyId    *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId   *Many2One `xmlrpc:"currency_id,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty"`
	Image1024    *String   `xmlrpc:"image_1024,omitempty"`
	Image128     *String   `xmlrpc:"image_128,omitempty"`
	Image1920    *String   `xmlrpc:"image_1920,omitempty"`
	Image256     *String   `xmlrpc:"image_256,omitempty"`
	Image512     *String   `xmlrpc:"image_512,omitempty"`
	Name         *String   `xmlrpc:"name,omitempty"`
	ProductCount *Int      `xmlrpc:"product_count,omitempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omitempty"`
}

// LunchProductCategorys represents array of lunch.product.category model.
type LunchProductCategorys []LunchProductCategory

// LunchProductCategoryModel is the odoo model name.
const LunchProductCategoryModel = "lunch.product.category"

// Many2One convert LunchProductCategory to *Many2One.
func (lpc *LunchProductCategory) Many2One() *Many2One {
	return NewMany2One(lpc.Id.Get(), "")
}

// CreateLunchProductCategory creates a new lunch.product.category model and returns its id.
func (c *Client) CreateLunchProductCategory(lpc *LunchProductCategory) (int64, error) {
	ids, err := c.CreateLunchProductCategorys([]*LunchProductCategory{lpc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLunchProductCategory creates a new lunch.product.category model and returns its id.
func (c *Client) CreateLunchProductCategorys(lpcs []*LunchProductCategory) ([]int64, error) {
	var vv []interface{}
	for _, v := range lpcs {
		vv = append(vv, v)
	}
	return c.Create(LunchProductCategoryModel, vv, nil)
}

// UpdateLunchProductCategory updates an existing lunch.product.category record.
func (c *Client) UpdateLunchProductCategory(lpc *LunchProductCategory) error {
	return c.UpdateLunchProductCategorys([]int64{lpc.Id.Get()}, lpc)
}

// UpdateLunchProductCategorys updates existing lunch.product.category records.
// All records (represented by ids) will be updated by lpc values.
func (c *Client) UpdateLunchProductCategorys(ids []int64, lpc *LunchProductCategory) error {
	return c.Update(LunchProductCategoryModel, ids, lpc, nil)
}

// DeleteLunchProductCategory deletes an existing lunch.product.category record.
func (c *Client) DeleteLunchProductCategory(id int64) error {
	return c.DeleteLunchProductCategorys([]int64{id})
}

// DeleteLunchProductCategorys deletes existing lunch.product.category records.
func (c *Client) DeleteLunchProductCategorys(ids []int64) error {
	return c.Delete(LunchProductCategoryModel, ids)
}

// GetLunchProductCategory gets lunch.product.category existing record.
func (c *Client) GetLunchProductCategory(id int64) (*LunchProductCategory, error) {
	lpcs, err := c.GetLunchProductCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lpcs)[0]), nil
}

// GetLunchProductCategorys gets lunch.product.category existing records.
func (c *Client) GetLunchProductCategorys(ids []int64) (*LunchProductCategorys, error) {
	lpcs := &LunchProductCategorys{}
	if err := c.Read(LunchProductCategoryModel, ids, nil, lpcs); err != nil {
		return nil, err
	}
	return lpcs, nil
}

// FindLunchProductCategory finds lunch.product.category record by querying it with criteria.
func (c *Client) FindLunchProductCategory(criteria *Criteria) (*LunchProductCategory, error) {
	lpcs := &LunchProductCategorys{}
	if err := c.SearchRead(LunchProductCategoryModel, criteria, NewOptions().Limit(1), lpcs); err != nil {
		return nil, err
	}
	return &((*lpcs)[0]), nil
}

// FindLunchProductCategorys finds lunch.product.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchProductCategorys(criteria *Criteria, options *Options) (*LunchProductCategorys, error) {
	lpcs := &LunchProductCategorys{}
	if err := c.SearchRead(LunchProductCategoryModel, criteria, options, lpcs); err != nil {
		return nil, err
	}
	return lpcs, nil
}

// FindLunchProductCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchProductCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(LunchProductCategoryModel, criteria, options)
}

// FindLunchProductCategoryId finds record id by querying it with criteria.
func (c *Client) FindLunchProductCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LunchProductCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
