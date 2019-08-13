package recurly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	accountsRoot = "/accounts"
)

// AccountCode Recurly account code
type AccountCode string

func urlEncode(accountCode AccountCode) string {
	return fmt.Sprintf("code-%s", url.PathEscape(string(accountCode)))
}

// Account account
// See https://partner-docs.recurly.com/v2018-08-09#tag/account
type Account struct {
	BaseObject
	Code                 AccountCode  `json:"code"`
	State                AccountState `json:"state"`
	HostedLoginToken     string       `json:"hosted_login_token"`
	ShippingAddresses    []Address    `json:"shipping_addresses"`
	Username             *string      `json:"username,omitifempty"`
	Email                *string      `json:"email,omitifempty"`
	PreferredLocal       *string      `json:"preferred_locale,omitifempty"`
	CCEmails             *string      `json:"cc_emails,omitifempty"`
	FirstName            *string      `json:"first_name,omitifempty"`
	LastName             *string      `json:"last_name,omitifempty"`
	Company              *string      `json:"company"`
	VATNumber            *string      `json:"vat_number"`
	TaxExempt            *bool        `json:"tax_exempt"`
	ExceptionCertificate *string      `json:"exception_certificate"`
	ParentAccountII      *string      `json:"parent_account_id"`
	BillTo               *string      `json:"bill_to"`
	Address              *Address     `json:"address,omitifemtpy"`
	BillingInfo          *BillingInfo `json:"billing_info"`
	CreatedAt            time.Time    `json:"created_at"`
	UpdatedAt            time.Time    `json:"updated_at"`
	DeletedAt            *time.Time   `json:"deleted_at"`
	CustomFields         CustomFields `json:"custom_fields"`
}

type AccountList struct {
	ListMetadata
	Data []*Account `json:"data"`
}

type ListMetadata struct {
	ObjectName string `json:"object"`
	HasMore    bool   `json:"has_more"`
	Next       string `json:"next"`
}

type AccountState string

const (
	AccountStateActive = "active"
	AccountStateClosed = "closed"
)

type BaseObject struct {
	ID     ObjectID `json:"id"`
	Object string   `json:"object"`
}

type ObjectID string

type AccountListParams struct {
	BaseListParams

	Subscriber *bool
	PastDue    *bool
}

func (list *AccountListParams) Options() []KeyValue {
	// TODO: coerce list to BaseListParams and call its Options()
	var options []KeyValue

	if list.IDs != nil && len(list.IDs) > 0 {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.IDs, ",")})
	}
	if list.Limit > 0 {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(list.Limit)})
	}
	options = append(options, KeyValue{Key: "order", Value: string(list.Order)})
	options = append(options, KeyValue{Key: "sort", Value: string(list.Sort)})
	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}
	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}
	if list.Subscriber != nil {
		options = append(options, KeyValue{Key: "subscriber", Value: formatBool(*list.Subscriber)})
	}
	if list.PastDue != nil {
		options = append(options, KeyValue{Key: "past_due", Value: formatBool(*list.PastDue)})
	}

	return options
}

// ToInt64 converts an ObjectID to an int64 for more efficient storage
func (id ObjectID) ToInt64() int64 {
	intID, _ := strconv.ParseInt(string(id), 36, 64)
	return intID
}

func (account *Account) String() string {
	data, err := json.Marshal(&account)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// GetAccountByID returns an account with the given account ID
func (c *Client) GetAccountByID(accountID int64) (*Account, error) {
	account := &Account{}
	err := c.Call(http.MethodGet, fmt.Sprintf("%s/%d", accountsRoot, accountID), nil, account)
	return account, err
}

// GetAccountByAccountCode returns an account with the given account code
func (c *Client) GetAccountByAccountCode(accountCode AccountCode) (*Account, error) {
	account := &Account{}
	err := c.Call(http.MethodGet, fmt.Sprintf("%s/%s", accountsRoot, urlEncode(accountCode)), nil, account)
	return account, err
}

// ListAccounts returns an array of accounts
func (c *Client) ListAccounts(params *AccountListParams) (*AccountList, error) {
	accounts := &AccountList{}
	err := c.Call(http.MethodGet, accountsRoot, params, accounts)
	return accounts, err
}
