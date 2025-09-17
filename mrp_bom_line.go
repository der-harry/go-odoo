package odoo

// MrpBomLine represents mrp.bom.line model.
type MrpBomLine struct {
	AllowedOperationIds                         *Relation  `xmlrpc:"allowed_operation_ids,omitempty"`
	AttachmentsCount                            *Int       `xmlrpc:"attachments_count,omitempty"`
	BomId                                       *Many2One  `xmlrpc:"bom_id,omitempty"`
	BomProductTemplateAttributeValueIds         *Relation  `xmlrpc:"bom_product_template_attribute_value_ids,omitempty"`
	ChildBomId                                  *Many2One  `xmlrpc:"child_bom_id,omitempty"`
	ChildLineIds                                *Relation  `xmlrpc:"child_line_ids,omitempty"`
	CompanyId                                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CostShare                                   *Float     `xmlrpc:"cost_share,omitempty"`
	CreateDate                                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName                                 *String    `xmlrpc:"display_name,omitempty"`
	Id                                          *Int       `xmlrpc:"id,omitempty"`
	ManualConsumption                           *Bool      `xmlrpc:"manual_consumption,omitempty"`
	OperationId                                 *Many2One  `xmlrpc:"operation_id,omitempty"`
	ParentProductTmplId                         *Many2One  `xmlrpc:"parent_product_tmpl_id,omitempty"`
	PossibleBomProductTemplateAttributeValueIds *Relation  `xmlrpc:"possible_bom_product_template_attribute_value_ids,omitempty"`
	ProductId                                   *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductQty                                  *Float     `xmlrpc:"product_qty,omitempty"`
	ProductTmplId                               *Many2One  `xmlrpc:"product_tmpl_id,omitempty"`
	ProductUomCategoryId                        *Many2One  `xmlrpc:"product_uom_category_id,omitempty"`
	ProductUomId                                *Many2One  `xmlrpc:"product_uom_id,omitempty"`
	Sequence                                    *Int       `xmlrpc:"sequence,omitempty"`
	Tracking                                    *Selection `xmlrpc:"tracking,omitempty"`
	WriteDate                                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MrpBomLines represents array of mrp.bom.line model.
type MrpBomLines []MrpBomLine

// MrpBomLineModel is the odoo model name.
const MrpBomLineModel = "mrp.bom.line"

// Many2One convert MrpBomLine to *Many2One.
func (mbl *MrpBomLine) Many2One() *Many2One {
	return NewMany2One(mbl.Id.Get(), "")
}

// CreateMrpBomLine creates a new mrp.bom.line model and returns its id.
func (c *Client) CreateMrpBomLine(mbl *MrpBomLine) (int64, error) {
	ids, err := c.CreateMrpBomLines([]*MrpBomLine{mbl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpBomLine creates a new mrp.bom.line model and returns its id.
func (c *Client) CreateMrpBomLines(mbls []*MrpBomLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range mbls {
		vv = append(vv, v)
	}
	return c.Create(MrpBomLineModel, vv, nil)
}

// UpdateMrpBomLine updates an existing mrp.bom.line record.
func (c *Client) UpdateMrpBomLine(mbl *MrpBomLine) error {
	return c.UpdateMrpBomLines([]int64{mbl.Id.Get()}, mbl)
}

// UpdateMrpBomLines updates existing mrp.bom.line records.
// All records (represented by ids) will be updated by mbl values.
func (c *Client) UpdateMrpBomLines(ids []int64, mbl *MrpBomLine) error {
	return c.Update(MrpBomLineModel, ids, mbl, nil)
}

// DeleteMrpBomLine deletes an existing mrp.bom.line record.
func (c *Client) DeleteMrpBomLine(id int64) error {
	return c.DeleteMrpBomLines([]int64{id})
}

// DeleteMrpBomLines deletes existing mrp.bom.line records.
func (c *Client) DeleteMrpBomLines(ids []int64) error {
	return c.Delete(MrpBomLineModel, ids)
}

// GetMrpBomLine gets mrp.bom.line existing record.
func (c *Client) GetMrpBomLine(id int64) (*MrpBomLine, error) {
	mbls, err := c.GetMrpBomLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mbls)[0]), nil
}

// GetMrpBomLines gets mrp.bom.line existing records.
func (c *Client) GetMrpBomLines(ids []int64) (*MrpBomLines, error) {
	mbls := &MrpBomLines{}
	if err := c.Read(MrpBomLineModel, ids, nil, mbls); err != nil {
		return nil, err
	}
	return mbls, nil
}

// FindMrpBomLine finds mrp.bom.line record by querying it with criteria.
func (c *Client) FindMrpBomLine(criteria *Criteria) (*MrpBomLine, error) {
	mbls := &MrpBomLines{}
	if err := c.SearchRead(MrpBomLineModel, criteria, NewOptions().Limit(1), mbls); err != nil {
		return nil, err
	}
	return &((*mbls)[0]), nil
}

// FindMrpBomLines finds mrp.bom.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpBomLines(criteria *Criteria, options *Options) (*MrpBomLines, error) {
	mbls := &MrpBomLines{}
	if err := c.SearchRead(MrpBomLineModel, criteria, options, mbls); err != nil {
		return nil, err
	}
	return mbls, nil
}

// FindMrpBomLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpBomLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpBomLineModel, criteria, options)
}

// FindMrpBomLineId finds record id by querying it with criteria.
func (c *Client) FindMrpBomLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpBomLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
