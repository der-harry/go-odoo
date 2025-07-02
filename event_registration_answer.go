package odoo

// EventRegistrationAnswer represents event.registration.answer model.
type EventRegistrationAnswer struct {
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	EventId        *Many2One  `xmlrpc:"event_id,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	PartnerId      *Many2One  `xmlrpc:"partner_id,omitempty"`
	QuestionId     *Many2One  `xmlrpc:"question_id,omitempty"`
	QuestionType   *Selection `xmlrpc:"question_type,omitempty"`
	RegistrationId *Many2One  `xmlrpc:"registration_id,omitempty"`
	ValueAnswerId  *Many2One  `xmlrpc:"value_answer_id,omitempty"`
	ValueTextBox   *String    `xmlrpc:"value_text_box,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// EventRegistrationAnswers represents array of event.registration.answer model.
type EventRegistrationAnswers []EventRegistrationAnswer

// EventRegistrationAnswerModel is the odoo model name.
const EventRegistrationAnswerModel = "event.registration.answer"

// Many2One convert EventRegistrationAnswer to *Many2One.
func (era *EventRegistrationAnswer) Many2One() *Many2One {
	return NewMany2One(era.Id.Get(), "")
}

// CreateEventRegistrationAnswer creates a new event.registration.answer model and returns its id.
func (c *Client) CreateEventRegistrationAnswer(era *EventRegistrationAnswer) (int64, error) {
	ids, err := c.CreateEventRegistrationAnswers([]*EventRegistrationAnswer{era})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventRegistrationAnswer creates a new event.registration.answer model and returns its id.
func (c *Client) CreateEventRegistrationAnswers(eras []*EventRegistrationAnswer) ([]int64, error) {
	var vv []interface{}
	for _, v := range eras {
		vv = append(vv, v)
	}
	return c.Create(EventRegistrationAnswerModel, vv, nil)
}

// UpdateEventRegistrationAnswer updates an existing event.registration.answer record.
func (c *Client) UpdateEventRegistrationAnswer(era *EventRegistrationAnswer) error {
	return c.UpdateEventRegistrationAnswers([]int64{era.Id.Get()}, era)
}

// UpdateEventRegistrationAnswers updates existing event.registration.answer records.
// All records (represented by ids) will be updated by era values.
func (c *Client) UpdateEventRegistrationAnswers(ids []int64, era *EventRegistrationAnswer) error {
	return c.Update(EventRegistrationAnswerModel, ids, era, nil)
}

// DeleteEventRegistrationAnswer deletes an existing event.registration.answer record.
func (c *Client) DeleteEventRegistrationAnswer(id int64) error {
	return c.DeleteEventRegistrationAnswers([]int64{id})
}

// DeleteEventRegistrationAnswers deletes existing event.registration.answer records.
func (c *Client) DeleteEventRegistrationAnswers(ids []int64) error {
	return c.Delete(EventRegistrationAnswerModel, ids)
}

// GetEventRegistrationAnswer gets event.registration.answer existing record.
func (c *Client) GetEventRegistrationAnswer(id int64) (*EventRegistrationAnswer, error) {
	eras, err := c.GetEventRegistrationAnswers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*eras)[0]), nil
}

// GetEventRegistrationAnswers gets event.registration.answer existing records.
func (c *Client) GetEventRegistrationAnswers(ids []int64) (*EventRegistrationAnswers, error) {
	eras := &EventRegistrationAnswers{}
	if err := c.Read(EventRegistrationAnswerModel, ids, nil, eras); err != nil {
		return nil, err
	}
	return eras, nil
}

// FindEventRegistrationAnswer finds event.registration.answer record by querying it with criteria.
func (c *Client) FindEventRegistrationAnswer(criteria *Criteria) (*EventRegistrationAnswer, error) {
	eras := &EventRegistrationAnswers{}
	if err := c.SearchRead(EventRegistrationAnswerModel, criteria, NewOptions().Limit(1), eras); err != nil {
		return nil, err
	}
	return &((*eras)[0]), nil
}

// FindEventRegistrationAnswers finds event.registration.answer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventRegistrationAnswers(criteria *Criteria, options *Options) (*EventRegistrationAnswers, error) {
	eras := &EventRegistrationAnswers{}
	if err := c.SearchRead(EventRegistrationAnswerModel, criteria, options, eras); err != nil {
		return nil, err
	}
	return eras, nil
}

// FindEventRegistrationAnswerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventRegistrationAnswerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(EventRegistrationAnswerModel, criteria, options)
}

// FindEventRegistrationAnswerId finds record id by querying it with criteria.
func (c *Client) FindEventRegistrationAnswerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventRegistrationAnswerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
