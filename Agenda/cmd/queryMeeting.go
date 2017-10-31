// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
    "Agenda/ops"
    "github.com/spf13/cobra"
)

// queryMeetingCmd represents the queryMeeting command
var queryMeetingCmd = &cobra.Command{
    Use:   "queryMeeting -s [StartDate] -e [EndDate]",
    Short: "To query the meeting during [StartDate] and [EndDate]",
    Long: `usage of using this command is to query the meeting during [StartDate] and [EndDate]`,
    Run: func(cmd *cobra.Command, args []string) {
        startDate, _ := cmd.Flags().GetString("StartDate")
        endDate, _ := cmd.Flags().GetString("EndDate")
        ops.Query_meeting(startDate, endDate)
    },
}

func init() {
    RootCmd.AddCommand(queryMeetingCmd)
    queryMeetingCmd.Flags().StringP("StartDate", "s", "Anonymous", "startDate of the meeting")
    queryMeetingCmd.Flags().StringP("EndDate", "e", "default_endtime", "endDate of the meeting")
}
