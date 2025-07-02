package odoo

// EventStage represents event.stage model.
type EventStage struct {
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	Description   *String   `xmlrpc:"description,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Fold          *Bool     `xmlrpc:"fold,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	LegendBlocked *String   `xmlrpc:"legend_blocked,omitempty"`
	LegendDone    *String   `xmlrpc:"legend_done,omitempty"`
	LegendNormal  *String   `xmlrpc:"legend_normal,omitempty"`
	Name          *String   `xmlrpc:"name,omitempty"`
	PipeEnd       *Bool     `xmlrpc:"pipe_end,omitempty"`
	Sequence      *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// EventStages represents array of event.stage model.
type EventStages []EventStage

// EventStageModel is the odoo model name.
const EventStageModel = "event.stage"

// Many2One convert EventStage to *Many2One.
func (es *EventStage) Many2One() *Many2One {
	return NewMany2One(es.Id.Get(), "")
}

// CreateEventStage creates a new event.stage model and returns its id.
func (c *Client) CreateEventStage(es *EventStage) (int64, error) {
	ids, err := c.CreateEventStages([]*EventStage{es})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventStage creates a new event.stage model and returns its id.
func (c *Client) CreateEventStages(ess []*EventStage) ([]int64, error) {
	var vv []interface{}
	for _, v := range ess {
		vv = append(vv, v)
	}
	return c.Create(EventStageModel, vv, nil)
}

// UpdateEventStage updates an existing event.stage record.
func (c *Client) UpdateEventStage(es *EventStage) error {
	return c.UpdateEventStages([]int64{es.Id.Get()}, es)
}

// UpdateEventStages updates existing event.stage records.
// All records (represented by ids) will be updated by es values.
func (c *Client) UpdateEventStages(ids []int64, es *EventStage) error {
	return c.Update(EventStageModel, ids, es, nil)
}

// DeleteEventStage deletes an existing event.stage record.
func (c *Client) DeleteEventStage(id int64) error {
	return c.DeleteEventStages([]int64{id})
}

// DeleteEventStages deletes existing event.stage records.
func (c *Client) DeleteEventStages(ids []int64) error {
	return c.Delete(EventStageModel, ids)
}

// GetEventStage gets event.stage existing record.
func (c *Client) GetEventStage(id int64) (*EventStage, error) {
	ess, err := c.GetEventStages([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ess)[0]), nil
}

// GetEventStages gets event.stage existing records.
func (c *Client) GetEventStages(ids []int64) (*EventStages, error) {
	ess := &EventStages{}
	if err := c.Read(EventStageModel, ids, nil, ess); err != nil {
		return nil, err
	}
	return ess, nil
}

// FindEventStage finds event.stage record by querying it with criteria.
func (c *Client) FindEventStage(criteria *Criteria) (*EventStage, error) {
	ess := &EventStages{}
	if err := c.SearchRead(EventStageModel, criteria, NewOptions().Limit(1), ess); err != nil {
		return nil, err
	}
	return &((*ess)[0]), nil
}

// FindEventStages finds event.stage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventStages(criteria *Criteria, options *Options) (*EventStages, error) {
	ess := &EventStages{}
	if err := c.SearchRead(EventStageModel, criteria, options, ess); err != nil {
		return nil, err
	}
	return ess, nil
}

// FindEventStageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventStageIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(EventStageModel, criteria, options)
}

// FindEventStageId finds record id by querying it with criteria.
func (c *Client) FindEventStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
