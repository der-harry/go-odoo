package odoo

// EventTag represents event.tag model.
type EventTag struct {
	CategoryId       *Many2One `xmlrpc:"category_id,omitempty"`
	CategorySequence *Int      `xmlrpc:"category_sequence,omitempty"`
	Color            *Int      `xmlrpc:"color,omitempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	Name             *String   `xmlrpc:"name,omitempty"`
	Sequence         *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omitempty"`
}

// EventTags represents array of event.tag model.
type EventTags []EventTag

// EventTagModel is the odoo model name.
const EventTagModel = "event.tag"

// Many2One convert EventTag to *Many2One.
func (et *EventTag) Many2One() *Many2One {
	return NewMany2One(et.Id.Get(), "")
}

// CreateEventTag creates a new event.tag model and returns its id.
func (c *Client) CreateEventTag(et *EventTag) (int64, error) {
	ids, err := c.CreateEventTags([]*EventTag{et})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventTag creates a new event.tag model and returns its id.
func (c *Client) CreateEventTags(ets []*EventTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range ets {
		vv = append(vv, v)
	}
	return c.Create(EventTagModel, vv, nil)
}

// UpdateEventTag updates an existing event.tag record.
func (c *Client) UpdateEventTag(et *EventTag) error {
	return c.UpdateEventTags([]int64{et.Id.Get()}, et)
}

// UpdateEventTags updates existing event.tag records.
// All records (represented by ids) will be updated by et values.
func (c *Client) UpdateEventTags(ids []int64, et *EventTag) error {
	return c.Update(EventTagModel, ids, et, nil)
}

// DeleteEventTag deletes an existing event.tag record.
func (c *Client) DeleteEventTag(id int64) error {
	return c.DeleteEventTags([]int64{id})
}

// DeleteEventTags deletes existing event.tag records.
func (c *Client) DeleteEventTags(ids []int64) error {
	return c.Delete(EventTagModel, ids)
}

// GetEventTag gets event.tag existing record.
func (c *Client) GetEventTag(id int64) (*EventTag, error) {
	ets, err := c.GetEventTags([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ets)[0]), nil
}

// GetEventTags gets event.tag existing records.
func (c *Client) GetEventTags(ids []int64) (*EventTags, error) {
	ets := &EventTags{}
	if err := c.Read(EventTagModel, ids, nil, ets); err != nil {
		return nil, err
	}
	return ets, nil
}

// FindEventTag finds event.tag record by querying it with criteria.
func (c *Client) FindEventTag(criteria *Criteria) (*EventTag, error) {
	ets := &EventTags{}
	if err := c.SearchRead(EventTagModel, criteria, NewOptions().Limit(1), ets); err != nil {
		return nil, err
	}
	return &((*ets)[0]), nil
}

// FindEventTags finds event.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventTags(criteria *Criteria, options *Options) (*EventTags, error) {
	ets := &EventTags{}
	if err := c.SearchRead(EventTagModel, criteria, options, ets); err != nil {
		return nil, err
	}
	return ets, nil
}

// FindEventTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(EventTagModel, criteria, options)
}

// FindEventTagId finds record id by querying it with criteria.
func (c *Client) FindEventTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
