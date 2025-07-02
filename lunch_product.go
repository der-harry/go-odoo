package odoo

// LunchProduct represents lunch.product model.
type LunchProduct struct {
	Active          *Bool     `xmlrpc:"active,omitempty"`
	CategoryId      *Many2One `xmlrpc:"category_id,omitempty"`
	CompanyId       *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId      *Many2One `xmlrpc:"currency_id,omitempty"`
	Description     *String   `xmlrpc:"description,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	FavoriteUserIds *Relation `xmlrpc:"favorite_user_ids,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	Image1024       *String   `xmlrpc:"image_1024,omitempty"`
	Image128        *String   `xmlrpc:"image_128,omitempty"`
	Image1920       *String   `xmlrpc:"image_1920,omitempty"`
	Image256        *String   `xmlrpc:"image_256,omitempty"`
	Image512        *String   `xmlrpc:"image_512,omitempty"`
	IsAvailableAt   *Many2One `xmlrpc:"is_available_at,omitempty"`
	IsFavorite      *Bool     `xmlrpc:"is_favorite,omitempty"`
	IsNew           *Bool     `xmlrpc:"is_new,omitempty"`
	LastOrderDate   *Time     `xmlrpc:"last_order_date,omitempty"`
	Name            *String   `xmlrpc:"name,omitempty"`
	NewUntil        *Time     `xmlrpc:"new_until,omitempty"`
	Price           *Float    `xmlrpc:"price,omitempty"`
	ProductImage    *String   `xmlrpc:"product_image,omitempty"`
	SupplierId      *Many2One `xmlrpc:"supplier_id,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// LunchProducts represents array of lunch.product model.
type LunchProducts []LunchProduct

// LunchProductModel is the odoo model name.
const LunchProductModel = "lunch.product"

// Many2One convert LunchProduct to *Many2One.
func (lp *LunchProduct) Many2One() *Many2One {
	return NewMany2One(lp.Id.Get(), "")
}

// CreateLunchProduct creates a new lunch.product model and returns its id.
func (c *Client) CreateLunchProduct(lp *LunchProduct) (int64, error) {
	ids, err := c.CreateLunchProducts([]*LunchProduct{lp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLunchProduct creates a new lunch.product model and returns its id.
func (c *Client) CreateLunchProducts(lps []*LunchProduct) ([]int64, error) {
	var vv []interface{}
	for _, v := range lps {
		vv = append(vv, v)
	}
	return c.Create(LunchProductModel, vv, nil)
}

// UpdateLunchProduct updates an existing lunch.product record.
func (c *Client) UpdateLunchProduct(lp *LunchProduct) error {
	return c.UpdateLunchProducts([]int64{lp.Id.Get()}, lp)
}

// UpdateLunchProducts updates existing lunch.product records.
// All records (represented by ids) will be updated by lp values.
func (c *Client) UpdateLunchProducts(ids []int64, lp *LunchProduct) error {
	return c.Update(LunchProductModel, ids, lp, nil)
}

// DeleteLunchProduct deletes an existing lunch.product record.
func (c *Client) DeleteLunchProduct(id int64) error {
	return c.DeleteLunchProducts([]int64{id})
}

// DeleteLunchProducts deletes existing lunch.product records.
func (c *Client) DeleteLunchProducts(ids []int64) error {
	return c.Delete(LunchProductModel, ids)
}

// GetLunchProduct gets lunch.product existing record.
func (c *Client) GetLunchProduct(id int64) (*LunchProduct, error) {
	lps, err := c.GetLunchProducts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lps)[0]), nil
}

// GetLunchProducts gets lunch.product existing records.
func (c *Client) GetLunchProducts(ids []int64) (*LunchProducts, error) {
	lps := &LunchProducts{}
	if err := c.Read(LunchProductModel, ids, nil, lps); err != nil {
		return nil, err
	}
	return lps, nil
}

// FindLunchProduct finds lunch.product record by querying it with criteria.
func (c *Client) FindLunchProduct(criteria *Criteria) (*LunchProduct, error) {
	lps := &LunchProducts{}
	if err := c.SearchRead(LunchProductModel, criteria, NewOptions().Limit(1), lps); err != nil {
		return nil, err
	}
	return &((*lps)[0]), nil
}

// FindLunchProducts finds lunch.product records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchProducts(criteria *Criteria, options *Options) (*LunchProducts, error) {
	lps := &LunchProducts{}
	if err := c.SearchRead(LunchProductModel, criteria, options, lps); err != nil {
		return nil, err
	}
	return lps, nil
}

// FindLunchProductIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchProductIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(LunchProductModel, criteria, options)
}

// FindLunchProductId finds record id by querying it with criteria.
func (c *Client) FindLunchProductId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LunchProductModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
