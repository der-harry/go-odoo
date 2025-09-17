package odoo

// MrpBatchProduce represents mrp.batch.produce model.
type MrpBatchProduce struct {
	ComponentSeparator    *String   `xmlrpc:"component_separator,omitempty"`
	CreateDate            *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName           *String   `xmlrpc:"display_name,omitempty"`
	Id                    *Int      `xmlrpc:"id,omitempty"`
	LotName               *String   `xmlrpc:"lot_name,omitempty"`
	LotQty                *Int      `xmlrpc:"lot_qty,omitempty"`
	LotsQuantitySeparator *String   `xmlrpc:"lots_quantity_separator,omitempty"`
	LotsSeparator         *String   `xmlrpc:"lots_separator,omitempty"`
	ProductionId          *Many2One `xmlrpc:"production_id,omitempty"`
	ProductionText        *String   `xmlrpc:"production_text,omitempty"`
	ProductionTextHelp    *String   `xmlrpc:"production_text_help,omitempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MrpBatchProduces represents array of mrp.batch.produce model.
type MrpBatchProduces []MrpBatchProduce

// MrpBatchProduceModel is the odoo model name.
const MrpBatchProduceModel = "mrp.batch.produce"

// Many2One convert MrpBatchProduce to *Many2One.
func (mbp *MrpBatchProduce) Many2One() *Many2One {
	return NewMany2One(mbp.Id.Get(), "")
}

// CreateMrpBatchProduce creates a new mrp.batch.produce model and returns its id.
func (c *Client) CreateMrpBatchProduce(mbp *MrpBatchProduce) (int64, error) {
	ids, err := c.CreateMrpBatchProduces([]*MrpBatchProduce{mbp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpBatchProduce creates a new mrp.batch.produce model and returns its id.
func (c *Client) CreateMrpBatchProduces(mbps []*MrpBatchProduce) ([]int64, error) {
	var vv []interface{}
	for _, v := range mbps {
		vv = append(vv, v)
	}
	return c.Create(MrpBatchProduceModel, vv, nil)
}

// UpdateMrpBatchProduce updates an existing mrp.batch.produce record.
func (c *Client) UpdateMrpBatchProduce(mbp *MrpBatchProduce) error {
	return c.UpdateMrpBatchProduces([]int64{mbp.Id.Get()}, mbp)
}

// UpdateMrpBatchProduces updates existing mrp.batch.produce records.
// All records (represented by ids) will be updated by mbp values.
func (c *Client) UpdateMrpBatchProduces(ids []int64, mbp *MrpBatchProduce) error {
	return c.Update(MrpBatchProduceModel, ids, mbp, nil)
}

// DeleteMrpBatchProduce deletes an existing mrp.batch.produce record.
func (c *Client) DeleteMrpBatchProduce(id int64) error {
	return c.DeleteMrpBatchProduces([]int64{id})
}

// DeleteMrpBatchProduces deletes existing mrp.batch.produce records.
func (c *Client) DeleteMrpBatchProduces(ids []int64) error {
	return c.Delete(MrpBatchProduceModel, ids)
}

// GetMrpBatchProduce gets mrp.batch.produce existing record.
func (c *Client) GetMrpBatchProduce(id int64) (*MrpBatchProduce, error) {
	mbps, err := c.GetMrpBatchProduces([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mbps)[0]), nil
}

// GetMrpBatchProduces gets mrp.batch.produce existing records.
func (c *Client) GetMrpBatchProduces(ids []int64) (*MrpBatchProduces, error) {
	mbps := &MrpBatchProduces{}
	if err := c.Read(MrpBatchProduceModel, ids, nil, mbps); err != nil {
		return nil, err
	}
	return mbps, nil
}

// FindMrpBatchProduce finds mrp.batch.produce record by querying it with criteria.
func (c *Client) FindMrpBatchProduce(criteria *Criteria) (*MrpBatchProduce, error) {
	mbps := &MrpBatchProduces{}
	if err := c.SearchRead(MrpBatchProduceModel, criteria, NewOptions().Limit(1), mbps); err != nil {
		return nil, err
	}
	return &((*mbps)[0]), nil
}

// FindMrpBatchProduces finds mrp.batch.produce records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpBatchProduces(criteria *Criteria, options *Options) (*MrpBatchProduces, error) {
	mbps := &MrpBatchProduces{}
	if err := c.SearchRead(MrpBatchProduceModel, criteria, options, mbps); err != nil {
		return nil, err
	}
	return mbps, nil
}

// FindMrpBatchProduceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpBatchProduceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpBatchProduceModel, criteria, options)
}

// FindMrpBatchProduceId finds record id by querying it with criteria.
func (c *Client) FindMrpBatchProduceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpBatchProduceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
