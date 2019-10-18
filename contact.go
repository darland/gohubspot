package gohubspot

import (
	"fmt"
	"net/url"
)

type ContactsService service

type Contact struct {
	Vid          int        `json:"vid"`
	CanonicalVid int        `json:"canonical-vid"`
	MergedVids   []int      `json:"merged-vids"`
	PortalID     int        `json:"portal-id"`
	IsContact    bool       `json:"is-contact"`
	ProfileToken string     `json:"profile-token"`
	ProfileURL   string     `json:"profile-url"`
	Properties   Properties `json:"properties"`
}

type Contacts struct {
	Contacts  []Contact `json:"contacts"`
	HasMore   bool      `json:"has-more"`
	VidOffset int       `json:"vid-offset"`
}

func (s *ContactsService) Create(properties Properties) (*IdentityProfile, error) {
	urlPath := "/contacts/v1/contact"
	res := new(IdentityProfile)
	err := s.client.RunPost(urlPath, properties, res)
	return res, err
}

func (s *ContactsService) Update(contactID int, properties Properties) error {
	urlPath := fmt.Sprintf("/contacts/v1/contact/vid/%d/profile", contactID)
	return s.client.RunPost(urlPath, properties, nil)
}

func (s *ContactsService) UpdateByEmail(email string, properties Properties) error {
	urlPath := fmt.Sprintf("/contacts/v1/contact/email/%s/profile", email)
	return s.client.RunPost(urlPath, properties, nil)
}

func (s *ContactsService) CreateOrUpdateByEmail(email string, properties Properties) (*Vid, error) {
	urlPath := fmt.Sprintf("/contacts/v1/contact/createOrUpdate/email/%s", email)

	res := new(Vid)
	err := s.client.RunPost(urlPath, properties, res)
	return res, err
}

func (s *ContactsService) DeleteById(id int) (*ContactDeleteResult, error) {
	urlPath := fmt.Sprintf("/contacts/v1/contact/vid/%d", id)

	res := new(ContactDeleteResult)
	err := s.client.RunDelete(urlPath, res)
	return res, err
}

func (s *ContactsService) DeleteByEmail(email string) (*ContactDeleteResult, error) {
	urlPath := fmt.Sprintf("/contacts/v1/contact/email/%s", email)

	res := new(ContactDeleteResult)
	err := s.client.RunDelete(urlPath, res)
	return res, err
}

func (s *ContactsService) Merge(primaryID, secondaryID int) error {
	urlPath := fmt.Sprintf("/contacts/v1/contact/merge-vids/%d/", primaryID)
	secondary := struct {
		SecondaryID int `json:"vidToMerge"`
	}{
		SecondaryID: secondaryID,
	}

	return s.client.RunPost(urlPath, secondary, nil)
}

func (s *ContactsService) GetByToken(token string) (*Contact, error) {
	urlPath := fmt.Sprintf("/contacts/v1/contact/utk/%s/profile", token)
	res := new(Contact)
	err := s.client.RunGet(urlPath, res)
	return res, err
}

func (s *ContactsService) GetAllContacts(params *url.Values) (*Contacts, error) {
	urlPath := "/contacts/v1/lists/all/contacts/all?"
	if params != nil {
		urlPath += params.Encode()
	}
	res := new(Contacts)
	err := s.client.RunGet(urlPath, res)
	return res, err
}
