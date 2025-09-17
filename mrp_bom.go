package odoo

// MrpBom represents mrp.bom model.
type MrpBom struct {
	Active                                   *Bool      `xmlrpc:"active,omitempty"`
	AllowOperationDependencies               *Bool      `xmlrpc:"allow_operation_dependencies,omitempty"`
	BomLineIds                               *Relation  `xmlrpc:"bom_line_ids,omitempty"`
	ByproductIds                             *Relation  `xmlrpc:"byproduct_ids,omitempty"`
	Code                                     *String    `xmlrpc:"code,omitempty"`
	CompanyId                                *Many2One  `xmlrpc:"company_id,omitempty"`
	Consumption                              *Selection `xmlrpc:"consumption,omitempty"`
	CreateDate                               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                                *Many2One  `xmlrpc:"create_uid,omitempty"`
	DaysToPrepareMo                          *Int       `xmlrpc:"days_to_prepare_mo,omitempty"`
	DisplayName                              *String    `xmlrpc:"display_name,omitempty"`
	HasMessage                               *Bool      `xmlrpc:"has_message,omitempty"`
	Id                                       *Int       `xmlrpc:"id,omitempty"`
	MessageAttachmentCount                   *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds                       *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                          *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter                   *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError                       *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                               *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                        *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction                        *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter                 *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                        *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	OperationIds                             *Relation  `xmlrpc:"operation_ids,omitempty"`
	PickingTypeId                            *Many2One  `xmlrpc:"picking_type_id,omitempty"`
	PossibleProductTemplateAttributeValueIds *Relation  `xmlrpc:"possible_product_template_attribute_value_ids,omitempty"`
	ProduceDelay                             *Int       `xmlrpc:"produce_delay,omitempty"`
	ProductId                                *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductQty                               *Float     `xmlrpc:"product_qty,omitempty"`
	ProductTmplId                            *Many2One  `xmlrpc:"product_tmpl_id,omitempty"`
	ProductUomCategoryId                     *Many2One  `xmlrpc:"product_uom_category_id,omitempty"`
	ProductUomId                             *Many2One  `xmlrpc:"product_uom_id,omitempty"`
	ProjectId                                *Many2One  `xmlrpc:"project_id,omitempty"`
	RatingIds                                *Relation  `xmlrpc:"rating_ids,omitempty"`
	ReadyToProduce                           *Selection `xmlrpc:"ready_to_produce,omitempty"`
	Sequence                                 *Int       `xmlrpc:"sequence,omitempty"`
	Type                                     *Selection `xmlrpc:"type,omitempty"`
	WebsiteMessageIds                        *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MrpBoms represents array of mrp.bom model.
type MrpBoms []MrpBom

// MrpBomModel is the odoo model name.
const MrpBomModel = "mrp.bom"

// Many2One convert MrpBom to *Many2One.
func (mb *MrpBom) Many2One() *Many2One {
	return NewMany2One(mb.Id.Get(), "")
}

// CreateMrpBom creates a new mrp.bom model and returns its id.
func (c *Client) CreateMrpBom(mb *MrpBom) (int64, error) {
	ids, err := c.CreateMrpBoms([]*MrpBom{mb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpBom creates a new mrp.bom model and returns its id.
func (c *Client) CreateMrpBoms(mbs []*MrpBom) ([]int64, error) {
	var vv []interface{}
	for _, v := range mbs {
		vv = append(vv, v)
	}
	return c.Create(MrpBomModel, vv, nil)
}

// UpdateMrpBom updates an existing mrp.bom record.
func (c *Client) UpdateMrpBom(mb *MrpBom) error {
	return c.UpdateMrpBoms([]int64{mb.Id.Get()}, mb)
}

// UpdateMrpBoms updates existing mrp.bom records.
// All records (represented by ids) will be updated by mb values.
func (c *Client) UpdateMrpBoms(ids []int64, mb *MrpBom) error {
	return c.Update(MrpBomModel, ids, mb, nil)
}

// DeleteMrpBom deletes an existing mrp.bom record.
func (c *Client) DeleteMrpBom(id int64) error {
	return c.DeleteMrpBoms([]int64{id})
}

// DeleteMrpBoms deletes existing mrp.bom records.
func (c *Client) DeleteMrpBoms(ids []int64) error {
	return c.Delete(MrpBomModel, ids)
}

// GetMrpBom gets mrp.bom existing record.
func (c *Client) GetMrpBom(id int64) (*MrpBom, error) {
	mbs, err := c.GetMrpBoms([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mbs)[0]), nil
}

// GetMrpBoms gets mrp.bom existing records.
func (c *Client) GetMrpBoms(ids []int64) (*MrpBoms, error) {
	mbs := &MrpBoms{}
	if err := c.Read(MrpBomModel, ids, nil, mbs); err != nil {
		return nil, err
	}
	return mbs, nil
}

// FindMrpBom finds mrp.bom record by querying it with criteria.
func (c *Client) FindMrpBom(criteria *Criteria) (*MrpBom, error) {
	mbs := &MrpBoms{}
	if err := c.SearchRead(MrpBomModel, criteria, NewOptions().Limit(1), mbs); err != nil {
		return nil, err
	}
	return &((*mbs)[0]), nil
}

// FindMrpBoms finds mrp.bom records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpBoms(criteria *Criteria, options *Options) (*MrpBoms, error) {
	mbs := &MrpBoms{}
	if err := c.SearchRead(MrpBomModel, criteria, options, mbs); err != nil {
		return nil, err
	}
	return mbs, nil
}

// FindMrpBomIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpBomIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpBomModel, criteria, options)
}

// FindMrpBomId finds record id by querying it with criteria.
func (c *Client) FindMrpBomId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpBomModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
