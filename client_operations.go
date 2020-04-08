package recurly

import (
	"net/http"
	"time"
	//  "net/url"
)

const (
	// APIVersion is the current Recurly API Version
	APIVersion = "v2019-10-10"
)

type ListSitesParams struct {
	Params

	Ids   *[]string
	Limit *int
	Order *string
	Sort  *string
}

func (list *ListSitesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List sites
// Returns: A list of sites.
func (c *Client) ListSites(params *ListSitesParams) *SiteList {
	path := "/sites"
	path = BuildUrl(path, params)
	return &SiteList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// Fetch a site
// Returns: A site.
func (c *Client) GetSite(siteId string) (*Site, error) {
	path := c.InterpolatePath("/sites/{site_id}", siteId)
	result := &Site{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

type ListAccountsParams struct {
	Params

	Ids        *[]string
	Limit      *int
	Order      *string
	Sort       *string
	BeginTime  *time.Time
	EndTime    *time.Time
	Email      *string
	Subscriber *bool
	PastDue    *string
}

func (list *ListAccountsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List a site's accounts
// Returns: A list of the site's accounts.
func (c *Client) ListAccounts(params *ListAccountsParams) *AccountList {
	path := "/accounts"
	path = BuildUrl(path, params)
	return &AccountList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// Create an account
// Returns: An account.
func (c *Client) CreateAccount(body *AccountCreate) (*Account, error) {
	path := c.InterpolatePath("/accounts")
	result := &Account{}
	err := c.Call(http.MethodPost, path, body, result)
	return result, err
}

// Fetch an account
// Returns: An account.
func (c *Client) GetAccount(accountId string) (*Account, error) {
	path := c.InterpolatePath("/accounts/{account_id}", accountId)
	result := &Account{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

// Modify an account
// Returns: An account.
func (c *Client) UpdateAccount(accountId string, body *AccountUpdate) (*Account, error) {
	path := c.InterpolatePath("/accounts/{account_id}", accountId)
	result := &Account{}
	err := c.Call(http.MethodPut, path, body, result)
	return result, err
}

// Deactivate an account
// Returns: An account.
func (c *Client) DeactivateAccount(accountId string) (*Account, error) {
	path := c.InterpolatePath("/accounts/{account_id}", accountId)
	result := &Account{}
	err := c.Call(http.MethodDelete, path, nil, result)
	return result, err
}

// Fetch an account's acquisition data
// Returns: An account's acquisition data.
func (c *Client) GetAccountAcquisition(accountId string) (*AccountAcquisition, error) {
	path := c.InterpolatePath("/accounts/{account_id}/acquisition", accountId)
	result := &AccountAcquisition{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

// Update an account's acquisition data
// Returns: An account's updated acquisition data.
func (c *Client) UpdateAccountAcquisition(accountId string, body *AccountAcquisitionUpdatable) (*AccountAcquisition, error) {
	path := c.InterpolatePath("/accounts/{account_id}/acquisition", accountId)
	result := &AccountAcquisition{}
	err := c.Call(http.MethodPut, path, body, result)
	return result, err
}

// Remove an account's acquisition data
// Returns: Acquisition data was succesfully deleted.
func (c *Client) RemoveAccountAcquisition(accountId string) (*Empty, error) {
	path := c.InterpolatePath("/accounts/{account_id}/acquisition", accountId)
	result := &Empty{}
	err := c.Call(http.MethodDelete, path, nil, result)
	return result, err
}

// Reactivate an inactive account
// Returns: An account.
func (c *Client) ReactivateAccount(accountId string) (*Account, error) {
	path := c.InterpolatePath("/accounts/{account_id}/reactivate", accountId)
	result := &Account{}
	err := c.Call(http.MethodPut, path, nil, result)
	return result, err
}

// Fetch an account's balance and past due status
// Returns: An account's balance.
func (c *Client) GetAccountBalance(accountId string) (*AccountBalance, error) {
	path := c.InterpolatePath("/accounts/{account_id}/balance", accountId)
	result := &AccountBalance{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

// Fetch an account's billing information
// Returns: An account's billing information.
func (c *Client) GetBillingInfo(accountId string) (*BillingInfo, error) {
	path := c.InterpolatePath("/accounts/{account_id}/billing_info", accountId)
	result := &BillingInfo{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

// Set an account's billing information
// Returns: Updated billing information.
func (c *Client) UpdateBillingInfo(accountId string, body *BillingInfoCreate) (*BillingInfo, error) {
	path := c.InterpolatePath("/accounts/{account_id}/billing_info", accountId)
	result := &BillingInfo{}
	err := c.Call(http.MethodPut, path, body, result)
	return result, err
}

// Remove an account's billing information
// Returns: Billing information deleted
func (c *Client) RemoveBillingInfo(accountId string) (*Empty, error) {
	path := c.InterpolatePath("/accounts/{account_id}/billing_info", accountId)
	result := &Empty{}
	err := c.Call(http.MethodDelete, path, nil, result)
	return result, err
}

type ListAccountCouponRedemptionsParams struct {
	Params

	Ids       *[]string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
}

func (list *ListAccountCouponRedemptionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// Show the coupon redemptions for an account
// Returns: A list of the the coupon redemptions on an account.
func (c *Client) ListAccountCouponRedemptions(accountId string, params *ListAccountCouponRedemptionsParams) *CouponRedemptionList {
	path := c.InterpolatePath("/accounts/{account_id}/coupon_redemptions", accountId)
	path = BuildUrl(path, params)
	return &CouponRedemptionList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// Show the coupon redemption that is active on an account
// Returns: An active coupon redemption on an account.
func (c *Client) GetActiveCouponRedemption(accountId string) (*CouponRedemption, error) {
	path := c.InterpolatePath("/accounts/{account_id}/coupon_redemptions/active", accountId)
	result := &CouponRedemption{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

// Generate an active coupon redemption on an account
// Returns: Returns the new coupon redemption.
func (c *Client) CreateCouponRedemption(accountId string, body *CouponRedemptionCreate) (*CouponRedemption, error) {
	path := c.InterpolatePath("/accounts/{account_id}/coupon_redemptions/active", accountId)
	result := &CouponRedemption{}
	err := c.Call(http.MethodPost, path, body, result)
	return result, err
}

// Delete the active coupon redemption from an account
// Returns: Coupon redemption deleted.
func (c *Client) RemoveCouponRedemption(accountId string) (*CouponRedemption, error) {
	path := c.InterpolatePath("/accounts/{account_id}/coupon_redemptions/active", accountId)
	result := &CouponRedemption{}
	err := c.Call(http.MethodDelete, path, nil, result)
	return result, err
}

type ListAccountCreditPaymentsParams struct {
	Params

	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
}

func (list *ListAccountCreditPaymentsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List an account's credit payments
// Returns: A list of the account's credit payments.
func (c *Client) ListAccountCreditPayments(accountId string, params *ListAccountCreditPaymentsParams) *CreditPaymentList {
	path := c.InterpolatePath("/accounts/{account_id}/credit_payments", accountId)
	path = BuildUrl(path, params)
	return &CreditPaymentList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

type ListAccountInvoicesParams struct {
	Params

	Ids       *[]string
	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
	Type      *string
}

func (list *ListAccountInvoicesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List an account's invoices
// Returns: A list of the account's invoices.
func (c *Client) ListAccountInvoices(accountId string, params *ListAccountInvoicesParams) *InvoiceList {
	path := c.InterpolatePath("/accounts/{account_id}/invoices", accountId)
	path = BuildUrl(path, params)
	return &InvoiceList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// Create an invoice for pending line items
// Returns: Returns the new invoices.
func (c *Client) CreateInvoice(accountId string, body *InvoiceCreate) (*InvoiceCollection, error) {
	path := c.InterpolatePath("/accounts/{account_id}/invoices", accountId)
	result := &InvoiceCollection{}
	err := c.Call(http.MethodPost, path, body, result)
	return result, err
}

// Preview new invoice for pending line items
// Returns: Returns the invoice previews.
func (c *Client) PreviewInvoice(accountId string, body *InvoiceCreate) (*InvoiceCollection, error) {
	path := c.InterpolatePath("/accounts/{account_id}/invoices/preview", accountId)
	result := &InvoiceCollection{}
	err := c.Call(http.MethodPost, path, body, result)
	return result, err
}

type ListAccountLineItemsParams struct {
	Params

	Ids       *[]string
	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
	Original  *string
	State     *string
	Type      *string
}

func (list *ListAccountLineItemsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List an account's line items
// Returns: A list of the account's line items.
func (c *Client) ListAccountLineItems(accountId string, params *ListAccountLineItemsParams) *LineItemList {
	path := c.InterpolatePath("/accounts/{account_id}/line_items", accountId)
	path = BuildUrl(path, params)
	return &LineItemList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// Create a new line item for the account
// Returns: Returns the new line item.
func (c *Client) CreateLineItem(accountId string, body *LineItemCreate) (*LineItem, error) {
	path := c.InterpolatePath("/accounts/{account_id}/line_items", accountId)
	result := &LineItem{}
	err := c.Call(http.MethodPost, path, body, result)
	return result, err
}

type ListAccountNotesParams struct {
	Params

	Ids *[]string
}

func (list *ListAccountNotesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// Fetch a list of an account's notes
// Returns: A list of an account's notes.
func (c *Client) ListAccountNotes(accountId string, params *ListAccountNotesParams) *AccountNoteList {
	path := c.InterpolatePath("/accounts/{account_id}/notes", accountId)
	path = BuildUrl(path, params)
	return &AccountNoteList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// Fetch an account note
// Returns: An account note.
func (c *Client) GetAccountNote(accountId string, accountNoteId string) (*AccountNote, error) {
	path := c.InterpolatePath("/accounts/{account_id}/notes/{account_note_id}", accountId, accountNoteId)
	result := &AccountNote{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

type ListShippingAddressesParams struct {
	Params

	Ids       *[]string
	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
}

func (list *ListShippingAddressesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// Fetch a list of an account's shipping addresses
// Returns: A list of an account's shipping addresses.
func (c *Client) ListShippingAddresses(accountId string, params *ListShippingAddressesParams) *ShippingAddressList {
	path := c.InterpolatePath("/accounts/{account_id}/shipping_addresses", accountId)
	path = BuildUrl(path, params)
	return &ShippingAddressList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// Create a new shipping address for the account
// Returns: Returns the new shipping address.
func (c *Client) CreateShippingAddress(accountId string, body *ShippingAddressCreate) (*ShippingAddress, error) {
	path := c.InterpolatePath("/accounts/{account_id}/shipping_addresses", accountId)
	result := &ShippingAddress{}
	err := c.Call(http.MethodPost, path, body, result)
	return result, err
}

// Fetch an account's shipping address
// Returns: A shipping address.
func (c *Client) GetShippingAddress(accountId string, shippingAddressId string) (*ShippingAddress, error) {
	path := c.InterpolatePath("/accounts/{account_id}/shipping_addresses/{shipping_address_id}", accountId, shippingAddressId)
	result := &ShippingAddress{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

// Update an account's shipping address
// Returns: The updated shipping address.
func (c *Client) UpdateShippingAddress(accountId string, shippingAddressId string, body *ShippingAddressUpdate) (*ShippingAddress, error) {
	path := c.InterpolatePath("/accounts/{account_id}/shipping_addresses/{shipping_address_id}", accountId, shippingAddressId)
	result := &ShippingAddress{}
	err := c.Call(http.MethodPut, path, body, result)
	return result, err
}

// Remove an account's shipping address
// Returns: Shipping address deleted.
func (c *Client) RemoveShippingAddress(accountId string, shippingAddressId string) (*Empty, error) {
	path := c.InterpolatePath("/accounts/{account_id}/shipping_addresses/{shipping_address_id}", accountId, shippingAddressId)
	result := &Empty{}
	err := c.Call(http.MethodDelete, path, nil, result)
	return result, err
}

type ListAccountSubscriptionsParams struct {
	Params

	Ids       *[]string
	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
	State     *string
}

func (list *ListAccountSubscriptionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List an account's subscriptions
// Returns: A list of the account's subscriptions.
func (c *Client) ListAccountSubscriptions(accountId string, params *ListAccountSubscriptionsParams) *SubscriptionList {
	path := c.InterpolatePath("/accounts/{account_id}/subscriptions", accountId)
	path = BuildUrl(path, params)
	return &SubscriptionList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

type ListAccountTransactionsParams struct {
	Params

	Ids       *[]string
	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
	Type      *string
	Success   *string
}

func (list *ListAccountTransactionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List an account's transactions
// Returns: A list of the account's transactions.
func (c *Client) ListAccountTransactions(accountId string, params *ListAccountTransactionsParams) *TransactionList {
	path := c.InterpolatePath("/accounts/{account_id}/transactions", accountId)
	path = BuildUrl(path, params)
	return &TransactionList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

type ListChildAccountsParams struct {
	Params

	Ids        *[]string
	Limit      *int
	Order      *string
	Sort       *string
	BeginTime  *time.Time
	EndTime    *time.Time
	Email      *string
	Subscriber *bool
	PastDue    *string
}

func (list *ListChildAccountsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List an account's child accounts
// Returns: A list of an account's child accounts.
func (c *Client) ListChildAccounts(accountId string, params *ListChildAccountsParams) *AccountList {
	path := c.InterpolatePath("/accounts/{account_id}/accounts", accountId)
	path = BuildUrl(path, params)
	return &AccountList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

type ListAccountAcquisitionParams struct {
	Params

	Ids       *[]string
	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
}

func (list *ListAccountAcquisitionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List a site's account acquisition data
// Returns: A list of the site's account acquisition data.
func (c *Client) ListAccountAcquisition(params *ListAccountAcquisitionParams) *AccountAcquisitionList {
	path := "/acquisitions"
	path = BuildUrl(path, params)
	return &AccountAcquisitionList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

type ListCouponsParams struct {
	Params

	Ids       *[]string
	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
}

func (list *ListCouponsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List a site's coupons
// Returns: A list of the site's coupons.
func (c *Client) ListCoupons(params *ListCouponsParams) *CouponList {
	path := "/coupons"
	path = BuildUrl(path, params)
	return &CouponList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// Create a new coupon
// Returns: A new coupon.
func (c *Client) CreateCoupon(body *CouponCreate) (*Coupon, error) {
	path := c.InterpolatePath("/coupons")
	result := &Coupon{}
	err := c.Call(http.MethodPost, path, body, result)
	return result, err
}

// Fetch a coupon
// Returns: A coupon.
func (c *Client) GetCoupon(couponId string) (*Coupon, error) {
	path := c.InterpolatePath("/coupons/{coupon_id}", couponId)
	result := &Coupon{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

// Update an active coupon
// Returns: The updated coupon.
func (c *Client) UpdateCoupon(couponId string, body *CouponUpdate) (*Coupon, error) {
	path := c.InterpolatePath("/coupons/{coupon_id}", couponId)
	result := &Coupon{}
	err := c.Call(http.MethodPut, path, body, result)
	return result, err
}

// Expire a coupon
// Returns: The expired Coupon
func (c *Client) DeactivateCoupon(couponId string) (*Coupon, error) {
	path := c.InterpolatePath("/coupons/{coupon_id}", couponId)
	result := &Coupon{}
	err := c.Call(http.MethodDelete, path, nil, result)
	return result, err
}

type ListUniqueCouponCodesParams struct {
	Params

	Ids       *[]string
	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
}

func (list *ListUniqueCouponCodesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List unique coupon codes associated with a bulk coupon
// Returns: A list of unique coupon codes that were generated
func (c *Client) ListUniqueCouponCodes(couponId string, params *ListUniqueCouponCodesParams) *UniqueCouponCodeList {
	path := c.InterpolatePath("/coupons/{coupon_id}/unique_coupon_codes", couponId)
	path = BuildUrl(path, params)
	return &UniqueCouponCodeList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

type ListCreditPaymentsParams struct {
	Params

	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
}

func (list *ListCreditPaymentsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List a site's credit payments
// Returns: A list of the site's credit payments.
func (c *Client) ListCreditPayments(params *ListCreditPaymentsParams) *CreditPaymentList {
	path := "/credit_payments"
	path = BuildUrl(path, params)
	return &CreditPaymentList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// Fetch a credit payment
// Returns: A credit payment.
func (c *Client) GetCreditPayment(creditPaymentId string) (*CreditPayment, error) {
	path := c.InterpolatePath("/credit_payments/{credit_payment_id}", creditPaymentId)
	result := &CreditPayment{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

type ListCustomFieldDefinitionsParams struct {
	Params

	Ids         *[]string
	Limit       *int
	Order       *string
	Sort        *string
	BeginTime   *time.Time
	EndTime     *time.Time
	RelatedType *string
}

func (list *ListCustomFieldDefinitionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List a site's custom field definitions
// Returns: A list of the site's custom field definitions.
func (c *Client) ListCustomFieldDefinitions(params *ListCustomFieldDefinitionsParams) *CustomFieldDefinitionList {
	path := "/custom_field_definitions"
	path = BuildUrl(path, params)
	return &CustomFieldDefinitionList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// Fetch an custom field definition
// Returns: An custom field definition.
func (c *Client) GetCustomFieldDefinition(customFieldDefinitionId string) (*CustomFieldDefinition, error) {
	path := c.InterpolatePath("/custom_field_definitions/{custom_field_definition_id}", customFieldDefinitionId)
	result := &CustomFieldDefinition{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

type ListItemsParams struct {
	Params

	Ids       *[]string
	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
	State     *string
}

func (list *ListItemsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List a site's items
// Returns: A list of the site's items.
func (c *Client) ListItems(params *ListItemsParams) *ItemList {
	path := "/items"
	path = BuildUrl(path, params)
	return &ItemList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// Create a new item
// Returns: A new item.
func (c *Client) CreateItem(body *ItemCreate) (*Item, error) {
	path := c.InterpolatePath("/items")
	result := &Item{}
	err := c.Call(http.MethodPost, path, body, result)
	return result, err
}

// Fetch an item
// Returns: An item.
func (c *Client) GetItem(itemId string) (*Item, error) {
	path := c.InterpolatePath("/items/{item_id}", itemId)
	result := &Item{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

// Update an active item
// Returns: The updated item.
func (c *Client) UpdateItem(itemId string, body *ItemUpdate) (*Item, error) {
	path := c.InterpolatePath("/items/{item_id}", itemId)
	result := &Item{}
	err := c.Call(http.MethodPut, path, body, result)
	return result, err
}

// Deactivate an item
// Returns: An item.
func (c *Client) DeactivateItem(itemId string) (*Item, error) {
	path := c.InterpolatePath("/items/{item_id}", itemId)
	result := &Item{}
	err := c.Call(http.MethodDelete, path, nil, result)
	return result, err
}

// Reactivate an inactive item
// Returns: An item.
func (c *Client) ReactivateItem(itemId string) (*Item, error) {
	path := c.InterpolatePath("/items/{item_id}/reactivate", itemId)
	result := &Item{}
	err := c.Call(http.MethodPut, path, nil, result)
	return result, err
}

type ListInvoicesParams struct {
	Params

	Ids       *[]string
	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
	Type      *string
}

func (list *ListInvoicesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List a site's invoices
// Returns: A list of the site's invoices.
func (c *Client) ListInvoices(params *ListInvoicesParams) *InvoiceList {
	path := "/invoices"
	path = BuildUrl(path, params)
	return &InvoiceList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// Fetch an invoice
// Returns: An invoice.
func (c *Client) GetInvoice(invoiceId string) (*Invoice, error) {
	path := c.InterpolatePath("/invoices/{invoice_id}", invoiceId)
	result := &Invoice{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

// Update an invoice
// Returns: An invoice.
func (c *Client) PutInvoice(invoiceId string, body *InvoiceUpdatable) (*Invoice, error) {
	path := c.InterpolatePath("/invoices/{invoice_id}", invoiceId)
	result := &Invoice{}
	err := c.Call(http.MethodPut, path, body, result)
	return result, err
}

// Fetch an invoice as a PDF
// Returns: An invoice as a PDF.
func (c *Client) GetInvoicePdf(invoiceId string) (*BinaryFile, error) {
	path := c.InterpolatePath("/invoices/{invoice_id}.pdf", invoiceId)
	result := &BinaryFile{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

type CollectInvoiceParams struct {
	Params

	Body *InvoiceCollect
}

func (list *CollectInvoiceParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// Collect a pending or past due, automatic invoice
// Returns: The updated invoice.
func (c *Client) CollectInvoice(invoiceId string, params *CollectInvoiceParams) (*Invoice, error) {
	path := c.InterpolatePath("/invoices/{invoice_id}/collect", invoiceId)
	result := &Invoice{}
	err := c.Call(http.MethodPut, path, params, result)
	return result, err
}

// Mark an open invoice as failed
// Returns: The updated invoice.
func (c *Client) FailInvoice(invoiceId string) (*Invoice, error) {
	path := c.InterpolatePath("/invoices/{invoice_id}/mark_failed", invoiceId)
	result := &Invoice{}
	err := c.Call(http.MethodPut, path, nil, result)
	return result, err
}

// Mark an open invoice as successful
// Returns: The updated invoice.
func (c *Client) MarkInvoiceSuccessful(invoiceId string) (*Invoice, error) {
	path := c.InterpolatePath("/invoices/{invoice_id}/mark_successful", invoiceId)
	result := &Invoice{}
	err := c.Call(http.MethodPut, path, nil, result)
	return result, err
}

// Reopen a closed, manual invoice
// Returns: The updated invoice.
func (c *Client) ReopenInvoice(invoiceId string) (*Invoice, error) {
	path := c.InterpolatePath("/invoices/{invoice_id}/reopen", invoiceId)
	result := &Invoice{}
	err := c.Call(http.MethodPut, path, nil, result)
	return result, err
}

// Void a credit invoice.
// Returns: The updated invoice.
func (c *Client) VoidInvoice(invoiceId string) (*Invoice, error) {
	path := c.InterpolatePath("/invoices/{invoice_id}/void", invoiceId)
	result := &Invoice{}
	err := c.Call(http.MethodPut, path, nil, result)
	return result, err
}

type ListInvoiceLineItemsParams struct {
	Params

	Ids       *[]string
	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
	Original  *string
	State     *string
	Type      *string
}

func (list *ListInvoiceLineItemsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List an invoice's line items
// Returns: A list of the invoice's line items.
func (c *Client) ListInvoiceLineItems(invoiceId string, params *ListInvoiceLineItemsParams) *LineItemList {
	path := c.InterpolatePath("/invoices/{invoice_id}/line_items", invoiceId)
	path = BuildUrl(path, params)
	return &LineItemList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

type ListInvoiceCouponRedemptionsParams struct {
	Params

	Ids       *[]string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
}

func (list *ListInvoiceCouponRedemptionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// Show the coupon redemptions applied to an invoice
// Returns: A list of the the coupon redemptions associated with the invoice.
func (c *Client) ListInvoiceCouponRedemptions(invoiceId string, params *ListInvoiceCouponRedemptionsParams) *CouponRedemptionList {
	path := c.InterpolatePath("/invoices/{invoice_id}/coupon_redemptions", invoiceId)
	path = BuildUrl(path, params)
	return &CouponRedemptionList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// List an invoice's related credit or charge invoices
// Returns: A list of the credit or charge invoices associated with the invoice.
func (c *Client) ListRelatedInvoices(invoiceId string) *InvoiceList {
	path := c.InterpolatePath("/invoices/{invoice_id}/related_invoices", invoiceId)
	return &InvoiceList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// Refund an invoice
// Returns: Returns the new credit invoice.
func (c *Client) RefundInvoice(invoiceId string, body *InvoiceRefund) (*Invoice, error) {
	path := c.InterpolatePath("/invoices/{invoice_id}/refund", invoiceId)
	result := &Invoice{}
	err := c.Call(http.MethodPost, path, body, result)
	return result, err
}

type ListLineItemsParams struct {
	Params

	Ids       *[]string
	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
	Original  *string
	State     *string
	Type      *string
}

func (list *ListLineItemsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List a site's line items
// Returns: A list of the site's line items.
func (c *Client) ListLineItems(params *ListLineItemsParams) *LineItemList {
	path := "/line_items"
	path = BuildUrl(path, params)
	return &LineItemList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// Fetch a line item
// Returns: A line item.
func (c *Client) GetLineItem(lineItemId string) (*LineItem, error) {
	path := c.InterpolatePath("/line_items/{line_item_id}", lineItemId)
	result := &LineItem{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

// Delete an uninvoiced line item
// Returns: Line item deleted.
func (c *Client) RemoveLineItem(lineItemId string) (*Empty, error) {
	path := c.InterpolatePath("/line_items/{line_item_id}", lineItemId)
	result := &Empty{}
	err := c.Call(http.MethodDelete, path, nil, result)
	return result, err
}

type ListPlansParams struct {
	Params

	Ids       *[]string
	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
	State     *string
}

func (list *ListPlansParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List a site's plans
// Returns: A list of plans.
func (c *Client) ListPlans(params *ListPlansParams) *PlanList {
	path := "/plans"
	path = BuildUrl(path, params)
	return &PlanList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// Create a plan
// Returns: A plan.
func (c *Client) CreatePlan(body *PlanCreate) (*Plan, error) {
	path := c.InterpolatePath("/plans")
	result := &Plan{}
	err := c.Call(http.MethodPost, path, body, result)
	return result, err
}

// Fetch a plan
// Returns: A plan.
func (c *Client) GetPlan(planId string) (*Plan, error) {
	path := c.InterpolatePath("/plans/{plan_id}", planId)
	result := &Plan{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

// Update a plan
// Returns: A plan.
func (c *Client) UpdatePlan(planId string, body *PlanUpdate) (*Plan, error) {
	path := c.InterpolatePath("/plans/{plan_id}", planId)
	result := &Plan{}
	err := c.Call(http.MethodPut, path, body, result)
	return result, err
}

// Remove a plan
// Returns: Plan deleted
func (c *Client) RemovePlan(planId string) (*Plan, error) {
	path := c.InterpolatePath("/plans/{plan_id}", planId)
	result := &Plan{}
	err := c.Call(http.MethodDelete, path, nil, result)
	return result, err
}

type ListPlanAddOnsParams struct {
	Params

	Ids       *[]string
	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
	State     *string
}

func (list *ListPlanAddOnsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List a plan's add-ons
// Returns: A list of add-ons.
func (c *Client) ListPlanAddOns(planId string, params *ListPlanAddOnsParams) *AddOnList {
	path := c.InterpolatePath("/plans/{plan_id}/add_ons", planId)
	path = BuildUrl(path, params)
	return &AddOnList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// Create an add-on
// Returns: An add-on.
func (c *Client) CreatePlanAddOn(planId string, body *AddOnCreate) (*AddOn, error) {
	path := c.InterpolatePath("/plans/{plan_id}/add_ons", planId)
	result := &AddOn{}
	err := c.Call(http.MethodPost, path, body, result)
	return result, err
}

// Fetch a plan's add-on
// Returns: An add-on.
func (c *Client) GetPlanAddOn(planId string, addOnId string) (*AddOn, error) {
	path := c.InterpolatePath("/plans/{plan_id}/add_ons/{add_on_id}", planId, addOnId)
	result := &AddOn{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

// Update an add-on
// Returns: An add-on.
func (c *Client) UpdatePlanAddOn(planId string, addOnId string, body *AddOnUpdate) (*AddOn, error) {
	path := c.InterpolatePath("/plans/{plan_id}/add_ons/{add_on_id}", planId, addOnId)
	result := &AddOn{}
	err := c.Call(http.MethodPut, path, body, result)
	return result, err
}

// Remove an add-on
// Returns: Add-on deleted
func (c *Client) RemovePlanAddOn(planId string, addOnId string) (*AddOn, error) {
	path := c.InterpolatePath("/plans/{plan_id}/add_ons/{add_on_id}", planId, addOnId)
	result := &AddOn{}
	err := c.Call(http.MethodDelete, path, nil, result)
	return result, err
}

type ListAddOnsParams struct {
	Params

	Ids       *[]string
	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
	State     *string
}

func (list *ListAddOnsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List a site's add-ons
// Returns: A list of add-ons.
func (c *Client) ListAddOns(params *ListAddOnsParams) *AddOnList {
	path := "/add_ons"
	path = BuildUrl(path, params)
	return &AddOnList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// Fetch an add-on
// Returns: An add-on.
func (c *Client) GetAddOn(addOnId string) (*AddOn, error) {
	path := c.InterpolatePath("/add_ons/{add_on_id}", addOnId)
	result := &AddOn{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

type ListShippingMethodsParams struct {
	Params

	Ids       *[]string
	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
}

func (list *ListShippingMethodsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List a site's shipping methods
// Returns: A list of the site's shipping methods.
func (c *Client) ListShippingMethods(params *ListShippingMethodsParams) *ShippingMethodList {
	path := "/shipping_methods"
	path = BuildUrl(path, params)
	return &ShippingMethodList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// Fetch a shipping method
// Returns: A shipping_method.
func (c *Client) GetShippingMethod(id string) (*ShippingMethod, error) {
	path := c.InterpolatePath("/shipping_methods/{id}", id)
	result := &ShippingMethod{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

type ListSubscriptionsParams struct {
	Params

	Ids       *[]string
	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
	State     *string
}

func (list *ListSubscriptionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List a site's subscriptions
// Returns: A list of the site's subscriptions.
func (c *Client) ListSubscriptions(params *ListSubscriptionsParams) *SubscriptionList {
	path := "/subscriptions"
	path = BuildUrl(path, params)
	return &SubscriptionList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// Create a new subscription
// Returns: A subscription.
func (c *Client) CreateSubscription(body *SubscriptionCreate) (*Subscription, error) {
	path := c.InterpolatePath("/subscriptions")
	result := &Subscription{}
	err := c.Call(http.MethodPost, path, body, result)
	return result, err
}

// Fetch a subscription
// Returns: A subscription.
func (c *Client) GetSubscription(subscriptionId string) (*Subscription, error) {
	path := c.InterpolatePath("/subscriptions/{subscription_id}", subscriptionId)
	result := &Subscription{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

// Modify a subscription
// Returns: A subscription.
func (c *Client) ModifySubscription(subscriptionId string, body *SubscriptionUpdate) (*Subscription, error) {
	path := c.InterpolatePath("/subscriptions/{subscription_id}", subscriptionId)
	result := &Subscription{}
	err := c.Call(http.MethodPut, path, body, result)
	return result, err
}

type TerminateSubscriptionParams struct {
	Params

	Refund *string
}

func (list *TerminateSubscriptionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// Terminate a subscription
// Returns: An expired subscription.
func (c *Client) TerminateSubscription(subscriptionId string, params *TerminateSubscriptionParams) (*Subscription, error) {
	path := c.InterpolatePath("/subscriptions/{subscription_id}", subscriptionId)
	result := &Subscription{}
	err := c.Call(http.MethodDelete, path, params, result)
	return result, err
}

type CancelSubscriptionParams struct {
	Params

	Body *SubscriptionCancel
}

func (list *CancelSubscriptionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// Cancel a subscription
// Returns: A canceled or failed subscription.
func (c *Client) CancelSubscription(subscriptionId string, params *CancelSubscriptionParams) (*Subscription, error) {
	path := c.InterpolatePath("/subscriptions/{subscription_id}/cancel", subscriptionId)
	result := &Subscription{}
	err := c.Call(http.MethodPut, path, params, result)
	return result, err
}

// Reactivate a canceled subscription
// Returns: An active subscription.
func (c *Client) ReactivateSubscription(subscriptionId string) (*Subscription, error) {
	path := c.InterpolatePath("/subscriptions/{subscription_id}/reactivate", subscriptionId)
	result := &Subscription{}
	err := c.Call(http.MethodPut, path, nil, result)
	return result, err
}

// Pause subscription
// Returns: A subscription.
func (c *Client) PauseSubscription(subscriptionId string, body *SubscriptionPause) (*Subscription, error) {
	path := c.InterpolatePath("/subscriptions/{subscription_id}/pause", subscriptionId)
	result := &Subscription{}
	err := c.Call(http.MethodPut, path, body, result)
	return result, err
}

// Resume subscription
// Returns: A subscription.
func (c *Client) ResumeSubscription(subscriptionId string) (*Subscription, error) {
	path := c.InterpolatePath("/subscriptions/{subscription_id}/resume", subscriptionId)
	result := &Subscription{}
	err := c.Call(http.MethodPut, path, nil, result)
	return result, err
}

// Convert trial subscription
// Returns: A subscription.
func (c *Client) ConvertTrial(subscriptionId string) (*Subscription, error) {
	path := c.InterpolatePath("/subscriptions/{subscription_id}/convert_trial", subscriptionId)
	result := &Subscription{}
	err := c.Call(http.MethodPut, path, nil, result)
	return result, err
}

// Fetch a subscription's pending change
// Returns: A subscription's pending change.
func (c *Client) GetSubscriptionChange(subscriptionId string) (*SubscriptionChange, error) {
	path := c.InterpolatePath("/subscriptions/{subscription_id}/change", subscriptionId)
	result := &SubscriptionChange{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

// Create a new subscription change
// Returns: A subscription change.
func (c *Client) CreateSubscriptionChange(subscriptionId string, body *SubscriptionChangeCreate) (*SubscriptionChange, error) {
	path := c.InterpolatePath("/subscriptions/{subscription_id}/change", subscriptionId)
	result := &SubscriptionChange{}
	err := c.Call(http.MethodPost, path, body, result)
	return result, err
}

// Delete the pending subscription change
// Returns: Subscription change was deleted.
func (c *Client) RemoveSubscriptionChange(subscriptionId string) (*Empty, error) {
	path := c.InterpolatePath("/subscriptions/{subscription_id}/change", subscriptionId)
	result := &Empty{}
	err := c.Call(http.MethodDelete, path, nil, result)
	return result, err
}

type ListSubscriptionInvoicesParams struct {
	Params

	Ids       *[]string
	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
	Type      *string
}

func (list *ListSubscriptionInvoicesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List a subscription's invoices
// Returns: A list of the subscription's invoices.
func (c *Client) ListSubscriptionInvoices(subscriptionId string, params *ListSubscriptionInvoicesParams) *InvoiceList {
	path := c.InterpolatePath("/subscriptions/{subscription_id}/invoices", subscriptionId)
	path = BuildUrl(path, params)
	return &InvoiceList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

type ListSubscriptionLineItemsParams struct {
	Params

	Ids       *[]string
	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
	Original  *string
	State     *string
	Type      *string
}

func (list *ListSubscriptionLineItemsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List a subscription's line items
// Returns: A list of the subscription's line items.
func (c *Client) ListSubscriptionLineItems(subscriptionId string, params *ListSubscriptionLineItemsParams) *LineItemList {
	path := c.InterpolatePath("/subscriptions/{subscription_id}/line_items", subscriptionId)
	path = BuildUrl(path, params)
	return &LineItemList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

type ListSubscriptionCouponRedemptionsParams struct {
	Params

	Ids       *[]string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
}

func (list *ListSubscriptionCouponRedemptionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// Show the coupon redemptions for a subscription
// Returns: A list of the the coupon redemptions on a subscription.
func (c *Client) ListSubscriptionCouponRedemptions(subscriptionId string, params *ListSubscriptionCouponRedemptionsParams) *CouponRedemptionList {
	path := c.InterpolatePath("/subscriptions/{subscription_id}/coupon_redemptions", subscriptionId)
	path = BuildUrl(path, params)
	return &CouponRedemptionList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

type ListTransactionsParams struct {
	Params

	Ids       *[]string
	Limit     *int
	Order     *string
	Sort      *string
	BeginTime *time.Time
	EndTime   *time.Time
	Type      *string
	Success   *string
}

func (list *ListTransactionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// List a site's transactions
// Returns: A list of the site's transactions.
func (c *Client) ListTransactions(params *ListTransactionsParams) *TransactionList {
	path := "/transactions"
	path = BuildUrl(path, params)
	return &TransactionList{
		client:       c,
		nextPagePath: path,
		HasMore:      true,
	}
}

// Fetch a transaction
// Returns: A transaction.
func (c *Client) GetTransaction(transactionId string) (*Transaction, error) {
	path := c.InterpolatePath("/transactions/{transaction_id}", transactionId)
	result := &Transaction{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

// Fetch a unique coupon code
// Returns: A unique coupon code.
func (c *Client) GetUniqueCouponCode(uniqueCouponCodeId string) (*UniqueCouponCode, error) {
	path := c.InterpolatePath("/unique_coupon_codes/{unique_coupon_code_id}", uniqueCouponCodeId)
	result := &UniqueCouponCode{}
	err := c.Call(http.MethodGet, path, nil, result)
	return result, err
}

// Deactivate a unique coupon code
// Returns: A unique coupon code.
func (c *Client) DeactivateUniqueCouponCode(uniqueCouponCodeId string) (*UniqueCouponCode, error) {
	path := c.InterpolatePath("/unique_coupon_codes/{unique_coupon_code_id}", uniqueCouponCodeId)
	result := &UniqueCouponCode{}
	err := c.Call(http.MethodDelete, path, nil, result)
	return result, err
}

// Restore a unique coupon code
// Returns: A unique coupon code.
func (c *Client) ReactivateUniqueCouponCode(uniqueCouponCodeId string) (*UniqueCouponCode, error) {
	path := c.InterpolatePath("/unique_coupon_codes/{unique_coupon_code_id}/restore", uniqueCouponCodeId)
	result := &UniqueCouponCode{}
	err := c.Call(http.MethodPut, path, nil, result)
	return result, err
}

// Create a new purchase
// Returns: Returns the new invoices
func (c *Client) CreatePurchase(body *PurchaseCreate) (*InvoiceCollection, error) {
	path := c.InterpolatePath("/purchases")
	result := &InvoiceCollection{}
	err := c.Call(http.MethodPost, path, body, result)
	return result, err
}

// Preview a new purchase
// Returns: Returns preview of the new invoices
func (c *Client) PreviewPurchase(body *PurchaseCreate) (*InvoiceCollection, error) {
	path := c.InterpolatePath("/purchases/preview")
	result := &InvoiceCollection{}
	err := c.Call(http.MethodPost, path, body, result)
	return result, err
}
