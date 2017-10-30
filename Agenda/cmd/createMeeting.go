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
    "github.com/spf13/cobra"
    "Agenda/ops"
)

var createMeetingCmd = &cobra.Command{
    Use:   "createMeeting -t [Title] -p [Participator] -s [StartTime] -e [EndTime]",
    Short: "create a new meeting",
    Long: `usage of using this command is to create a new meeting`,
    Run: func(cmd *cobra.Command, args []string) {
        title, _ := cmd.Flags().GetString("Title")
        participators, _ := cmd.Flags().GetStringSlice("Participators")
        startDate, _ := cmd.Flags().GetString("StartDate")
        endDate, _ := cmd.Flags().GetString("EndDate")
        ops.Create_meeting(title, participators, startDate, endDate)
    },
}

func init() {
    RootCmd.AddCommand(createMeetingCmd)
    createMeetingCmd.Flags().StringP("Title", "t", "Anonymous", "title of the meeting")
    createMeetingCmd.Flags().StringSliceP("Participators", "p", []string{}, "participators of the meeting")
    createMeetingCmd.Flags().StringP("StartDate", "s", "Anonymous", "startDate of the meeting")
    createMeetingCmd.Flags().StringP("EndDate", "e", "Anonymous", "endDate of the meeting")
}
