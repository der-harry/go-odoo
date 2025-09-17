package odoo

// MrpBomByproduct represents mrp.bom.byproduct model.
type MrpBomByproduct struct {
	AllowedOperationIds                         *Relation `xmlrpc:"allowed_operation_ids,omitempty"`
	BomId                                       *Many2One `xmlrpc:"bom_id,omitempty"`
	BomProductTemplateAttributeValueIds         *Relation `xmlrpc:"bom_product_template_attribute_value_ids,omitempty"`
	CompanyId                                   *Many2One `xmlrpc:"company_id,omitempty"`
	CostShare                                   *Float    `xmlrpc:"cost_share,omitempty"`
	CreateDate                                  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                                   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName                                 *String   `xmlrpc:"display_name,omitempty"`
	Id                                          *Int      `xmlrpc:"id,omitempty"`
	OperationId                                 *Many2One `xmlrpc:"operation_id,omitempty"`
	PossibleBomProductTemplateAttributeValueIds *Relation `xmlrpc:"possible_bom_product_template_attribute_value_ids,omitempty"`
	ProductId                                   *Many2One `xmlrpc:"product_id,omitempty"`
	ProductQty                                  *Float    `xmlrpc:"product_qty,omitempty"`
	ProductUomCategoryId                        *Many2One `xmlrpc:"product_uom_category_id,omitempty"`
	ProductUomId                                *Many2One `xmlrpc:"product_uom_id,omitempty"`
	Sequence                                    *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate                                   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                                    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MrpBomByproducts represents array of mrp.bom.byproduct model.
type MrpBomByproducts []MrpBomByproduct

// MrpBomByproductModel is the odoo model name.
const MrpBomByproductModel = "mrp.bom.byproduct"

// Many2One convert MrpBomByproduct to *Many2One.
func (mbb *MrpBomByproduct) Many2One() *Many2One {
	return NewMany2One(mbb.Id.Get(), "")
}

// CreateMrpBomByproduct creates a new mrp.bom.byproduct model and returns its id.
func (c *Client) CreateMrpBomByproduct(mbb *MrpBomByproduct) (int64, error) {
	ids, err := c.CreateMrpBomByproducts([]*MrpBomByproduct{mbb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpBomByproduct creates a new mrp.bom.byproduct model and returns its id.
func (c *Client) CreateMrpBomByproducts(mbbs []*MrpBomByproduct) ([]int64, error) {
	var vv []interface{}
	for _, v := range mbbs {
		vv = append(vv, v)
	}
	return c.Create(MrpBomByproductModel, vv, nil)
}

// UpdateMrpBomByproduct updates an existing mrp.bom.byproduct record.
func (c *Client) UpdateMrpBomByproduct(mbb *MrpBomByproduct) error {
	return c.UpdateMrpBomByproducts([]int64{mbb.Id.Get()}, mbb)
}

// UpdateMrpBomByproducts updates existing mrp.bom.byproduct records.
// All records (represented by ids) will be updated by mbb values.
func (c *Client) UpdateMrpBomByproducts(ids []int64, mbb *MrpBomByproduct) error {
	return c.Update(MrpBomByproductModel, ids, mbb, nil)
}

// DeleteMrpBomByproduct deletes an existing mrp.bom.byproduct record.
func (c *Client) DeleteMrpBomByproduct(id int64) error {
	return c.DeleteMrpBomByproducts([]int64{id})
}

// DeleteMrpBomByproducts deletes existing mrp.bom.byproduct records.
func (c *Client) DeleteMrpBomByproducts(ids []int64) error {
	return c.Delete(MrpBomByproductModel, ids)
}

// GetMrpBomByproduct gets mrp.bom.byproduct existing record.
func (c *Client) GetMrpBomByproduct(id int64) (*MrpBomByproduct, error) {
	mbbs, err := c.GetMrpBomByproducts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mbbs)[0]), nil
}

// GetMrpBomByproducts gets mrp.bom.byproduct existing records.
func (c *Client) GetMrpBomByproducts(ids []int64) (*MrpBomByproducts, error) {
	mbbs := &MrpBomByproducts{}
	if err := c.Read(MrpBomByproductModel, ids, nil, mbbs); err != nil {
		return nil, err
	}
	return mbbs, nil
}

// FindMrpBomByproduct finds mrp.bom.byproduct record by querying it with criteria.
func (c *Client) FindMrpBomByproduct(criteria *Criteria) (*MrpBomByproduct, error) {
	mbbs := &MrpBomByproducts{}
	if err := c.SearchRead(MrpBomByproductModel, criteria, NewOptions().Limit(1), mbbs); err != nil {
		return nil, err
	}
	return &((*mbbs)[0]), nil
}

// FindMrpBomByproducts finds mrp.bom.byproduct records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpBomByproducts(criteria *Criteria, options *Options) (*MrpBomByproducts, error) {
	mbbs := &MrpBomByproducts{}
	if err := c.SearchRead(MrpBomByproductModel, criteria, options, mbbs); err != nil {
		return nil, err
	}
	return mbbs, nil
}

// FindMrpBomByproductIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpBomByproductIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpBomByproductModel, criteria, options)
}

// FindMrpBomByproductId finds record id by querying it with criteria.
func (c *Client) FindMrpBomByproductId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpBomByproductModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
