package odoo

// ProductPricelist represents product.pricelist model.
type ProductPricelist struct {
	Active                      *Bool      `xmlrpc:"active,omitempty"`
	ActivityCalendarEventId     *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryGroupIds             *Relation  `xmlrpc:"country_group_ids,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	ItemIds                     *Relation  `xmlrpc:"item_ids,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	Sequence                    *Int       `xmlrpc:"sequence,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ProductPricelists represents array of product.pricelist model.
type ProductPricelists []ProductPricelist

// ProductPricelistModel is the odoo model name.
const ProductPricelistModel = "product.pricelist"

// Many2One convert ProductPricelist to *Many2One.
func (pp *ProductPricelist) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreateProductPricelist creates a new product.pricelist model and returns its id.
func (c *Client) CreateProductPricelist(pp *ProductPricelist) (int64, error) {
	ids, err := c.CreateProductPricelists([]*ProductPricelist{pp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductPricelist creates a new product.pricelist model and returns its id.
func (c *Client) CreateProductPricelists(pps []*ProductPricelist) ([]int64, error) {
	var vv []interface{}
	for _, v := range pps {
		vv = append(vv, v)
	}
	return c.Create(ProductPricelistModel, vv, nil)
}

// UpdateProductPricelist updates an existing product.pricelist record.
func (c *Client) UpdateProductPricelist(pp *ProductPricelist) error {
	return c.UpdateProductPricelists([]int64{pp.Id.Get()}, pp)
}

// UpdateProductPricelists updates existing product.pricelist records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProductPricelists(ids []int64, pp *ProductPricelist) error {
	return c.Update(ProductPricelistModel, ids, pp, nil)
}

// DeleteProductPricelist deletes an existing product.pricelist record.
func (c *Client) DeleteProductPricelist(id int64) error {
	return c.DeleteProductPricelists([]int64{id})
}

// DeleteProductPricelists deletes existing product.pricelist records.
func (c *Client) DeleteProductPricelists(ids []int64) error {
	return c.Delete(ProductPricelistModel, ids)
}

// GetProductPricelist gets product.pricelist existing record.
func (c *Client) GetProductPricelist(id int64) (*ProductPricelist, error) {
	pps, err := c.GetProductPricelists([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pps)[0]), nil
}

// GetProductPricelists gets product.pricelist existing records.
func (c *Client) GetProductPricelists(ids []int64) (*ProductPricelists, error) {
	pps := &ProductPricelists{}
	if err := c.Read(ProductPricelistModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductPricelist finds product.pricelist record by querying it with criteria.
func (c *Client) FindProductPricelist(criteria *Criteria) (*ProductPricelist, error) {
	pps := &ProductPricelists{}
	if err := c.SearchRead(ProductPricelistModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	return &((*pps)[0]), nil
}

// FindProductPricelists finds product.pricelist records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPricelists(criteria *Criteria, options *Options) (*ProductPricelists, error) {
	pps := &ProductPricelists{}
	if err := c.SearchRead(ProductPricelistModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductPricelistIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPricelistIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductPricelistModel, criteria, options)
}

// FindProductPricelistId finds record id by querying it with criteria.
func (c *Client) FindProductPricelistId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductPricelistModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
