/*
 * Copyright (C) 2019 Intel Corporation
 * SPDX-License-Identifier: BSD-3-Clause
 */
package constants

const (
	ServiceName                   = "WLS"
	ConfigDir                     = "/etc/workload-service/"
	TrustedJWTSigningCertsDir     = ConfigDir + "jwt/"
	TrustedCaCertsDir             = ConfigDir + "cacerts/"
	TLSCertPath                   = ConfigDir + "tls-cert.pem"
	TLSKeyPath                    = ConfigDir + "tls.key"
	CertApproverGroupName         = "CertApprover"
	DefaultWlsTlsCn               = "WLS TLS Certificate"
	DefaultWlsTlsSan              = "127.0.0.1,localhost"
	DefaultWlsCertOrganization    = "INTEL"
	DefaultWlsCertCountry         = "US"
	DefaultWlsCertProvince        = "SF"
	DefaultWlsCertLocality        = "SC"
	DefaultKeyAlgorithm           = "rsa"
	DefaultKeyAlgorithmLength     = 3072
	FlavorImageRetrievalGroupName = "FlavorsImageRetrieval"
	AdministratorGroupName        = "Administrator"
	ReportCreationGroupName       = "ReportsCreate"
	BearerToken                   = "BEARER_TOKEN"
	CmsBaseUrlEnv                 = "CMS_BASE_URL"
	WlsTLsCertCnEnv               = "WLS_TLS_CERT_CN"
	WlsCertOrgEnv                 = "WLS_CERT_ORG"
	WlsCertCountryEnv             = "WLS_CERT_COUNTRY"
	WlsCertProvinceEnv            = "WLS_CERT_PROVINCE"
	WlsCertLocalityEnv            = "WLS_CERT_LOCALITY"
)

// State represents whether or not a daemon is running or not
type State bool

const (
	// Stopped is the default nil value, indicating not running
	Stopped State = false
	// Running means the daemon is active
	Running State = true
)
