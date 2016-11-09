package k8s

import "net/http"

func Swagger(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte(api))
}

const api = `{
  "swaggerVersion": "1.2",
  "apiVersion": "v1",
  "basePath": "https://172.17.0.6:443",
  "resourcePath": "/api/v1",
  "apis": [
   {
    "path": "/api/v1/namespaces/{namespace}/bindings",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.Binding",
      "method": "POST",
      "summary": "create a Binding",
      "nickname": "createNamespacedBinding",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Binding",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Binding"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/componentstatuses",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.ComponentStatusList",
      "method": "GET",
      "summary": "list objects of kind ComponentStatus",
      "nickname": "listNamespacedComponentStatus",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ComponentStatusList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/componentstatuses/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.ComponentStatus",
      "method": "GET",
      "summary": "read the specified ComponentStatus",
      "nickname": "readNamespacedComponentStatus",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ComponentStatus",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ComponentStatus"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/configmaps",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.ConfigMapList",
      "method": "GET",
      "summary": "list or watch objects of kind ConfigMap",
      "nickname": "listNamespacedConfigMap",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ConfigMapList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.ConfigMap",
      "method": "POST",
      "summary": "create a ConfigMap",
      "nickname": "createNamespacedConfigMap",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.ConfigMap",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ConfigMap"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete collection of ConfigMap",
      "nickname": "deletecollectionNamespacedConfigMap",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/configmaps",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of ConfigMap",
      "nickname": "watchNamespacedConfigMapList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/configmaps/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.ConfigMap",
      "method": "GET",
      "summary": "read the specified ConfigMap",
      "nickname": "readNamespacedConfigMap",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "export",
        "description": "Should this value be exported.  Export strips fields that a user can not specify.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "exact",
        "description": "Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ConfigMap",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ConfigMap"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.ConfigMap",
      "method": "PUT",
      "summary": "replace the specified ConfigMap",
      "nickname": "replaceNamespacedConfigMap",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.ConfigMap",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ConfigMap",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ConfigMap"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.ConfigMap",
      "method": "PATCH",
      "summary": "partially update the specified ConfigMap",
      "nickname": "patchNamespacedConfigMap",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "unversioned.Patch",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ConfigMap",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ConfigMap"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "application/json-patch+json",
       "application/merge-patch+json",
       "application/strategic-merge-patch+json"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete a ConfigMap",
      "nickname": "deleteNamespacedConfigMap",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.DeleteOptions",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ConfigMap",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/configmaps/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch changes to an object of kind ConfigMap",
      "nickname": "watchNamespacedConfigMap",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ConfigMap",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/configmaps",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.ConfigMapList",
      "method": "GET",
      "summary": "list or watch objects of kind ConfigMap",
      "nickname": "listConfigMap",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ConfigMapList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/configmaps",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of ConfigMap",
      "nickname": "watchConfigMapList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/endpoints",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.EndpointsList",
      "method": "GET",
      "summary": "list or watch objects of kind Endpoints",
      "nickname": "listNamespacedEndpoints",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.EndpointsList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Endpoints",
      "method": "POST",
      "summary": "create a Endpoints",
      "nickname": "createNamespacedEndpoints",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Endpoints",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Endpoints"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete collection of Endpoints",
      "nickname": "deletecollectionNamespacedEndpoints",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/endpoints",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of Endpoints",
      "nickname": "watchNamespacedEndpointsList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/endpoints/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.Endpoints",
      "method": "GET",
      "summary": "read the specified Endpoints",
      "nickname": "readNamespacedEndpoints",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "export",
        "description": "Should this value be exported.  Export strips fields that a user can not specify.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "exact",
        "description": "Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Endpoints",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Endpoints"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Endpoints",
      "method": "PUT",
      "summary": "replace the specified Endpoints",
      "nickname": "replaceNamespacedEndpoints",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Endpoints",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Endpoints",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Endpoints"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Endpoints",
      "method": "PATCH",
      "summary": "partially update the specified Endpoints",
      "nickname": "patchNamespacedEndpoints",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "unversioned.Patch",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Endpoints",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Endpoints"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "application/json-patch+json",
       "application/merge-patch+json",
       "application/strategic-merge-patch+json"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete a Endpoints",
      "nickname": "deleteNamespacedEndpoints",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.DeleteOptions",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Endpoints",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/endpoints/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch changes to an object of kind Endpoints",
      "nickname": "watchNamespacedEndpoints",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Endpoints",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/endpoints",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.EndpointsList",
      "method": "GET",
      "summary": "list or watch objects of kind Endpoints",
      "nickname": "listEndpoints",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.EndpointsList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/endpoints",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of Endpoints",
      "nickname": "watchEndpointsList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/events",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.EventList",
      "method": "GET",
      "summary": "list or watch objects of kind Event",
      "nickname": "listNamespacedEvent",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.EventList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Event",
      "method": "POST",
      "summary": "create a Event",
      "nickname": "createNamespacedEvent",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Event",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Event"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete collection of Event",
      "nickname": "deletecollectionNamespacedEvent",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/events",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of Event",
      "nickname": "watchNamespacedEventList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/events/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.Event",
      "method": "GET",
      "summary": "read the specified Event",
      "nickname": "readNamespacedEvent",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "export",
        "description": "Should this value be exported.  Export strips fields that a user can not specify.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "exact",
        "description": "Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Event",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Event"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Event",
      "method": "PUT",
      "summary": "replace the specified Event",
      "nickname": "replaceNamespacedEvent",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Event",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Event",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Event"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Event",
      "method": "PATCH",
      "summary": "partially update the specified Event",
      "nickname": "patchNamespacedEvent",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "unversioned.Patch",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Event",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Event"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "application/json-patch+json",
       "application/merge-patch+json",
       "application/strategic-merge-patch+json"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete a Event",
      "nickname": "deleteNamespacedEvent",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.DeleteOptions",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Event",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/events/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch changes to an object of kind Event",
      "nickname": "watchNamespacedEvent",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Event",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/events",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.EventList",
      "method": "GET",
      "summary": "list or watch objects of kind Event",
      "nickname": "listEvent",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.EventList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/events",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of Event",
      "nickname": "watchEventList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/limitranges",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.LimitRangeList",
      "method": "GET",
      "summary": "list or watch objects of kind LimitRange",
      "nickname": "listNamespacedLimitRange",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.LimitRangeList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.LimitRange",
      "method": "POST",
      "summary": "create a LimitRange",
      "nickname": "createNamespacedLimitRange",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.LimitRange",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.LimitRange"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete collection of LimitRange",
      "nickname": "deletecollectionNamespacedLimitRange",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/limitranges",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of LimitRange",
      "nickname": "watchNamespacedLimitRangeList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/limitranges/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.LimitRange",
      "method": "GET",
      "summary": "read the specified LimitRange",
      "nickname": "readNamespacedLimitRange",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "export",
        "description": "Should this value be exported.  Export strips fields that a user can not specify.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "exact",
        "description": "Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the LimitRange",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.LimitRange"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.LimitRange",
      "method": "PUT",
      "summary": "replace the specified LimitRange",
      "nickname": "replaceNamespacedLimitRange",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.LimitRange",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the LimitRange",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.LimitRange"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.LimitRange",
      "method": "PATCH",
      "summary": "partially update the specified LimitRange",
      "nickname": "patchNamespacedLimitRange",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "unversioned.Patch",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the LimitRange",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.LimitRange"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "application/json-patch+json",
       "application/merge-patch+json",
       "application/strategic-merge-patch+json"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete a LimitRange",
      "nickname": "deleteNamespacedLimitRange",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.DeleteOptions",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the LimitRange",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/limitranges/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch changes to an object of kind LimitRange",
      "nickname": "watchNamespacedLimitRange",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the LimitRange",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/limitranges",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.LimitRangeList",
      "method": "GET",
      "summary": "list or watch objects of kind LimitRange",
      "nickname": "listLimitRange",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.LimitRangeList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/limitranges",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of LimitRange",
      "nickname": "watchLimitRangeList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.NamespaceList",
      "method": "GET",
      "summary": "list or watch objects of kind Namespace",
      "nickname": "listNamespacedNamespace",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.NamespaceList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Namespace",
      "method": "POST",
      "summary": "create a Namespace",
      "nickname": "createNamespacedNamespace",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Namespace",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Namespace"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete collection of Namespace",
      "nickname": "deletecollectionNamespacedNamespace",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of Namespace",
      "nickname": "watchNamespacedNamespaceList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.Namespace",
      "method": "GET",
      "summary": "read the specified Namespace",
      "nickname": "readNamespacedNamespace",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "export",
        "description": "Should this value be exported.  Export strips fields that a user can not specify.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "exact",
        "description": "Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Namespace",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Namespace"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Namespace",
      "method": "PUT",
      "summary": "replace the specified Namespace",
      "nickname": "replaceNamespacedNamespace",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Namespace",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Namespace",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Namespace"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Namespace",
      "method": "PATCH",
      "summary": "partially update the specified Namespace",
      "nickname": "patchNamespacedNamespace",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "unversioned.Patch",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Namespace",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Namespace"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "application/json-patch+json",
       "application/merge-patch+json",
       "application/strategic-merge-patch+json"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete a Namespace",
      "nickname": "deleteNamespacedNamespace",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.DeleteOptions",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Namespace",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch changes to an object of kind Namespace",
      "nickname": "watchNamespacedNamespace",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Namespace",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{name}/finalize",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.Namespace",
      "method": "PUT",
      "summary": "replace finalize of the specified Namespace",
      "nickname": "replaceNamespacedNamespaceFinalize",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Namespace",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Namespace",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Namespace"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{name}/status",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.Namespace",
      "method": "PUT",
      "summary": "replace status of the specified Namespace",
      "nickname": "replaceNamespacedNamespaceStatus",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Namespace",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Namespace",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Namespace"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/nodes",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.NodeList",
      "method": "GET",
      "summary": "list or watch objects of kind Node",
      "nickname": "listNamespacedNode",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.NodeList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Node",
      "method": "POST",
      "summary": "create a Node",
      "nickname": "createNamespacedNode",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Node",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Node"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete collection of Node",
      "nickname": "deletecollectionNamespacedNode",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/nodes",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of Node",
      "nickname": "watchNamespacedNodeList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/nodes/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.Node",
      "method": "GET",
      "summary": "read the specified Node",
      "nickname": "readNamespacedNode",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "export",
        "description": "Should this value be exported.  Export strips fields that a user can not specify.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "exact",
        "description": "Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Node"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Node",
      "method": "PUT",
      "summary": "replace the specified Node",
      "nickname": "replaceNamespacedNode",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Node",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Node"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Node",
      "method": "PATCH",
      "summary": "partially update the specified Node",
      "nickname": "patchNamespacedNode",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "unversioned.Patch",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Node"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "application/json-patch+json",
       "application/merge-patch+json",
       "application/strategic-merge-patch+json"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete a Node",
      "nickname": "deleteNamespacedNode",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.DeleteOptions",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/nodes/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch changes to an object of kind Node",
      "nickname": "watchNamespacedNode",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/proxy/nodes/{name}/{path}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "string",
      "method": "GET",
      "summary": "proxy GET requests to Node",
      "nickname": "proxyGETNamespacedNode",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "PUT",
      "summary": "proxy PUT requests to Node",
      "nickname": "proxyPUTNamespacedNode",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "POST",
      "summary": "proxy POST requests to Node",
      "nickname": "proxyPOSTNamespacedNode",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "DELETE",
      "summary": "proxy DELETE requests to Node",
      "nickname": "proxyDELETENamespacedNode",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "HEAD",
      "summary": "proxy HEAD requests to Node",
      "nickname": "proxyHEADNamespacedNode",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "OPTIONS",
      "summary": "proxy OPTIONS requests to Node",
      "nickname": "proxyOPTIONSNamespacedNode",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/proxy/nodes/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "string",
      "method": "GET",
      "summary": "proxy GET requests to Node",
      "nickname": "proxyGETNamespacedNode",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "PUT",
      "summary": "proxy PUT requests to Node",
      "nickname": "proxyPUTNamespacedNode",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "POST",
      "summary": "proxy POST requests to Node",
      "nickname": "proxyPOSTNamespacedNode",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "DELETE",
      "summary": "proxy DELETE requests to Node",
      "nickname": "proxyDELETENamespacedNode",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "HEAD",
      "summary": "proxy HEAD requests to Node",
      "nickname": "proxyHEADNamespacedNode",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "OPTIONS",
      "summary": "proxy OPTIONS requests to Node",
      "nickname": "proxyOPTIONSNamespacedNode",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/nodes/{name}/proxy",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "string",
      "method": "GET",
      "summary": "connect GET requests to proxy of Node",
      "nickname": "connectGetNamespacedNodeProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to node.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "POST",
      "summary": "connect POST requests to proxy of Node",
      "nickname": "connectPostNamespacedNodeProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to node.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "PUT",
      "summary": "connect PUT requests to proxy of Node",
      "nickname": "connectPutNamespacedNodeProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to node.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "DELETE",
      "summary": "connect DELETE requests to proxy of Node",
      "nickname": "connectDeleteNamespacedNodeProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to node.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "HEAD",
      "summary": "connect HEAD requests to proxy of Node",
      "nickname": "connectHeadNamespacedNodeProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to node.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "OPTIONS",
      "summary": "connect OPTIONS requests to proxy of Node",
      "nickname": "connectOptionsNamespacedNodeProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to node.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/nodes/{name}/proxy/{path}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "string",
      "method": "GET",
      "summary": "connect GET requests to proxy of Node",
      "nickname": "connectGetNamespacedNodeProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to node.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "POST",
      "summary": "connect POST requests to proxy of Node",
      "nickname": "connectPostNamespacedNodeProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to node.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "PUT",
      "summary": "connect PUT requests to proxy of Node",
      "nickname": "connectPutNamespacedNodeProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to node.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "DELETE",
      "summary": "connect DELETE requests to proxy of Node",
      "nickname": "connectDeleteNamespacedNodeProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to node.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "HEAD",
      "summary": "connect HEAD requests to proxy of Node",
      "nickname": "connectHeadNamespacedNodeProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to node.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "OPTIONS",
      "summary": "connect OPTIONS requests to proxy of Node",
      "nickname": "connectOptionsNamespacedNodeProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to node.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/nodes/{name}/status",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.Node",
      "method": "PUT",
      "summary": "replace status of the specified Node",
      "nickname": "replaceNamespacedNodeStatus",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Node",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Node",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Node"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/persistentvolumeclaims",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.PersistentVolumeClaimList",
      "method": "GET",
      "summary": "list or watch objects of kind PersistentVolumeClaim",
      "nickname": "listNamespacedPersistentVolumeClaim",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PersistentVolumeClaimList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.PersistentVolumeClaim",
      "method": "POST",
      "summary": "create a PersistentVolumeClaim",
      "nickname": "createNamespacedPersistentVolumeClaim",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.PersistentVolumeClaim",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PersistentVolumeClaim"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete collection of PersistentVolumeClaim",
      "nickname": "deletecollectionNamespacedPersistentVolumeClaim",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/persistentvolumeclaims",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of PersistentVolumeClaim",
      "nickname": "watchNamespacedPersistentVolumeClaimList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/persistentvolumeclaims/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.PersistentVolumeClaim",
      "method": "GET",
      "summary": "read the specified PersistentVolumeClaim",
      "nickname": "readNamespacedPersistentVolumeClaim",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "export",
        "description": "Should this value be exported.  Export strips fields that a user can not specify.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "exact",
        "description": "Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the PersistentVolumeClaim",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PersistentVolumeClaim"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.PersistentVolumeClaim",
      "method": "PUT",
      "summary": "replace the specified PersistentVolumeClaim",
      "nickname": "replaceNamespacedPersistentVolumeClaim",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.PersistentVolumeClaim",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the PersistentVolumeClaim",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PersistentVolumeClaim"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.PersistentVolumeClaim",
      "method": "PATCH",
      "summary": "partially update the specified PersistentVolumeClaim",
      "nickname": "patchNamespacedPersistentVolumeClaim",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "unversioned.Patch",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the PersistentVolumeClaim",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PersistentVolumeClaim"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "application/json-patch+json",
       "application/merge-patch+json",
       "application/strategic-merge-patch+json"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete a PersistentVolumeClaim",
      "nickname": "deleteNamespacedPersistentVolumeClaim",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.DeleteOptions",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the PersistentVolumeClaim",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/persistentvolumeclaims/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch changes to an object of kind PersistentVolumeClaim",
      "nickname": "watchNamespacedPersistentVolumeClaim",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the PersistentVolumeClaim",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/persistentvolumeclaims",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.PersistentVolumeClaimList",
      "method": "GET",
      "summary": "list or watch objects of kind PersistentVolumeClaim",
      "nickname": "listPersistentVolumeClaim",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PersistentVolumeClaimList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/persistentvolumeclaims",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of PersistentVolumeClaim",
      "nickname": "watchPersistentVolumeClaimList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/persistentvolumeclaims/{name}/status",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.PersistentVolumeClaim",
      "method": "PUT",
      "summary": "replace status of the specified PersistentVolumeClaim",
      "nickname": "replaceNamespacedPersistentVolumeClaimStatus",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.PersistentVolumeClaim",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the PersistentVolumeClaim",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PersistentVolumeClaim"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/persistentvolumes",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.PersistentVolumeList",
      "method": "GET",
      "summary": "list or watch objects of kind PersistentVolume",
      "nickname": "listNamespacedPersistentVolume",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PersistentVolumeList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.PersistentVolume",
      "method": "POST",
      "summary": "create a PersistentVolume",
      "nickname": "createNamespacedPersistentVolume",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.PersistentVolume",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PersistentVolume"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete collection of PersistentVolume",
      "nickname": "deletecollectionNamespacedPersistentVolume",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/persistentvolumes",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of PersistentVolume",
      "nickname": "watchNamespacedPersistentVolumeList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/persistentvolumes/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.PersistentVolume",
      "method": "GET",
      "summary": "read the specified PersistentVolume",
      "nickname": "readNamespacedPersistentVolume",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "export",
        "description": "Should this value be exported.  Export strips fields that a user can not specify.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "exact",
        "description": "Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the PersistentVolume",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PersistentVolume"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.PersistentVolume",
      "method": "PUT",
      "summary": "replace the specified PersistentVolume",
      "nickname": "replaceNamespacedPersistentVolume",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.PersistentVolume",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the PersistentVolume",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PersistentVolume"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.PersistentVolume",
      "method": "PATCH",
      "summary": "partially update the specified PersistentVolume",
      "nickname": "patchNamespacedPersistentVolume",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "unversioned.Patch",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the PersistentVolume",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PersistentVolume"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "application/json-patch+json",
       "application/merge-patch+json",
       "application/strategic-merge-patch+json"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete a PersistentVolume",
      "nickname": "deleteNamespacedPersistentVolume",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.DeleteOptions",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the PersistentVolume",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/persistentvolumes/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch changes to an object of kind PersistentVolume",
      "nickname": "watchNamespacedPersistentVolume",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the PersistentVolume",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/persistentvolumes/{name}/status",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.PersistentVolume",
      "method": "PUT",
      "summary": "replace status of the specified PersistentVolume",
      "nickname": "replaceNamespacedPersistentVolumeStatus",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.PersistentVolume",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the PersistentVolume",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PersistentVolume"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/pods",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.PodList",
      "method": "GET",
      "summary": "list or watch objects of kind Pod",
      "nickname": "listNamespacedPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PodList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Pod",
      "method": "POST",
      "summary": "create a Pod",
      "nickname": "createNamespacedPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Pod",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Pod"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete collection of Pod",
      "nickname": "deletecollectionNamespacedPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/pods",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of Pod",
      "nickname": "watchNamespacedPodList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/pods/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.Pod",
      "method": "GET",
      "summary": "read the specified Pod",
      "nickname": "readNamespacedPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "export",
        "description": "Should this value be exported.  Export strips fields that a user can not specify.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "exact",
        "description": "Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Pod"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Pod",
      "method": "PUT",
      "summary": "replace the specified Pod",
      "nickname": "replaceNamespacedPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Pod",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Pod"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Pod",
      "method": "PATCH",
      "summary": "partially update the specified Pod",
      "nickname": "patchNamespacedPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "unversioned.Patch",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Pod"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "application/json-patch+json",
       "application/merge-patch+json",
       "application/strategic-merge-patch+json"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete a Pod",
      "nickname": "deleteNamespacedPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.DeleteOptions",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/pods/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch changes to an object of kind Pod",
      "nickname": "watchNamespacedPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/proxy/namespaces/{namespace}/pods/{name}/{path}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "string",
      "method": "GET",
      "summary": "proxy GET requests to Pod",
      "nickname": "proxyGETNamespacedPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "PUT",
      "summary": "proxy PUT requests to Pod",
      "nickname": "proxyPUTNamespacedPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "POST",
      "summary": "proxy POST requests to Pod",
      "nickname": "proxyPOSTNamespacedPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "DELETE",
      "summary": "proxy DELETE requests to Pod",
      "nickname": "proxyDELETENamespacedPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "HEAD",
      "summary": "proxy HEAD requests to Pod",
      "nickname": "proxyHEADNamespacedPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "OPTIONS",
      "summary": "proxy OPTIONS requests to Pod",
      "nickname": "proxyOPTIONSNamespacedPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/proxy/namespaces/{namespace}/pods/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "string",
      "method": "GET",
      "summary": "proxy GET requests to Pod",
      "nickname": "proxyGETNamespacedPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "PUT",
      "summary": "proxy PUT requests to Pod",
      "nickname": "proxyPUTNamespacedPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "POST",
      "summary": "proxy POST requests to Pod",
      "nickname": "proxyPOSTNamespacedPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "DELETE",
      "summary": "proxy DELETE requests to Pod",
      "nickname": "proxyDELETENamespacedPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "HEAD",
      "summary": "proxy HEAD requests to Pod",
      "nickname": "proxyHEADNamespacedPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "OPTIONS",
      "summary": "proxy OPTIONS requests to Pod",
      "nickname": "proxyOPTIONSNamespacedPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/pods",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.PodList",
      "method": "GET",
      "summary": "list or watch objects of kind Pod",
      "nickname": "listPod",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PodList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/pods",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of Pod",
      "nickname": "watchPodList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/pods/{name}/attach",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "string",
      "method": "GET",
      "summary": "connect GET requests to attach of Pod",
      "nickname": "connectGetNamespacedPodAttach",
      "parameters": [
       {
        "type": "boolean",
        "paramType": "query",
        "name": "stdin",
        "description": "Stdin if true, redirects the standard input stream of the pod for this call. Defaults to false.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "stdout",
        "description": "Stdout if true indicates that stdout is to be redirected for the attach call. Defaults to true.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "stderr",
        "description": "Stderr if true indicates that stderr is to be redirected for the attach call. Defaults to true.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "tty",
        "description": "TTY if true indicates that a tty will be allocated for the attach call. This is passed through the container runtime so the tty is allocated on the worker node by the container runtime. Defaults to false.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "container",
        "description": "The container in which to execute the command. Defaults to only container if there is only one container in the pod.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "POST",
      "summary": "connect POST requests to attach of Pod",
      "nickname": "connectPostNamespacedPodAttach",
      "parameters": [
       {
        "type": "boolean",
        "paramType": "query",
        "name": "stdin",
        "description": "Stdin if true, redirects the standard input stream of the pod for this call. Defaults to false.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "stdout",
        "description": "Stdout if true indicates that stdout is to be redirected for the attach call. Defaults to true.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "stderr",
        "description": "Stderr if true indicates that stderr is to be redirected for the attach call. Defaults to true.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "tty",
        "description": "TTY if true indicates that a tty will be allocated for the attach call. This is passed through the container runtime so the tty is allocated on the worker node by the container runtime. Defaults to false.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "container",
        "description": "The container in which to execute the command. Defaults to only container if there is only one container in the pod.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/pods/{name}/binding",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.Binding",
      "method": "POST",
      "summary": "create binding of a Binding",
      "nickname": "createNamespacedBindingBinding",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Binding",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Binding",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Binding"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/pods/{name}/exec",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "string",
      "method": "GET",
      "summary": "connect GET requests to exec of Pod",
      "nickname": "connectGetNamespacedPodExec",
      "parameters": [
       {
        "type": "boolean",
        "paramType": "query",
        "name": "stdin",
        "description": "Redirect the standard input stream of the pod for this call. Defaults to false.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "stdout",
        "description": "Redirect the standard output stream of the pod for this call. Defaults to true.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "stderr",
        "description": "Redirect the standard error stream of the pod for this call. Defaults to true.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "tty",
        "description": "TTY if true indicates that a tty will be allocated for the exec call. Defaults to false.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "container",
        "description": "Container in which to execute the command. Defaults to only container if there is only one container in the pod.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "command",
        "description": "Command is the remote command to execute. argv array. Not executed within a shell.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "POST",
      "summary": "connect POST requests to exec of Pod",
      "nickname": "connectPostNamespacedPodExec",
      "parameters": [
       {
        "type": "boolean",
        "paramType": "query",
        "name": "stdin",
        "description": "Redirect the standard input stream of the pod for this call. Defaults to false.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "stdout",
        "description": "Redirect the standard output stream of the pod for this call. Defaults to true.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "stderr",
        "description": "Redirect the standard error stream of the pod for this call. Defaults to true.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "tty",
        "description": "TTY if true indicates that a tty will be allocated for the exec call. Defaults to false.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "container",
        "description": "Container in which to execute the command. Defaults to only container if there is only one container in the pod.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "command",
        "description": "Command is the remote command to execute. argv array. Not executed within a shell.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/pods/{name}/log",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.Pod",
      "method": "GET",
      "summary": "read log of the specified Pod",
      "nickname": "readNamespacedPodLog",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "container",
        "description": "The container for which to stream logs. Defaults to only container if there is one container in the pod.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "follow",
        "description": "Follow the log stream of the pod. Defaults to false.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "previous",
        "description": "Return previous terminated container logs. Defaults to false.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "sinceSeconds",
        "description": "A relative time in seconds before the current time from which to show logs. If this value precedes the time a pod was started, only logs since the pod start will be returned. If this value is in the future, no logs will be returned. Only one of sinceSeconds or sinceTime may be specified.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "sinceTime",
        "description": "An RFC3339 timestamp from which to show logs. If this value precedes the time a pod was started, only logs since the pod start will be returned. If this value is in the future, no logs will be returned. Only one of sinceSeconds or sinceTime may be specified.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "timestamps",
        "description": "If true, add an RFC3339 or RFC3339Nano timestamp at the beginning of every line of log output. Defaults to false.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "tailLines",
        "description": "If set, the number of lines from the end of the logs to show. If not specified, logs are shown from the creation of the container or sinceSeconds or sinceTime",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "limitBytes",
        "description": "If set, the number of bytes to read from the server before terminating the log output. This may not display a complete final line of logging, and may return slightly more or slightly less than the specified limit.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Pod"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/pods/{name}/portforward",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "string",
      "method": "GET",
      "summary": "connect GET requests to portforward of Pod",
      "nickname": "connectGetNamespacedPodPortforward",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "POST",
      "summary": "connect POST requests to portforward of Pod",
      "nickname": "connectPostNamespacedPodPortforward",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/pods/{name}/proxy",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "string",
      "method": "GET",
      "summary": "connect GET requests to proxy of Pod",
      "nickname": "connectGetNamespacedPodProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to pod.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "POST",
      "summary": "connect POST requests to proxy of Pod",
      "nickname": "connectPostNamespacedPodProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to pod.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "PUT",
      "summary": "connect PUT requests to proxy of Pod",
      "nickname": "connectPutNamespacedPodProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to pod.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "DELETE",
      "summary": "connect DELETE requests to proxy of Pod",
      "nickname": "connectDeleteNamespacedPodProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to pod.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "HEAD",
      "summary": "connect HEAD requests to proxy of Pod",
      "nickname": "connectHeadNamespacedPodProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to pod.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "OPTIONS",
      "summary": "connect OPTIONS requests to proxy of Pod",
      "nickname": "connectOptionsNamespacedPodProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to pod.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/pods/{name}/proxy/{path}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "string",
      "method": "GET",
      "summary": "connect GET requests to proxy of Pod",
      "nickname": "connectGetNamespacedPodProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to pod.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "POST",
      "summary": "connect POST requests to proxy of Pod",
      "nickname": "connectPostNamespacedPodProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to pod.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "PUT",
      "summary": "connect PUT requests to proxy of Pod",
      "nickname": "connectPutNamespacedPodProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to pod.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "DELETE",
      "summary": "connect DELETE requests to proxy of Pod",
      "nickname": "connectDeleteNamespacedPodProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to pod.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "HEAD",
      "summary": "connect HEAD requests to proxy of Pod",
      "nickname": "connectHeadNamespacedPodProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to pod.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "OPTIONS",
      "summary": "connect OPTIONS requests to proxy of Pod",
      "nickname": "connectOptionsNamespacedPodProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the URL path to use for the current proxy request to pod.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/pods/{name}/status",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.Pod",
      "method": "PUT",
      "summary": "replace status of the specified Pod",
      "nickname": "replaceNamespacedPodStatus",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Pod",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Pod",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Pod"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/podtemplates",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.PodTemplateList",
      "method": "GET",
      "summary": "list or watch objects of kind PodTemplate",
      "nickname": "listNamespacedPodTemplate",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PodTemplateList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.PodTemplate",
      "method": "POST",
      "summary": "create a PodTemplate",
      "nickname": "createNamespacedPodTemplate",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.PodTemplate",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PodTemplate"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete collection of PodTemplate",
      "nickname": "deletecollectionNamespacedPodTemplate",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/podtemplates",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of PodTemplate",
      "nickname": "watchNamespacedPodTemplateList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/podtemplates/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.PodTemplate",
      "method": "GET",
      "summary": "read the specified PodTemplate",
      "nickname": "readNamespacedPodTemplate",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "export",
        "description": "Should this value be exported.  Export strips fields that a user can not specify.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "exact",
        "description": "Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the PodTemplate",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PodTemplate"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.PodTemplate",
      "method": "PUT",
      "summary": "replace the specified PodTemplate",
      "nickname": "replaceNamespacedPodTemplate",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.PodTemplate",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the PodTemplate",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PodTemplate"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.PodTemplate",
      "method": "PATCH",
      "summary": "partially update the specified PodTemplate",
      "nickname": "patchNamespacedPodTemplate",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "unversioned.Patch",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the PodTemplate",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PodTemplate"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "application/json-patch+json",
       "application/merge-patch+json",
       "application/strategic-merge-patch+json"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete a PodTemplate",
      "nickname": "deleteNamespacedPodTemplate",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.DeleteOptions",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the PodTemplate",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/podtemplates/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch changes to an object of kind PodTemplate",
      "nickname": "watchNamespacedPodTemplate",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the PodTemplate",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/podtemplates",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.PodTemplateList",
      "method": "GET",
      "summary": "list or watch objects of kind PodTemplate",
      "nickname": "listPodTemplate",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.PodTemplateList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/podtemplates",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of PodTemplate",
      "nickname": "watchPodTemplateList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/replicationcontrollers",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.ReplicationControllerList",
      "method": "GET",
      "summary": "list or watch objects of kind ReplicationController",
      "nickname": "listNamespacedReplicationController",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ReplicationControllerList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.ReplicationController",
      "method": "POST",
      "summary": "create a ReplicationController",
      "nickname": "createNamespacedReplicationController",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.ReplicationController",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ReplicationController"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete collection of ReplicationController",
      "nickname": "deletecollectionNamespacedReplicationController",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/replicationcontrollers",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of ReplicationController",
      "nickname": "watchNamespacedReplicationControllerList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/replicationcontrollers/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.ReplicationController",
      "method": "GET",
      "summary": "read the specified ReplicationController",
      "nickname": "readNamespacedReplicationController",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "export",
        "description": "Should this value be exported.  Export strips fields that a user can not specify.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "exact",
        "description": "Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ReplicationController",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ReplicationController"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.ReplicationController",
      "method": "PUT",
      "summary": "replace the specified ReplicationController",
      "nickname": "replaceNamespacedReplicationController",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.ReplicationController",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ReplicationController",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ReplicationController"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.ReplicationController",
      "method": "PATCH",
      "summary": "partially update the specified ReplicationController",
      "nickname": "patchNamespacedReplicationController",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "unversioned.Patch",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ReplicationController",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ReplicationController"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "application/json-patch+json",
       "application/merge-patch+json",
       "application/strategic-merge-patch+json"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete a ReplicationController",
      "nickname": "deleteNamespacedReplicationController",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.DeleteOptions",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ReplicationController",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/replicationcontrollers/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch changes to an object of kind ReplicationController",
      "nickname": "watchNamespacedReplicationController",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ReplicationController",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/replicationcontrollers",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.ReplicationControllerList",
      "method": "GET",
      "summary": "list or watch objects of kind ReplicationController",
      "nickname": "listReplicationController",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ReplicationControllerList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/replicationcontrollers",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of ReplicationController",
      "nickname": "watchReplicationControllerList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/replicationcontrollers/{name}/scale",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.Scale",
      "method": "GET",
      "summary": "read scale of the specified Scale",
      "nickname": "readNamespacedScaleScale",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Scale",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Scale"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Scale",
      "method": "PUT",
      "summary": "replace scale of the specified Scale",
      "nickname": "replaceNamespacedScaleScale",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Scale",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Scale",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Scale"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Scale",
      "method": "PATCH",
      "summary": "partially update scale of the specified Scale",
      "nickname": "patchNamespacedScaleScale",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "unversioned.Patch",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Scale",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Scale"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "application/json-patch+json",
       "application/merge-patch+json",
       "application/strategic-merge-patch+json"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/replicationcontrollers/{name}/status",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.ReplicationController",
      "method": "PUT",
      "summary": "replace status of the specified ReplicationController",
      "nickname": "replaceNamespacedReplicationControllerStatus",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.ReplicationController",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ReplicationController",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ReplicationController"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/resourcequotas",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.ResourceQuotaList",
      "method": "GET",
      "summary": "list or watch objects of kind ResourceQuota",
      "nickname": "listNamespacedResourceQuota",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ResourceQuotaList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.ResourceQuota",
      "method": "POST",
      "summary": "create a ResourceQuota",
      "nickname": "createNamespacedResourceQuota",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.ResourceQuota",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ResourceQuota"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete collection of ResourceQuota",
      "nickname": "deletecollectionNamespacedResourceQuota",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/resourcequotas",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of ResourceQuota",
      "nickname": "watchNamespacedResourceQuotaList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/resourcequotas/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.ResourceQuota",
      "method": "GET",
      "summary": "read the specified ResourceQuota",
      "nickname": "readNamespacedResourceQuota",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "export",
        "description": "Should this value be exported.  Export strips fields that a user can not specify.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "exact",
        "description": "Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ResourceQuota",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ResourceQuota"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.ResourceQuota",
      "method": "PUT",
      "summary": "replace the specified ResourceQuota",
      "nickname": "replaceNamespacedResourceQuota",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.ResourceQuota",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ResourceQuota",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ResourceQuota"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.ResourceQuota",
      "method": "PATCH",
      "summary": "partially update the specified ResourceQuota",
      "nickname": "patchNamespacedResourceQuota",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "unversioned.Patch",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ResourceQuota",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ResourceQuota"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "application/json-patch+json",
       "application/merge-patch+json",
       "application/strategic-merge-patch+json"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete a ResourceQuota",
      "nickname": "deleteNamespacedResourceQuota",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.DeleteOptions",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ResourceQuota",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/resourcequotas/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch changes to an object of kind ResourceQuota",
      "nickname": "watchNamespacedResourceQuota",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ResourceQuota",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/resourcequotas",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.ResourceQuotaList",
      "method": "GET",
      "summary": "list or watch objects of kind ResourceQuota",
      "nickname": "listResourceQuota",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ResourceQuotaList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/resourcequotas",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of ResourceQuota",
      "nickname": "watchResourceQuotaList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/resourcequotas/{name}/status",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.ResourceQuota",
      "method": "PUT",
      "summary": "replace status of the specified ResourceQuota",
      "nickname": "replaceNamespacedResourceQuotaStatus",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.ResourceQuota",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ResourceQuota",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ResourceQuota"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/secrets",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.SecretList",
      "method": "GET",
      "summary": "list or watch objects of kind Secret",
      "nickname": "listNamespacedSecret",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.SecretList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Secret",
      "method": "POST",
      "summary": "create a Secret",
      "nickname": "createNamespacedSecret",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Secret",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Secret"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete collection of Secret",
      "nickname": "deletecollectionNamespacedSecret",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/secrets",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of Secret",
      "nickname": "watchNamespacedSecretList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/secrets/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.Secret",
      "method": "GET",
      "summary": "read the specified Secret",
      "nickname": "readNamespacedSecret",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "export",
        "description": "Should this value be exported.  Export strips fields that a user can not specify.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "exact",
        "description": "Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Secret",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Secret"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Secret",
      "method": "PUT",
      "summary": "replace the specified Secret",
      "nickname": "replaceNamespacedSecret",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Secret",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Secret",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Secret"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Secret",
      "method": "PATCH",
      "summary": "partially update the specified Secret",
      "nickname": "patchNamespacedSecret",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "unversioned.Patch",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Secret",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Secret"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "application/json-patch+json",
       "application/merge-patch+json",
       "application/strategic-merge-patch+json"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete a Secret",
      "nickname": "deleteNamespacedSecret",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.DeleteOptions",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Secret",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/secrets/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch changes to an object of kind Secret",
      "nickname": "watchNamespacedSecret",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Secret",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/secrets",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.SecretList",
      "method": "GET",
      "summary": "list or watch objects of kind Secret",
      "nickname": "listSecret",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.SecretList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/secrets",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of Secret",
      "nickname": "watchSecretList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/serviceaccounts",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.ServiceAccountList",
      "method": "GET",
      "summary": "list or watch objects of kind ServiceAccount",
      "nickname": "listNamespacedServiceAccount",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ServiceAccountList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.ServiceAccount",
      "method": "POST",
      "summary": "create a ServiceAccount",
      "nickname": "createNamespacedServiceAccount",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.ServiceAccount",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ServiceAccount"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete collection of ServiceAccount",
      "nickname": "deletecollectionNamespacedServiceAccount",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/serviceaccounts",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of ServiceAccount",
      "nickname": "watchNamespacedServiceAccountList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/serviceaccounts/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.ServiceAccount",
      "method": "GET",
      "summary": "read the specified ServiceAccount",
      "nickname": "readNamespacedServiceAccount",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "export",
        "description": "Should this value be exported.  Export strips fields that a user can not specify.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "exact",
        "description": "Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ServiceAccount",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ServiceAccount"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.ServiceAccount",
      "method": "PUT",
      "summary": "replace the specified ServiceAccount",
      "nickname": "replaceNamespacedServiceAccount",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.ServiceAccount",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ServiceAccount",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ServiceAccount"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.ServiceAccount",
      "method": "PATCH",
      "summary": "partially update the specified ServiceAccount",
      "nickname": "patchNamespacedServiceAccount",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "unversioned.Patch",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ServiceAccount",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ServiceAccount"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "application/json-patch+json",
       "application/merge-patch+json",
       "application/strategic-merge-patch+json"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete a ServiceAccount",
      "nickname": "deleteNamespacedServiceAccount",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.DeleteOptions",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ServiceAccount",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/serviceaccounts/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch changes to an object of kind ServiceAccount",
      "nickname": "watchNamespacedServiceAccount",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the ServiceAccount",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/serviceaccounts",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.ServiceAccountList",
      "method": "GET",
      "summary": "list or watch objects of kind ServiceAccount",
      "nickname": "listServiceAccount",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ServiceAccountList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/serviceaccounts",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of ServiceAccount",
      "nickname": "watchServiceAccountList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/services",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.ServiceList",
      "method": "GET",
      "summary": "list or watch objects of kind Service",
      "nickname": "listNamespacedService",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ServiceList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Service",
      "method": "POST",
      "summary": "create a Service",
      "nickname": "createNamespacedService",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Service",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Service"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/services",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of Service",
      "nickname": "watchNamespacedServiceList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/services/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.Service",
      "method": "GET",
      "summary": "read the specified Service",
      "nickname": "readNamespacedService",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "export",
        "description": "Should this value be exported.  Export strips fields that a user can not specify.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "exact",
        "description": "Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Service"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Service",
      "method": "PUT",
      "summary": "replace the specified Service",
      "nickname": "replaceNamespacedService",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Service",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Service"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "v1.Service",
      "method": "PATCH",
      "summary": "partially update the specified Service",
      "nickname": "patchNamespacedService",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "unversioned.Patch",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Service"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "application/json-patch+json",
       "application/merge-patch+json",
       "application/strategic-merge-patch+json"
      ]
     },
     {
      "type": "unversioned.Status",
      "method": "DELETE",
      "summary": "delete a Service",
      "nickname": "deleteNamespacedService",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "unversioned.Status"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/namespaces/{namespace}/services/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch changes to an object of kind Service",
      "nickname": "watchNamespacedService",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/proxy/namespaces/{namespace}/services/{name}/{path}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "string",
      "method": "GET",
      "summary": "proxy GET requests to Service",
      "nickname": "proxyGETNamespacedService",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "PUT",
      "summary": "proxy PUT requests to Service",
      "nickname": "proxyPUTNamespacedService",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "POST",
      "summary": "proxy POST requests to Service",
      "nickname": "proxyPOSTNamespacedService",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "DELETE",
      "summary": "proxy DELETE requests to Service",
      "nickname": "proxyDELETENamespacedService",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "HEAD",
      "summary": "proxy HEAD requests to Service",
      "nickname": "proxyHEADNamespacedService",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "OPTIONS",
      "summary": "proxy OPTIONS requests to Service",
      "nickname": "proxyOPTIONSNamespacedService",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/proxy/namespaces/{namespace}/services/{name}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "string",
      "method": "GET",
      "summary": "proxy GET requests to Service",
      "nickname": "proxyGETNamespacedService",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "PUT",
      "summary": "proxy PUT requests to Service",
      "nickname": "proxyPUTNamespacedService",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "POST",
      "summary": "proxy POST requests to Service",
      "nickname": "proxyPOSTNamespacedService",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "DELETE",
      "summary": "proxy DELETE requests to Service",
      "nickname": "proxyDELETENamespacedService",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "HEAD",
      "summary": "proxy HEAD requests to Service",
      "nickname": "proxyHEADNamespacedService",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "OPTIONS",
      "summary": "proxy OPTIONS requests to Service",
      "nickname": "proxyOPTIONSNamespacedService",
      "parameters": [
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/services",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.ServiceList",
      "method": "GET",
      "summary": "list or watch objects of kind Service",
      "nickname": "listService",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.ServiceList"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/watch/services",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "json.WatchEvent",
      "method": "GET",
      "summary": "watch individual changes to a list of Service",
      "nickname": "watchServiceList",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "labelSelector",
        "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "fieldSelector",
        "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "boolean",
        "paramType": "query",
        "name": "watch",
        "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "query",
        "name": "resourceVersion",
        "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "integer",
        "paramType": "query",
        "name": "timeoutSeconds",
        "description": "Timeout for the list/watch call.",
        "required": false,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "json.WatchEvent"
       }
      ],
      "produces": [
       "application/json"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/services/{name}/proxy",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "string",
      "method": "GET",
      "summary": "connect GET requests to proxy of Service",
      "nickname": "connectGetNamespacedServiceProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the part of URLs that include service endpoints, suffixes, and parameters to use for the current proxy request to service. For example, the whole request URL is http://localhost/api/v1/namespaces/kube-system/services/elasticsearch-logging/_search?q=user:kimchy. Path is _search?q=user:kimchy.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "POST",
      "summary": "connect POST requests to proxy of Service",
      "nickname": "connectPostNamespacedServiceProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the part of URLs that include service endpoints, suffixes, and parameters to use for the current proxy request to service. For example, the whole request URL is http://localhost/api/v1/namespaces/kube-system/services/elasticsearch-logging/_search?q=user:kimchy. Path is _search?q=user:kimchy.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "PUT",
      "summary": "connect PUT requests to proxy of Service",
      "nickname": "connectPutNamespacedServiceProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the part of URLs that include service endpoints, suffixes, and parameters to use for the current proxy request to service. For example, the whole request URL is http://localhost/api/v1/namespaces/kube-system/services/elasticsearch-logging/_search?q=user:kimchy. Path is _search?q=user:kimchy.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "DELETE",
      "summary": "connect DELETE requests to proxy of Service",
      "nickname": "connectDeleteNamespacedServiceProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the part of URLs that include service endpoints, suffixes, and parameters to use for the current proxy request to service. For example, the whole request URL is http://localhost/api/v1/namespaces/kube-system/services/elasticsearch-logging/_search?q=user:kimchy. Path is _search?q=user:kimchy.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "HEAD",
      "summary": "connect HEAD requests to proxy of Service",
      "nickname": "connectHeadNamespacedServiceProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the part of URLs that include service endpoints, suffixes, and parameters to use for the current proxy request to service. For example, the whole request URL is http://localhost/api/v1/namespaces/kube-system/services/elasticsearch-logging/_search?q=user:kimchy. Path is _search?q=user:kimchy.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "OPTIONS",
      "summary": "connect OPTIONS requests to proxy of Service",
      "nickname": "connectOptionsNamespacedServiceProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the part of URLs that include service endpoints, suffixes, and parameters to use for the current proxy request to service. For example, the whole request URL is http://localhost/api/v1/namespaces/kube-system/services/elasticsearch-logging/_search?q=user:kimchy. Path is _search?q=user:kimchy.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/services/{name}/proxy/{path}",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "string",
      "method": "GET",
      "summary": "connect GET requests to proxy of Service",
      "nickname": "connectGetNamespacedServiceProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the part of URLs that include service endpoints, suffixes, and parameters to use for the current proxy request to service. For example, the whole request URL is http://localhost/api/v1/namespaces/kube-system/services/elasticsearch-logging/_search?q=user:kimchy. Path is _search?q=user:kimchy.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "POST",
      "summary": "connect POST requests to proxy of Service",
      "nickname": "connectPostNamespacedServiceProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the part of URLs that include service endpoints, suffixes, and parameters to use for the current proxy request to service. For example, the whole request URL is http://localhost/api/v1/namespaces/kube-system/services/elasticsearch-logging/_search?q=user:kimchy. Path is _search?q=user:kimchy.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "PUT",
      "summary": "connect PUT requests to proxy of Service",
      "nickname": "connectPutNamespacedServiceProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the part of URLs that include service endpoints, suffixes, and parameters to use for the current proxy request to service. For example, the whole request URL is http://localhost/api/v1/namespaces/kube-system/services/elasticsearch-logging/_search?q=user:kimchy. Path is _search?q=user:kimchy.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "DELETE",
      "summary": "connect DELETE requests to proxy of Service",
      "nickname": "connectDeleteNamespacedServiceProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the part of URLs that include service endpoints, suffixes, and parameters to use for the current proxy request to service. For example, the whole request URL is http://localhost/api/v1/namespaces/kube-system/services/elasticsearch-logging/_search?q=user:kimchy. Path is _search?q=user:kimchy.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "HEAD",
      "summary": "connect HEAD requests to proxy of Service",
      "nickname": "connectHeadNamespacedServiceProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the part of URLs that include service endpoints, suffixes, and parameters to use for the current proxy request to service. For example, the whole request URL is http://localhost/api/v1/namespaces/kube-system/services/elasticsearch-logging/_search?q=user:kimchy. Path is _search?q=user:kimchy.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     },
     {
      "type": "string",
      "method": "OPTIONS",
      "summary": "connect OPTIONS requests to proxy of Service",
      "nickname": "connectOptionsNamespacedServiceProxy",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "path",
        "description": "Path is the part of URLs that include service endpoints, suffixes, and parameters to use for the current proxy request to service. For example, the whole request URL is http://localhost/api/v1/namespaces/kube-system/services/elasticsearch-logging/_search?q=user:kimchy. Path is _search?q=user:kimchy.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "path",
        "description": "path to the resource",
        "required": true,
        "allowMultiple": false
       }
      ],
      "produces": [
       "*/*"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1/namespaces/{namespace}/services/{name}/status",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "v1.Service",
      "method": "PUT",
      "summary": "replace status of the specified Service",
      "nickname": "replaceNamespacedServiceStatus",
      "parameters": [
       {
        "type": "string",
        "paramType": "query",
        "name": "pretty",
        "description": "If 'true', then the output is pretty printed.",
        "required": false,
        "allowMultiple": false
       },
       {
        "type": "v1.Service",
        "paramType": "body",
        "name": "body",
        "description": "",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "namespace",
        "description": "object name and auth scope, such as for teams and projects",
        "required": true,
        "allowMultiple": false
       },
       {
        "type": "string",
        "paramType": "path",
        "name": "name",
        "description": "name of the Service",
        "required": true,
        "allowMultiple": false
       }
      ],
      "responseMessages": [
       {
        "code": 200,
        "message": "OK",
        "responseModel": "v1.Service"
       }
      ],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "*/*"
      ]
     }
    ]
   },
   {
    "path": "/api/v1",
    "description": "API at /api/v1",
    "operations": [
     {
      "type": "void",
      "method": "GET",
      "summary": "get available resources",
      "nickname": "getAPIResources",
      "parameters": [],
      "produces": [
       "application/json",
       "application/yaml"
      ],
      "consumes": [
       "application/json",
       "application/yaml"
      ]
     }
    ]
   }
  ],
  "models": {
   "v1.Binding": {
    "id": "v1.Binding",
    "description": "Binding ties one object to another. For example, a pod is bound to a node by a scheduler.",
    "required": [
     "target"
    ],
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "v1.ObjectMeta",
      "description": "Standard object's metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "target": {
      "$ref": "v1.ObjectReference",
      "description": "The target object that you want to bind to the standard object."
     }
    }
   },
   "v1.ObjectMeta": {
    "id": "v1.ObjectMeta",
    "description": "ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.",
    "properties": {
     "name": {
      "type": "string",
      "description": "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://releases.k8s.io/release-1.2/docs/user-guide/identifiers.md#names"
     },
     "generateName": {
      "type": "string",
      "description": "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.\n\nIf this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header).\n\nApplied only if Name is not specified. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#idempotency"
     },
     "namespace": {
      "type": "string",
      "description": "Namespace defines the space within each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.\n\nMust be a DNS_LABEL. Cannot be updated. More info: http://releases.k8s.io/release-1.2/docs/user-guide/namespaces.md"
     },
     "selfLink": {
      "type": "string",
      "description": "SelfLink is a URL representing this object. Populated by the system. Read-only."
     },
     "uid": {
      "type": "string",
      "description": "UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.\n\nPopulated by the system. Read-only. More info: http://releases.k8s.io/release-1.2/docs/user-guide/identifiers.md#uids"
     },
     "resourceVersion": {
      "type": "string",
      "description": "An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.\n\nPopulated by the system. Read-only. Value must be treated as opaque by clients and . More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#concurrency-control-and-consistency"
     },
     "generation": {
      "type": "integer",
      "format": "int64",
      "description": "A sequence number representing a specific generation of the desired state. Populated by the system. Read-only."
     },
     "creationTimestamp": {
      "type": "string",
      "description": "CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.\n\nPopulated by the system. Read-only. Null for lists. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "deletionTimestamp": {
      "type": "string",
      "description": "DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is set by the server when a graceful deletion is requested by the user, and is not directly settable by a client. The resource will be deleted (no longer visible from resource lists, and not reachable by name) after the time in this field. Once set, this value may not be unset or be set further into the future, although it may be shortened or the resource may be deleted prior to this time. For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react by sending a graceful termination signal to the containers in the pod. Once the resource is deleted in the API, the Kubelet will send a hard termination signal to the container. If not set, graceful deletion of the object has not been requested.\n\nPopulated by the system when a graceful deletion is requested. Read-only. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "deletionGracePeriodSeconds": {
      "type": "integer",
      "format": "int64",
      "description": "Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only."
     },
     "labels": {
      "type": "any",
      "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://releases.k8s.io/release-1.2/docs/user-guide/labels.md"
     },
     "annotations": {
      "type": "any",
      "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://releases.k8s.io/release-1.2/docs/user-guide/annotations.md"
     }
    }
   },
   "v1.ObjectReference": {
    "id": "v1.ObjectReference",
    "description": "ObjectReference contains enough information to let you inspect or modify the referred object.",
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind of the referent. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "namespace": {
      "type": "string",
      "description": "Namespace of the referent. More info: http://releases.k8s.io/release-1.2/docs/user-guide/namespaces.md"
     },
     "name": {
      "type": "string",
      "description": "Name of the referent. More info: http://releases.k8s.io/release-1.2/docs/user-guide/identifiers.md#names"
     },
     "uid": {
      "type": "string",
      "description": "UID of the referent. More info: http://releases.k8s.io/release-1.2/docs/user-guide/identifiers.md#uids"
     },
     "apiVersion": {
      "type": "string",
      "description": "API version of the referent."
     },
     "resourceVersion": {
      "type": "string",
      "description": "Specific resourceVersion to which this reference is made, if any. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#concurrency-control-and-consistency"
     },
     "fieldPath": {
      "type": "string",
      "description": "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: \"spec.containers{name}\" (where \"name\" refers to the name of the container that triggered the event) or if no container name is specified \"spec.containers[2]\" (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object."
     }
    }
   },
   "v1.ComponentStatusList": {
    "id": "v1.ComponentStatusList",
    "description": "Status of all the conditions for the component as a list of ComponentStatus objects.",
    "required": [
     "items"
    ],
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "unversioned.ListMeta",
      "description": "Standard list metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "items": {
      "type": "array",
      "items": {
       "$ref": "v1.ComponentStatus"
      },
      "description": "List of ComponentStatus objects."
     }
    }
   },
   "unversioned.ListMeta": {
    "id": "unversioned.ListMeta",
    "description": "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
    "properties": {
     "selfLink": {
      "type": "string",
      "description": "SelfLink is a URL representing this object. Populated by the system. Read-only."
     },
     "resourceVersion": {
      "type": "string",
      "description": "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#concurrency-control-and-consistency"
     }
    }
   },
   "v1.ComponentStatus": {
    "id": "v1.ComponentStatus",
    "description": "ComponentStatus (and ComponentStatusList) holds the cluster validation info.",
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "v1.ObjectMeta",
      "description": "Standard object's metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "conditions": {
      "type": "array",
      "items": {
       "$ref": "v1.ComponentCondition"
      },
      "description": "List of component conditions observed"
     }
    }
   },
   "v1.ComponentCondition": {
    "id": "v1.ComponentCondition",
    "description": "Information about the condition of a component.",
    "required": [
     "type",
     "status"
    ],
    "properties": {
     "type": {
      "type": "string",
      "description": "Type of condition for a component. Valid value: \"Healthy\""
     },
     "status": {
      "type": "string",
      "description": "Status of the condition for a component. Valid values for \"Healthy\": \"True\", \"False\", or \"Unknown\"."
     },
     "message": {
      "type": "string",
      "description": "Message about the condition for a component. For example, information about a health check."
     },
     "error": {
      "type": "string",
      "description": "Condition error code for a component. For example, a health check error code."
     }
    }
   },
   "v1.ConfigMapList": {
    "id": "v1.ConfigMapList",
    "description": "ConfigMapList is a resource containing a list of ConfigMap objects.",
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "unversioned.ListMeta",
      "description": "More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "items": {
      "type": "array",
      "items": {
       "$ref": "v1.ConfigMap"
      },
      "description": "Items is the list of ConfigMaps."
     }
    }
   },
   "v1.ConfigMap": {
    "id": "v1.ConfigMap",
    "description": "ConfigMap holds configuration data for pods to consume.",
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "v1.ObjectMeta",
      "description": "Standard object's metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "data": {
      "type": "any",
      "description": "Data contains the configuration data. Each key must be a valid DNS_SUBDOMAIN with an optional leading dot."
     }
    }
   },
   "unversioned.Status": {
    "id": "unversioned.Status",
    "description": "Status is a return value for calls that don't return other objects.",
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "unversioned.ListMeta",
      "description": "Standard list metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "status": {
      "type": "string",
      "description": "Status of the operation. One of: \"Success\" or \"Failure\". More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#spec-and-status"
     },
     "message": {
      "type": "string",
      "description": "A human-readable description of the status of this operation."
     },
     "reason": {
      "type": "string",
      "description": "A machine-readable description of why this operation is in the \"Failure\" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it."
     },
     "details": {
      "$ref": "unversioned.StatusDetails",
      "description": "Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type."
     },
     "code": {
      "type": "integer",
      "format": "int32",
      "description": "Suggested HTTP return code for this status, 0 if not set."
     }
    }
   },
   "unversioned.StatusDetails": {
    "id": "unversioned.StatusDetails",
    "description": "StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.",
    "properties": {
     "name": {
      "type": "string",
      "description": "The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described)."
     },
     "group": {
      "type": "string",
      "description": "The group attribute of the resource associated with the status StatusReason."
     },
     "kind": {
      "type": "string",
      "description": "The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "causes": {
      "type": "array",
      "items": {
       "$ref": "unversioned.StatusCause"
      },
      "description": "The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes."
     },
     "retryAfterSeconds": {
      "type": "integer",
      "format": "int32",
      "description": "If specified, the time in seconds before the operation should be retried."
     }
    }
   },
   "unversioned.StatusCause": {
    "id": "unversioned.StatusCause",
    "description": "StatusCause provides more information about an api.Status failure, including cases when multiple errors are encountered.",
    "properties": {
     "reason": {
      "type": "string",
      "description": "A machine-readable description of the cause of the error. If this value is empty there is no information available."
     },
     "message": {
      "type": "string",
      "description": "A human-readable description of the cause of the error.  This field may be presented as-is to a reader."
     },
     "field": {
      "type": "string",
      "description": "The field of the resource that has caused this error, as named by its JSON serialization. May include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may appear more than once in an array of causes due to fields having multiple errors. Optional.\n\nExamples:\n  \"name\" - the field \"name\" on the current resource\n  \"items[0].name\" - the field \"name\" on the first array entry in \"items\""
     }
    }
   },
   "json.WatchEvent": {
    "id": "json.WatchEvent",
    "properties": {
     "type": {
      "type": "string",
      "description": "the type of watch event; may be ADDED, MODIFIED, DELETED, or ERROR"
     },
     "object": {
      "type": "string",
      "description": "the object being watched; will match the type of the resource endpoint or be a Status object if the type is ERROR"
     }
    }
   },
   "unversioned.Patch": {
    "id": "unversioned.Patch",
    "description": "Patch is provided to give a concrete name and type to the Kubernetes PATCH request body.",
    "properties": {}
   },
   "v1.DeleteOptions": {
    "id": "v1.DeleteOptions",
    "description": "DeleteOptions may be provided when deleting an API object",
    "required": [
     "gracePeriodSeconds"
    ],
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "gracePeriodSeconds": {
      "type": "integer",
      "format": "int64",
      "description": "The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately."
     }
    }
   },
   "v1.EndpointsList": {
    "id": "v1.EndpointsList",
    "description": "EndpointsList is a list of endpoints.",
    "required": [
     "items"
    ],
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "unversioned.ListMeta",
      "description": "Standard list metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "items": {
      "type": "array",
      "items": {
       "$ref": "v1.Endpoints"
      },
      "description": "List of endpoints."
     }
    }
   },
   "v1.Endpoints": {
    "id": "v1.Endpoints",
    "description": "Endpoints is a collection of endpoints that implement the actual service. Example:\n  Name: \"mysvc\",\n  Subsets: [\n    {\n      Addresses: [{\"ip\": \"10.10.1.1\"}, {\"ip\": \"10.10.2.2\"}],\n      Ports: [{\"name\": \"a\", \"port\": 8675}, {\"name\": \"b\", \"port\": 309}]\n    },\n    {\n      Addresses: [{\"ip\": \"10.10.3.3\"}],\n      Ports: [{\"name\": \"a\", \"port\": 93}, {\"name\": \"b\", \"port\": 76}]\n    },\n ]",
    "required": [
     "subsets"
    ],
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "v1.ObjectMeta",
      "description": "Standard object's metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "subsets": {
      "type": "array",
      "items": {
       "$ref": "v1.EndpointSubset"
      },
      "description": "The set of all endpoints is the union of all subsets. Addresses are placed into subsets according to the IPs they share. A single address with multiple ports, some of which are ready and some of which are not (because they come from different containers) will result in the address being displayed in different subsets for the different ports. No address will appear in both Addresses and NotReadyAddresses in the same subset. Sets of addresses and ports that comprise a service."
     }
    }
   },
   "v1.EndpointSubset": {
    "id": "v1.EndpointSubset",
    "description": "EndpointSubset is a group of addresses with a common set of ports. The expanded set of endpoints is the Cartesian product of Addresses x Ports. For example, given:\n  {\n    Addresses: [{\"ip\": \"10.10.1.1\"}, {\"ip\": \"10.10.2.2\"}],\n    Ports:     [{\"name\": \"a\", \"port\": 8675}, {\"name\": \"b\", \"port\": 309}]\n  }\nThe resulting set of endpoints can be viewed as:\n    a: [ 10.10.1.1:8675, 10.10.2.2:8675 ],\n    b: [ 10.10.1.1:309, 10.10.2.2:309 ]",
    "properties": {
     "addresses": {
      "type": "array",
      "items": {
       "$ref": "v1.EndpointAddress"
      },
      "description": "IP addresses which offer the related ports that are marked as ready. These endpoints should be considered safe for load balancers and clients to utilize."
     },
     "notReadyAddresses": {
      "type": "array",
      "items": {
       "$ref": "v1.EndpointAddress"
      },
      "description": "IP addresses which offer the related ports but are not currently marked as ready because they have not yet finished starting, have recently failed a readiness check, or have recently failed a liveness check."
     },
     "ports": {
      "type": "array",
      "items": {
       "$ref": "v1.EndpointPort"
      },
      "description": "Port numbers available on the related IP addresses."
     }
    }
   },
   "v1.EndpointAddress": {
    "id": "v1.EndpointAddress",
    "description": "EndpointAddress is a tuple that describes single IP address.",
    "required": [
     "ip"
    ],
    "properties": {
     "ip": {
      "type": "string",
      "description": "The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24)."
     },
     "targetRef": {
      "$ref": "v1.ObjectReference",
      "description": "Reference to object providing the endpoint."
     }
    }
   },
   "v1.EndpointPort": {
    "id": "v1.EndpointPort",
    "description": "EndpointPort is a tuple that describes a single port.",
    "required": [
     "port"
    ],
    "properties": {
     "name": {
      "type": "string",
      "description": "The name of this port (corresponds to ServicePort.Name). Must be a DNS_LABEL. Optional only if one port is defined."
     },
     "port": {
      "type": "integer",
      "format": "int32",
      "description": "The port number of the endpoint."
     },
     "protocol": {
      "type": "string",
      "description": "The IP protocol for this port. Must be UDP or TCP. Default is TCP."
     }
    }
   },
   "v1.EventList": {
    "id": "v1.EventList",
    "description": "EventList is a list of events.",
    "required": [
     "items"
    ],
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "unversioned.ListMeta",
      "description": "Standard list metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "items": {
      "type": "array",
      "items": {
       "$ref": "v1.Event"
      },
      "description": "List of events"
     }
    }
   },
   "v1.Event": {
    "id": "v1.Event",
    "description": "Event is a report of an event somewhere in the cluster.",
    "required": [
     "metadata",
     "involvedObject"
    ],
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "v1.ObjectMeta",
      "description": "Standard object's metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "involvedObject": {
      "$ref": "v1.ObjectReference",
      "description": "The object that this event is about."
     },
     "reason": {
      "type": "string",
      "description": "This should be a short, machine understandable string that gives the reason for the transition into the object's current status."
     },
     "message": {
      "type": "string",
      "description": "A human-readable description of the status of this operation."
     },
     "source": {
      "$ref": "v1.EventSource",
      "description": "The component reporting this event. Should be a short machine understandable string."
     },
     "firstTimestamp": {
      "type": "string",
      "description": "The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)"
     },
     "lastTimestamp": {
      "type": "string",
      "description": "The time at which the most recent occurrence of this event was recorded."
     },
     "count": {
      "type": "integer",
      "format": "int32",
      "description": "The number of times this event has occurred."
     },
     "type": {
      "type": "string",
      "description": "Type of this event (Normal, Warning), new types could be added in the future"
     }
    }
   },
   "v1.EventSource": {
    "id": "v1.EventSource",
    "description": "EventSource contains information for an event.",
    "properties": {
     "component": {
      "type": "string",
      "description": "Component from which the event is generated."
     },
     "host": {
      "type": "string",
      "description": "Host name on which the event is generated."
     }
    }
   },
   "v1.LimitRangeList": {
    "id": "v1.LimitRangeList",
    "description": "LimitRangeList is a list of LimitRange items.",
    "required": [
     "items"
    ],
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "unversioned.ListMeta",
      "description": "Standard list metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "items": {
      "type": "array",
      "items": {
       "$ref": "v1.LimitRange"
      },
      "description": "Items is a list of LimitRange objects. More info: http://releases.k8s.io/release-1.2/docs/design/admission_control_limit_range.md"
     }
    }
   },
   "v1.LimitRange": {
    "id": "v1.LimitRange",
    "description": "LimitRange sets resource usage limits for each kind of resource in a Namespace.",
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "v1.ObjectMeta",
      "description": "Standard object's metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "spec": {
      "$ref": "v1.LimitRangeSpec",
      "description": "Spec defines the limits enforced. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#spec-and-status"
     }
    }
   },
   "v1.LimitRangeSpec": {
    "id": "v1.LimitRangeSpec",
    "description": "LimitRangeSpec defines a min/max usage limit for resources that match on kind.",
    "required": [
     "limits"
    ],
    "properties": {
     "limits": {
      "type": "array",
      "items": {
       "$ref": "v1.LimitRangeItem"
      },
      "description": "Limits is the list of LimitRangeItem objects that are enforced."
     }
    }
   },
   "v1.LimitRangeItem": {
    "id": "v1.LimitRangeItem",
    "description": "LimitRangeItem defines a min/max usage limit for any resource that matches on kind.",
    "properties": {
     "type": {
      "type": "string",
      "description": "Type of resource that this limit applies to."
     },
     "max": {
      "type": "any",
      "description": "Max usage constraints on this kind by resource name."
     },
     "min": {
      "type": "any",
      "description": "Min usage constraints on this kind by resource name."
     },
     "default": {
      "type": "any",
      "description": "Default resource requirement limit value by resource name if resource limit is omitted."
     },
     "defaultRequest": {
      "type": "any",
      "description": "DefaultRequest is the default resource requirement request value by resource name if resource request is omitted."
     },
     "maxLimitRequestRatio": {
      "type": "any",
      "description": "MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource."
     }
    }
   },
   "v1.NamespaceList": {
    "id": "v1.NamespaceList",
    "description": "NamespaceList is a list of Namespaces.",
    "required": [
     "items"
    ],
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "unversioned.ListMeta",
      "description": "Standard list metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "items": {
      "type": "array",
      "items": {
       "$ref": "v1.Namespace"
      },
      "description": "Items is the list of Namespace objects in the list. More info: http://releases.k8s.io/release-1.2/docs/user-guide/namespaces.md"
     }
    }
   },
   "v1.Namespace": {
    "id": "v1.Namespace",
    "description": "Namespace provides a scope for Names. Use of multiple namespaces is optional.",
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "v1.ObjectMeta",
      "description": "Standard object's metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "spec": {
      "$ref": "v1.NamespaceSpec",
      "description": "Spec defines the behavior of the Namespace. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#spec-and-status"
     },
     "status": {
      "$ref": "v1.NamespaceStatus",
      "description": "Status describes the current status of a Namespace. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#spec-and-status"
     }
    }
   },
   "v1.NamespaceSpec": {
    "id": "v1.NamespaceSpec",
    "description": "NamespaceSpec describes the attributes on a Namespace.",
    "properties": {
     "finalizers": {
      "type": "array",
      "items": {
       "$ref": "v1.FinalizerName"
      },
      "description": "Finalizers is an opaque list of values that must be empty to permanently remove object from storage. More info: http://releases.k8s.io/release-1.2/docs/design/namespaces.md#finalizers"
     }
    }
   },
   "v1.FinalizerName": {
    "id": "v1.FinalizerName",
    "properties": {}
   },
   "v1.NamespaceStatus": {
    "id": "v1.NamespaceStatus",
    "description": "NamespaceStatus is information about the current status of a Namespace.",
    "properties": {
     "phase": {
      "type": "string",
      "description": "Phase is the current lifecycle phase of the namespace. More info: http://releases.k8s.io/release-1.2/docs/design/namespaces.md#phases"
     }
    }
   },
   "v1.NodeList": {
    "id": "v1.NodeList",
    "description": "NodeList is the whole list of all Nodes which have been registered with master.",
    "required": [
     "items"
    ],
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "unversioned.ListMeta",
      "description": "Standard list metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "items": {
      "type": "array",
      "items": {
       "$ref": "v1.Node"
      },
      "description": "List of nodes"
     }
    }
   },
   "v1.Node": {
    "id": "v1.Node",
    "description": "Node is a worker node in Kubernetes, formerly known as minion. Each node will have a unique identifier in the cache (i.e. in etcd).",
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "v1.ObjectMeta",
      "description": "Standard object's metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "spec": {
      "$ref": "v1.NodeSpec",
      "description": "Spec defines the behavior of a node. http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#spec-and-status"
     },
     "status": {
      "$ref": "v1.NodeStatus",
      "description": "Most recently observed status of the node. Populated by the system. Read-only. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#spec-and-status"
     }
    }
   },
   "v1.NodeSpec": {
    "id": "v1.NodeSpec",
    "description": "NodeSpec describes the attributes that a node is created with.",
    "properties": {
     "podCIDR": {
      "type": "string",
      "description": "PodCIDR represents the pod IP range assigned to the node."
     },
     "externalID": {
      "type": "string",
      "description": "External ID of the node assigned by some machine database (e.g. a cloud provider). Deprecated."
     },
     "providerID": {
      "type": "string",
      "description": "ID of the node assigned by the cloud provider in the format: \u003cProviderName\u003e://\u003cProviderSpecificNodeID\u003e"
     },
     "unschedulable": {
      "type": "boolean",
      "description": "Unschedulable controls node schedulability of new pods. By default, node is schedulable. More info: http://releases.k8s.io/release-1.2/docs/admin/node.md#manual-node-administration\""
     }
    }
   },
   "v1.NodeStatus": {
    "id": "v1.NodeStatus",
    "description": "NodeStatus is information about the current status of a node.",
    "required": [
     "images"
    ],
    "properties": {
     "capacity": {
      "type": "any",
      "description": "Capacity represents the total resources of a node. More info: http://releases.k8s.io/release-1.2/docs/user-guide/persistent-volumes.md#capacity for more details."
     },
     "allocatable": {
      "type": "any",
      "description": "Allocatable represents the resources of a node that are available for scheduling. Defaults to Capacity."
     },
     "phase": {
      "type": "string",
      "description": "NodePhase is the recently observed lifecycle phase of the node. More info: http://releases.k8s.io/release-1.2/docs/admin/node.md#node-phase"
     },
     "conditions": {
      "type": "array",
      "items": {
       "$ref": "v1.NodeCondition"
      },
      "description": "Conditions is an array of current observed node conditions. More info: http://releases.k8s.io/release-1.2/docs/admin/node.md#node-condition"
     },
     "addresses": {
      "type": "array",
      "items": {
       "$ref": "v1.NodeAddress"
      },
      "description": "List of addresses reachable to the node. Queried from cloud provider, if available. More info: http://releases.k8s.io/release-1.2/docs/admin/node.md#node-addresses"
     },
     "daemonEndpoints": {
      "$ref": "v1.NodeDaemonEndpoints",
      "description": "Endpoints of daemons running on the Node."
     },
     "nodeInfo": {
      "$ref": "v1.NodeSystemInfo",
      "description": "Set of ids/uuids to uniquely identify the node. More info: http://releases.k8s.io/release-1.2/docs/admin/node.md#node-info"
     },
     "images": {
      "type": "array",
      "items": {
       "$ref": "v1.ContainerImage"
      },
      "description": "List of container images on this node"
     }
    }
   },
   "v1.NodeCondition": {
    "id": "v1.NodeCondition",
    "description": "NodeCondition contains condition infromation for a node.",
    "required": [
     "type",
     "status"
    ],
    "properties": {
     "type": {
      "type": "string",
      "description": "Type of node condition."
     },
     "status": {
      "type": "string",
      "description": "Status of the condition, one of True, False, Unknown."
     },
     "lastHeartbeatTime": {
      "type": "string",
      "description": "Last time we got an update on a given condition."
     },
     "lastTransitionTime": {
      "type": "string",
      "description": "Last time the condition transit from one status to another."
     },
     "reason": {
      "type": "string",
      "description": "(brief) reason for the condition's last transition."
     },
     "message": {
      "type": "string",
      "description": "Human readable message indicating details about last transition."
     }
    }
   },
   "v1.NodeAddress": {
    "id": "v1.NodeAddress",
    "description": "NodeAddress contains information for the node's address.",
    "required": [
     "type",
     "address"
    ],
    "properties": {
     "type": {
      "type": "string",
      "description": "Node address type, one of Hostname, ExternalIP or InternalIP."
     },
     "address": {
      "type": "string",
      "description": "The node address."
     }
    }
   },
   "v1.NodeDaemonEndpoints": {
    "id": "v1.NodeDaemonEndpoints",
    "description": "NodeDaemonEndpoints lists ports opened by daemons running on the Node.",
    "properties": {
     "kubeletEndpoint": {
      "$ref": "v1.DaemonEndpoint",
      "description": "Endpoint on which Kubelet is listening."
     }
    }
   },
   "v1.DaemonEndpoint": {
    "id": "v1.DaemonEndpoint",
    "description": "DaemonEndpoint contains information about a single Daemon endpoint.",
    "required": [
     "Port"
    ],
    "properties": {
     "Port": {
      "type": "integer",
      "format": "int32",
      "description": "Port number of the given endpoint."
     }
    }
   },
   "v1.NodeSystemInfo": {
    "id": "v1.NodeSystemInfo",
    "description": "NodeSystemInfo is a set of ids/uuids to uniquely identify the node.",
    "required": [
     "machineID",
     "systemUUID",
     "bootID",
     "kernelVersion",
     "osImage",
     "containerRuntimeVersion",
     "kubeletVersion",
     "kubeProxyVersion"
    ],
    "properties": {
     "machineID": {
      "type": "string",
      "description": "Machine ID reported by the node."
     },
     "systemUUID": {
      "type": "string",
      "description": "System UUID reported by the node."
     },
     "bootID": {
      "type": "string",
      "description": "Boot ID reported by the node."
     },
     "kernelVersion": {
      "type": "string",
      "description": "Kernel Version reported by the node from 'uname -r' (e.g. 3.16.0-0.bpo.4-amd64)."
     },
     "osImage": {
      "type": "string",
      "description": "OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy))."
     },
     "containerRuntimeVersion": {
      "type": "string",
      "description": "ContainerRuntime Version reported by the node through runtime remote API (e.g. docker://1.5.0)."
     },
     "kubeletVersion": {
      "type": "string",
      "description": "Kubelet Version reported by the node."
     },
     "kubeProxyVersion": {
      "type": "string",
      "description": "KubeProxy Version reported by the node."
     }
    }
   },
   "v1.ContainerImage": {
    "id": "v1.ContainerImage",
    "description": "Describe a container image",
    "required": [
     "names"
    ],
    "properties": {
     "names": {
      "type": "array",
      "items": {
       "type": "string"
      },
      "description": "Names by which this image is known. e.g. [\"gcr.io/google_containers/hyperkube:v1.0.7\", \"dockerhub.io/google_containers/hyperkube:v1.0.7\"]"
     },
     "sizeBytes": {
      "type": "integer",
      "format": "int64",
      "description": "The size of the image in bytes."
     }
    }
   },
   "v1.PersistentVolumeClaimList": {
    "id": "v1.PersistentVolumeClaimList",
    "description": "PersistentVolumeClaimList is a list of PersistentVolumeClaim items.",
    "required": [
     "items"
    ],
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "unversioned.ListMeta",
      "description": "Standard list metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "items": {
      "type": "array",
      "items": {
       "$ref": "v1.PersistentVolumeClaim"
      },
      "description": "A list of persistent volume claims. More info: http://releases.k8s.io/release-1.2/docs/user-guide/persistent-volumes.md#persistentvolumeclaims"
     }
    }
   },
   "v1.PersistentVolumeClaim": {
    "id": "v1.PersistentVolumeClaim",
    "description": "PersistentVolumeClaim is a user's request for and claim to a persistent volume",
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "v1.ObjectMeta",
      "description": "Standard object's metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "spec": {
      "$ref": "v1.PersistentVolumeClaimSpec",
      "description": "Spec defines the desired characteristics of a volume requested by a pod author. More info: http://releases.k8s.io/release-1.2/docs/user-guide/persistent-volumes.md#persistentvolumeclaims"
     },
     "status": {
      "$ref": "v1.PersistentVolumeClaimStatus",
      "description": "Status represents the current information/status of a persistent volume claim. Read-only. More info: http://releases.k8s.io/release-1.2/docs/user-guide/persistent-volumes.md#persistentvolumeclaims"
     }
    }
   },
   "v1.PersistentVolumeClaimSpec": {
    "id": "v1.PersistentVolumeClaimSpec",
    "description": "PersistentVolumeClaimSpec describes the common attributes of storage devices and allows a Source for provider-specific attributes",
    "properties": {
     "accessModes": {
      "type": "array",
      "items": {
       "$ref": "v1.PersistentVolumeAccessMode"
      },
      "description": "AccessModes contains the desired access modes the volume should have. More info: http://releases.k8s.io/release-1.2/docs/user-guide/persistent-volumes.md#access-modes-1"
     },
     "resources": {
      "$ref": "v1.ResourceRequirements",
      "description": "Resources represents the minimum resources the volume should have. More info: http://releases.k8s.io/release-1.2/docs/user-guide/persistent-volumes.md#resources"
     },
     "volumeName": {
      "type": "string",
      "description": "VolumeName is the binding reference to the PersistentVolume backing this claim."
     }
    }
   },
   "v1.PersistentVolumeAccessMode": {
    "id": "v1.PersistentVolumeAccessMode",
    "properties": {}
   },
   "v1.ResourceRequirements": {
    "id": "v1.ResourceRequirements",
    "description": "ResourceRequirements describes the compute resource requirements.",
    "properties": {
     "limits": {
      "type": "any",
      "description": "Limits describes the maximum amount of compute resources allowed. More info: http://releases.k8s.io/release-1.2/docs/design/resources.md#resource-specifications"
     },
     "requests": {
      "type": "any",
      "description": "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: http://releases.k8s.io/release-1.2/docs/design/resources.md#resource-specifications"
     }
    }
   },
   "v1.PersistentVolumeClaimStatus": {
    "id": "v1.PersistentVolumeClaimStatus",
    "description": "PersistentVolumeClaimStatus is the current status of a persistent volume claim.",
    "properties": {
     "phase": {
      "type": "string",
      "description": "Phase represents the current phase of PersistentVolumeClaim."
     },
     "accessModes": {
      "type": "array",
      "items": {
       "$ref": "v1.PersistentVolumeAccessMode"
      },
      "description": "AccessModes contains the actual access modes the volume backing the PVC has. More info: http://releases.k8s.io/release-1.2/docs/user-guide/persistent-volumes.md#access-modes-1"
     },
     "capacity": {
      "type": "any",
      "description": "Represents the actual resources of the underlying volume."
     }
    }
   },
   "v1.PersistentVolumeList": {
    "id": "v1.PersistentVolumeList",
    "description": "PersistentVolumeList is a list of PersistentVolume items.",
    "required": [
     "items"
    ],
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "unversioned.ListMeta",
      "description": "Standard list metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "items": {
      "type": "array",
      "items": {
       "$ref": "v1.PersistentVolume"
      },
      "description": "List of persistent volumes. More info: http://releases.k8s.io/release-1.2/docs/user-guide/persistent-volumes.md"
     }
    }
   },
   "v1.PersistentVolume": {
    "id": "v1.PersistentVolume",
    "description": "PersistentVolume (PV) is a storage resource provisioned by an administrator. It is analogous to a node. More info: http://releases.k8s.io/release-1.2/docs/user-guide/persistent-volumes.md",
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "v1.ObjectMeta",
      "description": "Standard object's metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "spec": {
      "$ref": "v1.PersistentVolumeSpec",
      "description": "Spec defines a specification of a persistent volume owned by the cluster. Provisioned by an administrator. More info: http://releases.k8s.io/release-1.2/docs/user-guide/persistent-volumes.md#persistent-volumes"
     },
     "status": {
      "$ref": "v1.PersistentVolumeStatus",
      "description": "Status represents the current information/status for the persistent volume. Populated by the system. Read-only. More info: http://releases.k8s.io/release-1.2/docs/user-guide/persistent-volumes.md#persistent-volumes"
     }
    }
   },
   "v1.PersistentVolumeSpec": {
    "id": "v1.PersistentVolumeSpec",
    "description": "PersistentVolumeSpec is the specification of a persistent volume.",
    "properties": {
     "capacity": {
      "type": "any",
      "description": "A description of the persistent volume's resources and capacity. More info: http://releases.k8s.io/release-1.2/docs/user-guide/persistent-volumes.md#capacity"
     },
     "gcePersistentDisk": {
      "$ref": "v1.GCEPersistentDiskVolumeSource",
      "description": "GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#gcepersistentdisk"
     },
     "awsElasticBlockStore": {
      "$ref": "v1.AWSElasticBlockStoreVolumeSource",
      "description": "AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#awselasticblockstore"
     },
     "hostPath": {
      "$ref": "v1.HostPathVolumeSource",
      "description": "HostPath represents a directory on the host. Provisioned by a developer or tester. This is useful for single-node development and testing only! On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#hostpath"
     },
     "glusterfs": {
      "$ref": "v1.GlusterfsVolumeSource",
      "description": "Glusterfs represents a Glusterfs volume that is attached to a host and exposed to the pod. Provisioned by an admin. More info: http://releases.k8s.io/release-1.2/examples/glusterfs/README.md"
     },
     "nfs": {
      "$ref": "v1.NFSVolumeSource",
      "description": "NFS represents an NFS mount on the host. Provisioned by an admin. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#nfs"
     },
     "rbd": {
      "$ref": "v1.RBDVolumeSource",
      "description": "RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: http://releases.k8s.io/release-1.2/examples/rbd/README.md"
     },
     "iscsi": {
      "$ref": "v1.ISCSIVolumeSource",
      "description": "ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin."
     },
     "cinder": {
      "$ref": "v1.CinderVolumeSource",
      "description": "Cinder represents a cinder volume attached and mounted on kubelets host machine More info: http://releases.k8s.io/release-1.2/examples/mysql-cinder-pd/README.md"
     },
     "cephfs": {
      "$ref": "v1.CephFSVolumeSource",
      "description": "CephFS represents a Ceph FS mount on the host that shares a pod's lifetime"
     },
     "fc": {
      "$ref": "v1.FCVolumeSource",
      "description": "FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod."
     },
     "flocker": {
      "$ref": "v1.FlockerVolumeSource",
      "description": "Flocker represents a Flocker volume attached to a kubelet's host machine and exposed to the pod for its usage. This depends on the Flocker control service being running"
     },
     "flexVolume": {
      "$ref": "v1.FlexVolumeSource",
      "description": "FlexVolume represents a generic volume resource that is provisioned/attached using a exec based plugin. This is an alpha feature and may change in future."
     },
     "azureFile": {
      "$ref": "v1.AzureFileVolumeSource",
      "description": "AzureFile represents an Azure File Service mount on the host and bind mount to the pod."
     },
     "accessModes": {
      "type": "array",
      "items": {
       "$ref": "v1.PersistentVolumeAccessMode"
      },
      "description": "AccessModes contains all ways the volume can be mounted. More info: http://releases.k8s.io/release-1.2/docs/user-guide/persistent-volumes.md#access-modes"
     },
     "claimRef": {
      "$ref": "v1.ObjectReference",
      "description": "ClaimRef is part of a bi-directional binding between PersistentVolume and PersistentVolumeClaim. Expected to be non-nil when bound. claim.VolumeName is the authoritative bind between PV and PVC. More info: http://releases.k8s.io/release-1.2/docs/user-guide/persistent-volumes.md#binding"
     },
     "persistentVolumeReclaimPolicy": {
      "type": "string",
      "description": "What happens to a persistent volume when released from its claim. Valid options are Retain (default) and Recycle. Recyling must be supported by the volume plugin underlying this persistent volume. More info: http://releases.k8s.io/release-1.2/docs/user-guide/persistent-volumes.md#recycling-policy"
     }
    }
   },
   "v1.GCEPersistentDiskVolumeSource": {
    "id": "v1.GCEPersistentDiskVolumeSource",
    "description": "Represents a Persistent Disk resource in Google Compute Engine.\n\nA GCE PD must exist before mounting to a container. The disk must also be in the same GCE project and zone as the kubelet. A GCE PD can only be mounted as read/write once or read-only many times. GCE PDs support ownership management and SELinux relabeling.",
    "required": [
     "pdName"
    ],
    "properties": {
     "pdName": {
      "type": "string",
      "description": "Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#gcepersistentdisk"
     },
     "fsType": {
      "type": "string",
      "description": "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#gcepersistentdisk"
     },
     "partition": {
      "type": "integer",
      "format": "int32",
      "description": "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as \"1\". Similarly, the volume partition for /dev/sda is \"0\" (or you can leave the property empty). More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#gcepersistentdisk"
     },
     "readOnly": {
      "type": "boolean",
      "description": "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#gcepersistentdisk"
     }
    }
   },
   "v1.AWSElasticBlockStoreVolumeSource": {
    "id": "v1.AWSElasticBlockStoreVolumeSource",
    "description": "Represents a Persistent Disk resource in AWS.\n\nAn AWS EBS disk must exist before mounting to a container. The disk must also be in the same AWS zone as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS EBS volumes support ownership management and SELinux relabeling.",
    "required": [
     "volumeID"
    ],
    "properties": {
     "volumeID": {
      "type": "string",
      "description": "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#awselasticblockstore"
     },
     "fsType": {
      "type": "string",
      "description": "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#awselasticblockstore"
     },
     "partition": {
      "type": "integer",
      "format": "int32",
      "description": "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as \"1\". Similarly, the volume partition for /dev/sda is \"0\" (or you can leave the property empty)."
     },
     "readOnly": {
      "type": "boolean",
      "description": "Specify \"true\" to force and set the ReadOnly property in VolumeMounts to \"true\". If omitted, the default is \"false\". More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#awselasticblockstore"
     }
    }
   },
   "v1.HostPathVolumeSource": {
    "id": "v1.HostPathVolumeSource",
    "description": "Represents a host path mapped into a pod. Host path volumes do not support ownership management or SELinux relabeling.",
    "required": [
     "path"
    ],
    "properties": {
     "path": {
      "type": "string",
      "description": "Path of the directory on the host. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#hostpath"
     }
    }
   },
   "v1.GlusterfsVolumeSource": {
    "id": "v1.GlusterfsVolumeSource",
    "description": "Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or SELinux relabeling.",
    "required": [
     "endpoints",
     "path"
    ],
    "properties": {
     "endpoints": {
      "type": "string",
      "description": "EndpointsName is the endpoint name that details Glusterfs topology. More info: http://releases.k8s.io/release-1.2/examples/glusterfs/README.md#create-a-pod"
     },
     "path": {
      "type": "string",
      "description": "Path is the Glusterfs volume path. More info: http://releases.k8s.io/release-1.2/examples/glusterfs/README.md#create-a-pod"
     },
     "readOnly": {
      "type": "boolean",
      "description": "ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: http://releases.k8s.io/release-1.2/examples/glusterfs/README.md#create-a-pod"
     }
    }
   },
   "v1.NFSVolumeSource": {
    "id": "v1.NFSVolumeSource",
    "description": "Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.",
    "required": [
     "server",
     "path"
    ],
    "properties": {
     "server": {
      "type": "string",
      "description": "Server is the hostname or IP address of the NFS server. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#nfs"
     },
     "path": {
      "type": "string",
      "description": "Path that is exported by the NFS server. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#nfs"
     },
     "readOnly": {
      "type": "boolean",
      "description": "ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#nfs"
     }
    }
   },
   "v1.RBDVolumeSource": {
    "id": "v1.RBDVolumeSource",
    "description": "Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.",
    "required": [
     "monitors",
     "image",
     "pool",
     "user",
     "keyring",
     "secretRef"
    ],
    "properties": {
     "monitors": {
      "type": "array",
      "items": {
       "type": "string"
      },
      "description": "A collection of Ceph monitors. More info: http://releases.k8s.io/release-1.2/examples/rbd/README.md#how-to-use-it"
     },
     "image": {
      "type": "string",
      "description": "The rados image name. More info: http://releases.k8s.io/release-1.2/examples/rbd/README.md#how-to-use-it"
     },
     "fsType": {
      "type": "string",
      "description": "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#rbd"
     },
     "pool": {
      "type": "string",
      "description": "The rados pool name. Default is rbd. More info: http://releases.k8s.io/release-1.2/examples/rbd/README.md#how-to-use-it."
     },
     "user": {
      "type": "string",
      "description": "The rados user name. Default is admin. More info: http://releases.k8s.io/release-1.2/examples/rbd/README.md#how-to-use-it"
     },
     "keyring": {
      "type": "string",
      "description": "Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: http://releases.k8s.io/release-1.2/examples/rbd/README.md#how-to-use-it"
     },
     "secretRef": {
      "$ref": "v1.LocalObjectReference",
      "description": "SecretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is empty. More info: http://releases.k8s.io/release-1.2/examples/rbd/README.md#how-to-use-it"
     },
     "readOnly": {
      "type": "boolean",
      "description": "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: http://releases.k8s.io/release-1.2/examples/rbd/README.md#how-to-use-it"
     }
    }
   },
   "v1.LocalObjectReference": {
    "id": "v1.LocalObjectReference",
    "description": "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
    "properties": {
     "name": {
      "type": "string",
      "description": "Name of the referent. More info: http://releases.k8s.io/release-1.2/docs/user-guide/identifiers.md#names"
     }
    }
   },
   "v1.ISCSIVolumeSource": {
    "id": "v1.ISCSIVolumeSource",
    "description": "Represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.",
    "required": [
     "targetPortal",
     "iqn",
     "lun"
    ],
    "properties": {
     "targetPortal": {
      "type": "string",
      "description": "iSCSI target portal. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260)."
     },
     "iqn": {
      "type": "string",
      "description": "Target iSCSI Qualified Name."
     },
     "lun": {
      "type": "integer",
      "format": "int32",
      "description": "iSCSI target lun number."
     },
     "iscsiInterface": {
      "type": "string",
      "description": "Optional: Defaults to 'default' (tcp). iSCSI interface name that uses an iSCSI transport."
     },
     "fsType": {
      "type": "string",
      "description": "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#iscsi"
     },
     "readOnly": {
      "type": "boolean",
      "description": "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false."
     }
    }
   },
   "v1.CinderVolumeSource": {
    "id": "v1.CinderVolumeSource",
    "description": "Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.",
    "required": [
     "volumeID"
    ],
    "properties": {
     "volumeID": {
      "type": "string",
      "description": "volume id used to identify the volume in cinder More info: http://releases.k8s.io/release-1.2/examples/mysql-cinder-pd/README.md"
     },
     "fsType": {
      "type": "string",
      "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: http://releases.k8s.io/release-1.2/examples/mysql-cinder-pd/README.md"
     },
     "readOnly": {
      "type": "boolean",
      "description": "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: http://releases.k8s.io/release-1.2/examples/mysql-cinder-pd/README.md"
     }
    }
   },
   "v1.CephFSVolumeSource": {
    "id": "v1.CephFSVolumeSource",
    "description": "Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.",
    "required": [
     "monitors"
    ],
    "properties": {
     "monitors": {
      "type": "array",
      "items": {
       "type": "string"
      },
      "description": "Required: Monitors is a collection of Ceph monitors More info: http://releases.k8s.io/release-1.2/examples/cephfs/README.md#how-to-use-it"
     },
     "path": {
      "type": "string",
      "description": "Optional: Used as the mounted root, rather than the full Ceph tree, default is /"
     },
     "user": {
      "type": "string",
      "description": "Optional: User is the rados user name, default is admin More info: http://releases.k8s.io/release-1.2/examples/cephfs/README.md#how-to-use-it"
     },
     "secretFile": {
      "type": "string",
      "description": "Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: http://releases.k8s.io/release-1.2/examples/cephfs/README.md#how-to-use-it"
     },
     "secretRef": {
      "$ref": "v1.LocalObjectReference",
      "description": "Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: http://releases.k8s.io/release-1.2/examples/cephfs/README.md#how-to-use-it"
     },
     "readOnly": {
      "type": "boolean",
      "description": "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: http://releases.k8s.io/release-1.2/examples/cephfs/README.md#how-to-use-it"
     }
    }
   },
   "v1.FCVolumeSource": {
    "id": "v1.FCVolumeSource",
    "description": "Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.",
    "required": [
     "targetWWNs",
     "lun"
    ],
    "properties": {
     "targetWWNs": {
      "type": "array",
      "items": {
       "type": "string"
      },
      "description": "Required: FC target world wide names (WWNs)"
     },
     "lun": {
      "type": "integer",
      "format": "int32",
      "description": "Required: FC target lun number"
     },
     "fsType": {
      "type": "string",
      "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified."
     },
     "readOnly": {
      "type": "boolean",
      "description": "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts."
     }
    }
   },
   "v1.FlockerVolumeSource": {
    "id": "v1.FlockerVolumeSource",
    "description": "Represents a Flocker volume mounted by the Flocker agent. Flocker volumes do not support ownership management or SELinux relabeling.",
    "required": [
     "datasetName"
    ],
    "properties": {
     "datasetName": {
      "type": "string",
      "description": "Required: the volume name. This is going to be store on metadata -\u003e name on the payload for Flocker"
     }
    }
   },
   "v1.FlexVolumeSource": {
    "id": "v1.FlexVolumeSource",
    "description": "FlexVolume represents a generic volume resource that is provisioned/attached using a exec based plugin. This is an alpha feature and may change in future.",
    "required": [
     "driver"
    ],
    "properties": {
     "driver": {
      "type": "string",
      "description": "Driver is the name of the driver to use for this volume."
     },
     "fsType": {
      "type": "string",
      "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". The default filesystem depends on FlexVolume script."
     },
     "secretRef": {
      "$ref": "v1.LocalObjectReference",
      "description": "Optional: SecretRef is reference to the authentication secret for User, default is empty."
     },
     "readOnly": {
      "type": "boolean",
      "description": "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts."
     },
     "options": {
      "type": "any",
      "description": "Optional: Extra command options if any."
     }
    }
   },
   "v1.AzureFileVolumeSource": {
    "id": "v1.AzureFileVolumeSource",
    "description": "AzureFile represents an Azure File Service mount on the host and bind mount to the pod.",
    "required": [
     "secretName",
     "shareName"
    ],
    "properties": {
     "secretName": {
      "type": "string",
      "description": "the name of secret that contains Azure Storage Account Name and Key"
     },
     "shareName": {
      "type": "string",
      "description": "Share Name"
     },
     "readOnly": {
      "type": "boolean",
      "description": "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts."
     }
    }
   },
   "v1.PersistentVolumeStatus": {
    "id": "v1.PersistentVolumeStatus",
    "description": "PersistentVolumeStatus is the current status of a persistent volume.",
    "properties": {
     "phase": {
      "type": "string",
      "description": "Phase indicates if a volume is available, bound to a claim, or released by a claim. More info: http://releases.k8s.io/release-1.2/docs/user-guide/persistent-volumes.md#phase"
     },
     "message": {
      "type": "string",
      "description": "A human-readable message indicating details about why the volume is in this state."
     },
     "reason": {
      "type": "string",
      "description": "Reason is a brief CamelCase string that describes any failure and is meant for machine parsing and tidy display in the CLI."
     }
    }
   },
   "v1.PodList": {
    "id": "v1.PodList",
    "description": "PodList is a list of Pods.",
    "required": [
     "items"
    ],
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "unversioned.ListMeta",
      "description": "Standard list metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "items": {
      "type": "array",
      "items": {
       "$ref": "v1.Pod"
      },
      "description": "List of pods. More info: http://releases.k8s.io/release-1.2/docs/user-guide/pods.md"
     }
    }
   },
   "v1.Pod": {
    "id": "v1.Pod",
    "description": "Pod is a collection of containers that can run on a host. This resource is created by clients and scheduled onto hosts.",
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "v1.ObjectMeta",
      "description": "Standard object's metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "spec": {
      "$ref": "v1.PodSpec",
      "description": "Specification of the desired behavior of the pod. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#spec-and-status"
     },
     "status": {
      "$ref": "v1.PodStatus",
      "description": "Most recently observed status of the pod. This data may not be up to date. Populated by the system. Read-only. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#spec-and-status"
     }
    }
   },
   "v1.PodSpec": {
    "id": "v1.PodSpec",
    "description": "PodSpec is a description of a pod.",
    "required": [
     "containers"
    ],
    "properties": {
     "volumes": {
      "type": "array",
      "items": {
       "$ref": "v1.Volume"
      },
      "description": "List of volumes that can be mounted by containers belonging to the pod. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md"
     },
     "containers": {
      "type": "array",
      "items": {
       "$ref": "v1.Container"
      },
      "description": "List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated. More info: http://releases.k8s.io/release-1.2/docs/user-guide/containers.md"
     },
     "restartPolicy": {
      "type": "string",
      "description": "Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to Always. More info: http://releases.k8s.io/release-1.2/docs/user-guide/pod-states.md#restartpolicy"
     },
     "terminationGracePeriodSeconds": {
      "type": "integer",
      "format": "int64",
      "description": "Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds."
     },
     "activeDeadlineSeconds": {
      "type": "integer",
      "format": "int64",
      "description": "Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer."
     },
     "dnsPolicy": {
      "type": "string",
      "description": "Set DNS policy for containers within the pod. One of 'ClusterFirst' or 'Default'. Defaults to \"ClusterFirst\"."
     },
     "nodeSelector": {
      "type": "any",
      "description": "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: http://releases.k8s.io/release-1.2/docs/user-guide/node-selection/README.md"
     },
     "serviceAccountName": {
      "type": "string",
      "description": "ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: http://releases.k8s.io/release-1.2/docs/design/service_accounts.md"
     },
     "serviceAccount": {
      "type": "string",
      "description": "DeprecatedServiceAccount is a depreciated alias for ServiceAccountName. Deprecated: Use serviceAccountName instead."
     },
     "nodeName": {
      "type": "string",
      "description": "NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements."
     },
     "hostNetwork": {
      "type": "boolean",
      "description": "Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified. Default to false."
     },
     "hostPID": {
      "type": "boolean",
      "description": "Use the host's pid namespace. Optional: Default to false."
     },
     "hostIPC": {
      "type": "boolean",
      "description": "Use the host's ipc namespace. Optional: Default to false."
     },
     "securityContext": {
      "$ref": "v1.PodSecurityContext",
      "description": "SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty.  See type description for default values of each field."
     },
     "imagePullSecrets": {
      "type": "array",
      "items": {
       "$ref": "v1.LocalObjectReference"
      },
      "description": "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. More info: http://releases.k8s.io/release-1.2/docs/user-guide/images.md#specifying-imagepullsecrets-on-a-pod"
     }
    }
   },
   "v1.Volume": {
    "id": "v1.Volume",
    "description": "Volume represents a named volume in a pod that may be accessed by any container in the pod.",
    "required": [
     "name"
    ],
    "properties": {
     "name": {
      "type": "string",
      "description": "Volume's name. Must be a DNS_LABEL and unique within the pod. More info: http://releases.k8s.io/release-1.2/docs/user-guide/identifiers.md#names"
     },
     "hostPath": {
      "$ref": "v1.HostPathVolumeSource",
      "description": "HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#hostpath"
     },
     "emptyDir": {
      "$ref": "v1.EmptyDirVolumeSource",
      "description": "EmptyDir represents a temporary directory that shares a pod's lifetime. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#emptydir"
     },
     "gcePersistentDisk": {
      "$ref": "v1.GCEPersistentDiskVolumeSource",
      "description": "GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#gcepersistentdisk"
     },
     "awsElasticBlockStore": {
      "$ref": "v1.AWSElasticBlockStoreVolumeSource",
      "description": "AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#awselasticblockstore"
     },
     "gitRepo": {
      "$ref": "v1.GitRepoVolumeSource",
      "description": "GitRepo represents a git repository at a particular revision."
     },
     "secret": {
      "$ref": "v1.SecretVolumeSource",
      "description": "Secret represents a secret that should populate this volume. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#secrets"
     },
     "nfs": {
      "$ref": "v1.NFSVolumeSource",
      "description": "NFS represents an NFS mount on the host that shares a pod's lifetime More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#nfs"
     },
     "iscsi": {
      "$ref": "v1.ISCSIVolumeSource",
      "description": "ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: http://releases.k8s.io/release-1.2/examples/iscsi/README.md"
     },
     "glusterfs": {
      "$ref": "v1.GlusterfsVolumeSource",
      "description": "Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: http://releases.k8s.io/release-1.2/examples/glusterfs/README.md"
     },
     "persistentVolumeClaim": {
      "$ref": "v1.PersistentVolumeClaimVolumeSource",
      "description": "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: http://releases.k8s.io/release-1.2/docs/user-guide/persistent-volumes.md#persistentvolumeclaims"
     },
     "rbd": {
      "$ref": "v1.RBDVolumeSource",
      "description": "RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: http://releases.k8s.io/release-1.2/examples/rbd/README.md"
     },
     "flexVolume": {
      "$ref": "v1.FlexVolumeSource",
      "description": "FlexVolume represents a generic volume resource that is provisioned/attached using a exec based plugin. This is an alpha feature and may change in future."
     },
     "cinder": {
      "$ref": "v1.CinderVolumeSource",
      "description": "Cinder represents a cinder volume attached and mounted on kubelets host machine More info: http://releases.k8s.io/release-1.2/examples/mysql-cinder-pd/README.md"
     },
     "cephfs": {
      "$ref": "v1.CephFSVolumeSource",
      "description": "CephFS represents a Ceph FS mount on the host that shares a pod's lifetime"
     },
     "flocker": {
      "$ref": "v1.FlockerVolumeSource",
      "description": "Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running"
     },
     "downwardAPI": {
      "$ref": "v1.DownwardAPIVolumeSource",
      "description": "DownwardAPI represents downward API about the pod that should populate this volume"
     },
     "fc": {
      "$ref": "v1.FCVolumeSource",
      "description": "FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod."
     },
     "azureFile": {
      "$ref": "v1.AzureFileVolumeSource",
      "description": "AzureFile represents an Azure File Service mount on the host and bind mount to the pod."
     },
     "configMap": {
      "$ref": "v1.ConfigMapVolumeSource",
      "description": "ConfigMap represents a configMap that should populate this volume"
     }
    }
   },
   "v1.EmptyDirVolumeSource": {
    "id": "v1.EmptyDirVolumeSource",
    "description": "Represents an empty directory for a pod. Empty directory volumes support ownership management and SELinux relabeling.",
    "properties": {
     "medium": {
      "type": "string",
      "description": "What type of storage medium should back this directory. The default is \"\" which means to use the node's default medium. Must be an empty string (default) or Memory. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#emptydir"
     }
    }
   },
   "v1.GitRepoVolumeSource": {
    "id": "v1.GitRepoVolumeSource",
    "description": "Represents a volume that is populated with the contents of a git repository. Git repo volumes do not support ownership management. Git repo volumes support SELinux relabeling.",
    "required": [
     "repository"
    ],
    "properties": {
     "repository": {
      "type": "string",
      "description": "Repository URL"
     },
     "revision": {
      "type": "string",
      "description": "Commit hash for the specified revision."
     },
     "directory": {
      "type": "string",
      "description": "Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name."
     }
    }
   },
   "v1.SecretVolumeSource": {
    "id": "v1.SecretVolumeSource",
    "description": "Adapts a Secret into a volume.\n\nThe contents of the target Secret's Data field will be presented in a volume as files using the keys in the Data field as the file names. Secret volumes support ownership management and SELinux relabeling.",
    "properties": {
     "secretName": {
      "type": "string",
      "description": "Name of the secret in the pod's namespace to use. More info: http://releases.k8s.io/release-1.2/docs/user-guide/volumes.md#secrets"
     }
    }
   },
   "v1.PersistentVolumeClaimVolumeSource": {
    "id": "v1.PersistentVolumeClaimVolumeSource",
    "description": "PersistentVolumeClaimVolumeSource references the user's PVC in the same namespace. This volume finds the bound PV and mounts that volume for the pod. A PersistentVolumeClaimVolumeSource is, essentially, a wrapper around another type of volume that is owned by someone else (the system).",
    "required": [
     "claimName"
    ],
    "properties": {
     "claimName": {
      "type": "string",
      "description": "ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: http://releases.k8s.io/release-1.2/docs/user-guide/persistent-volumes.md#persistentvolumeclaims"
     },
     "readOnly": {
      "type": "boolean",
      "description": "Will force the ReadOnly setting in VolumeMounts. Default false."
     }
    }
   },
   "v1.DownwardAPIVolumeSource": {
    "id": "v1.DownwardAPIVolumeSource",
    "description": "DownwardAPIVolumeSource represents a volume containing downward API info. Downward API volumes support ownership management and SELinux relabeling.",
    "properties": {
     "items": {
      "type": "array",
      "items": {
       "$ref": "v1.DownwardAPIVolumeFile"
      },
      "description": "Items is a list of downward API volume file"
     }
    }
   },
   "v1.DownwardAPIVolumeFile": {
    "id": "v1.DownwardAPIVolumeFile",
    "description": "DownwardAPIVolumeFile represents information to create the file containing the pod field",
    "required": [
     "path",
     "fieldRef"
    ],
    "properties": {
     "path": {
      "type": "string",
      "description": "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'"
     },
     "fieldRef": {
      "$ref": "v1.ObjectFieldSelector",
      "description": "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported."
     }
    }
   },
   "v1.ObjectFieldSelector": {
    "id": "v1.ObjectFieldSelector",
    "description": "ObjectFieldSelector selects an APIVersioned field of an object.",
    "required": [
     "fieldPath"
    ],
    "properties": {
     "apiVersion": {
      "type": "string",
      "description": "Version of the schema the FieldPath is written in terms of, defaults to \"v1\"."
     },
     "fieldPath": {
      "type": "string",
      "description": "Path of the field to select in the specified API version."
     }
    }
   },
   "v1.ConfigMapVolumeSource": {
    "id": "v1.ConfigMapVolumeSource",
    "description": "Adapts a ConfigMap into a volume.\n\nThe contents of the target ConfigMap's Data field will be presented in a volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. ConfigMap volumes support ownership management and SELinux relabeling.",
    "properties": {
     "name": {
      "type": "string",
      "description": "Name of the referent. More info: http://releases.k8s.io/release-1.2/docs/user-guide/identifiers.md#names"
     },
     "items": {
      "type": "array",
      "items": {
       "$ref": "v1.KeyToPath"
      },
      "description": "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error. Paths must be relative and may not contain the '..' path or start with '..'."
     }
    }
   },
   "v1.KeyToPath": {
    "id": "v1.KeyToPath",
    "description": "Maps a string key to a path within a volume.",
    "required": [
     "key",
     "path"
    ],
    "properties": {
     "key": {
      "type": "string",
      "description": "The key to project."
     },
     "path": {
      "type": "string",
      "description": "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'."
     }
    }
   },
   "v1.Container": {
    "id": "v1.Container",
    "description": "A single application container that you want to run within a pod.",
    "required": [
     "name"
    ],
    "properties": {
     "name": {
      "type": "string",
      "description": "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated."
     },
     "image": {
      "type": "string",
      "description": "Docker image name. More info: http://releases.k8s.io/release-1.2/docs/user-guide/images.md"
     },
     "command": {
      "type": "array",
      "items": {
       "type": "string"
      },
      "description": "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: http://releases.k8s.io/release-1.2/docs/user-guide/containers.md#containers-and-commands"
     },
     "args": {
      "type": "array",
      "items": {
       "type": "string"
      },
      "description": "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: http://releases.k8s.io/release-1.2/docs/user-guide/containers.md#containers-and-commands"
     },
     "workingDir": {
      "type": "string",
      "description": "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated."
     },
     "ports": {
      "type": "array",
      "items": {
       "$ref": "v1.ContainerPort"
      },
      "description": "List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default \"0.0.0.0\" address inside a container will be accessible from the network. Cannot be updated."
     },
     "env": {
      "type": "array",
      "items": {
       "$ref": "v1.EnvVar"
      },
      "description": "List of environment variables to set in the container. Cannot be updated."
     },
     "resources": {
      "$ref": "v1.ResourceRequirements",
      "description": "Compute Resources required by this container. Cannot be updated. More info: http://releases.k8s.io/release-1.2/docs/user-guide/persistent-volumes.md#resources"
     },
     "volumeMounts": {
      "type": "array",
      "items": {
       "$ref": "v1.VolumeMount"
      },
      "description": "Pod volumes to mount into the container's filesyste. Cannot be updated."
     },
     "livenessProbe": {
      "$ref": "v1.Probe",
      "description": "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: http://releases.k8s.io/release-1.2/docs/user-guide/pod-states.md#container-probes"
     },
     "readinessProbe": {
      "$ref": "v1.Probe",
      "description": "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: http://releases.k8s.io/release-1.2/docs/user-guide/pod-states.md#container-probes"
     },
     "lifecycle": {
      "$ref": "v1.Lifecycle",
      "description": "Actions that the management system should take in response to container lifecycle events. Cannot be updated."
     },
     "terminationMessagePath": {
      "type": "string",
      "description": "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Defaults to /dev/termination-log. Cannot be updated."
     },
     "imagePullPolicy": {
      "type": "string",
      "description": "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: http://releases.k8s.io/release-1.2/docs/user-guide/images.md#updating-images"
     },
     "securityContext": {
      "$ref": "v1.SecurityContext",
      "description": "Security options the pod should run with. More info: http://releases.k8s.io/release-1.2/docs/design/security_context.md"
     },
     "stdin": {
      "type": "boolean",
      "description": "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false."
     },
     "stdinOnce": {
      "type": "boolean",
      "description": "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false"
     },
     "tty": {
      "type": "boolean",
      "description": "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false."
     }
    }
   },
   "v1.ContainerPort": {
    "id": "v1.ContainerPort",
    "description": "ContainerPort represents a network port in a single container.",
    "required": [
     "containerPort"
    ],
    "properties": {
     "name": {
      "type": "string",
      "description": "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services."
     },
     "hostPort": {
      "type": "integer",
      "format": "int32",
      "description": "Number of port to expose on the host. If specified, this must be a valid port number, 0 \u003c x \u003c 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this."
     },
     "containerPort": {
      "type": "integer",
      "format": "int32",
      "description": "Number of port to expose on the pod's IP address. This must be a valid port number, 0 \u003c x \u003c 65536."
     },
     "protocol": {
      "type": "string",
      "description": "Protocol for port. Must be UDP or TCP. Defaults to \"TCP\"."
     },
     "hostIP": {
      "type": "string",
      "description": "What host IP to bind the external port to."
     }
    }
   },
   "v1.EnvVar": {
    "id": "v1.EnvVar",
    "description": "EnvVar represents an environment variable present in a Container.",
    "required": [
     "name"
    ],
    "properties": {
     "name": {
      "type": "string",
      "description": "Name of the environment variable. Must be a C_IDENTIFIER."
     },
     "value": {
      "type": "string",
      "description": "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to \"\"."
     },
     "valueFrom": {
      "$ref": "v1.EnvVarSource",
      "description": "Source for the environment variable's value. Cannot be used if value is not empty."
     }
    }
   },
   "v1.EnvVarSource": {
    "id": "v1.EnvVarSource",
    "description": "EnvVarSource represents a source for the value of an EnvVar.",
    "properties": {
     "fieldRef": {
      "$ref": "v1.ObjectFieldSelector",
      "description": "Selects a field of the pod; only name and namespace are supported."
     },
     "configMapKeyRef": {
      "$ref": "v1.ConfigMapKeySelector",
      "description": "Selects a key of a ConfigMap."
     },
     "secretKeyRef": {
      "$ref": "v1.SecretKeySelector",
      "description": "Selects a key of a secret in the pod's namespace"
     }
    }
   },
   "v1.ConfigMapKeySelector": {
    "id": "v1.ConfigMapKeySelector",
    "description": "Selects a key from a ConfigMap.",
    "required": [
     "key"
    ],
    "properties": {
     "name": {
      "type": "string",
      "description": "Name of the referent. More info: http://releases.k8s.io/release-1.2/docs/user-guide/identifiers.md#names"
     },
     "key": {
      "type": "string",
      "description": "The key to select."
     }
    }
   },
   "v1.SecretKeySelector": {
    "id": "v1.SecretKeySelector",
    "description": "SecretKeySelector selects a key of a Secret.",
    "required": [
     "key"
    ],
    "properties": {
     "name": {
      "type": "string",
      "description": "Name of the referent. More info: http://releases.k8s.io/release-1.2/docs/user-guide/identifiers.md#names"
     },
     "key": {
      "type": "string",
      "description": "The key of the secret to select from.  Must be a valid secret key."
     }
    }
   },
   "v1.VolumeMount": {
    "id": "v1.VolumeMount",
    "description": "VolumeMount describes a mounting of a Volume within a container.",
    "required": [
     "name",
     "mountPath"
    ],
    "properties": {
     "name": {
      "type": "string",
      "description": "This must match the Name of a Volume."
     },
     "readOnly": {
      "type": "boolean",
      "description": "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false."
     },
     "mountPath": {
      "type": "string",
      "description": "Path within the container at which the volume should be mounted.  Must not contain ':'."
     }
    }
   },
   "v1.Probe": {
    "id": "v1.Probe",
    "description": "Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.",
    "properties": {
     "exec": {
      "$ref": "v1.ExecAction",
      "description": "One and only one of the following should be specified. Exec specifies the action to take."
     },
     "httpGet": {
      "$ref": "v1.HTTPGetAction",
      "description": "HTTPGet specifies the http request to perform."
     },
     "tcpSocket": {
      "$ref": "v1.TCPSocketAction",
      "description": "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported"
     },
     "initialDelaySeconds": {
      "type": "integer",
      "format": "int32",
      "description": "Number of seconds after the container has started before liveness probes are initiated. More info: http://releases.k8s.io/release-1.2/docs/user-guide/pod-states.md#container-probes"
     },
     "timeoutSeconds": {
      "type": "integer",
      "format": "int32",
      "description": "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: http://releases.k8s.io/release-1.2/docs/user-guide/pod-states.md#container-probes"
     },
     "periodSeconds": {
      "type": "integer",
      "format": "int32",
      "description": "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1."
     },
     "successThreshold": {
      "type": "integer",
      "format": "int32",
      "description": "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1."
     },
     "failureThreshold": {
      "type": "integer",
      "format": "int32",
      "description": "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1."
     }
    }
   },
   "v1.ExecAction": {
    "id": "v1.ExecAction",
    "description": "ExecAction describes a \"run in container\" action.",
    "properties": {
     "command": {
      "type": "array",
      "items": {
       "type": "string"
      },
      "description": "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy."
     }
    }
   },
   "v1.HTTPGetAction": {
    "id": "v1.HTTPGetAction",
    "description": "HTTPGetAction describes an action based on HTTP Get requests.",
    "required": [
     "port"
    ],
    "properties": {
     "path": {
      "type": "string",
      "description": "Path to access on the HTTP server."
     },
     "port": {
      "type": "string",
      "description": "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME."
     },
     "host": {
      "type": "string",
      "description": "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead."
     },
     "scheme": {
      "type": "string",
      "description": "Scheme to use for connecting to the host. Defaults to HTTP."
     },
     "httpHeaders": {
      "type": "array",
      "items": {
       "$ref": "v1.HTTPHeader"
      },
      "description": "Custom headers to set in the request. HTTP allows repeated headers."
     }
    }
   },
   "v1.HTTPHeader": {
    "id": "v1.HTTPHeader",
    "description": "HTTPHeader describes a custom header to be used in HTTP probes",
    "required": [
     "name",
     "value"
    ],
    "properties": {
     "name": {
      "type": "string",
      "description": "The header field name"
     },
     "value": {
      "type": "string",
      "description": "The header field value"
     }
    }
   },
   "v1.TCPSocketAction": {
    "id": "v1.TCPSocketAction",
    "description": "TCPSocketAction describes an action based on opening a socket",
    "required": [
     "port"
    ],
    "properties": {
     "port": {
      "type": "string",
      "description": "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME."
     }
    }
   },
   "v1.Lifecycle": {
    "id": "v1.Lifecycle",
    "description": "Lifecycle describes actions that the management system should take in response to container lifecycle events. For the PostStart and PreStop lifecycle handlers, management of the container blocks until the action is complete, unless the container process fails, in which case the handler is aborted.",
    "properties": {
     "postStart": {
      "$ref": "v1.Handler",
      "description": "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: http://releases.k8s.io/release-1.2/docs/user-guide/container-environment.md#hook-details"
     },
     "preStop": {
      "$ref": "v1.Handler",
      "description": "PreStop is called immediately before a container is terminated. The container is terminated after the handler completes. The reason for termination is passed to the handler. Regardless of the outcome of the handler, the container is eventually terminated. Other management of the container blocks until the hook completes. More info: http://releases.k8s.io/release-1.2/docs/user-guide/container-environment.md#hook-details"
     }
    }
   },
   "v1.Handler": {
    "id": "v1.Handler",
    "description": "Handler defines a specific action that should be taken",
    "properties": {
     "exec": {
      "$ref": "v1.ExecAction",
      "description": "One and only one of the following should be specified. Exec specifies the action to take."
     },
     "httpGet": {
      "$ref": "v1.HTTPGetAction",
      "description": "HTTPGet specifies the http request to perform."
     },
     "tcpSocket": {
      "$ref": "v1.TCPSocketAction",
      "description": "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported"
     }
    }
   },
   "v1.SecurityContext": {
    "id": "v1.SecurityContext",
    "description": "SecurityContext holds security configuration that will be applied to a container. Some fields are present in both SecurityContext and PodSecurityContext.  When both are set, the values in SecurityContext take precedence.",
    "properties": {
     "capabilities": {
      "$ref": "v1.Capabilities",
      "description": "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime."
     },
     "privileged": {
      "type": "boolean",
      "description": "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false."
     },
     "seLinuxOptions": {
      "$ref": "v1.SELinuxOptions",
      "description": "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence."
     },
     "runAsUser": {
      "type": "integer",
      "format": "int64",
      "description": "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence."
     },
     "runAsNonRoot": {
      "type": "boolean",
      "description": "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence."
     },
     "readOnlyRootFilesystem": {
      "type": "boolean",
      "description": "Whether this container has a read-only root filesystem. Default is false."
     }
    }
   },
   "v1.Capabilities": {
    "id": "v1.Capabilities",
    "description": "Adds and removes POSIX capabilities from running containers.",
    "properties": {
     "add": {
      "type": "array",
      "items": {
       "$ref": "v1.Capability"
      },
      "description": "Added capabilities"
     },
     "drop": {
      "type": "array",
      "items": {
       "$ref": "v1.Capability"
      },
      "description": "Removed capabilities"
     }
    }
   },
   "v1.Capability": {
    "id": "v1.Capability",
    "properties": {}
   },
   "v1.SELinuxOptions": {
    "id": "v1.SELinuxOptions",
    "description": "SELinuxOptions are the labels to be applied to the container",
    "properties": {
     "user": {
      "type": "string",
      "description": "User is a SELinux user label that applies to the container."
     },
     "role": {
      "type": "string",
      "description": "Role is a SELinux role label that applies to the container."
     },
     "type": {
      "type": "string",
      "description": "Type is a SELinux type label that applies to the container."
     },
     "level": {
      "type": "string",
      "description": "Level is SELinux level label that applies to the container."
     }
    }
   },
   "v1.PodSecurityContext": {
    "id": "v1.PodSecurityContext",
    "description": "PodSecurityContext holds pod-level security attributes and common container settings. Some fields are also present in container.securityContext.  Field values of container.securityContext take precedence over field values of PodSecurityContext.",
    "properties": {
     "seLinuxOptions": {
      "$ref": "v1.SELinuxOptions",
      "description": "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container."
     },
     "runAsUser": {
      "type": "integer",
      "format": "int64",
      "description": "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container."
     },
     "runAsNonRoot": {
      "type": "boolean",
      "description": "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence."
     },
     "supplementalGroups": {
      "type": "array",
      "items": {
       "$ref": "integer"
      },
      "description": "A list of groups applied to the first process run in each container, in addition to the container's primary GID.  If unspecified, no groups will be added to any container."
     },
     "fsGroup": {
      "type": "integer",
      "format": "int64",
      "description": "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:\n\n1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw "
     }
    }
   },
   "integer": {
    "id": "integer",
    "properties": {}
   },
   "v1.PodStatus": {
    "id": "v1.PodStatus",
    "description": "PodStatus represents information about the status of a pod. Status may trail the actual state of a system.",
    "properties": {
     "phase": {
      "type": "string",
      "description": "Current condition of the pod. More info: http://releases.k8s.io/release-1.2/docs/user-guide/pod-states.md#pod-phase"
     },
     "conditions": {
      "type": "array",
      "items": {
       "$ref": "v1.PodCondition"
      },
      "description": "Current service state of pod. More info: http://releases.k8s.io/release-1.2/docs/user-guide/pod-states.md#pod-conditions"
     },
     "message": {
      "type": "string",
      "description": "A human readable message indicating details about why the pod is in this condition."
     },
     "reason": {
      "type": "string",
      "description": "A brief CamelCase message indicating details about why the pod is in this state. e.g. 'OutOfDisk'"
     },
     "hostIP": {
      "type": "string",
      "description": "IP address of the host to which the pod is assigned. Empty if not yet scheduled."
     },
     "podIP": {
      "type": "string",
      "description": "IP address allocated to the pod. Routable at least within the cluster. Empty if not yet allocated."
     },
     "startTime": {
      "type": "string",
      "description": "RFC 3339 date and time at which the object was acknowledged by the Kubelet. This is before the Kubelet pulled the container image(s) for the pod."
     },
     "containerStatuses": {
      "type": "array",
      "items": {
       "$ref": "v1.ContainerStatus"
      },
      "description": "The list has one entry per container in the manifest. Each entry is currently the output of docker inspect. More info: http://releases.k8s.io/release-1.2/docs/user-guide/pod-states.md#container-statuses"
     }
    }
   },
   "v1.PodCondition": {
    "id": "v1.PodCondition",
    "description": "PodCondition contains details for the current condition of this pod.",
    "required": [
     "type",
     "status"
    ],
    "properties": {
     "type": {
      "type": "string",
      "description": "Type is the type of the condition. Currently only Ready. More info: http://releases.k8s.io/release-1.2/docs/user-guide/pod-states.md#pod-conditions"
     },
     "status": {
      "type": "string",
      "description": "Status is the status of the condition. Can be True, False, Unknown. More info: http://releases.k8s.io/release-1.2/docs/user-guide/pod-states.md#pod-conditions"
     },
     "lastProbeTime": {
      "type": "string",
      "description": "Last time we probed the condition."
     },
     "lastTransitionTime": {
      "type": "string",
      "description": "Last time the condition transitioned from one status to another."
     },
     "reason": {
      "type": "string",
      "description": "Unique, one-word, CamelCase reason for the condition's last transition."
     },
     "message": {
      "type": "string",
      "description": "Human-readable message indicating details about last transition."
     }
    }
   },
   "v1.ContainerStatus": {
    "id": "v1.ContainerStatus",
    "description": "ContainerStatus contains details for the current status of this container.",
    "required": [
     "name",
     "ready",
     "restartCount",
     "image",
     "imageID"
    ],
    "properties": {
     "name": {
      "type": "string",
      "description": "This must be a DNS_LABEL. Each container in a pod must have a unique name. Cannot be updated."
     },
     "state": {
      "$ref": "v1.ContainerState",
      "description": "Details about the container's current condition."
     },
     "lastState": {
      "$ref": "v1.ContainerState",
      "description": "Details about the container's last termination condition."
     },
     "ready": {
      "type": "boolean",
      "description": "Specifies whether the container has passed its readiness probe."
     },
     "restartCount": {
      "type": "integer",
      "format": "int32",
      "description": "The number of times the container has been restarted, currently based on the number of dead containers that have not yet been removed. Note that this is calculated from dead containers. But those containers are subject to garbage collection. This value will get capped at 5 by GC."
     },
     "image": {
      "type": "string",
      "description": "The image the container is running. More info: http://releases.k8s.io/release-1.2/docs/user-guide/images.md"
     },
     "imageID": {
      "type": "string",
      "description": "ImageID of the container's image."
     },
     "containerID": {
      "type": "string",
      "description": "Container's ID in the format 'docker://\u003ccontainer_id\u003e'. More info: http://releases.k8s.io/release-1.2/docs/user-guide/container-environment.md#container-information"
     }
    }
   },
   "v1.ContainerState": {
    "id": "v1.ContainerState",
    "description": "ContainerState holds a possible state of container. Only one of its members may be specified. If none of them is specified, the default one is ContainerStateWaiting.",
    "properties": {
     "waiting": {
      "$ref": "v1.ContainerStateWaiting",
      "description": "Details about a waiting container"
     },
     "running": {
      "$ref": "v1.ContainerStateRunning",
      "description": "Details about a running container"
     },
     "terminated": {
      "$ref": "v1.ContainerStateTerminated",
      "description": "Details about a terminated container"
     }
    }
   },
   "v1.ContainerStateWaiting": {
    "id": "v1.ContainerStateWaiting",
    "description": "ContainerStateWaiting is a waiting state of a container.",
    "properties": {
     "reason": {
      "type": "string",
      "description": "(brief) reason the container is not yet running."
     },
     "message": {
      "type": "string",
      "description": "Message regarding why the container is not yet running."
     }
    }
   },
   "v1.ContainerStateRunning": {
    "id": "v1.ContainerStateRunning",
    "description": "ContainerStateRunning is a running state of a container.",
    "properties": {
     "startedAt": {
      "type": "string",
      "description": "Time at which the container was last (re-)started"
     }
    }
   },
   "v1.ContainerStateTerminated": {
    "id": "v1.ContainerStateTerminated",
    "description": "ContainerStateTerminated is a terminated state of a container.",
    "required": [
     "exitCode"
    ],
    "properties": {
     "exitCode": {
      "type": "integer",
      "format": "int32",
      "description": "Exit status from the last termination of the container"
     },
     "signal": {
      "type": "integer",
      "format": "int32",
      "description": "Signal from the last termination of the container"
     },
     "reason": {
      "type": "string",
      "description": "(brief) reason from the last termination of the container"
     },
     "message": {
      "type": "string",
      "description": "Message regarding the last termination of the container"
     },
     "startedAt": {
      "type": "string",
      "description": "Time at which previous execution of the container started"
     },
     "finishedAt": {
      "type": "string",
      "description": "Time at which the container last terminated"
     },
     "containerID": {
      "type": "string",
      "description": "Container's ID in the format 'docker://\u003ccontainer_id\u003e'"
     }
    }
   },
   "v1.PodTemplateList": {
    "id": "v1.PodTemplateList",
    "description": "PodTemplateList is a list of PodTemplates.",
    "required": [
     "items"
    ],
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "unversioned.ListMeta",
      "description": "Standard list metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "items": {
      "type": "array",
      "items": {
       "$ref": "v1.PodTemplate"
      },
      "description": "List of pod templates"
     }
    }
   },
   "v1.PodTemplate": {
    "id": "v1.PodTemplate",
    "description": "PodTemplate describes a template for creating copies of a predefined pod.",
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "v1.ObjectMeta",
      "description": "Standard object's metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "template": {
      "$ref": "v1.PodTemplateSpec",
      "description": "Template defines the pods that will be created from this pod template. http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#spec-and-status"
     }
    }
   },
   "v1.PodTemplateSpec": {
    "id": "v1.PodTemplateSpec",
    "description": "PodTemplateSpec describes the data a pod should have when created from a template",
    "properties": {
     "metadata": {
      "$ref": "v1.ObjectMeta",
      "description": "Standard object's metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "spec": {
      "$ref": "v1.PodSpec",
      "description": "Specification of the desired behavior of the pod. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#spec-and-status"
     }
    }
   },
   "v1.ReplicationControllerList": {
    "id": "v1.ReplicationControllerList",
    "description": "ReplicationControllerList is a collection of replication controllers.",
    "required": [
     "items"
    ],
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "unversioned.ListMeta",
      "description": "Standard list metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "items": {
      "type": "array",
      "items": {
       "$ref": "v1.ReplicationController"
      },
      "description": "List of replication controllers. More info: http://releases.k8s.io/release-1.2/docs/user-guide/replication-controller.md"
     }
    }
   },
   "v1.ReplicationController": {
    "id": "v1.ReplicationController",
    "description": "ReplicationController represents the configuration of a replication controller.",
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "v1.ObjectMeta",
      "description": "If the Labels of a ReplicationController are empty, they are defaulted to be the same as the Pod(s) that the replication controller manages. Standard object's metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "spec": {
      "$ref": "v1.ReplicationControllerSpec",
      "description": "Spec defines the specification of the desired behavior of the replication controller. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#spec-and-status"
     },
     "status": {
      "$ref": "v1.ReplicationControllerStatus",
      "description": "Status is the most recently observed status of the replication controller. This data may be out of date by some window of time. Populated by the system. Read-only. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#spec-and-status"
     }
    }
   },
   "v1.ReplicationControllerSpec": {
    "id": "v1.ReplicationControllerSpec",
    "description": "ReplicationControllerSpec is the specification of a replication controller.",
    "properties": {
     "replicas": {
      "type": "integer",
      "format": "int32",
      "description": "Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: http://releases.k8s.io/release-1.2/docs/user-guide/replication-controller.md#what-is-a-replication-controller"
     },
     "selector": {
      "type": "any",
      "description": "Selector is a label query over pods that should match the Replicas count. If Selector is empty, it is defaulted to the labels present on the Pod template. Label keys and values that must match in order to be controlled by this replication controller, if empty defaulted to labels on Pod template. More info: http://releases.k8s.io/release-1.2/docs/user-guide/labels.md#label-selectors"
     },
     "template": {
      "$ref": "v1.PodTemplateSpec",
      "description": "Template is the object that describes the pod that will be created if insufficient replicas are detected. This takes precedence over a TemplateRef. More info: http://releases.k8s.io/release-1.2/docs/user-guide/replication-controller.md#pod-template"
     }
    }
   },
   "v1.ReplicationControllerStatus": {
    "id": "v1.ReplicationControllerStatus",
    "description": "ReplicationControllerStatus represents the current status of a replication controller.",
    "required": [
     "replicas"
    ],
    "properties": {
     "replicas": {
      "type": "integer",
      "format": "int32",
      "description": "Replicas is the most recently oberved number of replicas. More info: http://releases.k8s.io/release-1.2/docs/user-guide/replication-controller.md#what-is-a-replication-controller"
     },
     "fullyLabeledReplicas": {
      "type": "integer",
      "format": "int32",
      "description": "The number of pods that have labels matching the labels of the pod template of the replication controller."
     },
     "observedGeneration": {
      "type": "integer",
      "format": "int64",
      "description": "ObservedGeneration reflects the generation of the most recently observed replication controller."
     }
    }
   },
   "v1.Scale": {
    "id": "v1.Scale",
    "description": "Scale represents a scaling request for a resource.",
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "v1.ObjectMeta",
      "description": "Standard object metadata; More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata."
     },
     "spec": {
      "$ref": "v1.ScaleSpec",
      "description": "defines the behavior of the scale. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#spec-and-status."
     },
     "status": {
      "$ref": "v1.ScaleStatus",
      "description": "current status of the scale. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#spec-and-status. Read-only."
     }
    }
   },
   "v1.ScaleSpec": {
    "id": "v1.ScaleSpec",
    "description": "ScaleSpec describes the attributes of a scale subresource.",
    "properties": {
     "replicas": {
      "type": "integer",
      "format": "int32",
      "description": "desired number of instances for the scaled object."
     }
    }
   },
   "v1.ScaleStatus": {
    "id": "v1.ScaleStatus",
    "description": "ScaleStatus represents the current status of a scale subresource.",
    "required": [
     "replicas"
    ],
    "properties": {
     "replicas": {
      "type": "integer",
      "format": "int32",
      "description": "actual number of observed instances of the scaled object."
     },
     "selector": {
      "type": "string",
      "description": "label query over pods that should match the replicas count. This is same as the label selector but in the string format to avoid introspection by clients. The string will be in the same format as the query-param syntax. More info about label selectors: http://releases.k8s.io/release-1.2/docs/user-guide/labels.md#label-selectors"
     }
    }
   },
   "v1.ResourceQuotaList": {
    "id": "v1.ResourceQuotaList",
    "description": "ResourceQuotaList is a list of ResourceQuota items.",
    "required": [
     "items"
    ],
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "unversioned.ListMeta",
      "description": "Standard list metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "items": {
      "type": "array",
      "items": {
       "$ref": "v1.ResourceQuota"
      },
      "description": "Items is a list of ResourceQuota objects. More info: http://releases.k8s.io/release-1.2/docs/design/admission_control_resource_quota.md#admissioncontrol-plugin-resourcequota"
     }
    }
   },
   "v1.ResourceQuota": {
    "id": "v1.ResourceQuota",
    "description": "ResourceQuota sets aggregate quota restrictions enforced per namespace",
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "v1.ObjectMeta",
      "description": "Standard object's metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "spec": {
      "$ref": "v1.ResourceQuotaSpec",
      "description": "Spec defines the desired quota. http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#spec-and-status"
     },
     "status": {
      "$ref": "v1.ResourceQuotaStatus",
      "description": "Status defines the actual enforced quota and its current usage. http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#spec-and-status"
     }
    }
   },
   "v1.ResourceQuotaSpec": {
    "id": "v1.ResourceQuotaSpec",
    "description": "ResourceQuotaSpec defines the desired hard limits to enforce for Quota.",
    "properties": {
     "hard": {
      "type": "any",
      "description": "Hard is the set of desired hard limits for each named resource. More info: http://releases.k8s.io/release-1.2/docs/design/admission_control_resource_quota.md#admissioncontrol-plugin-resourcequota"
     },
     "scopes": {
      "type": "array",
      "items": {
       "$ref": "v1.ResourceQuotaScope"
      },
      "description": "A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects."
     }
    }
   },
   "v1.ResourceQuotaScope": {
    "id": "v1.ResourceQuotaScope",
    "properties": {}
   },
   "v1.ResourceQuotaStatus": {
    "id": "v1.ResourceQuotaStatus",
    "description": "ResourceQuotaStatus defines the enforced hard limits and observed use.",
    "properties": {
     "hard": {
      "type": "any",
      "description": "Hard is the set of enforced hard limits for each named resource. More info: http://releases.k8s.io/release-1.2/docs/design/admission_control_resource_quota.md#admissioncontrol-plugin-resourcequota"
     },
     "used": {
      "type": "any",
      "description": "Used is the current observed total usage of the resource in the namespace."
     }
    }
   },
   "v1.SecretList": {
    "id": "v1.SecretList",
    "description": "SecretList is a list of Secret.",
    "required": [
     "items"
    ],
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "unversioned.ListMeta",
      "description": "Standard list metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "items": {
      "type": "array",
      "items": {
       "$ref": "v1.Secret"
      },
      "description": "Items is a list of secret objects. More info: http://releases.k8s.io/release-1.2/docs/user-guide/secrets.md"
     }
    }
   },
   "v1.Secret": {
    "id": "v1.Secret",
    "description": "Secret holds secret data of a certain type. The total bytes of the values in the Data field must be less than MaxSecretSize bytes.",
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "v1.ObjectMeta",
      "description": "Standard object's metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "data": {
      "type": "any",
      "description": "Data contains the secret data. Each key must be a valid DNS_SUBDOMAIN or leading dot followed by valid DNS_SUBDOMAIN. The serialized form of the secret data is a base64 encoded string, representing the arbitrary (possibly non-string) data value here. Described in https://tools.ietf.org/html/rfc4648#section-4"
     },
     "type": {
      "type": "string",
      "description": "Used to facilitate programmatic handling of secret data."
     }
    }
   },
   "v1.ServiceAccountList": {
    "id": "v1.ServiceAccountList",
    "description": "ServiceAccountList is a list of ServiceAccount objects",
    "required": [
     "items"
    ],
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "unversioned.ListMeta",
      "description": "Standard list metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "items": {
      "type": "array",
      "items": {
       "$ref": "v1.ServiceAccount"
      },
      "description": "List of ServiceAccounts. More info: http://releases.k8s.io/release-1.2/docs/design/service_accounts.md#service-accounts"
     }
    }
   },
   "v1.ServiceAccount": {
    "id": "v1.ServiceAccount",
    "description": "ServiceAccount binds together: * a name, understood by users, and perhaps by peripheral systems, for an identity * a principal that can be authenticated and authorized * a set of secrets",
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "v1.ObjectMeta",
      "description": "Standard object's metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "secrets": {
      "type": "array",
      "items": {
       "$ref": "v1.ObjectReference"
      },
      "description": "Secrets is the list of secrets allowed to be used by pods running using this ServiceAccount. More info: http://releases.k8s.io/release-1.2/docs/user-guide/secrets.md"
     },
     "imagePullSecrets": {
      "type": "array",
      "items": {
       "$ref": "v1.LocalObjectReference"
      },
      "description": "ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: http://releases.k8s.io/release-1.2/docs/user-guide/secrets.md#manually-specifying-an-imagepullsecret"
     }
    }
   },
   "v1.ServiceList": {
    "id": "v1.ServiceList",
    "description": "ServiceList holds a list of services.",
    "required": [
     "items"
    ],
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "unversioned.ListMeta",
      "description": "Standard list metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "items": {
      "type": "array",
      "items": {
       "$ref": "v1.Service"
      },
      "description": "List of services"
     }
    }
   },
   "v1.Service": {
    "id": "v1.Service",
    "description": "Service is a named abstraction of software service (for example, mysql) consisting of local port (for example 3306) that the proxy listens on, and the selector that determines which pods will answer requests sent through the proxy.",
    "properties": {
     "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#types-kinds"
     },
     "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#resources"
     },
     "metadata": {
      "$ref": "v1.ObjectMeta",
      "description": "Standard object's metadata. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#metadata"
     },
     "spec": {
      "$ref": "v1.ServiceSpec",
      "description": "Spec defines the behavior of a service. http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#spec-and-status"
     },
     "status": {
      "$ref": "v1.ServiceStatus",
      "description": "Most recently observed status of the service. Populated by the system. Read-only. More info: http://releases.k8s.io/release-1.2/docs/devel/api-conventions.md#spec-and-status"
     }
    }
   },
   "v1.ServiceSpec": {
    "id": "v1.ServiceSpec",
    "description": "ServiceSpec describes the attributes that a user creates on a service.",
    "required": [
     "ports"
    ],
    "properties": {
     "ports": {
      "type": "array",
      "items": {
       "$ref": "v1.ServicePort"
      },
      "description": "The list of ports that are exposed by this service. More info: http://releases.k8s.io/release-1.2/docs/user-guide/services.md#virtual-ips-and-service-proxies"
     },
     "selector": {
      "type": "any",
      "description": "This service will route traffic to pods having labels matching this selector. Label keys and values that must match in order to receive traffic for this service. If empty, all pods are selected, if not specified, endpoints must be manually specified. More info: http://releases.k8s.io/release-1.2/docs/user-guide/services.md#overview"
     },
     "clusterIP": {
      "type": "string",
      "description": "ClusterIP is usually assigned by the master and is the IP address of the service. If specified, it will be allocated to the service if it is unused or else creation of the service will fail. Valid values are None, empty string (\"\"), or a valid IP address. 'None' can be specified for a headless service when proxying is not required. Cannot be updated. More info: http://releases.k8s.io/release-1.2/docs/user-guide/services.md#virtual-ips-and-service-proxies"
     },
     "type": {
      "type": "string",
      "description": "Type of exposed service. Must be ClusterIP, NodePort, or LoadBalancer. Defaults to ClusterIP. More info: http://releases.k8s.io/release-1.2/docs/user-guide/services.md#external-services"
     },
     "externalIPs": {
      "type": "array",
      "items": {
       "type": "string"
      },
      "description": "externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service.  These IPs are not managed by Kubernetes.  The user is responsible for ensuring that traffic arrives at a node with this IP.  A common example is external load-balancers that are not part of the Kubernetes system.  A previous form of this functionality exists as the deprecatedPublicIPs field.  When using this field, callers should also clear the deprecatedPublicIPs field."
     },
     "deprecatedPublicIPs": {
      "type": "array",
      "items": {
       "type": "string"
      },
      "description": "deprecatedPublicIPs is deprecated and replaced by the externalIPs field with almost the exact same semantics.  This field is retained in the v1 API for compatibility until at least 8/20/2016.  It will be removed from any new API revisions.  If both deprecatedPublicIPs *and* externalIPs are set, deprecatedPublicIPs is used."
     },
     "sessionAffinity": {
      "type": "string",
      "description": "Supports \"ClientIP\" and \"None\". Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: http://releases.k8s.io/release-1.2/docs/user-guide/services.md#virtual-ips-and-service-proxies"
     },
     "loadBalancerIP": {
      "type": "string",
      "description": "Only applies to Service Type: LoadBalancer LoadBalancer will get created with the IP specified in this field. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature."
     }
    }
   },
   "v1.ServicePort": {
    "id": "v1.ServicePort",
    "description": "ServicePort contains information on service's port.",
    "required": [
     "port"
    ],
    "properties": {
     "name": {
      "type": "string",
      "description": "The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. This maps to the 'Name' field in EndpointPort objects. Optional if only one ServicePort is defined on this service."
     },
     "protocol": {
      "type": "string",
      "description": "The IP protocol for this port. Supports \"TCP\" and \"UDP\". Default is TCP."
     },
     "port": {
      "type": "integer",
      "format": "int32",
      "description": "The port that will be exposed by this service."
     },
     "targetPort": {
      "type": "string",
      "description": "Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field. More info: http://releases.k8s.io/release-1.2/docs/user-guide/services.md#defining-a-service"
     },
     "nodePort": {
      "type": "integer",
      "format": "int32",
      "description": "The port on each node on which this service is exposed when type=NodePort or LoadBalancer. Usually assigned by the system. If specified, it will be allocated to the service if unused or else creation of the service will fail. Default is to auto-allocate a port if the ServiceType of this Service requires one. More info: http://releases.k8s.io/release-1.2/docs/user-guide/services.md#type--nodeport"
     }
    }
   },
   "v1.ServiceStatus": {
    "id": "v1.ServiceStatus",
    "description": "ServiceStatus represents the current status of a service.",
    "properties": {
     "loadBalancer": {
      "$ref": "v1.LoadBalancerStatus",
      "description": "LoadBalancer contains the current status of the load-balancer, if one is present."
     }
    }
   },
   "v1.LoadBalancerStatus": {
    "id": "v1.LoadBalancerStatus",
    "description": "LoadBalancerStatus represents the status of a load-balancer.",
    "properties": {
     "ingress": {
      "type": "array",
      "items": {
       "$ref": "v1.LoadBalancerIngress"
      },
      "description": "Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points."
     }
    }
   },
   "v1.LoadBalancerIngress": {
    "id": "v1.LoadBalancerIngress",
    "description": "LoadBalancerIngress represents the status of a load-balancer ingress point: traffic intended for the service should be sent to an ingress point.",
    "properties": {
     "ip": {
      "type": "string",
      "description": "IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)"
     },
     "hostname": {
      "type": "string",
      "description": "Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)"
     }
    }
   }
  }
 }`
