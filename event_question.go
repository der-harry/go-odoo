package odoo

// EventQuestion represents event.question model.
type EventQuestion struct {
	AnswerIds         *Relation  `xmlrpc:"answer_ids,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	EventId           *Many2One  `xmlrpc:"event_id,omitempty"`
	EventTypeId       *Many2One  `xmlrpc:"event_type_id,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	IsMandatoryAnswer *Bool      `xmlrpc:"is_mandatory_answer,omitempty"`
	OncePerOrder      *Bool      `xmlrpc:"once_per_order,omitempty"`
	QuestionType      *Selection `xmlrpc:"question_type,omitempty"`
	Sequence          *Int       `xmlrpc:"sequence,omitempty"`
	Title             *String    `xmlrpc:"title,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// EventQuestions represents array of event.question model.
type EventQuestions []EventQuestion

// EventQuestionModel is the odoo model name.
const EventQuestionModel = "event.question"

// Many2One convert EventQuestion to *Many2One.
func (eq *EventQuestion) Many2One() *Many2One {
	return NewMany2One(eq.Id.Get(), "")
}

// CreateEventQuestion creates a new event.question model and returns its id.
func (c *Client) CreateEventQuestion(eq *EventQuestion) (int64, error) {
	ids, err := c.CreateEventQuestions([]*EventQuestion{eq})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventQuestion creates a new event.question model and returns its id.
func (c *Client) CreateEventQuestions(eqs []*EventQuestion) ([]int64, error) {
	var vv []interface{}
	for _, v := range eqs {
		vv = append(vv, v)
	}
	return c.Create(EventQuestionModel, vv, nil)
}

// UpdateEventQuestion updates an existing event.question record.
func (c *Client) UpdateEventQuestion(eq *EventQuestion) error {
	return c.UpdateEventQuestions([]int64{eq.Id.Get()}, eq)
}

// UpdateEventQuestions updates existing event.question records.
// All records (represented by ids) will be updated by eq values.
func (c *Client) UpdateEventQuestions(ids []int64, eq *EventQuestion) error {
	return c.Update(EventQuestionModel, ids, eq, nil)
}

// DeleteEventQuestion deletes an existing event.question record.
func (c *Client) DeleteEventQuestion(id int64) error {
	return c.DeleteEventQuestions([]int64{id})
}

// DeleteEventQuestions deletes existing event.question records.
func (c *Client) DeleteEventQuestions(ids []int64) error {
	return c.Delete(EventQuestionModel, ids)
}

// GetEventQuestion gets event.question existing record.
func (c *Client) GetEventQuestion(id int64) (*EventQuestion, error) {
	eqs, err := c.GetEventQuestions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*eqs)[0]), nil
}

// GetEventQuestions gets event.question existing records.
func (c *Client) GetEventQuestions(ids []int64) (*EventQuestions, error) {
	eqs := &EventQuestions{}
	if err := c.Read(EventQuestionModel, ids, nil, eqs); err != nil {
		return nil, err
	}
	return eqs, nil
}

// FindEventQuestion finds event.question record by querying it with criteria.
func (c *Client) FindEventQuestion(criteria *Criteria) (*EventQuestion, error) {
	eqs := &EventQuestions{}
	if err := c.SearchRead(EventQuestionModel, criteria, NewOptions().Limit(1), eqs); err != nil {
		return nil, err
	}
	return &((*eqs)[0]), nil
}

// FindEventQuestions finds event.question records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventQuestions(criteria *Criteria, options *Options) (*EventQuestions, error) {
	eqs := &EventQuestions{}
	if err := c.SearchRead(EventQuestionModel, criteria, options, eqs); err != nil {
		return nil, err
	}
	return eqs, nil
}

// FindEventQuestionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventQuestionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(EventQuestionModel, criteria, options)
}

// FindEventQuestionId finds record id by querying it with criteria.
func (c *Client) FindEventQuestionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventQuestionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
