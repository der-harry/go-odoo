package odoo

// EventTypeTicket represents event.type.ticket model.
type EventTypeTicket struct {
	CreateDate   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId   *Many2One `xmlrpc:"currency_id,omitempty"`
	Description  *String   `xmlrpc:"description,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty"`
	EventTypeId  *Many2One `xmlrpc:"event_type_id,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty"`
	Name         *String   `xmlrpc:"name,omitempty"`
	Price        *Float    `xmlrpc:"price,omitempty"`
	PriceReduce  *Float    `xmlrpc:"price_reduce,omitempty"`
	ProductId    *Many2One `xmlrpc:"product_id,omitempty"`
	SeatsLimited *Bool     `xmlrpc:"seats_limited,omitempty"`
	SeatsMax     *Int      `xmlrpc:"seats_max,omitempty"`
	Sequence     *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omitempty"`
}

// EventTypeTickets represents array of event.type.ticket model.
type EventTypeTickets []EventTypeTicket

// EventTypeTicketModel is the odoo model name.
const EventTypeTicketModel = "event.type.ticket"

// Many2One convert EventTypeTicket to *Many2One.
func (ett *EventTypeTicket) Many2One() *Many2One {
	return NewMany2One(ett.Id.Get(), "")
}

// CreateEventTypeTicket creates a new event.type.ticket model and returns its id.
func (c *Client) CreateEventTypeTicket(ett *EventTypeTicket) (int64, error) {
	ids, err := c.CreateEventTypeTickets([]*EventTypeTicket{ett})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventTypeTicket creates a new event.type.ticket model and returns its id.
func (c *Client) CreateEventTypeTickets(etts []*EventTypeTicket) ([]int64, error) {
	var vv []interface{}
	for _, v := range etts {
		vv = append(vv, v)
	}
	return c.Create(EventTypeTicketModel, vv, nil)
}

// UpdateEventTypeTicket updates an existing event.type.ticket record.
func (c *Client) UpdateEventTypeTicket(ett *EventTypeTicket) error {
	return c.UpdateEventTypeTickets([]int64{ett.Id.Get()}, ett)
}

// UpdateEventTypeTickets updates existing event.type.ticket records.
// All records (represented by ids) will be updated by ett values.
func (c *Client) UpdateEventTypeTickets(ids []int64, ett *EventTypeTicket) error {
	return c.Update(EventTypeTicketModel, ids, ett, nil)
}

// DeleteEventTypeTicket deletes an existing event.type.ticket record.
func (c *Client) DeleteEventTypeTicket(id int64) error {
	return c.DeleteEventTypeTickets([]int64{id})
}

// DeleteEventTypeTickets deletes existing event.type.ticket records.
func (c *Client) DeleteEventTypeTickets(ids []int64) error {
	return c.Delete(EventTypeTicketModel, ids)
}

// GetEventTypeTicket gets event.type.ticket existing record.
func (c *Client) GetEventTypeTicket(id int64) (*EventTypeTicket, error) {
	etts, err := c.GetEventTypeTickets([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*etts)[0]), nil
}

// GetEventTypeTickets gets event.type.ticket existing records.
func (c *Client) GetEventTypeTickets(ids []int64) (*EventTypeTickets, error) {
	etts := &EventTypeTickets{}
	if err := c.Read(EventTypeTicketModel, ids, nil, etts); err != nil {
		return nil, err
	}
	return etts, nil
}

// FindEventTypeTicket finds event.type.ticket record by querying it with criteria.
func (c *Client) FindEventTypeTicket(criteria *Criteria) (*EventTypeTicket, error) {
	etts := &EventTypeTickets{}
	if err := c.SearchRead(EventTypeTicketModel, criteria, NewOptions().Limit(1), etts); err != nil {
		return nil, err
	}
	return &((*etts)[0]), nil
}

// FindEventTypeTickets finds event.type.ticket records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventTypeTickets(criteria *Criteria, options *Options) (*EventTypeTickets, error) {
	etts := &EventTypeTickets{}
	if err := c.SearchRead(EventTypeTicketModel, criteria, options, etts); err != nil {
		return nil, err
	}
	return etts, nil
}

// FindEventTypeTicketIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventTypeTicketIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(EventTypeTicketModel, criteria, options)
}

// FindEventTypeTicketId finds record id by querying it with criteria.
func (c *Client) FindEventTypeTicketId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventTypeTicketModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
