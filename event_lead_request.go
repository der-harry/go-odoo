package odoo

// EventLeadRequest represents event.lead.request model.
type EventLeadRequest struct {
	DisplayName             *String   `xmlrpc:"display_name,omitempty"`
	EventId                 *Many2One `xmlrpc:"event_id,omitempty"`
	Id                      *Int      `xmlrpc:"id,omitempty"`
	ProcessedRegistrationId *Int      `xmlrpc:"processed_registration_id,omitempty"`
}

// EventLeadRequests represents array of event.lead.request model.
type EventLeadRequests []EventLeadRequest

// EventLeadRequestModel is the odoo model name.
const EventLeadRequestModel = "event.lead.request"

// Many2One convert EventLeadRequest to *Many2One.
func (elr *EventLeadRequest) Many2One() *Many2One {
	return NewMany2One(elr.Id.Get(), "")
}

// CreateEventLeadRequest creates a new event.lead.request model and returns its id.
func (c *Client) CreateEventLeadRequest(elr *EventLeadRequest) (int64, error) {
	ids, err := c.CreateEventLeadRequests([]*EventLeadRequest{elr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventLeadRequest creates a new event.lead.request model and returns its id.
func (c *Client) CreateEventLeadRequests(elrs []*EventLeadRequest) ([]int64, error) {
	var vv []interface{}
	for _, v := range elrs {
		vv = append(vv, v)
	}
	return c.Create(EventLeadRequestModel, vv, nil)
}

// UpdateEventLeadRequest updates an existing event.lead.request record.
func (c *Client) UpdateEventLeadRequest(elr *EventLeadRequest) error {
	return c.UpdateEventLeadRequests([]int64{elr.Id.Get()}, elr)
}

// UpdateEventLeadRequests updates existing event.lead.request records.
// All records (represented by ids) will be updated by elr values.
func (c *Client) UpdateEventLeadRequests(ids []int64, elr *EventLeadRequest) error {
	return c.Update(EventLeadRequestModel, ids, elr, nil)
}

// DeleteEventLeadRequest deletes an existing event.lead.request record.
func (c *Client) DeleteEventLeadRequest(id int64) error {
	return c.DeleteEventLeadRequests([]int64{id})
}

// DeleteEventLeadRequests deletes existing event.lead.request records.
func (c *Client) DeleteEventLeadRequests(ids []int64) error {
	return c.Delete(EventLeadRequestModel, ids)
}

// GetEventLeadRequest gets event.lead.request existing record.
func (c *Client) GetEventLeadRequest(id int64) (*EventLeadRequest, error) {
	elrs, err := c.GetEventLeadRequests([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*elrs)[0]), nil
}

// GetEventLeadRequests gets event.lead.request existing records.
func (c *Client) GetEventLeadRequests(ids []int64) (*EventLeadRequests, error) {
	elrs := &EventLeadRequests{}
	if err := c.Read(EventLeadRequestModel, ids, nil, elrs); err != nil {
		return nil, err
	}
	return elrs, nil
}

// FindEventLeadRequest finds event.lead.request record by querying it with criteria.
func (c *Client) FindEventLeadRequest(criteria *Criteria) (*EventLeadRequest, error) {
	elrs := &EventLeadRequests{}
	if err := c.SearchRead(EventLeadRequestModel, criteria, NewOptions().Limit(1), elrs); err != nil {
		return nil, err
	}
	return &((*elrs)[0]), nil
}

// FindEventLeadRequests finds event.lead.request records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventLeadRequests(criteria *Criteria, options *Options) (*EventLeadRequests, error) {
	elrs := &EventLeadRequests{}
	if err := c.SearchRead(EventLeadRequestModel, criteria, options, elrs); err != nil {
		return nil, err
	}
	return elrs, nil
}

// FindEventLeadRequestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventLeadRequestIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(EventLeadRequestModel, criteria, options)
}

// FindEventLeadRequestId finds record id by querying it with criteria.
func (c *Client) FindEventLeadRequestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventLeadRequestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
