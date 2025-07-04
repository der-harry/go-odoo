package odoo

// CalendarEvent represents calendar.event model.
type CalendarEvent struct {
	AcceptedCount            *Int       `xmlrpc:"accepted_count,omitempty"`
	AccessToken              *String    `xmlrpc:"access_token,omitempty"`
	Active                   *Bool      `xmlrpc:"active,omitempty"`
	ActivityIds              *Relation  `xmlrpc:"activity_ids,omitempty"`
	AlarmIds                 *Relation  `xmlrpc:"alarm_ids,omitempty"`
	Allday                   *Bool      `xmlrpc:"allday,omitempty"`
	AttendeeIds              *Relation  `xmlrpc:"attendee_ids,omitempty"`
	AttendeesCount           *Int       `xmlrpc:"attendees_count,omitempty"`
	AwaitingCount            *Int       `xmlrpc:"awaiting_count,omitempty"`
	Byday                    *Selection `xmlrpc:"byday,omitempty"`
	CategIds                 *Relation  `xmlrpc:"categ_ids,omitempty"`
	Count                    *Int       `xmlrpc:"count,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrentAttendee          *Many2One  `xmlrpc:"current_attendee,omitempty"`
	CurrentStatus            *Selection `xmlrpc:"current_status,omitempty"`
	Day                      *Int       `xmlrpc:"day,omitempty"`
	DeclinedCount            *Int       `xmlrpc:"declined_count,omitempty"`
	Description              *String    `xmlrpc:"description,omitempty"`
	DisplayDescription       *Bool      `xmlrpc:"display_description,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	DisplayTime              *String    `xmlrpc:"display_time,omitempty"`
	Duration                 *Float     `xmlrpc:"duration,omitempty"`
	EndType                  *Selection `xmlrpc:"end_type,omitempty"`
	EventTz                  *Selection `xmlrpc:"event_tz,omitempty"`
	FollowRecurrence         *Bool      `xmlrpc:"follow_recurrence,omitempty"`
	Fri                      *Bool      `xmlrpc:"fri,omitempty"`
	HasMessage               *Bool      `xmlrpc:"has_message,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	Interval                 *Int       `xmlrpc:"interval,omitempty"`
	InvalidEmailPartnerIds   *Relation  `xmlrpc:"invalid_email_partner_ids,omitempty"`
	IsHighlighted            *Bool      `xmlrpc:"is_highlighted,omitempty"`
	IsOrganizerAlone         *Bool      `xmlrpc:"is_organizer_alone,omitempty"`
	Location                 *String    `xmlrpc:"location,omitempty"`
	MessageAttachmentCount   *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError          *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	Mon                      *Bool      `xmlrpc:"mon,omitempty"`
	MonthBy                  *Selection `xmlrpc:"month_by,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	OpportunityId            *Many2One  `xmlrpc:"opportunity_id,omitempty"`
	PartnerId                *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerIds               *Relation  `xmlrpc:"partner_ids,omitempty"`
	Privacy                  *Selection `xmlrpc:"privacy,omitempty"`
	RecurrenceId             *Many2One  `xmlrpc:"recurrence_id,omitempty"`
	RecurrenceUpdate         *Selection `xmlrpc:"recurrence_update,omitempty"`
	Recurrency               *Bool      `xmlrpc:"recurrency,omitempty"`
	ResId                    *Many2One  `xmlrpc:"res_id,omitempty"`
	ResModel                 *String    `xmlrpc:"res_model,omitempty"`
	ResModelId               *Many2One  `xmlrpc:"res_model_id,omitempty"`
	ResModelName             *String    `xmlrpc:"res_model_name,omitempty"`
	Rrule                    *String    `xmlrpc:"rrule,omitempty"`
	RruleType                *Selection `xmlrpc:"rrule_type,omitempty"`
	RruleTypeUi              *Selection `xmlrpc:"rrule_type_ui,omitempty"`
	Sat                      *Bool      `xmlrpc:"sat,omitempty"`
	ShouldShowStatus         *Bool      `xmlrpc:"should_show_status,omitempty"`
	ShowAs                   *Selection `xmlrpc:"show_as,omitempty"`
	Start                    *Time      `xmlrpc:"start,omitempty"`
	StartDate                *Time      `xmlrpc:"start_date,omitempty"`
	Stop                     *Time      `xmlrpc:"stop,omitempty"`
	StopDate                 *Time      `xmlrpc:"stop_date,omitempty"`
	Sun                      *Bool      `xmlrpc:"sun,omitempty"`
	TentativeCount           *Int       `xmlrpc:"tentative_count,omitempty"`
	Thu                      *Bool      `xmlrpc:"thu,omitempty"`
	Tue                      *Bool      `xmlrpc:"tue,omitempty"`
	Until                    *Time      `xmlrpc:"until,omitempty"`
	UserCanEdit              *Bool      `xmlrpc:"user_can_edit,omitempty"`
	UserId                   *Many2One  `xmlrpc:"user_id,omitempty"`
	VideocallChannelId       *Many2One  `xmlrpc:"videocall_channel_id,omitempty"`
	VideocallLocation        *String    `xmlrpc:"videocall_location,omitempty"`
	VideocallSource          *Selection `xmlrpc:"videocall_source,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty"`
	Wed                      *Bool      `xmlrpc:"wed,omitempty"`
	Weekday                  *Selection `xmlrpc:"weekday,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// CalendarEvents represents array of calendar.event model.
type CalendarEvents []CalendarEvent

// CalendarEventModel is the odoo model name.
const CalendarEventModel = "calendar.event"

// Many2One convert CalendarEvent to *Many2One.
func (ce *CalendarEvent) Many2One() *Many2One {
	return NewMany2One(ce.Id.Get(), "")
}

// CreateCalendarEvent creates a new calendar.event model and returns its id.
func (c *Client) CreateCalendarEvent(ce *CalendarEvent) (int64, error) {
	ids, err := c.CreateCalendarEvents([]*CalendarEvent{ce})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCalendarEvent creates a new calendar.event model and returns its id.
func (c *Client) CreateCalendarEvents(ces []*CalendarEvent) ([]int64, error) {
	var vv []interface{}
	for _, v := range ces {
		vv = append(vv, v)
	}
	return c.Create(CalendarEventModel, vv, nil)
}

// UpdateCalendarEvent updates an existing calendar.event record.
func (c *Client) UpdateCalendarEvent(ce *CalendarEvent) error {
	return c.UpdateCalendarEvents([]int64{ce.Id.Get()}, ce)
}

// UpdateCalendarEvents updates existing calendar.event records.
// All records (represented by ids) will be updated by ce values.
func (c *Client) UpdateCalendarEvents(ids []int64, ce *CalendarEvent) error {
	return c.Update(CalendarEventModel, ids, ce, nil)
}

// DeleteCalendarEvent deletes an existing calendar.event record.
func (c *Client) DeleteCalendarEvent(id int64) error {
	return c.DeleteCalendarEvents([]int64{id})
}

// DeleteCalendarEvents deletes existing calendar.event records.
func (c *Client) DeleteCalendarEvents(ids []int64) error {
	return c.Delete(CalendarEventModel, ids)
}

// GetCalendarEvent gets calendar.event existing record.
func (c *Client) GetCalendarEvent(id int64) (*CalendarEvent, error) {
	ces, err := c.GetCalendarEvents([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ces)[0]), nil
}

// GetCalendarEvents gets calendar.event existing records.
func (c *Client) GetCalendarEvents(ids []int64) (*CalendarEvents, error) {
	ces := &CalendarEvents{}
	if err := c.Read(CalendarEventModel, ids, nil, ces); err != nil {
		return nil, err
	}
	return ces, nil
}

// FindCalendarEvent finds calendar.event record by querying it with criteria.
func (c *Client) FindCalendarEvent(criteria *Criteria) (*CalendarEvent, error) {
	ces := &CalendarEvents{}
	if err := c.SearchRead(CalendarEventModel, criteria, NewOptions().Limit(1), ces); err != nil {
		return nil, err
	}
	return &((*ces)[0]), nil
}

// FindCalendarEvents finds calendar.event records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarEvents(criteria *Criteria, options *Options) (*CalendarEvents, error) {
	ces := &CalendarEvents{}
	if err := c.SearchRead(CalendarEventModel, criteria, options, ces); err != nil {
		return nil, err
	}
	return ces, nil
}

// FindCalendarEventIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarEventIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CalendarEventModel, criteria, options)
}

// FindCalendarEventId finds record id by querying it with criteria.
func (c *Client) FindCalendarEventId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarEventModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
