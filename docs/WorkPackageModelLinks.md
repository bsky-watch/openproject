# WorkPackageModelLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddAttachment** | Pointer to [**WorkPackageModelLinksAddAttachment**](WorkPackageModelLinksAddAttachment.md) |  | [optional] 
**AddComment** | Pointer to [**WorkPackageModelLinksAddComment**](WorkPackageModelLinksAddComment.md) |  | [optional] 
**AddRelation** | Pointer to [**WorkPackageModelLinksAddRelation**](WorkPackageModelLinksAddRelation.md) |  | [optional] 
**AddWatcher** | Pointer to [**WorkPackageModelLinksAddWatcher**](WorkPackageModelLinksAddWatcher.md) |  | [optional] 
**CustomActions** | Pointer to [**[]WorkPackageModelLinksCustomActionsInner**](WorkPackageModelLinksCustomActionsInner.md) |  | [optional] 
**PreviewMarkup** | Pointer to [**WorkPackageModelLinksPreviewMarkup**](WorkPackageModelLinksPreviewMarkup.md) |  | [optional] 
**RemoveWatcher** | Pointer to [**WorkPackageModelLinksRemoveWatcher**](WorkPackageModelLinksRemoveWatcher.md) |  | [optional] 
**Unwatch** | Pointer to [**WorkPackageModelLinksUnwatch**](WorkPackageModelLinksUnwatch.md) |  | [optional] 
**Update** | Pointer to [**WorkPackageModelLinksUpdate**](WorkPackageModelLinksUpdate.md) |  | [optional] 
**UpdateImmediately** | Pointer to [**WorkPackageModelLinksUpdateImmediately**](WorkPackageModelLinksUpdateImmediately.md) |  | [optional] 
**Watch** | Pointer to [**WorkPackageModelLinksWatch**](WorkPackageModelLinksWatch.md) |  | [optional] 
**Self** | Pointer to [**WorkPackageModelLinksSelf**](WorkPackageModelLinksSelf.md) |  | [optional] 
**Schema** | Pointer to [**WorkPackageModelLinksSchema**](WorkPackageModelLinksSchema.md) |  | [optional] 
**Ancestors** | Pointer to [**[]WorkPackageModelLinksAncestorsInner**](WorkPackageModelLinksAncestorsInner.md) |  | [optional] 
**Attachments** | Pointer to [**Link**](Link.md) |  | [optional] 
**Author** | Pointer to [**WorkPackageModelLinksAuthor**](WorkPackageModelLinksAuthor.md) |  | [optional] 
**Assignee** | Pointer to [**WorkPackageModelLinksAssignee**](WorkPackageModelLinksAssignee.md) |  | [optional] 
**AvailableWatchers** | Pointer to [**WorkPackageModelLinksAvailableWatchers**](WorkPackageModelLinksAvailableWatchers.md) |  | [optional] 
**Budget** | Pointer to [**WorkPackageModelLinksBudget**](WorkPackageModelLinksBudget.md) |  | [optional] 
**Category** | Pointer to [**WorkPackageModelLinksCategory**](WorkPackageModelLinksCategory.md) |  | [optional] 
**Children** | Pointer to [**[]WorkPackageModelLinksChildrenInner**](WorkPackageModelLinksChildrenInner.md) |  | [optional] 
**AddFileLink** | Pointer to [**WorkPackageModelLinksAddFileLink**](WorkPackageModelLinksAddFileLink.md) |  | [optional] 
**FileLinks** | Pointer to [**WorkPackageModelLinksFileLinks**](WorkPackageModelLinksFileLinks.md) |  | [optional] 
**Parent** | Pointer to [**WorkPackageModelLinksParent**](WorkPackageModelLinksParent.md) |  | [optional] 
**Priority** | Pointer to [**WorkPackageModelLinksPriority**](WorkPackageModelLinksPriority.md) |  | [optional] 
**Project** | Pointer to [**WorkPackageModelLinksProject**](WorkPackageModelLinksProject.md) |  | [optional] 
**Responsible** | Pointer to [**WorkPackageModelLinksResponsible**](WorkPackageModelLinksResponsible.md) |  | [optional] 
**Relations** | Pointer to [**WorkPackageModelLinksRelations**](WorkPackageModelLinksRelations.md) |  | [optional] 
**Revisions** | Pointer to [**WorkPackageModelLinksRevisions**](WorkPackageModelLinksRevisions.md) |  | [optional] 
**Status** | Pointer to [**WorkPackageModelLinksStatus**](WorkPackageModelLinksStatus.md) |  | [optional] 
**TimeEntries** | Pointer to [**WorkPackageModelLinksTimeEntries**](WorkPackageModelLinksTimeEntries.md) |  | [optional] 
**Type** | Pointer to [**WorkPackageModelLinksType**](WorkPackageModelLinksType.md) |  | [optional] 
**Version** | Pointer to [**WorkPackageModelLinksVersion**](WorkPackageModelLinksVersion.md) |  | [optional] 
**Watchers** | Pointer to [**WorkPackageModelLinksWatchers**](WorkPackageModelLinksWatchers.md) |  | [optional] 

## Methods

### NewWorkPackageModelLinks

`func NewWorkPackageModelLinks() *WorkPackageModelLinks`

NewWorkPackageModelLinks instantiates a new WorkPackageModelLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkPackageModelLinksWithDefaults

`func NewWorkPackageModelLinksWithDefaults() *WorkPackageModelLinks`

NewWorkPackageModelLinksWithDefaults instantiates a new WorkPackageModelLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddAttachment

`func (o *WorkPackageModelLinks) GetAddAttachment() WorkPackageModelLinksAddAttachment`

GetAddAttachment returns the AddAttachment field if non-nil, zero value otherwise.

### GetAddAttachmentOk

`func (o *WorkPackageModelLinks) GetAddAttachmentOk() (*WorkPackageModelLinksAddAttachment, bool)`

GetAddAttachmentOk returns a tuple with the AddAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAttachment

`func (o *WorkPackageModelLinks) SetAddAttachment(v WorkPackageModelLinksAddAttachment)`

SetAddAttachment sets AddAttachment field to given value.

### HasAddAttachment

`func (o *WorkPackageModelLinks) HasAddAttachment() bool`

HasAddAttachment returns a boolean if a field has been set.

### GetAddComment

`func (o *WorkPackageModelLinks) GetAddComment() WorkPackageModelLinksAddComment`

GetAddComment returns the AddComment field if non-nil, zero value otherwise.

### GetAddCommentOk

`func (o *WorkPackageModelLinks) GetAddCommentOk() (*WorkPackageModelLinksAddComment, bool)`

GetAddCommentOk returns a tuple with the AddComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddComment

`func (o *WorkPackageModelLinks) SetAddComment(v WorkPackageModelLinksAddComment)`

SetAddComment sets AddComment field to given value.

### HasAddComment

`func (o *WorkPackageModelLinks) HasAddComment() bool`

HasAddComment returns a boolean if a field has been set.

### GetAddRelation

`func (o *WorkPackageModelLinks) GetAddRelation() WorkPackageModelLinksAddRelation`

GetAddRelation returns the AddRelation field if non-nil, zero value otherwise.

### GetAddRelationOk

`func (o *WorkPackageModelLinks) GetAddRelationOk() (*WorkPackageModelLinksAddRelation, bool)`

GetAddRelationOk returns a tuple with the AddRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddRelation

`func (o *WorkPackageModelLinks) SetAddRelation(v WorkPackageModelLinksAddRelation)`

SetAddRelation sets AddRelation field to given value.

### HasAddRelation

`func (o *WorkPackageModelLinks) HasAddRelation() bool`

HasAddRelation returns a boolean if a field has been set.

### GetAddWatcher

`func (o *WorkPackageModelLinks) GetAddWatcher() WorkPackageModelLinksAddWatcher`

GetAddWatcher returns the AddWatcher field if non-nil, zero value otherwise.

### GetAddWatcherOk

`func (o *WorkPackageModelLinks) GetAddWatcherOk() (*WorkPackageModelLinksAddWatcher, bool)`

GetAddWatcherOk returns a tuple with the AddWatcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddWatcher

`func (o *WorkPackageModelLinks) SetAddWatcher(v WorkPackageModelLinksAddWatcher)`

SetAddWatcher sets AddWatcher field to given value.

### HasAddWatcher

`func (o *WorkPackageModelLinks) HasAddWatcher() bool`

HasAddWatcher returns a boolean if a field has been set.

### GetCustomActions

`func (o *WorkPackageModelLinks) GetCustomActions() []WorkPackageModelLinksCustomActionsInner`

GetCustomActions returns the CustomActions field if non-nil, zero value otherwise.

### GetCustomActionsOk

`func (o *WorkPackageModelLinks) GetCustomActionsOk() (*[]WorkPackageModelLinksCustomActionsInner, bool)`

GetCustomActionsOk returns a tuple with the CustomActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomActions

`func (o *WorkPackageModelLinks) SetCustomActions(v []WorkPackageModelLinksCustomActionsInner)`

SetCustomActions sets CustomActions field to given value.

### HasCustomActions

`func (o *WorkPackageModelLinks) HasCustomActions() bool`

HasCustomActions returns a boolean if a field has been set.

### GetPreviewMarkup

`func (o *WorkPackageModelLinks) GetPreviewMarkup() WorkPackageModelLinksPreviewMarkup`

GetPreviewMarkup returns the PreviewMarkup field if non-nil, zero value otherwise.

### GetPreviewMarkupOk

`func (o *WorkPackageModelLinks) GetPreviewMarkupOk() (*WorkPackageModelLinksPreviewMarkup, bool)`

GetPreviewMarkupOk returns a tuple with the PreviewMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewMarkup

`func (o *WorkPackageModelLinks) SetPreviewMarkup(v WorkPackageModelLinksPreviewMarkup)`

SetPreviewMarkup sets PreviewMarkup field to given value.

### HasPreviewMarkup

`func (o *WorkPackageModelLinks) HasPreviewMarkup() bool`

HasPreviewMarkup returns a boolean if a field has been set.

### GetRemoveWatcher

`func (o *WorkPackageModelLinks) GetRemoveWatcher() WorkPackageModelLinksRemoveWatcher`

GetRemoveWatcher returns the RemoveWatcher field if non-nil, zero value otherwise.

### GetRemoveWatcherOk

`func (o *WorkPackageModelLinks) GetRemoveWatcherOk() (*WorkPackageModelLinksRemoveWatcher, bool)`

GetRemoveWatcherOk returns a tuple with the RemoveWatcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveWatcher

`func (o *WorkPackageModelLinks) SetRemoveWatcher(v WorkPackageModelLinksRemoveWatcher)`

SetRemoveWatcher sets RemoveWatcher field to given value.

### HasRemoveWatcher

`func (o *WorkPackageModelLinks) HasRemoveWatcher() bool`

HasRemoveWatcher returns a boolean if a field has been set.

### GetUnwatch

`func (o *WorkPackageModelLinks) GetUnwatch() WorkPackageModelLinksUnwatch`

GetUnwatch returns the Unwatch field if non-nil, zero value otherwise.

### GetUnwatchOk

`func (o *WorkPackageModelLinks) GetUnwatchOk() (*WorkPackageModelLinksUnwatch, bool)`

GetUnwatchOk returns a tuple with the Unwatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnwatch

`func (o *WorkPackageModelLinks) SetUnwatch(v WorkPackageModelLinksUnwatch)`

SetUnwatch sets Unwatch field to given value.

### HasUnwatch

`func (o *WorkPackageModelLinks) HasUnwatch() bool`

HasUnwatch returns a boolean if a field has been set.

### GetUpdate

`func (o *WorkPackageModelLinks) GetUpdate() WorkPackageModelLinksUpdate`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *WorkPackageModelLinks) GetUpdateOk() (*WorkPackageModelLinksUpdate, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *WorkPackageModelLinks) SetUpdate(v WorkPackageModelLinksUpdate)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *WorkPackageModelLinks) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUpdateImmediately

`func (o *WorkPackageModelLinks) GetUpdateImmediately() WorkPackageModelLinksUpdateImmediately`

GetUpdateImmediately returns the UpdateImmediately field if non-nil, zero value otherwise.

### GetUpdateImmediatelyOk

`func (o *WorkPackageModelLinks) GetUpdateImmediatelyOk() (*WorkPackageModelLinksUpdateImmediately, bool)`

GetUpdateImmediatelyOk returns a tuple with the UpdateImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateImmediately

`func (o *WorkPackageModelLinks) SetUpdateImmediately(v WorkPackageModelLinksUpdateImmediately)`

SetUpdateImmediately sets UpdateImmediately field to given value.

### HasUpdateImmediately

`func (o *WorkPackageModelLinks) HasUpdateImmediately() bool`

HasUpdateImmediately returns a boolean if a field has been set.

### GetWatch

`func (o *WorkPackageModelLinks) GetWatch() WorkPackageModelLinksWatch`

GetWatch returns the Watch field if non-nil, zero value otherwise.

### GetWatchOk

`func (o *WorkPackageModelLinks) GetWatchOk() (*WorkPackageModelLinksWatch, bool)`

GetWatchOk returns a tuple with the Watch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatch

`func (o *WorkPackageModelLinks) SetWatch(v WorkPackageModelLinksWatch)`

SetWatch sets Watch field to given value.

### HasWatch

`func (o *WorkPackageModelLinks) HasWatch() bool`

HasWatch returns a boolean if a field has been set.

### GetSelf

`func (o *WorkPackageModelLinks) GetSelf() WorkPackageModelLinksSelf`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *WorkPackageModelLinks) GetSelfOk() (*WorkPackageModelLinksSelf, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *WorkPackageModelLinks) SetSelf(v WorkPackageModelLinksSelf)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *WorkPackageModelLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSchema

`func (o *WorkPackageModelLinks) GetSchema() WorkPackageModelLinksSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *WorkPackageModelLinks) GetSchemaOk() (*WorkPackageModelLinksSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *WorkPackageModelLinks) SetSchema(v WorkPackageModelLinksSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *WorkPackageModelLinks) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetAncestors

`func (o *WorkPackageModelLinks) GetAncestors() []WorkPackageModelLinksAncestorsInner`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *WorkPackageModelLinks) GetAncestorsOk() (*[]WorkPackageModelLinksAncestorsInner, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *WorkPackageModelLinks) SetAncestors(v []WorkPackageModelLinksAncestorsInner)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *WorkPackageModelLinks) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### GetAttachments

`func (o *WorkPackageModelLinks) GetAttachments() Link`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *WorkPackageModelLinks) GetAttachmentsOk() (*Link, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *WorkPackageModelLinks) SetAttachments(v Link)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *WorkPackageModelLinks) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetAuthor

`func (o *WorkPackageModelLinks) GetAuthor() WorkPackageModelLinksAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *WorkPackageModelLinks) GetAuthorOk() (*WorkPackageModelLinksAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *WorkPackageModelLinks) SetAuthor(v WorkPackageModelLinksAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *WorkPackageModelLinks) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetAssignee

`func (o *WorkPackageModelLinks) GetAssignee() WorkPackageModelLinksAssignee`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *WorkPackageModelLinks) GetAssigneeOk() (*WorkPackageModelLinksAssignee, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *WorkPackageModelLinks) SetAssignee(v WorkPackageModelLinksAssignee)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *WorkPackageModelLinks) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetAvailableWatchers

`func (o *WorkPackageModelLinks) GetAvailableWatchers() WorkPackageModelLinksAvailableWatchers`

GetAvailableWatchers returns the AvailableWatchers field if non-nil, zero value otherwise.

### GetAvailableWatchersOk

`func (o *WorkPackageModelLinks) GetAvailableWatchersOk() (*WorkPackageModelLinksAvailableWatchers, bool)`

GetAvailableWatchersOk returns a tuple with the AvailableWatchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableWatchers

`func (o *WorkPackageModelLinks) SetAvailableWatchers(v WorkPackageModelLinksAvailableWatchers)`

SetAvailableWatchers sets AvailableWatchers field to given value.

### HasAvailableWatchers

`func (o *WorkPackageModelLinks) HasAvailableWatchers() bool`

HasAvailableWatchers returns a boolean if a field has been set.

### GetBudget

`func (o *WorkPackageModelLinks) GetBudget() WorkPackageModelLinksBudget`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *WorkPackageModelLinks) GetBudgetOk() (*WorkPackageModelLinksBudget, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *WorkPackageModelLinks) SetBudget(v WorkPackageModelLinksBudget)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *WorkPackageModelLinks) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetCategory

`func (o *WorkPackageModelLinks) GetCategory() WorkPackageModelLinksCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *WorkPackageModelLinks) GetCategoryOk() (*WorkPackageModelLinksCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *WorkPackageModelLinks) SetCategory(v WorkPackageModelLinksCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *WorkPackageModelLinks) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetChildren

`func (o *WorkPackageModelLinks) GetChildren() []WorkPackageModelLinksChildrenInner`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *WorkPackageModelLinks) GetChildrenOk() (*[]WorkPackageModelLinksChildrenInner, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *WorkPackageModelLinks) SetChildren(v []WorkPackageModelLinksChildrenInner)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *WorkPackageModelLinks) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetAddFileLink

`func (o *WorkPackageModelLinks) GetAddFileLink() WorkPackageModelLinksAddFileLink`

GetAddFileLink returns the AddFileLink field if non-nil, zero value otherwise.

### GetAddFileLinkOk

`func (o *WorkPackageModelLinks) GetAddFileLinkOk() (*WorkPackageModelLinksAddFileLink, bool)`

GetAddFileLinkOk returns a tuple with the AddFileLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddFileLink

`func (o *WorkPackageModelLinks) SetAddFileLink(v WorkPackageModelLinksAddFileLink)`

SetAddFileLink sets AddFileLink field to given value.

### HasAddFileLink

`func (o *WorkPackageModelLinks) HasAddFileLink() bool`

HasAddFileLink returns a boolean if a field has been set.

### GetFileLinks

`func (o *WorkPackageModelLinks) GetFileLinks() WorkPackageModelLinksFileLinks`

GetFileLinks returns the FileLinks field if non-nil, zero value otherwise.

### GetFileLinksOk

`func (o *WorkPackageModelLinks) GetFileLinksOk() (*WorkPackageModelLinksFileLinks, bool)`

GetFileLinksOk returns a tuple with the FileLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLinks

`func (o *WorkPackageModelLinks) SetFileLinks(v WorkPackageModelLinksFileLinks)`

SetFileLinks sets FileLinks field to given value.

### HasFileLinks

`func (o *WorkPackageModelLinks) HasFileLinks() bool`

HasFileLinks returns a boolean if a field has been set.

### GetParent

`func (o *WorkPackageModelLinks) GetParent() WorkPackageModelLinksParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *WorkPackageModelLinks) GetParentOk() (*WorkPackageModelLinksParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *WorkPackageModelLinks) SetParent(v WorkPackageModelLinksParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *WorkPackageModelLinks) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPriority

`func (o *WorkPackageModelLinks) GetPriority() WorkPackageModelLinksPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WorkPackageModelLinks) GetPriorityOk() (*WorkPackageModelLinksPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WorkPackageModelLinks) SetPriority(v WorkPackageModelLinksPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *WorkPackageModelLinks) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProject

`func (o *WorkPackageModelLinks) GetProject() WorkPackageModelLinksProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *WorkPackageModelLinks) GetProjectOk() (*WorkPackageModelLinksProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *WorkPackageModelLinks) SetProject(v WorkPackageModelLinksProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *WorkPackageModelLinks) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetResponsible

`func (o *WorkPackageModelLinks) GetResponsible() WorkPackageModelLinksResponsible`

GetResponsible returns the Responsible field if non-nil, zero value otherwise.

### GetResponsibleOk

`func (o *WorkPackageModelLinks) GetResponsibleOk() (*WorkPackageModelLinksResponsible, bool)`

GetResponsibleOk returns a tuple with the Responsible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsible

`func (o *WorkPackageModelLinks) SetResponsible(v WorkPackageModelLinksResponsible)`

SetResponsible sets Responsible field to given value.

### HasResponsible

`func (o *WorkPackageModelLinks) HasResponsible() bool`

HasResponsible returns a boolean if a field has been set.

### GetRelations

`func (o *WorkPackageModelLinks) GetRelations() WorkPackageModelLinksRelations`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *WorkPackageModelLinks) GetRelationsOk() (*WorkPackageModelLinksRelations, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *WorkPackageModelLinks) SetRelations(v WorkPackageModelLinksRelations)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *WorkPackageModelLinks) HasRelations() bool`

HasRelations returns a boolean if a field has been set.

### GetRevisions

`func (o *WorkPackageModelLinks) GetRevisions() WorkPackageModelLinksRevisions`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *WorkPackageModelLinks) GetRevisionsOk() (*WorkPackageModelLinksRevisions, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *WorkPackageModelLinks) SetRevisions(v WorkPackageModelLinksRevisions)`

SetRevisions sets Revisions field to given value.

### HasRevisions

`func (o *WorkPackageModelLinks) HasRevisions() bool`

HasRevisions returns a boolean if a field has been set.

### GetStatus

`func (o *WorkPackageModelLinks) GetStatus() WorkPackageModelLinksStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkPackageModelLinks) GetStatusOk() (*WorkPackageModelLinksStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkPackageModelLinks) SetStatus(v WorkPackageModelLinksStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkPackageModelLinks) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeEntries

`func (o *WorkPackageModelLinks) GetTimeEntries() WorkPackageModelLinksTimeEntries`

GetTimeEntries returns the TimeEntries field if non-nil, zero value otherwise.

### GetTimeEntriesOk

`func (o *WorkPackageModelLinks) GetTimeEntriesOk() (*WorkPackageModelLinksTimeEntries, bool)`

GetTimeEntriesOk returns a tuple with the TimeEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEntries

`func (o *WorkPackageModelLinks) SetTimeEntries(v WorkPackageModelLinksTimeEntries)`

SetTimeEntries sets TimeEntries field to given value.

### HasTimeEntries

`func (o *WorkPackageModelLinks) HasTimeEntries() bool`

HasTimeEntries returns a boolean if a field has been set.

### GetType

`func (o *WorkPackageModelLinks) GetType() WorkPackageModelLinksType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkPackageModelLinks) GetTypeOk() (*WorkPackageModelLinksType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkPackageModelLinks) SetType(v WorkPackageModelLinksType)`

SetType sets Type field to given value.

### HasType

`func (o *WorkPackageModelLinks) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *WorkPackageModelLinks) GetVersion() WorkPackageModelLinksVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkPackageModelLinks) GetVersionOk() (*WorkPackageModelLinksVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkPackageModelLinks) SetVersion(v WorkPackageModelLinksVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkPackageModelLinks) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWatchers

`func (o *WorkPackageModelLinks) GetWatchers() WorkPackageModelLinksWatchers`

GetWatchers returns the Watchers field if non-nil, zero value otherwise.

### GetWatchersOk

`func (o *WorkPackageModelLinks) GetWatchersOk() (*WorkPackageModelLinksWatchers, bool)`

GetWatchersOk returns a tuple with the Watchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchers

`func (o *WorkPackageModelLinks) SetWatchers(v WorkPackageModelLinksWatchers)`

SetWatchers sets Watchers field to given value.

### HasWatchers

`func (o *WorkPackageModelLinks) HasWatchers() bool`

HasWatchers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


