package main

type SharedCollection struct {
	ShareId  string `json:"shareId"`
	ShareURL string `json:"shareURL"`
	NodeInfo struct {
		Id           string `json:"id"`
		Kind         string `json:"kind"` // "SHARED_COLLECTION"
		Version      int    `json:"version"`
		CreatedDate  string `json:"createdDate"`
		ModifiedDate string `json:"modifiedDate"`
		Name         string `json:"name"`
	} `json:"nodeInfo"`
	CreationDate string `json:"creationDate"`
	StatusCode   int    `json:"statusCode"`
}

type Resource struct {
	Data  []ResourceData `json:"data"`
	Count int            `json:"count"`
}

type ResourceData struct {
	//"protectedFolder":false,
	TempLink string `json:"tempLink,omitempty"`
	//"keywords":[]
	//"transforms":[],
	//"ownerId":"OWNERID"
	//"xAccntParentMap":{}
	//"eTagResponse":"ETAG"
	Id   string `json:"id"`
	Kind string `json:"kind"`
	//"xAccntParents":[]
	//"version":3
	//"labels":[]
	Properties struct {
		//"extension":"iso"
		Size        int64  `json:"size"`        //"size":1236461568
		ContentType string `json:"contentType"` //"contentType":"application/x-iso9660-image"
		//"version":1
		MD5 string `json:"md5"`
	} `json:"contentProperties"`
	//"createdDate":"2016-05-15T20:07:10.418Z"
	//"parentMap":{"SHARED_COLLECTION":["ID"],"FOLDER":["ID"]}
	//"createdBy":"acd_cli_oa-SOMETHINGSOMETHING"
	Restricted bool `json:"restricted"`
	//"modifiedDate":"2016-05-15T20:29:17.466Z"
	Name     string `json:"name"`
	IsShared bool   `json:"isShared"`
	//Parents []string `json:"parents"`
	Status string `json:"status"`
}
