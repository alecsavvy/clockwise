package db

import (
	corev1 "github.com/alecsavvy/clockwise/api/oap/core/v1"
)

type DB interface {
	StoreERN(address, version, tx_hash string, msg corev1.ElectronicReleaseNotification) error
	StorePIE(address, version, tx_hash string, msg corev1.PartyIdentificationAndEnrichment) error
	StoreMEAD(address, version, tx_hash string, msg corev1.MediaEnrichmentAndDescription) error

	GetERN(address string) (corev1.ElectronicReleaseNotification, error)
	GetPIE(address string) (corev1.PartyIdentificationAndEnrichment, error)
	GetMEAD(address string) (corev1.MediaEnrichmentAndDescription, error)
}
