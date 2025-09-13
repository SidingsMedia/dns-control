// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package dnscontrol

type Service interface {
	// <Handler>(<model> *model.<model>) error
}

type service struct {
	// Any resources go here
}

// func (service *Service) <Handler>(<model> *model.<Model>) error {
// 	// Handler logic here
// 	return nil
// }

func NewService() Service {
	return &service{
		//Resources
	}
}
