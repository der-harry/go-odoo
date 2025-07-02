package odoo

// WebsiteConfiguratorFeature represents website.configurator.feature model.
type WebsiteConfiguratorFeature struct {
	CreateDate                *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One `xmlrpc:"create_uid,omitempty"`
	Description               *String   `xmlrpc:"description,omitempty"`
	DisplayName               *String   `xmlrpc:"display_name,omitempty"`
	FeatureUrl                *String   `xmlrpc:"feature_url,omitempty"`
	IapPageCode               *String   `xmlrpc:"iap_page_code,omitempty"`
	Icon                      *String   `xmlrpc:"icon,omitempty"`
	Id                        *Int      `xmlrpc:"id,omitempty"`
	MenuCompany               *Bool     `xmlrpc:"menu_company,omitempty"`
	MenuSequence              *Int      `xmlrpc:"menu_sequence,omitempty"`
	ModuleId                  *Many2One `xmlrpc:"module_id,omitempty"`
	Name                      *String   `xmlrpc:"name,omitempty"`
	PageViewId                *Many2One `xmlrpc:"page_view_id,omitempty"`
	Sequence                  *Int      `xmlrpc:"sequence,omitempty"`
	WebsiteConfigPreselection *String   `xmlrpc:"website_config_preselection,omitempty"`
	WriteDate                 *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One `xmlrpc:"write_uid,omitempty"`
}

// WebsiteConfiguratorFeatures represents array of website.configurator.feature model.
type WebsiteConfiguratorFeatures []WebsiteConfiguratorFeature

// WebsiteConfiguratorFeatureModel is the odoo model name.
const WebsiteConfiguratorFeatureModel = "website.configurator.feature"

// Many2One convert WebsiteConfiguratorFeature to *Many2One.
func (wcf *WebsiteConfiguratorFeature) Many2One() *Many2One {
	return NewMany2One(wcf.Id.Get(), "")
}

// CreateWebsiteConfiguratorFeature creates a new website.configurator.feature model and returns its id.
func (c *Client) CreateWebsiteConfiguratorFeature(wcf *WebsiteConfiguratorFeature) (int64, error) {
	ids, err := c.CreateWebsiteConfiguratorFeatures([]*WebsiteConfiguratorFeature{wcf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsiteConfiguratorFeature creates a new website.configurator.feature model and returns its id.
func (c *Client) CreateWebsiteConfiguratorFeatures(wcfs []*WebsiteConfiguratorFeature) ([]int64, error) {
	var vv []interface{}
	for _, v := range wcfs {
		vv = append(vv, v)
	}
	return c.Create(WebsiteConfiguratorFeatureModel, vv, nil)
}

// UpdateWebsiteConfiguratorFeature updates an existing website.configurator.feature record.
func (c *Client) UpdateWebsiteConfiguratorFeature(wcf *WebsiteConfiguratorFeature) error {
	return c.UpdateWebsiteConfiguratorFeatures([]int64{wcf.Id.Get()}, wcf)
}

// UpdateWebsiteConfiguratorFeatures updates existing website.configurator.feature records.
// All records (represented by ids) will be updated by wcf values.
func (c *Client) UpdateWebsiteConfiguratorFeatures(ids []int64, wcf *WebsiteConfiguratorFeature) error {
	return c.Update(WebsiteConfiguratorFeatureModel, ids, wcf, nil)
}

// DeleteWebsiteConfiguratorFeature deletes an existing website.configurator.feature record.
func (c *Client) DeleteWebsiteConfiguratorFeature(id int64) error {
	return c.DeleteWebsiteConfiguratorFeatures([]int64{id})
}

// DeleteWebsiteConfiguratorFeatures deletes existing website.configurator.feature records.
func (c *Client) DeleteWebsiteConfiguratorFeatures(ids []int64) error {
	return c.Delete(WebsiteConfiguratorFeatureModel, ids)
}

// GetWebsiteConfiguratorFeature gets website.configurator.feature existing record.
func (c *Client) GetWebsiteConfiguratorFeature(id int64) (*WebsiteConfiguratorFeature, error) {
	wcfs, err := c.GetWebsiteConfiguratorFeatures([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wcfs)[0]), nil
}

// GetWebsiteConfiguratorFeatures gets website.configurator.feature existing records.
func (c *Client) GetWebsiteConfiguratorFeatures(ids []int64) (*WebsiteConfiguratorFeatures, error) {
	wcfs := &WebsiteConfiguratorFeatures{}
	if err := c.Read(WebsiteConfiguratorFeatureModel, ids, nil, wcfs); err != nil {
		return nil, err
	}
	return wcfs, nil
}

// FindWebsiteConfiguratorFeature finds website.configurator.feature record by querying it with criteria.
func (c *Client) FindWebsiteConfiguratorFeature(criteria *Criteria) (*WebsiteConfiguratorFeature, error) {
	wcfs := &WebsiteConfiguratorFeatures{}
	if err := c.SearchRead(WebsiteConfiguratorFeatureModel, criteria, NewOptions().Limit(1), wcfs); err != nil {
		return nil, err
	}
	return &((*wcfs)[0]), nil
}

// FindWebsiteConfiguratorFeatures finds website.configurator.feature records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteConfiguratorFeatures(criteria *Criteria, options *Options) (*WebsiteConfiguratorFeatures, error) {
	wcfs := &WebsiteConfiguratorFeatures{}
	if err := c.SearchRead(WebsiteConfiguratorFeatureModel, criteria, options, wcfs); err != nil {
		return nil, err
	}
	return wcfs, nil
}

// FindWebsiteConfiguratorFeatureIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteConfiguratorFeatureIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebsiteConfiguratorFeatureModel, criteria, options)
}

// FindWebsiteConfiguratorFeatureId finds record id by querying it with criteria.
func (c *Client) FindWebsiteConfiguratorFeatureId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteConfiguratorFeatureModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
