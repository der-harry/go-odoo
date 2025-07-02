package odoo

// EventMailRegistration represents event.mail.registration model.
type EventMailRegistration struct {
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	MailSent       *Bool     `xmlrpc:"mail_sent,omitempty"`
	RegistrationId *Many2One `xmlrpc:"registration_id,omitempty"`
	ScheduledDate  *Time     `xmlrpc:"scheduled_date,omitempty"`
	SchedulerId    *Many2One `xmlrpc:"scheduler_id,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// EventMailRegistrations represents array of event.mail.registration model.
type EventMailRegistrations []EventMailRegistration

// EventMailRegistrationModel is the odoo model name.
const EventMailRegistrationModel = "event.mail.registration"

// Many2One convert EventMailRegistration to *Many2One.
func (emr *EventMailRegistration) Many2One() *Many2One {
	return NewMany2One(emr.Id.Get(), "")
}

// CreateEventMailRegistration creates a new event.mail.registration model and returns its id.
func (c *Client) CreateEventMailRegistration(emr *EventMailRegistration) (int64, error) {
	ids, err := c.CreateEventMailRegistrations([]*EventMailRegistration{emr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventMailRegistration creates a new event.mail.registration model and returns its id.
func (c *Client) CreateEventMailRegistrations(emrs []*EventMailRegistration) ([]int64, error) {
	var vv []interface{}
	for _, v := range emrs {
		vv = append(vv, v)
	}
	return c.Create(EventMailRegistrationModel, vv, nil)
}

// UpdateEventMailRegistration updates an existing event.mail.registration record.
func (c *Client) UpdateEventMailRegistration(emr *EventMailRegistration) error {
	return c.UpdateEventMailRegistrations([]int64{emr.Id.Get()}, emr)
}

// UpdateEventMailRegistrations updates existing event.mail.registration records.
// All records (represented by ids) will be updated by emr values.
func (c *Client) UpdateEventMailRegistrations(ids []int64, emr *EventMailRegistration) error {
	return c.Update(EventMailRegistrationModel, ids, emr, nil)
}

// DeleteEventMailRegistration deletes an existing event.mail.registration record.
func (c *Client) DeleteEventMailRegistration(id int64) error {
	return c.DeleteEventMailRegistrations([]int64{id})
}

// DeleteEventMailRegistrations deletes existing event.mail.registration records.
func (c *Client) DeleteEventMailRegistrations(ids []int64) error {
	return c.Delete(EventMailRegistrationModel, ids)
}

// GetEventMailRegistration gets event.mail.registration existing record.
func (c *Client) GetEventMailRegistration(id int64) (*EventMailRegistration, error) {
	emrs, err := c.GetEventMailRegistrations([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*emrs)[0]), nil
}

// GetEventMailRegistrations gets event.mail.registration existing records.
func (c *Client) GetEventMailRegistrations(ids []int64) (*EventMailRegistrations, error) {
	emrs := &EventMailRegistrations{}
	if err := c.Read(EventMailRegistrationModel, ids, nil, emrs); err != nil {
		return nil, err
	}
	return emrs, nil
}

// FindEventMailRegistration finds event.mail.registration record by querying it with criteria.
func (c *Client) FindEventMailRegistration(criteria *Criteria) (*EventMailRegistration, error) {
	emrs := &EventMailRegistrations{}
	if err := c.SearchRead(EventMailRegistrationModel, criteria, NewOptions().Limit(1), emrs); err != nil {
		return nil, err
	}
	return &((*emrs)[0]), nil
}

// FindEventMailRegistrations finds event.mail.registration records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventMailRegistrations(criteria *Criteria, options *Options) (*EventMailRegistrations, error) {
	emrs := &EventMailRegistrations{}
	if err := c.SearchRead(EventMailRegistrationModel, criteria, options, emrs); err != nil {
		return nil, err
	}
	return emrs, nil
}

// FindEventMailRegistrationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventMailRegistrationIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(EventMailRegistrationModel, criteria, options)
}

// FindEventMailRegistrationId finds record id by querying it with criteria.
func (c *Client) FindEventMailRegistrationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventMailRegistrationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
