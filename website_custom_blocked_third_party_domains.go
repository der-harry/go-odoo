package odoo

// WebsiteCustomBlockedThirdPartyDomains represents website.custom_blocked_third_party_domains model.
type WebsiteCustomBlockedThirdPartyDomains struct {
	Content     *String   `xmlrpc:"content,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// WebsiteCustomBlockedThirdPartyDomainss represents array of website.custom_blocked_third_party_domains model.
type WebsiteCustomBlockedThirdPartyDomainss []WebsiteCustomBlockedThirdPartyDomains

// WebsiteCustomBlockedThirdPartyDomainsModel is the odoo model name.
const WebsiteCustomBlockedThirdPartyDomainsModel = "website.custom_blocked_third_party_domains"

// Many2One convert WebsiteCustomBlockedThirdPartyDomains to *Many2One.
func (wc *WebsiteCustomBlockedThirdPartyDomains) Many2One() *Many2One {
	return NewMany2One(wc.Id.Get(), "")
}

// CreateWebsiteCustomBlockedThirdPartyDomains creates a new website.custom_blocked_third_party_domains model and returns its id.
func (c *Client) CreateWebsiteCustomBlockedThirdPartyDomains(wc *WebsiteCustomBlockedThirdPartyDomains) (int64, error) {
	ids, err := c.CreateWebsiteCustomBlockedThirdPartyDomainss([]*WebsiteCustomBlockedThirdPartyDomains{wc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsiteCustomBlockedThirdPartyDomains creates a new website.custom_blocked_third_party_domains model and returns its id.
func (c *Client) CreateWebsiteCustomBlockedThirdPartyDomainss(wcs []*WebsiteCustomBlockedThirdPartyDomains) ([]int64, error) {
	var vv []interface{}
	for _, v := range wcs {
		vv = append(vv, v)
	}
	return c.Create(WebsiteCustomBlockedThirdPartyDomainsModel, vv, nil)
}

// UpdateWebsiteCustomBlockedThirdPartyDomains updates an existing website.custom_blocked_third_party_domains record.
func (c *Client) UpdateWebsiteCustomBlockedThirdPartyDomains(wc *WebsiteCustomBlockedThirdPartyDomains) error {
	return c.UpdateWebsiteCustomBlockedThirdPartyDomainss([]int64{wc.Id.Get()}, wc)
}

// UpdateWebsiteCustomBlockedThirdPartyDomainss updates existing website.custom_blocked_third_party_domains records.
// All records (represented by ids) will be updated by wc values.
func (c *Client) UpdateWebsiteCustomBlockedThirdPartyDomainss(ids []int64, wc *WebsiteCustomBlockedThirdPartyDomains) error {
	return c.Update(WebsiteCustomBlockedThirdPartyDomainsModel, ids, wc, nil)
}

// DeleteWebsiteCustomBlockedThirdPartyDomains deletes an existing website.custom_blocked_third_party_domains record.
func (c *Client) DeleteWebsiteCustomBlockedThirdPartyDomains(id int64) error {
	return c.DeleteWebsiteCustomBlockedThirdPartyDomainss([]int64{id})
}

// DeleteWebsiteCustomBlockedThirdPartyDomainss deletes existing website.custom_blocked_third_party_domains records.
func (c *Client) DeleteWebsiteCustomBlockedThirdPartyDomainss(ids []int64) error {
	return c.Delete(WebsiteCustomBlockedThirdPartyDomainsModel, ids)
}

// GetWebsiteCustomBlockedThirdPartyDomains gets website.custom_blocked_third_party_domains existing record.
func (c *Client) GetWebsiteCustomBlockedThirdPartyDomains(id int64) (*WebsiteCustomBlockedThirdPartyDomains, error) {
	wcs, err := c.GetWebsiteCustomBlockedThirdPartyDomainss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wcs)[0]), nil
}

// GetWebsiteCustomBlockedThirdPartyDomainss gets website.custom_blocked_third_party_domains existing records.
func (c *Client) GetWebsiteCustomBlockedThirdPartyDomainss(ids []int64) (*WebsiteCustomBlockedThirdPartyDomainss, error) {
	wcs := &WebsiteCustomBlockedThirdPartyDomainss{}
	if err := c.Read(WebsiteCustomBlockedThirdPartyDomainsModel, ids, nil, wcs); err != nil {
		return nil, err
	}
	return wcs, nil
}

// FindWebsiteCustomBlockedThirdPartyDomains finds website.custom_blocked_third_party_domains record by querying it with criteria.
func (c *Client) FindWebsiteCustomBlockedThirdPartyDomains(criteria *Criteria) (*WebsiteCustomBlockedThirdPartyDomains, error) {
	wcs := &WebsiteCustomBlockedThirdPartyDomainss{}
	if err := c.SearchRead(WebsiteCustomBlockedThirdPartyDomainsModel, criteria, NewOptions().Limit(1), wcs); err != nil {
		return nil, err
	}
	return &((*wcs)[0]), nil
}

// FindWebsiteCustomBlockedThirdPartyDomainss finds website.custom_blocked_third_party_domains records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteCustomBlockedThirdPartyDomainss(criteria *Criteria, options *Options) (*WebsiteCustomBlockedThirdPartyDomainss, error) {
	wcs := &WebsiteCustomBlockedThirdPartyDomainss{}
	if err := c.SearchRead(WebsiteCustomBlockedThirdPartyDomainsModel, criteria, options, wcs); err != nil {
		return nil, err
	}
	return wcs, nil
}

// FindWebsiteCustomBlockedThirdPartyDomainsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteCustomBlockedThirdPartyDomainsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebsiteCustomBlockedThirdPartyDomainsModel, criteria, options)
}

// FindWebsiteCustomBlockedThirdPartyDomainsId finds record id by querying it with criteria.
func (c *Client) FindWebsiteCustomBlockedThirdPartyDomainsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteCustomBlockedThirdPartyDomainsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
