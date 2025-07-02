package odoo

// EventEventTicket represents event.event.ticket model.
type EventEventTicket struct {
	Color             *String   `xmlrpc:"color,omitempty"`
	CompanyId         *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId        *Many2One `xmlrpc:"currency_id,omitempty"`
	Description       *String   `xmlrpc:"description,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	EndSaleDatetime   *Time     `xmlrpc:"end_sale_datetime,omitempty"`
	EventId           *Many2One `xmlrpc:"event_id,omitempty"`
	EventTypeId       *Many2One `xmlrpc:"event_type_id,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	IsExpired         *Bool     `xmlrpc:"is_expired,omitempty"`
	IsLaunched        *Bool     `xmlrpc:"is_launched,omitempty"`
	IsSoldOut         *Bool     `xmlrpc:"is_sold_out,omitempty"`
	Name              *String   `xmlrpc:"name,omitempty"`
	Price             *Float    `xmlrpc:"price,omitempty"`
	PriceIncl         *Float    `xmlrpc:"price_incl,omitempty"`
	PriceReduce       *Float    `xmlrpc:"price_reduce,omitempty"`
	PriceReduceTaxinc *Float    `xmlrpc:"price_reduce_taxinc,omitempty"`
	ProductId         *Many2One `xmlrpc:"product_id,omitempty"`
	RegistrationIds   *Relation `xmlrpc:"registration_ids,omitempty"`
	SaleAvailable     *Bool     `xmlrpc:"sale_available,omitempty"`
	SeatsAvailable    *Int      `xmlrpc:"seats_available,omitempty"`
	SeatsLimited      *Bool     `xmlrpc:"seats_limited,omitempty"`
	SeatsMax          *Int      `xmlrpc:"seats_max,omitempty"`
	SeatsReserved     *Int      `xmlrpc:"seats_reserved,omitempty"`
	SeatsTaken        *Int      `xmlrpc:"seats_taken,omitempty"`
	SeatsUsed         *Int      `xmlrpc:"seats_used,omitempty"`
	Sequence          *Int      `xmlrpc:"sequence,omitempty"`
	StartSaleDatetime *Time     `xmlrpc:"start_sale_datetime,omitempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omitempty"`
}

// EventEventTickets represents array of event.event.ticket model.
type EventEventTickets []EventEventTicket

// EventEventTicketModel is the odoo model name.
const EventEventTicketModel = "event.event.ticket"

// Many2One convert EventEventTicket to *Many2One.
func (eet *EventEventTicket) Many2One() *Many2One {
	return NewMany2One(eet.Id.Get(), "")
}

// CreateEventEventTicket creates a new event.event.ticket model and returns its id.
func (c *Client) CreateEventEventTicket(eet *EventEventTicket) (int64, error) {
	ids, err := c.CreateEventEventTickets([]*EventEventTicket{eet})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventEventTicket creates a new event.event.ticket model and returns its id.
func (c *Client) CreateEventEventTickets(eets []*EventEventTicket) ([]int64, error) {
	var vv []interface{}
	for _, v := range eets {
		vv = append(vv, v)
	}
	return c.Create(EventEventTicketModel, vv, nil)
}

// UpdateEventEventTicket updates an existing event.event.ticket record.
func (c *Client) UpdateEventEventTicket(eet *EventEventTicket) error {
	return c.UpdateEventEventTickets([]int64{eet.Id.Get()}, eet)
}

// UpdateEventEventTickets updates existing event.event.ticket records.
// All records (represented by ids) will be updated by eet values.
func (c *Client) UpdateEventEventTickets(ids []int64, eet *EventEventTicket) error {
	return c.Update(EventEventTicketModel, ids, eet, nil)
}

// DeleteEventEventTicket deletes an existing event.event.ticket record.
func (c *Client) DeleteEventEventTicket(id int64) error {
	return c.DeleteEventEventTickets([]int64{id})
}

// DeleteEventEventTickets deletes existing event.event.ticket records.
func (c *Client) DeleteEventEventTickets(ids []int64) error {
	return c.Delete(EventEventTicketModel, ids)
}

// GetEventEventTicket gets event.event.ticket existing record.
func (c *Client) GetEventEventTicket(id int64) (*EventEventTicket, error) {
	eets, err := c.GetEventEventTickets([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*eets)[0]), nil
}

// GetEventEventTickets gets event.event.ticket existing records.
func (c *Client) GetEventEventTickets(ids []int64) (*EventEventTickets, error) {
	eets := &EventEventTickets{}
	if err := c.Read(EventEventTicketModel, ids, nil, eets); err != nil {
		return nil, err
	}
	return eets, nil
}

// FindEventEventTicket finds event.event.ticket record by querying it with criteria.
func (c *Client) FindEventEventTicket(criteria *Criteria) (*EventEventTicket, error) {
	eets := &EventEventTickets{}
	if err := c.SearchRead(EventEventTicketModel, criteria, NewOptions().Limit(1), eets); err != nil {
		return nil, err
	}
	return &((*eets)[0]), nil
}

// FindEventEventTickets finds event.event.ticket records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventEventTickets(criteria *Criteria, options *Options) (*EventEventTickets, error) {
	eets := &EventEventTickets{}
	if err := c.SearchRead(EventEventTicketModel, criteria, options, eets); err != nil {
		return nil, err
	}
	return eets, nil
}

// FindEventEventTicketIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventEventTicketIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(EventEventTicketModel, criteria, options)
}

// FindEventEventTicketId finds record id by querying it with criteria.
func (c *Client) FindEventEventTicketId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventEventTicketModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
