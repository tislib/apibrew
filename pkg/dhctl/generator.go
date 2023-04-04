package dhctl

import (
	"github.com/gosimple/slug"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/pkg/generator/golang"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
	"os"
)

var generatorCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate - Generate codes",
	Run: func(cmd *cobra.Command, args []string) {
		parseRootFlags(cmd)

		namespace, err := cmd.Flags().GetString("namespace")
		check(err)

		path, err := cmd.Flags().GetString("path")
		check(err)

		pkg, err := cmd.Flags().GetString("package")
		check(err)

		resp, err := GetDhClient().GetResourceClient().List(cmd.Context(), &stub.ListResourceRequest{
			Token: GetDhClient().GetToken(),
		})

		check(err)

		var filteredResources []*model.Resource

		if len(args) == 0 {
			filteredResources = resp.Resources
		} else {
			for _, resource := range resp.Resources {
				if contains(args, resource.Name) {
					filteredResources = append(filteredResources, resource)
				}
			}
		}

		if pkg == "" {
			pkg = "model"
		}

		for _, resource := range filteredResources {
			if namespace != resource.Namespace {
				continue
			}

			code, err := golang.GenerateGoResourceCode(resource, golang.GenerateResourceCodeParams{
				Package:   pkg,
				Resources: resp.Resources,
			})

			if err != nil {
				log.Fatal(err)
			}

			resourceFileName := slug.Make(resource.Namespace) + "-" + slug.Make(resource.Name) + ".go"

			if resource.Namespace == "default" {
				resourceFileName = slug.Make(resource.Name) + ".go"
			}

			if err := os.MkdirAll(path, 0777); err != nil {
				log.Fatal(err)
			}

			err = os.WriteFile(path+"/"+resourceFileName, []byte(code), 0777)

			check(err)
		}
	},
}

func init() {
	generatorCmd.PersistentFlags().StringP("namespace", "n", "default", "Namespace")
	generatorCmd.PersistentFlags().StringP("path", "p", "", "Path")
	generatorCmd.PersistentFlags().String("package", "", "Package")
}
