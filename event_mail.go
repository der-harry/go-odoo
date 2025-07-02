package odoo

// EventMail represents event.mail model.
type EventMail struct {
	CreateDate          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	EventId             *Many2One  `xmlrpc:"event_id,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	IntervalNbr         *Int       `xmlrpc:"interval_nbr,omitempty"`
	IntervalType        *Selection `xmlrpc:"interval_type,omitempty"`
	IntervalUnit        *Selection `xmlrpc:"interval_unit,omitempty"`
	LastRegistrationId  *Many2One  `xmlrpc:"last_registration_id,omitempty"`
	MailCountDone       *Int       `xmlrpc:"mail_count_done,omitempty"`
	MailDone            *Bool      `xmlrpc:"mail_done,omitempty"`
	MailRegistrationIds *Relation  `xmlrpc:"mail_registration_ids,omitempty"`
	MailState           *Selection `xmlrpc:"mail_state,omitempty"`
	NotificationType    *Selection `xmlrpc:"notification_type,omitempty"`
	ScheduledDate       *Time      `xmlrpc:"scheduled_date,omitempty"`
	Sequence            *Int       `xmlrpc:"sequence,omitempty"`
	TemplateRef         *String    `xmlrpc:"template_ref,omitempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// EventMails represents array of event.mail model.
type EventMails []EventMail

// EventMailModel is the odoo model name.
const EventMailModel = "event.mail"

// Many2One convert EventMail to *Many2One.
func (em *EventMail) Many2One() *Many2One {
	return NewMany2One(em.Id.Get(), "")
}

// CreateEventMail creates a new event.mail model and returns its id.
func (c *Client) CreateEventMail(em *EventMail) (int64, error) {
	ids, err := c.CreateEventMails([]*EventMail{em})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventMail creates a new event.mail model and returns its id.
func (c *Client) CreateEventMails(ems []*EventMail) ([]int64, error) {
	var vv []interface{}
	for _, v := range ems {
		vv = append(vv, v)
	}
	return c.Create(EventMailModel, vv, nil)
}

// UpdateEventMail updates an existing event.mail record.
func (c *Client) UpdateEventMail(em *EventMail) error {
	return c.UpdateEventMails([]int64{em.Id.Get()}, em)
}

// UpdateEventMails updates existing event.mail records.
// All records (represented by ids) will be updated by em values.
func (c *Client) UpdateEventMails(ids []int64, em *EventMail) error {
	return c.Update(EventMailModel, ids, em, nil)
}

// DeleteEventMail deletes an existing event.mail record.
func (c *Client) DeleteEventMail(id int64) error {
	return c.DeleteEventMails([]int64{id})
}

// DeleteEventMails deletes existing event.mail records.
func (c *Client) DeleteEventMails(ids []int64) error {
	return c.Delete(EventMailModel, ids)
}

// GetEventMail gets event.mail existing record.
func (c *Client) GetEventMail(id int64) (*EventMail, error) {
	ems, err := c.GetEventMails([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ems)[0]), nil
}

// GetEventMails gets event.mail existing records.
func (c *Client) GetEventMails(ids []int64) (*EventMails, error) {
	ems := &EventMails{}
	if err := c.Read(EventMailModel, ids, nil, ems); err != nil {
		return nil, err
	}
	return ems, nil
}

// FindEventMail finds event.mail record by querying it with criteria.
func (c *Client) FindEventMail(criteria *Criteria) (*EventMail, error) {
	ems := &EventMails{}
	if err := c.SearchRead(EventMailModel, criteria, NewOptions().Limit(1), ems); err != nil {
		return nil, err
	}
	return &((*ems)[0]), nil
}

// FindEventMails finds event.mail records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventMails(criteria *Criteria, options *Options) (*EventMails, error) {
	ems := &EventMails{}
	if err := c.SearchRead(EventMailModel, criteria, options, ems); err != nil {
		return nil, err
	}
	return ems, nil
}

// FindEventMailIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventMailIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(EventMailModel, criteria, options)
}

// FindEventMailId finds record id by querying it with criteria.
func (c *Client) FindEventMailId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventMailModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
