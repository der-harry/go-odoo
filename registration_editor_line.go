package odoo

// RegistrationEditorLine represents registration.editor.line model.
type RegistrationEditorLine struct {
	CompanyId       *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	EditorId        *Many2One `xmlrpc:"editor_id,omitempty"`
	Email           *String   `xmlrpc:"email,omitempty"`
	EventId         *Many2One `xmlrpc:"event_id,omitempty"`
	EventTicketId   *Many2One `xmlrpc:"event_ticket_id,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	Name            *String   `xmlrpc:"name,omitempty"`
	Phone           *String   `xmlrpc:"phone,omitempty"`
	RegistrationId  *Many2One `xmlrpc:"registration_id,omitempty"`
	SaleOrderLineId *Many2One `xmlrpc:"sale_order_line_id,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// RegistrationEditorLines represents array of registration.editor.line model.
type RegistrationEditorLines []RegistrationEditorLine

// RegistrationEditorLineModel is the odoo model name.
const RegistrationEditorLineModel = "registration.editor.line"

// Many2One convert RegistrationEditorLine to *Many2One.
func (rel *RegistrationEditorLine) Many2One() *Many2One {
	return NewMany2One(rel.Id.Get(), "")
}

// CreateRegistrationEditorLine creates a new registration.editor.line model and returns its id.
func (c *Client) CreateRegistrationEditorLine(rel *RegistrationEditorLine) (int64, error) {
	ids, err := c.CreateRegistrationEditorLines([]*RegistrationEditorLine{rel})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateRegistrationEditorLine creates a new registration.editor.line model and returns its id.
func (c *Client) CreateRegistrationEditorLines(rels []*RegistrationEditorLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range rels {
		vv = append(vv, v)
	}
	return c.Create(RegistrationEditorLineModel, vv, nil)
}

// UpdateRegistrationEditorLine updates an existing registration.editor.line record.
func (c *Client) UpdateRegistrationEditorLine(rel *RegistrationEditorLine) error {
	return c.UpdateRegistrationEditorLines([]int64{rel.Id.Get()}, rel)
}

// UpdateRegistrationEditorLines updates existing registration.editor.line records.
// All records (represented by ids) will be updated by rel values.
func (c *Client) UpdateRegistrationEditorLines(ids []int64, rel *RegistrationEditorLine) error {
	return c.Update(RegistrationEditorLineModel, ids, rel, nil)
}

// DeleteRegistrationEditorLine deletes an existing registration.editor.line record.
func (c *Client) DeleteRegistrationEditorLine(id int64) error {
	return c.DeleteRegistrationEditorLines([]int64{id})
}

// DeleteRegistrationEditorLines deletes existing registration.editor.line records.
func (c *Client) DeleteRegistrationEditorLines(ids []int64) error {
	return c.Delete(RegistrationEditorLineModel, ids)
}

// GetRegistrationEditorLine gets registration.editor.line existing record.
func (c *Client) GetRegistrationEditorLine(id int64) (*RegistrationEditorLine, error) {
	rels, err := c.GetRegistrationEditorLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rels)[0]), nil
}

// GetRegistrationEditorLines gets registration.editor.line existing records.
func (c *Client) GetRegistrationEditorLines(ids []int64) (*RegistrationEditorLines, error) {
	rels := &RegistrationEditorLines{}
	if err := c.Read(RegistrationEditorLineModel, ids, nil, rels); err != nil {
		return nil, err
	}
	return rels, nil
}

// FindRegistrationEditorLine finds registration.editor.line record by querying it with criteria.
func (c *Client) FindRegistrationEditorLine(criteria *Criteria) (*RegistrationEditorLine, error) {
	rels := &RegistrationEditorLines{}
	if err := c.SearchRead(RegistrationEditorLineModel, criteria, NewOptions().Limit(1), rels); err != nil {
		return nil, err
	}
	return &((*rels)[0]), nil
}

// FindRegistrationEditorLines finds registration.editor.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindRegistrationEditorLines(criteria *Criteria, options *Options) (*RegistrationEditorLines, error) {
	rels := &RegistrationEditorLines{}
	if err := c.SearchRead(RegistrationEditorLineModel, criteria, options, rels); err != nil {
		return nil, err
	}
	return rels, nil
}

// FindRegistrationEditorLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindRegistrationEditorLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(RegistrationEditorLineModel, criteria, options)
}

// FindRegistrationEditorLineId finds record id by querying it with criteria.
func (c *Client) FindRegistrationEditorLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(RegistrationEditorLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
