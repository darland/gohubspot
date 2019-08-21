package gohubspot

import "fmt"

type EmailSubscriptionService service

func (s *EmailSubscriptionService) Update(email string, properties interface{}) error {
	url := fmt.Sprintf("/email/public/v1/subscriptions/%s", email)
	return s.client.RunPut(url, properties, nil)
}
