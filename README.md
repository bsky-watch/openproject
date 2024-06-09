# Go API client for openproject

You're looking at the current **stable** documentation of the OpenProject APIv3. If you're interested in the current
development version, please go to [github.com/opf](https://github.com/opf/openproject/tree/dev/docs/api/apiv3).

## Introduction

The documentation for the APIv3 is written according to the [OpenAPI 3.0 Specification](https://swagger.io/specification/).
You can either view the static version of this documentation on the [website](https://www.openproject.org/docs/api/introduction/)
or the interactive version, rendered with [OpenAPI Explorer](https://github.com/Rhosys/openapi-explorer/blob/main/README.md),
in your OpenProject installation under `/api/docs`.
In the latter you can try out the various API endpoints directly interacting with our OpenProject data.
Moreover you can access the specification source itself under `/api/v3/spec.json` and `/api/v3/spec.yml`
(e.g. [here](https://community.openproject.org/api/v3/spec.yml)).

The APIv3 is a hypermedia REST API, a shorthand for \"Hypermedia As The Engine Of Application State\" (HATEOAS).
This means that each endpoint of this API will have links to other resources or actions defined in the resulting body.

These related resources and actions for any given resource will be context sensitive. For example, only actions that the
authenticated user can take are being rendered. This can be used to dynamically identify actions that the user might take for any
given response.

As an example, if you fetch a work package through the [Work Package endpoint](https://www.openproject.org/docs/api/endpoints/work-packages/), the `update` link will only
be present when the user you authenticated has been granted a permission to update the work package in the assigned project.

## HAL+JSON

HAL is a simple format that gives a consistent and easy way to hyperlink between resources in your API.
Read more in the following specification: [https://tools.ietf.org/html/draft-kelly-json-hal-08](https://tools.ietf.org/html/draft-kelly-json-hal-08)

**OpenProject API implementation of HAL+JSON format** enriches JSON and introduces a few meta properties:

- `_type` - specifies the type of the resource (e.g.: WorkPackage, Project)
- `_links` - contains all related resource and action links available for the resource
- `_embedded` - contains all embedded objects

HAL does not guarantee that embedded resources are embedded in their full representation, they might as well be
partially represented (e.g. some properties can be left out).
However in this API you have the guarantee that whenever a resource is **embedded**, it is embedded in its **full representation**.

## API response structure

All API responses contain a single HAL+JSON object, even collections of objects are technically represented by
a single HAL+JSON object that itself contains its members. More details on collections can be found
in the [Collections Section](https://www.openproject.org/docs/api/collections/).

## Authentication

The API supports the following authentication schemes: OAuth2, session based authentication, and basic auth.

Depending on the settings of the OpenProject instance many resources can be accessed without being authenticated.
In case the instance requires authentication on all requests the client will receive an **HTTP 401** status code
in response to any request.

Otherwise unauthenticated clients have all the permissions of the anonymous user.

### Session-based Authentication

This means you have to login to OpenProject via the Web-Interface to be authenticated in the API.
This method is well-suited for clients acting within the browser, like the Angular-Client built into OpenProject.

In this case, you always need to pass the HTTP header `X-Requested-With \"XMLHttpRequest\"` for authentication.

### API Key through Basic Auth

Users can authenticate towards the API v3 using basic auth with the user name `apikey` (NOT your login) and the API key as the password.
Users can find their API key on their account page.

Example:

```shell
API_KEY=2519132cdf62dcf5a66fd96394672079f9e9cad1
curl -u apikey:$API_KEY https://community.openproject.org/api/v3/users/42
```

### OAuth2.0 authentication

OpenProject allows authentication and authorization with OAuth2 with *Authorization code flow*, as well as *Client credentials* operation modes.

To get started, you first need to register an application in the OpenProject OAuth administration section of your installation.
This will save an entry for your application with a client unique identifier (`client_id`) and an accompanying secret key (`client_secret`).

You can then use one the following guides to perform the supported OAuth 2.0 flows:

- [Authorization code flow](https://oauth.net/2/grant-types/authorization-code)

- [Authorization code flow with PKCE](https://doorkeeper.gitbook.io/guides/ruby-on-rails/pkce-flow), recommended for clients unable to keep the client_secret confidential.

- [Client credentials](https://oauth.net/2/grant-types/client-credentials/) - Requires an application to be bound to an impersonating user for non-public access

### Why not username and password?

The simplest way to do basic auth would be to use a user's username and password naturally.
However, OpenProject already has supported API keys in the past for the API v2, though not through basic auth.

Using **username and password** directly would have some advantages:

* It is intuitive for the user who then just has to provide those just as they would when logging into OpenProject.

* No extra logic for token management necessary.

On the other hand using **API keys** has some advantages too, which is why we went for that:

* If compromised while saved on an insecure client the user only has to regenerate the API key instead of changing their password, too.

* They are naturally long and random which makes them invulnerable to dictionary attacks and harder to crack in general.

Most importantly users may not actually have a password to begin with. Specifically when they have registered
through an OpenID Connect provider.

## Cross-Origin Resource Sharing (CORS)

By default, the OpenProject API is _not_ responding with any CORS headers.
If you want to allow cross-domain AJAX calls against your OpenProject instance, you need to enable CORS headers being returned.

Please see [our API settings documentation](https://www.openproject.org/docs/system-admin-guide/api-and-webhooks/) on
how to selectively enable CORS.

## Allowed HTTP methods

- `GET` - Get a single resource or collection of resources

- `POST` - Create a new resource or perform

- `PATCH` - Update a resource

- `DELETE` - Delete a resource

## Compression

Responses are compressed if requested by the client. Currently [gzip](https://www.gzip.org/) and [deflate](https://tools.ietf.org/html/rfc1951)
are supported. The client signals the desired compression by setting the [`Accept-Encoding` header](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.3).
If no `Accept-Encoding` header is send, `Accept-Encoding: identity` is assumed which will result in the API responding uncompressed.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3
- Package version: 1.0.0
- Generator version: 7.7.0-SNAPSHOT
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openproject "place/holder/openproject"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openproject.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openproject.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openproject.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openproject.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openproject.ContextOperationServerIndices` and `openproject.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openproject.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openproject.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://qa.openproject-edge.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ActionsCapabilitiesAPI* | [**ListActions**](docs/ActionsCapabilitiesAPI.md#listactions) | **Get** /api/v3/actions | List actions
*ActionsCapabilitiesAPI* | [**ListCapabilities**](docs/ActionsCapabilitiesAPI.md#listcapabilities) | **Get** /api/v3/capabilities | List capabilities
*ActionsCapabilitiesAPI* | [**ViewAction**](docs/ActionsCapabilitiesAPI.md#viewaction) | **Get** /api/v3/actions/{id} | View action
*ActionsCapabilitiesAPI* | [**ViewCapabilities**](docs/ActionsCapabilitiesAPI.md#viewcapabilities) | **Get** /api/v3/capabilities/{id} | View capabilities
*ActionsCapabilitiesAPI* | [**ViewGlobalContext**](docs/ActionsCapabilitiesAPI.md#viewglobalcontext) | **Get** /api/v3/capabilities/context/global | View global context
*ActivitiesAPI* | [**UpdateActivity**](docs/ActivitiesAPI.md#updateactivity) | **Patch** /api/v3/activities/{id} | Update activity
*ActivitiesAPI* | [**ViewActivity**](docs/ActivitiesAPI.md#viewactivity) | **Get** /api/v3/activities/{id} | View activity
*AttachmentsAPI* | [**AddAttachmentToMeeting**](docs/AttachmentsAPI.md#addattachmenttomeeting) | **Post** /api/v3/meetings/{id}/attachments | Add attachment to meeting
*AttachmentsAPI* | [**AddAttachmentToPost**](docs/AttachmentsAPI.md#addattachmenttopost) | **Post** /api/v3/posts/{id}/attachments | Add attachment to post
*AttachmentsAPI* | [**AddAttachmentToWikiPage**](docs/AttachmentsAPI.md#addattachmenttowikipage) | **Post** /api/v3/wiki_pages/{id}/attachments | Add attachment to wiki page
*AttachmentsAPI* | [**CreateAttachment**](docs/AttachmentsAPI.md#createattachment) | **Post** /api/v3/attachments | Create Attachment
*AttachmentsAPI* | [**CreateWorkPackageAttachment**](docs/AttachmentsAPI.md#createworkpackageattachment) | **Post** /api/v3/work_packages/{id}/attachments | Create work package attachment
*AttachmentsAPI* | [**DeleteAttachment**](docs/AttachmentsAPI.md#deleteattachment) | **Delete** /api/v3/attachments/{id} | Delete attachment
*AttachmentsAPI* | [**ListAttachmentsByMeeting**](docs/AttachmentsAPI.md#listattachmentsbymeeting) | **Get** /api/v3/meetings/{id}/attachments | List attachments by meeting
*AttachmentsAPI* | [**ListAttachmentsByPost**](docs/AttachmentsAPI.md#listattachmentsbypost) | **Get** /api/v3/posts/{id}/attachments | List attachments by post
*AttachmentsAPI* | [**ListAttachmentsByWikiPage**](docs/AttachmentsAPI.md#listattachmentsbywikipage) | **Get** /api/v3/wiki_pages/{id}/attachments | List attachments by wiki page
*AttachmentsAPI* | [**ListWorkPackageAttachments**](docs/AttachmentsAPI.md#listworkpackageattachments) | **Get** /api/v3/work_packages/{id}/attachments | List attachments by work package
*AttachmentsAPI* | [**ViewAttachment**](docs/AttachmentsAPI.md#viewattachment) | **Get** /api/v3/attachments/{id} | View attachment
*BudgetsAPI* | [**ViewBudget**](docs/BudgetsAPI.md#viewbudget) | **Get** /api/v3/budgets/{id} | view Budget
*BudgetsAPI* | [**ViewBudgetsOfAProject**](docs/BudgetsAPI.md#viewbudgetsofaproject) | **Get** /api/v3/projects/{id}/budgets | view Budgets of a Project
*CategoriesAPI* | [**ListCategoriesOfAProject**](docs/CategoriesAPI.md#listcategoriesofaproject) | **Get** /api/v3/projects/{id}/categories | List categories of a project
*CategoriesAPI* | [**ViewCategory**](docs/CategoriesAPI.md#viewcategory) | **Get** /api/v3/categories/{id} | View Category
*CollectionsAPI* | [**ViewAggregatedResult**](docs/CollectionsAPI.md#viewaggregatedresult) | **Get** /api/v3/examples | view aggregated result
*ConfigurationAPI* | [**ViewConfiguration**](docs/ConfigurationAPI.md#viewconfiguration) | **Get** /api/v3/configuration | View configuration
*CustomActionsAPI* | [**ExecuteCustomAction**](docs/CustomActionsAPI.md#executecustomaction) | **Post** /api/v3/custom_actions/{id}/execute | Execute custom action
*CustomActionsAPI* | [**GetCustomAction**](docs/CustomActionsAPI.md#getcustomaction) | **Get** /api/v3/custom_actions/{id} | Get a custom action
*CustomOptionsAPI* | [**ViewCustomOption**](docs/CustomOptionsAPI.md#viewcustomoption) | **Get** /api/v3/custom_options/{id} | View Custom Option
*DefaultAPI* | [**ViewValuesSchema**](docs/DefaultAPI.md#viewvaluesschema) | **Get** /api/v3/values/schema/{id} | View Values schema
*DocumentsAPI* | [**ListDocuments**](docs/DocumentsAPI.md#listdocuments) | **Get** /api/v3/documents | List Documents
*DocumentsAPI* | [**ViewDocument**](docs/DocumentsAPI.md#viewdocument) | **Get** /api/v3/documents/{id} | View document
*FileLinksAPI* | [**CreateStorage**](docs/FileLinksAPI.md#createstorage) | **Post** /api/v3/storages | Creates a storage.
*FileLinksAPI* | [**CreateStorageOauthCredentials**](docs/FileLinksAPI.md#createstorageoauthcredentials) | **Post** /api/v3/storages/{id}/oauth_client_credentials | Creates an oauth client credentials object for a storage.
*FileLinksAPI* | [**DeleteFileLink**](docs/FileLinksAPI.md#deletefilelink) | **Delete** /api/v3/file_links/{id} | Removes a file link.
*FileLinksAPI* | [**DeleteStorage**](docs/FileLinksAPI.md#deletestorage) | **Delete** /api/v3/storages/{id} | Delete a storage
*FileLinksAPI* | [**DownloadFileLink**](docs/FileLinksAPI.md#downloadfilelink) | **Get** /api/v3/file_links/{id}/download | Creates a download uri of the linked file.
*FileLinksAPI* | [**GetProjectStorage**](docs/FileLinksAPI.md#getprojectstorage) | **Get** /api/v3/project_storages/{id} | Gets a project storage
*FileLinksAPI* | [**GetStorage**](docs/FileLinksAPI.md#getstorage) | **Get** /api/v3/storages/{id} | Get a storage
*FileLinksAPI* | [**GetStorageFiles**](docs/FileLinksAPI.md#getstoragefiles) | **Get** /api/v3/storages/{id}/files | Gets files of a storage.
*FileLinksAPI* | [**ListProjectStorages**](docs/FileLinksAPI.md#listprojectstorages) | **Get** /api/v3/project_storages | Gets a list of project storages
*FileLinksAPI* | [**ListStorages**](docs/FileLinksAPI.md#liststorages) | **Get** /api/v3/storages | Get Storages
*FileLinksAPI* | [**OpenFileLink**](docs/FileLinksAPI.md#openfilelink) | **Get** /api/v3/file_links/{id}/open | Creates an opening uri of the linked file.
*FileLinksAPI* | [**OpenProjectStorage**](docs/FileLinksAPI.md#openprojectstorage) | **Get** /api/v3/project_storages/{id}/open | Open the project storage
*FileLinksAPI* | [**OpenStorage**](docs/FileLinksAPI.md#openstorage) | **Get** /api/v3/storages/{id}/open | Open the storage
*FileLinksAPI* | [**PrepareStorageFileUpload**](docs/FileLinksAPI.md#preparestoragefileupload) | **Post** /api/v3/storages/{id}/files/prepare_upload | Preparation of a direct upload of a file to the given storage.
*FileLinksAPI* | [**UpdateStorage**](docs/FileLinksAPI.md#updatestorage) | **Patch** /api/v3/storages/{id} | Update a storage
*FileLinksAPI* | [**ViewFileLink**](docs/FileLinksAPI.md#viewfilelink) | **Get** /api/v3/file_links/{id} | Gets a file link.
*FormsAPI* | [**ShowOrValidateForm**](docs/FormsAPI.md#showorvalidateform) | **Post** /api/v3/example/form | show or validate form
*GridsAPI* | [**CreateGrid**](docs/GridsAPI.md#creategrid) | **Post** /api/v3/grids | Create a grid
*GridsAPI* | [**GetGrid**](docs/GridsAPI.md#getgrid) | **Get** /api/v3/grids/{id} | Get a grid
*GridsAPI* | [**GridCreateForm**](docs/GridsAPI.md#gridcreateform) | **Post** /api/v3/grids/form | Grid Create Form
*GridsAPI* | [**GridUpdateForm**](docs/GridsAPI.md#gridupdateform) | **Post** /api/v3/grids/{id}/form | Grid Update Form
*GridsAPI* | [**ListGrids**](docs/GridsAPI.md#listgrids) | **Get** /api/v3/grids | List grids
*GridsAPI* | [**UpdateGrid**](docs/GridsAPI.md#updategrid) | **Patch** /api/v3/grids/{id} | Update a grid
*GroupsAPI* | [**CreateGroup**](docs/GroupsAPI.md#creategroup) | **Post** /api/v3/groups | Create group
*GroupsAPI* | [**DeleteGroup**](docs/GroupsAPI.md#deletegroup) | **Delete** /api/v3/groups/{id} | Delete group
*GroupsAPI* | [**GetGroup**](docs/GroupsAPI.md#getgroup) | **Get** /api/v3/groups/{id} | Get group
*GroupsAPI* | [**ListGroups**](docs/GroupsAPI.md#listgroups) | **Get** /api/v3/groups | List groups
*GroupsAPI* | [**UpdateGroup**](docs/GroupsAPI.md#updategroup) | **Patch** /api/v3/groups/{id} | Update group
*HelpTextsAPI* | [**GetHelpText**](docs/HelpTextsAPI.md#gethelptext) | **Get** /api/v3/help_texts/{id} | Get help text
*HelpTextsAPI* | [**ListHelpTexts**](docs/HelpTextsAPI.md#listhelptexts) | **Get** /api/v3/help_texts | List help texts
*MeetingsAPI* | [**ViewMeeting**](docs/MeetingsAPI.md#viewmeeting) | **Get** /api/v3/meetings/{id} | View Meeting Page
*MembershipsAPI* | [**CreateMembership**](docs/MembershipsAPI.md#createmembership) | **Post** /api/v3/memberships | Create a membership
*MembershipsAPI* | [**DeleteMembership**](docs/MembershipsAPI.md#deletemembership) | **Delete** /api/v3/memberships/{id} | Delete membership
*MembershipsAPI* | [**FormCreateMembership**](docs/MembershipsAPI.md#formcreatemembership) | **Post** /api/v3/memberships/form | Form create membership
*MembershipsAPI* | [**FormUpdateMembership**](docs/MembershipsAPI.md#formupdatemembership) | **Post** /api/v3/memberships/{id}/form | Form update membership
*MembershipsAPI* | [**GetMembership**](docs/MembershipsAPI.md#getmembership) | **Get** /api/v3/memberships/{id} | Get a membership
*MembershipsAPI* | [**GetMembershipSchema**](docs/MembershipsAPI.md#getmembershipschema) | **Get** /api/v3/memberships/schema | Schema membership
*MembershipsAPI* | [**GetMembershipsAvailableProjects**](docs/MembershipsAPI.md#getmembershipsavailableprojects) | **Get** /api/v3/memberships/available_projects | Available projects for memberships
*MembershipsAPI* | [**ListMemberships**](docs/MembershipsAPI.md#listmemberships) | **Get** /api/v3/memberships | List memberships
*MembershipsAPI* | [**UpdateMembership**](docs/MembershipsAPI.md#updatemembership) | **Patch** /api/v3/memberships/{id} | Update membership
*NewsAPI* | [**ListNews**](docs/NewsAPI.md#listnews) | **Get** /api/v3/news | List News
*NewsAPI* | [**ViewNews**](docs/NewsAPI.md#viewnews) | **Get** /api/v3/news/{id} | View news
*NotificationsAPI* | [**ListNotifications**](docs/NotificationsAPI.md#listnotifications) | **Get** /api/v3/notifications | Get notification collection
*NotificationsAPI* | [**ReadNotification**](docs/NotificationsAPI.md#readnotification) | **Post** /api/v3/notifications/{id}/read_ian | Read notification
*NotificationsAPI* | [**ReadNotifications**](docs/NotificationsAPI.md#readnotifications) | **Post** /api/v3/notifications/read_ian | Read all notifications
*NotificationsAPI* | [**UnreadNotification**](docs/NotificationsAPI.md#unreadnotification) | **Post** /api/v3/notifications/{id}/unread_ian | Unread notification
*NotificationsAPI* | [**UnreadNotifications**](docs/NotificationsAPI.md#unreadnotifications) | **Post** /api/v3/notifications/unread_ian | Unread all notifications
*NotificationsAPI* | [**ViewNotification**](docs/NotificationsAPI.md#viewnotification) | **Get** /api/v3/notifications/{id} | Get the notification
*NotificationsAPI* | [**ViewNotificationDetail**](docs/NotificationsAPI.md#viewnotificationdetail) | **Get** /api/v3/notifications/{notification_id}/details/{id} | Get a notification detail
*OAuth2API* | [**GetOauthApplication**](docs/OAuth2API.md#getoauthapplication) | **Get** /api/v3/oauth_applications/{id} | Get the oauth application.
*OAuth2API* | [**GetOauthClientCredentials**](docs/OAuth2API.md#getoauthclientcredentials) | **Get** /api/v3/oauth_client_credentials/{id} | Get the oauth client credentials object.
*PostsAPI* | [**ViewPost**](docs/PostsAPI.md#viewpost) | **Get** /api/v3/posts/{id} | View Post
*PreviewingAPI* | [**PreviewMarkdownDocument**](docs/PreviewingAPI.md#previewmarkdowndocument) | **Post** /api/v3/render/markdown | Preview Markdown document
*PreviewingAPI* | [**PreviewPlainDocument**](docs/PreviewingAPI.md#previewplaindocument) | **Post** /api/v3/render/plain | Preview plain document
*PrioritiesAPI* | [**ListAllPriorities**](docs/PrioritiesAPI.md#listallpriorities) | **Get** /api/v3/priorities | List all Priorities
*PrioritiesAPI* | [**ViewPriority**](docs/PrioritiesAPI.md#viewpriority) | **Get** /api/v3/priorities/{id} | View Priority
*ProjectsAPI* | [**CreateProject**](docs/ProjectsAPI.md#createproject) | **Post** /api/v3/projects | Create project
*ProjectsAPI* | [**CreateProjectCopy**](docs/ProjectsAPI.md#createprojectcopy) | **Post** /api/v3/projects/{id}/copy | Create project copy
*ProjectsAPI* | [**DeleteProject**](docs/ProjectsAPI.md#deleteproject) | **Delete** /api/v3/projects/{id} | Delete Project
*ProjectsAPI* | [**ListAvailableParentProjectCandidates**](docs/ProjectsAPI.md#listavailableparentprojectcandidates) | **Get** /api/v3/projects/available_parent_projects | List available parent project candidates
*ProjectsAPI* | [**ListProjects**](docs/ProjectsAPI.md#listprojects) | **Get** /api/v3/projects | List projects
*ProjectsAPI* | [**ListProjectsWithVersion**](docs/ProjectsAPI.md#listprojectswithversion) | **Get** /api/v3/versions/{id}/projects | List projects having version
*ProjectsAPI* | [**ProjectCopyForm**](docs/ProjectsAPI.md#projectcopyform) | **Post** /api/v3/projects/{id}/copy/form | Project copy form
*ProjectsAPI* | [**ProjectCreateForm**](docs/ProjectsAPI.md#projectcreateform) | **Post** /api/v3/projects/form | Project create form
*ProjectsAPI* | [**ProjectUpdateForm**](docs/ProjectsAPI.md#projectupdateform) | **Post** /api/v3/projects/{id}/form | Project update form
*ProjectsAPI* | [**UpdateProject**](docs/ProjectsAPI.md#updateproject) | **Patch** /api/v3/projects/{id} | Update Project
*ProjectsAPI* | [**ViewProject**](docs/ProjectsAPI.md#viewproject) | **Get** /api/v3/projects/{id} | View project
*ProjectsAPI* | [**ViewProjectSchema**](docs/ProjectsAPI.md#viewprojectschema) | **Get** /api/v3/projects/schema | View project schema
*ProjectsAPI* | [**ViewProjectStatus**](docs/ProjectsAPI.md#viewprojectstatus) | **Get** /api/v3/project_statuses/{id} | View project status
*QueriesAPI* | [**AvailableProjectsForQuery**](docs/QueriesAPI.md#availableprojectsforquery) | **Get** /api/v3/queries/available_projects | Available projects for query
*QueriesAPI* | [**CreateQuery**](docs/QueriesAPI.md#createquery) | **Post** /api/v3/queries | Create query
*QueriesAPI* | [**DeleteQuery**](docs/QueriesAPI.md#deletequery) | **Delete** /api/v3/queries/{id} | Delete query
*QueriesAPI* | [**EditQuery**](docs/QueriesAPI.md#editquery) | **Patch** /api/v3/queries/{id} | Edit Query
*QueriesAPI* | [**ListQueries**](docs/QueriesAPI.md#listqueries) | **Get** /api/v3/queries | List queries
*QueriesAPI* | [**QueryCreateForm**](docs/QueriesAPI.md#querycreateform) | **Post** /api/v3/queries/form | Query Create Form
*QueriesAPI* | [**QueryUpdateForm**](docs/QueriesAPI.md#queryupdateform) | **Post** /api/v3/queries/{id}/form | Query Update Form
*QueriesAPI* | [**StarQuery**](docs/QueriesAPI.md#starquery) | **Patch** /api/v3/queries/{id}/star | Star query
*QueriesAPI* | [**UnstarQuery**](docs/QueriesAPI.md#unstarquery) | **Patch** /api/v3/queries/{id}/unstar | Unstar query
*QueriesAPI* | [**ViewDefaultQuery**](docs/QueriesAPI.md#viewdefaultquery) | **Get** /api/v3/queries/default | View default query
*QueriesAPI* | [**ViewDefaultQueryForProject**](docs/QueriesAPI.md#viewdefaultqueryforproject) | **Get** /api/v3/projects/{id}/queries/default | View default query for project
*QueriesAPI* | [**ViewQuery**](docs/QueriesAPI.md#viewquery) | **Get** /api/v3/queries/{id} | View query
*QueriesAPI* | [**ViewSchemaForGlobalQueries**](docs/QueriesAPI.md#viewschemaforglobalqueries) | **Get** /api/v3/queries/schema | View schema for global queries
*QueriesAPI* | [**ViewSchemaForProjectQueries**](docs/QueriesAPI.md#viewschemaforprojectqueries) | **Get** /api/v3/projects/{id}/queries/schema | View schema for project queries
*QueryColumnsAPI* | [**ViewQueryColumn**](docs/QueryColumnsAPI.md#viewquerycolumn) | **Get** /api/v3/queries/columns/{id} | View Query Column
*QueryFilterInstanceSchemaAPI* | [**ListQueryFilterInstanceSchemas**](docs/QueryFilterInstanceSchemaAPI.md#listqueryfilterinstanceschemas) | **Get** /api/v3/queries/filter_instance_schemas | List Query Filter Instance Schemas
*QueryFilterInstanceSchemaAPI* | [**ListQueryFilterInstanceSchemasForProject**](docs/QueryFilterInstanceSchemaAPI.md#listqueryfilterinstanceschemasforproject) | **Get** /api/v3/projects/{id}/queries/filter_instance_schemas | List Query Filter Instance Schemas for Project
*QueryFilterInstanceSchemaAPI* | [**ViewQueryFilterInstanceSchema**](docs/QueryFilterInstanceSchemaAPI.md#viewqueryfilterinstanceschema) | **Get** /api/v3/queries/filter_instance_schemas/{id} | View Query Filter Instance Schema
*QueryFiltersAPI* | [**ViewQueryFilter**](docs/QueryFiltersAPI.md#viewqueryfilter) | **Get** /api/v3/queries/filters/{id} | View Query Filter
*QueryOperatorsAPI* | [**ViewQueryOperator**](docs/QueryOperatorsAPI.md#viewqueryoperator) | **Get** /api/v3/queries/operators/{id} | View Query Operator
*QuerySortBysAPI* | [**ViewQuerySortBy**](docs/QuerySortBysAPI.md#viewquerysortby) | **Get** /api/v3/queries/sort_bys/{id} | View Query Sort By
*RelationsAPI* | [**DeleteRelation**](docs/RelationsAPI.md#deleterelation) | **Delete** /api/v3/relations/{id} | Delete Relation
*RelationsAPI* | [**EditRelation**](docs/RelationsAPI.md#editrelation) | **Patch** /api/v3/relations/{id} | Edit Relation
*RelationsAPI* | [**ListRelations**](docs/RelationsAPI.md#listrelations) | **Get** /api/v3/relations | List Relations
*RelationsAPI* | [**RelationEditForm**](docs/RelationsAPI.md#relationeditform) | **Post** /api/v3/relations/{id}/form | Relation edit form
*RelationsAPI* | [**ViewRelation**](docs/RelationsAPI.md#viewrelation) | **Get** /api/v3/relations/{id} | View Relation
*RelationsAPI* | [**ViewRelationSchema**](docs/RelationsAPI.md#viewrelationschema) | **Get** /api/v3/relations/schema | View relation schema
*RelationsAPI* | [**ViewRelationSchemaForType**](docs/RelationsAPI.md#viewrelationschemafortype) | **Get** /api/v3/relations/schema/{type} | View relation schema for type
*RevisionsAPI* | [**ViewRevision**](docs/RevisionsAPI.md#viewrevision) | **Get** /api/v3/revisions/{id} | View revision
*RolesAPI* | [**ListRoles**](docs/RolesAPI.md#listroles) | **Get** /api/v3/roles | List roles
*RolesAPI* | [**ViewRole**](docs/RolesAPI.md#viewrole) | **Get** /api/v3/roles/{id} | View role
*RootAPI* | [**ViewRoot**](docs/RootAPI.md#viewroot) | **Get** /api/v3 | View root
*SchemasAPI* | [**ViewTheSchema**](docs/SchemasAPI.md#viewtheschema) | **Get** /api/v3/example/schema | view the schema
*StatusesAPI* | [**ListAllStatuses**](docs/StatusesAPI.md#listallstatuses) | **Get** /api/v3/statuses | List all Statuses
*StatusesAPI* | [**ViewStatus**](docs/StatusesAPI.md#viewstatus) | **Get** /api/v3/statuses/{id} | View Status
*TimeEntriesAPI* | [**AvailableProjectsForTimeEntries**](docs/TimeEntriesAPI.md#availableprojectsfortimeentries) | **Get** /api/v3/time_entries/available_projects | Available projects for time entries
*TimeEntriesAPI* | [**CreateTimeEntry**](docs/TimeEntriesAPI.md#createtimeentry) | **Post** /api/v3/time_entries | Create time entry
*TimeEntriesAPI* | [**DeleteTimeEntry**](docs/TimeEntriesAPI.md#deletetimeentry) | **Delete** /api/v3/time_entries/{id} | Delete time entry
*TimeEntriesAPI* | [**GetTimeEntry**](docs/TimeEntriesAPI.md#gettimeentry) | **Get** /api/v3/time_entries/{id} | Get time entry
*TimeEntriesAPI* | [**ListTimeEntries**](docs/TimeEntriesAPI.md#listtimeentries) | **Get** /api/v3/time_entries | List time entries
*TimeEntriesAPI* | [**TimeEntryCreateForm**](docs/TimeEntriesAPI.md#timeentrycreateform) | **Post** /api/v3/time_entries/form | Time entry create form
*TimeEntriesAPI* | [**TimeEntryUpdateForm**](docs/TimeEntriesAPI.md#timeentryupdateform) | **Post** /api/v3/time_entries/{id}/form | Time entry update form
*TimeEntriesAPI* | [**UpdateTimeEntry**](docs/TimeEntriesAPI.md#updatetimeentry) | **Patch** /api/v3/time_entries/{id} | update time entry
*TimeEntriesAPI* | [**ViewTimeEntrySchema**](docs/TimeEntriesAPI.md#viewtimeentryschema) | **Get** /api/v3/time_entries/schema | View time entry schema
*TimeEntryActivitiesAPI* | [**GetTimeEntriesActivity**](docs/TimeEntryActivitiesAPI.md#gettimeentriesactivity) | **Get** /api/v3/time_entries/activity/{id} | View time entries activity
*TypesAPI* | [**ListAllTypes**](docs/TypesAPI.md#listalltypes) | **Get** /api/v3/types | List all Types
*TypesAPI* | [**ListTypesAvailableInAProject**](docs/TypesAPI.md#listtypesavailableinaproject) | **Get** /api/v3/projects/{id}/types | List types available in a project
*TypesAPI* | [**ViewType**](docs/TypesAPI.md#viewtype) | **Get** /api/v3/types/{id} | View Type
*UserPreferencesAPI* | [**ShowMyPreferences**](docs/UserPreferencesAPI.md#showmypreferences) | **Get** /api/v3/my_preferences | Show my preferences
*UserPreferencesAPI* | [**UpdateUserPreferences**](docs/UserPreferencesAPI.md#updateuserpreferences) | **Patch** /api/v3/my_preferences | Update my preferences
*UsersAPI* | [**CreateUser**](docs/UsersAPI.md#createuser) | **Post** /api/v3/users | Create User
*UsersAPI* | [**DeleteUser**](docs/UsersAPI.md#deleteuser) | **Delete** /api/v3/users/{id} | Delete user
*UsersAPI* | [**ListUsers**](docs/UsersAPI.md#listusers) | **Get** /api/v3/users | List Users
*UsersAPI* | [**LockUser**](docs/UsersAPI.md#lockuser) | **Post** /api/v3/users/{id}/lock | Lock user
*UsersAPI* | [**UnlockUser**](docs/UsersAPI.md#unlockuser) | **Delete** /api/v3/users/{id}/lock | Unlock user
*UsersAPI* | [**UpdateUser**](docs/UsersAPI.md#updateuser) | **Patch** /api/v3/users/{id} | Update user
*UsersAPI* | [**UserUpdateForm**](docs/UsersAPI.md#userupdateform) | **Post** /api/v3/users/{id}/form | User update form
*UsersAPI* | [**ViewUser**](docs/UsersAPI.md#viewuser) | **Get** /api/v3/users/{id} | View user
*UsersAPI* | [**ViewUserSchema**](docs/UsersAPI.md#viewuserschema) | **Get** /api/v3/users/schema | View user schema
*VersionsAPI* | [**AvailableProjectsForVersions**](docs/VersionsAPI.md#availableprojectsforversions) | **Get** /api/v3/versions/available_projects | Available projects for versions
*VersionsAPI* | [**CreateVersion**](docs/VersionsAPI.md#createversion) | **Post** /api/v3/versions | Create version
*VersionsAPI* | [**DeleteVersion**](docs/VersionsAPI.md#deleteversion) | **Delete** /api/v3/versions/{id} | Delete version
*VersionsAPI* | [**ListVersions**](docs/VersionsAPI.md#listversions) | **Get** /api/v3/versions | List versions
*VersionsAPI* | [**ListVersionsAvailableInAProject**](docs/VersionsAPI.md#listversionsavailableinaproject) | **Get** /api/v3/projects/{id}/versions | List versions available in a project
*VersionsAPI* | [**UpdateVersion**](docs/VersionsAPI.md#updateversion) | **Patch** /api/v3/versions/{id} | Update Version
*VersionsAPI* | [**VersionCreateForm**](docs/VersionsAPI.md#versioncreateform) | **Post** /api/v3/versions/form | Version create form
*VersionsAPI* | [**VersionUpdateForm**](docs/VersionsAPI.md#versionupdateform) | **Post** /api/v3/versions/{id}/form | Version update form
*VersionsAPI* | [**ViewVersion**](docs/VersionsAPI.md#viewversion) | **Get** /api/v3/versions/{id} | View version
*VersionsAPI* | [**ViewVersionSchema**](docs/VersionsAPI.md#viewversionschema) | **Get** /api/v3/versions/schema | View version schema
*ViewsAPI* | [**CreateViews**](docs/ViewsAPI.md#createviews) | **Post** /api/v3/views/{id} | Create view
*ViewsAPI* | [**ListViews**](docs/ViewsAPI.md#listviews) | **Get** /api/v3/views | List views
*ViewsAPI* | [**ViewView**](docs/ViewsAPI.md#viewview) | **Get** /api/v3/views/{id} | View view
*WikiPagesAPI* | [**ViewWikiPage**](docs/WikiPagesAPI.md#viewwikipage) | **Get** /api/v3/wiki_pages/{id} | View Wiki Page
*WorkPackagesAPI* | [**AddWatcher**](docs/WorkPackagesAPI.md#addwatcher) | **Post** /api/v3/work_packages/{id}/watchers | Add watcher
*WorkPackagesAPI* | [**AvailableProjectsForWorkPackage**](docs/WorkPackagesAPI.md#availableprojectsforworkpackage) | **Get** /api/v3/work_packages/{id}/available_projects | Available projects for work package
*WorkPackagesAPI* | [**AvailableWatchers**](docs/WorkPackagesAPI.md#availablewatchers) | **Get** /api/v3/work_packages/{id}/available_watchers | Available watchers
*WorkPackagesAPI* | [**CommentWorkPackage**](docs/WorkPackagesAPI.md#commentworkpackage) | **Post** /api/v3/work_packages/{id}/activities | Comment work package
*WorkPackagesAPI* | [**CreateProjectWorkPackage**](docs/WorkPackagesAPI.md#createprojectworkpackage) | **Post** /api/v3/projects/{id}/work_packages | Create work package in project
*WorkPackagesAPI* | [**CreateRelation**](docs/WorkPackagesAPI.md#createrelation) | **Post** /api/v3/work_packages/{id}/relations | Create Relation
*WorkPackagesAPI* | [**CreateWorkPackage**](docs/WorkPackagesAPI.md#createworkpackage) | **Post** /api/v3/work_packages | Create Work Package
*WorkPackagesAPI* | [**CreateWorkPackageFileLink**](docs/WorkPackagesAPI.md#createworkpackagefilelink) | **Post** /api/v3/work_packages/{id}/file_links | Creates file links.
*WorkPackagesAPI* | [**DeleteWorkPackage**](docs/WorkPackagesAPI.md#deleteworkpackage) | **Delete** /api/v3/work_packages/{id} | Delete Work Package
*WorkPackagesAPI* | [**GetProjectWorkPackageCollection**](docs/WorkPackagesAPI.md#getprojectworkpackagecollection) | **Get** /api/v3/projects/{id}/work_packages | Get work packages of project
*WorkPackagesAPI* | [**ListAvailableRelationCandidates**](docs/WorkPackagesAPI.md#listavailablerelationcandidates) | **Get** /api/v3/work_packages/{id}/available_relation_candidates | Available relation candidates
*WorkPackagesAPI* | [**ListWatchers**](docs/WorkPackagesAPI.md#listwatchers) | **Get** /api/v3/work_packages/{id}/watchers | List watchers
*WorkPackagesAPI* | [**ListWorkPackageActivities**](docs/WorkPackagesAPI.md#listworkpackageactivities) | **Get** /api/v3/work_packages/{id}/activities | List work package activities
*WorkPackagesAPI* | [**ListWorkPackageFileLinks**](docs/WorkPackagesAPI.md#listworkpackagefilelinks) | **Get** /api/v3/work_packages/{id}/file_links | Gets all file links of a work package
*WorkPackagesAPI* | [**ListWorkPackageRelations**](docs/WorkPackagesAPI.md#listworkpackagerelations) | **Get** /api/v3/work_packages/{id}/relations | List relations
*WorkPackagesAPI* | [**ListWorkPackageSchemas**](docs/WorkPackagesAPI.md#listworkpackageschemas) | **Get** /api/v3/work_packages/schemas | List Work Package Schemas
*WorkPackagesAPI* | [**ListWorkPackages**](docs/WorkPackagesAPI.md#listworkpackages) | **Get** /api/v3/work_packages | List work packages
*WorkPackagesAPI* | [**ProjectAvailableAssignees**](docs/WorkPackagesAPI.md#projectavailableassignees) | **Get** /api/v3/projects/{id}/available_assignees | Project Available assignees
*WorkPackagesAPI* | [**RemoveWatcher**](docs/WorkPackagesAPI.md#removewatcher) | **Delete** /api/v3/work_packages/{id}/watchers/{user_id} | Remove watcher
*WorkPackagesAPI* | [**Revisions**](docs/WorkPackagesAPI.md#revisions) | **Get** /api/v3/work_packages/{id}/revisions | Revisions
*WorkPackagesAPI* | [**UpdateWorkPackage**](docs/WorkPackagesAPI.md#updateworkpackage) | **Patch** /api/v3/work_packages/{id} | Update a Work Package
*WorkPackagesAPI* | [**ViewWorkPackage**](docs/WorkPackagesAPI.md#viewworkpackage) | **Get** /api/v3/work_packages/{id} | View Work Package
*WorkPackagesAPI* | [**ViewWorkPackageSchema**](docs/WorkPackagesAPI.md#viewworkpackageschema) | **Get** /api/v3/work_packages/schemas/{identifier} | View Work Package Schema
*WorkPackagesAPI* | [**WorkPackageAvailableAssignees**](docs/WorkPackagesAPI.md#workpackageavailableassignees) | **Get** /api/v3/work_packages/{id}/available_assignees | Work Package Available assignees
*WorkPackagesAPI* | [**WorkPackageCreateForm**](docs/WorkPackagesAPI.md#workpackagecreateform) | **Post** /api/v3/work_packages/form | Work Package Create Form
*WorkPackagesAPI* | [**WorkPackageCreateFormForProject**](docs/WorkPackagesAPI.md#workpackagecreateformforproject) | **Post** /api/v3/projects/{id}/work_packages/form | Work Package Create Form For Project
*WorkPackagesAPI* | [**WorkPackageEditForm**](docs/WorkPackagesAPI.md#workpackageeditform) | **Post** /api/v3/work_packages/{id}/form | Work Package Edit Form
*WorkScheduleAPI* | [**CreateNonWorkingDay**](docs/WorkScheduleAPI.md#createnonworkingday) | **Post** /api/v3/days/non_working | Creates a non-working day (NOT IMPLEMENTED)
*WorkScheduleAPI* | [**DeleteNonWorkingDay**](docs/WorkScheduleAPI.md#deletenonworkingday) | **Delete** /api/v3/days/non_working/{date} | Removes a non-working day (NOT IMPLEMENTED)
*WorkScheduleAPI* | [**ListDays**](docs/WorkScheduleAPI.md#listdays) | **Get** /api/v3/days | Lists days
*WorkScheduleAPI* | [**ListNonWorkingDays**](docs/WorkScheduleAPI.md#listnonworkingdays) | **Get** /api/v3/days/non_working | Lists all non working days
*WorkScheduleAPI* | [**ListWeekDays**](docs/WorkScheduleAPI.md#listweekdays) | **Get** /api/v3/days/week | Lists week days
*WorkScheduleAPI* | [**UpdateNonWorkingDay**](docs/WorkScheduleAPI.md#updatenonworkingday) | **Patch** /api/v3/days/non_working/{date} | Update a non-working day attributes (NOT IMPLEMENTED)
*WorkScheduleAPI* | [**UpdateWeekDay**](docs/WorkScheduleAPI.md#updateweekday) | **Patch** /api/v3/days/week/{day} | Update a week day attributes (NOT IMPLEMENTED)
*WorkScheduleAPI* | [**UpdateWeekDays**](docs/WorkScheduleAPI.md#updateweekdays) | **Patch** /api/v3/days/week | Update week days (NOT IMPLEMENTED)
*WorkScheduleAPI* | [**ViewDay**](docs/WorkScheduleAPI.md#viewday) | **Get** /api/v3/days/{date} | View day
*WorkScheduleAPI* | [**ViewNonWorkingDay**](docs/WorkScheduleAPI.md#viewnonworkingday) | **Get** /api/v3/days/non_working/{date} | View a non-working day
*WorkScheduleAPI* | [**ViewWeekDay**](docs/WorkScheduleAPI.md#viewweekday) | **Get** /api/v3/days/week/{day} | View a week day


## Documentation For Models

 - [ActivityModel](docs/ActivityModel.md)
 - [ActivityModelComment](docs/ActivityModelComment.md)
 - [AddWatcherRequest](docs/AddWatcherRequest.md)
 - [AttachmentModel](docs/AttachmentModel.md)
 - [AttachmentModelDescription](docs/AttachmentModelDescription.md)
 - [AttachmentModelDigest](docs/AttachmentModelDigest.md)
 - [AttachmentModelLinks](docs/AttachmentModelLinks.md)
 - [AttachmentModelLinksAuthor](docs/AttachmentModelLinksAuthor.md)
 - [AttachmentModelLinksContainer](docs/AttachmentModelLinksContainer.md)
 - [AttachmentModelLinksDelete](docs/AttachmentModelLinksDelete.md)
 - [AttachmentModelLinksDownloadLocation](docs/AttachmentModelLinksDownloadLocation.md)
 - [AttachmentModelLinksSelf](docs/AttachmentModelLinksSelf.md)
 - [AttachmentsModel](docs/AttachmentsModel.md)
 - [AttachmentsModelAllOfEmbedded](docs/AttachmentsModelAllOfEmbedded.md)
 - [AttachmentsModelAllOfEmbeddedElements](docs/AttachmentsModelAllOfEmbeddedElements.md)
 - [AttachmentsModelAllOfLinks](docs/AttachmentsModelAllOfLinks.md)
 - [AttachmentsModelAllOfLinksSelf](docs/AttachmentsModelAllOfLinksSelf.md)
 - [AvailableAssigneesModel](docs/AvailableAssigneesModel.md)
 - [AvailableAssigneesModelAllOfEmbedded](docs/AvailableAssigneesModelAllOfEmbedded.md)
 - [AvailableAssigneesModelAllOfEmbeddedElements](docs/AvailableAssigneesModelAllOfEmbeddedElements.md)
 - [AvailableAssigneesModelAllOfLinks](docs/AvailableAssigneesModelAllOfLinks.md)
 - [AvailableAssigneesModelAllOfLinksSelf](docs/AvailableAssigneesModelAllOfLinksSelf.md)
 - [BudgetModel](docs/BudgetModel.md)
 - [BudgetModelLinks](docs/BudgetModelLinks.md)
 - [BudgetModelLinksSelf](docs/BudgetModelLinksSelf.md)
 - [CategoriesByProjectModel](docs/CategoriesByProjectModel.md)
 - [CategoriesByProjectModelAllOfEmbedded](docs/CategoriesByProjectModelAllOfEmbedded.md)
 - [CategoriesByProjectModelAllOfEmbeddedElements](docs/CategoriesByProjectModelAllOfEmbeddedElements.md)
 - [CategoriesByProjectModelAllOfLinks](docs/CategoriesByProjectModelAllOfLinks.md)
 - [CategoriesByProjectModelAllOfLinksSelf](docs/CategoriesByProjectModelAllOfLinksSelf.md)
 - [CategoryModel](docs/CategoryModel.md)
 - [CategoryModelLinks](docs/CategoryModelLinks.md)
 - [CategoryModelLinksDefaultAssignee](docs/CategoryModelLinksDefaultAssignee.md)
 - [CategoryModelLinksProject](docs/CategoryModelLinksProject.md)
 - [CategoryModelLinksSelf](docs/CategoryModelLinksSelf.md)
 - [CollectionModel](docs/CollectionModel.md)
 - [CollectionModelLinks](docs/CollectionModelLinks.md)
 - [CollectionModelLinksSelf](docs/CollectionModelLinksSelf.md)
 - [CommentWorkPackageRequest](docs/CommentWorkPackageRequest.md)
 - [CommentWorkPackageRequestComment](docs/CommentWorkPackageRequestComment.md)
 - [ConfigurationModel](docs/ConfigurationModel.md)
 - [CreateAttachmentRequestMetadata](docs/CreateAttachmentRequestMetadata.md)
 - [CreateRelationRequest](docs/CreateRelationRequest.md)
 - [CreateRelationRequestLinks](docs/CreateRelationRequestLinks.md)
 - [CreateViewsRequest](docs/CreateViewsRequest.md)
 - [CreateViewsRequestLinks](docs/CreateViewsRequestLinks.md)
 - [CreateViewsRequestLinksQuery](docs/CreateViewsRequestLinksQuery.md)
 - [CustomActionModel](docs/CustomActionModel.md)
 - [CustomActionModelLinks](docs/CustomActionModelLinks.md)
 - [CustomActionModelLinksExecuteImmediately](docs/CustomActionModelLinksExecuteImmediately.md)
 - [CustomActionModelLinksSelf](docs/CustomActionModelLinksSelf.md)
 - [CustomOptionModel](docs/CustomOptionModel.md)
 - [CustomOptionModelLinks](docs/CustomOptionModelLinks.md)
 - [CustomOptionModelLinksSelf](docs/CustomOptionModelLinksSelf.md)
 - [DayCollectionModel](docs/DayCollectionModel.md)
 - [DayCollectionModelAllOfEmbedded](docs/DayCollectionModelAllOfEmbedded.md)
 - [DayCollectionModelAllOfLinks](docs/DayCollectionModelAllOfLinks.md)
 - [DayCollectionModelAllOfLinksSelf](docs/DayCollectionModelAllOfLinksSelf.md)
 - [DayModel](docs/DayModel.md)
 - [DayModelLinks](docs/DayModelLinks.md)
 - [DayModelLinksWeekDay](docs/DayModelLinksWeekDay.md)
 - [DocumentModel](docs/DocumentModel.md)
 - [DocumentModelLinks](docs/DocumentModelLinks.md)
 - [DocumentModelLinksAttachments](docs/DocumentModelLinksAttachments.md)
 - [DocumentModelLinksProject](docs/DocumentModelLinksProject.md)
 - [DocumentModelLinksSelf](docs/DocumentModelLinksSelf.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [ErrorResponseEmbedded](docs/ErrorResponseEmbedded.md)
 - [ErrorResponseEmbeddedDetails](docs/ErrorResponseEmbeddedDetails.md)
 - [ExecuteCustomActionRequest](docs/ExecuteCustomActionRequest.md)
 - [ExecuteCustomActionRequestLinks](docs/ExecuteCustomActionRequestLinks.md)
 - [ExecuteCustomActionRequestLinksWorkPackage](docs/ExecuteCustomActionRequestLinksWorkPackage.md)
 - [FileLinkCollectionReadModel](docs/FileLinkCollectionReadModel.md)
 - [FileLinkCollectionReadModelAllOfEmbedded](docs/FileLinkCollectionReadModelAllOfEmbedded.md)
 - [FileLinkCollectionReadModelAllOfLinks](docs/FileLinkCollectionReadModelAllOfLinks.md)
 - [FileLinkCollectionReadModelAllOfLinksSelf](docs/FileLinkCollectionReadModelAllOfLinksSelf.md)
 - [FileLinkCollectionWriteModel](docs/FileLinkCollectionWriteModel.md)
 - [FileLinkCollectionWriteModelAllOfEmbedded](docs/FileLinkCollectionWriteModelAllOfEmbedded.md)
 - [FileLinkOriginDataModel](docs/FileLinkOriginDataModel.md)
 - [FileLinkReadModel](docs/FileLinkReadModel.md)
 - [FileLinkReadModelEmbedded](docs/FileLinkReadModelEmbedded.md)
 - [FileLinkReadModelLinks](docs/FileLinkReadModelLinks.md)
 - [FileLinkReadModelLinksContainer](docs/FileLinkReadModelLinksContainer.md)
 - [FileLinkReadModelLinksCreator](docs/FileLinkReadModelLinksCreator.md)
 - [FileLinkReadModelLinksDelete](docs/FileLinkReadModelLinksDelete.md)
 - [FileLinkReadModelLinksOriginOpen](docs/FileLinkReadModelLinksOriginOpen.md)
 - [FileLinkReadModelLinksOriginOpenLocation](docs/FileLinkReadModelLinksOriginOpenLocation.md)
 - [FileLinkReadModelLinksSelf](docs/FileLinkReadModelLinksSelf.md)
 - [FileLinkReadModelLinksStaticOriginDownload](docs/FileLinkReadModelLinksStaticOriginDownload.md)
 - [FileLinkReadModelLinksStaticOriginOpen](docs/FileLinkReadModelLinksStaticOriginOpen.md)
 - [FileLinkReadModelLinksStaticOriginOpenLocation](docs/FileLinkReadModelLinksStaticOriginOpenLocation.md)
 - [FileLinkReadModelLinksStatus](docs/FileLinkReadModelLinksStatus.md)
 - [FileLinkReadModelLinksStorage](docs/FileLinkReadModelLinksStorage.md)
 - [FileLinkWriteModel](docs/FileLinkWriteModel.md)
 - [FileLinkWriteModelLinks](docs/FileLinkWriteModelLinks.md)
 - [FileLinkWriteModelLinksOneOf](docs/FileLinkWriteModelLinksOneOf.md)
 - [FileLinkWriteModelLinksOneOf1](docs/FileLinkWriteModelLinksOneOf1.md)
 - [FileLinkWriteModelLinksOneOf1StorageUrl](docs/FileLinkWriteModelLinksOneOf1StorageUrl.md)
 - [Formattable](docs/Formattable.md)
 - [GridCollectionModel](docs/GridCollectionModel.md)
 - [GridCollectionModelAllOfEmbedded](docs/GridCollectionModelAllOfEmbedded.md)
 - [GridReadModel](docs/GridReadModel.md)
 - [GridReadModelLinks](docs/GridReadModelLinks.md)
 - [GridReadModelLinksAddAttachment](docs/GridReadModelLinksAddAttachment.md)
 - [GridReadModelLinksAttachments](docs/GridReadModelLinksAttachments.md)
 - [GridReadModelLinksDelete](docs/GridReadModelLinksDelete.md)
 - [GridReadModelLinksScope](docs/GridReadModelLinksScope.md)
 - [GridReadModelLinksSelf](docs/GridReadModelLinksSelf.md)
 - [GridReadModelLinksUpdate](docs/GridReadModelLinksUpdate.md)
 - [GridReadModelLinksUpdateImmediately](docs/GridReadModelLinksUpdateImmediately.md)
 - [GridWidgetModel](docs/GridWidgetModel.md)
 - [GridWriteModel](docs/GridWriteModel.md)
 - [GridWriteModelLinks](docs/GridWriteModelLinks.md)
 - [GroupCollectionModel](docs/GroupCollectionModel.md)
 - [GroupCollectionModelAllOfEmbedded](docs/GroupCollectionModelAllOfEmbedded.md)
 - [GroupCollectionModelAllOfLinks](docs/GroupCollectionModelAllOfLinks.md)
 - [GroupCollectionModelAllOfLinksSelf](docs/GroupCollectionModelAllOfLinksSelf.md)
 - [GroupModel](docs/GroupModel.md)
 - [GroupModelAllOfEmbedded](docs/GroupModelAllOfEmbedded.md)
 - [GroupModelAllOfLinks](docs/GroupModelAllOfLinks.md)
 - [GroupModelAllOfLinksDelete](docs/GroupModelAllOfLinksDelete.md)
 - [GroupModelAllOfLinksMembers](docs/GroupModelAllOfLinksMembers.md)
 - [GroupModelAllOfLinksUpdateImmediately](docs/GroupModelAllOfLinksUpdateImmediately.md)
 - [GroupWriteModel](docs/GroupWriteModel.md)
 - [GroupWriteModelLinks](docs/GroupWriteModelLinks.md)
 - [GroupWriteModelLinksMembersInner](docs/GroupWriteModelLinksMembersInner.md)
 - [HelpTextCollectionModel](docs/HelpTextCollectionModel.md)
 - [HelpTextCollectionModelAllOfEmbedded](docs/HelpTextCollectionModelAllOfEmbedded.md)
 - [HelpTextCollectionModelAllOfLinks](docs/HelpTextCollectionModelAllOfLinks.md)
 - [HelpTextCollectionModelAllOfLinksSelf](docs/HelpTextCollectionModelAllOfLinksSelf.md)
 - [HelpTextModel](docs/HelpTextModel.md)
 - [HelpTextModelLinks](docs/HelpTextModelLinks.md)
 - [HelpTextModelLinksAddAttachment](docs/HelpTextModelLinksAddAttachment.md)
 - [HelpTextModelLinksAttachments](docs/HelpTextModelLinksAttachments.md)
 - [HelpTextModelLinksEditText](docs/HelpTextModelLinksEditText.md)
 - [HelpTextModelLinksSelf](docs/HelpTextModelLinksSelf.md)
 - [Link](docs/Link.md)
 - [ListAvailableParentProjectCandidatesModel](docs/ListAvailableParentProjectCandidatesModel.md)
 - [ListAvailableParentProjectCandidatesModelAllOfEmbedded](docs/ListAvailableParentProjectCandidatesModelAllOfEmbedded.md)
 - [ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements](docs/ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements.md)
 - [ListAvailableParentProjectCandidatesModelAllOfLinks](docs/ListAvailableParentProjectCandidatesModelAllOfLinks.md)
 - [ListAvailableParentProjectCandidatesModelAllOfLinksSelf](docs/ListAvailableParentProjectCandidatesModelAllOfLinksSelf.md)
 - [MeetingModel](docs/MeetingModel.md)
 - [MeetingModelLinks](docs/MeetingModelLinks.md)
 - [MeetingModelLinksAddAttachment](docs/MeetingModelLinksAddAttachment.md)
 - [MeetingModelLinksAttachments](docs/MeetingModelLinksAttachments.md)
 - [MeetingModelLinksAuthor](docs/MeetingModelLinksAuthor.md)
 - [MeetingModelLinksProject](docs/MeetingModelLinksProject.md)
 - [MeetingModelLinksSelf](docs/MeetingModelLinksSelf.md)
 - [MembershipCollectionModel](docs/MembershipCollectionModel.md)
 - [MembershipCollectionModelAllOfEmbedded](docs/MembershipCollectionModelAllOfEmbedded.md)
 - [MembershipFormModel](docs/MembershipFormModel.md)
 - [MembershipFormModelEmbedded](docs/MembershipFormModelEmbedded.md)
 - [MembershipFormModelEmbeddedValidationError](docs/MembershipFormModelEmbeddedValidationError.md)
 - [MembershipFormModelLinks](docs/MembershipFormModelLinks.md)
 - [MembershipFormModelLinksCommit](docs/MembershipFormModelLinksCommit.md)
 - [MembershipFormModelLinksSelf](docs/MembershipFormModelLinksSelf.md)
 - [MembershipFormModelLinksValidateInner](docs/MembershipFormModelLinksValidateInner.md)
 - [MembershipReadModel](docs/MembershipReadModel.md)
 - [MembershipReadModelEmbedded](docs/MembershipReadModelEmbedded.md)
 - [MembershipReadModelEmbeddedPrincipal](docs/MembershipReadModelEmbeddedPrincipal.md)
 - [MembershipReadModelLinks](docs/MembershipReadModelLinks.md)
 - [MembershipReadModelLinksPrincipal](docs/MembershipReadModelLinksPrincipal.md)
 - [MembershipReadModelLinksProject](docs/MembershipReadModelLinksProject.md)
 - [MembershipReadModelLinksRolesInner](docs/MembershipReadModelLinksRolesInner.md)
 - [MembershipReadModelLinksSchema](docs/MembershipReadModelLinksSchema.md)
 - [MembershipReadModelLinksSelf](docs/MembershipReadModelLinksSelf.md)
 - [MembershipReadModelLinksUpdate](docs/MembershipReadModelLinksUpdate.md)
 - [MembershipReadModelLinksUpdateImmediately](docs/MembershipReadModelLinksUpdateImmediately.md)
 - [MembershipSchemaModel](docs/MembershipSchemaModel.md)
 - [MembershipWriteModel](docs/MembershipWriteModel.md)
 - [MembershipWriteModelLinks](docs/MembershipWriteModelLinks.md)
 - [MembershipWriteModelLinksPrincipal](docs/MembershipWriteModelLinksPrincipal.md)
 - [MembershipWriteModelLinksProject](docs/MembershipWriteModelLinksProject.md)
 - [MembershipWriteModelMeta](docs/MembershipWriteModelMeta.md)
 - [MembershipWriteModelMetaNotificationMessage](docs/MembershipWriteModelMetaNotificationMessage.md)
 - [NewsModel](docs/NewsModel.md)
 - [NewsModelLinks](docs/NewsModelLinks.md)
 - [NewsModelLinksAuthor](docs/NewsModelLinksAuthor.md)
 - [NewsModelLinksProject](docs/NewsModelLinksProject.md)
 - [NewsModelLinksSelf](docs/NewsModelLinksSelf.md)
 - [NonWorkingDayCollectionModel](docs/NonWorkingDayCollectionModel.md)
 - [NonWorkingDayCollectionModelAllOfEmbedded](docs/NonWorkingDayCollectionModelAllOfEmbedded.md)
 - [NonWorkingDayCollectionModelAllOfLinks](docs/NonWorkingDayCollectionModelAllOfLinks.md)
 - [NonWorkingDayCollectionModelAllOfLinksSelf](docs/NonWorkingDayCollectionModelAllOfLinksSelf.md)
 - [NonWorkingDayModel](docs/NonWorkingDayModel.md)
 - [NonWorkingDayModelLinks](docs/NonWorkingDayModelLinks.md)
 - [NonWorkingDayModelLinksSelf](docs/NonWorkingDayModelLinksSelf.md)
 - [NotificationCollectionModel](docs/NotificationCollectionModel.md)
 - [NotificationCollectionModelAllOfEmbedded](docs/NotificationCollectionModelAllOfEmbedded.md)
 - [NotificationCollectionModelAllOfLinks](docs/NotificationCollectionModelAllOfLinks.md)
 - [NotificationCollectionModelAllOfLinksChangeSize](docs/NotificationCollectionModelAllOfLinksChangeSize.md)
 - [NotificationCollectionModelAllOfLinksJumpTo](docs/NotificationCollectionModelAllOfLinksJumpTo.md)
 - [NotificationCollectionModelAllOfLinksSelf](docs/NotificationCollectionModelAllOfLinksSelf.md)
 - [NotificationModel](docs/NotificationModel.md)
 - [NotificationModelDetailsInner](docs/NotificationModelDetailsInner.md)
 - [NotificationModelEmbedded](docs/NotificationModelEmbedded.md)
 - [NotificationModelEmbeddedResource](docs/NotificationModelEmbeddedResource.md)
 - [NotificationModelLinks](docs/NotificationModelLinks.md)
 - [NotificationModelLinksActivity](docs/NotificationModelLinksActivity.md)
 - [NotificationModelLinksActor](docs/NotificationModelLinksActor.md)
 - [NotificationModelLinksDetailsInner](docs/NotificationModelLinksDetailsInner.md)
 - [NotificationModelLinksProject](docs/NotificationModelLinksProject.md)
 - [NotificationModelLinksReadIAN](docs/NotificationModelLinksReadIAN.md)
 - [NotificationModelLinksResource](docs/NotificationModelLinksResource.md)
 - [NotificationModelLinksSelf](docs/NotificationModelLinksSelf.md)
 - [NotificationModelLinksUnreadIAN](docs/NotificationModelLinksUnreadIAN.md)
 - [OAuthApplicationReadModel](docs/OAuthApplicationReadModel.md)
 - [OAuthApplicationReadModelLinks](docs/OAuthApplicationReadModelLinks.md)
 - [OAuthApplicationReadModelLinksIntegration](docs/OAuthApplicationReadModelLinksIntegration.md)
 - [OAuthApplicationReadModelLinksOwner](docs/OAuthApplicationReadModelLinksOwner.md)
 - [OAuthApplicationReadModelLinksRedirectUri](docs/OAuthApplicationReadModelLinksRedirectUri.md)
 - [OAuthApplicationReadModelLinksSelf](docs/OAuthApplicationReadModelLinksSelf.md)
 - [OAuthClientCredentialsReadModel](docs/OAuthClientCredentialsReadModel.md)
 - [OAuthClientCredentialsReadModelLinks](docs/OAuthClientCredentialsReadModelLinks.md)
 - [OAuthClientCredentialsReadModelLinksIntegration](docs/OAuthClientCredentialsReadModelLinksIntegration.md)
 - [OAuthClientCredentialsReadModelLinksSelf](docs/OAuthClientCredentialsReadModelLinksSelf.md)
 - [OAuthClientCredentialsWriteModel](docs/OAuthClientCredentialsWriteModel.md)
 - [PaginatedCollectionModel](docs/PaginatedCollectionModel.md)
 - [PaginatedCollectionModelAllOfLinks](docs/PaginatedCollectionModelAllOfLinks.md)
 - [PaginatedCollectionModelAllOfLinksChangeSize](docs/PaginatedCollectionModelAllOfLinksChangeSize.md)
 - [PaginatedCollectionModelAllOfLinksJumpTo](docs/PaginatedCollectionModelAllOfLinksJumpTo.md)
 - [PlaceholderUserCollectionModel](docs/PlaceholderUserCollectionModel.md)
 - [PlaceholderUserCollectionModelAllOfEmbedded](docs/PlaceholderUserCollectionModelAllOfEmbedded.md)
 - [PlaceholderUserCollectionModelAllOfLinks](docs/PlaceholderUserCollectionModelAllOfLinks.md)
 - [PlaceholderUserCollectionModelAllOfLinksSelf](docs/PlaceholderUserCollectionModelAllOfLinksSelf.md)
 - [PlaceholderUserCreateModel](docs/PlaceholderUserCreateModel.md)
 - [PlaceholderUserModel](docs/PlaceholderUserModel.md)
 - [PlaceholderUserModelAllOfLinks](docs/PlaceholderUserModelAllOfLinks.md)
 - [PlaceholderUserModelAllOfLinksDelete](docs/PlaceholderUserModelAllOfLinksDelete.md)
 - [PlaceholderUserModelAllOfLinksShowUser](docs/PlaceholderUserModelAllOfLinksShowUser.md)
 - [PlaceholderUserModelAllOfLinksUpdateImmediately](docs/PlaceholderUserModelAllOfLinksUpdateImmediately.md)
 - [PostModel](docs/PostModel.md)
 - [PostModelLinks](docs/PostModelLinks.md)
 - [PostModelLinksAddAttachment](docs/PostModelLinksAddAttachment.md)
 - [PrincipalCollectionModel](docs/PrincipalCollectionModel.md)
 - [PrincipalCollectionModelAllOfEmbedded](docs/PrincipalCollectionModelAllOfEmbedded.md)
 - [PrincipalCollectionModelAllOfEmbeddedElements](docs/PrincipalCollectionModelAllOfEmbeddedElements.md)
 - [PrincipalCollectionModelAllOfLinks](docs/PrincipalCollectionModelAllOfLinks.md)
 - [PrincipalCollectionModelAllOfLinksSelf](docs/PrincipalCollectionModelAllOfLinksSelf.md)
 - [PrincipalModel](docs/PrincipalModel.md)
 - [PrincipalModelLinks](docs/PrincipalModelLinks.md)
 - [PrincipalModelLinksMemberships](docs/PrincipalModelLinksMemberships.md)
 - [PrincipalModelLinksSelf](docs/PrincipalModelLinksSelf.md)
 - [PriorityCollectionModel](docs/PriorityCollectionModel.md)
 - [PriorityCollectionModelAllOfEmbedded](docs/PriorityCollectionModelAllOfEmbedded.md)
 - [PriorityCollectionModelAllOfLinks](docs/PriorityCollectionModelAllOfLinks.md)
 - [PriorityCollectionModelAllOfLinksSelf](docs/PriorityCollectionModelAllOfLinksSelf.md)
 - [PriorityCollectionModelAllOfLinksSelfAllOfSelf](docs/PriorityCollectionModelAllOfLinksSelfAllOfSelf.md)
 - [PriorityModel](docs/PriorityModel.md)
 - [PriorityModelLinks](docs/PriorityModelLinks.md)
 - [PriorityModelLinksSelf](docs/PriorityModelLinksSelf.md)
 - [ProjectCollectionModel](docs/ProjectCollectionModel.md)
 - [ProjectCollectionModelAllOfEmbedded](docs/ProjectCollectionModelAllOfEmbedded.md)
 - [ProjectCollectionModelAllOfLinks](docs/ProjectCollectionModelAllOfLinks.md)
 - [ProjectCollectionModelAllOfLinksRepresentations](docs/ProjectCollectionModelAllOfLinksRepresentations.md)
 - [ProjectCollectionModelAllOfLinksSelf](docs/ProjectCollectionModelAllOfLinksSelf.md)
 - [ProjectModel](docs/ProjectModel.md)
 - [ProjectModelLinks](docs/ProjectModelLinks.md)
 - [ProjectModelLinksAncestorsInner](docs/ProjectModelLinksAncestorsInner.md)
 - [ProjectModelLinksCategories](docs/ProjectModelLinksCategories.md)
 - [ProjectModelLinksCreateWorkPackage](docs/ProjectModelLinksCreateWorkPackage.md)
 - [ProjectModelLinksCreateWorkPackageImmediately](docs/ProjectModelLinksCreateWorkPackageImmediately.md)
 - [ProjectModelLinksDelete](docs/ProjectModelLinksDelete.md)
 - [ProjectModelLinksMemberships](docs/ProjectModelLinksMemberships.md)
 - [ProjectModelLinksParent](docs/ProjectModelLinksParent.md)
 - [ProjectModelLinksProjectStorages](docs/ProjectModelLinksProjectStorages.md)
 - [ProjectModelLinksSelf](docs/ProjectModelLinksSelf.md)
 - [ProjectModelLinksStatus](docs/ProjectModelLinksStatus.md)
 - [ProjectModelLinksStoragesInner](docs/ProjectModelLinksStoragesInner.md)
 - [ProjectModelLinksTypes](docs/ProjectModelLinksTypes.md)
 - [ProjectModelLinksUpdate](docs/ProjectModelLinksUpdate.md)
 - [ProjectModelLinksUpdateImmediately](docs/ProjectModelLinksUpdateImmediately.md)
 - [ProjectModelLinksVersions](docs/ProjectModelLinksVersions.md)
 - [ProjectModelLinksWorkPackages](docs/ProjectModelLinksWorkPackages.md)
 - [ProjectModelStatusExplanation](docs/ProjectModelStatusExplanation.md)
 - [ProjectStorageCollectionModel](docs/ProjectStorageCollectionModel.md)
 - [ProjectStorageCollectionModelAllOfEmbedded](docs/ProjectStorageCollectionModelAllOfEmbedded.md)
 - [ProjectStorageCollectionModelAllOfLinks](docs/ProjectStorageCollectionModelAllOfLinks.md)
 - [ProjectStorageCollectionModelAllOfLinksSelf](docs/ProjectStorageCollectionModelAllOfLinksSelf.md)
 - [ProjectStorageModel](docs/ProjectStorageModel.md)
 - [ProjectStorageModelLinks](docs/ProjectStorageModelLinks.md)
 - [ProjectStorageModelLinksCreator](docs/ProjectStorageModelLinksCreator.md)
 - [ProjectStorageModelLinksOpen](docs/ProjectStorageModelLinksOpen.md)
 - [ProjectStorageModelLinksOpenWithConnectionEnsured](docs/ProjectStorageModelLinksOpenWithConnectionEnsured.md)
 - [ProjectStorageModelLinksProject](docs/ProjectStorageModelLinksProject.md)
 - [ProjectStorageModelLinksProjectFolder](docs/ProjectStorageModelLinksProjectFolder.md)
 - [ProjectStorageModelLinksSelf](docs/ProjectStorageModelLinksSelf.md)
 - [ProjectStorageModelLinksStorage](docs/ProjectStorageModelLinksStorage.md)
 - [QueryColumnModel](docs/QueryColumnModel.md)
 - [QueryCreateForm](docs/QueryCreateForm.md)
 - [QueryFilterInstanceSchemaModel](docs/QueryFilterInstanceSchemaModel.md)
 - [QueryFilterInstanceSchemaModelLinks](docs/QueryFilterInstanceSchemaModelLinks.md)
 - [QueryFilterInstanceSchemaModelLinksFilter](docs/QueryFilterInstanceSchemaModelLinksFilter.md)
 - [QueryFilterInstanceSchemaModelLinksSelf](docs/QueryFilterInstanceSchemaModelLinksSelf.md)
 - [QueryFilterModel](docs/QueryFilterModel.md)
 - [QueryModel](docs/QueryModel.md)
 - [QueryModelLinks](docs/QueryModelLinks.md)
 - [QueryModelLinksStar](docs/QueryModelLinksStar.md)
 - [QueryModelLinksUnstar](docs/QueryModelLinksUnstar.md)
 - [QueryModelLinksUpdate](docs/QueryModelLinksUpdate.md)
 - [QueryModelLinksUpdateImmediately](docs/QueryModelLinksUpdateImmediately.md)
 - [QueryOperatorModel](docs/QueryOperatorModel.md)
 - [QuerySortByModel](docs/QuerySortByModel.md)
 - [QueryUpdateForm](docs/QueryUpdateForm.md)
 - [RelationModel](docs/RelationModel.md)
 - [RelationModelLinks](docs/RelationModelLinks.md)
 - [RelationModelLinksDelete](docs/RelationModelLinksDelete.md)
 - [RelationModelLinksFrom](docs/RelationModelLinksFrom.md)
 - [RelationModelLinksSchema](docs/RelationModelLinksSchema.md)
 - [RelationModelLinksSelf](docs/RelationModelLinksSelf.md)
 - [RelationModelLinksTo](docs/RelationModelLinksTo.md)
 - [RelationModelLinksUpdate](docs/RelationModelLinksUpdate.md)
 - [RelationModelLinksUpdateImmediately](docs/RelationModelLinksUpdateImmediately.md)
 - [RevisionModel](docs/RevisionModel.md)
 - [RevisionModelLinks](docs/RevisionModelLinks.md)
 - [RevisionModelLinksAuthor](docs/RevisionModelLinksAuthor.md)
 - [RevisionModelLinksProject](docs/RevisionModelLinksProject.md)
 - [RevisionModelLinksSelf](docs/RevisionModelLinksSelf.md)
 - [RevisionModelLinksShowRevision](docs/RevisionModelLinksShowRevision.md)
 - [RevisionModelMessage](docs/RevisionModelMessage.md)
 - [RoleModel](docs/RoleModel.md)
 - [RoleModelLinks](docs/RoleModelLinks.md)
 - [RoleModelLinksSelf](docs/RoleModelLinksSelf.md)
 - [RootModel](docs/RootModel.md)
 - [RootModelLinks](docs/RootModelLinks.md)
 - [RootModelLinksConfiguration](docs/RootModelLinksConfiguration.md)
 - [RootModelLinksMemberships](docs/RootModelLinksMemberships.md)
 - [RootModelLinksPriorities](docs/RootModelLinksPriorities.md)
 - [RootModelLinksRelations](docs/RootModelLinksRelations.md)
 - [RootModelLinksSelf](docs/RootModelLinksSelf.md)
 - [RootModelLinksStatuses](docs/RootModelLinksStatuses.md)
 - [RootModelLinksTimeEntries](docs/RootModelLinksTimeEntries.md)
 - [RootModelLinksTypes](docs/RootModelLinksTypes.md)
 - [RootModelLinksUser](docs/RootModelLinksUser.md)
 - [RootModelLinksUserPreferences](docs/RootModelLinksUserPreferences.md)
 - [RootModelLinksWorkPackages](docs/RootModelLinksWorkPackages.md)
 - [SchemaModel](docs/SchemaModel.md)
 - [SchemaModelLinks](docs/SchemaModelLinks.md)
 - [SchemaModelLinksSelf](docs/SchemaModelLinksSelf.md)
 - [SchemaPropertyModel](docs/SchemaPropertyModel.md)
 - [ShowOrValidateFormRequest](docs/ShowOrValidateFormRequest.md)
 - [StatusCollectionModel](docs/StatusCollectionModel.md)
 - [StatusCollectionModelAllOfEmbedded](docs/StatusCollectionModelAllOfEmbedded.md)
 - [StatusCollectionModelAllOfLinks](docs/StatusCollectionModelAllOfLinks.md)
 - [StatusCollectionModelAllOfLinksSelf](docs/StatusCollectionModelAllOfLinksSelf.md)
 - [StatusModel](docs/StatusModel.md)
 - [StatusModelLinks](docs/StatusModelLinks.md)
 - [StatusModelLinksSelf](docs/StatusModelLinksSelf.md)
 - [StorageCollectionModel](docs/StorageCollectionModel.md)
 - [StorageCollectionModelAllOfEmbedded](docs/StorageCollectionModelAllOfEmbedded.md)
 - [StorageCollectionModelAllOfLinks](docs/StorageCollectionModelAllOfLinks.md)
 - [StorageCollectionModelAllOfLinksSelf](docs/StorageCollectionModelAllOfLinksSelf.md)
 - [StorageFileModel](docs/StorageFileModel.md)
 - [StorageFileModelAllOfLinks](docs/StorageFileModelAllOfLinks.md)
 - [StorageFileModelAllOfLinksSelf](docs/StorageFileModelAllOfLinksSelf.md)
 - [StorageFileUploadLinkModel](docs/StorageFileUploadLinkModel.md)
 - [StorageFileUploadLinkModelLinks](docs/StorageFileUploadLinkModelLinks.md)
 - [StorageFileUploadLinkModelLinksDestination](docs/StorageFileUploadLinkModelLinksDestination.md)
 - [StorageFileUploadLinkModelLinksSelf](docs/StorageFileUploadLinkModelLinksSelf.md)
 - [StorageFileUploadPreparationModel](docs/StorageFileUploadPreparationModel.md)
 - [StorageFilesModel](docs/StorageFilesModel.md)
 - [StorageFilesModelParent](docs/StorageFilesModelParent.md)
 - [StorageReadModel](docs/StorageReadModel.md)
 - [StorageReadModelEmbedded](docs/StorageReadModelEmbedded.md)
 - [StorageReadModelLinks](docs/StorageReadModelLinks.md)
 - [StorageReadModelLinksAuthorizationState](docs/StorageReadModelLinksAuthorizationState.md)
 - [StorageReadModelLinksAuthorize](docs/StorageReadModelLinksAuthorize.md)
 - [StorageReadModelLinksOauthApplication](docs/StorageReadModelLinksOauthApplication.md)
 - [StorageReadModelLinksOauthClientCredentials](docs/StorageReadModelLinksOauthClientCredentials.md)
 - [StorageReadModelLinksOpen](docs/StorageReadModelLinksOpen.md)
 - [StorageReadModelLinksOrigin](docs/StorageReadModelLinksOrigin.md)
 - [StorageReadModelLinksSelf](docs/StorageReadModelLinksSelf.md)
 - [StorageReadModelLinksType](docs/StorageReadModelLinksType.md)
 - [StorageWriteModel](docs/StorageWriteModel.md)
 - [StorageWriteModelLinks](docs/StorageWriteModelLinks.md)
 - [StorageWriteModelLinksOrigin](docs/StorageWriteModelLinksOrigin.md)
 - [StorageWriteModelLinksType](docs/StorageWriteModelLinksType.md)
 - [TimeEntryActivityModel](docs/TimeEntryActivityModel.md)
 - [TimeEntryActivityModelEmbedded](docs/TimeEntryActivityModelEmbedded.md)
 - [TimeEntryActivityModelLinks](docs/TimeEntryActivityModelLinks.md)
 - [TimeEntryActivityModelLinksProjectsInner](docs/TimeEntryActivityModelLinksProjectsInner.md)
 - [TimeEntryActivityModelLinksSelf](docs/TimeEntryActivityModelLinksSelf.md)
 - [TimeEntryCollectionModel](docs/TimeEntryCollectionModel.md)
 - [TimeEntryCollectionModelAllOfEmbedded](docs/TimeEntryCollectionModelAllOfEmbedded.md)
 - [TimeEntryCollectionModelAllOfLinks](docs/TimeEntryCollectionModelAllOfLinks.md)
 - [TimeEntryCollectionModelAllOfLinksSelf](docs/TimeEntryCollectionModelAllOfLinksSelf.md)
 - [TimeEntryModel](docs/TimeEntryModel.md)
 - [TimeEntryModelComment](docs/TimeEntryModelComment.md)
 - [TimeEntryModelLinks](docs/TimeEntryModelLinks.md)
 - [TimeEntryModelLinksActivity](docs/TimeEntryModelLinksActivity.md)
 - [TimeEntryModelLinksDelete](docs/TimeEntryModelLinksDelete.md)
 - [TimeEntryModelLinksProject](docs/TimeEntryModelLinksProject.md)
 - [TimeEntryModelLinksSchema](docs/TimeEntryModelLinksSchema.md)
 - [TimeEntryModelLinksSelf](docs/TimeEntryModelLinksSelf.md)
 - [TimeEntryModelLinksUpdate](docs/TimeEntryModelLinksUpdate.md)
 - [TimeEntryModelLinksUpdateImmediately](docs/TimeEntryModelLinksUpdateImmediately.md)
 - [TimeEntryModelLinksUser](docs/TimeEntryModelLinksUser.md)
 - [TimeEntryModelLinksWorkPackage](docs/TimeEntryModelLinksWorkPackage.md)
 - [TypeModel](docs/TypeModel.md)
 - [TypeModelLinks](docs/TypeModelLinks.md)
 - [TypeModelLinksSelf](docs/TypeModelLinksSelf.md)
 - [TypesByProjectModel](docs/TypesByProjectModel.md)
 - [TypesByProjectModelAllOfEmbedded](docs/TypesByProjectModelAllOfEmbedded.md)
 - [TypesByProjectModelAllOfEmbeddedElements](docs/TypesByProjectModelAllOfEmbeddedElements.md)
 - [TypesByProjectModelAllOfLinks](docs/TypesByProjectModelAllOfLinks.md)
 - [TypesByProjectModelAllOfLinksSelf](docs/TypesByProjectModelAllOfLinksSelf.md)
 - [UpdateActivityRequest](docs/UpdateActivityRequest.md)
 - [UpdateActivityRequestLinks](docs/UpdateActivityRequestLinks.md)
 - [UpdateUserPreferencesRequest](docs/UpdateUserPreferencesRequest.md)
 - [UserCollectionModel](docs/UserCollectionModel.md)
 - [UserCollectionModelAllOfEmbedded](docs/UserCollectionModelAllOfEmbedded.md)
 - [UserCollectionModelAllOfLinks](docs/UserCollectionModelAllOfLinks.md)
 - [UserCollectionModelAllOfLinksSelf](docs/UserCollectionModelAllOfLinksSelf.md)
 - [UserCreateModel](docs/UserCreateModel.md)
 - [UserModel](docs/UserModel.md)
 - [UserModelAllOfLinks](docs/UserModelAllOfLinks.md)
 - [UserModelAllOfLinksAuthSource](docs/UserModelAllOfLinksAuthSource.md)
 - [UserModelAllOfLinksDelete](docs/UserModelAllOfLinksDelete.md)
 - [UserModelAllOfLinksLock](docs/UserModelAllOfLinksLock.md)
 - [UserModelAllOfLinksShowUser](docs/UserModelAllOfLinksShowUser.md)
 - [UserModelAllOfLinksUnlock](docs/UserModelAllOfLinksUnlock.md)
 - [UserModelAllOfLinksUpdateImmediately](docs/UserModelAllOfLinksUpdateImmediately.md)
 - [ValuesPropertyModel](docs/ValuesPropertyModel.md)
 - [ValuesPropertyModelLinks](docs/ValuesPropertyModelLinks.md)
 - [ValuesPropertyModelLinksSchema](docs/ValuesPropertyModelLinksSchema.md)
 - [ValuesPropertyModelLinksSelf](docs/ValuesPropertyModelLinksSelf.md)
 - [VersionModel](docs/VersionModel.md)
 - [VersionModelLinks](docs/VersionModelLinks.md)
 - [VersionModelLinksAvailableInProjects](docs/VersionModelLinksAvailableInProjects.md)
 - [VersionModelLinksDefiningProject](docs/VersionModelLinksDefiningProject.md)
 - [VersionModelLinksSelf](docs/VersionModelLinksSelf.md)
 - [VersionModelLinksUpdate](docs/VersionModelLinksUpdate.md)
 - [VersionModelLinksUpdateImmediately](docs/VersionModelLinksUpdateImmediately.md)
 - [VersionsByProjectModel](docs/VersionsByProjectModel.md)
 - [VersionsByProjectModelAllOfEmbedded](docs/VersionsByProjectModelAllOfEmbedded.md)
 - [VersionsByProjectModelAllOfEmbeddedElements](docs/VersionsByProjectModelAllOfEmbeddedElements.md)
 - [VersionsByProjectModelAllOfLinks](docs/VersionsByProjectModelAllOfLinks.md)
 - [VersionsByProjectModelAllOfLinksSelf](docs/VersionsByProjectModelAllOfLinksSelf.md)
 - [WatchersModel](docs/WatchersModel.md)
 - [WatchersModelAllOfEmbedded](docs/WatchersModelAllOfEmbedded.md)
 - [WatchersModelAllOfEmbeddedElements](docs/WatchersModelAllOfEmbeddedElements.md)
 - [WatchersModelAllOfLinks](docs/WatchersModelAllOfLinks.md)
 - [WatchersModelAllOfLinksSelf](docs/WatchersModelAllOfLinksSelf.md)
 - [WeekDayCollectionModel](docs/WeekDayCollectionModel.md)
 - [WeekDayCollectionModelAllOfEmbedded](docs/WeekDayCollectionModelAllOfEmbedded.md)
 - [WeekDayCollectionModelAllOfLinks](docs/WeekDayCollectionModelAllOfLinks.md)
 - [WeekDayCollectionModelAllOfLinksSelf](docs/WeekDayCollectionModelAllOfLinksSelf.md)
 - [WeekDayCollectionWriteModel](docs/WeekDayCollectionWriteModel.md)
 - [WeekDayCollectionWriteModelEmbedded](docs/WeekDayCollectionWriteModelEmbedded.md)
 - [WeekDayCollectionWriteModelEmbeddedElementsInner](docs/WeekDayCollectionWriteModelEmbeddedElementsInner.md)
 - [WeekDayModel](docs/WeekDayModel.md)
 - [WeekDaySelfLinkModel](docs/WeekDaySelfLinkModel.md)
 - [WeekDaySelfLinkModelSelf](docs/WeekDaySelfLinkModelSelf.md)
 - [WeekDayWriteModel](docs/WeekDayWriteModel.md)
 - [WikiPageModel](docs/WikiPageModel.md)
 - [WikiPageModelLinks](docs/WikiPageModelLinks.md)
 - [WikiPageModelLinksAddAttachment](docs/WikiPageModelLinksAddAttachment.md)
 - [WorkPackageActivitiesModel](docs/WorkPackageActivitiesModel.md)
 - [WorkPackageActivitiesModelAllOfEmbedded](docs/WorkPackageActivitiesModelAllOfEmbedded.md)
 - [WorkPackageModel](docs/WorkPackageModel.md)
 - [WorkPackageModelDescription](docs/WorkPackageModelDescription.md)
 - [WorkPackageModelLinks](docs/WorkPackageModelLinks.md)
 - [WorkPackageModelLinksAddAttachment](docs/WorkPackageModelLinksAddAttachment.md)
 - [WorkPackageModelLinksAddComment](docs/WorkPackageModelLinksAddComment.md)
 - [WorkPackageModelLinksAddFileLink](docs/WorkPackageModelLinksAddFileLink.md)
 - [WorkPackageModelLinksAddRelation](docs/WorkPackageModelLinksAddRelation.md)
 - [WorkPackageModelLinksAddWatcher](docs/WorkPackageModelLinksAddWatcher.md)
 - [WorkPackageModelLinksAncestorsInner](docs/WorkPackageModelLinksAncestorsInner.md)
 - [WorkPackageModelLinksAssignee](docs/WorkPackageModelLinksAssignee.md)
 - [WorkPackageModelLinksAuthor](docs/WorkPackageModelLinksAuthor.md)
 - [WorkPackageModelLinksAvailableWatchers](docs/WorkPackageModelLinksAvailableWatchers.md)
 - [WorkPackageModelLinksBudget](docs/WorkPackageModelLinksBudget.md)
 - [WorkPackageModelLinksCategory](docs/WorkPackageModelLinksCategory.md)
 - [WorkPackageModelLinksChildrenInner](docs/WorkPackageModelLinksChildrenInner.md)
 - [WorkPackageModelLinksCustomActionsInner](docs/WorkPackageModelLinksCustomActionsInner.md)
 - [WorkPackageModelLinksFileLinks](docs/WorkPackageModelLinksFileLinks.md)
 - [WorkPackageModelLinksParent](docs/WorkPackageModelLinksParent.md)
 - [WorkPackageModelLinksPreviewMarkup](docs/WorkPackageModelLinksPreviewMarkup.md)
 - [WorkPackageModelLinksPriority](docs/WorkPackageModelLinksPriority.md)
 - [WorkPackageModelLinksProject](docs/WorkPackageModelLinksProject.md)
 - [WorkPackageModelLinksRelations](docs/WorkPackageModelLinksRelations.md)
 - [WorkPackageModelLinksRemoveWatcher](docs/WorkPackageModelLinksRemoveWatcher.md)
 - [WorkPackageModelLinksResponsible](docs/WorkPackageModelLinksResponsible.md)
 - [WorkPackageModelLinksRevisions](docs/WorkPackageModelLinksRevisions.md)
 - [WorkPackageModelLinksSchema](docs/WorkPackageModelLinksSchema.md)
 - [WorkPackageModelLinksSelf](docs/WorkPackageModelLinksSelf.md)
 - [WorkPackageModelLinksStatus](docs/WorkPackageModelLinksStatus.md)
 - [WorkPackageModelLinksTimeEntries](docs/WorkPackageModelLinksTimeEntries.md)
 - [WorkPackageModelLinksType](docs/WorkPackageModelLinksType.md)
 - [WorkPackageModelLinksUnwatch](docs/WorkPackageModelLinksUnwatch.md)
 - [WorkPackageModelLinksUpdate](docs/WorkPackageModelLinksUpdate.md)
 - [WorkPackageModelLinksUpdateImmediately](docs/WorkPackageModelLinksUpdateImmediately.md)
 - [WorkPackageModelLinksVersion](docs/WorkPackageModelLinksVersion.md)
 - [WorkPackageModelLinksWatch](docs/WorkPackageModelLinksWatch.md)
 - [WorkPackageModelLinksWatchers](docs/WorkPackageModelLinksWatchers.md)
 - [WorkPackagePatchModel](docs/WorkPackagePatchModel.md)
 - [WorkPackagePatchModelLinks](docs/WorkPackagePatchModelLinks.md)
 - [WorkPackagesModel](docs/WorkPackagesModel.md)
 - [WorkPackagesModelAllOfEmbedded](docs/WorkPackagesModelAllOfEmbedded.md)
 - [WorkPackagesModelAllOfLinks](docs/WorkPackagesModelAllOfLinks.md)
 - [WorkPackagesModelAllOfLinksSelf](docs/WorkPackagesModelAllOfLinksSelf.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### BasicAuth

- **Type**: HTTP basic authentication

Example

```go
auth := context.WithValue(context.Background(), openproject.ContextBasicAuth, openproject.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



