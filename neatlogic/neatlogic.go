// Package neatlogic provides a Go SDK for interacting with the NeatLogic API.
// It includes functionality for authentication, CMDB entity management, and other
// NeatLogic-specific operations.
package neatlogic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/hejingwen098/neatapi/auth"
	"github.com/hejingwen098/neatapi/common"
)

// LRequest represents the login request structure for NeatLogic authentication.
type LRequest struct {
	// UserID is the user identifier for authentication.
	UserID string `json:"userid"`
	// Password is the encrypted password for authentication.
	Password string `json:"password"`
}

// LResponse represents the response structure for NeatLogic login operations.
type LResponse struct {
	// Status indicates the operation status (OK or ERROR).
	Status string `json:"Status"`
	// Message provides additional information about the operation.
	Message string `json:"Message"`
	// JwtToken is the authentication token for subsequent API calls.
	JwtToken string `json:"JwtToken"`
}

// NeatClient represents a client for interacting with the NeatLogic API.
type NeatClient struct {
	// Client is the HTTP client used for making API requests.
	Client *http.Client
	// NeatlogicUri is the base URL for the NeatLogic API.
	NeatlogicUri string
	// JwtToken is the authentication token for API requests.
	JwtToken string
}

// CRequestBody represents the request body structure for searching CMDB entities with filters.
type CRequestBody struct {
	// PageSize specifies the number of items per page in the response.
	PageSize int `json:"pageSize"`
	// CiId is the configuration item ID for the search.
	CiId int `json:"ciId"`
	// NeedAction indicates if action information is needed.
	NeedAction bool `json:"needAction"`
	// NeedExpand indicates if expanded information is needed.
	NeedExpand bool `json:"needExpand"`
	// NeedActionType indicates if action type information is needed.
	NeedActionType bool `json:"needActionType"`
	// NeedCheck indicates if check information is needed.
	NeedCheck bool `json:"needCheck"`
	// Mode specifies the search mode.
	Mode string `json:"mode"`
	// CurrentPage indicates the current page number for pagination.
	CurrentPage int `json:"currentPage"`
	// AttrFilterList contains attribute filters for the search.
	AttrFilterList []map[string]interface{} `json:"attrFilterList"`
	// GlobalAttrFilterList contains global attribute filters for the search.
	GlobalAttrFilterList []map[string]interface{} `json:"globalAttrFilterList"`
	// RelFilterList contains relationship filters for the search.
	RelFilterList []map[string]interface{} `json:"relFilterList"`
	// Keyword is the search keyword for filtering results.
	Keyword string `json:"keyword"`
}

// CRequest represents the request structure for CMDB entity operations.
type CRequest struct {
	// CiId is the configuration item ID.
	CiId int64 `json:"ciId"`
	// CiEntityId is the configuration item entity ID.
	CiEntityId int64 `json:"ciEntityId"`
	// Keyword is the search keyword.
	Keyword string `json:"keyword"`
	// PageSize specifies the number of items per page.
	PageSize int `json:"pageSize"`
	// CurrentPage indicates the current page number for pagination.
	CurrentPage int `json:"currentPage"`
	// LimitRelEntity indicates if relationship entities should be limited.
	LimitRelEntity bool `json:"limitRelEntity"`
	// LimitAttrEntity indicates if attribute entities should be limited.
	LimitAttrEntity bool `json:"limitAttrEntity"`
}

// CResponse represents the response structure for CMDB entity search operations.
type CResponse struct {
	// Status indicates the operation status (OK or ERROR).
	Status string `json:"Status"`
	// CReturn contains the actual return data.
	CReturn CReturn `json:"Return"`
	// TimeCost is the time cost of the operation in milliseconds.
	TimeCost int64 `json:"TimeCost"`
}

// CReturn represents the return data structure for CMDB operations.
type CReturn struct {
	// Name is the name of the configuration item.
	Name string `json:"name"`
	// ID is the identifier of the configuration item.
	ID int64 `json:"id"`
	// PageCount is the total number of pages in the result set.
	PageCount int `json:"pageCount"`
	// RowNum is the total number of rows in the result set.
	RowNum int `json:"rowNum"`
	// PageSize is the number of items per page.
	PageSize int `json:"pageSize"`
	// CurrentPage indicates the current page number.
	CurrentPage int `json:"currentPage"`
	// TbodyList contains the list of CMDB entities in the current page.
	TbodyList []TbodyList `json:"tbodyList"`
}

// GetCientityResponse represents the response structure for getting a single CMDB entity.
type GetCientityResponse struct {
	// Status indicates the operation status (OK or ERROR).
	Status string `json:"Status"`
	// GetcientityReturn contains the retrieved CMDB entity data.
	GetcientityReturn TbodyList `json:"Return"`
	// TimeCost is the time cost of the operation in milliseconds.
	TimeCost int64 `json:"TimeCost"`
}

// TbodyList represents the structure of a single CMDB entity in the response.
type TbodyList struct {
	// CiIcon is the icon for the configuration item.
	CiIcon string `json:"ciIcon"`
	// GlobalAttrEntityData contains global attribute entity data.
	GlobalAttrEntityData map[string]interface{} `json:"globalAttrEntityData"`
	// TypeName is the type name of the configuration item.
	TypeName string `json:"typeName"`
	// Type is the type identifier of the configuration item.
	Type int64 `json:"type"`
	// InspectStatus is the inspection status of the entity.
	InspectStatus string `json:"inspectStatus"`
	// UUID is the universally unique identifier of the entity.
	UUID string `json:"uuid"`
	// CiName is the name of the configuration item.
	CiName string `json:"ciName"`
	// CiId is the configuration item ID.
	CiId int64 `json:"ciId"`
	// RenewTime is the renewal time of the entity.
	RenewTime string `json:"renewTime"`
	// MaxRelEntityCount is the maximum count of relationship entities.
	MaxRelEntityCount int `json:"maxRelEntityCount"`
	// RelEntityData contains relationship entity data.
	RelEntityData map[string]interface{} `json:"relEntityData"`
	// Name is the name of the entity.
	Name string `json:"name"`
	// AttrEntityData contains attribute entity data.
	AttrEntityData map[string]interface{} `json:"attrEntityData"`
	// ID is the identifier of the entity.
	ID int64 `json:"id"`
	// MaxAttrEntityCount is the maximum count of attribute entities.
	MaxAttrEntityCount int `json:"maxAttrEntityCount"`
	// IsVirtual indicates if the entity is virtual.
	IsVirtual int `json:"isVirtual"`
	// CiLabel is the label of the configuration item.
	CiLabel string `json:"ciLabel"`
	// MonitorStatus is the monitoring status of the entity.
	MonitorStatus string `json:"monitorStatus"`
	// AuthData contains authorization data for the entity.
	AuthData map[string]bool `json:"authData"`
}

// AResponse represents the response structure for attribute search operations.
type AResponse struct {
	// Status indicates the operation status (OK or ERROR).
	Status string `json:"Status"`
	// AReturn contains the list of attribute search results.
	AReturn []AReturn `json:"Return"`
	// TimeCost is the time cost of the operation in milliseconds.
	TimeCost int64 `json:"TimeCost"`
}

// AReturn represents a single attribute search result.
type AReturn struct {
	// Name is the name of the attribute.
	Name string `json:"name"`
	// ID is the identifier of the attribute.
	ID int64 `json:"id"`
}

// NewNeatClient creates a new NeatClient instance with default configuration.
// It performs authentication and initializes the client with the JWT token.
//
// Returns:
//   - *NeatClient: A new client instance ready to make API calls
//   - Panics if authentication fails
func NewNeatClient() *NeatClient {
	token, err := auth.Login()
	if err != nil {
		panic(err)
	}
	return &NeatClient{
		Client:       &http.Client{},
		NeatlogicUri: common.NeatlogicUri,
		JwtToken:     token,
	}
}

// NewNeatClientWithConfigPath creates a new NeatClient instance with a custom configuration file path.
// It performs authentication using the specified configuration and initializes the client with the JWT token.
//
// Parameters:
//   - configPath: Path to the configuration file to use
//
// Returns:
//   - *NeatClient: A new client instance ready to make API calls
//   - Panics if authentication fails
func NewNeatClientWithConfigPath(configPath string) *NeatClient {
	token, err := auth.LoginWithConfigPath(configPath)
	if err != nil {
		panic(err)
	}
	return &NeatClient{
		Client:       &http.Client{},
		NeatlogicUri: common.NeatlogicUri,
		JwtToken:     token,
	}
}

// GetAllCientity retrieves all CMDB entities for a given configuration item ID.
// It automatically handles pagination to retrieve all entities.
//
// Parameters:
//   - ciId: The configuration item ID to search for
//
// Returns:
//   - []TbodyList: A slice of all CMDB entities found
//   - error: An error if the operation fails
func (c *NeatClient) GetAllCientity(ciId int64) ([]TbodyList, error) {
	var allCientity []TbodyList
	currentPage := 1

	for {
		url := fmt.Sprintf("%s/api/rest/cmdb/cientity/search", c.NeatlogicUri)
		reqbody := CRequest{
			CiId:        ciId,
			CurrentPage: currentPage,
			PageSize:    100,
		}
		jsonData, err := json.Marshal(reqbody)
		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			return nil, err
		}

		resp, err := c.SendRequest(req)
		if err != nil {
			return nil, err
		}
		var respBody CResponse
		if err := json.Unmarshal(resp, &respBody); err != nil {
			return nil, err
		}
		allCientity = append(allCientity, respBody.CReturn.TbodyList...)
		if currentPage >= respBody.CReturn.PageCount {
			break
		}
		currentPage++
	}
	return allCientity, nil
}

// SearchCientityByFilter retrieves CMDB entities based on filter criteria.
// It automatically handles pagination to retrieve all matching entities.
//
// Parameters:
//   - reqbody: The request body containing filter criteria
//
// Returns:
//   - []TbodyList: A slice of CMDB entities that match the filter criteria
//   - error: An error if the operation fails
func (c *NeatClient) SearchCientityByFilter(reqbody CRequestBody) ([]TbodyList, error) {
	var allCientity []TbodyList
	currentPage := 1

	for {
		url := fmt.Sprintf("%s/api/rest/cmdb/cientity/search", c.NeatlogicUri)
		reqbody.CurrentPage = currentPage
		jsonData, err := json.Marshal(reqbody)
		if err != nil {
			return nil, err
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			return nil, err
		}
		resp, err := c.SendRequest(req)
		if err != nil {
			return nil, err
		}
		var respBody CResponse
		if err := json.Unmarshal(resp, &respBody); err != nil {
			return nil, err
		}
		allCientity = append(allCientity, respBody.CReturn.TbodyList...)
		if currentPage >= respBody.CReturn.PageCount {
			break
		}
		currentPage++
	}

	return allCientity, nil
}

// SearchCientityByKeyword searches for CMDB entities using a keyword and configuration item ID.
// It automatically handles pagination to retrieve all matching entities.
//
// Parameters:
//   - ciId: The configuration item ID to search within
//   - keyword: The keyword to search for
//
// Returns:
//   - []TbodyList: A slice of CMDB entities that match the search criteria
//   - error: An error if the operation fails
func (c *NeatClient) SearchCientityByKeyword(ciId int64, keyword string) ([]TbodyList, error) {
	var allCientity []TbodyList
	currentPage := 1

	for {
		url := fmt.Sprintf("%s/api/rest/cmdb/cientity/search", c.NeatlogicUri)

		// Build request body
		reqbody := CRequest{
			CiId:        ciId,
			CurrentPage: currentPage,
			Keyword:     keyword,
		}
		jsonData, err := json.Marshal(reqbody)
		if err != nil {
			return nil, err
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			return nil, err
		}
		resp, err := c.SendRequest(req)
		if err != nil {
			return nil, err
		}
		var respBody CResponse
		if err := json.Unmarshal(resp, &respBody); err != nil {
			return nil, err
		}
		allCientity = append(allCientity, respBody.CReturn.TbodyList...)
		if currentPage >= respBody.CReturn.PageCount {
			break
		}
		currentPage++
	}

	return allCientity, nil
}

// GetCientity retrieves a specific CMDB entity by its configuration item ID and entity ID.
// It returns the complete entity information without limiting relationship or attribute entities.
//
// Parameters:
//   - ciId: The configuration item ID
//   - ciEntityId: The specific entity ID to retrieve
//
// Returns:
//   - TbodyList: The requested CMDB entity
//   - error: An error if the operation fails
func (c *NeatClient) GetCientity(ciId int64, ciEntityId int64) (TbodyList, error) {
	// Do not limit RelEntity and AttrEntity
	url := fmt.Sprintf("%s/api/rest/cmdb/cientity/get", c.NeatlogicUri)

	// Build request body
	reqbody := CRequest{
		CiId:            ciId,
		CiEntityId:      ciEntityId,
		LimitRelEntity:  false,
		LimitAttrEntity: false,
	}
	jsonData, err := json.Marshal(reqbody)
	if err != nil {
		return TbodyList{}, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return TbodyList{}, err
	}
	resp, err := c.SendRequest(req)
	if err != nil {
		return TbodyList{}, err
	}
	var respBody GetCientityResponse
	if err := json.Unmarshal(resp, &respBody); err != nil {
		return TbodyList{}, err
	}
	return respBody.GetcientityReturn, nil
}

// SendRequest sends an HTTP request with JWT authentication headers.
// It adds the required headers and processes the response.
//
// Parameters:
//   - req: The HTTP request to send
//
// Returns:
//   - []byte: The response body as bytes
//   - error: An error if the operation fails
func (c *NeatClient) SendRequest(req *http.Request) ([]byte, error) {
	// Set JWT authentication headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.JwtToken)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse response
	respBody, err := ParseResourceResponse(resp)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

// ParseResourceResponse parses an HTTP response and returns the response body.
// It checks the status code and returns an error if the request was not successful.
//
// Parameters:
//   - resp: The HTTP response to parse
//
// Returns:
//   - []byte: The response body as bytes
//   - error: An error if the operation fails or status code is not OK
func ParseResourceResponse(resp *http.Response) ([]byte, error) {
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

// SearchTargetAttr searches for target attributes based on a request body and attribute ID.
// It allows searching for specific attributes with the given criteria.
//
// Parameters:
//   - reqbody: The request body containing search criteria (keyword is typically required)
//   - attrId: The attribute ID to search for
//
// Returns:
//   - []AReturn: A slice of attribute search results
//   - error: An error if the operation fails
func (c *NeatClient) SearchTargetAttr(reqbody CRequestBody, attrId string) ([]AReturn, error) {
	// This function is used to search for target attributes
	// It requires a request body (containing keyword) and an attribute ID
	apiurl := fmt.Sprintf("%s/api/rest/cmdb/attr/targetci/search", c.NeatlogicUri)
	parmas := url.Values{}
	parmas.Add("attrId", attrId)
	jsonData, err := json.Marshal(reqbody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", apiurl+"?"+parmas.Encode(), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	resp, err := c.SendRequest(req)
	if err != nil {
		return nil, err
	}
	var respBody AResponse
	if err := json.Unmarshal(resp, &respBody); err != nil {
		return nil, err
	}

	return respBody.AReturn, nil
}
