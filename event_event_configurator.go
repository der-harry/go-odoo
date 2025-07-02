package odoo

// EventEventConfigurator represents event.event.configurator model.
type EventEventConfigurator struct {
	CreateDate          *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName         *String   `xmlrpc:"display_name,omitempty"`
	EventId             *Many2One `xmlrpc:"event_id,omitempty"`
	EventTicketId       *Many2One `xmlrpc:"event_ticket_id,omitempty"`
	HasAvailableTickets *Bool     `xmlrpc:"has_available_tickets,omitempty"`
	Id                  *Int      `xmlrpc:"id,omitempty"`
	ProductId           *Many2One `xmlrpc:"product_id,omitempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omitempty"`
}

// EventEventConfigurators represents array of event.event.configurator model.
type EventEventConfigurators []EventEventConfigurator

// EventEventConfiguratorModel is the odoo model name.
const EventEventConfiguratorModel = "event.event.configurator"

// Many2One convert EventEventConfigurator to *Many2One.
func (eec *EventEventConfigurator) Many2One() *Many2One {
	return NewMany2One(eec.Id.Get(), "")
}

// CreateEventEventConfigurator creates a new event.event.configurator model and returns its id.
func (c *Client) CreateEventEventConfigurator(eec *EventEventConfigurator) (int64, error) {
	ids, err := c.CreateEventEventConfigurators([]*EventEventConfigurator{eec})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventEventConfigurator creates a new event.event.configurator model and returns its id.
func (c *Client) CreateEventEventConfigurators(eecs []*EventEventConfigurator) ([]int64, error) {
	var vv []interface{}
	for _, v := range eecs {
		vv = append(vv, v)
	}
	return c.Create(EventEventConfiguratorModel, vv, nil)
}

// UpdateEventEventConfigurator updates an existing event.event.configurator record.
func (c *Client) UpdateEventEventConfigurator(eec *EventEventConfigurator) error {
	return c.UpdateEventEventConfigurators([]int64{eec.Id.Get()}, eec)
}

// UpdateEventEventConfigurators updates existing event.event.configurator records.
// All records (represented by ids) will be updated by eec values.
func (c *Client) UpdateEventEventConfigurators(ids []int64, eec *EventEventConfigurator) error {
	return c.Update(EventEventConfiguratorModel, ids, eec, nil)
}

// DeleteEventEventConfigurator deletes an existing event.event.configurator record.
func (c *Client) DeleteEventEventConfigurator(id int64) error {
	return c.DeleteEventEventConfigurators([]int64{id})
}

// DeleteEventEventConfigurators deletes existing event.event.configurator records.
func (c *Client) DeleteEventEventConfigurators(ids []int64) error {
	return c.Delete(EventEventConfiguratorModel, ids)
}

// GetEventEventConfigurator gets event.event.configurator existing record.
func (c *Client) GetEventEventConfigurator(id int64) (*EventEventConfigurator, error) {
	eecs, err := c.GetEventEventConfigurators([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*eecs)[0]), nil
}

// GetEventEventConfigurators gets event.event.configurator existing records.
func (c *Client) GetEventEventConfigurators(ids []int64) (*EventEventConfigurators, error) {
	eecs := &EventEventConfigurators{}
	if err := c.Read(EventEventConfiguratorModel, ids, nil, eecs); err != nil {
		return nil, err
	}
	return eecs, nil
}

// FindEventEventConfigurator finds event.event.configurator record by querying it with criteria.
func (c *Client) FindEventEventConfigurator(criteria *Criteria) (*EventEventConfigurator, error) {
	eecs := &EventEventConfigurators{}
	if err := c.SearchRead(EventEventConfiguratorModel, criteria, NewOptions().Limit(1), eecs); err != nil {
		return nil, err
	}
	return &((*eecs)[0]), nil
}

// FindEventEventConfigurators finds event.event.configurator records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventEventConfigurators(criteria *Criteria, options *Options) (*EventEventConfigurators, error) {
	eecs := &EventEventConfigurators{}
	if err := c.SearchRead(EventEventConfiguratorModel, criteria, options, eecs); err != nil {
		return nil, err
	}
	return eecs, nil
}

// FindEventEventConfiguratorIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventEventConfiguratorIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(EventEventConfiguratorModel, criteria, options)
}

// FindEventEventConfiguratorId finds record id by querying it with criteria.
func (c *Client) FindEventEventConfiguratorId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventEventConfiguratorModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
