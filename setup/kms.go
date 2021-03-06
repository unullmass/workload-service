/*
 * Copyright (C) 2019 Intel Corporation
 * SPDX-License-Identifier: BSD-3-Clause
 */
package setup

import (
	"errors"
	"fmt"
	csetup "intel/isecl/lib/common/setup"
	"intel/isecl/workload-service/config"
	"strings"
)

// KMSConnection is a setup task for setting up the connection to the Key Management Server (KMS)
type KMSConnection struct{}

// Run will run the KMS Connection setup task, but will skip if Validate() returns no errors
func (kms KMSConnection) Run(c csetup.Context) error {
	if kms.Validate(c) == nil {
		fmt.Println("KMS connection already setup, skipping ...")
		return nil
	}
	fmt.Println("Setting up KMS connection ...")
	var err error
	var kmsURL string
	if kmsURL, err = c.GetenvString(config.KMS_URL, "Key Management Server URL"); err != nil {
		return err
	}
	if strings.HasSuffix(kmsURL, "/") {
		config.Configuration.KMS.URL = kmsURL
	} else {
		config.Configuration.KMS.URL = kmsURL + "/"
	}
	if config.Configuration.KMS.User, err = c.GetenvString(config.KMS_USER, "Key Management Server User"); err != nil {
		return err
	}
	if config.Configuration.KMS.Password, err = c.GetenvSecret(config.KMS_PASSWORD, "Key Management Server Password"); err != nil {
		return err
	}

	return config.Save()
}

// Validate checks whether or not the KMS Connection setup task was completed successfully
func (kms KMSConnection) Validate(c csetup.Context) error {
	k := &config.Configuration.KMS
	if k.URL == "" {
		return errors.New("KMS Connection: URL is not set")
	}
	if k.User == "" {
		return errors.New("KMS Connection: User is not set")
	}
	if k.Password == "" {
		return errors.New("KMS Connection: Password is not set ")
	}
	return nil
}
