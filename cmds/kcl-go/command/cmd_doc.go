// Copyright 2021 The KCL Authors. All rights reserved.

package command

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"kcl-lang.io/kcl-go/pkg/tools/gen"
)

const version = "v0.0.1"

func NewDocCmd() *cli.Command {
	return &cli.Command{
		Hidden: false,
		Name:   "doc",
		Usage:  "show documentation for package or symbol",
		UsageText: `# Generate document for current package
kcl-go doc generate

# Start a local KCL document server
kcl-go doc start`,
		Subcommands: []*cli.Command{
			{
				Name:  "generate",
				Usage: "generates documents from code and examples",
				UsageText: `# Generate Markdown document for current package
kcl-go doc generate

# Generate Html document for current package
kcl-go doc generate --format html

# Generate Markdown document for specific package
kcl-go doc generate --file-path <package path>

# Generate Markdown document for specific package to a <target directory>
kcl-go doc generate --file-path <package path> --target <target directory>`,
				Flags: []cli.Flag{
					// todo: look for packages recursive
					// todo: package path list
					&cli.StringFlag{
						Name: "file-path",
						Usage: `Relative or absolute path to the KCL package root when running kcl-doc command from
	outside of the KCL package root directory.
	If not specified, the current work directory will be used as the KCL package root.`,
					},
					&cli.BoolFlag{
						Name:  "ignore-deprecated",
						Usage: "do not generate documentation for deprecated schemas",
						Value: false,
					},
					&cli.StringFlag{
						Name:  "format",
						Usage: "The document format to generate. Supported values: markdown, html, openapi",
						Value: string(gen.Markdown),
					},
					&cli.StringFlag{
						Name:  "target",
						Usage: "If not specified, the current work directory will be used. A docs/ folder will be created under the target directory",
					},
					&cli.BoolFlag{
						Name:  "escape-html",
						Usage: "whether to escape html symbols when the output format is markdown. Always scape when the output format is html. Default to false",
					},
					&cli.StringFlag{
						Name:  "template",
						Usage: "The template directory based on the KCL package root. If not specified, the built-in templates will be used.",
					},
				},
				Action: func(context *cli.Context) error {
					opts := gen.GenOpts{
						Path:             context.String("file-path"),
						IgnoreDeprecated: context.Bool("ignore-deprecated"),
						Format:           context.String("format"),
						Target:           context.String("target"),
						EscapeHtml:       context.Bool("escape-html"),
						TemplateDir:      context.String("template"),
					}

					genContext, err := opts.ValidateComplete()
					if err != nil {
						fmt.Println(fmt.Errorf("generate failed: %s", err))
						return err
					}

					err = genContext.GenDoc()
					if err != nil {
						fmt.Println(fmt.Errorf("generate failed: %s", err))
						return err
					} else {
						fmt.Println(fmt.Sprintf("Generate Complete! Check generated docs in %s", genContext.Target))
						return nil
					}
				},
			},
			{
				Name:  "start",
				Usage: "starts a document website locally",
				Action: func(context *cli.Context) error {
					fmt.Println("not implemented")
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: "version",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				cli.ShowSubcommandHelpAndExit(c, 1)
				return nil
			}
			arg := c.Args().First()
			if arg == "version" {
				fmt.Println(version)
			}
			return nil
		},
	}
}
