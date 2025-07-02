package odoo

// RegistrationEditor represents registration.editor model.
type RegistrationEditor struct {
	CreateDate           *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName          *String   `xmlrpc:"display_name,omitempty"`
	EventRegistrationIds *Relation `xmlrpc:"event_registration_ids,omitempty"`
	Id                   *Int      `xmlrpc:"id,omitempty"`
	SaleOrderId          *Many2One `xmlrpc:"sale_order_id,omitempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omitempty"`
}

// RegistrationEditors represents array of registration.editor model.
type RegistrationEditors []RegistrationEditor

// RegistrationEditorModel is the odoo model name.
const RegistrationEditorModel = "registration.editor"

// Many2One convert RegistrationEditor to *Many2One.
func (re *RegistrationEditor) Many2One() *Many2One {
	return NewMany2One(re.Id.Get(), "")
}

// CreateRegistrationEditor creates a new registration.editor model and returns its id.
func (c *Client) CreateRegistrationEditor(re *RegistrationEditor) (int64, error) {
	ids, err := c.CreateRegistrationEditors([]*RegistrationEditor{re})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateRegistrationEditor creates a new registration.editor model and returns its id.
func (c *Client) CreateRegistrationEditors(res []*RegistrationEditor) ([]int64, error) {
	var vv []interface{}
	for _, v := range res {
		vv = append(vv, v)
	}
	return c.Create(RegistrationEditorModel, vv, nil)
}

// UpdateRegistrationEditor updates an existing registration.editor record.
func (c *Client) UpdateRegistrationEditor(re *RegistrationEditor) error {
	return c.UpdateRegistrationEditors([]int64{re.Id.Get()}, re)
}

// UpdateRegistrationEditors updates existing registration.editor records.
// All records (represented by ids) will be updated by re values.
func (c *Client) UpdateRegistrationEditors(ids []int64, re *RegistrationEditor) error {
	return c.Update(RegistrationEditorModel, ids, re, nil)
}

// DeleteRegistrationEditor deletes an existing registration.editor record.
func (c *Client) DeleteRegistrationEditor(id int64) error {
	return c.DeleteRegistrationEditors([]int64{id})
}

// DeleteRegistrationEditors deletes existing registration.editor records.
func (c *Client) DeleteRegistrationEditors(ids []int64) error {
	return c.Delete(RegistrationEditorModel, ids)
}

// GetRegistrationEditor gets registration.editor existing record.
func (c *Client) GetRegistrationEditor(id int64) (*RegistrationEditor, error) {
	res, err := c.GetRegistrationEditors([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*res)[0]), nil
}

// GetRegistrationEditors gets registration.editor existing records.
func (c *Client) GetRegistrationEditors(ids []int64) (*RegistrationEditors, error) {
	res := &RegistrationEditors{}
	if err := c.Read(RegistrationEditorModel, ids, nil, res); err != nil {
		return nil, err
	}
	return res, nil
}

// FindRegistrationEditor finds registration.editor record by querying it with criteria.
func (c *Client) FindRegistrationEditor(criteria *Criteria) (*RegistrationEditor, error) {
	res := &RegistrationEditors{}
	if err := c.SearchRead(RegistrationEditorModel, criteria, NewOptions().Limit(1), res); err != nil {
		return nil, err
	}
	return &((*res)[0]), nil
}

// FindRegistrationEditors finds registration.editor records by querying it
// and filtering it with criteria and options.
func (c *Client) FindRegistrationEditors(criteria *Criteria, options *Options) (*RegistrationEditors, error) {
	res := &RegistrationEditors{}
	if err := c.SearchRead(RegistrationEditorModel, criteria, options, res); err != nil {
		return nil, err
	}
	return res, nil
}

// FindRegistrationEditorIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindRegistrationEditorIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(RegistrationEditorModel, criteria, options)
}

// FindRegistrationEditorId finds record id by querying it with criteria.
func (c *Client) FindRegistrationEditorId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(RegistrationEditorModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
