package odoo

// MrpConsumptionWarningLine represents mrp.consumption.warning.line model.
type MrpConsumptionWarningLine struct {
	Consumption             *Selection `xmlrpc:"consumption,omitempty"`
	CreateDate              *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid               *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName             *String    `xmlrpc:"display_name,omitempty"`
	Id                      *Int       `xmlrpc:"id,omitempty"`
	MrpConsumptionWarningId *Many2One  `xmlrpc:"mrp_consumption_warning_id,omitempty"`
	MrpProductionId         *Many2One  `xmlrpc:"mrp_production_id,omitempty"`
	ProductConsumedQtyUom   *Float     `xmlrpc:"product_consumed_qty_uom,omitempty"`
	ProductExpectedQtyUom   *Float     `xmlrpc:"product_expected_qty_uom,omitempty"`
	ProductId               *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductUomId            *Many2One  `xmlrpc:"product_uom_id,omitempty"`
	WriteDate               *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MrpConsumptionWarningLines represents array of mrp.consumption.warning.line model.
type MrpConsumptionWarningLines []MrpConsumptionWarningLine

// MrpConsumptionWarningLineModel is the odoo model name.
const MrpConsumptionWarningLineModel = "mrp.consumption.warning.line"

// Many2One convert MrpConsumptionWarningLine to *Many2One.
func (mcwl *MrpConsumptionWarningLine) Many2One() *Many2One {
	return NewMany2One(mcwl.Id.Get(), "")
}

// CreateMrpConsumptionWarningLine creates a new mrp.consumption.warning.line model and returns its id.
func (c *Client) CreateMrpConsumptionWarningLine(mcwl *MrpConsumptionWarningLine) (int64, error) {
	ids, err := c.CreateMrpConsumptionWarningLines([]*MrpConsumptionWarningLine{mcwl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpConsumptionWarningLine creates a new mrp.consumption.warning.line model and returns its id.
func (c *Client) CreateMrpConsumptionWarningLines(mcwls []*MrpConsumptionWarningLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range mcwls {
		vv = append(vv, v)
	}
	return c.Create(MrpConsumptionWarningLineModel, vv, nil)
}

// UpdateMrpConsumptionWarningLine updates an existing mrp.consumption.warning.line record.
func (c *Client) UpdateMrpConsumptionWarningLine(mcwl *MrpConsumptionWarningLine) error {
	return c.UpdateMrpConsumptionWarningLines([]int64{mcwl.Id.Get()}, mcwl)
}

// UpdateMrpConsumptionWarningLines updates existing mrp.consumption.warning.line records.
// All records (represented by ids) will be updated by mcwl values.
func (c *Client) UpdateMrpConsumptionWarningLines(ids []int64, mcwl *MrpConsumptionWarningLine) error {
	return c.Update(MrpConsumptionWarningLineModel, ids, mcwl, nil)
}

// DeleteMrpConsumptionWarningLine deletes an existing mrp.consumption.warning.line record.
func (c *Client) DeleteMrpConsumptionWarningLine(id int64) error {
	return c.DeleteMrpConsumptionWarningLines([]int64{id})
}

// DeleteMrpConsumptionWarningLines deletes existing mrp.consumption.warning.line records.
func (c *Client) DeleteMrpConsumptionWarningLines(ids []int64) error {
	return c.Delete(MrpConsumptionWarningLineModel, ids)
}

// GetMrpConsumptionWarningLine gets mrp.consumption.warning.line existing record.
func (c *Client) GetMrpConsumptionWarningLine(id int64) (*MrpConsumptionWarningLine, error) {
	mcwls, err := c.GetMrpConsumptionWarningLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mcwls)[0]), nil
}

// GetMrpConsumptionWarningLines gets mrp.consumption.warning.line existing records.
func (c *Client) GetMrpConsumptionWarningLines(ids []int64) (*MrpConsumptionWarningLines, error) {
	mcwls := &MrpConsumptionWarningLines{}
	if err := c.Read(MrpConsumptionWarningLineModel, ids, nil, mcwls); err != nil {
		return nil, err
	}
	return mcwls, nil
}

// FindMrpConsumptionWarningLine finds mrp.consumption.warning.line record by querying it with criteria.
func (c *Client) FindMrpConsumptionWarningLine(criteria *Criteria) (*MrpConsumptionWarningLine, error) {
	mcwls := &MrpConsumptionWarningLines{}
	if err := c.SearchRead(MrpConsumptionWarningLineModel, criteria, NewOptions().Limit(1), mcwls); err != nil {
		return nil, err
	}
	return &((*mcwls)[0]), nil
}

// FindMrpConsumptionWarningLines finds mrp.consumption.warning.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpConsumptionWarningLines(criteria *Criteria, options *Options) (*MrpConsumptionWarningLines, error) {
	mcwls := &MrpConsumptionWarningLines{}
	if err := c.SearchRead(MrpConsumptionWarningLineModel, criteria, options, mcwls); err != nil {
		return nil, err
	}
	return mcwls, nil
}

// FindMrpConsumptionWarningLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpConsumptionWarningLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpConsumptionWarningLineModel, criteria, options)
}

// FindMrpConsumptionWarningLineId finds record id by querying it with criteria.
func (c *Client) FindMrpConsumptionWarningLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpConsumptionWarningLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
