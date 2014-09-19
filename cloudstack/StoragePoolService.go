//
// Copyright 2014, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cloudstack

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type ListStorageProvidersParams struct {
	p map[string]interface{}
}

func (p *ListStorageProvidersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *ListStorageProvidersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListStorageProvidersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListStorageProvidersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListStorageProvidersParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storagePoolType"] = v
	return
}

// You should always use this function to get a new ListStorageProvidersParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewListStorageProvidersParams(storagePoolType string) *ListStorageProvidersParams {
	p := &ListStorageProvidersParams{}
	p.p = make(map[string]interface{})
	p.p["storagePoolType"] = storagePoolType
	return p
}

// Lists storage providers.
func (s *StoragePoolService) ListStorageProviders(p *ListStorageProvidersParams) (*ListStorageProvidersResponse, error) {
	resp, err := s.cs.newRequest("listStorageProviders", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListStorageProvidersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListStorageProvidersResponse struct {
	Count            int                `json:"count"`
	StorageProviders []*StorageProvider `json:"storageprovider"`
}

type StorageProvider struct {
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
}

type EnableStorageMaintenanceParams struct {
	p map[string]interface{}
}

func (p *EnableStorageMaintenanceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *EnableStorageMaintenanceParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new EnableStorageMaintenanceParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewEnableStorageMaintenanceParams(id string) *EnableStorageMaintenanceParams {
	p := &EnableStorageMaintenanceParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Puts storage pool into maintenance state
func (s *StoragePoolService) EnableStorageMaintenance(p *EnableStorageMaintenanceParams) (*EnableStorageMaintenanceResponse, error) {
	resp, err := s.cs.newRequest("enableStorageMaintenance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableStorageMaintenanceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type EnableStorageMaintenanceResponse struct {
	JobID                string            `json:"jobid,omitempty"`
	Disksizeused         int               `json:"disksizeused,omitempty"`
	Zonename             string            `json:"zonename,omitempty"`
	Id                   string            `json:"id,omitempty"`
	Zoneid               string            `json:"zoneid,omitempty"`
	Podid                string            `json:"podid,omitempty"`
	Created              string            `json:"created,omitempty"`
	Scope                string            `json:"scope,omitempty"`
	Name                 string            `json:"name,omitempty"`
	Ipaddress            string            `json:"ipaddress,omitempty"`
	Disksizetotal        int               `json:"disksizetotal,omitempty"`
	Suitableformigration bool              `json:"suitableformigration,omitempty"`
	Clusterid            string            `json:"clusterid,omitempty"`
	State                string            `json:"state,omitempty"`
	Storagecapabilities  map[string]string `json:"storagecapabilities,omitempty"`
	Capacityiops         int               `json:"capacityiops,omitempty"`
	Hypervisor           string            `json:"hypervisor,omitempty"`
	Disksizeallocated    int               `json:"disksizeallocated,omitempty"`
	Tags                 string            `json:"tags,omitempty"`
	Path                 string            `json:"path,omitempty"`
	Clustername          string            `json:"clustername,omitempty"`
	Podname              string            `json:"podname,omitempty"`
	Type                 string            `json:"type,omitempty"`
}

type CancelStorageMaintenanceParams struct {
	p map[string]interface{}
}

func (p *CancelStorageMaintenanceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *CancelStorageMaintenanceParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new CancelStorageMaintenanceParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewCancelStorageMaintenanceParams(id string) *CancelStorageMaintenanceParams {
	p := &CancelStorageMaintenanceParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Cancels maintenance for primary storage
func (s *StoragePoolService) CancelStorageMaintenance(p *CancelStorageMaintenanceParams) (*CancelStorageMaintenanceResponse, error) {
	resp, err := s.cs.newRequest("cancelStorageMaintenance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CancelStorageMaintenanceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type CancelStorageMaintenanceResponse struct {
	JobID                string            `json:"jobid,omitempty"`
	Zonename             string            `json:"zonename,omitempty"`
	Disksizetotal        int               `json:"disksizetotal,omitempty"`
	Storagecapabilities  map[string]string `json:"storagecapabilities,omitempty"`
	Tags                 string            `json:"tags,omitempty"`
	Podid                string            `json:"podid,omitempty"`
	Zoneid               string            `json:"zoneid,omitempty"`
	Ipaddress            string            `json:"ipaddress,omitempty"`
	Clustername          string            `json:"clustername,omitempty"`
	Name                 string            `json:"name,omitempty"`
	Clusterid            string            `json:"clusterid,omitempty"`
	Type                 string            `json:"type,omitempty"`
	Podname              string            `json:"podname,omitempty"`
	Disksizeused         int               `json:"disksizeused,omitempty"`
	Path                 string            `json:"path,omitempty"`
	Disksizeallocated    int               `json:"disksizeallocated,omitempty"`
	Capacityiops         int               `json:"capacityiops,omitempty"`
	Id                   string            `json:"id,omitempty"`
	State                string            `json:"state,omitempty"`
	Created              string            `json:"created,omitempty"`
	Hypervisor           string            `json:"hypervisor,omitempty"`
	Scope                string            `json:"scope,omitempty"`
	Suitableformigration bool              `json:"suitableformigration,omitempty"`
}
