package odoo

// DynamicFieldWidgets represents dynamic.field.widgets model.
type DynamicFieldWidgets struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DataType    *String   `xmlrpc:"data_type,omitempty"`
	Description *String   `xmlrpc:"description,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// DynamicFieldWidgetss represents array of dynamic.field.widgets model.
type DynamicFieldWidgetss []DynamicFieldWidgets

// DynamicFieldWidgetsModel is the odoo model name.
const DynamicFieldWidgetsModel = "dynamic.field.widgets"

// Many2One convert DynamicFieldWidgets to *Many2One.
func (dfw *DynamicFieldWidgets) Many2One() *Many2One {
	return NewMany2One(dfw.Id.Get(), "")
}

// CreateDynamicFieldWidgets creates a new dynamic.field.widgets model and returns its id.
func (c *Client) CreateDynamicFieldWidgets(dfw *DynamicFieldWidgets) (int64, error) {
	ids, err := c.CreateDynamicFieldWidgetss([]*DynamicFieldWidgets{dfw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDynamicFieldWidgets creates a new dynamic.field.widgets model and returns its id.
func (c *Client) CreateDynamicFieldWidgetss(dfws []*DynamicFieldWidgets) ([]int64, error) {
	var vv []interface{}
	for _, v := range dfws {
		vv = append(vv, v)
	}
	return c.Create(DynamicFieldWidgetsModel, vv, nil)
}

// UpdateDynamicFieldWidgets updates an existing dynamic.field.widgets record.
func (c *Client) UpdateDynamicFieldWidgets(dfw *DynamicFieldWidgets) error {
	return c.UpdateDynamicFieldWidgetss([]int64{dfw.Id.Get()}, dfw)
}

// UpdateDynamicFieldWidgetss updates existing dynamic.field.widgets records.
// All records (represented by ids) will be updated by dfw values.
func (c *Client) UpdateDynamicFieldWidgetss(ids []int64, dfw *DynamicFieldWidgets) error {
	return c.Update(DynamicFieldWidgetsModel, ids, dfw, nil)
}

// DeleteDynamicFieldWidgets deletes an existing dynamic.field.widgets record.
func (c *Client) DeleteDynamicFieldWidgets(id int64) error {
	return c.DeleteDynamicFieldWidgetss([]int64{id})
}

// DeleteDynamicFieldWidgetss deletes existing dynamic.field.widgets records.
func (c *Client) DeleteDynamicFieldWidgetss(ids []int64) error {
	return c.Delete(DynamicFieldWidgetsModel, ids)
}

// GetDynamicFieldWidgets gets dynamic.field.widgets existing record.
func (c *Client) GetDynamicFieldWidgets(id int64) (*DynamicFieldWidgets, error) {
	dfws, err := c.GetDynamicFieldWidgetss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dfws)[0]), nil
}

// GetDynamicFieldWidgetss gets dynamic.field.widgets existing records.
func (c *Client) GetDynamicFieldWidgetss(ids []int64) (*DynamicFieldWidgetss, error) {
	dfws := &DynamicFieldWidgetss{}
	if err := c.Read(DynamicFieldWidgetsModel, ids, nil, dfws); err != nil {
		return nil, err
	}
	return dfws, nil
}

// FindDynamicFieldWidgets finds dynamic.field.widgets record by querying it with criteria.
func (c *Client) FindDynamicFieldWidgets(criteria *Criteria) (*DynamicFieldWidgets, error) {
	dfws := &DynamicFieldWidgetss{}
	if err := c.SearchRead(DynamicFieldWidgetsModel, criteria, NewOptions().Limit(1), dfws); err != nil {
		return nil, err
	}
	return &((*dfws)[0]), nil
}

// FindDynamicFieldWidgetss finds dynamic.field.widgets records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDynamicFieldWidgetss(criteria *Criteria, options *Options) (*DynamicFieldWidgetss, error) {
	dfws := &DynamicFieldWidgetss{}
	if err := c.SearchRead(DynamicFieldWidgetsModel, criteria, options, dfws); err != nil {
		return nil, err
	}
	return dfws, nil
}

// FindDynamicFieldWidgetsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDynamicFieldWidgetsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DynamicFieldWidgetsModel, criteria, options)
}

// FindDynamicFieldWidgetsId finds record id by querying it with criteria.
func (c *Client) FindDynamicFieldWidgetsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DynamicFieldWidgetsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
