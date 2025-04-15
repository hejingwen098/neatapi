package neatlogic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hejingwen098/neatapi/auth"
	"github.com/hejingwen098/neatapi/common"
)

// NeatLogic Login 结构体
type LRequest struct {
	UserID   string `json:"userid"`
	Password string `json:"password"`
}

type LResponse struct {
	Status   string `json:"Status"`
	Message  string `json:"Message"`
	JwtToken string `json:"JwtToken"`
}
type NeatClient struct {
	Client       *http.Client
	NeatlogicUri string
	JwtToken     string
}

// CMDB Cientity 结构体
type CRequestBody struct {
	PageSize       int    `json:"pageSize"`
	CiId           int    `json:"ciId"`
	NeedAction     bool   `json:"needAction"`
	NeedExpand     bool   `json:"needExpand"`
	NeedActionType bool   `json:"needActionType"`
	NeedCheck      bool   `json:"needCheck"`
	Mode           string `json:"mode"`
	CurrentPage    int    `json:"currentPage"`
	// What J8 struct is this
	AttrFilterList       map[string]interface{} `json:"attrFilterList"`
	GlobalAttrFilterList map[string]interface{} `json:"globalAttrFilterList"`
	RelFilterList        map[string]interface{} `json:"relFilterList"`
	Keyword              string                 `json:"keyword"`
}

type CRequest struct {
	CiId        int    `json:"ciId"`
	CiEntityId  int    `json:"ciEntityId"`
	Keyword     string `json:"keyword"`
	PageSize    int    `json:"pageSize"`
	CurrentPage int    `json:"currentPage"`
}

type CResponse struct {
	Status   string  `json:"Status"`
	CReturn  CReturn `json:"Return"`
	TimeCost int64   `json:"TimeCost"`
}

type CReturn struct {
	PageCount   int         `json:"pageCount"`
	RowNum      int         `json:"rowNum"`
	PageSize    int         `json:"pageSize"`
	CurrentPage int         `json:"currentPage"`
	TbodyList   []TbodyList `json:"tbodyList"`
}

type TbodyList struct {
	CiIcon               string                 `json:"ciIcon"`
	GlobalAttrEntityData map[string]interface{} `json:"globalAttrEntityData"`
	TypeName             string                 `json:"typeName"`
	Type                 int64                  `json:"type"`
	InspectStatus        string                 `json:"inspectStatus"`
	UUID                 string                 `json:"uuid"`
	CiName               string                 `json:"ciName"`
	CiId                 int64                  `json:"ciId"`
	RenewTime            string                 `json:"renewTime"`
	MaxRelEntityCount    int                    `json:"maxRelEntityCount"`
	RelEntityData        map[string]interface{} `json:"relEntityData"`
	Name                 string                 `json:"name"`
	AttrEntityData       map[string]interface{} `json:"attrEntityData"`
	ID                   int64                  `json:"id"`
	MaxAttrEntityCount   int                    `json:"maxAttrEntityCount"`
	IsVirtual            int                    `json:"isVirtual"`
	CiLabel              string                 `json:"ciLabel"`
	MonitorStatus        string                 `json:"monitorStatus"`
	AuthData             map[string]bool        `json:"authData"`
}

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

func (c *NeatClient) GetAllCientity(ciId int) ([]TbodyList, error) {
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

// 垃圾请求body，想用自己拼吧，爷不伺候
func (c *NeatClient) SearchCientityByFilter(repbody CRequestBody) ([]TbodyList, error) {
	var allCientity []TbodyList
	currentPage := 1

	for {
		url := fmt.Sprintf("%s/api/rest/cmdb/cientity/search", c.NeatlogicUri)

		jsonData, err := json.Marshal(repbody)
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

func (c *NeatClient) SearchCientityByKeyword(ciId int, keyword string) ([]TbodyList, error) {
	var allCientity []TbodyList
	currentPage := 1

	for {
		url := fmt.Sprintf("%s/api/rest/cmdb/cientity/search", c.NeatlogicUri)

		// 构建请求body
		reqbody := CRequest{
			CiId:    ciId,
			Keyword: keyword,
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

func (c *NeatClient) GetCientity(ciId int, ciEntityId int) ([]TbodyList, error) {
	url := fmt.Sprintf("%s/api/rest/cmdb/cientity/get", c.NeatlogicUri)

	// 构建请求body
	reqbody := CRequest{
		CiId:       ciId,
		CiEntityId: ciEntityId,
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
	return respBody.CReturn.TbodyList, nil
}
func (c *NeatClient) SendRequest(req *http.Request) ([]byte, error) {
	// 设置JWT认证头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.JwtToken)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解析响应
	respBody, err := ParseResourceResponse(resp)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func ParseResourceResponse(resp *http.Response) ([]byte, error) {
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
