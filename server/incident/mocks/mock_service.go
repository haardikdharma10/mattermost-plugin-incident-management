// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-incident-management/server/incident (interfaces: Service)

// Package mock_incident is a generated GoMock package.
package mock_incident

import (
	gomock "github.com/golang/mock/gomock"
	incident "github.com/mattermost/mattermost-plugin-incident-management/server/incident"
	playbook "github.com/mattermost/mattermost-plugin-incident-management/server/playbook"
	model "github.com/mattermost/mattermost-server/v5/model"
	reflect "reflect"
	time "time"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// AddChecklistItem mocks base method
func (m *MockService) AddChecklistItem(arg0, arg1 string, arg2 int, arg3 playbook.ChecklistItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddChecklistItem", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddChecklistItem indicates an expected call of AddChecklistItem
func (mr *MockServiceMockRecorder) AddChecklistItem(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddChecklistItem", reflect.TypeOf((*MockService)(nil).AddChecklistItem), arg0, arg1, arg2, arg3)
}

// ChangeActiveStage mocks base method
func (m *MockService) ChangeActiveStage(arg0, arg1 string, arg2 int) (*incident.Incident, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeActiveStage", arg0, arg1, arg2)
	ret0, _ := ret[0].(*incident.Incident)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeActiveStage indicates an expected call of ChangeActiveStage
func (mr *MockServiceMockRecorder) ChangeActiveStage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeActiveStage", reflect.TypeOf((*MockService)(nil).ChangeActiveStage), arg0, arg1, arg2)
}

// ChangeCommander mocks base method
func (m *MockService) ChangeCommander(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeCommander", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeCommander indicates an expected call of ChangeCommander
func (mr *MockServiceMockRecorder) ChangeCommander(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeCommander", reflect.TypeOf((*MockService)(nil).ChangeCommander), arg0, arg1, arg2)
}

// ChangeCreationDate mocks base method
func (m *MockService) ChangeCreationDate(arg0 string, arg1 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeCreationDate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeCreationDate indicates an expected call of ChangeCreationDate
func (mr *MockServiceMockRecorder) ChangeCreationDate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeCreationDate", reflect.TypeOf((*MockService)(nil).ChangeCreationDate), arg0, arg1)
}

// CreateIncident mocks base method
func (m *MockService) CreateIncident(arg0 *incident.Incident, arg1 string, arg2 bool) (*incident.Incident, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIncident", arg0, arg1, arg2)
	ret0, _ := ret[0].(*incident.Incident)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIncident indicates an expected call of CreateIncident
func (mr *MockServiceMockRecorder) CreateIncident(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIncident", reflect.TypeOf((*MockService)(nil).CreateIncident), arg0, arg1, arg2)
}

// EndIncident mocks base method
func (m *MockService) EndIncident(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndIncident", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EndIncident indicates an expected call of EndIncident
func (mr *MockServiceMockRecorder) EndIncident(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndIncident", reflect.TypeOf((*MockService)(nil).EndIncident), arg0, arg1)
}

// GetChecklistAutocomplete mocks base method
func (m *MockService) GetChecklistAutocomplete(arg0 string) ([]model.AutocompleteListItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChecklistAutocomplete", arg0)
	ret0, _ := ret[0].([]model.AutocompleteListItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChecklistAutocomplete indicates an expected call of GetChecklistAutocomplete
func (mr *MockServiceMockRecorder) GetChecklistAutocomplete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChecklistAutocomplete", reflect.TypeOf((*MockService)(nil).GetChecklistAutocomplete), arg0)
}

// GetCommanders mocks base method
func (m *MockService) GetCommanders(arg0 incident.RequesterInfo, arg1 incident.HeaderFilterOptions) ([]incident.CommanderInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommanders", arg0, arg1)
	ret0, _ := ret[0].([]incident.CommanderInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommanders indicates an expected call of GetCommanders
func (mr *MockServiceMockRecorder) GetCommanders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommanders", reflect.TypeOf((*MockService)(nil).GetCommanders), arg0, arg1)
}

// GetIncident mocks base method
func (m *MockService) GetIncident(arg0 string) (*incident.Incident, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncident", arg0)
	ret0, _ := ret[0].(*incident.Incident)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncident indicates an expected call of GetIncident
func (mr *MockServiceMockRecorder) GetIncident(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncident", reflect.TypeOf((*MockService)(nil).GetIncident), arg0)
}

// GetIncidentFromRecentUpdatePost mocks base method
func (m *MockService) GetIncidentFromRecentUpdatePost(arg0 *model.Post) (*incident.Incident, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncidentFromRecentUpdatePost", arg0)
	ret0, _ := ret[0].(*incident.Incident)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncidentFromRecentUpdatePost indicates an expected call of GetIncidentFromRecentUpdatePost
func (mr *MockServiceMockRecorder) GetIncidentFromRecentUpdatePost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncidentFromRecentUpdatePost", reflect.TypeOf((*MockService)(nil).GetIncidentFromRecentUpdatePost), arg0)
}

// GetIncidentIDForChannel mocks base method
func (m *MockService) GetIncidentIDForChannel(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncidentIDForChannel", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncidentIDForChannel indicates an expected call of GetIncidentIDForChannel
func (mr *MockServiceMockRecorder) GetIncidentIDForChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncidentIDForChannel", reflect.TypeOf((*MockService)(nil).GetIncidentIDForChannel), arg0)
}

// GetIncidentMetadata mocks base method
func (m *MockService) GetIncidentMetadata(arg0 string) (*incident.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncidentMetadata", arg0)
	ret0, _ := ret[0].(*incident.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncidentMetadata indicates an expected call of GetIncidentMetadata
func (mr *MockServiceMockRecorder) GetIncidentMetadata(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncidentMetadata", reflect.TypeOf((*MockService)(nil).GetIncidentMetadata), arg0)
}

// GetIncidents mocks base method
func (m *MockService) GetIncidents(arg0 incident.RequesterInfo, arg1 incident.HeaderFilterOptions) (*incident.GetIncidentsResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncidents", arg0, arg1)
	ret0, _ := ret[0].(*incident.GetIncidentsResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncidents indicates an expected call of GetIncidents
func (mr *MockServiceMockRecorder) GetIncidents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncidents", reflect.TypeOf((*MockService)(nil).GetIncidents), arg0, arg1)
}

// HandleReminder mocks base method
func (m *MockService) HandleReminder(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleReminder", arg0)
}

// HandleReminder indicates an expected call of HandleReminder
func (mr *MockServiceMockRecorder) HandleReminder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleReminder", reflect.TypeOf((*MockService)(nil).HandleReminder), arg0)
}

// IsCommander mocks base method
func (m *MockService) IsCommander(arg0, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCommander", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsCommander indicates an expected call of IsCommander
func (mr *MockServiceMockRecorder) IsCommander(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCommander", reflect.TypeOf((*MockService)(nil).IsCommander), arg0, arg1)
}

// ModifyCheckedState mocks base method
func (m *MockService) ModifyCheckedState(arg0, arg1, arg2 string, arg3, arg4 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyCheckedState", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyCheckedState indicates an expected call of ModifyCheckedState
func (mr *MockServiceMockRecorder) ModifyCheckedState(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyCheckedState", reflect.TypeOf((*MockService)(nil).ModifyCheckedState), arg0, arg1, arg2, arg3, arg4)
}

// MoveChecklistItem mocks base method
func (m *MockService) MoveChecklistItem(arg0, arg1 string, arg2, arg3, arg4 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveChecklistItem", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveChecklistItem indicates an expected call of MoveChecklistItem
func (mr *MockServiceMockRecorder) MoveChecklistItem(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveChecklistItem", reflect.TypeOf((*MockService)(nil).MoveChecklistItem), arg0, arg1, arg2, arg3, arg4)
}

// NukeDB mocks base method
func (m *MockService) NukeDB() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NukeDB")
	ret0, _ := ret[0].(error)
	return ret0
}

// NukeDB indicates an expected call of NukeDB
func (mr *MockServiceMockRecorder) NukeDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NukeDB", reflect.TypeOf((*MockService)(nil).NukeDB))
}

// OpenCreateIncidentDialog mocks base method
func (m *MockService) OpenCreateIncidentDialog(arg0, arg1, arg2, arg3, arg4 string, arg5 []playbook.Playbook, arg6 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenCreateIncidentDialog", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenCreateIncidentDialog indicates an expected call of OpenCreateIncidentDialog
func (mr *MockServiceMockRecorder) OpenCreateIncidentDialog(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenCreateIncidentDialog", reflect.TypeOf((*MockService)(nil).OpenCreateIncidentDialog), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// OpenEndIncidentDialog mocks base method
func (m *MockService) OpenEndIncidentDialog(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenEndIncidentDialog", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenEndIncidentDialog indicates an expected call of OpenEndIncidentDialog
func (mr *MockServiceMockRecorder) OpenEndIncidentDialog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenEndIncidentDialog", reflect.TypeOf((*MockService)(nil).OpenEndIncidentDialog), arg0, arg1)
}

// OpenNextStageDialog mocks base method
func (m *MockService) OpenNextStageDialog(arg0 string, arg1 int, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenNextStageDialog", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenNextStageDialog indicates an expected call of OpenNextStageDialog
func (mr *MockServiceMockRecorder) OpenNextStageDialog(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenNextStageDialog", reflect.TypeOf((*MockService)(nil).OpenNextStageDialog), arg0, arg1, arg2)
}

// OpenUpdateStatusDialog mocks base method
func (m *MockService) OpenUpdateStatusDialog(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUpdateStatusDialog", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenUpdateStatusDialog indicates an expected call of OpenUpdateStatusDialog
func (mr *MockServiceMockRecorder) OpenUpdateStatusDialog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUpdateStatusDialog", reflect.TypeOf((*MockService)(nil).OpenUpdateStatusDialog), arg0, arg1)
}

// RemoveChecklistItem mocks base method
func (m *MockService) RemoveChecklistItem(arg0, arg1 string, arg2, arg3 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveChecklistItem", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveChecklistItem indicates an expected call of RemoveChecklistItem
func (mr *MockServiceMockRecorder) RemoveChecklistItem(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveChecklistItem", reflect.TypeOf((*MockService)(nil).RemoveChecklistItem), arg0, arg1, arg2, arg3)
}

// RemoveReminder mocks base method
func (m *MockService) RemoveReminder(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveReminder", arg0)
}

// RemoveReminder indicates an expected call of RemoveReminder
func (mr *MockServiceMockRecorder) RemoveReminder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveReminder", reflect.TypeOf((*MockService)(nil).RemoveReminder), arg0)
}

// RemoveReminderPost mocks base method
func (m *MockService) RemoveReminderPost(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveReminderPost", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveReminderPost indicates an expected call of RemoveReminderPost
func (mr *MockServiceMockRecorder) RemoveReminderPost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveReminderPost", reflect.TypeOf((*MockService)(nil).RemoveReminderPost), arg0)
}

// RenameChecklistItem mocks base method
func (m *MockService) RenameChecklistItem(arg0, arg1 string, arg2, arg3 int, arg4, arg5 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameChecklistItem", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameChecklistItem indicates an expected call of RenameChecklistItem
func (mr *MockServiceMockRecorder) RenameChecklistItem(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameChecklistItem", reflect.TypeOf((*MockService)(nil).RenameChecklistItem), arg0, arg1, arg2, arg3, arg4, arg5)
}

// RestartIncident mocks base method
func (m *MockService) RestartIncident(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestartIncident", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestartIncident indicates an expected call of RestartIncident
func (mr *MockServiceMockRecorder) RestartIncident(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestartIncident", reflect.TypeOf((*MockService)(nil).RestartIncident), arg0, arg1)
}

// RunChecklistItemSlashCommand mocks base method
func (m *MockService) RunChecklistItemSlashCommand(arg0, arg1 string, arg2, arg3 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunChecklistItemSlashCommand", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunChecklistItemSlashCommand indicates an expected call of RunChecklistItemSlashCommand
func (mr *MockServiceMockRecorder) RunChecklistItemSlashCommand(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunChecklistItemSlashCommand", reflect.TypeOf((*MockService)(nil).RunChecklistItemSlashCommand), arg0, arg1, arg2, arg3)
}

// SetAssignee mocks base method
func (m *MockService) SetAssignee(arg0, arg1, arg2 string, arg3, arg4 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAssignee", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAssignee indicates an expected call of SetAssignee
func (mr *MockServiceMockRecorder) SetAssignee(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAssignee", reflect.TypeOf((*MockService)(nil).SetAssignee), arg0, arg1, arg2, arg3, arg4)
}

// SetReminder mocks base method
func (m *MockService) SetReminder(arg0 string, arg1 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReminder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReminder indicates an expected call of SetReminder
func (mr *MockServiceMockRecorder) SetReminder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReminder", reflect.TypeOf((*MockService)(nil).SetReminder), arg0, arg1)
}

// ToggleCheckedState mocks base method
func (m *MockService) ToggleCheckedState(arg0, arg1 string, arg2, arg3 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToggleCheckedState", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ToggleCheckedState indicates an expected call of ToggleCheckedState
func (mr *MockServiceMockRecorder) ToggleCheckedState(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToggleCheckedState", reflect.TypeOf((*MockService)(nil).ToggleCheckedState), arg0, arg1, arg2, arg3)
}

// UpdateStatus mocks base method
func (m *MockService) UpdateStatus(arg0, arg1 string, arg2 incident.StatusUpdateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockServiceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockService)(nil).UpdateStatus), arg0, arg1, arg2)
}
