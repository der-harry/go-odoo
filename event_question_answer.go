package odoo

// EventQuestionAnswer represents event.question.answer model.
type EventQuestionAnswer struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	QuestionId  *Many2One `xmlrpc:"question_id,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// EventQuestionAnswers represents array of event.question.answer model.
type EventQuestionAnswers []EventQuestionAnswer

// EventQuestionAnswerModel is the odoo model name.
const EventQuestionAnswerModel = "event.question.answer"

// Many2One convert EventQuestionAnswer to *Many2One.
func (eqa *EventQuestionAnswer) Many2One() *Many2One {
	return NewMany2One(eqa.Id.Get(), "")
}

// CreateEventQuestionAnswer creates a new event.question.answer model and returns its id.
func (c *Client) CreateEventQuestionAnswer(eqa *EventQuestionAnswer) (int64, error) {
	ids, err := c.CreateEventQuestionAnswers([]*EventQuestionAnswer{eqa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventQuestionAnswer creates a new event.question.answer model and returns its id.
func (c *Client) CreateEventQuestionAnswers(eqas []*EventQuestionAnswer) ([]int64, error) {
	var vv []interface{}
	for _, v := range eqas {
		vv = append(vv, v)
	}
	return c.Create(EventQuestionAnswerModel, vv, nil)
}

// UpdateEventQuestionAnswer updates an existing event.question.answer record.
func (c *Client) UpdateEventQuestionAnswer(eqa *EventQuestionAnswer) error {
	return c.UpdateEventQuestionAnswers([]int64{eqa.Id.Get()}, eqa)
}

// UpdateEventQuestionAnswers updates existing event.question.answer records.
// All records (represented by ids) will be updated by eqa values.
func (c *Client) UpdateEventQuestionAnswers(ids []int64, eqa *EventQuestionAnswer) error {
	return c.Update(EventQuestionAnswerModel, ids, eqa, nil)
}

// DeleteEventQuestionAnswer deletes an existing event.question.answer record.
func (c *Client) DeleteEventQuestionAnswer(id int64) error {
	return c.DeleteEventQuestionAnswers([]int64{id})
}

// DeleteEventQuestionAnswers deletes existing event.question.answer records.
func (c *Client) DeleteEventQuestionAnswers(ids []int64) error {
	return c.Delete(EventQuestionAnswerModel, ids)
}

// GetEventQuestionAnswer gets event.question.answer existing record.
func (c *Client) GetEventQuestionAnswer(id int64) (*EventQuestionAnswer, error) {
	eqas, err := c.GetEventQuestionAnswers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*eqas)[0]), nil
}

// GetEventQuestionAnswers gets event.question.answer existing records.
func (c *Client) GetEventQuestionAnswers(ids []int64) (*EventQuestionAnswers, error) {
	eqas := &EventQuestionAnswers{}
	if err := c.Read(EventQuestionAnswerModel, ids, nil, eqas); err != nil {
		return nil, err
	}
	return eqas, nil
}

// FindEventQuestionAnswer finds event.question.answer record by querying it with criteria.
func (c *Client) FindEventQuestionAnswer(criteria *Criteria) (*EventQuestionAnswer, error) {
	eqas := &EventQuestionAnswers{}
	if err := c.SearchRead(EventQuestionAnswerModel, criteria, NewOptions().Limit(1), eqas); err != nil {
		return nil, err
	}
	return &((*eqas)[0]), nil
}

// FindEventQuestionAnswers finds event.question.answer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventQuestionAnswers(criteria *Criteria, options *Options) (*EventQuestionAnswers, error) {
	eqas := &EventQuestionAnswers{}
	if err := c.SearchRead(EventQuestionAnswerModel, criteria, options, eqas); err != nil {
		return nil, err
	}
	return eqas, nil
}

// FindEventQuestionAnswerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventQuestionAnswerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(EventQuestionAnswerModel, criteria, options)
}

// FindEventQuestionAnswerId finds record id by querying it with criteria.
func (c *Client) FindEventQuestionAnswerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventQuestionAnswerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
