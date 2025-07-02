package odoo

// WebsiteControllerPage represents website.controller.page model.
type WebsiteControllerPage struct {
	Active                    *Bool      `xmlrpc:"active,omitempty"`
	Arch                      *String    `xmlrpc:"arch,omitempty"`
	ArchBase                  *String    `xmlrpc:"arch_base,omitempty"`
	ArchDb                    *String    `xmlrpc:"arch_db,omitempty"`
	ArchFs                    *String    `xmlrpc:"arch_fs,omitempty"`
	ArchPrev                  *String    `xmlrpc:"arch_prev,omitempty"`
	ArchUpdated               *Bool      `xmlrpc:"arch_updated,omitempty"`
	CanPublish                *Bool      `xmlrpc:"can_publish,omitempty"`
	ControllerPageIds         *Relation  `xmlrpc:"controller_page_ids,omitempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omitempty"`
	CustomizeShow             *Bool      `xmlrpc:"customize_show,omitempty"`
	DefaultLayout             *Selection `xmlrpc:"default_layout,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty"`
	FirstPageId               *Many2One  `xmlrpc:"first_page_id,omitempty"`
	GroupsId                  *Relation  `xmlrpc:"groups_id,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty"`
	InheritChildrenIds        *Relation  `xmlrpc:"inherit_children_ids,omitempty"`
	InheritId                 *Many2One  `xmlrpc:"inherit_id,omitempty"`
	IsPublished               *Bool      `xmlrpc:"is_published,omitempty"`
	IsSeoOptimized            *Bool      `xmlrpc:"is_seo_optimized,omitempty"`
	Key                       *String    `xmlrpc:"key,omitempty"`
	MenuIds                   *Relation  `xmlrpc:"menu_ids,omitempty"`
	Mode                      *Selection `xmlrpc:"mode,omitempty"`
	Model                     *String    `xmlrpc:"model,omitempty"`
	ModelDataId               *Many2One  `xmlrpc:"model_data_id,omitempty"`
	ModelId                   *Many2One  `xmlrpc:"model_id,omitempty"`
	Name                      *String    `xmlrpc:"name,omitempty"`
	NameSlugified             *String    `xmlrpc:"name_slugified,omitempty"`
	PageIds                   *Relation  `xmlrpc:"page_ids,omitempty"`
	Priority                  *Int       `xmlrpc:"priority,omitempty"`
	RecordDomain              *String    `xmlrpc:"record_domain,omitempty"`
	RecordViewId              *Many2One  `xmlrpc:"record_view_id,omitempty"`
	SeoName                   *String    `xmlrpc:"seo_name,omitempty"`
	ThemeTemplateId           *Many2One  `xmlrpc:"theme_template_id,omitempty"`
	Track                     *Bool      `xmlrpc:"track,omitempty"`
	Type                      *Selection `xmlrpc:"type,omitempty"`
	UrlDemo                   *String    `xmlrpc:"url_demo,omitempty"`
	ViewId                    *Many2One  `xmlrpc:"view_id,omitempty"`
	Visibility                *Selection `xmlrpc:"visibility,omitempty"`
	VisibilityPassword        *String    `xmlrpc:"visibility_password,omitempty"`
	VisibilityPasswordDisplay *String    `xmlrpc:"visibility_password_display,omitempty"`
	WarningInfo               *String    `xmlrpc:"warning_info,omitempty"`
	WebsiteId                 *Many2One  `xmlrpc:"website_id,omitempty"`
	WebsiteMetaDescription    *String    `xmlrpc:"website_meta_description,omitempty"`
	WebsiteMetaKeywords       *String    `xmlrpc:"website_meta_keywords,omitempty"`
	WebsiteMetaOgImg          *String    `xmlrpc:"website_meta_og_img,omitempty"`
	WebsiteMetaTitle          *String    `xmlrpc:"website_meta_title,omitempty"`
	WebsitePublished          *Bool      `xmlrpc:"website_published,omitempty"`
	WebsiteUrl                *String    `xmlrpc:"website_url,omitempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omitempty"`
	XmlId                     *String    `xmlrpc:"xml_id,omitempty"`
}

// WebsiteControllerPages represents array of website.controller.page model.
type WebsiteControllerPages []WebsiteControllerPage

// WebsiteControllerPageModel is the odoo model name.
const WebsiteControllerPageModel = "website.controller.page"

// Many2One convert WebsiteControllerPage to *Many2One.
func (wcp *WebsiteControllerPage) Many2One() *Many2One {
	return NewMany2One(wcp.Id.Get(), "")
}

// CreateWebsiteControllerPage creates a new website.controller.page model and returns its id.
func (c *Client) CreateWebsiteControllerPage(wcp *WebsiteControllerPage) (int64, error) {
	ids, err := c.CreateWebsiteControllerPages([]*WebsiteControllerPage{wcp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsiteControllerPage creates a new website.controller.page model and returns its id.
func (c *Client) CreateWebsiteControllerPages(wcps []*WebsiteControllerPage) ([]int64, error) {
	var vv []interface{}
	for _, v := range wcps {
		vv = append(vv, v)
	}
	return c.Create(WebsiteControllerPageModel, vv, nil)
}

// UpdateWebsiteControllerPage updates an existing website.controller.page record.
func (c *Client) UpdateWebsiteControllerPage(wcp *WebsiteControllerPage) error {
	return c.UpdateWebsiteControllerPages([]int64{wcp.Id.Get()}, wcp)
}

// UpdateWebsiteControllerPages updates existing website.controller.page records.
// All records (represented by ids) will be updated by wcp values.
func (c *Client) UpdateWebsiteControllerPages(ids []int64, wcp *WebsiteControllerPage) error {
	return c.Update(WebsiteControllerPageModel, ids, wcp, nil)
}

// DeleteWebsiteControllerPage deletes an existing website.controller.page record.
func (c *Client) DeleteWebsiteControllerPage(id int64) error {
	return c.DeleteWebsiteControllerPages([]int64{id})
}

// DeleteWebsiteControllerPages deletes existing website.controller.page records.
func (c *Client) DeleteWebsiteControllerPages(ids []int64) error {
	return c.Delete(WebsiteControllerPageModel, ids)
}

// GetWebsiteControllerPage gets website.controller.page existing record.
func (c *Client) GetWebsiteControllerPage(id int64) (*WebsiteControllerPage, error) {
	wcps, err := c.GetWebsiteControllerPages([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wcps)[0]), nil
}

// GetWebsiteControllerPages gets website.controller.page existing records.
func (c *Client) GetWebsiteControllerPages(ids []int64) (*WebsiteControllerPages, error) {
	wcps := &WebsiteControllerPages{}
	if err := c.Read(WebsiteControllerPageModel, ids, nil, wcps); err != nil {
		return nil, err
	}
	return wcps, nil
}

// FindWebsiteControllerPage finds website.controller.page record by querying it with criteria.
func (c *Client) FindWebsiteControllerPage(criteria *Criteria) (*WebsiteControllerPage, error) {
	wcps := &WebsiteControllerPages{}
	if err := c.SearchRead(WebsiteControllerPageModel, criteria, NewOptions().Limit(1), wcps); err != nil {
		return nil, err
	}
	return &((*wcps)[0]), nil
}

// FindWebsiteControllerPages finds website.controller.page records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteControllerPages(criteria *Criteria, options *Options) (*WebsiteControllerPages, error) {
	wcps := &WebsiteControllerPages{}
	if err := c.SearchRead(WebsiteControllerPageModel, criteria, options, wcps); err != nil {
		return nil, err
	}
	return wcps, nil
}

// FindWebsiteControllerPageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteControllerPageIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebsiteControllerPageModel, criteria, options)
}

// FindWebsiteControllerPageId finds record id by querying it with criteria.
func (c *Client) FindWebsiteControllerPageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteControllerPageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
