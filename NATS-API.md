# NATS configuration proxy 0.0.1 documentation

* License: [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0)

This application dynamically changes the configuration of a running NATS server

## Table of Contents

* [Operations](#operations)
  * [SUB accounts.v1.auth.idents](#sub-accountsv1authidents-operation)
  * [PUB accounts.v1.auth.idents.get](#pub-accountsv1authidentsget-operation)
  * [PUB accounts.v1.auth.idents.delete](#pub-accountsv1authidentsdelete-operation)
  * [SUB accounts.v1.auth.idents.{name}](#sub-accountsv1authidentsname-operation)
  * [PUB accounts.v1.auth.idents.{name}.get](#pub-accountsv1authidentsnameget-operation)
  * [PUB accounts.v1.auth.idents.{name}.update](#pub-accountsv1authidentsnameupdate-operation)
  * [PUB accounts.v1.auth.idents.{name}.create](#pub-accountsv1authidentsnamecreate-operation)
  * [PUB accounts.v1.auth.idents.{name}.delete](#pub-accountsv1authidentsnamedelete-operation)
  * [SUB accounts.v1.auth.perms](#sub-accountsv1authperms-operation)
  * [PUB accounts.v1.auth.perms.get](#pub-accountsv1authpermsget-operation)
  * [PUB accounts.v1.auth.perms.delete](#pub-accountsv1authpermsdelete-operation)
  * [SUB accounts.v1.auth.perms.{name}](#sub-accountsv1authpermsname-operation)
  * [PUB accounts.v1.auth.perms.{name}.get](#pub-accountsv1authpermsnameget-operation)
  * [PUB accounts.v1.auth.perms.{name}.update](#pub-accountsv1authpermsnameupdate-operation)
  * [PUB accounts.v1.auth.perms.{name}.delete](#pub-accountsv1authpermsnamedelete-operation)
  * [SUB accounts.v1.auth.accounts](#sub-accountsv1authaccounts-operation)
  * [PUB accounts.v1.auth.accounts.get](#pub-accountsv1authaccountsget-operation)
  * [SUB accounts.v1.auth.accounts.{name}](#sub-accountsv1authaccountsname-operation)
  * [PUB accounts.v1.auth.accounts.{name}.get](#pub-accountsv1authaccountsnameget-operation)
  * [PUB accounts.v1.auth.accounts.{name}.update](#pub-accountsv1authaccountsnameupdate-operation)
  * [PUB accounts.v1.auth.accounts.{name}.create](#pub-accountsv1authaccountsnamecreate-operation)
  * [PUB accounts.v1.auth.accounts.{name}.delete](#pub-accountsv1authaccountsnamedelete-operation)
  * [PUB accounts.v.auth.jetstream.update](#pub-accountsvauthjetstreamupdate-operation)
  * [PUB accounts.v.auth.jetstream.create](#pub-accountsvauthjetstreamcreate-operation)
  * [PUB accounts.v.auth.jetstream.delete](#pub-accountsvauthjetstreamdelete-operation)

## Operations

### SUB `accounts.v1.auth.idents` Operation

Get list of identities

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| identities | Identity related | - |

#### Message `Identity`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | oneOf | - | - | - | **additional properties are allowed** |
| 0 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| username | string | - | - | - | - |
| password | string | - | - | - | - |
| 1 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| 1.nkey | string | - | - | - | - |
| 2 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| 2.username | string | - | - | - | - |
| 2.permissions | string | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "username": "string",
  "password": "string"
}
```



### PUB `accounts.v1.auth.idents.get` Operation

Get list of identities

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| identities | Identity related | - |

#### Message `Identity`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | oneOf | - | - | - | **additional properties are allowed** |
| 0 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| username | string | - | - | - | - |
| password | string | - | - | - | - |
| 1 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| 1.nkey | string | - | - | - | - |
| 2 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| 2.username | string | - | - | - | - |
| 2.permissions | string | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "username": "string",
  "password": "string"
}
```



### PUB `accounts.v1.auth.idents.delete` Operation

Delete all identities

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| identities | Identity related | - |

#### Message `Identity`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | oneOf | - | - | - | **additional properties are allowed** |
| 0 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| username | string | - | - | - | - |
| password | string | - | - | - | - |
| 1 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| 1.nkey | string | - | - | - | - |
| 2 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| 2.username | string | - | - | - | - |
| 2.permissions | string | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "username": "string",
  "password": "string"
}
```



### SUB `accounts.v1.auth.idents.{name}` Operation

Get specific identity w/ permissions

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| identities | Identity related | - |

#### Parameters

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| name | string | name of the account to look up | - | - | **required** |


#### Message `Identity`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | oneOf | - | - | - | **additional properties are allowed** |
| 0 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| username | string | - | - | - | - |
| password | string | - | - | - | - |
| 1 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| 1.nkey | string | - | - | - | - |
| 2 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| 2.username | string | - | - | - | - |
| 2.permissions | string | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "username": "string",
  "password": "string"
}
```



### PUB `accounts.v1.auth.idents.{name}.get` Operation

Get specific identity w/ permissions

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| identities | Identity related | - |

#### Parameters

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| name | string | name of the account to look up | - | - | **required** |


#### Message `Identity`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | oneOf | - | - | - | **additional properties are allowed** |
| 0 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| username | string | - | - | - | - |
| password | string | - | - | - | - |
| 1 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| 1.nkey | string | - | - | - | - |
| 2 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| 2.username | string | - | - | - | - |
| 2.permissions | string | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "username": "string",
  "password": "string"
}
```



### PUB `accounts.v1.auth.idents.{name}.update` Operation

Update the specific identify

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| identities | Identity related | - |

#### Parameters

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| name | string | name of the account to look up | - | - | **required** |


#### Message `Identity`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | oneOf | - | - | - | **additional properties are allowed** |
| 0 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| username | string | - | - | - | - |
| password | string | - | - | - | - |
| 1 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| 1.nkey | string | - | - | - | - |
| 2 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| 2.username | string | - | - | - | - |
| 2.permissions | string | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "username": "string",
  "password": "string"
}
```



### PUB `accounts.v1.auth.idents.{name}.create` Operation

Create the specific identify

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| identities | Identity related | - |

#### Parameters

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| name | string | name of the account to look up | - | - | **required** |


#### Message `Identity`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | oneOf | - | - | - | **additional properties are allowed** |
| 0 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| username | string | - | - | - | - |
| password | string | - | - | - | - |
| 1 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| 1.nkey | string | - | - | - | - |
| 2 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| 2.username | string | - | - | - | - |
| 2.permissions | string | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "username": "string",
  "password": "string"
}
```



### PUB `accounts.v1.auth.idents.{name}.delete` Operation

Delete the specific identify

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| identities | Identity related | - |

#### Parameters

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| name | string | name of the account to look up | - | - | **required** |


#### Message `Identity`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | oneOf | - | - | - | **additional properties are allowed** |
| 0 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| username | string | - | - | - | - |
| password | string | - | - | - | - |
| 1 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| 1.nkey | string | - | - | - | - |
| 2 (oneOf item) | object | - | - | - | **additional properties are NOT allowed** |
| 2.username | string | - | - | - | - |
| 2.permissions | string | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "username": "string",
  "password": "string"
}
```



### SUB `accounts.v1.auth.perms` Operation

Get list of named permission sets

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| permissions | Permission related | - |

#### Message `Permissions`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are NOT allowed** |
| publish | object | - | - | - | **additional properties are NOT allowed** |
| publish.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| publish.allow.0 (index) | string | - | - | - | - |
| publish.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| publish.deny.0 (index) | string | - | - | - | - |
| subscribe | object | - | - | - | **additional properties are NOT allowed** |
| subscribe.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| subscribe.allow.0 (index) | string | - | - | - | - |
| subscribe.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| subscribe.deny.0 (index) | string | - | - | - | - |
| responses | object | - | - | - | **additional properties are NOT allowed** |
| responses.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| responses.allow.0 (index) | string | - | - | - | - |
| responses.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| responses.deny.0 (index) | string | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "publish": {
    "allow": [
      "string"
    ],
    "deny": [
      "string"
    ]
  },
  "subscribe": {
    "allow": [
      "string"
    ],
    "deny": [
      "string"
    ]
  },
  "responses": {
    "allow": [
      "string"
    ],
    "deny": [
      "string"
    ]
  }
}
```



### PUB `accounts.v1.auth.perms.get` Operation

Get list of named permission sets

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| permissions | Permission related | - |

#### Message `EmptyMessage`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | null | - | - | - | - |

> Examples of payload _(generated)_

```json
""
```



### PUB `accounts.v1.auth.perms.delete` Operation

Delete all permissions

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| permissions | Permission related | - |

#### Message `Permissions`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are NOT allowed** |
| publish | object | - | - | - | **additional properties are NOT allowed** |
| publish.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| publish.allow.0 (index) | string | - | - | - | - |
| publish.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| publish.deny.0 (index) | string | - | - | - | - |
| subscribe | object | - | - | - | **additional properties are NOT allowed** |
| subscribe.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| subscribe.allow.0 (index) | string | - | - | - | - |
| subscribe.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| subscribe.deny.0 (index) | string | - | - | - | - |
| responses | object | - | - | - | **additional properties are NOT allowed** |
| responses.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| responses.allow.0 (index) | string | - | - | - | - |
| responses.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| responses.deny.0 (index) | string | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "publish": {
    "allow": [
      "string"
    ],
    "deny": [
      "string"
    ]
  },
  "subscribe": {
    "allow": [
      "string"
    ],
    "deny": [
      "string"
    ]
  },
  "responses": {
    "allow": [
      "string"
    ],
    "deny": [
      "string"
    ]
  }
}
```



### SUB `accounts.v1.auth.perms.{name}` Operation

Get specific permissions

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| permissions | Permission related | - |

#### Parameters

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| name | string | name of the account to look up | - | - | **required** |


#### Message `Permissions`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are NOT allowed** |
| publish | object | - | - | - | **additional properties are NOT allowed** |
| publish.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| publish.allow.0 (index) | string | - | - | - | - |
| publish.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| publish.deny.0 (index) | string | - | - | - | - |
| subscribe | object | - | - | - | **additional properties are NOT allowed** |
| subscribe.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| subscribe.allow.0 (index) | string | - | - | - | - |
| subscribe.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| subscribe.deny.0 (index) | string | - | - | - | - |
| responses | object | - | - | - | **additional properties are NOT allowed** |
| responses.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| responses.allow.0 (index) | string | - | - | - | - |
| responses.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| responses.deny.0 (index) | string | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "publish": {
    "allow": [
      "string"
    ],
    "deny": [
      "string"
    ]
  },
  "subscribe": {
    "allow": [
      "string"
    ],
    "deny": [
      "string"
    ]
  },
  "responses": {
    "allow": [
      "string"
    ],
    "deny": [
      "string"
    ]
  }
}
```



### PUB `accounts.v1.auth.perms.{name}.get` Operation

Get specific permission sets

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| permissions | Permission related | - |

#### Parameters

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| name | string | name of the account to look up | - | - | **required** |


#### Message `EmptyMessage`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | null | - | - | - | - |

> Examples of payload _(generated)_

```json
""
```



### PUB `accounts.v1.auth.perms.{name}.update` Operation

Update the named permission sets

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| permissions | Permission related | - |

#### Parameters

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| name | string | name of the account to look up | - | - | **required** |


#### Message `Permissions`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are NOT allowed** |
| publish | object | - | - | - | **additional properties are NOT allowed** |
| publish.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| publish.allow.0 (index) | string | - | - | - | - |
| publish.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| publish.deny.0 (index) | string | - | - | - | - |
| subscribe | object | - | - | - | **additional properties are NOT allowed** |
| subscribe.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| subscribe.allow.0 (index) | string | - | - | - | - |
| subscribe.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| subscribe.deny.0 (index) | string | - | - | - | - |
| responses | object | - | - | - | **additional properties are NOT allowed** |
| responses.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| responses.allow.0 (index) | string | - | - | - | - |
| responses.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| responses.deny.0 (index) | string | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "publish": {
    "allow": [
      "string"
    ],
    "deny": [
      "string"
    ]
  },
  "subscribe": {
    "allow": [
      "string"
    ],
    "deny": [
      "string"
    ]
  },
  "responses": {
    "allow": [
      "string"
    ],
    "deny": [
      "string"
    ]
  }
}
```



### PUB `accounts.v1.auth.perms.{name}.delete` Operation

Delete named permission sets

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| permissions | Permission related | - |

#### Parameters

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| name | string | name of the account to look up | - | - | **required** |


#### Message `Permissions`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are NOT allowed** |
| publish | object | - | - | - | **additional properties are NOT allowed** |
| publish.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| publish.allow.0 (index) | string | - | - | - | - |
| publish.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| publish.deny.0 (index) | string | - | - | - | - |
| subscribe | object | - | - | - | **additional properties are NOT allowed** |
| subscribe.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| subscribe.allow.0 (index) | string | - | - | - | - |
| subscribe.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| subscribe.deny.0 (index) | string | - | - | - | - |
| responses | object | - | - | - | **additional properties are NOT allowed** |
| responses.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| responses.allow.0 (index) | string | - | - | - | - |
| responses.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| responses.deny.0 (index) | string | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "publish": {
    "allow": [
      "string"
    ],
    "deny": [
      "string"
    ]
  },
  "subscribe": {
    "allow": [
      "string"
    ],
    "deny": [
      "string"
    ]
  },
  "responses": {
    "allow": [
      "string"
    ],
    "deny": [
      "string"
    ]
  }
}
```



### SUB `accounts.v1.auth.accounts` Operation

Get list of accounts

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| accounts | Account related | - |

#### Message `Accounts`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | array<object> | - | - | - | - |
| users | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| users.0.username | string | - | - | - | - |
| users.0.password | string | - | - | - | - |
| users.0.nkey | string | - | - | - | - |
| users.0.permissions | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.publish | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.publish.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.publish.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.publish.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.publish.deny.0 (index) | string | - | - | - | - |
| users.0.permissions.subscribe | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.subscribe.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.subscribe.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.subscribe.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.subscribe.deny.0 (index) | string | - | - | - | - |
| users.0.permissions.responses | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.responses.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.responses.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.responses.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.responses.deny.0 (index) | string | - | - | - | - |
| exports | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| exports.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| exports.0.stream | string | - | - | - | - |
| exports.0.service | string | - | - | - | - |
| exports.0.Accounts | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| exports.0.Accounts.0 (index) | string | - | - | - | - |
| exports.0.response | string | - | - | - | - |
| imports | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| imports.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.service | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.service.acocount | string | - | - | - | - |
| imports.0.service.subject | string | - | - | - | - |
| imports.0.stream | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.stream.acocount | string | - | - | - | - |
| imports.0.stream.subject | string | - | - | - | - |
| imports.0.prefix | string | - | - | - | - |
| imports.0.to | string | - | - | - | - |
| jetstream | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| jetstream.0 (index) | object | - | - | - | - |
| jetstream.0.enabled | boolean | - | - | - | - |
| jetstream.0.password | string | - | - | - | - |
| jetstream.0.nkey | string | - | - | - | - |
| jetstream.0 (additional properties) | object | - | - | - | **additional properties are NOT allowed** |
| jetstream.0.max_mem | integer | - | - | - | - |
| jetstream.0.max_file | integer | - | - | - | - |
| jetstream.0.max_streams | integer | - | - | - | - |
| jetstream.0.max_consumers | integer | - | - | - | - |

> Examples of payload _(generated)_

```json
[
  {
    "users": [
      {
        "username": "string",
        "password": "string",
        "nkey": "string",
        "permissions": {
          "publish": {
            "allow": [
              "string"
            ],
            "deny": [
              "string"
            ]
          },
          "subscribe": {
            "allow": [
              "string"
            ],
            "deny": [
              "string"
            ]
          },
          "responses": {
            "allow": [
              "string"
            ],
            "deny": [
              "string"
            ]
          }
        }
      }
    ],
    "exports": [
      {
        "stream": "string",
        "service": "string",
        "Accounts": [
          "string"
        ],
        "response": "string"
      }
    ],
    "imports": [
      {
        "service": {
          "acocount": "string",
          "subject": "string"
        },
        "stream": {
          "acocount": "string",
          "subject": "string"
        },
        "prefix": "string",
        "to": "string"
      }
    ],
    "jetstream": [
      {
        "enabled": true,
        "password": "string",
        "nkey": "string",
        "property1": {
          "max_mem": 0,
          "max_file": 0,
          "max_streams": 0,
          "max_consumers": 0
        },
        "property2": {
          "max_mem": 0,
          "max_file": 0,
          "max_streams": 0,
          "max_consumers": 0
        }
      }
    ]
  }
]
```



### PUB `accounts.v1.auth.accounts.get` Operation

Get list of accounts

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| accounts | Account related | - |

#### Message `Accounts`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | array<object> | - | - | - | - |
| users | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| users.0.username | string | - | - | - | - |
| users.0.password | string | - | - | - | - |
| users.0.nkey | string | - | - | - | - |
| users.0.permissions | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.publish | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.publish.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.publish.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.publish.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.publish.deny.0 (index) | string | - | - | - | - |
| users.0.permissions.subscribe | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.subscribe.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.subscribe.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.subscribe.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.subscribe.deny.0 (index) | string | - | - | - | - |
| users.0.permissions.responses | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.responses.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.responses.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.responses.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.responses.deny.0 (index) | string | - | - | - | - |
| exports | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| exports.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| exports.0.stream | string | - | - | - | - |
| exports.0.service | string | - | - | - | - |
| exports.0.Accounts | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| exports.0.Accounts.0 (index) | string | - | - | - | - |
| exports.0.response | string | - | - | - | - |
| imports | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| imports.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.service | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.service.acocount | string | - | - | - | - |
| imports.0.service.subject | string | - | - | - | - |
| imports.0.stream | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.stream.acocount | string | - | - | - | - |
| imports.0.stream.subject | string | - | - | - | - |
| imports.0.prefix | string | - | - | - | - |
| imports.0.to | string | - | - | - | - |
| jetstream | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| jetstream.0 (index) | object | - | - | - | - |
| jetstream.0.enabled | boolean | - | - | - | - |
| jetstream.0.password | string | - | - | - | - |
| jetstream.0.nkey | string | - | - | - | - |
| jetstream.0 (additional properties) | object | - | - | - | **additional properties are NOT allowed** |
| jetstream.0.max_mem | integer | - | - | - | - |
| jetstream.0.max_file | integer | - | - | - | - |
| jetstream.0.max_streams | integer | - | - | - | - |
| jetstream.0.max_consumers | integer | - | - | - | - |

> Examples of payload _(generated)_

```json
[
  {
    "users": [
      {
        "username": "string",
        "password": "string",
        "nkey": "string",
        "permissions": {
          "publish": {
            "allow": [
              "string"
            ],
            "deny": [
              "string"
            ]
          },
          "subscribe": {
            "allow": [
              "string"
            ],
            "deny": [
              "string"
            ]
          },
          "responses": {
            "allow": [
              "string"
            ],
            "deny": [
              "string"
            ]
          }
        }
      }
    ],
    "exports": [
      {
        "stream": "string",
        "service": "string",
        "Accounts": [
          "string"
        ],
        "response": "string"
      }
    ],
    "imports": [
      {
        "service": {
          "acocount": "string",
          "subject": "string"
        },
        "stream": {
          "acocount": "string",
          "subject": "string"
        },
        "prefix": "string",
        "to": "string"
      }
    ],
    "jetstream": [
      {
        "enabled": true,
        "password": "string",
        "nkey": "string",
        "property1": {
          "max_mem": 0,
          "max_file": 0,
          "max_streams": 0,
          "max_consumers": 0
        },
        "property2": {
          "max_mem": 0,
          "max_file": 0,
          "max_streams": 0,
          "max_consumers": 0
        }
      }
    ]
  }
]
```



### SUB `accounts.v1.auth.accounts.{name}` Operation

Get named account

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| accounts | Account related | - |

#### Parameters

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| name | string | name of the account to look up | - | - | **required** |


#### Message `Account`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are NOT allowed** |
| users | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| users.0.username | string | - | - | - | - |
| users.0.password | string | - | - | - | - |
| users.0.nkey | string | - | - | - | - |
| users.0.permissions | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.publish | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.publish.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.publish.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.publish.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.publish.deny.0 (index) | string | - | - | - | - |
| users.0.permissions.subscribe | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.subscribe.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.subscribe.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.subscribe.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.subscribe.deny.0 (index) | string | - | - | - | - |
| users.0.permissions.responses | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.responses.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.responses.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.responses.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.responses.deny.0 (index) | string | - | - | - | - |
| exports | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| exports.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| exports.0.stream | string | - | - | - | - |
| exports.0.service | string | - | - | - | - |
| exports.0.Accounts | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| exports.0.Accounts.0 (index) | string | - | - | - | - |
| exports.0.response | string | - | - | - | - |
| imports | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| imports.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.service | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.service.acocount | string | - | - | - | - |
| imports.0.service.subject | string | - | - | - | - |
| imports.0.stream | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.stream.acocount | string | - | - | - | - |
| imports.0.stream.subject | string | - | - | - | - |
| imports.0.prefix | string | - | - | - | - |
| imports.0.to | string | - | - | - | - |
| jetstream | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| jetstream.0 (index) | object | - | - | - | - |
| jetstream.0.enabled | boolean | - | - | - | - |
| jetstream.0.password | string | - | - | - | - |
| jetstream.0.nkey | string | - | - | - | - |
| jetstream.0 (additional properties) | object | - | - | - | **additional properties are NOT allowed** |
| jetstream.0.max_mem | integer | - | - | - | - |
| jetstream.0.max_file | integer | - | - | - | - |
| jetstream.0.max_streams | integer | - | - | - | - |
| jetstream.0.max_consumers | integer | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "users": [
    {
      "username": "string",
      "password": "string",
      "nkey": "string",
      "permissions": {
        "publish": {
          "allow": [
            "string"
          ],
          "deny": [
            "string"
          ]
        },
        "subscribe": {
          "allow": [
            "string"
          ],
          "deny": [
            "string"
          ]
        },
        "responses": {
          "allow": [
            "string"
          ],
          "deny": [
            "string"
          ]
        }
      }
    }
  ],
  "exports": [
    {
      "stream": "string",
      "service": "string",
      "Accounts": [
        "string"
      ],
      "response": "string"
    }
  ],
  "imports": [
    {
      "service": {
        "acocount": "string",
        "subject": "string"
      },
      "stream": {
        "acocount": "string",
        "subject": "string"
      },
      "prefix": "string",
      "to": "string"
    }
  ],
  "jetstream": [
    {
      "enabled": true,
      "password": "string",
      "nkey": "string",
      "property1": {
        "max_mem": 0,
        "max_file": 0,
        "max_streams": 0,
        "max_consumers": 0
      },
      "property2": {
        "max_mem": 0,
        "max_file": 0,
        "max_streams": 0,
        "max_consumers": 0
      }
    }
  ]
}
```



### PUB `accounts.v1.auth.accounts.{name}.get` Operation

Get named account

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| accounts | Account related | - |

#### Parameters

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| name | string | name of the account to look up | - | - | **required** |


#### Message `Account`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are NOT allowed** |
| users | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| users.0.username | string | - | - | - | - |
| users.0.password | string | - | - | - | - |
| users.0.nkey | string | - | - | - | - |
| users.0.permissions | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.publish | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.publish.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.publish.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.publish.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.publish.deny.0 (index) | string | - | - | - | - |
| users.0.permissions.subscribe | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.subscribe.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.subscribe.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.subscribe.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.subscribe.deny.0 (index) | string | - | - | - | - |
| users.0.permissions.responses | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.responses.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.responses.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.responses.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.responses.deny.0 (index) | string | - | - | - | - |
| exports | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| exports.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| exports.0.stream | string | - | - | - | - |
| exports.0.service | string | - | - | - | - |
| exports.0.Accounts | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| exports.0.Accounts.0 (index) | string | - | - | - | - |
| exports.0.response | string | - | - | - | - |
| imports | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| imports.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.service | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.service.acocount | string | - | - | - | - |
| imports.0.service.subject | string | - | - | - | - |
| imports.0.stream | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.stream.acocount | string | - | - | - | - |
| imports.0.stream.subject | string | - | - | - | - |
| imports.0.prefix | string | - | - | - | - |
| imports.0.to | string | - | - | - | - |
| jetstream | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| jetstream.0 (index) | object | - | - | - | - |
| jetstream.0.enabled | boolean | - | - | - | - |
| jetstream.0.password | string | - | - | - | - |
| jetstream.0.nkey | string | - | - | - | - |
| jetstream.0 (additional properties) | object | - | - | - | **additional properties are NOT allowed** |
| jetstream.0.max_mem | integer | - | - | - | - |
| jetstream.0.max_file | integer | - | - | - | - |
| jetstream.0.max_streams | integer | - | - | - | - |
| jetstream.0.max_consumers | integer | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "users": [
    {
      "username": "string",
      "password": "string",
      "nkey": "string",
      "permissions": {
        "publish": {
          "allow": [
            "string"
          ],
          "deny": [
            "string"
          ]
        },
        "subscribe": {
          "allow": [
            "string"
          ],
          "deny": [
            "string"
          ]
        },
        "responses": {
          "allow": [
            "string"
          ],
          "deny": [
            "string"
          ]
        }
      }
    }
  ],
  "exports": [
    {
      "stream": "string",
      "service": "string",
      "Accounts": [
        "string"
      ],
      "response": "string"
    }
  ],
  "imports": [
    {
      "service": {
        "acocount": "string",
        "subject": "string"
      },
      "stream": {
        "acocount": "string",
        "subject": "string"
      },
      "prefix": "string",
      "to": "string"
    }
  ],
  "jetstream": [
    {
      "enabled": true,
      "password": "string",
      "nkey": "string",
      "property1": {
        "max_mem": 0,
        "max_file": 0,
        "max_streams": 0,
        "max_consumers": 0
      },
      "property2": {
        "max_mem": 0,
        "max_file": 0,
        "max_streams": 0,
        "max_consumers": 0
      }
    }
  ]
}
```



### PUB `accounts.v1.auth.accounts.{name}.update` Operation

Update named account

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| accounts | Account related | - |

#### Parameters

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| name | string | name of the account to look up | - | - | **required** |


#### Message `Account`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are NOT allowed** |
| users | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| users.0.username | string | - | - | - | - |
| users.0.password | string | - | - | - | - |
| users.0.nkey | string | - | - | - | - |
| users.0.permissions | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.publish | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.publish.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.publish.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.publish.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.publish.deny.0 (index) | string | - | - | - | - |
| users.0.permissions.subscribe | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.subscribe.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.subscribe.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.subscribe.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.subscribe.deny.0 (index) | string | - | - | - | - |
| users.0.permissions.responses | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.responses.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.responses.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.responses.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.responses.deny.0 (index) | string | - | - | - | - |
| exports | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| exports.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| exports.0.stream | string | - | - | - | - |
| exports.0.service | string | - | - | - | - |
| exports.0.Accounts | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| exports.0.Accounts.0 (index) | string | - | - | - | - |
| exports.0.response | string | - | - | - | - |
| imports | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| imports.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.service | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.service.acocount | string | - | - | - | - |
| imports.0.service.subject | string | - | - | - | - |
| imports.0.stream | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.stream.acocount | string | - | - | - | - |
| imports.0.stream.subject | string | - | - | - | - |
| imports.0.prefix | string | - | - | - | - |
| imports.0.to | string | - | - | - | - |
| jetstream | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| jetstream.0 (index) | object | - | - | - | - |
| jetstream.0.enabled | boolean | - | - | - | - |
| jetstream.0.password | string | - | - | - | - |
| jetstream.0.nkey | string | - | - | - | - |
| jetstream.0 (additional properties) | object | - | - | - | **additional properties are NOT allowed** |
| jetstream.0.max_mem | integer | - | - | - | - |
| jetstream.0.max_file | integer | - | - | - | - |
| jetstream.0.max_streams | integer | - | - | - | - |
| jetstream.0.max_consumers | integer | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "users": [
    {
      "username": "string",
      "password": "string",
      "nkey": "string",
      "permissions": {
        "publish": {
          "allow": [
            "string"
          ],
          "deny": [
            "string"
          ]
        },
        "subscribe": {
          "allow": [
            "string"
          ],
          "deny": [
            "string"
          ]
        },
        "responses": {
          "allow": [
            "string"
          ],
          "deny": [
            "string"
          ]
        }
      }
    }
  ],
  "exports": [
    {
      "stream": "string",
      "service": "string",
      "Accounts": [
        "string"
      ],
      "response": "string"
    }
  ],
  "imports": [
    {
      "service": {
        "acocount": "string",
        "subject": "string"
      },
      "stream": {
        "acocount": "string",
        "subject": "string"
      },
      "prefix": "string",
      "to": "string"
    }
  ],
  "jetstream": [
    {
      "enabled": true,
      "password": "string",
      "nkey": "string",
      "property1": {
        "max_mem": 0,
        "max_file": 0,
        "max_streams": 0,
        "max_consumers": 0
      },
      "property2": {
        "max_mem": 0,
        "max_file": 0,
        "max_streams": 0,
        "max_consumers": 0
      }
    }
  ]
}
```



### PUB `accounts.v1.auth.accounts.{name}.create` Operation

Create named account

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| accounts | Account related | - |

#### Parameters

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| name | string | name of the account to look up | - | - | **required** |


#### Message `Account`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are NOT allowed** |
| users | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| users.0.username | string | - | - | - | - |
| users.0.password | string | - | - | - | - |
| users.0.nkey | string | - | - | - | - |
| users.0.permissions | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.publish | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.publish.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.publish.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.publish.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.publish.deny.0 (index) | string | - | - | - | - |
| users.0.permissions.subscribe | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.subscribe.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.subscribe.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.subscribe.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.subscribe.deny.0 (index) | string | - | - | - | - |
| users.0.permissions.responses | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.responses.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.responses.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.responses.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.responses.deny.0 (index) | string | - | - | - | - |
| exports | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| exports.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| exports.0.stream | string | - | - | - | - |
| exports.0.service | string | - | - | - | - |
| exports.0.Accounts | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| exports.0.Accounts.0 (index) | string | - | - | - | - |
| exports.0.response | string | - | - | - | - |
| imports | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| imports.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.service | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.service.acocount | string | - | - | - | - |
| imports.0.service.subject | string | - | - | - | - |
| imports.0.stream | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.stream.acocount | string | - | - | - | - |
| imports.0.stream.subject | string | - | - | - | - |
| imports.0.prefix | string | - | - | - | - |
| imports.0.to | string | - | - | - | - |
| jetstream | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| jetstream.0 (index) | object | - | - | - | - |
| jetstream.0.enabled | boolean | - | - | - | - |
| jetstream.0.password | string | - | - | - | - |
| jetstream.0.nkey | string | - | - | - | - |
| jetstream.0 (additional properties) | object | - | - | - | **additional properties are NOT allowed** |
| jetstream.0.max_mem | integer | - | - | - | - |
| jetstream.0.max_file | integer | - | - | - | - |
| jetstream.0.max_streams | integer | - | - | - | - |
| jetstream.0.max_consumers | integer | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "users": [
    {
      "username": "string",
      "password": "string",
      "nkey": "string",
      "permissions": {
        "publish": {
          "allow": [
            "string"
          ],
          "deny": [
            "string"
          ]
        },
        "subscribe": {
          "allow": [
            "string"
          ],
          "deny": [
            "string"
          ]
        },
        "responses": {
          "allow": [
            "string"
          ],
          "deny": [
            "string"
          ]
        }
      }
    }
  ],
  "exports": [
    {
      "stream": "string",
      "service": "string",
      "Accounts": [
        "string"
      ],
      "response": "string"
    }
  ],
  "imports": [
    {
      "service": {
        "acocount": "string",
        "subject": "string"
      },
      "stream": {
        "acocount": "string",
        "subject": "string"
      },
      "prefix": "string",
      "to": "string"
    }
  ],
  "jetstream": [
    {
      "enabled": true,
      "password": "string",
      "nkey": "string",
      "property1": {
        "max_mem": 0,
        "max_file": 0,
        "max_streams": 0,
        "max_consumers": 0
      },
      "property2": {
        "max_mem": 0,
        "max_file": 0,
        "max_streams": 0,
        "max_consumers": 0
      }
    }
  ]
}
```



### PUB `accounts.v1.auth.accounts.{name}.delete` Operation

Delete named account

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| accounts | Account related | - |

#### Parameters

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| name | string | name of the account to look up | - | - | **required** |


#### Message `Account`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are NOT allowed** |
| users | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| users.0.username | string | - | - | - | - |
| users.0.password | string | - | - | - | - |
| users.0.nkey | string | - | - | - | - |
| users.0.permissions | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.publish | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.publish.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.publish.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.publish.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.publish.deny.0 (index) | string | - | - | - | - |
| users.0.permissions.subscribe | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.subscribe.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.subscribe.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.subscribe.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.subscribe.deny.0 (index) | string | - | - | - | - |
| users.0.permissions.responses | object | - | - | - | **additional properties are NOT allowed** |
| users.0.permissions.responses.allow | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.responses.allow.0 (index) | string | - | - | - | - |
| users.0.permissions.responses.deny | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| users.0.permissions.responses.deny.0 (index) | string | - | - | - | - |
| exports | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| exports.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| exports.0.stream | string | - | - | - | - |
| exports.0.service | string | - | - | - | - |
| exports.0.Accounts | tuple<string, ...optional<any>> | - | - | - | **additional items are allowed** |
| exports.0.Accounts.0 (index) | string | - | - | - | - |
| exports.0.response | string | - | - | - | - |
| imports | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| imports.0 (index) | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.service | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.service.acocount | string | - | - | - | - |
| imports.0.service.subject | string | - | - | - | - |
| imports.0.stream | object | - | - | - | **additional properties are NOT allowed** |
| imports.0.stream.acocount | string | - | - | - | - |
| imports.0.stream.subject | string | - | - | - | - |
| imports.0.prefix | string | - | - | - | - |
| imports.0.to | string | - | - | - | - |
| jetstream | tuple<object, ...optional<any>> | - | - | - | **additional items are allowed** |
| jetstream.0 (index) | object | - | - | - | - |
| jetstream.0.enabled | boolean | - | - | - | - |
| jetstream.0.password | string | - | - | - | - |
| jetstream.0.nkey | string | - | - | - | - |
| jetstream.0 (additional properties) | object | - | - | - | **additional properties are NOT allowed** |
| jetstream.0.max_mem | integer | - | - | - | - |
| jetstream.0.max_file | integer | - | - | - | - |
| jetstream.0.max_streams | integer | - | - | - | - |
| jetstream.0.max_consumers | integer | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "users": [
    {
      "username": "string",
      "password": "string",
      "nkey": "string",
      "permissions": {
        "publish": {
          "allow": [
            "string"
          ],
          "deny": [
            "string"
          ]
        },
        "subscribe": {
          "allow": [
            "string"
          ],
          "deny": [
            "string"
          ]
        },
        "responses": {
          "allow": [
            "string"
          ],
          "deny": [
            "string"
          ]
        }
      }
    }
  ],
  "exports": [
    {
      "stream": "string",
      "service": "string",
      "Accounts": [
        "string"
      ],
      "response": "string"
    }
  ],
  "imports": [
    {
      "service": {
        "acocount": "string",
        "subject": "string"
      },
      "stream": {
        "acocount": "string",
        "subject": "string"
      },
      "prefix": "string",
      "to": "string"
    }
  ],
  "jetstream": [
    {
      "enabled": true,
      "password": "string",
      "nkey": "string",
      "property1": {
        "max_mem": 0,
        "max_file": 0,
        "max_streams": 0,
        "max_consumers": 0
      },
      "property2": {
        "max_mem": 0,
        "max_file": 0,
        "max_streams": 0,
        "max_consumers": 0
      }
    }
  ]
}
```



### PUB `accounts.v.auth.jetstream.update` Operation

Update JetStream config

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| jetstream | Jetstream related | - |

#### Message `JetstreamConfig`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| enabled | boolean | - | - | - | - |
| password | string | - | - | - | - |
| nkey | string | - | - | - | - |
| (additional properties) | object | - | - | - | **additional properties are NOT allowed** |
| max_mem | integer | - | - | - | - |
| max_file | integer | - | - | - | - |
| max_streams | integer | - | - | - | - |
| max_consumers | integer | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "enabled": true,
  "password": "string",
  "nkey": "string",
  "property1": {
    "max_mem": 0,
    "max_file": 0,
    "max_streams": 0,
    "max_consumers": 0
  },
  "property2": {
    "max_mem": 0,
    "max_file": 0,
    "max_streams": 0,
    "max_consumers": 0
  }
}
```



### PUB `accounts.v.auth.jetstream.create` Operation

Create JetStream config

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| jetstream | Jetstream related | - |

#### Message `JetstreamConfig`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| enabled | boolean | - | - | - | - |
| password | string | - | - | - | - |
| nkey | string | - | - | - | - |
| (additional properties) | object | - | - | - | **additional properties are NOT allowed** |
| max_mem | integer | - | - | - | - |
| max_file | integer | - | - | - | - |
| max_streams | integer | - | - | - | - |
| max_consumers | integer | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "enabled": true,
  "password": "string",
  "nkey": "string",
  "property1": {
    "max_mem": 0,
    "max_file": 0,
    "max_streams": 0,
    "max_consumers": 0
  },
  "property2": {
    "max_mem": 0,
    "max_file": 0,
    "max_streams": 0,
    "max_consumers": 0
  }
}
```



### PUB `accounts.v.auth.jetstream.delete` Operation

Delete JetStream config

##### Operation tags

| Name | Description | Documentation |
|---|---|---|
| jetstream | Jetstream related | - |

#### Message `JetstreamConfig`

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| enabled | boolean | - | - | - | - |
| password | string | - | - | - | - |
| nkey | string | - | - | - | - |
| (additional properties) | object | - | - | - | **additional properties are NOT allowed** |
| max_mem | integer | - | - | - | - |
| max_file | integer | - | - | - | - |
| max_streams | integer | - | - | - | - |
| max_consumers | integer | - | - | - | - |

> Examples of payload _(generated)_

```json
{
  "enabled": true,
  "password": "string",
  "nkey": "string",
  "property1": {
    "max_mem": 0,
    "max_file": 0,
    "max_streams": 0,
    "max_consumers": 0
  },
  "property2": {
    "max_mem": 0,
    "max_file": 0,
    "max_streams": 0,
    "max_consumers": 0
  }
}
```



