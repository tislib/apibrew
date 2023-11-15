package apbr

import (
	"errors"
	"github.com/apibrew/apibrew/pkg/apbr/flags"
	"github.com/apibrew/apibrew/pkg/apbr/output"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "watch - watch resource/record docs: https://apibrew.io/docs/cli#apply",
	RunE: func(cmd *cobra.Command, args []string) error {
		parseRootFlags(cmd)

		f := getFlag(cmd, "format", true)
		o := getFlag(cmd, "output", false)
		forApply, err := cmd.Flags().GetBool("for-apply")

		if err != nil {
			return err
		}

		filters, err := cmd.Flags().GetStringSlice("filter")

		if err != nil {
			return err
		}

		selectorFlags.Filters = filters

		limit, err := cmd.PersistentFlags().GetInt64("limit")

		if err != nil {
			return err
		}
		selectorFlags.Limit = limit

		packRecords, err := cmd.PersistentFlags().GetBool("pack-records")

		if err != nil {
			return err
		}
		selectorFlags.PackRecords = packRecords

		offset, err := cmd.PersistentFlags().GetInt64("offset")

		if err != nil {
			return err
		}
		selectorFlags.Offset = offset

		if err != nil {
			return err
		}

		var selection = &flags.SelectedRecordsResult{}

		err = selectorFlags.Parse(selection, cmd, args)

		if err != nil {
			return err
		}

		var w io.Writer
		if o == "" || o == "-" {
			w = os.Stdout
		} else {
			var err error
			var wf *os.File
			wf, err = os.Create(o)

			if err != nil {
				log.Fatal(err)
			}

			w = wf

			defer func() {
				err = wf.Close()

				if err != nil {
					log.Fatal(err)
				}
			}()
		}

		writer := output.NewOutputWriter(f, w, map[string]string{
			"for-apply": boolToString(forApply),
		})

		if writer.IsBinary() && o == "" {
			return errors.New("format is binary but output is not specified")
		}

		// not implemented

		return nil
	},
}

func initWatchCmd() {
	watchCmd.PersistentFlags().StringP("format", "f", "console", "format")
	watchCmd.PersistentFlags().StringP("output", "o", "", "output")
	watchCmd.PersistentFlags().Int64("limit", 100, "limit")
	watchCmd.PersistentFlags().Int64("offset", 0, "offset")
	watchCmd.PersistentFlags().StringSlice("filter", nil, "filter")
	watchCmd.PersistentFlags().Bool("append", false, "append")
	watchCmd.PersistentFlags().Bool("pack-records", false, "pack-records")
	watchCmd.PersistentFlags().Bool("for-apply", false, "Prepare for apply")
	selectorFlags.Declare(watchCmd)
}
