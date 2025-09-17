package odoo

// StockWarnInsufficientQtyUnbuild represents stock.warn.insufficient.qty.unbuild model.
type StockWarnInsufficientQtyUnbuild struct {
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	LocationId     *Many2One `xmlrpc:"location_id,omitempty"`
	ProductId      *Many2One `xmlrpc:"product_id,omitempty"`
	ProductUomName *String   `xmlrpc:"product_uom_name,omitempty"`
	QuantIds       *Relation `xmlrpc:"quant_ids,omitempty"`
	Quantity       *Float    `xmlrpc:"quantity,omitempty"`
	UnbuildId      *Many2One `xmlrpc:"unbuild_id,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockWarnInsufficientQtyUnbuilds represents array of stock.warn.insufficient.qty.unbuild model.
type StockWarnInsufficientQtyUnbuilds []StockWarnInsufficientQtyUnbuild

// StockWarnInsufficientQtyUnbuildModel is the odoo model name.
const StockWarnInsufficientQtyUnbuildModel = "stock.warn.insufficient.qty.unbuild"

// Many2One convert StockWarnInsufficientQtyUnbuild to *Many2One.
func (swiqu *StockWarnInsufficientQtyUnbuild) Many2One() *Many2One {
	return NewMany2One(swiqu.Id.Get(), "")
}

// CreateStockWarnInsufficientQtyUnbuild creates a new stock.warn.insufficient.qty.unbuild model and returns its id.
func (c *Client) CreateStockWarnInsufficientQtyUnbuild(swiqu *StockWarnInsufficientQtyUnbuild) (int64, error) {
	ids, err := c.CreateStockWarnInsufficientQtyUnbuilds([]*StockWarnInsufficientQtyUnbuild{swiqu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockWarnInsufficientQtyUnbuild creates a new stock.warn.insufficient.qty.unbuild model and returns its id.
func (c *Client) CreateStockWarnInsufficientQtyUnbuilds(swiqus []*StockWarnInsufficientQtyUnbuild) ([]int64, error) {
	var vv []interface{}
	for _, v := range swiqus {
		vv = append(vv, v)
	}
	return c.Create(StockWarnInsufficientQtyUnbuildModel, vv, nil)
}

// UpdateStockWarnInsufficientQtyUnbuild updates an existing stock.warn.insufficient.qty.unbuild record.
func (c *Client) UpdateStockWarnInsufficientQtyUnbuild(swiqu *StockWarnInsufficientQtyUnbuild) error {
	return c.UpdateStockWarnInsufficientQtyUnbuilds([]int64{swiqu.Id.Get()}, swiqu)
}

// UpdateStockWarnInsufficientQtyUnbuilds updates existing stock.warn.insufficient.qty.unbuild records.
// All records (represented by ids) will be updated by swiqu values.
func (c *Client) UpdateStockWarnInsufficientQtyUnbuilds(ids []int64, swiqu *StockWarnInsufficientQtyUnbuild) error {
	return c.Update(StockWarnInsufficientQtyUnbuildModel, ids, swiqu, nil)
}

// DeleteStockWarnInsufficientQtyUnbuild deletes an existing stock.warn.insufficient.qty.unbuild record.
func (c *Client) DeleteStockWarnInsufficientQtyUnbuild(id int64) error {
	return c.DeleteStockWarnInsufficientQtyUnbuilds([]int64{id})
}

// DeleteStockWarnInsufficientQtyUnbuilds deletes existing stock.warn.insufficient.qty.unbuild records.
func (c *Client) DeleteStockWarnInsufficientQtyUnbuilds(ids []int64) error {
	return c.Delete(StockWarnInsufficientQtyUnbuildModel, ids)
}

// GetStockWarnInsufficientQtyUnbuild gets stock.warn.insufficient.qty.unbuild existing record.
func (c *Client) GetStockWarnInsufficientQtyUnbuild(id int64) (*StockWarnInsufficientQtyUnbuild, error) {
	swiqus, err := c.GetStockWarnInsufficientQtyUnbuilds([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*swiqus)[0]), nil
}

// GetStockWarnInsufficientQtyUnbuilds gets stock.warn.insufficient.qty.unbuild existing records.
func (c *Client) GetStockWarnInsufficientQtyUnbuilds(ids []int64) (*StockWarnInsufficientQtyUnbuilds, error) {
	swiqus := &StockWarnInsufficientQtyUnbuilds{}
	if err := c.Read(StockWarnInsufficientQtyUnbuildModel, ids, nil, swiqus); err != nil {
		return nil, err
	}
	return swiqus, nil
}

// FindStockWarnInsufficientQtyUnbuild finds stock.warn.insufficient.qty.unbuild record by querying it with criteria.
func (c *Client) FindStockWarnInsufficientQtyUnbuild(criteria *Criteria) (*StockWarnInsufficientQtyUnbuild, error) {
	swiqus := &StockWarnInsufficientQtyUnbuilds{}
	if err := c.SearchRead(StockWarnInsufficientQtyUnbuildModel, criteria, NewOptions().Limit(1), swiqus); err != nil {
		return nil, err
	}
	return &((*swiqus)[0]), nil
}

// FindStockWarnInsufficientQtyUnbuilds finds stock.warn.insufficient.qty.unbuild records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockWarnInsufficientQtyUnbuilds(criteria *Criteria, options *Options) (*StockWarnInsufficientQtyUnbuilds, error) {
	swiqus := &StockWarnInsufficientQtyUnbuilds{}
	if err := c.SearchRead(StockWarnInsufficientQtyUnbuildModel, criteria, options, swiqus); err != nil {
		return nil, err
	}
	return swiqus, nil
}

// FindStockWarnInsufficientQtyUnbuildIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockWarnInsufficientQtyUnbuildIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockWarnInsufficientQtyUnbuildModel, criteria, options)
}

// FindStockWarnInsufficientQtyUnbuildId finds record id by querying it with criteria.
func (c *Client) FindStockWarnInsufficientQtyUnbuildId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockWarnInsufficientQtyUnbuildModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
