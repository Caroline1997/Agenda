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
    "log"
    "os"
)

// queryMeetingCmd represents the queryMeeting command
var queryMeetingCmd = &cobra.Command{
    Use:   "queryMeeting -s [StartDate] -e [EndDate]",
    Short: "To query the meeting during [StartDate] and [EndDate]",
    Long: `usage of using this command is to query the meeting during [StartDate] and [EndDate]`,
    Run: func(cmd *cobra.Command, args []string) {
        file,err := os.OpenFile("testquery.log",os.O_WRONLY|os.O_APPEND, 0666)
        if err != nil {
             log.Fatalln("Fail to create testquery.log file!")
        }
        querylog := log.New(file,"",log.LstdFlags|log.Llongfile)
        startDate, _ := cmd.Flags().GetString("StartDate")
        querylog.Printf("[queryMeeting] 1.Get the meeting's StartDate.\n")
        endDate, _ := cmd.Flags().GetString("EndDate")
        querylog.Printf("[queryMeeting] 2.Get the meeting's EndDate.\n")
        ops.Query_meeting(startDate, endDate)
    },
}

func init() {
    RootCmd.AddCommand(queryMeetingCmd)
    queryMeetingCmd.Flags().StringP("StartDate", "s", "Anonymous", "startDate of the meeting")
    queryMeetingCmd.Flags().StringP("EndDate", "e", "default_endtime", "endDate of the meeting")
}
