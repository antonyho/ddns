package models

import "time"

// RecordType defines the type of a DNS record.
type RecordType string

const (
	A     RecordType = "A"
	AAAA  RecordType = "AAAA"
	CNAME RecordType = "CNAME"
	MX    RecordType = "MX"
	TXT   RecordType = "TXT"
	NS    RecordType = "NS"
)

// DNSRecord represents a single DNS record.
type DNSRecord struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	DomainID  uint       `json:"domain_id" gorm:"index"`
	Subdomain string     `json:"subdomain" gorm:"not null"`
	Type      RecordType `json:"type" gorm:"not null"`
	Value     string     `json:"value" gorm:"not null"`
	TTL       uint       `json:"ttl" gorm:"not null;default:60"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
