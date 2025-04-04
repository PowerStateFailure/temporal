// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package isactivitytaskvalid

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/api/serviceerror"
	persistencespb "go.temporal.io/server/api/persistence/v1"
	"go.temporal.io/server/common"
	"go.temporal.io/server/service/history/api"
	historyi "go.temporal.io/server/service/history/interfaces"
	"go.uber.org/mock/gomock"
)

type (
	apiSuite struct {
		suite.Suite
		*require.Assertions

		controller      *gomock.Controller
		workflowLease   api.WorkflowLease
		workflowContext *historyi.MockWorkflowContext
		mutableState    *historyi.MockMutableState
	}
)

func TestAPISuite(t *testing.T) {
	s := new(apiSuite)
	suite.Run(t, s)
}

func (s *apiSuite) SetupTest() {
	s.Assertions = require.New(s.T())

	s.controller = gomock.NewController(s.T())
	s.workflowContext = historyi.NewMockWorkflowContext(s.controller)
	s.mutableState = historyi.NewMockMutableState(s.controller)
	s.workflowLease = api.NewWorkflowLease(
		s.workflowContext,
		func(err error) {},
		s.mutableState,
	)
}

func (s *apiSuite) TeardownTest() {
	s.controller.Finish()
}

func (s *apiSuite) TestWorkflowCompleted() {
	s.mutableState.EXPECT().IsWorkflowExecutionRunning().Return(false)

	_, err := isActivityTaskValid(s.workflowLease, rand.Int63())
	s.Error(err)
	s.IsType(&serviceerror.NotFound{}, err)
}

func (s *apiSuite) TestWorkflowRunning_ActivityTaskNotStarted() {
	s.mutableState.EXPECT().IsWorkflowExecutionRunning().Return(true)
	activityScheduleEventID := rand.Int63()
	s.mutableState.EXPECT().GetActivityInfo(activityScheduleEventID).Return(&persistencespb.ActivityInfo{
		ScheduledEventId: activityScheduleEventID,
		StartedEventId:   common.EmptyEventID,
	}, true)

	valid, err := isActivityTaskValid(s.workflowLease, activityScheduleEventID)
	s.NoError(err)
	s.True(valid)
}

func (s *apiSuite) TestWorkflowRunning_ActivityTaskStarted() {
	s.mutableState.EXPECT().IsWorkflowExecutionRunning().Return(true)
	activityScheduleEventID := rand.Int63()
	s.mutableState.EXPECT().GetActivityInfo(activityScheduleEventID).Return(&persistencespb.ActivityInfo{
		ScheduledEventId: activityScheduleEventID,
		StartedEventId:   activityScheduleEventID + 1,
	}, true)

	valid, err := isActivityTaskValid(s.workflowLease, activityScheduleEventID)
	s.NoError(err)
	s.False(valid)
}

func (s *apiSuite) TestWorkflowRunning_ActivityTaskMissing() {
	s.mutableState.EXPECT().IsWorkflowExecutionRunning().Return(true)
	activityScheduleEventID := rand.Int63()
	s.mutableState.EXPECT().GetActivityInfo(activityScheduleEventID).Return(nil, false)

	valid, err := isActivityTaskValid(s.workflowLease, activityScheduleEventID)
	s.NoError(err)
	s.False(valid)
}
